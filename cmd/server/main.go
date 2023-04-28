package main

import (
	"fmt"
	"github.com/motolies/telegram-gpt-go/pkg/customLog"
	"github.com/motolies/telegram-gpt-go/pkg/telegram"
	"os"
	"time"
)

func main() {
	bot, err := telegram.InitializeServer(os.Getenv("TELEGRAM_BOT_TOKEN"), os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		customLog.ColorLog(err.Error(), customLog.ERROR)
		return
	}

	for {
		if err := bot.Run(); err != nil {
			customLog.ColorLog(fmt.Sprintf("Bot encountered an error: %v", err), customLog.ERROR)
			time.Sleep(time.Second * 5)
		}
	}
}
