package internal

import "github.com/gofiber/fiber/v2"

func CreateRoutes(fiberApp *fiber.App, manager *ControllerManager) {
	apiRouter := fiberApp.Group("api")

	apiRouter.Get("/", manager.UserController.Greet)
}
