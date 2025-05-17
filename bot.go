package tgbot

import (
	"context"
	"log/slog"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func SendMessage(chatID int64, text string, buttons ...[]Button) {
	msg := tu.Message(tu.ID(chatID), text)
	addKeyboard(msg, buttons)

	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		slog.Error("can't send the message", "chat", chatID, "err", err, "full_msg", msg)
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
		slog.Error("can't send the message", "chat", chatID, "err", err, "full_msg", msg)
	}
}

func DeleteMessage(chatID int64, messageID int) {
	err := bot.DeleteMessage(context.Background(), tu.Delete(tu.ID(chatID), messageID))
	if err != nil {
		slog.Error("can't delete the message", "chat", chatID, "msg_id", messageID, "err", err)
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
