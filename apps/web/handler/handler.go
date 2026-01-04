package handler

import (
	"strings"

	"github.com/ansufw/go-mazer/apps/web/config"
	"github.com/ansufw/go-mazer/views/pages/menu"
	"github.com/ansufw/go-mazer/views/pages/menu/components"
	"github.com/ansufw/go-mazer/views/pages/menu/extra"
	"github.com/ansufw/go-mazer/views/pages/menu/layouts"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func (h *Handler) Home(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return c.RedirectToRoute("index", fiber.Map{}, 302)
}

func (h *Handler) Index(c *fiber.Ctx) error {
	sidebar, err := config.LoadSidebar()
	if err != nil {
		return c.Status(500).SendString("Failed to load sidebar")
	}

	filename := strings.Trim(c.Path(), "/")

	c.Set("Content-Type", "text/html")
	return menu.Index(sidebar, filename).Render(c.Context(), c.Response().BodyWriter())
}

func (h *Handler) Accordion(c *fiber.Ctx) error {
	sidebar, err := config.LoadSidebar()
	if err != nil {
		return c.Status(500).SendString("Failed to load sidebar")
	}

	filename := strings.Trim(c.Path(), "/")

	c.Set("Content-Type", "text/html")
	return components.Accordion(sidebar, filename).Render(c.Context(), c.Response().BodyWriter())
}

func (h *Handler) Avatar(c *fiber.Ctx) error {
	sidebar, err := config.LoadSidebar()
	if err != nil {
		return c.Status(500).SendString("Failed to load sidebar")
	}

	filename := strings.Trim(c.Path(), "/")

	c.Set("Content-Type", "text/html")
	return extra.Avatar(sidebar, filename).Render(c.Context(), c.Response().BodyWriter())
}

func (h *Handler) SingleVertical(c *fiber.Ctx) error {

	filename := strings.Trim(c.Path(), "/")

	c.Set("Content-Type", "text/html")
	return layouts.SingleVertical(nil, filename).Render(c.Context(), c.Response().BodyWriter())
}
