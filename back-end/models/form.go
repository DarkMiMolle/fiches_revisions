package models

import (
	"fmt"
	"time"
)

type Level int

type Percent float64

func (pc Percent) Of(other Percent) Percent {
	return (pc * other) / 10_000
}
func (pc Percent) String() string {
	return fmt.Sprintf("%.2f %%", float64(pc))
}

type Form struct {
	Question       string    `json:"question" bson:"question"`
	Answers        []string  `json:"answers" bson:"answers"`
	Tips           []string  `json:"tips" bson:"tips"`
	Level          Level     `json:"level" bson:"level"`
	LastDateWorked time.Time `json:"lastDateWorked" bson:"lastDateWorked"`
	SuccessRate    Percent   `json:"successRate"`
}
