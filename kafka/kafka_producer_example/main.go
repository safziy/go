package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors =true
	config.Version = sarama.V3_1_0_0

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9093"}, config)
	if err != nil {
		fmt.Printf("producer create producer error:%v", err)
		return
	}
	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic:     "kafka_go_test",
		Key:       sarama.StringEncoder("go_test"),
	}

	value := "this is message"
	for {
		fmt.Scanln(&value)
		msg.Value = sarama.ByteEncoder(value)
		fmt.Printf("input [%s]\n", value)

		// send to chain
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d, timestamp: %s\n", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %v\n", fail.Err)
		}
	}
}

func main()  {
	producer()
}
