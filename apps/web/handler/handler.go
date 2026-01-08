package handler

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/ansufw/go-mazer/apps/web/data"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func (h *Handler) Home(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return c.RedirectToRoute("index", fiber.Map{}, 302)
}

func (h *Handler) render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func (h *Handler) RenderPage(f func(data.TemplData) templ.Component) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		filename := strings.Trim(c.Path(), "/")
		strMap := map[data.TdKey]string{data.PathnameStrKey: filename}
		td := data.TemplData{
			StringMap: strMap,
		}
		return h.render(c, f(td))
	}
}
