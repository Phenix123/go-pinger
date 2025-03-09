package service

import (
	"github.com/Phenix123/go-pinger/config"
	"github.com/Phenix123/go-pinger/internal/pinger"
	"github.com/Phenix123/go-pinger/pkg/tg_bot"
	"github.com/sirupsen/logrus"
	"time"
)

func Run(cfg *config.Config, log *logrus.Logger) {
	tgBot := tg_bot.NewTGBot(cfg.BotToken, cfg.ChatID)

	p := pinger.NewPinger(cfg.Host)

	ticker := time.NewTicker(time.Minute) // Check every minute
	defer ticker.Stop()

	log.Info("Service started, checking host every minute")

	for {
		select {
		case <-ticker.C:
			statusCode := p.Ping()

			log.WithFields(logrus.Fields{
				"host":        cfg.Host,
				"status_code": statusCode,
			}).Info("Ping result")

			tgBot.CheckResp(statusCode)
		}
	}
}
