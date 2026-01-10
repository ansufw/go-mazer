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
	app.Get("/component-navs", h.RenderPage(pages.NavsTabs)).Name("navs")
	app.Get("/component-navbar", h.RenderPage(pages.Navbar)).Name("navbar")
	app.Get("/component-pagination", h.RenderPage(pages.Pagination)).Name("pagination")
	app.Get("/component-placeholder", h.RenderPage(pages.Placeholder)).Name("placeholder")
	app.Get("/component-pill", h.RenderPage(pages.Pill)).Name("pill")
	app.Get("/component-progress", h.RenderPage(pages.Progress)).Name("progress")
	app.Get("/component-spinner", h.RenderPage(pages.Spinner)).Name("spinner")
	app.Get("/component-toast", h.RenderPage(pages.Toast)).Name("toasts")
	app.Get("/component-tooltip", h.RenderPage(pages.Tooltip)).Name("tooltip")

	app.Get("/extra-component-avatar", h.RenderPage(pages.Avatar)).Name("avatar")
	app.Get("/extra-component-comment", h.RenderPage(pages.Comment)).Name("comment")
	app.Get("/extra-component-divider", h.RenderPage(pages.Divider)).Name("divider")
	app.Get("/extra-component-date-picker", h.RenderPage(pages.DatePicker)).Name("date-picker")
	app.Get("/extra-component-flag", h.RenderPage(pages.Flag)).Name("flag")
	app.Get("/extra-component-sweetalert", h.RenderPage(pages.SweetAlert)).Name("sweetalert")
	app.Get("/extra-component-toastify", h.RenderPage(pages.Toastify)).Name("toastify")
	app.Get("/extra-component-rating", h.RenderPage(pages.Rating)).Name("rating")
	app.Get("/extra-component-glightbox", h.RenderPage(pages.GLightbox)).Name("glightbox")

	app.Get("/single-vertical", h.RenderPage(pages.SingleVertical)).Name("single-vertical")
	app.Get("/single-horizontal", h.RenderPage(pages.LayoutHorizontal)).Name("horizontal")
	app.Get("/vertical-navbar", h.RenderPage(pages.LayoutVerticalNavbar)).Name("vertical-navbar")
	app.Get("/default-layout", h.RenderPage(pages.LayoutDefault)).Name("default-layout")
	app.Get("/rtl-layout", h.RenderPage(pages.LayoutRtl)).Name("rtl-layout")

	app.Get("/form-element-checkbox", h.RenderPage(pages.FormElementCheckbox)).Name("form-element-checkbox")
	app.Get("/form-element-input", h.RenderPage(pages.FormElementInput)).Name("form-element-input")
	app.Get("/form-element-input-group", h.RenderPage(pages.FormElementInputGroup)).Name("form-element-input-group")
	app.Get("/form-element-select", h.RenderPage(pages.FormElementSelect)).Name("form-element-select")
	app.Get("/form-element-radio", h.RenderPage(pages.FormElementRadio)).Name("form-element-radio")
	app.Get("/form-element-checkbox", h.RenderPage(pages.FormElementCheckbox)).Name("form-element-checkbox")
	app.Get("/form-element-textarea", h.RenderPage(pages.FormElementTextarea)).Name("form-element-textarea")

	app.Get("/form-layout", h.RenderPage(pages.FormLayout)).Name("form-layout")

	app.Get("/form-validation-parsley", h.RenderPage(pages.FormValidationParsley)).Name("form-validation-parsley")

	app.Get("/form-editor-quill", h.RenderPage(pages.FormEditorQuill)).Name("form-editor-quill")
	app.Get("/form-editor-ckeditor", h.RenderPage(pages.FormEditorCKE)).Name("form-editor-ckeditor")
	app.Get("/form-editor-summernote", h.RenderPage(pages.FormEditorSummernote)).Name("form-editor-summernote")
	app.Get("/form-editor-tinymce", h.RenderPage(pages.FormEditorTinyMCE)).Name("form-editor-tinymce")

	app.Get("/table", h.RenderPage(pages.Table)).Name("table")

	app.Get("/table-datatable", h.RenderPage(pages.TableDatatable)).Name("table-datatable")
	app.Get("/table-datatable-jquery", h.RenderPage(pages.TableDatatableJquery)).Name("table-datatable-jquery")

	app.Get("/ui-widgets-chatbox", h.RenderPage(pages.Chatbox)).Name("ui-widgets-chatbox")
	app.Get("/ui-widgets-pricing", h.RenderPage(pages.Pricing)).Name("ui-widgets-pricing")
	app.Get("/ui-widgets-todolist", h.RenderPage(pages.TodoList)).Name("ui-widgets-todolist")

	app.Get("/ui-icons-bootstrap-icons", h.RenderPage(pages.IconsBootstrapIcons)).Name("ui-icons-bootstrap-icons")
	app.Get("/ui-icons-fontawesome", h.RenderPage(pages.IconsFontawesome)).Name("ui-icons-fontawesome")
	app.Get("/ui-icons-dripicons", h.RenderPage(pages.IconsDripicons)).Name("ui-icons-dripicons")

	app.Get("/ui-chart-chartjs", h.RenderPage(pages.ChartJS)).Name("ui-chart-chartjs")
	app.Get("/ui-chart-apexcharts", h.RenderPage(pages.Apexcharts)).Name("ui-chart-apexcharts")

}
