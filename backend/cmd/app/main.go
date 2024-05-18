package main

import (
	"context"
	"github.com/mag1c0/L0/backend/internal/app"
)

const configDir = "config"

func main() {
	ctx := context.Background()

	app.Run(ctx, configDir)
}
