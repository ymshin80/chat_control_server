package app

import (
	"chat_golang_control/config"
	"chat_golang_control/network"
	"chat_golang_control/repository"
	"chat_golang_control/service"
)

type App struct {
	cfg *config.Config
	repository *repository.Repository
	service *service.Service
	network	*network.Server
}

func NewApp(cfg *config.Config) *App {
	a := &App{cfg: cfg}

	var err error

	if a.repository, err = repository.NewRepository(cfg); err != nil {
		panic(err)
	} else {
		a.service = service.NewService(a.repository)
		a.network = network.NewNetwork(a.service, cfg.Info.Port)
	} 

	return a
}

func (a *App) Start() error {
	return a.network.Start()
}