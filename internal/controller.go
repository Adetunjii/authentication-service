package internal

import "github.com/Adetunjii/lookapp/auth-service/internal/users"

type ControllerManager struct {
	UserController users.ControllerManager
}

func NewControllerManager(service *ServiceManager) *ControllerManager {

	userController := users.Controller{
		Service: service.UserService,
	}

	return &ControllerManager{
		UserController: &userController,
	}
}
