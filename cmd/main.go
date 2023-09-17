package main

import (
	"github.com/dbzyuzin/tgbot"
)

func main() {
	tgbot.RegisterHandler(func(msg tgbot.Message) {
		tgbot.SendMessageWithKeyboard(msg.ChatID, "🎲", [][]tgbot.Button{
			{tgbot.Button{"Окей", "okay-data-id"}},
			{tgbot.Button{"Окей", "okay-data-id2"}},
		})
	})

	tgbot.RegisterHandler(func(callback tgbot.Callback) {
		tgbot.SendMessage(callback.Message.ChatID, "Кнопка нажата")
	})

	tgbot.Start()
}
