package botapp

import (
	"github.com/k6mil6/control-bot/internal/bot"
	"github.com/k6mil6/control-bot/internal/config"
)

type App struct {
	bot *bot.Bot
}

func New(config *config.Config) (*App, error) {
	tgBot, err := bot.New(config.Bot.Token, config.Bot.Timeout, config.Bot.OwnerID)
	if err != nil {
		return nil, err
	}

	return &App{
		bot: tgBot,
	}, nil
}

func (a *App) Start() error {
	return a.bot.Start()
}

func (a *App) Stop() {
	a.bot.Stop()
}
