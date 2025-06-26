package tgbot

import (
	"log"
	"strings"

	"github.com/caarlos0/env/v10"
)

type config struct {
	BotToken   string `env:"BOT_TOKEN,required"`
	AppURL     string `env:"APP_URL"`
	ServerPort int    `env:"SERVER_PORT"`
}

var cfg config

func init() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}
	cfg.AppURL = strings.TrimSuffix(cfg.AppURL, "/")
}

var handlers struct {
	messageHandlers  []func(Message)
	callbackHandlers []func(Callback)
}

func RegisterHandler(handler any) {
	switch fnc := handler.(type) {
	case func(Message):
		handlers.messageHandlers = append(handlers.messageHandlers, fnc)
	case func(Callback):
		handlers.callbackHandlers = append(handlers.callbackHandlers, fnc)
	default:
		myPanic("unknown handler type", "Передан не верный аргумент в функцию RegisterHandler."+
			" Должна быть функция, принимающая обновление, например сообщение.")
	}
}
