package main

import (
	"embed"
	"github.com/hcancelik/go-vue-test/app"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	app.EmbedFrontend = frontend

	port := 3000

	app.StartServer(port)
}
