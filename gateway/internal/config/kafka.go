package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type kafkaConfig struct {
	topic     string
	partition int
	Host      string `env:"KAFKA_HOST"`
	Port      string `env:"KAFKA_PORT"`
}

func NewKafkaConfig(topic string, partition int) (KafkaConfig, error) {
	var cfg kafkaConfig

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		println(err.Error())
	}
	cfg.topic = topic
	cfg.partition = partition
	log.Printf("%+v\n", cfg)
	return cfg, err
}

func (cfg kafkaConfig) GetTopic() string {
	return cfg.topic
}

func (cfg kafkaConfig) GetPartition() int {
	return cfg.partition
}

func (cfg kafkaConfig) GetAddress() string {

	return fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
}

type KafkaConfig interface {
	GetTopic() string
	GetPartition() int
	GetAddress() string
}
