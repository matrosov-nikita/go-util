package util

import (
	"net/http"
)

func NewResponseWriterDecorator(writer http.ResponseWriter) *ResponseWriterDecorator {
	return &ResponseWriterDecorator{writer: writer}
}

type ResponseWriterDecorator struct {
	writer http.ResponseWriter

	StatusCode int
}

func (d *ResponseWriterDecorator) Header() http.Header {
	return d.writer.Header()
}

func (d *ResponseWriterDecorator) Write(body []byte) (int, error) {
	return d.writer.Write(body)
}

func (d *ResponseWriterDecorator) WriteHeader(statusCode int) {
	d.StatusCode = statusCode
	d.writer.WriteHeader(statusCode)
}
