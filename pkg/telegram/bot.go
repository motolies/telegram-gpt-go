package telegram

import (
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/motolies/telegram-gpt-go/pkg/openai"
	"log"
)

// 이곳에서 텔레그램의 에코봇을 적용한다.
// 텔레그램의 에코봇은 사용자가 보낸 메시지를 그대로 돌려주는 기능을 한다.

type ChatBot struct {
	BotToken    string
	OpenAIToken string
}

func InitializeServer(botToken string, aiToken string) (*ChatBot, error) {
	return &ChatBot{
		BotToken:    botToken,
		OpenAIToken: aiToken,
	}, nil
}

func (b *ChatBot) Run() error {
	bot, err := telegramApi.NewBotAPI(b.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telegramApi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	// 다른 스레드로 돌아야 하나?
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// update.Message.Text를 가지고 openai에 요청을 보내고, 그 결과를 다시 돌려준다.
			aiResponse := openai.Call(b.OpenAIToken, update.Message.Text)

			msg := telegramApi.NewMessage(update.Message.Chat.ID, aiResponse)
			bot.Send(msg)
		}
	}
	return nil
}
