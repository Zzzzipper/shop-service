package custom_error

import (
	"fmt"
)

type CustomError struct {
	ErrCode    string
	ErrMessage string
}

func New(errCode, errMessage string) error {
	return &CustomError{
		ErrCode:    errCode,
		ErrMessage: errMessage,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrCode, e.ErrMessage)
}
