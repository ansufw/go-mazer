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
			{Name: "Navbar", RouteName: "navbar"},
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
			{Name: "Input", RouteName: "form-element-input"},
			{Name: "Input Group", RouteName: "form-element-input-group"},
			{Name: "Select", RouteName: "form-element-select"},
			{Name: "Radio", RouteName: "form-element-radio"},
			{Name: "Checkbox", RouteName: "form-element-checkbox"},
			{Name: "Textarea", RouteName: "form-element-textarea"},
		},
	},
	{
		Name: "Form Layout",
		URL:  "form-layout",
		Icon: "file-earmark-medical-fill",
	},
	{
		Name: "Form Validation",
		Icon: "journal-check",
		Key:  "form-validation",
		Submenu: []SidebarItem{
			{Name: "Parsley", URL: "form-validation-parsley"},
		},
	},
	{
		Name: "Form Editor",
		Icon: "pen-fill",
		Key:  "form-editor",
		Submenu: []SidebarItem{
			{Name: "Quill", RouteName: "form-editor-quill"},
			{Name: "CKEditor", RouteName: "form-editor-ckeditor"},
			{Name: "Summernote", RouteName: "form-editor-summernote"},
			{Name: "TinyMCE", RouteName: "form-editor-tinymce"},
		},
	},
	{
		Name: "Table",
		URL:  "table",
		Icon: "grid-1x2-fill",
	},
	{
		Name: "Datatables",
		Key:  "table-datatable",
		Icon: "file-earmark-spreadsheet-fill",
		Submenu: []SidebarItem{
			{Name: "Datatable", RouteName: "table-datatable"},
			{Name: "Datatable (jQuery)", RouteName: "table-datatable-jquery"},
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
			{Name: "Chatbox", RouteName: "ui-widgets-chatbox"},
			{Name: "Pricing", RouteName: "ui-widgets-pricing"},
			{Name: "To-do List", RouteName: "ui-widgets-todolist"},
		},
	},
	{
		Name: "Icons",
		Key:  "ui-icons",
		Icon: "egg-fill",
		Submenu: []SidebarItem{
			{Name: "Bootstrap Icons ", RouteName: "ui-icons-bootstrap-icons"},
			{Name: "Fontawesome", RouteName: "ui-icons-fontawesome"},
			{Name: "Dripicons", RouteName: "ui-icons-dripicons"},
		},
	},
	{
		Name: "Charts",
		Key:  "ui-chart",
		Icon: "bar-chart-fill",
		Submenu: []SidebarItem{
			{Name: "ChartJS", RouteName: "ui-chart-chartjs"},
			{Name: "Apexcharts", RouteName: "ui-chart-apexcharts"},
		},
	},
	{
		Name:      "File Uploader",
		Key:       "ui-file",
		Icon:      "cloud-arrow-up-fill",
		RouteName: "ui-file-uploader",
	},
	{
		Name: "Maps",
		Key:  "ui-map",
		Icon: "map-fill",
		Submenu: []SidebarItem{
			{Name: "Google Map", RouteName: "ui-map-google-map"},
			{Name: "JS Vector Map", RouteName: "ui-map-jsvectormap"},
			{Name: "Leaflet Map", RouteName: "ui-map-leaflet"},
			{Name: "OpenLayers Map", RouteName: "ui-map-openlayers"},
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
					{Name: "Second Level", RouteName: "ui-multi-level-menu"},
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
		Name:      "Email Application",
		Key:       "application-email",
		Icon:      "envelope-fill",
		RouteName: "application-email",
	},
	{
		Name:      "Chat Application",
		Key:       "application-chat",
		Icon:      "chat-dots-fill",
		RouteName: "application-chat",
	},
	{
		Name:      "Photo Gallery",
		Key:       "application-gallery",
		Icon:      "image-fill",
		RouteName: "application-gallery",
	},
	{
		Name:      "Checkout Page",
		Key:       "application-checkout",
		Icon:      "basket-fill",
		RouteName: "application-checkout",
	},
	{
		Name: "Account",
		Key:  "account",
		Icon: "person-circle",
		Submenu: []SidebarItem{
			{Name: "Profile", RouteName: "account-profile"},
			{Name: "Security", RouteName: "account-security"},
		},
	},
	{
		Name: "Authentication",
		Key:  "auth",
		Icon: "person-badge-fill",
		Submenu: []SidebarItem{
			{Name: "Login", RouteName: "auth-login"},
			{Name: "Register", RouteName: "auth-register"},
			{Name: "Forgot Password", RouteName: "auth-forgot-password"},
		},
	},
	{
		Name: "Errors",
		Key:  "error",
		Icon: "x-octagon-fill",
		Submenu: []SidebarItem{
			{Name: "403", RouteName: "error-403"},
			{Name: "404", RouteName: "error-404"},
			{Name: "500", RouteName: "error-500"},
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
