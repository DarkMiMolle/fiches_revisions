package errors

import (
	"fmt"
	"github.com/go-stack/stack"
	"strings"
)

type App struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	stack   []string
}

func (err App) Error() string {
	return err.Message
}

func (err App) GetStackString() string {
	return strings.Join(err.stack, "\n|-> ")
}
func (err App) Stack() App {
	frame := stack.Caller(1).Frame()
	err.stack = append(err.stack, fmt.Sprintf("%v:%v (%v)", frame.File, frame.Line, frame.Function))
	return err
}
