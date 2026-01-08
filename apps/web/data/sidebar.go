package data

import (
	"context"
	"log"

	"github.com/ansufw/go-mazer/apps/web/config"
)

// GetSidebar returns the base URL for web endpoints
func GetSidebar(ctx context.Context) []config.SidebarItem {
	td := getTemplData(ctx, TemplDataKey)

	items, ok := td.Data[SidebarDataKey]
	if !ok {
		log.Printf("[GetSidebar] sidebar data: %v not available", SidebarDataKey)
		return []config.SidebarItem{}
	}

	return items.([]config.SidebarItem)
}

func SetSidebar(ctx context.Context, items []config.SidebarItem) {
	td, ok := getCtx[*TemplData](ctx, TemplDataKey)
	if !ok {
		log.Fatal("[SetSidebar] templ-data not available")
	}
	if td.Data == nil {
		td.Data = make(map[TdKey]any)
	}
	td.Data[SidebarDataKey] = items
}
