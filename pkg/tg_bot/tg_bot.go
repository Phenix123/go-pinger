package tg_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

type TgBot struct {
	botToken string
	chatID   string
	bot      *tgbotapi.BotAPI
}

func NewTGBot(botToken string, chatID string) *TgBot {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	return &TgBot{
		botToken: botToken,
		chatID:   chatID,
		bot:      bot,
	}
}

func (b *TgBot) CheckResp(statusCode int) {
	if statusCode >= 200 {
		err := b.SendNotification("Response status is: " + strconv.Itoa(statusCode))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (b *TgBot) SendNotification(body string) error {
	chatID, err := strconv.ParseInt(b.chatID, 10, 64)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatID, body)
	_, err = b.bot.Send(msg)
	return err
}
