package repository

import (
	"errors"
	"github.com/Adetunjii/go-microservices/auth-service/db"
	"github.com/Adetunjii/go-microservices/auth-service/internal/user"
	"github.com/Adetunjii/go-microservices/auth-service/utils"
)

type Repository struct {
	UserRepository user.Repository
}

func NewRepository(logger utils.AppLogger, dbInstance *db.DB) *Repository {

	if dbInstance.Name() != "CoreDB" {
		logger.Fatal("no handler for this database", errors.New("invalid database"))
	}

	userCollection := dbInstance.GetCollection("user")
	userRepository := user.NewUserRepository(userCollection)

	return &Repository{
		userRepository,
	}
}
