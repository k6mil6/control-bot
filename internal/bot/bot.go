package bot

import (
	"errors"
	"github.com/k6mil6/control-bot/internal/bot/handlers/commands"
	"github.com/k6mil6/control-bot/internal/bot/replies"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"time"
)

type Bot struct {
	botAPI  *telebot.Bot
	ownerID int64
}

func New(token string, timeout time.Duration, ownerID int64) (*Bot, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: timeout},
	})

	if err != nil {
		return nil, err
	}

	return &Bot{
		botAPI:  bot,
		ownerID: ownerID,
	}, nil
}

func (b *Bot) Start() error {
	adminOnly := b.botAPI.Group()
	adminOnly.Use(middleware.Whitelist(b.ownerID))

	// Register handlers
	adminOnly.Handle("/start", commands.StartHandler)
	adminOnly.Handle("/turn_off", commands.TurnOffHandler)
	adminOnly.Handle("/restart", commands.RestartHandler)

	_, err := b.botAPI.Send(&telebot.User{ID: b.ownerID}, replies.TurnOn)
	if err != nil {
		return err
	}

	// Start bot
	b.botAPI.Start()

	return nil
}

func (b *Bot) Stop() {
	b.botAPI.Stop()
}
