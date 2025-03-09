package service

import (
	"context"
	"github.com/Phenix123/go-pinger/config"
	"github.com/Phenix123/go-pinger/internal/pinger"
	"github.com/Phenix123/go-pinger/pkg/tg_bot"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config, log *logrus.Logger) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	tgBot := tg_bot.NewTGBot(cfg.BotToken, cfg.ChatID)

	p := pinger.NewPinger(cfg.Host)

	duration := time.Duration(cfg.PingInterval) * time.Second
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	log.Info("Service started, checking host every ", cfg.PingInterval, " seconds")

	// Run in a goroutine to handle pinging and logging
	go func() {
		for {
			select {
			case <-ticker.C:
				statusCode := p.Ping()

				log.WithFields(logrus.Fields{
					"host":        cfg.Host,
					"status_code": statusCode,
				}).Info("Ping result")

				tgBot.CheckResp(statusCode)

			case <-ctx.Done():
				log.Info("Shutting down ticker...")
				return
			}
		}
	}()

	// Wait for termination signal (Ctrl+C or SIGTERM)
	sigReceived := <-signalChan
	log.Infof("Received signal %s, shutting down...", sigReceived)

	cancel()

	time.Sleep(2 * time.Second)
	log.Info("Service stopped gracefully.")
}
