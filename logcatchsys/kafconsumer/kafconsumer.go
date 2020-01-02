package kafconsumer

import (
	"context"
	"fmt"
	"golang-/logcatchsys/logconfig"
	"sync"

	"github.com/Shopify/sarama"
)

type TopicData struct {
	Topic       string
	Partition   int32
	KafConsumer sarama.Consumer
	Ctx         context.Context
	Cancel      context.CancelFunc
}

var topicMap map[string]map[int32]*TopicData
var topicSet map[string]bool

func init() {
	topicMap = make(map[string]map[int32]*TopicData)
	topicSet = make(map[string]bool)
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
	//wg := sync.WaitGroup{}
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
	defer consumer.Close()
	topicSet = ConstructTopicSet()
	ConsumeTopic(consumer)
}

func ConsumeTopic(consumer sarama.Consumer) {
	wg := sync.WaitGroup{}
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
			wg.Add(1)
			go func(sarama.PartitionConsumer) {
				defer wg.Done()
				//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
				for msg := range pc.Messages() {
					fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n",
						msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				}
			}(pc)
		}
	}
	wg.Wait()
}
