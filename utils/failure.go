package utils

import (
	"fmt"
	"net/http"
)

type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (f *Failure) Error() string {
	return fmt.Sprintf("%s : %s", http.StatusText(f.Code), f.Message)
}

func BadRequest(err error) error {
	if err != nil {
		return &Failure{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return nil
}

func NotFoundError(err error) error {
	if err != nil {
		return &Failure{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
	}
	return nil
}

func UnprocessableEntity(err error) error {
	if err != nil {
		return &Failure{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}
	return nil
}

func InternalServerError(err error) error {
	if err != nil {
		return &Failure{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func GetCode(err error) int {
	if f, ok := err.(*Failure); ok {
		return f.Code
	}
	return http.StatusInternalServerError
}

func GetMessage(err error) string {
	if f, ok := err.(*Failure); ok {
		return f.Message
	}
	return err.Error()
}
