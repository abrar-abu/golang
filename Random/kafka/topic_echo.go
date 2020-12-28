/*
https://github.com/sohamkamani/golang-kafka-example/blob/master/main.go
*/
package main

import (
	"flag"
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	var topicName string
	flag.StringVar(&topicName, "t", "movies", "topic")
	flag.Parse()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Subscribe Topic: " + topicName)

	c.SubscribeTopics([]string{topicName, "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: ", msg.TopicPartition)
			fmt.Printf("\t headers-> %v\n", msg.Headers)
			fmt.Printf("\t Key    -> %s\n", string(msg.Key))
			fmt.Printf("\t Value  -> %s\n", string(msg.Value))
			fmt.Printf("\t Timestamp  -> %v\n", msg.Timestamp)
			fmt.Printf("\t TimestampType  -> %v\n", msg.TimestampType)
			fmt.Printf("\t Opaque  -> %v\n", msg.Opaque)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}