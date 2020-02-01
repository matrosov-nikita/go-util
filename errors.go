package util

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

type Error struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) WriteResponse(writer http.ResponseWriter, producer runtime.Producer) {
	writer.WriteHeader(e.StatusCode)
	if err := producer.Produce(writer, e); err != nil {
		panic(err)
	}
}

func NewError(statusCode int, message string) *Error {
	return &Error{StatusCode: statusCode, Message: message}
}

func ConvertErrorToResponder(err error) middleware.Responder {
	if httpError, ok := err.(*Error); ok {
		return httpError
	}
	return NewError(http.StatusInternalServerError, err.Error())
}
