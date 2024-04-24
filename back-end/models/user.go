package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (pwd Password) Match(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password)) == nil
}

type User struct {
	Pseudo string ` json:"pseudo" bson:"pseudo"`
	Name   string `json:"name" bson:"name"`
}

func (u User) MongoIDFilter() bson.M {
	return bson.M{"pseudo": u.Pseudo}
}

type AuthUser struct {
	User     `bson:",inline"`
	Password Password `json:"-" bson:"password"` // hash
}
