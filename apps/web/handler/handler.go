package handler

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/ansufw/go-mazer/views/pages"
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

func (h *Handler) RenderPage(f func(string) templ.Component) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		filename := strings.Trim(c.Path(), "/")
		return h.render(c, f(filename))
	}
}

func (h *Handler) SingleVertical(c *fiber.Ctx) error {
	filename := strings.Trim(c.Path(), "/")
	return h.render(c, pages.SingleVertical(filename))
}
