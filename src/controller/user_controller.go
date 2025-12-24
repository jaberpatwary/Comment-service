package controller

import (
	"comment/src/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	_UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{

		_UserService: userService,
	}
}

// Create
func (u *UserController) CreateUser(c *fiber.Ctx) error {

	user, err := u._UserService.CreateUser(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
