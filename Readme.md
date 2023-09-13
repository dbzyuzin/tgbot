# Tgbot

Пакет предоставляющий минимальный набор методов для работы с ботами в телеграмм.  
Задача пакета дать возможность создавать ботов тем кто только учит язык.

### Как добавить библиотеку к себе
```shell
go get github.com/dbzyuzin/tgbot@latestSS
```

### Самой простой бот
```go
package main

import (
	"github.com/dbzyuzin/tgbot"
)

func main() {
	tgbot.RegisterHandler(func(msg tgbot.Message) {
		tgbot.SendMessage(msg.ChatID, msg.Text)
	})

	tgbot.Start()
}
```