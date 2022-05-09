package user

type UserDto struct {
	ID        string `json:"_id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Username  string `json:"username" bson:"username"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
}

type CreateUserDto struct {
	ID        string `json:"_id" bson:"_id,omitempty"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Username  string `json:"username" bson:"username"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}
