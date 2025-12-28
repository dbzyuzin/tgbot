# tgbot

Go-библиотека для создания Telegram ботов. Обёртка над [telego](https://github.com/mymmrac/telego) с упрощённым API для начинающих.

## Структура

```
├── start.go      # Запуск бота, long polling / webhook
├── config.go     # Конфигурация через env, регистрация хендлеров
├── router.go     # Роутинг сообщений, команд, колбэков
├── chat.go       # Интерфейс Chat для работы с чатом
├── bot.go        # Глобальные функции отправки, хелперы форматирования
├── ds.go         # Структуры данных (Message, User, Button, Callback)
├── webhook.go    # Настройка вебхуков, HTTP сервер
├── webapp.go     # Поддержка Mini Apps
├── helpers.go    # Вспомогательные функции
└── example/      # Пример использования
```

## Переменные окружения

- `BOT_TOKEN` (required) - токен бота
- `APP_URL` - URL для вебхуков (без него используется long polling)
- `WEB_APP_URL` - URL Mini App (по умолчанию = APP_URL)
- `SERVER_PORT` - порт HTTP сервера (default: 8080)

## Режимы работы

**Long polling** (локальная разработка):
- Только `BOT_TOKEN`
- HTTP сервер запускается если есть `WebApp()` или `APIHandler()`

**Webhooks** (продакшн):
- `BOT_TOKEN` + `APP_URL`
- Вебхук устанавливается на `{APP_URL}/api/bot/webhook`

## API

Хендлеры регистрируются через:
- `CommandHandler(cmd, handler)` - команды (/start, /help)
- `MessageHandler(handler)` - текстовые сообщения
- `CallbackHandler(handler)` - нажатия inline-кнопок
- `UnknownCommandHandler(handler)` - неизвестные команды
- `EditHandler(handler)` - редактирование сообщений

Отправка через `Chat` интерфейс или глобальные функции:
- `SendText/SendHTML/SendMarkdown` - отправка
- `ReplyText/ReplyHTML/ReplyMarkdown` - ответ на сообщение
- `DeleteMessage`, `UpdateKeyboard`

## Сборка

```bash
go build ./...
```
