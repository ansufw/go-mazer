package main

import (
	"log"

	"github.com/ansufw/go-mazer/apps/web"
)

func main() {
	a, err := web.New("web")
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}

	if err := a.Run(); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}
