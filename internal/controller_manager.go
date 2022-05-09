package internal

import "github.com/Adetunjii/go-microservices/auth-service/internal/user"

type ControllerManager struct {
	UserController user.ControllerInterface
}

func NewControllerManager(serviceManager *ServiceManager) *ControllerManager {

	userController := user.NewUserController(serviceManager.UserService)

	return &ControllerManager{
		UserController: userController,
	}
}
