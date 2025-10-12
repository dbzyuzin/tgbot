package tgbot

import (
	"fmt"
	"log/slog"
)

func myPanic(eng, ruFormat string, params ...any) {
	fmt.Printf("\n\n"+ruFormat+"\n\n", params...)
	slog.Error(eng)
	panic(eng)
}
