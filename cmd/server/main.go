package main

import (
	"github.com/motolies/telegram-gpt-go/pkg/telegram"
	"os"
)

func main() {
	bot, err := telegram.InitializeServer(os.Getenv("TELEGRAM_BOT_TOKEN"), "b")
	if err != nil {
		panic(err)
	}
	bot.Run()
}
