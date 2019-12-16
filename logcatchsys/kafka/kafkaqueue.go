package kafkaqueue

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func CreateKafkaProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()

	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("create producer failed, ", err.Error())
		return nil, err
	}
	fmt.Println("create kafka producer success")

	return producer, nil
}

type ProducerKaf struct {
	Producer sarama.SyncProducer
}

func (p *ProducerKaf) PutIntoKafka(keystr string, valstr string) {
	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic: keystr,
		Key:   sarama.StringEncoder(keystr),
		Value: sarama.StringEncoder(valstr),
	}
	partition, offset, err := p.Producer.SendMessage(msg)

	if err != nil {
		fmt.Println("Send message Fail")
		fmt.Println(err.Error())
	}
	fmt.Printf("Partition = %d, offset=%d, msgvalue=%s \n", partition, offset, valstr)

}

func CreateKafkaConsumer() {

}
