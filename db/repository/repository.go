package repository

import (
	"errors"
	"github.com/Adetunjii/lookapp/auth-service/db"
	"github.com/Adetunjii/lookapp/auth-service/utils"
)

type Repository struct {
	UserRepository UserRepository
}

func NewRepository(logger utils.AppLogger, dbInstance *db.DB) *Repository {

	if dbInstance.Name() != "CoreDB" {
		logger.Fatal("no handler for this database", errors.New("invalid database"))
	}

	userRepository := UserRepository{}
	userRepository.Collection = dbInstance.GetCollection("users")

	return &Repository{
		userRepository,
	}
}
