package main

import (
	"context"
	"github.com/k6mil6/control-bot/internal/app"
	"github.com/k6mil6/control-bot/internal/config"
	"github.com/k6mil6/control-bot/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	log := logger.SetupLogger(cfg.ENV)
	log = log.With("env", cfg.ENV)

	application, err := app.New(cfg)
	if err != nil {
		log.Error("cannot create app: " + err.Error())
		return
	}

	go application.MustStart()
	log.Info("bot started")

	<-ctx.Done()

	application.Stop()
	log.Info("bot stopped")
}
