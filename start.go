package tgbot

import (
	"context"
	"errors"
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
	bot, err = telego.NewBot(cfg().BotToken)
	if err != nil {
		myPanic(err.Error(), "Не удалось подключиться к телеграмм, проверь токен.")
	}

	me, err := bot.GetMe(ctx)
	if err != nil {
		myPanic(err.Error(), "can't get bot info")
		return
	}

	var updates <-chan telego.Update

	if cfg().AppURL != "" {
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
			handleMessage(ctx, update)
		case update.CallbackQuery != nil && update.CallbackQuery.Message.IsAccessible():
			handleCallback(ctx, update)
		default:
			slog.Warn("unknown update")
		}
	}
}

func mapMessage(msg *telego.Message) Message {
	htmlText := msg.Text
	if len(msg.Entities) > 0 {
		htmlText = convertToHTML(msg.Text, msg.Entities)
	}

	var entities []MessageEntity
	for _, entity := range msg.Entities {
		entities = append(entities, MessageEntity{
			Type:   entity.Type,
			Offset: entity.Offset,
			Length: entity.Length,
			URL:    entity.URL,
		})
	}

	return Message{
		ID:       msg.MessageID,
		User:     mapUser(msg.From),
		ChatID:   msg.Chat.ID,
		Text:     msg.Text,
		HTMLText: htmlText,
		Entities: entities,
		IsTextOnly: msg.Video == nil && msg.Photo == nil && msg.Audio == nil &&
			msg.Document == nil && msg.Sticker == nil && msg.Voice == nil &&
			msg.VideoNote == nil && msg.Animation == nil && msg.Contact == nil,
	}
}

// convertToHTML конвертирует текст с entities в HTML
func convertToHTML(text string, entities []telego.MessageEntity) string {
	if len(entities) == 0 {
		return text
	}

	// Сортируем entities по offset в обратном порядке для правильной вставки тегов
	runes := []rune(text)
	result := make([]rune, len(runes))
	copy(result, runes)

	// Обрабатываем entities от конца к началу
	for i := len(entities) - 1; i >= 0; i-- {
		entity := entities[i]
		start := entity.Offset
		end := entity.Offset + entity.Length

		if start >= 0 && end <= len(result) {
			var openTag, closeTag string
			switch entity.Type {
			case "bold":
				openTag, closeTag = "<b>", "</b>"
			case "italic":
				openTag, closeTag = "<i>", "</i>"
			case "code":
				openTag, closeTag = "<code>", "</code>"
			case "pre":
				openTag, closeTag = "<pre>", "</pre>"
			case "text_link":
				openTag = `<a href="` + entity.URL + `">`
				closeTag = "</a>"
			case "spoiler":
				openTag, closeTag = "<spoiler>", "</spoiler>"
			case "url":
				openTag, closeTag = "", "" // URL уже видны как есть
			default:
				continue
			}

			// Вставляем теги
			result = append(result[:end], append([]rune(closeTag), result[end:]...)...)
			result = append(result[:start], append([]rune(openTag), result[start:]...)...)
		}
	}

	return string(result)
}

func mapUser(user *telego.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.Username,
	}
}
