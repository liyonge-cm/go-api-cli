package app

import (
	"context"

	"github.com/liyonge-cm/go-api-cli-prj/config"

	"go.uber.org/zap"
)

type Service interface {
	Start()
	Close()
	WithLogger(*zap.Logger)
}

type App struct {
	ctx         context.Context
	config      *config.Config
	appServices []Service
	Logger      *zap.Logger
}

func NewApp(config *config.Config) *App {
	return &App{
		ctx:         context.Background(),
		config:      config,
		appServices: []Service{},
	}
}

func (s *App) WithLogger(log *zap.Logger) {
	s.Logger = log
}

func (s *App) RegistService(service Service) {
	s.appServices = append(s.appServices, service)
}

func (s *App) StartServices() {
	for _, service := range s.appServices {
		service.WithLogger(s.Logger)
		service.Start()
	}
}

func (s *App) CloseServices() {
	for _, service := range s.appServices {
		service.Close()
	}
}
