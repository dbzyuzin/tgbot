# Tgbot

Пакет предоставляющий минимальный набор методов для работы с ботами в телеграмм.  
Задача пакета дать возможность создавать ботов тем кто только учит язык.

### Как добавить библиотеку к себе
```shell
go get 
```

### Самой простой бот
```go
package main

import (
	"tgbot"
)

func main() {
	tgbot.RegisterHandler(func(msg tgbot.Message) {
		tgbot.ReplyMessage(msg.ChatID, msg.ID, "Привет!")
	})

	tgbot.Start()
}
```