package tgbot

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type SendOption func(*sendOptions)

type sendOptions struct {
	buttons   [][]Button
	replyTo   int
	parseMode string
	silent    bool
}

func WithButtons(buttons ...Button) SendOption {
	return func(o *sendOptions) {
		o.buttons = append(o.buttons, buttons)
	}
}

func WithKeyboard(rows ...[]Button) SendOption {
	return func(o *sendOptions) {
		o.buttons = append(o.buttons, rows...)
	}
}

func WithReply(msgID int) SendOption {
	return func(o *sendOptions) {
		o.replyTo = msgID
	}
}

func WithHTML() SendOption {
	return func(o *sendOptions) {
		o.parseMode = telego.ModeHTML
	}
}

func WithHTMLIf(condition bool) SendOption {
	return func(o *sendOptions) {
		if condition {
			o.parseMode = telego.ModeHTML
		}
	}
}

func WithMarkdown() SendOption {
	return func(o *sendOptions) {
		o.parseMode = telego.ModeMarkdownV2
	}
}

func WithMarkdownIf(condition bool) SendOption {
	return func(o *sendOptions) {
		if condition {
			o.parseMode = telego.ModeMarkdownV2
		}
	}
}

func WithSilent() SendOption {
	return func(o *sendOptions) {
		o.silent = true
	}
}

func WithSilentIf(condition bool) SendOption {
	return func(o *sendOptions) {
		if condition {
			o.silent = true
		}
	}
}

func WithWebAppButton(text string) SendOption {
	return func(o *sendOptions) {
		o.buttons = append(o.buttons, []Button{{Text: text, webApp: true}})
	}
}

func Send(chatID int64, text string, opts ...SendOption) error {
	return SendCtx(context.Background(), chatID, text, opts...)
}

func SendCtx(ctx context.Context, chatID int64, text string, opts ...SendOption) error {
	var o sendOptions
	for _, opt := range opts {
		opt(&o)
	}

	msg := tu.Message(tu.ID(chatID), text)

	if o.parseMode != "" {
		msg.ParseMode = o.parseMode
	}

	if keyboard := createInlineKeyboard(o.buttons); keyboard != nil {
		msg.ReplyMarkup = keyboard
	}

	if o.replyTo != 0 {
		msg.ReplyParameters = &telego.ReplyParameters{
			MessageID: o.replyTo,
		}
	}

	if o.silent {
		msg.DisableNotification = true
	}

	_, err := bot.SendMessage(ctx, msg)
	if err != nil {
		slog.Error("can't send the message", "chat", chatID, "err", err)
		return fmt.Errorf("can't send the message: %w", err)
	}

	return nil
}

func DeleteMessage(chatID int64, messageID int) error {
	err := bot.DeleteMessage(context.Background(), tu.Delete(tu.ID(chatID), messageID))
	if err != nil {
		slog.Error("can't delete the message", "chat", chatID, "msg_id", messageID, "err", err)
		return fmt.Errorf("can't delete the message: %w", err)
	}
	return nil
}

func UpdateKeyboard(chatID int64, messageID int, buttons ...[]Button) error {
	keyboard := createInlineKeyboard(buttons)

	_, err := bot.EditMessageReplyMarkup(context.Background(), &telego.EditMessageReplyMarkupParams{
		ChatID:      tu.ID(chatID),
		MessageID:   messageID,
		ReplyMarkup: keyboard,
	})
	if err != nil {
		slog.Error("can't edit the keyboard", "chat", chatID, "msg_id", messageID, "err", err)
		return fmt.Errorf("can't edit the keyboard: %w", err)
	}
	return nil
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
