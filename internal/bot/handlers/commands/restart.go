package commands

import (
	"github.com/k6mil6/control-bot/internal/bot/replies"
	"github.com/k6mil6/control-bot/internal/service/system"
	"gopkg.in/telebot.v3"
)

func RestartHandler(ctx telebot.Context) error {
	err := system.Restart()
	if err != nil {
		return ctx.Send(replies.RestartFailed)
	}

	return ctx.Send(replies.Restart)
}
