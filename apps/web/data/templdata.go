package data

import (
	"context"
	"log"
)

type CtxKey string
type TdKey string

const (
	SidebarDataKey TdKey = "sidebar_data"
	AppnameStrKey  TdKey = "appname_str"
	PathnameStrKey TdKey = "pathname_str"
)

const (
	TemplDataKey CtxKey = "template-data"
)

// TemplData holds common data that can be passed to all templates
type TemplData struct {
	StringMap map[TdKey]string
	IntMap    map[TdKey]int
	FloatMap  map[TdKey]float32
	Data      map[TdKey]any
}

func getCtx[T any](ctx context.Context, key CtxKey) (T, bool) {
	var zero T
	val := ctx.Value(key)
	if val == nil {
		return zero, false
	}

	result, ok := val.(T)
	return result, ok
}

func getTemplData(ctx context.Context, key CtxKey) *TemplData {
	td, ok := getCtx[*TemplData](ctx, key)
	if !ok {
		log.Fatal("[getTemplData] templ-data not available")
	}
	return td
}
