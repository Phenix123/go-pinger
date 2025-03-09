package main

import (
	"github.com/Phenix123/go-pinger/config"
	"github.com/Phenix123/go-pinger/internal/service"
	"github.com/Phenix123/go-pinger/pkg/logger"
	"github.com/joho/godotenv"
)

// send get method to host got from .env and send to tg notification if response status is >400
func main() {
	log := logger.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewConfig()

	service.Run(cfg, log)
}
