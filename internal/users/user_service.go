package users

import (
	"github.com/Adetunjii/lookapp/auth-service/db/repository"
	"github.com/Adetunjii/lookapp/auth-service/utils"
)

type Service struct {
	logger     utils.AppLogger
	repository repository.UserRepository
}

func NewService(logger utils.AppLogger, userRepository repository.UserRepository) *Service {
	return &Service{
		logger:     logger,
		repository: userRepository,
	}
}
