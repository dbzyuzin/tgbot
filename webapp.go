package tgbot

import "embed"

var webappFS embed.FS
var webappRoot string
var webappEnabled bool

func WebApp(fs embed.FS, root string) {
	webappFS = fs
	webappRoot = root
	webappEnabled = true
}
