package internal

import (
	"github.com/Adetunjii/lookapp/auth-service/db/repository"
	"github.com/Adetunjii/lookapp/auth-service/internal/users"
	"github.com/Adetunjii/lookapp/auth-service/utils"
)

type ServiceManager struct {
	UserService users.Service
}

func NewServiceManager(logger utils.AppLogger, repository repository.Repository) *ServiceManager {

	userService := users.NewService(logger, repository.UserRepository)

	return &ServiceManager{
		UserService: *userService,
	}
}
