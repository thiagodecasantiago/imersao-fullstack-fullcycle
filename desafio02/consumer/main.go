package main

import (
  "runtime"
  "path/filepath"
  "github.com/joho/godotenv"
  "log"
  "os"
  "fmt"
  ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
  DeliveryChan chan ckafka.Event
}

func main()  {
  err := loadEnv()
  if err != nil {
    log.Fatalf("Error loading .env files")
  }
 
  deliveryChan := make(chan ckafka.Event)
  consumer := NewKafkaConsumer(deliveryChan)

  consumer.Consume()
}

func loadEnv() error{
  _, b, _, _ := runtime.Caller(0)
  basepath := filepath.Dir(b)

  err := godotenv.Load(basepath + "/../.env")
  return err
}

func NewKafkaConsumer(deliveryChan chan ckafka.Event) *KafkaConsumer {
  return &KafkaConsumer{
    DeliveryChan: deliveryChan,
  }
}

func (k *KafkaConsumer) Consume() {
  configMap := &ckafka.ConfigMap{
    "bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
    "group.id":          os.Getenv("kafkaConsumerGroupId"),
    "auto.offset.reset": "earliest",
  }
  c, err := ckafka.NewConsumer(configMap)
  defer c.Close()

  if err != nil {
    panic(err)
  }

  topics := []string{os.Getenv("kafkaTopic")}
  c.SubscribeTopics(topics, nil)

  fmt.Println("kafka consumer has been started")
  for {
    msg, err := c.ReadMessage(-1)
    if err == nil {
      k.processMessage(msg)
    } else {
      // The client will automatically try to recover from all errors.
      fmt.Printf("Consumer error: %v (%v)\n", err, msg)
    }
  }
}

func (k *KafkaConsumer) processMessage(msg *ckafka.Message) {
  fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
}
