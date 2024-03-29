package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/septiyanandika/go-kafka-example/plain/producer"
)

const (
	kafkaHost     = "127.0.0.1:9092"
	kafkaUserName = ""
	kafkaPassword = ""
)

func main() {

	kafkaConfig := initKafkaConfig(kafkaUserName, kafkaPassword)
	producers, err := sarama.NewSyncProducer([]string{kafkaHost}, kafkaConfig)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := producers.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	fmt.Println("Success create kafka sync-producer")

	kafka := &producer.KafkaProducer{
		Producer: producers,
	}

	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("message test number %v", i)
		err := kafka.SendMessage("test-topic", msg)
		if err != nil {
			panic(err)
		}
	}
}

func initKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
