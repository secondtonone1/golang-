package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/Shopify/sarama"
)

func main() {
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
		panic(err)
	}

	defer producer.Close()

	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		//Topic: "test",//包含了消息的主题
		Partition: int32(10),                   //
		Key:       sarama.StringEncoder("key"), //
	}

	

	inputReader := bufio.NewReader(os.Stdin)
	for{
		value, _ , err := inputReader.ReadLine()
    	if err != nil {
        	fmt.Printf("error:", err.Error())
        	return
    	}
		msgType , _, err  := inputReader.ReadLine()
		msg.Topic = string(msgType)
		fmt.Println("topic is : ",msg.Topic)
		fmt.Println("value is : ",string(value))
		msg.Value = sarama.ByteEncoder(value)
        partition, offset, err := producer.SendMessage(msg)

        if err != nil {
			fmt.Println("Send message Fail")
			fmt.Println(err.Error())
        }
        fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	}
}
