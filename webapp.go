package tgbot

import (
	"embed"
	"net/http"
)

var webappFS embed.FS
var webappRoot string
var webappEnabled bool
var apiHandler http.Handler

func WebApp(fs embed.FS, root string) {
	webappFS = fs
	webappRoot = root
	webappEnabled = true
}

func APIHandler(handler http.Handler) {
	apiHandler = handler
}
