package router

import (
	"comment/src/controller"
	"comment/src/service"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(v1 fiber.Router, u service.UserService) {
	// Initialize the UserController with the UserService
	userController := controller.NewUserController(u)
	// Define user-related routes
	userGroup := v1.Group("/users")
	comGroup.Use(middleware.JwtConfig()) //jwt router
	userGroup.Post("/", userController.CreateUser)

}
