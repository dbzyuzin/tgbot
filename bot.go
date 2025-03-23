package tgbot

import (
	"context"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func SendMessage(chatID int64, text string, buttons ...[]Button) {
	msg := tu.Message(tu.ID(chatID), text)
	addKeyboard(msg, buttons)

	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Не удалось отправить сообщение в чат", chatID)
	}
}

func ReplyMessage(chatID int64, msgID int, text string, buttons ...[]Button) {
	msg := tu.Message(tu.ID(chatID), text)
	addKeyboard(msg, buttons)

	msg.ReplyParameters = &telego.ReplyParameters{
		MessageID: msgID,
	}
	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		fmt.Println("Не удалось отправить сообщение в чат", chatID)
	}
}

func DeleteMessage(chatID int64, messageID int) {
	err := bot.DeleteMessage(context.Background(), tu.Delete(tu.ID(chatID), messageID))
	if err != nil {
		fmt.Printf("Не удалось удалить сообщение %d в чатe %d", messageID, chatID)
	}
}

func addKeyboard(msg *telego.SendMessageParams, buttons [][]Button) {
	var markupTable [][]telego.InlineKeyboardButton
	for _, row := range buttons {
		r := make([]telego.InlineKeyboardButton, 0, len(row))
		for _, btn := range row {
			r = append(r, tu.InlineKeyboardButton(btn.Text).WithCallbackData(btn.Data))
		}
		markupTable = append(markupTable, r)
	}

	if len(buttons) != 0 && len(buttons[0]) != 0 {
		msg.ReplyMarkup = tu.InlineKeyboard(
			markupTable...,
		)
	}
}
