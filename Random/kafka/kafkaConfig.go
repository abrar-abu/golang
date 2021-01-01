package main

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
	//"strconv"
	"time"
)

// Below are the default setting to create a partion
var brokers string 			= "localhost:9092" 
var topic string			= "demo"
var numberOfPartion int		= 1
var replicationFactor int	= 1


func createKafkaConfig(brokers string) (*kafka.AdminClient, error) {
	// config := &kafka.AdminClient{}
	config, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		// fmt.Printf("Failed to create Admin client: %s\n", err)
		// os.Exit(1)
		return nil, errors.New("Failed to create Admin client")
	}
	return config, nil
}


func createKafkaTopic(maxDur time.Duration, topic string, numberOfPartion int, replicationFactor int) ([]kafka.TopicResult, error) {
	config, err := createKafkaConfig(brokers)
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	results, err := config.CreateTopics(
		ctx,
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     numberOfPartion,
			ReplicationFactor: replicationFactor}},
		// Admin options
		kafka.SetAdminOperationTimeout(maxDur))
		if err != nil {
			// fmt.Printf("Failed to create topic: %v\n", err)
			// os.Exit(1)
			fmt.Printf("Failed to create topic: %v\n", err)
			return nil ,errors.New("Failed to create topic")
		}
	return results, nil
}

/*
func listAllTheTopics() {
	config, err := createKafkaConfig(brokers)
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
}
*/
func deleteKafkaTopic(topics string,maxDur time.Duration) ([]kafka.TopicResult, error) {
	config, err := createKafkaConfig(brokers)
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	results, err := config.DeleteTopics(ctx, []string{topics}, kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to delete topics: %v\n", err)
		//os.Exit(1)
		return nil ,errors.New("Failed to delete topic")
	}
	return results, nil
}


func main() {

	config, err := createKafkaConfig(brokers)
	// config, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("config 				- Type: %T Value: %v\n", config, config)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Printf("ctx 				- Type: %T Value: %v\n", ctx, ctx)
	
	// Create topics on cluster.
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}
	fmt.Printf("maxDur 				- Type: %T Value: %v\n", maxDur, maxDur)

	results, err := createKafkaTopic(maxDur,topic,numberOfPartion,replicationFactor)
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		os.Exit(1)
	}
	/*
	results, err := config.CreateTopics(
		ctx,
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     numberOfPartion,
			ReplicationFactor: replicationFactor}},
		// Admin options
		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		os.Exit(1)
	}
	*/
	fmt.Printf("results 			- Type: %T Value: %v\n", results, results)
	deleteResults, err := deleteKafkaTopic(topic,maxDur)
	if err != nil {
		fmt.Printf("Failed to delete topics: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("deleteResults 			- Type: %T Value: %v\n", deleteResults, deleteResults)
}