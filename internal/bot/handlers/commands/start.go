package commands

import (
	"github.com/k6mil6/control-bot/internal/bot/replies"
	"gopkg.in/telebot.v3"
)

func StartHandler(ctx telebot.Context) error {
	return ctx.Send(replies.Greeting)
}
