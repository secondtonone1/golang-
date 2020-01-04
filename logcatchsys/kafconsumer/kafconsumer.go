package kafconsumer

import (
	"context"
	"fmt"
	"golang-/logcatchsys/logconfig"
	"sync"

	"github.com/Shopify/sarama"
)

type TopicPart struct {
	Topic     string
	Partition int32
}

type TopicData struct {
	TPartition  *TopicPart
	KafConsumer sarama.PartitionConsumer
	Ctx         context.Context
	Cancel      context.CancelFunc
}

var topicMap map[string]map[int32]*TopicData
var topicSet map[string]bool
var readExitOnce sync.Once
var topicChan chan *TopicPart

func init() {
	topicMap = make(map[string]map[int32]*TopicData)
	topicSet = make(map[string]bool)
	topicChan = make(chan *TopicPart, 20)
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

func GetMsgFromKafka() {
	fmt.Println("kafka consumer begin ...")
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	var kafkaddr = "localhost:9092"
	kafkaconf, _ := logconfig.ReadConfig(logconfig.InitVipper(), "kafkaconfig.kafkaaddr")
	if kafkaconf != nil {
		kafkaddr = kafkaconf.(string)
	}
	//创建消费者
	consumer, err := sarama.NewConsumer([]string{kafkaddr}, config)
	if err != nil {
		fmt.Println("consumer create failed, error is ", err.Error())
		return
	}
	defer func(consumer sarama.Consumer) {
		if err := recover(); err != nil {
			fmt.Println("consumer panic error ", err)
		}
		consumer.Close()
		topicSet = nil
		//回收所有协程
		for _, val := range topicMap {
			for _, valt := range val {
				valt.Cancel()
			}
		}

		topicMap = nil
	}(consumer)
	topicSetTmp := ConstructTopicSet()
	if topicSetTmp == nil {
		fmt.Println("construct topic set error ")
		return
	}
	topicSet = topicSetTmp
	ConsumeTopic(consumer)
}

func ConsumeTopic(consumer sarama.Consumer) {

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
			defer pc.AsyncClose()

			topicData := new(TopicData)
			topicData.Ctx, topicData.Cancel = context.WithCancel(context.Background())
			topicData.KafConsumer = pc
			topicData.TPartition = new(TopicPart)
			topicData.TPartition.Partition = int32(partition)
			topicData.TPartition.Topic = key
			_, okm := topicMap[key]
			if !okm {
				topicMap[key] = make(map[int32]*TopicData)
			}
			topicMap[key][int32(partition)] = topicData
			go ReadFromEtcd(topicData)

		}
	}
	for {
		select {
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
			go ReadFromEtcd(tp)
		}

	}
}

func ReadFromEtcd(topicData *TopicData) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("consumer message panic %s, topic is %s, part is %d\n", err,
				topicData.TPartition.Topic, topicData.TPartition.Partition)
			topicChan <- topicData.TPartition
		}
	}()
	fmt.Printf("kafka consumer begin to read message, topic is %s, part is %d\n", topicData.TPartition.Topic,
		topicData.TPartition.Partition)
	panic("test panic")
	for {
		select {
		case msg, ok := <-topicData.KafConsumer.Messages():
			if !ok {
				fmt.Println("etcd message chan closed ")
				return
			}
			fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		case <-topicData.Ctx.Done():
			fmt.Println("receive exited from parent goroutine !")
			return
		}
	}
}
