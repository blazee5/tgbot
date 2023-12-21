package main

import (
	"github.com/blazee5/tgbot/config"
	"github.com/blazee5/tgbot/internal/handler"
	"github.com/blazee5/tgbot/internal/repository"
	"github.com/blazee5/tgbot/internal/service"
	"github.com/blazee5/tgbot/pkg/db/postgres"
	"github.com/blazee5/tgbot/pkg/db/redis"
	"github.com/blazee5/tgbot/pkg/logger"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	log := logger.NewLogger()

	settings := tele.Settings{
		Token:  cfg.Bot.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(settings)
	if err != nil {
		log.Fatal(err)
		return
	}

	db := postgres.NewPostgres(cfg)
	rdb := redis.NewRedis(cfg)

	repo := repository.NewRepository(db, rdb)
	services := service.NewService(log, repo)
	handlers := handler.NewHandler(log, services, b, cfg)

	handlers.Register()

	b.Use(middleware.Logger())

	log.Info("bot starting...")

	b.Start()
}
