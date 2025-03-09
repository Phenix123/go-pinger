package main

import (
	"github.com/Phenix123/go-pinger/config"
	"github.com/Phenix123/go-pinger/internal/pinger"
	"github.com/Phenix123/go-pinger/pkg/tg_bot"
	"github.com/joho/godotenv"
	"log"
)

// send get method to host got from .env and send to tg notification if response status is >400
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewConfig()

	tgBot := tg_bot.NewTGBot(cfg.BotToken, cfg.ChatID)

	statusCode := pinger.NewPinger(cfg.Host).Ping()

	tgBot.CheckResp(statusCode)
}
