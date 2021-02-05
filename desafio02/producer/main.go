package main

import (
  "runtime"
  "path/filepath"
  "github.com/joho/godotenv"
  "log"
  "os"
  "fmt"
  "errors"
  ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main()  {
  err := loadEnv()
  if err != nil {
    log.Fatalf("Error loading .env files")
  }
  
  deliveryChan := make(chan ckafka.Event)
  producer := NewKafkaProducer();
  defer producer.Close()
  go DeliveryReport(deliveryChan)

  msg, err := getMessageFromCLI()
  if err != nil {
    fmt.Println("Error reading message:", err)
    fmt.Println("To publish a message execute 'go run producer/main.go \"some message\"'")
    return
  }

  err = Publish(msg, os.Getenv("kafkaTopic"), producer, deliveryChan)
  if err != nil {
    fmt.Println("Error publishing message", err)
  }
  
  producer.Flush(15 * 1000)
}

func loadEnv() error{
  _, b, _, _ := runtime.Caller(0)
  basepath := filepath.Dir(b)

  err := godotenv.Load(basepath + "/../.env")
  return err
}

func getMessageFromCLI() (string, error) {
  var err error
  if len(os.Args) < 2 {
    err = errors.New("Empty message")
    return "", err
  }
  msg := os.Args[1]
  if len(msg) == 0 {
    err = errors.New("Empty message")
    return "", err
  }
  return msg, nil
}

func NewKafkaProducer() *ckafka.Producer {
  configMap := &ckafka.ConfigMap{
    "bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
  }
  p, err := ckafka.NewProducer(configMap)
  if err != nil {
    panic(err)
  }
  return p
}

func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
  message := &ckafka.Message{
    TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
    Value:          []byte(msg),
  }
  err := producer.Produce(message, deliveryChan)
  if err != nil {
    return err
  }
  return nil
}

func DeliveryReport(deliveryChan chan ckafka.Event) {
  for e := range deliveryChan {
    switch ev := e.(type) {
    case *ckafka.Message:
      if ev.TopicPartition.Error != nil {
        fmt.Println("Delivery failed:", ev.TopicPartition)
      } else {
        fmt.Println("Delivered message to:", ev.TopicPartition)
      }
    }
  }
}
