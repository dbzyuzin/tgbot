package tgbot

import (
	"context"
	"log/slog"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/caarlos0/env/v10"
)

type config struct {
	BotToken   string `env:"BOT_TOKEN,required"`
	AppURL     string `env:"APP_URL"`
	WebAppURL  string `env:"WEB_APP_URL"`
	ServerPort int    `env:"SERVER_PORT" envDefault:"8080"`
}

var (
	cfgInstance config
	once        sync.Once
	commandRe   = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,32}$`)
)

func cfg() config {
	once.Do(func() {
		if err := env.Parse(&cfgInstance); err != nil {
			slog.Error("Failed to parse environment variables", "error", err)
			os.Exit(1)
		}

		cfgInstance.AppURL = strings.TrimSuffix(cfgInstance.AppURL, "/")
		cfgInstance.WebAppURL = strings.TrimSuffix(cfgInstance.WebAppURL, "/")
		if cfgInstance.WebAppURL == "" {
			cfgInstance.WebAppURL = cfgInstance.AppURL
		}
	})

	return cfgInstance
}

func AppURL() string {
	return cfg().AppURL
}

var handlers = struct {
	unknownCommand   func(context.Context, Chat, Message)
	commandHandlers  map[string]func(context.Context, Chat, Message)
	messageHandlers  []func(context.Context, Chat, Message)
	callbackHandlers []func(context.Context, Chat, Callback)
	editHandlers     []func(context.Context, Chat, Message)
}{
	commandHandlers: make(map[string]func(context.Context, Chat, Message)),
}

func CommandHandler(command string, handler func(context.Context, Chat, Message)) {
	if handler == nil {
		myPanic("nil handler", "Передан nil в функцию CommandHandler")
	}

	command = strings.TrimPrefix(command, "/")

	if !commandRe.MatchString(command) {
		myPanic("invalid command", "Невалидная команда: "+command+
			". Допустимы только латинские буквы, цифры, _ и - (макс. 32 символа)")
	}

	if _, ok := handlers.commandHandlers[command]; ok {
		myPanic("duplicate command", "Команда "+command+" уже зарегистрирована")
	}

	handlers.commandHandlers[command] = handler
}

func UnknownCommandHandler(handler func(context.Context, Chat, Message)) {
	if handler == nil {
		myPanic("nil handler", "Передан nil в функцию UnknownCommandHandler")
	}

	handlers.unknownCommand = handler
}

func MessageHandler(handler func(context.Context, Chat, Message)) {
	if handler == nil {
		myPanic("nil handler", "Передан nil в функцию MessageHandler")
	}

	handlers.messageHandlers = append(handlers.messageHandlers, handler)
}

func CallbackHandler(handler func(context.Context, Chat, Callback)) {
	if handler == nil {
		myPanic("nil handler", "Передан nil в функцию CallbackHandler")
	}

	handlers.callbackHandlers = append(handlers.callbackHandlers, handler)
}

func EditHandler(handler func(context.Context, Chat, Message)) {
	if handler == nil {
		myPanic("nil handler", "Передан nil в функцию EditHandler")
	}

	handlers.editHandlers = append(handlers.editHandlers, handler)
}
