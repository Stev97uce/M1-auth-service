package event

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

// Función para enviar eventos a Kafka
func SendEventToKafka(event string) {
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaBroker},
		Topic:   "session-logs",
	})
	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(event),
	})
	if err != nil {
		log.Println("Error sending to Kafka: ", err)
	}
	log.Println("Event sent to Kafka:", event)
}
