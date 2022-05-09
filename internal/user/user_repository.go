package user

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	*mongo.Collection
}

func (r *Repository) FindByID(id string) (UserDto, error) {
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

func (r *Repository) Create(arg CreateUserDto) (UserDto, error) {
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
