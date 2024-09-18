package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type ScyllaConfig interface {
	GetAddress() string
	GetKeyspace() string
}

// TODO [16.09.2024] мб вместо только env сделать нормальный конфигуратор
type scyllaConfig struct {
	Host     string `env:"SCYLLA_HOST"`
	Port     string `env:"SCYLLA_PORT"`
	Keyspace string `env:"SCYLLA_KEYSPACE"`
}

func NewScyllaConfig() (ScyllaConfig, error) {
	var cfg scyllaConfig

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		println(err.Error())
	}
	log.Printf("%+v\n", cfg)
	return cfg, err
}

func (cfg scyllaConfig) GetAddress() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

func (cfg scyllaConfig) GetKeyspace() string {
	return cfg.Keyspace
}
