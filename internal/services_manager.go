package internal

import (
	"github.com/Adetunjii/go-microservices/auth-service/db/repository"
	"github.com/Adetunjii/go-microservices/auth-service/internal/user"
	"github.com/Adetunjii/go-microservices/auth-service/utils"
)

type ServiceManager struct {
	UserService user.Service
}

func NewServiceManager(logger utils.AppLogger, repository repository.Repository) *ServiceManager {

	userService := user.NewService(logger, repository.UserRepository)

	return &ServiceManager{
		UserService: userService,
	}
}
