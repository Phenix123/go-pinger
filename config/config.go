package config

import "os"

type Config struct {
	Host     string
	ChatID   string
	BotToken string
}

func NewConfig() *Config {
	return &Config{
		Host:     os.Getenv("HOST"),
		ChatID:   os.Getenv("CHAT_ID"),
		BotToken: os.Getenv("BOT_TOKEN"),
	}
}
