# tgbot

Пакет для создания Telegram ботов на Go с минимальным API. Подходит для тех, кто только учит язык.

## Установка

```shell
go get github.com/dbzyuzin/tgbot@latest
```

## Переменные окружения

| Переменная | Обязательная | Описание |
|------------|--------------|----------|
| `BOT_TOKEN` | да | Токен бота от @BotFather |
| `APP_URL` | нет | URL для вебхуков (без него — long polling) |
| `WEB_APP_URL` | нет | URL Mini App (по умолчанию = APP_URL) |
| `SERVER_PORT` | нет | Порт HTTP сервера (default: 8080) |

**VSCode / терминал:**
```shell
export BOT_TOKEN=твой_токен
```

**Replit:** Tools → Secrets → New Secret

## Простой бот

```go
package main

import (
	"context"
	"github.com/dbzyuzin/tgbot"
)

func main() {
	tgbot.MessageHandler(func(ctx context.Context, chat tgbot.Chat, msg tgbot.Message) {
		chat.Send(msg.Text)
	})

	tgbot.Start()
}
```

## Обработка команд

```go
tgbot.CommandHandler("start", func(ctx context.Context, chat tgbot.Chat, msg tgbot.Message) {
	chat.Send("Привет!")
})

tgbot.UnknownCommandHandler(func(ctx context.Context, chat tgbot.Chat, msg tgbot.Message) {
	chat.Send("Неизвестная команда")
})
```

## Кнопки и колбэки

```go
tgbot.MessageHandler(func(ctx context.Context, chat tgbot.Chat, msg tgbot.Message) {
	chat.Send("Выбери:", tgbot.WithButtons(
		tgbot.Button{Text: "Да", Data: "yes"},
		tgbot.Button{Text: "Нет", Data: "no"},
	))
})

tgbot.CallbackHandler(func(ctx context.Context, chat tgbot.Chat, cb tgbot.Callback) {
	if cb.Data == "yes" {
		chat.Send("Отлично!")
	}
})
```

## Форматирование

```go
chat.Send("<b>Жирный</b> и <i>курсив</i>", tgbot.WithHTML())
chat.Send("*Жирный* и _курсив_", tgbot.WithMarkdown())

// Хелперы
chat.Send(tgbot.Bold("важно")+" текст", tgbot.WithHTML())
chat.Send(tgbot.Link("ссылка", "https://example.com"), tgbot.WithHTML())
```

## Mini App

```go
//go:embed webapp
var webappFiles embed.FS

func main() {
	tgbot.WebApp(webappFiles, "webapp")

	tgbot.MessageHandler(func(ctx context.Context, chat tgbot.Chat, msg tgbot.Message) {
		chat.Send("Открой приложение:", tgbot.WithWebAppButton("Open"))
	})

	tgbot.Start()
}
```

## API Handler

Можно подключить свой HTTP handler (gin, chi, etc.):

```go
g := gin.New()
g.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
tgbot.APIHandler(g)
```

Роуты будут доступны по `/api/*`.

## Локальная разработка Mini App

1. Запусти туннель: `npx localtunnel --port 8080`
2. Установи `APP_URL=https://xxx.loca.lt`
3. Запусти бота: `go run .`
