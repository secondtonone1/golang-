package kafconsumer

import (
	"context"
	"fmt"
	"golang-/logcatchsys/etcdconsumer"
	"golang-/logcatchsys/logconfig"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/olivere/elastic"
	"go.etcd.io/etcd/clientv3"
)

var consumer_once sync.Once

type TopicPart struct {
	Topic     string
	Partition int32
}

type LogData struct {
	Topic string
	Log   string
	Id    string
}

type TopicData struct {
	TPartition  *TopicPart
	KafConsumer sarama.PartitionConsumer
	Ctx         context.Context
	Cancel      context.CancelFunc
}

var topicMap map[string]map[int32]*TopicData
var topicSet map[string]bool
var etcd_topicSet map[string]bool
var etcd_topicMap map[string]map[int32]*TopicData
var readExitOnce sync.Once
var topicChan chan *TopicPart
var etcd_topicChan chan *TopicPart
var consumer_list []sarama.Consumer
var etcdcli *clientv3.Client

func init() {
	topicMap = make(map[string]map[int32]*TopicData)
	etcd_topicMap = make(map[string]map[int32]*TopicData)
	topicSet = make(map[string]bool)
	etcd_topicSet = make(map[string]bool)
	topicChan = make(chan *TopicPart, 20)
	etcd_topicChan = make(chan *TopicPart, 20)
	consumer_list = make([]sarama.Consumer, 0, 20)
}

func ConstructTopicSet() map[string]bool {
	topicSetTmp := make(map[string]bool)
	configtopics, _ := logconfig.ReadConfig(logconfig.InitVipper(), "collectlogs")
	if configtopics == nil {
		goto CONFTOPIC
	}
	for _, configtopic := range configtopics.([]interface{}) {
		confmap := configtopic.(map[interface{}]interface{})
		for key, val := range confmap {
			if key.(string) == "logtopic" {
				topicSetTmp[val.(string)] = true
			}
		}
	}
CONFTOPIC:
	return topicSetTmp
}

func initConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	var kafkaddr = "localhost:9092"
	kafkaconf, _ := logconfig.ReadConfig(logconfig.InitVipper(), "kafkaconfig.kafkaaddr")
	if kafkaconf != nil {
		kafkaddr = kafkaconf.(string)
	}
	//创建消费者
	consumer, err := sarama.NewConsumer([]string{kafkaddr}, config)
	return consumer, err
}

func GetMsgFromKafka() {
	fmt.Println("kafka consumer begin ...")
	consumer, err := initConsumer()
	if err != nil {
		fmt.Println("consumer create failed, error is ", err.Error())
		return
	}
	consumer_list = append(consumer_list, consumer)
	etcdcli, err = etcdconsumer.InitEtcdClient()
	if err != nil {
		fmt.Println("etcd client init failed!")
		return
	}

	etcd_topictmp, err := etcdconsumer.GetTopicSet(etcdcli)
	if err != nil {
		fmt.Println("etcd topic set get error")
		return
	}

	fmt.Println(etcd_topictmp)
	etcd_topicSet = etcd_topictmp.(map[string]bool)
	defer func(consumerlist []sarama.Consumer) {
		if err := recover(); err != nil {
			fmt.Println("consumer panic error ", err)
		}
		topicSet = nil
		//回收所有协程
		for _, val := range topicMap {
			for _, valt := range val {
				defer valt.KafConsumer.AsyncClose()
				valt.Cancel()
			}
		}

		topicMap = nil

		for _, val := range etcd_topicMap {
			for _, valt := range val {
				defer valt.KafConsumer.AsyncClose()
				valt.Cancel()
			}
		}
		etcd_topicMap = nil
		for _, consumer := range consumerlist {
			consumer.Close()
		}

	}(consumer_list)
	topicSetTmp := ConstructTopicSet()
	if topicSetTmp == nil {
		fmt.Println("construct topic set error ")
		return
	}
	topicSet = topicSetTmp

	ConsumeTopic(consumer)

}

func ConvertSet2Map(consumer sarama.Consumer, topicSet map[string]bool,
	topicMaps map[string]map[int32]*TopicData, topic_chan chan *TopicPart) {
	for key, _ := range topicSet {
		partitionList, err := consumer.Partitions(key)
		if err != nil {
			fmt.Println("get consumer partitions failed")
			fmt.Println("error is ", err.Error())
			continue
		}

		for partition := range partitionList {
			pc, err := consumer.ConsumePartition(key, int32(partition), sarama.OffsetNewest)
			if err != nil {
				fmt.Println("consume partition error is ", err.Error())
				continue
			}
			//	defer pc.AsyncClose()
			topicData := new(TopicData)
			topicData.Ctx, topicData.Cancel = context.WithCancel(context.Background())
			topicData.KafConsumer = pc
			topicData.TPartition = new(TopicPart)
			topicData.TPartition.Partition = int32(partition)
			topicData.TPartition.Topic = key
			_, okm := topicMaps[key]
			if !okm {
				topicMaps[key] = make(map[int32]*TopicData)
			}
			topicMaps[key][int32(partition)] = topicData
			go PutIntoES(topicData, topic_chan)

		}
	}
}

func ConsumeTopic(consumer sarama.Consumer) {
	ConvertSet2Map(consumer, topicSet, topicMap, topicChan)
	ConvertSet2Map(consumer, etcd_topicSet, etcd_topicMap, etcd_topicChan)
	//监听配置文件
	ctx, cancel := context.WithCancel(context.Background())
	pathChan := make(chan interface{})
	etcdChan := make(chan interface{})
	go logconfig.WatchConfig(ctx, logconfig.InitVipper(), pathChan, etcdChan)
	defer func(cancel context.CancelFunc) {
		consumer_once.Do(func() {
			if err := recover(); err != nil {
				fmt.Println("consumer main goroutine panic, ", err)
			}
			cancel()
		})

	}(cancel)

	for {
		select {
		//检测监控路径的协程崩溃，重启
		case topicpart := <-topicChan:
			fmt.Printf("receive goroutine exited, topic is %s, partition is %d\n",
				topicpart.Topic, topicpart.Partition)
			//重启消费者读取数据的协程
			val, ok := topicMap[topicpart.Topic]
			if !ok {
				continue
			}
			tp, ok := val[topicpart.Partition]
			if !ok {
				continue
			}
			tp.Ctx, tp.Cancel = context.WithCancel(context.Background())
			go PutIntoES(tp, topicChan)
		//检测etcd配置解析后，监控路径的协程崩溃，重启
		case topicpart := <-etcd_topicChan:
			fmt.Printf("receive goroutine exited, topic is %s, partition is %d\n",
				topicpart.Topic, topicpart.Partition)
			//重启消费者读取数据的协程
			val, ok := etcd_topicMap[topicpart.Topic]
			if !ok {
				continue
			}
			tp, ok := val[topicpart.Partition]
			if !ok {
				continue
			}
			tp.Ctx, tp.Cancel = context.WithCancel(context.Background())
			go PutIntoES(tp, etcd_topicChan)
		//检测vipper监控返回配置的更新
		case pathchange, ok := <-pathChan:
			if !ok {
				fmt.Println("vipper watch goroutine exited")
				goto LOOPEND
			}
			//fmt.Println(pathchange)
			topicSetTemp := make(map[string]bool)
			for _, chval := range pathchange.([]interface{}) {
				for logkey, logval := range chval.(map[interface{}]interface{}) {
					if logkey.(string) == "logtopic" {
						topicSetTemp[logval.(string)] = true
					}
				}
			}
			UpdateTopicLogRoutine(topicSetTemp)

			//fmt.Println(topicSetTemp)
		case etcdchange, ok := <-etcdChan:
			if !ok {
				fmt.Println("vipper watch goroutine extied")
				goto LOOPEND
			}
			fmt.Println(etcdchange)
			topicsetTemp, err := etcdconsumer.GetTopicSet(etcdcli)
			if err != nil {
				continue
			}
			UpdateEtcdTopicLogRoutine(topicsetTemp.(map[string]bool))
		}
	}
LOOPEND:
	fmt.Printf("for exited ")
}

func UpdateTopicLogRoutine(newTopicSet map[string]bool) {
	for oldkey, oldval := range topicMap {
		_, ok := newTopicSet[oldkey]
		//旧的key在新的map中不存在
		if !ok {
			//一个topic会有很多partition，所以遍历关闭监控这个topic的所有partition的goroutine
			for _, ele := range oldval {
				ele.KafConsumer.AsyncClose()
				ele.Cancel()
			}
			//将该topic所有分区移除map
			delete(topicMap, oldkey)
			continue
		}
	}

	addkeySet := make(map[string]bool)
	for newkey, _ := range newTopicSet {
		_, ok := topicMap[newkey]
		//旧的map中没有新的key，则做新增处理
		if !ok {
			addkeySet[newkey] = true
			continue
		}
	}
	consumer, err := initConsumer()
	if err != nil {
		fmt.Println("init kafka consumer failed")
		return
	}
	consumer_list = append(consumer_list, consumer)
	//将新增的keyset转化为map，并启动协程
	ConvertSet2Map(consumer, addkeySet, topicMap, topicChan)
}

func UpdateEtcdTopicLogRoutine(newTopicSet map[string]bool) {
	for oldkey, oldval := range etcd_topicMap {
		_, ok := newTopicSet[oldkey]
		//旧的key在新的map中不存在
		if !ok {
			//一个topic会有很多partition，所以遍历关闭监控这个topic的所有partition的goroutine
			for _, ele := range oldval {
				ele.KafConsumer.AsyncClose()
				ele.Cancel()
			}
			//将该topic所有分区移除map
			delete(etcd_topicMap, oldkey)
			continue
		}
	}

	addkeySet := make(map[string]bool)
	for newkey, _ := range newTopicSet {
		_, ok := etcd_topicMap[newkey]
		//旧的map中没有新的key，则做新增处理
		if !ok {
			addkeySet[newkey] = true
			continue
		}
	}
	consumer, err := initConsumer()
	if err != nil {
		fmt.Println("init kafka consumer failed")
		return
	}
	consumer_list = append(consumer_list, consumer)
	//将新增的keyset转化为map，并启动协程
	ConvertSet2Map(consumer, addkeySet, etcd_topicMap, topicChan)
}

//
func PutIntoES(topicData *TopicData, topic_chan chan *TopicPart) {

	fmt.Printf("kafka consumer begin to read message, topic is %s, part is %d\n", topicData.TPartition.Topic,
		topicData.TPartition.Partition)

	logger := log.New(os.Stdout, "LOGCAT", log.LstdFlags|log.Lshortfile)
	elastiaddr, _ := logconfig.ReadConfig(logconfig.InitVipper(), "elasticconfig.elasticaddr")
	if elastiaddr == nil {
		elastiaddr = "localhost:9200"
	}

	esClient, err := elastic.NewClient(elastic.SetURL("http://"+elastiaddr.(string)),
		elastic.SetErrorLog(logger))
	if err != nil {
		// Handle error
		logger.Println("create elestic client error ", err.Error())
		return
	}

	info, code, err := esClient.Ping("http://" + elastiaddr.(string)).Do(context.Background())
	if err != nil {
		logger.Println("elestic search ping error, ", err.Error())
		esClient.Stop()
		esClient = nil
		return
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := esClient.ElasticsearchVersion("http://" + elastiaddr.(string))
	if err != nil {
		fmt.Println("elestic search version get failed, ", err.Error())
		esClient.Stop()
		esClient = nil
		return
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	defer func(esClient *elastic.Client) {
		if err := recover(); err != nil {
			fmt.Printf("consumer message panic %s, topic is %s, part is %d\n", err,
				topicData.TPartition.Topic, topicData.TPartition.Partition)
			topic_chan <- topicData.TPartition
		}

	}(esClient)

	var typestr = "catlog"
	typeconf, _ := logconfig.ReadConfig(logconfig.InitVipper(), "elasticconfig.typestr")
	if typeconf != nil {
		typestr = typeconf.(string)
	}

	for {
		select {
		case msg, ok := <-topicData.KafConsumer.Messages():
			if !ok {
				fmt.Println("etcd message chan closed ")
				return
			}
			fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			idstr := strconv.FormatInt(int64(msg.Partition), 10) + strconv.FormatInt(msg.Offset, 10)
			logdata := &LogData{Topic: msg.Topic, Log: string(msg.Value), Id: idstr}
			createIndex, err := esClient.Index().Index(msg.Topic).Type(typestr).Id(idstr).BodyJson(logdata).Do(context.Background())

			if err != nil {
				logger.Println("create index failed, ", err.Error())
				continue
			}
			fmt.Println("create index success, ", createIndex)

		case <-topicData.Ctx.Done():
			fmt.Println("es goroutine receive exited from parent goroutine !")
			return
		}
	}
}
