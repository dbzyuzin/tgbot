package tgbot

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println("Не удалось отправить сообщение в чат", chatID)
	}
}

func SendMessageWithKeyboard(chatID int64, text string, buttons [][]Button) {
	msg := tgbotapi.NewMessage(chatID, text)

	markupTable := [][]tgbotapi.InlineKeyboardButton{}
	for _, row := range buttons {
		r := make([]tgbotapi.InlineKeyboardButton, 0, len(row))
		for _, btn := range row {
			r = append(r, tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.Data))
		}
		markupTable = append(markupTable, r)
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(markupTable...)

	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println(err)
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
