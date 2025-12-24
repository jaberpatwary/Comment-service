package service

import (
	models "comment/src/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(c *fiber.Ctx) (*models.User, error)
	//GetAll(c *fiber.Ctx, params *validation.QueryUser) ([]models.User, error)
	//GetByUserId(c *fiber.Ctx, id string) (*models.User, error)
	//GetByPhoneNumber(c *fiber.Ctx)
	//UpdateUser(c *fiber.Ctx, req *validation.UpdateUser2, id string) (*models.User, error)
	//DeleteUser(c *fiber.Ctx, id string) error
}

// Define methods for expert service

type userService struct {
	DB *gorm.DB
}

// DB servie init
func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}

// Create
func (s *userService) CreateUser(c *fiber.Ctx) (*models.User, error) {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
