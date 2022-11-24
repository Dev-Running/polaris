package kafka

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func RunKAFKA(a *kafka.Consumer) {
	for {
		m, err := a.ReadMessage(100 * time.Millisecond)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.TopicPartition.Offset, string(m.Key), string(m.Value))
	}
	a.Close()
}
