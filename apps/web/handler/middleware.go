package handler

import (
	"github.com/ansufw/go-mazer/apps/web/config"
	"github.com/ansufw/go-mazer/apps/web/data"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CtxTempl(appName string) fiber.Handler {

	return func(c *fiber.Ctx) (err error) {

		stringMap := map[data.TdKey]string{
			data.AppnameStrKey: appName,
		}

		td := &data.TemplData{
			Data:      make(map[data.TdKey]any),
			StringMap: stringMap,
		}

		c.Context().SetUserValue(data.TemplDataKey, td)

		return c.Next()
	}
}

func (h *Handler) WithSidebar(c *fiber.Ctx) error {

	sidebar, err := config.GetSidebar(c.App())
	if err != nil {
		return c.Status(500).SendString("Failed to load sidebar")
	}
	data.SetSidebar(c.Context(), sidebar)
	return c.Next()
}
