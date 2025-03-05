package tg_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

type TgBot struct {
	botToken string
	chatID   string
}

func NewTGBot(botToken string, chatID string) *TgBot {
	return &TgBot{botToken: botToken, chatID: chatID}
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
	token := b.botToken
	chatID, err := strconv.ParseInt(b.chatID, 10, 64)
	if err != nil {
		return err
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, body)
	_, err = bot.Send(msg)
	return err
}
