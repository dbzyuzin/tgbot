package main

/*
Для тестирования Mini App локально:
1. Запусти туннель: npx localtunnel --port 8080
2. Скопируй полученный URL (например https://xxx.loca.lt)
3. Установи переменные окружения:
   BOT_TOKEN=твой_токен
   APP_URL=https://xxx.loca.lt
   SERVER_PORT=8080
4. Запусти: go run .
5. Отправь боту /app чтобы открыть Mini App
*/

import (
	"context"
	"embed"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dbzyuzin/tgbot"
	"github.com/gin-gonic/gin"
)

//go:embed webapp
var webappFiles embed.FS

func main() {
	tgbot.WebApp(webappFiles, "webapp")

	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(gin.Recovery())

	g.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	tgbot.APIHandler(g)

	tgbot.MessageHandler(func(ctx context.Context, chat tgbot.Chat, msg tgbot.Message) {
		chat.SendText("🎲", []tgbot.Button{{Text: "Окей", Data: "okay-data-id"}},
			[]tgbot.Button{{Text: "Окей", Data: "okay-data-id2"}},
			[]tgbot.Button{tgbot.WebAppButton("Open App")},
		)
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
