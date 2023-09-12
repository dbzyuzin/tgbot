package tgbot

import (
	"os"
)

const (
	tokenEnvName = "BOT_TOKEN"
)

var cfg struct {
	BotToken string
}

func init() {
	cfg.BotToken = os.Getenv(tokenEnvName)
	if cfg.BotToken == "" {
		myPanic("empty tgbot token", "Токен для бота не установлен! Проверь имя токена, должно быть \"%s\"", tokenEnvName)
	}

}

var handlers struct {
	newMessage func(Message)
}

func RegisterHandler(handler any) {
	switch fnc := handler.(type) {
	case func(Message):
		handlers.newMessage = fnc
	default:
		myPanic("unknown handler type", "Передан не верный аргумент в функцию RegisterHandler."+
			" Должна быть функция, принимающая обновление, например сообщение.")
	}
}
