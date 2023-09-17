package tgbot

import (
	"fmt"
	"strconv"

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

func DeleteMessage(chatID int64, message_id int) {
	_, err := bot.MakeRequest("deleteMessage", map[string]string{
		"chat_id":    strconv.FormatInt(chatID, 10),
		"message_id": strconv.FormatInt(int64(message_id), 10),
	})
	if err != nil {
		fmt.Printf("Не удалось удалить сообщение %d в чатe %d", message_id, chatID)
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
