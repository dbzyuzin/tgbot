package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

// Start запускает бота и замораживает поток до ручной остановки
func Start() {
	var err error
	bot, err = tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		myPanic(err.Error(), "Не удалось подключиться к телеграмм, проверь токен.")
	}
	log.Printf("Бот \"%s\" запустился успешно", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && handlers.newMessage != nil {
			//fmt.Printf("msg: [%s] %s\n", update.Message.From.UserName, update.Message.Text)
			handlers.newMessage(Message{
				ID: update.Message.MessageID,
				User: User{
					FirstName: update.Message.From.FirstName,
					LastName:  update.Message.From.LastName,
					UserName:  update.Message.From.UserName,
				},
				ChatID: update.Message.Chat.ID,
				Text:   update.Message.Text,
			})
		}
	}
}
