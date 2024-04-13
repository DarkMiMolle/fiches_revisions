package models

type Form struct {
	Question string   `json:"question" bson:"question"`
	Answers  []string `json:"answers" bson:"answers"`
}
