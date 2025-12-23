package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config.Load() // ⭐ এই লাইনটাই সবচেয়ে দরকার

	app := setupFiberApp()
	db := setupDatabase()
	defer closeDatabase(db)
	setupRoutes(app, db)

	address := fmt.Sprintf("%s:%d", config.AppHost, config.AppPort)

	serverErrors := make(chan error, 1)
	go startServer(app, address, serverErrors)
	handleGracefulShutdown(ctx, app, serverErrors)
}
