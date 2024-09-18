package app

import (
	"awesome-chat/gateway/internal/api"
	"awesome-chat/gateway/internal/config"
	"awesome-chat/gateway/internal/repository"
	msgRepo "awesome-chat/gateway/internal/repository/message"
	"awesome-chat/gateway/internal/service"
	msgService "awesome-chat/gateway/internal/service/message"
	"context"
	"log"
)

type serviceProvider struct {
	msgRepository repository.MessageRepository
	msgService    service.MessageService
	httpServer    *api.HttpServer
	httpConfig    config.HttpConfig
	scyllaConfig  config.ScyllaConfig
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	var err error
	if s.msgRepository == nil {
		s.msgRepository, err = msgRepo.NewRepository(
			s.ScyllaConfig(),
		)
	}
	if err != nil {
		log.Panicf("failed to init msg repo: %v", err)
	}
	return s.msgRepository
}

func (s *serviceProvider) MessageService(ctx context.Context) service.MessageService {
	if s.msgService == nil {
		s.msgService = msgService.NewService(
			s.MessageRepository(ctx),
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
