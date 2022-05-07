package main

import (
	"fmt"
	"github.com/Adetunjii/lookapp/auth-service/db"
	"github.com/Adetunjii/lookapp/auth-service/db/repository"
	"github.com/Adetunjii/lookapp/auth-service/internal"
	"github.com/Adetunjii/lookapp/auth-service/utils"
	"github.com/gofiber/fiber/v2"
)

var (
	logger *utils.Logger
)

func init() {
	logger = utils.NewLogger()
}

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		logger.Fatal("could not read config file", err)
	}

	app := fiber.New()
	mongoUri := config.MongoUri
	mongoClient := db.NewMongoClient(logger, mongoUri)
	database := mongoClient.ConnectToDatabase("CoreDB")
	repos := repository.NewRepository(logger, database)
	serviceManager := internal.NewServiceManager(logger, *repos)
	controllerManager := internal.NewControllerManager(serviceManager)

	internal.CreateRoutes(app, controllerManager)
	startServer(app, config.Port)
}

func startServer(fiberApp *fiber.App, port string) {

	logger.Info(fmt.Sprintf("Server is running on port %s 🚀🚀", port))

	err := fiberApp.Listen(":" + port)
	if err != nil {
		logger.Fatal("error during listening", err)
	}
}
