package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func consumer()  {
	fmt.Println("consumer")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V3_1_0_0

	// consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9093"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("kafka_go_test", 0, sarama.OffsetOldest)
	//partitionConsumer, err := consumer.ConsumePartition("kafka_go_test", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}

func main() {
	consumer()
}
