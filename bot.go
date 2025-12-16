package tgbot

import (
	"context"
	"log/slog"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func SendMessage(chatID int64, text string, buttons ...[]Button) {
	msg := tu.Message(tu.ID(chatID), text)

	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		slog.Error("can't send the message", "chat", chatID, "err", err, "full_msg", msg)
	}
}

// SendMessageHTML отправляет сообщение с HTML форматированием
func SendMessageHTML(chatID int64, htmlText string, buttons ...[]Button) {
	msg := tu.Message(tu.ID(chatID), htmlText)
	msg.ParseMode = telego.ModeHTML
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		slog.Error("can't send HTML message", "chat", chatID, "err", err, "full_msg", msg)
	}
}

// SendMessageMarkdown отправляет сообщение с Markdown форматированием
func SendMessageMarkdown(chatID int64, markdownText string, buttons ...[]Button) {
	msg := tu.Message(tu.ID(chatID), markdownText)
	msg.ParseMode = telego.ModeMarkdownV2
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		slog.Error("can't send Markdown message", "chat", chatID, "err", err, "full_msg", msg)
	}
}

func ReplyMessage(chatID int64, msgID int, text string, buttons ...[]Button) {
	msg := tu.Message(tu.ID(chatID), text)
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	msg.ReplyParameters = &telego.ReplyParameters{
		MessageID: msgID,
	}
	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		slog.Error("can't send the message", "chat", chatID, "err", err, "full_msg", msg)
	}
}

// ReplyMessageHTML отвечает на сообщение с HTML форматированием
func ReplyMessageHTML(chatID int64, msgID int, htmlText string, buttons ...[]Button) {
	msg := tu.Message(tu.ID(chatID), htmlText)
	msg.ParseMode = telego.ModeHTML
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}
	msg.ReplyParameters = &telego.ReplyParameters{
		MessageID: msgID,
	}

	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		slog.Error("can't reply with HTML message", "chat", chatID, "err", err, "full_msg", msg)
	}
}

// ReplyMessageMarkdown отвечает на сообщение с Markdown форматированием
func ReplyMessageMarkdown(chatID int64, msgID int, markdownText string, buttons ...[]Button) {
	msg := tu.Message(tu.ID(chatID), markdownText)
	msg.ParseMode = telego.ModeMarkdownV2
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}
	msg.ReplyParameters = &telego.ReplyParameters{
		MessageID: msgID,
	}

	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		slog.Error("can't reply with Markdown message", "chat", chatID, "err", err, "full_msg", msg)
	}
}

func DeleteMessage(chatID int64, messageID int) {
	err := bot.DeleteMessage(context.Background(), tu.Delete(tu.ID(chatID), messageID))
	if err != nil {
		slog.Error("can't delete the message", "chat", chatID, "msg_id", messageID, "err", err)
	}
}

func UpdateKeyboard(chatID int64, messageID int, buttons ...[]Button) {
	keyboard := createInlineKeyboard(buttons)

	_, err := bot.EditMessageReplyMarkup(context.Background(), &telego.EditMessageReplyMarkupParams{
		ChatID:      tu.ID(chatID),
		MessageID:   messageID,
		ReplyMarkup: keyboard,
	})
	if err != nil {
		slog.Error("can't edit the keyboard", "chat", chatID, "msg_id", messageID, "err", err)
	}
}

func createInlineKeyboard(buttons [][]Button) *telego.InlineKeyboardMarkup {
	var markupTable [][]telego.InlineKeyboardButton
	for _, row := range buttons {
		r := make([]telego.InlineKeyboardButton, 0, len(row))
		for _, btn := range row {
			if btn.webApp {
				r = append(r, telego.InlineKeyboardButton{
					Text:   btn.Text,
					WebApp: &telego.WebAppInfo{URL: cfg().WebAppURL},
				})
			} else {
				r = append(r, tu.InlineKeyboardButton(btn.Text).WithCallbackData(btn.Data))
			}
		}
		markupTable = append(markupTable, r)
	}

	if len(buttons) != 0 && len(buttons[0]) != 0 {
		return tu.InlineKeyboard(markupTable...)
	}

	return nil
}

// HTML форматирование - вспомогательные функции
func Bold(text string) string {
	return "<b>" + text + "</b>"
}

func Italic(text string) string {
	return "<i>" + text + "</i>"
}

func Code(text string) string {
	return "<code>" + text + "</code>"
}

func Pre(text string) string {
	return "<pre>" + text + "</pre>"
}

func Link(text, url string) string {
	return `<a href="` + url + `">` + text + `</a>`
}

func Spoiler(text string) string {
	return "<tg-spoiler>" + text + "</tg-spoiler>"
}

// Markdown форматирование - вспомогательные функции
func BoldMD(text string) string {
	return "*" + text + "*"
}

func ItalicMD(text string) string {
	return "_" + text + "_"
}

func CodeMD(text string) string {
	return "`" + text + "`"
}

func PreMD(text string) string {
	return "```\n" + text + "\n```"
}

func LinkMD(text, url string) string {
	return "[" + text + "](" + url + ")"
}

func SpoilerMD(text string) string {
	return "||" + text + "||"
}
