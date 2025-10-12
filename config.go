package tgbot

import (
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/caarlos0/env/v10"
)

type config struct {
	BotToken   string `env:"BOT_TOKEN,required"`
	AppURL     string `env:"APP_URL"`
	ServerPort int    `env:"SERVER_PORT"`
}

var (
	cfgInstance config
	once        sync.Once
)

func cfg() config {
	once.Do(func() {
		if err := env.Parse(&cfgInstance); err != nil {
			slog.Error("Failed to parse environment variables", "error", err)
			os.Exit(1)
		}
		cfgInstance.AppURL = strings.TrimSuffix(cfgInstance.AppURL, "/")
	})

	return cfgInstance
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
