package models

type Collection struct {
	User  Email  `json:"-" bson:"user"`
	Name  string `json:"name" bson:"name"`
	Forms []Form `json:"forms" bson:"forms"`
}
