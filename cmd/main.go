package main

import (
	"github.com/dbzyuzin/tgbot"
)

func main() {
	tgbot.RegisterHandler(func(msg tgbot.Message) {
		tgbot.ReplyMessage(msg.ChatID, msg.ID, "🎲")
	})

	tgbot.Start()
}
