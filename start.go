package tgbot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/mymmrac/telego"
)

var allowedUpdates = []string{"message", "callback_query"}

var bot *telego.Bot

// Start запускает бота и замораживает поток до ручной остановки
func Start() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var err error
	bot, err = telego.NewBot(cfg.BotToken)
	if err != nil {
		myPanic(err.Error(), "Не удалось подключиться к телеграмм, проверь токен.")
	}

	me, err := bot.GetMe(ctx)
	if err != nil {
		myPanic(err.Error(), "can't get bot info")
		return
	}

	var updates <-chan telego.Update

	if cfg.AppURL != "" {
		err = SetWebhook(ctx)
		if err != nil && !errors.Is(err, ErrNoAppUrl) {
			myPanic(err.Error(), "can't set webhook")
			return
		}
		mux := http.NewServeMux()
		updates, err = bot.UpdatesViaWebhook(ctx,
			telego.WebhookHTTPServeMux(mux, webhookPath, bot.SecretToken()))
		if err != nil {
			myPanic(err.Error(), "can't start webhook updates")
		}

		mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		go func() {
			err = StartWebhookServer(ctx, mux)
			if err != nil {
				myPanic(err.Error(), "can't start webhook server")
			}
		}()
	} else {
		info, err := bot.GetWebhookInfo(ctx)
		if err != nil {
			myPanic(err.Error(), "can't delete webhook")
		}

		if info.URL != "" {
			err = bot.DeleteWebhook(ctx, &telego.DeleteWebhookParams{})
			if err != nil {
				myPanic(err.Error(), "can't delete webhook")
			}
			slog.Info("webhook removed")
		}

		updates, err = bot.UpdatesViaLongPolling(ctx, &telego.GetUpdatesParams{
			AllowedUpdates: allowedUpdates,
		})
		if err != nil {
			myPanic(err.Error(), "can't start long polling updates")
		}
	}

	slog.Info("bot started", "name", me.Username)

	for update := range updates {
		switch {
		case update.Message != nil:
			handleMessage(update)
		case update.CallbackQuery != nil && update.CallbackQuery.Message.IsAccessible():
			handleCallback(update)
		default:
			fmt.Println("unknown update")
		}
	}
}

func handleMessage(update telego.Update) {
	if len(handlers.messageHandlers) == 0 {
		return
	}
	
	msg := mapMessage(update.Message)
	for _, handler := range handlers.messageHandlers {
		handler(msg)
	}
}

func handleCallback(update telego.Update) {
	if len(handlers.callbackHandlers) == 0 {
		return
	}

	// todo: check how to get callback without a msg
	msg, ok := update.CallbackQuery.Message.(*telego.Message)
	if !ok {
		return
	}

	callback := Callback{
		Data:    update.CallbackQuery.Data,
		Message: mapMessage(msg),
	}
	
	for _, handler := range handlers.callbackHandlers {
		handler(callback)
	}

	_ = bot.AnswerCallbackQuery(context.Background(), &telego.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})
}

func mapMessage(msg *telego.Message) Message {
	return Message{
		ID:     msg.MessageID,
		User:   mapUser(msg.From),
		ChatID: msg.Chat.ID,
		Text:   msg.Text,
	}
}

func mapUser(user *telego.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.Username,
	}
}
