package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wbtech-l0/internal/domain"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "wbtech-topic",
		GroupID:  "wbtech-l0",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	fmt.Println("start consuming ... !!")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigchan
		fmt.Printf("Received signal %v: shutting down...\n", sig)

		if err := reader.Close(); err != nil {
			fmt.Println("Error closing Kafka reader:", err)
		}

		os.Exit(0)
	}()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		message := &domain.Order{}
		err = json.Unmarshal(msg.Value, message)

		// TODO: реализовать функционал записи в БД. Разобраться, как это устроено в gorm или в стандартной библиотеке.
		// TODO: дополнительно: изучить, как управлять подтверждением обработки сообщений (на случай сбоя при записи в БД).

		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Received a message: ", message.Payment.PaymentDt)
		}
	}
}
