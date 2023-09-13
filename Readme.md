# Tgbot

Пакет предоставляющий минимальный набор методов для работы с ботами в телеграмм.  
Задача пакета дать возможность создавать ботов тем кто только учит язык.

### Как добавить библиотеку к себе
Скопировать комманду ниже к себе в консоль и запустить. После обновлений библиотеки нужно будет снова выполнять это действие, чтобы увидеть последние изменения.
```shell
go get github.com/dbzyuzin/tgbot@latest
```

### Передать токен для бота
Локально в vscode:
```shell
export BOT_TOKEN={ТВОЙ ТОКЕН}
```
В Replit нужно перейти слева в `Tools -> Secrets`. Справа нажать NewSecret и заполнить `BOT_TOKEN` как ключ, а в качестве значения вставить свой токен.

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