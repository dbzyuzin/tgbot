package tgbot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	//msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println("Не удалось отправить сообщение в чат", chatID)
	}
}

func ReplyMessage(chatID int64, msgID int, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = msgID

	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println("Не удалось отправить сообщение в чат", chatID)
	}
}
