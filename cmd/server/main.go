package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	updates, err := bot.GetUpdatesChan(tgbotapi.UpdateConfig{})
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		_, err := bot.Send(msg)
		if err != nil {
			return
		}
	}
}
