package models

type Email string

type Password struct {
	value string
}

type User struct {
	Email    Email    `json:"email" bson:"email"`
	Name     string   `json:"name" bson:"name"`
	Password Password `json:"-" bson:"password"` // hash
}
