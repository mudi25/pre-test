package errors

import (
	"fmt"
)

type errors struct {
	Code    int
	Message string
}

func (e *errors) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
func FromError(err error) (int, string) {
	if err == nil {
		return 200, "ok"
	}
	v, ok := err.(*errors)
	if !ok || v == nil {
		return 500, err.Error()
	}
	return v.Code, v.Message
}

func BadRequest() *errors {
	return &errors{Code: 400, Message: "Invalid input, make sure all inputs are correct"}
}
func IdNotFound() *errors {
	return &errors{Code: 404, Message: "id not found"}
}
func NotFound() *errors {
	return &errors{Code: 404, Message: "not found"}
}
func Internal() *errors {
	return &errors{Code: 500, Message: "the system is busy, please try again later"}
}
