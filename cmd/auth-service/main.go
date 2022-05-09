package main

import (
	"fmt"
	"github.com/Adetunjii/go-microservices/auth-service/db"
	"github.com/Adetunjii/go-microservices/auth-service/db/repository"
	"github.com/Adetunjii/go-microservices/auth-service/internal"
	"github.com/Adetunjii/go-microservices/auth-service/utils"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	logger            *utils.Logger
	fiberLoggerConfig = fiberLogger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path} \n",
	}
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

	//middlewares

	app.Use(fiberRecover.New()) //global error handler
	app.Use(fiberLogger.New(fiberLoggerConfig))
	createRoutes(app, controllerManager)
	startServer(app, config.Port)
}

func createRoutes(fiberApp *fiber.App, manager *internal.ControllerManager) {
	apiRouter := fiberApp.Group("api")

	apiRouter.Post("/register", manager.UserController.Register)
}

func startServer(fiberApp *fiber.App, port string) {

	logger.Info(fmt.Sprintf("Server is running on port %s 🚀🚀", port))

	err := fiberApp.Listen(":" + port)
	if err != nil {
		logger.Fatal("error during listening", err)
	}
}
