package user

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	*mongo.Collection
}

type Repository interface {
	FindByID(id string) (UserDto, error)
	Create(arg CreateUserDto) (UserDto, error)
}

func NewUserRepository(collection *mongo.Collection) *userRepository {
	return &userRepository{collection}
}

func (r *userRepository) FindByID(id string) (UserDto, error) {
	var result UserDto

	filter := bson.M{"_id": id}
	res, err := r.Find(context.TODO(), filter)
	if err != nil {
		return result, err
	}

	err = res.Decode(&result)
	if err != nil {
		return result, err
	}

	return result, err
}

func (r *userRepository) Create(arg CreateUserDto) (UserDto, error) {
	data, err := bson.Marshal(arg)

	user := UserDto{}

	if err != nil {
		return user, err
	}

	res, err := r.InsertOne(context.TODO(), data)

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		panic(errors.New("an error occurred"))
	}

	user.ID = id.Hex()
	user.Username = arg.Username
	user.Email = arg.Email
	user.FirstName = arg.FirstName
	user.LastName = arg.LastName

	return user, err
}
