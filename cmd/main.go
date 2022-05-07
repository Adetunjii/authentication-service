package main

import (
	"github.com/Adetunjii/lookapp/auth-service/utils"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

var logger *utils.Logger

func init() {
	logger = utils.NewLogger()
}

func main() {
	app := fiber.New()

	startServer(app)
}

func startServer(fiberApp *fiber.App) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		logger.Fatal("could not read config file", err)
	}

	log.Printf("Server is running on port %s 🚀🚀", config.Port)

	err = fiberApp.Listen(":" + config.Port)
	if err != nil {
		// @TODO: log error
		os.Exit(1)
	}
}
