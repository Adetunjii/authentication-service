package users

import "github.com/gofiber/fiber/v2"

type Controller struct {
	Service Service
}

type ControllerManager interface {
	Greet(c *fiber.Ctx) error
}

func (ctrl *Controller) Greet(c *fiber.Ctx) error {
	err := c.SendString("Hello world!!")
	if err != nil {
		return err
	}
	return nil
}
