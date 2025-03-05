package main

import (
	"github.com/joho/godotenv"
	"log"
	"siteChecker/config"
	"siteChecker/internal/pinger"
	"siteChecker/pkg/tg_bot"
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
