package main

import (
	"fmt"

	"comment/src/config"
	"comment/src/database"
	"comment/src/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1️⃣ ENV check
	fmt.Println("ENV:", config.AppEnv)

	// 2️⃣ Database connection
	db := database.Connect(config.DBHost, config.DBName)
	if db != nil {
		fmt.Println("✅ DB connected")
	}

	// 3️⃣ Fiber setup
	app := fiber.New()

	// 4️⃣ Test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber is running 🚀")
	})

	// 🔥🔥🔥 5️⃣ ROUTES REGISTER (THIS WAS MISSING)
	router.Routes(app, db)

	// 6️⃣ Server start
	address := fmt.Sprintf("%s:%d", config.AppHost, config.AppPort)
	fmt.Println("🚀 Server running on", address)

	if err := app.Listen(address); err != nil {
		panic(err)
	}
}
