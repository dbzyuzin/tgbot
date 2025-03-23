package tgbot

import (
	"context"
	"fmt"
	"github.com/mymmrac/telego"
	"log"
)

var bot *telego.Bot

// Start запускает бота и замораживает поток до ручной остановки
func Start() {
	var err error
	bot, err = telego.NewBot(cfg.BotToken)
	if err != nil {
		myPanic(err.Error(), "Не удалось подключиться к телеграмм, проверь токен.")
	}

	updates, _ := bot.UpdatesViaLongPolling(context.Background(), nil)

	me, err := bot.GetMe(context.Background())
	if err != nil {
		myPanic(err.Error(), "Не удалось подключиться к телеграмм, проверь токен.")
		return
	}

	log.Printf("bot \"%s\" started", me.Username)

	for update := range updates {
		switch {
		case update.Message != nil:
			handleMessage(update)
		case update.CallbackQuery != nil && update.CallbackQuery.Message.IsAccessible():
			handleCallback(update)
		default:
			fmt.Println("unknown update")
		}
	}
}

func handleMessage(update telego.Update) {
	if handlers.newMessage == nil {
		return
	}
	handlers.newMessage(mapMessage(update.Message))
}

func handleCallback(update telego.Update) {
	if handlers.newCallback == nil {
		return
	}

	// todo: check how to get callback without a msg
	msg, ok := update.CallbackQuery.Message.(*telego.Message)
	if !ok {
		return
	}

	handlers.newCallback(Callback{
		Data:    update.CallbackQuery.Data,
		Message: mapMessage(msg),
	})

	_ = bot.AnswerCallbackQuery(context.Background(), &telego.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})
}

func mapMessage(msg *telego.Message) Message {
	return Message{
		ID:     msg.MessageID,
		User:   mapUser(msg.From),
		ChatID: msg.Chat.ID,
		Text:   msg.Text,
	}
}

func mapUser(user *telego.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.Username,
	}
}
