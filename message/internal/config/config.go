package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

// TODO сделать красиво с областью видимости
type httpConfig struct {
	Port string `yaml:"port" env:"HTTP_PORT"`
}

type HttpConfig interface {
	GetPort() string
	GetAddress() string
}

func (cfg httpConfig) GetPort() string {
	return cfg.Port
}

func (cfg httpConfig) GetAddress() string {
	return fmt.Sprintf(":%v", cfg.Port)
}

func NewHttpConfig() (HttpConfig, error) {
	var cfg httpConfig
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		println(err.Error())
	}
	log.Printf("%+v\n", cfg)
	return cfg, err
}
