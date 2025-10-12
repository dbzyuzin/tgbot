package tgbot

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/mymmrac/telego"
)

var webhookPath = "/api/bot/webhook"
var ErrNoAppUrl = errors.New("no app url provided")

func SetWebhook(ctx context.Context) error {
	if bot == nil {
		return errors.New("bot is nil")
	}

	if cfg().AppURL == "" {
		return ErrNoAppUrl
	}

	wh, err := bot.GetWebhookInfo(ctx)
	if err != nil {
		return fmt.Errorf("get webhook error: %w", err)
	}

	url := cfg().AppURL + webhookPath

	if wh.URL == url {
		slog.Info("webhook already set", "url", url)
		return nil
	}

	err = bot.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:            url,
		SecretToken:    bot.SecretToken(),
		AllowedUpdates: allowedUpdates,
	})
	if err != nil {
		return fmt.Errorf("set webhook error: %w", err)
	}

	slog.Info("webhook set", "url", url)

	return nil
}

func StartWebhookServer(ctx context.Context, mux *http.ServeMux) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg().ServerPort),
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := server.Shutdown(ctx); err != nil {
			log.Println("Webhook server shutdown error:", err)
		}
	}()

	slog.Info("starting webhook server", "address", server.Addr)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("webhook server error", "error", err)
		return err
	}

	return nil
}
