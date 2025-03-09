package config

import (
	"os"
	"strconv"
)

type Config struct {
	Host         string
	ChatID       string
	BotToken     string
	PingInterval int
}

func NewConfig() *Config {
	p, err := strconv.Atoi(os.Getenv("PING_INTERVAL"))
	if err != nil {
		panic(err)
	}
	return &Config{
		Host:         os.Getenv("HOST"),
		ChatID:       os.Getenv("CHAT_ID"),
		BotToken:     os.Getenv("BOT_TOKEN"),
		PingInterval: p,
	}
}
