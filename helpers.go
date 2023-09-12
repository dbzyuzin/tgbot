package tgbot

import (
	"fmt"
	"log"
)

func myPanic(eng, ruFormat string, params ...any) {
	fmt.Printf("\n\n"+ruFormat+"\n\n", params...)
	log.Panicln(eng)
}
