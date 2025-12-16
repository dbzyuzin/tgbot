package tgbot

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type Chat interface {
	ID() int64
}

type chat struct {
	ID int64
}

func (ch *chat) SendText(text string, buttons ...[]Button) error {
	return ch.SendTextCtx(context.Background(), text, buttons...)
}

func (ch *chat) SendTextCtx(ctx context.Context, text string, buttons ...[]Button) error {
	msg := tu.Message(tu.ID(ch.ID), text)

	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		slog.Error("can't send the message", "chat", ch.ID, "err", err, "full_msg", msg)
		return fmt.Errorf("can't send the message: %w", err)
	}

	return nil
}

func (ch *chat) SendHTML(chatID int64, htmlText string, buttons ...[]Button) error {
	return ch.SendHTMLCtx(context.Background(), htmlText, buttons...)
}

func (ch *chat) SendHTMLCtx(ctx context.Context, htmlText string, buttons ...[]Button) error {
	msg := tu.Message(tu.ID(ch.ID), htmlText)
	msg.ParseMode = telego.ModeHTML
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	_, err := bot.SendMessage(ctx, msg)
	if err != nil {
		slog.Error("can't send HTML message", "chat", ch.ID, "err", err, "full_msg", msg)
		return fmt.Errorf("can't send HTML message: %w", err)
	}

	return nil
}

func (ch *chat) SendMarkdown(chatID int64, markdownText string, buttons ...[]Button) error {
	return ch.SendMarkdownCtx(context.Background(), chatID, markdownText, buttons...)
}

func (ch *chat) SendMarkdownCtx(ctx context.Context, markdownText string, buttons ...[]Button) error {
	msg := tu.Message(tu.ID(ch.ID), markdownText)
	msg.ParseMode = telego.ModeMarkdownV2
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	_, err := bot.SendMessage(ctx, msg)
	if err != nil {
		slog.Error("can't send Markdown message", "chat", ch.ID, "err", err, "full_msg", msg)
		return fmt.Errorf("can't send Markdown message: %w", err)
	}

	return nil
}

func (ch *chat) ReplyMessage(chatID int64, msgID int, text string, buttons ...[]Button) {
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
func (ch *chat) ReplyMessageHTML(chatID int64, msgID int, htmlText string, buttons ...[]Button) {
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
func (ch *chat) ReplyMessageMarkdown(chatID int64, msgID int, markdownText string, buttons ...[]Button) {
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

func (ch *chat) DeleteMessage(chatID int64, messageID int) {
	err := bot.DeleteMessage(context.Background(), tu.Delete(tu.ID(chatID), messageID))
	if err != nil {
		slog.Error("can't delete the message", "chat", chatID, "msg_id", messageID, "err", err)
	}
}

func (ch *chat) UpdateKeyboard(chatID int64, messageID int, buttons ...[]Button) {
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
