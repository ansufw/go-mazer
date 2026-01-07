package data

import (
	"context"
	"log"

	"github.com/ansufw/go-mazer/apps/web/config"
)

// TemplData holds common data that can be passed to all templates
type TemplData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
}

func GetCtx[T any](ctx context.Context, key string) (T, bool) {
	var zero T
	val := ctx.Value(key)
	if val == nil {
		return zero, false
	}

	result, ok := val.(T)
	return result, ok
}

// GetSidebar returns the base URL for web endpoints
func GetSidebar(ctx context.Context, key string) []config.SidebarItem {

	td, ok := GetCtx[*TemplData](ctx, key)
	if !ok {
		log.Fatal("[GetSidebar] templ-data not available")
	}

	if td.Data["sidebarItems"] == nil {
		return []config.SidebarItem{}
	}

	return td.Data["sidebarItems"].([]config.SidebarItem)
}

func SetSidebar(ctx context.Context, key string, items []config.SidebarItem) {
	td, ok := GetCtx[*TemplData](ctx, key)
	if !ok {
		log.Fatal("[SetSidebar] templ-data not available")
	}
	if td.Data == nil {
		td.Data = make(map[string]any)
	}
	td.Data["sidebarItems"] = items
}
