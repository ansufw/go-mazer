package handler

import (
	"github.com/ansufw/go-mazer/apps/web/config"
	"github.com/ansufw/go-mazer/apps/web/data"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CtxTempl() fiber.Handler {

	return func(c *fiber.Ctx) (err error) {
		td := &data.TemplData{
			Data: make(map[string]any),
		}
		c.Context().SetUserValue("template-data", td)

		return c.Next()
	}
}

func (h *Handler) WithSidebar(c *fiber.Ctx) error {

	sidebar, err := config.GetSidebar(c.App())
	if err != nil {
		return c.Status(500).SendString("Failed to load sidebar")
	}
	data.SetSidebar(c.Context(), "template-data", sidebar)
	return c.Next()
}
