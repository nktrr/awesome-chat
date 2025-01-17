package app

import (
	"awesome-chat/gateway/internal/api"
	"awesome-chat/gateway/internal/config"
	"awesome-chat/gateway/internal/kafka"
	"awesome-chat/gateway/internal/service"
	msgService "awesome-chat/gateway/internal/service/message"
	"context"
	"log"
)

type serviceProvider struct {
	msgProducerCfg config.KafkaConfig
	msgProducer    *kafka.Kafka
	msgService     service.MessageService
	httpServer     *api.HttpServer
	httpConfig     config.HttpConfig
	scyllaConfig   config.ScyllaConfig
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) MessageService(ctx context.Context) service.MessageService {
	if s.msgService == nil {
		s.msgService = msgService.NewService(
			s.MessagesProducer(ctx),
		)
	}
	return s.msgService
}

func (s *serviceProvider) HttpServer(ctx context.Context) *api.HttpServer {
	if s.httpServer == nil {
		s.httpServer = api.NewHttpServer(
			s.HttpConfig(),
			s.MessageService(ctx),
		)
	}
	return s.httpServer
}

func (s *serviceProvider) MessagesProducer(ctx context.Context) *kafka.Kafka {
	if s.msgProducer == nil {
		producer, err := kafka.NewKafka(ctx, s.MessagesProducerConfig())
		if err != nil {
			log.Panicf("failed to connect messages producer: %v", err)
		}
		s.msgProducer = producer
	}
	return s.msgProducer
}

func (s *serviceProvider) MessagesProducerConfig() config.KafkaConfig {
	if s.msgProducerCfg == nil {
		cfg, err := config.NewKafkaConfig(
			"messages",
			0,
		)
		if err != nil {
			log.Panicf("failed to get kafka config: %v", err)
		}
		s.msgProducerCfg = cfg
	}
	return s.msgProducerCfg
}

func (s *serviceProvider) HttpConfig() config.HttpConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHttpConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}
	return s.httpConfig
}

func (s *serviceProvider) ScyllaConfig() config.ScyllaConfig {
	if s.scyllaConfig == nil {
		cfg, err := config.NewScyllaConfig()
		if err != nil {
			log.Fatalf("failed to get scylla config: %s", err.Error())
		}
		s.scyllaConfig = cfg
	}
	return s.scyllaConfig
}
