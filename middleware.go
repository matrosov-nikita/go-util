package util

import (
	"context"
	"encoding/json"
	"github.com/facebookgo/stack"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

type timer interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

type realClock struct{}

func (rc *realClock) Now() time.Time {
	return time.Now()
}

func (rc *realClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Middleware is a middleware handler that logs the request as it goes in and the response as it goes out.
type Middleware struct {
	// Logger is the log.Logger instance used to log messages with the Logger middleware
	Logger logrus.FieldLogger

	requestIDContextKey string

	logStarting bool

	clock timer

	// Exclude URLs from logging
	excludeURLs []string
}

// NewDefaultMiddleware returns a new *Middleware which writes to a given logrus addedLogs.
func NewDefaultMiddleware(logger logrus.FieldLogger, requestIDContextKey string) func(handler http.Handler) http.Handler {
	mw := &Middleware{
		Logger:              logger,
		requestIDContextKey: requestIDContextKey,
		logStarting:         true,
		clock:               &realClock{},
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mw.ServeHTTP(w, r, next.ServeHTTP)
		})
	}
}

// SetLogStarting accepts a bool to control the logging of "started handling
// request" prior to passing to the next middleware
func (m *Middleware) SetLogStarting(v bool) {
	m.logStarting = v
}

// ExcludeURL adds a new URL u to be ignored during logging. The URL u is parsed, hence the returned error
func (m *Middleware) ExcludeURL(u string) error {
	if _, err := url.Parse(u); err != nil {
		return err
	}
	m.excludeURLs = append(m.excludeURLs, u)
	return nil
}

// ExcludedURLs returns the list of excluded URLs for this middleware
func (m *Middleware) ExcludedURLs() []string {
	return m.excludeURLs
}

var contentType = http.CanonicalHeaderKey("Content-Type")

const (
	applicationJSON            = "application/json"
	applicationJSONCharsetUTF8 = applicationJSON + "; charset=utf-8"
)

type jsonError struct {
	Message  interface{} `json:",omitempty"`
	Location string      `json:",omitempty"`
}

func (m *Middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	for _, u := range m.excludeURLs {
		if r.URL.Path == u {
			next(rw, r)
			return
		}
	}

	start := m.clock.Now()
	entry := m.Logger

	// request id
	reqID := r.Header.Get("X-Request-Id")
	if reqID == "" {
		reqID = RandomID()
	}

	ctx := context.WithValue(r.Context(), m.requestIDContextKey, reqID)
	r = r.WithContext(ctx)
	entry = entry.WithField("xRequestId", reqID)

	// Try to get the real IP
	remoteAddr := r.RemoteAddr
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		remoteAddr = realIP
	}

	entry = entry.WithFields(logrus.Fields{
		"request": r.RequestURI,
		"method":  r.Method,
		"remote":  remoteAddr,
	})

	if m.logStarting {
		entry.Debug("started handling request")
	}

	decorator := NewResponseWriterDecorator(rw)
	decorator.writer.Header().Add("X-Request-Id", reqID)

	defer func() {
		if err := recover(); err != nil {
			frames := stack.Callers(3)

			entry.WithFields(logrus.Fields{"stack": frames, "err": err}).Error("panic while serving request")

			decorator.WriteHeader(http.StatusInternalServerError)
			decorator.Header().Set(contentType, applicationJSONCharsetUTF8)
			e := jsonError{Message: err, Location: frames[0].String()}
			json.NewEncoder(decorator).Encode(e)
		}

		latency := m.clock.Since(start)
		entry.WithFields(logrus.Fields{
			"status": decorator.StatusCode,
			"took":   latency,
		}).Debug("completed handling request")
	}()

	next(decorator, r)
}
