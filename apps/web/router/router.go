package router

import (
	"github.com/ansufw/go-mazer/apps/web/handler"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, h *handler.Handler) {
	app.Static("assets", "./views/assets")

	app.Get("/", h.Home).Name("home")
	app.Get("/index", h.Index).Name("index")
	app.Get("/component-accordion", h.Accordion).Name("accordion")

	app.Get("/extra-component-avatar", h.Avatar).Name("avatar")

	app.Get("/single-vertical", h.SingleVertical).Name("single-vertical")

}
