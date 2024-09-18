package kafka

import (
	"awesome-chat/gateway/internal/config"
	"awesome-chat/gateway/internal/model"
	"context"
	"encoding/json"
)
import "github.com/segmentio/kafka-go"

type Kafka struct {
	cfg  config.KafkaConfig
	conn *kafka.Conn
}

func NewKafka(ctx context.Context, cfg config.KafkaConfig) (*Kafka, error) {
	conn, err := kafka.DialLeader(
		ctx,
		"tcp",
		cfg.GetAddress(),
		cfg.GetTopic(),
		cfg.GetPartition())

	if err != nil {
		return nil, err
	}

	return &Kafka{
		cfg:  cfg,
		conn: conn,
	}, nil
}

func (k *Kafka) Write(message *model.Message) error {
	// TODO [19.09.2024] поменять маршал на gob или что-то другое
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// TODO [19.09.2024] добавить батч запись
	_, err = k.conn.WriteMessages(
		kafka.Message{Value: b},
	)
	return err
}
