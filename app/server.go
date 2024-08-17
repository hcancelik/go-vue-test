package app

import (
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

var EmbedFrontend embed.FS

func StartServer(port int) {
	server := fiber.New(fiber.Config{
		AppName:               "Test",
		DisableStartupMessage: true,
		DisableKeepalive:      true,
	})

	frontend := getFrontendAssets()

	// Routes
	apiRoutes := server.Group("/api")
	apiRoutes.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World From Backend!",
		})
	})
	// Add more routes here

	server.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(frontend),
	}))

	server.Hooks().OnListen(func(listenData fiber.ListenData) error {
		if fiber.IsChild() {
			return nil
		}

		url := fmt.Sprintf("http://localhost:%s", listenData.Port)
		fmt.Printf("App is running on %s", url)

		return nil
	})

	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		StartServer(port + 1)
	}
}
