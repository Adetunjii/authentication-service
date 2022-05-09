package user

import (
	"github.com/Adetunjii/go-microservices/auth-service/utils"
)

type Service interface {
	CreateNewUser(arg CreateUserDto) (UserDto, error)
}

type service struct {
	logger     utils.AppLogger
	repository Repository
}

func NewService(logger utils.AppLogger, userRepository Repository) *service {
	return &service{
		logger:     logger,
		repository: userRepository,
	}
}

func (service *service) CreateNewUser(arg CreateUserDto) (UserDto, error) {
	return service.repository.Create(arg)
}
