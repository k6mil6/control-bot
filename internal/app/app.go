package app

import (
	botapp "github.com/k6mil6/control-bot/internal/app/bot"
	"github.com/k6mil6/control-bot/internal/config"
)

type App struct {
	BotApp *botapp.App
}

func New(config *config.Config) (*App, error) {
	botApp, err := botapp.New(config)
	if err != nil {
		return nil, err
	}

	return &App{
		BotApp: botApp,
	}, nil
}

func (a *App) Start() error {
	return a.BotApp.Start()
}

func (a *App) MustStart() {
	err := a.Start()
	if err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	a.BotApp.Stop()
}
