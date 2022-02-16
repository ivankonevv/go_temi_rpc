package models

import (
	"fmt"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email          string `bson:"email"`
	FirstName      string `bson:"first_name"`
	LastName       string `bson:"last_name"`
	HashedPassword string `bson:"hashed_password"`
	Role           string `bson:"role"`
}

type GetUser struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Role      string             `bson:"role"`
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

func (user *User) Validate(password string) (bool, error) {
	if len(user.Email) < 8 && len(password) < 8 {
		return false, fmt.Errorf("email/password must be at least 8 characters")
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return false, fmt.Errorf("email is not valid")
	}
	return true, nil
}
