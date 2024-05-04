package errors

import (
	"fmt"
	"github.com/go-stack/stack"
	"strings"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	stack   []string
}

func (err Error) Error() string {
	return err.Message
}

func (err Error) GetStackString() string {
	return strings.Join(err.stack, "\n|-> ")
}
func (err Error) Stack() Error {
	frame := stack.Caller(1).Frame()
	err.stack = append(err.stack, fmt.Sprintf("%v:%v (%v)", frame.File, frame.Line, frame.Function))
	return err
}

const AuthTokenExpiredError = 1
