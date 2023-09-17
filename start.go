package tgbot

import (
	"fmt"
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
	u.AllowedUpdates = []string{"message", "callback_query"}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		switch {
		case update.Message != nil:
			handleMessage(update)
		case update.CallbackQuery != nil:
			handleCallback(update)
		default:
			fmt.Println("unknown update")
		}
	}
}

func handleMessage(update tgbotapi.Update) {
	if handlers.newMessage == nil {
		return
	}
	handlers.newMessage(mapMessage(update.Message))
}

func handleCallback(update tgbotapi.Update) {
	if handlers.newCallback == nil {
		return
	}
	handlers.newCallback(Callback{
		Data:    update.CallbackQuery.Data,
		Message: mapMessage(update.CallbackQuery.Message),
	})
}

func mapMessage(msg *tgbotapi.Message) Message {
	return Message{
		ID:     msg.MessageID,
		User:   mapUser(msg.From),
		ChatID: msg.Chat.ID,
		Text:   msg.Text,
	}
}

func mapUser(user *tgbotapi.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
	}

}
