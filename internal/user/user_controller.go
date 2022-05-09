package user

import "github.com/gofiber/fiber/v2"

type Controller struct {
	Service Service
}

type ControllerInterface interface {
	Register(c *fiber.Ctx) error
}

func NewUserController(service Service) *Controller {
	return &Controller{Service: service}
}

func (ctrl Controller) Register(c *fiber.Ctx) error {
	input := CreateUserDto{}

	if err := c.BodyParser(&input); err != nil {
		panic(err)
	}

	result, err := ctrl.Service.CreateNewUser(input)
	if err != nil {
		panic(err)
	}

	return c.JSON(result)
}
