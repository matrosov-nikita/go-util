package util

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

type key string

const (
	RequestIDContextKey     key = "requestID"
	OperationNameContextKey key = "operationName"
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
	// Name is the name of the application as recorded in latency metrics
	Name   string
	Before func(logrus.FieldLogger, *http.Request, string) *logrus.Entry
	After  func(logrus.FieldLogger, *ResponseWriterDecorator, time.Duration, string) *logrus.Entry

	logStarting bool

	clock timer

	// Exclude URLs from logging
	excludeURLs []string
}

func MiddlewareHandler(mw *Middleware) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mw.ServeHTTP(w, r, next.ServeHTTP)
		})
	}
}

// NewMiddlewareFromLogger returns a new *Middleware which writes to a given logrus logger.
func NewMiddlewareFromLogger(logger logrus.FieldLogger, name string) *Middleware {
	return &Middleware{
		Logger: logger,
		Name:   name,
		Before: DefaultBefore,
		After:  DefaultAfter,

		logStarting: true,
		clock:       &realClock{},
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

func (m *Middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if m.Before == nil {
		m.Before = DefaultBefore
	}

	if m.After == nil {
		m.After = DefaultAfter
	}

	for _, u := range m.excludeURLs {
		if r.URL.Path == u {
			next(rw, r)
			return
		}
	}

	start := m.clock.Now()

	// Try to get the real IP
	remoteAddr := r.RemoteAddr
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		remoteAddr = realIP
	}

	entry := m.Logger

	reqID := r.Header.Get("X-Request-Id")
	if reqID == "" {
		reqID = RandomID()
	}
	entry = entry.WithField("request_id", reqID)

	entry = m.Before(entry, r, remoteAddr)

	if m.logStarting {
		entry.Info("started handling request")
	}

	decorator := NewResponseWriterDecorator(rw)
	decorator.writer.Header().Add("X-Request-Id", reqID)
	next(decorator, r)

	latency := m.clock.Since(start)
	m.After(entry, decorator, latency, m.Name).Info("completed handling request")
}

// BeforeFunc is the func type used to modify or replace the *logrus.Entry prior
// to calling the next func in the middleware chain
type BeforeFunc func(*logrus.Entry, *http.Request, string) *logrus.Entry

// AfterFunc is the func type used to modify or replace the *logrus.Entry after
// calling the next func in the middleware chain
type AfterFunc func(*logrus.Entry, http.ResponseWriter, time.Duration, string) *logrus.Entry

// DefaultBefore is the default func assigned to *Middleware.Before
func DefaultBefore(entry logrus.FieldLogger, req *http.Request, remoteAddr string) *logrus.Entry {
	return entry.WithFields(logrus.Fields{
		"request": req.RequestURI,
		"method":  req.Method,
		"remote":  remoteAddr,
	})
}

// DefaultAfter is the default func assigned to *Middleware.After
func DefaultAfter(entry logrus.FieldLogger, res *ResponseWriterDecorator, latency time.Duration, name string) *logrus.Entry {
	return entry.WithFields(logrus.Fields{
		"status":                                res.StatusCode,
		"took":                                  latency,
		fmt.Sprintf("measure#%s.latency", name): latency.Nanoseconds(),
	})
}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := AssignRequestIDToContextIfNeeded(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AssignRequestIDToContextIfNeeded(ctx context.Context) context.Context {
	_, ok := ctx.Value(RequestIDContextKey).(string)
	if !ok {
		ctx = context.WithValue(ctx, RequestIDContextKey, RandomID())
	}
	return ctx
}

func AssignOperationNameToContext(ctx context.Context, operationName string) context.Context {
	return context.WithValue(ctx, OperationNameContextKey, operationName)
}
