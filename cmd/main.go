package main

import (
	"fmt"

	"github.com/dbzyuzin/tgbot"
)

func main() {
	tgbot.RegisterHandler(func(msg tgbot.Message) {
		tgbot.SendMessage(msg.ChatID, "🎲", []tgbot.Button{
			{Text: "Окей", Data: "okay-data-id"},
			{Text: "Окей", Data: "okay-data-id2"},
		})
	})

	tgbot.RegisterHandler(func(callback tgbot.Callback) {
		fmt.Println(callback.Message)
		tgbot.SendMessage(callback.Message.ChatID, "Кнопка нажата: "+callback.Data)
	})

	tgbot.Start()
}
