package data

import (
	"context"
	"log"
)

// GetAppName returns the base URL for web endpoints
func GetAppName(ctx context.Context) string {
	td := getTemplData(ctx, TemplDataKey)

	items, ok := td.StringMap[AppnameStrKey]
	if !ok {
		log.Printf("[GetAppName] sidebar data: %v not available", AppnameStrKey)
		return ""
	}

	return items
}
