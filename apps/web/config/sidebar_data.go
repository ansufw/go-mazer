package config

// SidebarData contains the sidebar menu configuration
var SidebarData = []SidebarItem{
	{
		Name:    "Menu",
		IsTitle: boolPtr(true),
	},
	{
		Name:      "Dashboard",
		RouteName: "index",
		Icon:      "grid-fill",
	},
	{
		Name: "Components",
		Key:  "component",
		Icon: "stack",
		Submenu: []SidebarItem{
			{Name: "Accordion", RouteName: "accordion"},
			{Name: "Alert", RouteName: "alert"},
			{Name: "Badge", RouteName: "badge"},
			{Name: "Breadcrumb", RouteName: "breadcrumb"},
			{Name: "Button", RouteName: "button"},
			{Name: "Card", RouteName: "card"},
			{Name: "Carousel", RouteName: "carousel"},
			{Name: "Collapse", RouteName: "collapse"},
			{Name: "Dropdown", RouteName: "dropdown"},
			{Name: "List Group", RouteName: "list-group"},
			{Name: "Modal", RouteName: "modal"},
			{Name: "Navs", RouteName: "navs"},
			{Name: "Pagination", RouteName: "pagination"},
			{Name: "Placeholder", RouteName: "placeholder"},
			{Name: "Progress", RouteName: "progress"},
			{Name: "Spinner", RouteName: "spinner"},
			{Name: "Toasts", RouteName: "toasts"},
			{Name: "Tooltip", RouteName: "tooltip"},
		},
	},
	{
		Name: "Extra Components",
		Key:  "extra-component",
		Icon: "collection-fill",
		Submenu: []SidebarItem{
			{Name: "Avatar", RouteName: "avatar"},
			{Name: "Comment", RouteName: "comment"},
			{Name: "Divider", RouteName: "divider"},
			{Name: "Date Picker", RouteName: "date-picker"},
			{Name: "Flag", RouteName: "flag"},
			{Name: "Sweet Alert", RouteName: "sweetalert"},
			{Name: "Toastify", RouteName: "toastify"},
			{Name: "Rating", RouteName: "rating"},
			{Name: "GLightbox", RouteName: "glightbox"},
		},
	},
	{
		Name: "Layouts",
		Key:  "layout",
		Icon: "grid-1x2-fill",
		Submenu: []SidebarItem{
			{Name: "Default Layout", RouteName: "default-layout"},
			{Name: "1 Column", RouteName: "single-vertical"},
			{Name: "Vertical Navbar", RouteName: "vertical-navbar"},
			{Name: "RTL Layout", RouteName: "rtl-layout"},
			{Name: "Horizontal Menu", RouteName: "horizontal"},
		},
	},
	{
		Name:    "Forms & Tables",
		IsTitle: boolPtr(true),
	},
	{
		Name: "Form Elements",
		Key:  "form-element",
		Icon: "hexagon-fill",
		Submenu: []SidebarItem{
			{Name: "Input", URL: "form-element-input.html"},
			{Name: "Input Group", URL: "form-element-input-group.html"},
			{Name: "Select", URL: "form-element-select.html"},
			{Name: "Radio", URL: "form-element-radio.html"},
			{Name: "Checkbox", URL: "form-element-checkbox.html"},
			{Name: "Textarea", URL: "form-element-textarea.html"},
		},
	},
	{
		Name: "Form Layout",
		URL:  "form-layout.html",
		Icon: "file-earmark-medical-fill",
	},
	{
		Name: "Form Validation",
		Icon: "journal-check",
		Key:  "form-validation",
		Submenu: []SidebarItem{
			{Name: "Parsley", URL: "form-validation-parsley.html"},
		},
	},
	{
		Name: "Form Editor",
		Icon: "pen-fill",
		Key:  "form-editor",
		Submenu: []SidebarItem{
			{Name: "Quill", URL: "form-editor-quill.html"},
			{Name: "CKEditor", URL: "form-editor-ckeditor.html"},
			{Name: "Summernote", URL: "form-editor-summernote.html"},
			{Name: "TinyMCE", URL: "form-editor-tinymce.html"},
		},
	},
	{
		Name: "Table",
		URL:  "table.html",
		Icon: "grid-1x2-fill",
	},
	{
		Name: "Datatables",
		Key:  "table-datatable",
		Icon: "file-earmark-spreadsheet-fill",
		Submenu: []SidebarItem{
			{Name: "Datatable", URL: "table-datatable.html"},
			{Name: "Datatable (jQuery)", URL: "table-datatable-jquery.html"},
		},
	},
	{
		Name:    "Extra UI",
		IsTitle: boolPtr(true),
	},
	{
		Name: "Widgets",
		Key:  "ui-widgets",
		Icon: "pentagon-fill",
		Submenu: []SidebarItem{
			{Name: "Chatbox", URL: "ui-widgets-chatbox.html"},
			{Name: "Pricing", URL: "ui-widgets-pricing.html"},
			{Name: "To-do List", URL: "ui-widgets-todolist.html"},
		},
	},
	{
		Name: "Icons",
		Key:  "ui-icons",
		Icon: "egg-fill",
		Submenu: []SidebarItem{
			{Name: "Bootstrap Icons ", URL: "ui-icons-bootstrap-icons.html"},
			{Name: "Fontawesome", URL: "ui-icons-fontawesome.html"},
			{Name: "Dripicons", URL: "ui-icons-dripicons.html"},
		},
	},
	{
		Name: "Charts",
		Key:  "ui-chart",
		Icon: "bar-chart-fill",
		Submenu: []SidebarItem{
			{Name: "ChartJS", URL: "ui-chart-chartjs.html"},
			{Name: "Apexcharts", URL: "ui-chart-apexcharts.html"},
		},
	},
	{
		Name: "File Uploader",
		Key:  "ui-file",
		Icon: "cloud-arrow-up-fill",
		URL:  "ui-file-uploader.html",
	},
	{
		Name: "Maps",
		Key:  "ui-map",
		Icon: "map-fill",
		Submenu: []SidebarItem{
			{Name: "Google Map", URL: "ui-map-google-map.html"},
			{Name: "JS Vector Map", URL: "ui-map-jsvectormap.html"},
			{Name: "Leaflet Map", URL: "ui-map-leaflet.html"},
			{Name: "OpenLayers Map", URL: "ui-map-openlayers.html"},
		},
	},
	{
		Name: "Multi-level Menu",
		Key:  "ui-multi-level-menu",
		Icon: "three-dots",
		Submenu: []SidebarItem{
			{
				Name: "First Level",
				Key:  "ui-multi-level-menu",
				URL:  "#",
				Submenu: []SidebarItem{
					{Name: "Second Level", URL: "ui-multi-level-menu.html"},
					{Name: "Second Level Menu", URL: "#"},
				},
			},
			{
				Name: "Another Menu",
				URL:  "#",
				Submenu: []SidebarItem{
					{Name: "Second Level Menu", URL: "#"},
				},
			},
		},
	},
	{
		Name:    "Pages",
		IsTitle: boolPtr(true),
	},
	{
		Name: "Email Application",
		Key:  "application-email",
		Icon: "envelope-fill",
		URL:  "application-email.html",
	},
	{
		Name: "Chat Application",
		Key:  "application-chat",
		Icon: "chat-dots-fill",
		URL:  "application-chat.html",
	},
	{
		Name: "Photo Gallery",
		Key:  "application-gallery",
		Icon: "image-fill",
		URL:  "application-gallery.html",
	},
	{
		Name: "Checkout Page",
		Key:  "application-checkout",
		Icon: "basket-fill",
		URL:  "application-checkout.html",
	},
	{
		Name: "Account",
		Key:  "account",
		Icon: "person-circle",
		Submenu: []SidebarItem{
			{Name: "Profile", URL: "account-profile.html"},
			{Name: "Security", URL: "account-security.html"},
		},
	},
	{
		Name: "Authentication",
		Key:  "auth",
		Icon: "person-badge-fill",
		Submenu: []SidebarItem{
			{Name: "Login", URL: "auth-login.html"},
			{Name: "Register", URL: "auth-register.html"},
			{Name: "Forgot Password", URL: "auth-forgot-password.html"},
		},
	},
	{
		Name: "Errors",
		Key:  "error",
		Icon: "x-octagon-fill",
		Submenu: []SidebarItem{
			{Name: "403", URL: "error-403.html"},
			{Name: "404", URL: "error-404.html"},
			{Name: "500", URL: "error-500.html"},
		},
	},
	{
		Name:    "Raise Support",
		IsTitle: boolPtr(true),
	},
	{
		Name: "Documentation",
		Key:  "error",
		Icon: "life-preserver",
		URL:  "https://zuramai.github.io/mazer/docs",
	},
	{
		Name: "Contribute",
		Key:  "error",
		URL:  "https://github.com/zuramai/mazer/blob/main/CONTRIBUTING.md",
		Icon: "puzzle",
	},
	{
		Name: "Donate",
		Key:  "error",
		URL:  "https://github.com/sponsors/zuramai",
		Icon: "cash",
	},
}

// boolPtr returns a pointer to a bool value
func boolPtr(b bool) *bool {
	return &b
}
