package models

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	stack   []string
}

func (err Error) Error() string {
	return err.Message
}

func (err Error) GetStackString() []string {
	return err.stack
}
func (err Error) AddStack(stack string) Error {
	err.stack = append(err.stack, stack)
	return err
}

func NewCodeError(code ErrorCode, msg string) Error {
	return Error{Code: code, Message: msg}
}
func NewGenericError(msg string) Error {
	return NewCodeError(genericError, msg)
}

type ErrorCode int

const (
	genericError ErrorCode = iota // for documentation only
	AuthTokenExpiredError
)

func (ec ErrorCode) Make() Error {
	switch ec {
	case AuthTokenExpiredError:
		return NewCodeError(ec, "auth token expired")
	}

	return NewCodeError(ec, "error in the program - say dev to the dev generic error should be made with NewCodeError")
}
