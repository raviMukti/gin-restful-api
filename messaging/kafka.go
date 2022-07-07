package messaging

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
	"github.com/raviMukti/gin-restful-api/helper"
)

var (
	Brokers []string
	Topic   string
)

func InitKafkaConfig() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	Brokers = []string{os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT")}
	Topic = os.Getenv("KAFKA_TOPIC")
}

func newKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V2_8_0_0
	return conf
}

func NewKafkaSyncProducer() sarama.SyncProducer {
	kafka, err := sarama.NewSyncProducer(Brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("KAFKA ERROR : %s\n", err)
		os.Exit(-1)
	}

	return kafka
}

func NewKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(Brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("KAFKA ERROR: %s\n", err)
		os.Exit(-1)
	}

	return consumer
}

func SendMsg(kafka sarama.SyncProducer, event interface{}) error {
	json, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msgLog := &sarama.ProducerMessage{
		Topic: Topic,
		Value: sarama.StringEncoder(string(json)),
	}

	partition, offset, err := kafka.SendMessage(msgLog)
	if err != nil {
		fmt.Printf("KAFKA ERROR: %s\n", err)
	}

	fmt.Printf("MESSAGE: %+v\n", event)
	fmt.Printf("MESSAGE IS STORED IN PARTITION %d, OFFSET %d\n",
		partition, offset)

	return nil
}
