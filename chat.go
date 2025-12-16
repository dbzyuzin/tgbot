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
	Bot() *telego.Bot

	SendText(text string, buttons ...[]Button) error
	SendTextCtx(ctx context.Context, text string, buttons ...[]Button) error
	SendHTML(htmlText string, buttons ...[]Button) error
	SendHTMLCtx(ctx context.Context, htmlText string, buttons ...[]Button) error
	SendMarkdown(markdownText string, buttons ...[]Button) error
	SendMarkdownCtx(ctx context.Context, markdownText string, buttons ...[]Button) error

	ReplyText(msgID int, text string, buttons ...[]Button) error
	ReplyTextCtx(ctx context.Context, msgID int, text string, buttons ...[]Button) error
	ReplyHTML(msgID int, htmlText string, buttons ...[]Button) error
	ReplyHTMLCtx(ctx context.Context, msgID int, htmlText string, buttons ...[]Button) error
	ReplyMarkdown(msgID int, markdownText string, buttons ...[]Button) error
	ReplyMarkdownCtx(ctx context.Context, msgID int, markdownText string, buttons ...[]Button) error

	DeleteMessage(messageID int) error
	DeleteMessageCtx(ctx context.Context, messageID int) error
	UpdateKeyboard(messageID int, buttons ...[]Button) error
	UpdateKeyboardCtx(ctx context.Context, messageID int, buttons ...[]Button) error
}

type chat struct {
	id int64
}

func (ch *chat) ID() int64 {
	return ch.id
}

func (ch *chat) Bot() *telego.Bot {
	return bot
}

func (ch *chat) SendText(text string, buttons ...[]Button) error {
	return ch.SendTextCtx(context.Background(), text, buttons...)
}

func (ch *chat) SendTextCtx(ctx context.Context, text string, buttons ...[]Button) error {
	msg := tu.Message(tu.ID(ch.id), text)

	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	_, err := bot.SendMessage(context.Background(), msg)
	if err != nil {
		slog.Error("can't send the message", "chat", ch.id, "err", err, "full_msg", msg)
		return fmt.Errorf("can't send the message: %w", err)
	}

	return nil
}

func (ch *chat) SendHTML(htmlText string, buttons ...[]Button) error {
	return ch.SendHTMLCtx(context.Background(), htmlText, buttons...)
}

func (ch *chat) SendHTMLCtx(ctx context.Context, htmlText string, buttons ...[]Button) error {
	msg := tu.Message(tu.ID(ch.id), htmlText)
	msg.ParseMode = telego.ModeHTML
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	_, err := bot.SendMessage(ctx, msg)
	if err != nil {
		slog.Error("can't send HTML message", "chat", ch.id, "err", err, "full_msg", msg)
		return fmt.Errorf("can't send HTML message: %w", err)
	}

	return nil
}

func (ch *chat) SendMarkdown(markdownText string, buttons ...[]Button) error {
	return ch.SendMarkdownCtx(context.Background(), markdownText, buttons...)
}

func (ch *chat) SendMarkdownCtx(ctx context.Context, markdownText string, buttons ...[]Button) error {
	msg := tu.Message(tu.ID(ch.id), markdownText)
	msg.ParseMode = telego.ModeMarkdownV2
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = createInlineKeyboard(buttons)
	}

	_, err := bot.SendMessage(ctx, msg)
	if err != nil {
		slog.Error("can't send Markdown message", "chat", ch.id, "err", err, "full_msg", msg)
		return fmt.Errorf("can't send Markdown message: %w", err)
	}

	return nil
}

func (ch *chat) ReplyText(msgID int, text string, buttons ...[]Button) error {
	return ch.ReplyTextCtx(context.Background(), msgID, text, buttons...)
}

func (ch *chat) ReplyTextCtx(ctx context.Context, msgID int, text string, buttons ...[]Button) error {
	msg := tu.Message(tu.ID(ch.id), text)
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = keyboard
	}

	msg.ReplyParameters = &telego.ReplyParameters{
		MessageID: msgID,
	}
	_, err := bot.SendMessage(ctx, msg)
	if err != nil {
		slog.Error("can't send the message", "chat", ch.id, "err", err, "full_msg", msg)
		return fmt.Errorf("can't send the message: %w", err)
	}
	return nil
}

func (ch *chat) ReplyHTML(msgID int, htmlText string, buttons ...[]Button) error {
	return ch.ReplyHTMLCtx(context.Background(), msgID, htmlText, buttons...)
}

func (ch *chat) ReplyHTMLCtx(ctx context.Context, msgID int, htmlText string, buttons ...[]Button) error {
	msg := tu.Message(tu.ID(ch.id), htmlText)
	msg.ParseMode = telego.ModeHTML
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = keyboard
	}
	msg.ReplyParameters = &telego.ReplyParameters{
		MessageID: msgID,
	}

	_, err := bot.SendMessage(ctx, msg)
	if err != nil {
		slog.Error("can't reply with HTML message", "chat", ch.id, "err", err, "full_msg", msg)
		return fmt.Errorf("can't reply with HTML message: %w", err)
	}
	return nil
}

func (ch *chat) ReplyMarkdown(msgID int, markdownText string, buttons ...[]Button) error {
	return ch.ReplyMarkdownCtx(context.Background(), msgID, markdownText, buttons...)
}

func (ch *chat) ReplyMarkdownCtx(ctx context.Context, msgID int, markdownText string, buttons ...[]Button) error {
	msg := tu.Message(tu.ID(ch.id), markdownText)
	msg.ParseMode = telego.ModeMarkdownV2
	keyboard := createInlineKeyboard(buttons)
	if keyboard != nil {
		msg.ReplyMarkup = keyboard
	}
	msg.ReplyParameters = &telego.ReplyParameters{
		MessageID: msgID,
	}

	_, err := bot.SendMessage(ctx, msg)
	if err != nil {
		slog.Error("can't reply with Markdown message", "chat", ch.id, "err", err, "full_msg", msg)
		return fmt.Errorf("can't reply with Markdown message: %w", err)
	}
	return nil
}

func (ch *chat) DeleteMessage(messageID int) error {
	return ch.DeleteMessageCtx(context.Background(), messageID)
}

func (ch *chat) DeleteMessageCtx(ctx context.Context, messageID int) error {
	err := bot.DeleteMessage(ctx, tu.Delete(tu.ID(ch.id), messageID))
	if err != nil {
		slog.Error("can't delete the message", "chat", ch.id, "msg_id", messageID, "err", err)
		return fmt.Errorf("can't delete the message: %w", err)
	}
	return nil
}

func (ch *chat) UpdateKeyboard(messageID int, buttons ...[]Button) error {
	return ch.UpdateKeyboardCtx(context.Background(), messageID, buttons...)
}

func (ch *chat) UpdateKeyboardCtx(ctx context.Context, messageID int, buttons ...[]Button) error {
	keyboard := createInlineKeyboard(buttons)

	_, err := bot.EditMessageReplyMarkup(ctx, &telego.EditMessageReplyMarkupParams{
		ChatID:      tu.ID(ch.id),
		MessageID:   messageID,
		ReplyMarkup: keyboard,
	})
	if err != nil {
		slog.Error("can't edit the keyboard", "chat", ch.id, "msg_id", messageID, "err", err)
		return fmt.Errorf("can't edit the keyboard: %w", err)
	}
	return nil
}
