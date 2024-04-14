package models

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type Email string

func (email Email) IsValid() bool {
	match, err := regexp.Match(`^[\w-.]+@([\w-]+\.)+[\w-]{2,}$`, []byte(email))
	if err != nil {
		panic(err.Error())
	}
	return match
}

type Password string

func (pwd Password) Match(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password)) == nil
}

type User struct {
	Email    Email    `json:"email" bson:"email"`
	Name     string   `json:"name" bson:"name"`
	Password Password `json:"-" bson:"password"` // hash
}
