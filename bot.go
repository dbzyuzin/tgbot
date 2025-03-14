package tgbot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func SendMessage(chatID int64, text string, buttons ...[]Button) {
	msg := tgbotapi.NewMessage(chatID, text)

	addKeyboard(&msg, buttons)

	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Не удалось отправить сообщение в чат", chatID)
	}
}

func ReplyMessage(chatID int64, msgID int, text string, buttons ...[]Button) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = msgID

	addKeyboard(&msg, buttons)

	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println("Не удалось отправить сообщение в чат", chatID)
	}
}

func DeleteMessage(chatID int64, messageID int) {
	_, err := bot.MakeRequest("deleteMessage", map[string]string{
		"chat_id":    strconv.FormatInt(chatID, 10),
		"message_id": strconv.FormatInt(int64(messageID), 10),
	})
	if err != nil {
		fmt.Printf("Не удалось удалить сообщение %d в чатe %d", messageID, chatID)
	}
}

func addKeyboard(msg *tgbotapi.MessageConfig, buttons [][]Button) {
	var markupTable [][]tgbotapi.InlineKeyboardButton
	for _, row := range buttons {
		r := make([]tgbotapi.InlineKeyboardButton, 0, len(row))
		for _, btn := range row {
			r = append(r, tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.Data))
		}
		markupTable = append(markupTable, r)
	}
	if len(buttons) != 0 && len(buttons[0]) != 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(markupTable...)
	}
}
