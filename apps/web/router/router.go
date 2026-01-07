package router

import (
	"github.com/ansufw/go-mazer/apps/web/handler"
	"github.com/ansufw/go-mazer/views/pages"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, h *handler.Handler) {
	app.Static("assets", "./views/assets")

	app.Use(h.WithSidebar)

	app.Get("/", h.Home).Name("home")
	app.Get("/index", h.RenderPage(pages.Index)).Name("index")
	app.Get("/component-accordion", h.RenderPage(pages.Accordion)).Name("accordion")
	app.Get("/component-alert", h.RenderPage(pages.Alert)).Name("alert")
	app.Get("/component-badge", h.RenderPage(pages.Badge)).Name("badge")

	app.Get("/extra-component-avatar", h.RenderPage(pages.Avatar)).Name("avatar")

	app.Get("/single-vertical", h.RenderPage(pages.SingleVertical)).Name("single-vertical")

}
