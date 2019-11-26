package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func main(){
	fmt.Println("consumer begin...")
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	wg  :=sync.WaitGroup{}
	//创建消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"},config)
	if err != nil {
		fmt.Println("consumer create failed, error is ", err.Error())
		return
	}
	defer consumer.Close()
	
	//Partitions(topic):该方法返回了该topic的所有分区id
    partitionList, err := consumer.Partitions("test")
    if err != nil {
		fmt.Println("get consumer partitions failed")
		fmt.Println("error is ", err.Error())
		return
    }

	for partition := range partitionList {
		//ConsumePartition方法根据主题，
		//分区和给定的偏移量创建创建了相应的分区消费者
		//如果该分区消费者已经消费了该信息将会返回error
		//OffsetNewest消费最新数据
        pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
        if err != nil {
            panic(err)
		}
		//异步关闭，保证数据落盘
        defer pc.AsyncClose()
        wg.Add(1)
        go func(sarama.PartitionConsumer) {
            defer wg.Done()
            //Messages()该方法返回一个消费消息类型的只读通道，由代理产生
            for msg := range pc.Messages() {
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", 
				msg.Topic,msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
            }
        }(pc)
    }
    wg.Wait()
    consumer.Close()
	
}

