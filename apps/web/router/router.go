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
	app.Get("/component-breadcrumb", h.RenderPage(pages.Breadcrumb)).Name("breadcrumb")
	app.Get("/component-button", h.RenderPage(pages.Button)).Name("button")
	app.Get("/component-card", h.RenderPage(pages.Card)).Name("card")
	app.Get("/component-carousel", h.RenderPage(pages.Carousel)).Name("carousel")
	app.Get("/component-collapse", h.RenderPage(pages.Collapse)).Name("collapse")
	app.Get("/component-dropdown", h.RenderPage(pages.Dropdown)).Name("dropdown")
	app.Get("/component-list-group", h.RenderPage(pages.ListGroup)).Name("list-group")
	app.Get("/component-modal", h.RenderPage(pages.Modal)).Name("modal")
	app.Get("/component-navs-tabs", h.RenderPage(pages.NavsTabs)).Name("navs-tabs")
	app.Get("/component-navbar", h.RenderPage(pages.Navbar)).Name("navbar")
	app.Get("/component-pagination", h.RenderPage(pages.Pagination)).Name("pagination")
	app.Get("/component-pill", h.RenderPage(pages.Pill)).Name("pill")
	app.Get("/component-progress", h.RenderPage(pages.Progress)).Name("progress")
	app.Get("/component-spinner", h.RenderPage(pages.Spinner)).Name("spinner")
	app.Get("/component-toast", h.RenderPage(pages.Toast)).Name("toast")
	app.Get("/component-tooltip", h.RenderPage(pages.Tooltip)).Name("tooltip")

	app.Get("/extra-component-avatar", h.RenderPage(pages.Avatar)).Name("avatar")

	app.Get("/single-vertical", h.RenderPage(pages.SingleVertical)).Name("single-vertical")

}
