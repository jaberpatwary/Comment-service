package router

import (
	"comment/src/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB) {

	UserService := service.NewUserService(db)

	v1 := app.Group("/v1")

	UserRoutes(v1, UserService)

	// TODO: add another routes here...

}
