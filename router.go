package tgbot

import (
	"context"
	"strings"

	"github.com/mymmrac/telego"
)

func handleMessage(ctx context.Context, update telego.Update) {
	text := strings.TrimSpace(update.Message.Text)

	ch := chat{update.Message.Chat.ID}

	if !strings.HasPrefix(text, "/") {
		if len(handlers.messageHandlers) == 0 {
			return
		}

		msg := mapMessage(update.Message)
		for _, handler := range handlers.messageHandlers {
			handler(ctx, &ch, msg)
		}
	}

	userCommand := strings.Trim(text, "/")

	found := false

	for cmd, handler := range handlers.commandHandlers {
		if strings.HasPrefix(userCommand, cmd) {
			handler(ctx, &ch, mapMessage(update.Message))
			found = true
			break
		}
	}

	if !found && handlers.unknownCommand != nil {
		handlers.unknownCommand(ctx, &ch, mapMessage(update.Message))
	}
}

func handleEditedMessage(ctx context.Context, update telego.Update) {
	if len(handlers.editHandlers) == 0 {
		return
	}

	ch := chat{update.EditedMessage.Chat.ID}

	msg := mapMessage(update.EditedMessage)
	for _, handler := range handlers.editHandlers {
		handler(ctx, &ch, msg)
	}
}

func handleCallback(ctx context.Context, update telego.Update) {
	if len(handlers.callbackHandlers) == 0 {
		return
	}

	// todo: check how to get callback without a msg
	msg, ok := update.CallbackQuery.Message.(*telego.Message)
	if !ok {
		return
	}

	ch := chat{msg.Chat.ID}

	callback := Callback{
		Data:    update.CallbackQuery.Data,
		Message: mapMessage(msg),
	}

	for _, handler := range handlers.callbackHandlers {
		handler(ctx, &ch, callback)
	}

	_ = bot.AnswerCallbackQuery(context.Background(), &telego.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})
}
