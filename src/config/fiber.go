package config

import "github.com/gofiber/fiber/v2"

// FiberConfig returns Fiber configuration
func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:       "Comment Service",
		CaseSensitive: true,
	}
}
