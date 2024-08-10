package commands

import (
	"github.com/k6mil6/control-bot/internal/bot/replies"
	"github.com/k6mil6/control-bot/internal/service/system"
	"gopkg.in/telebot.v3"
)

func TurnOffHandler(ctx telebot.Context) error {
	err := system.TurnOff()
	if err != nil {
		return ctx.Send(replies.TurnOffFailed)
	}

	return ctx.Send(replies.TurnedOff)
}
