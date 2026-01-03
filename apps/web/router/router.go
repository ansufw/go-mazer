package router

import (
	"github.com/ansufw/go-mazer/apps/web/config"
	"github.com/ansufw/go-mazer/views/pages"
	"github.com/ansufw/go-mazer/views/pages/menu"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Static("assets", "./views/assets")

	app.Get("/", func(c *fiber.Ctx) error {
		sidebar, err := config.LoadSidebar()
		if err != nil {
			return c.Status(500).SendString("Failed to load sidebar")
		}

		filename := c.Path()

		c.Set("Content-Type", "text/html")
		return pages.Home(sidebar, filename).Render(c.Context(), c.Response().BodyWriter())
	}).Name("home")

	app.Get("/index", func(c *fiber.Ctx) error {
		sidebar, err := config.LoadSidebar()
		if err != nil {
			return c.Status(500).SendString("Failed to load sidebar")
		}

		filename := c.Path()

		c.Set("Content-Type", "text/html")
		return menu.Index(sidebar, filename).Render(c.Context(), c.Response().BodyWriter())
	})

}
