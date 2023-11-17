package main

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

func CreateReader() *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:29092"},
		Topic:     "notifications",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	return r
}

func main() {
	fmt.Println("Starting consumer")
	reader := CreateReader()
	fmt.Println("Created reader")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
