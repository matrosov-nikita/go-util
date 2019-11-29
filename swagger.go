package util

import (
	"net/http"
	"net/http/httptest"
)

func NewTransport(h http.Handler) http.RoundTripper {
	return &handlerTransport{h: h}
}

type handlerTransport struct {
	h http.Handler
}

func (s *handlerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := httptest.NewRecorder()
	s.h.ServeHTTP(resp, req)
	return resp.Result(), nil
}
