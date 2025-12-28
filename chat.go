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

	Send(text string, opts ...SendOption) error
	SendCtx(ctx context.Context, text string, opts ...SendOption) error

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

func (ch *chat) Send(text string, opts ...SendOption) error {
	return ch.SendCtx(context.Background(), text, opts...)
}

func (ch *chat) SendCtx(ctx context.Context, text string, opts ...SendOption) error {
	var o sendOptions
	for _, opt := range opts {
		opt(&o)
	}

	msg := tu.Message(tu.ID(ch.id), text)

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
		slog.Error("can't send the message", "chat", ch.id, "err", err)
		return fmt.Errorf("can't send the message: %w", err)
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
