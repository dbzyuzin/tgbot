package main

import (
	"context"
	"embed"
	"fmt"
	"log/slog"

	"github.com/dbzyuzin/tgbot"
)

//go:embed webapp
var webappFiles embed.FS

func main() {
	tgbot.WebApp(webappFiles, "webapp")

	tgbot.MessageHandler(func(ctx context.Context, chat tgbot.Chat, msg tgbot.Message) {
		chat.SendText("🎲", []tgbot.Button{
			{Text: "Окей", Data: "okay-data-id"},
			{Text: "Окей", Data: "okay-data-id2"},
		})
	})

	tgbot.CallbackHandler(func(ctx context.Context, chat tgbot.Chat, callback tgbot.Callback) {
		fmt.Println(callback.Message)
		tgbot.SendMessage(callback.Message.ChatID, "Кнопка нажата: "+callback.Data)
	})

	tgbot.CommandHandler("hello", func(ctx context.Context, chat tgbot.Chat, msg tgbot.Message) {
		chat.SendText("Привет!")
	})

	tgbot.UnknownCommandHandler(func(ctx context.Context, c tgbot.Chat, m tgbot.Message) {
		slog.Info("user used unknown command", "text", m.Text)
	})

	tgbot.Start()
}
