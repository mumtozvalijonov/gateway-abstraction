package server

import (
	"context"
	"example/microabstraction/internal/service"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaServer struct {
	sumService *service.SumService
}

func (s *KafkaServer) Run() {
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	log.Println("connected to kafka")

	for {
		batch := conn.ReadBatch(0, 1e6) // fetch 10KB min, 1MB max

		b := make([]byte, 10e3) // 10KB max per message
		for {
			n, err := batch.Read(b)
			if err != nil {
				break
			}
			fmt.Println(string(b[:n]))
		}

		if err := batch.Close(); err != nil {
			log.Fatal("failed to close batch:", err)
			break
		}
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func NewKafkaServer(service *service.SumService) *KafkaServer {
	return &KafkaServer{sumService: service}
}
