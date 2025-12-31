package router

import (
	"fmt"

	"github.com/ansufw/go-mazer/apps/web/config"
	"github.com/ansufw/go-mazer/views/pages"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Static("assets", "./views/assets")

	// Add favicon route for debugging
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		fmt.Println("DEBUG: Favicon.ico route accessed")
		return c.Status(404).SendString("Favicon not found")
	})

	app.Get("/", func(c *fiber.Ctx) error {

		sidebar, err := config.LoadSidebar()
		if err != nil {
			return c.Status(500).SendString("Failed to load sidebar")
		}

		c.Set("Content-Type", "text/html")
		return pages.Home(sidebar).Render(c.Context(), c.Response().BodyWriter())
	}).Name("home")

}
