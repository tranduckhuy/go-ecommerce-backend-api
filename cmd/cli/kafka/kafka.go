package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19094" // Port mapped in docker-compose for external access
	kafkaTopic = "test-topic"
)

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(message, stockType string) *StockInfo {
	return &StockInfo{
		Message: message,
		Type:    stockType,
	}
}

func actionStock(c *gin.Context) {

	s := newStock(c.Query("message"), c.Query("type"))
	body := make(map[string]any)
	body["action"] = "action"
	body["info"] = s

	jsonBody, err := json.Marshal(body)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err = kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to write message to Kafka"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "Message sent to Kafka"})
}

func RegisterConsumerATC(id int) {

	kafkaGroupId := "group-id-" + strconv.Itoa(id)

	// Initialize the Kafka reader
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer %d started to read from topic %s\n", id, kafkaTopic)

	// Start consuming messages
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer %d error: %v\n", id, err)
		}

		fmt.Printf("Consumer %d, received message: %s, topic: %v, partition: %d, offset: %d, time: %s\n",
			id, string(m.Value), m.Topic, m.Partition, m.Offset, m.Time)
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("/action/stock", actionStock)

	// Register multiple consumers
	for i := 1; i <= 3; i++ {
		go RegisterConsumerATC(i)
	}

	r.Run(":8080")
}
