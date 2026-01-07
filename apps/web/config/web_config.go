package config

import (
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// SidebarItem represents a single item in the sidebar menu
type SidebarItem struct {
	Name      string        `json:"name"`
	IsTitle   *bool         `json:"isTitle,omitempty"`
	Key       string        `json:"key,omitempty"`
	Icon      string        `json:"icon,omitempty"`
	URL       string        `json:"url,omitempty"`
	RouteName string        `json:"routeName,omitempty"`
	Submenu   []SidebarItem `json:"submenu,omitempty"`
}

// IsTitleItem checks if this item is a title
func (s *SidebarItem) IsTitleItem() bool {
	return s.IsTitle != nil && *s.IsTitle
}

// HasSubmenu checks if this item has a submenu
func (s *SidebarItem) HasSubmenu() bool {
	return len(s.Submenu) > 0
}

// IsLink checks if this item is a simple link
func (s *SidebarItem) IsLink() bool {
	return s.URL != "" && !s.HasSubmenu()
}

type Sidebar struct {
	Items []SidebarItem `json:"items"`
}

var (
	sidebarCache []SidebarItem
	cacheOnce    sync.Once
)

// resolveSidebar recursively updates the URL of sidebar items based on RouteName
func resolveSidebar(app *fiber.App, items []SidebarItem) {
	for i := range items {
		if items[i].RouteName != "" {
			route := app.GetRoute(items[i].RouteName)
			// Fiber's GetRoute returns a route which might have parameters.
			// Assuming we are just getting static paths or handled ones.
			// Also checking if route exists.
			if route.Method != "" { // Check if route is valid
				// Trim leading slash to match existing convention if necessary,
				// but typically sidebar URLs in this project seem to be relative or absolute paths.
				// Based on previous files, they were like "index", "component-accordion".
				// app.GetRoute().Path returns e.g. "/index".
				// So we should maybe trim the leading slash.
				items[i].URL = strings.TrimPrefix(route.Path, "/")
			}
		}
		if items[i].HasSubmenu() {
			resolveSidebar(app, items[i].Submenu)
		}
	}
}

// GetSidebar loads sidebar configuration from embedded data with caching, resolving routes dynamically
func GetSidebar(app *fiber.App) ([]SidebarItem, error) {
	cacheOnce.Do(func() {
		// Create a deep copy or just modify SidebarData if it's safe (SidebarData IS the source).
		// Since SidebarData is a global var in this package, modifying it directly is permanent in memory for the process.
		// That's fine for caching.
		sidebarCache = SidebarData
		resolveSidebar(app, sidebarCache)
	})

	return sidebarCache, nil
}
