package data

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

// GetPathname returns the base URL for web endpoints
func GetPathname(ctx context.Context) string {
	td := getTemplData(ctx, TemplDataKey)

	items, ok := td.StringMap[PathnameStrKey]
	if !ok {
		log.Printf("[GetPathname] pathname data: %v not available", PathnameStrKey)
		return ""
	}

	return items
}

func SetPathname(ctx *fiber.Ctx, pathname string) {
	td := getTemplData(ctx.Context(), TemplDataKey)
	td.StringMap[PathnameStrKey] = pathname
	ctx.Context().SetUserValue(TemplDataKey, td)
}
