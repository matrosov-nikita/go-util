package util

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMW_log_panic(t *testing.T) {
	testLog := newTestLog()
	mw := NewDefaultMiddleware(testLog)

	mh := &testHandler{
		t: t,
		testHandler: func(t *testing.T, rw http.ResponseWriter, r *http.Request) {
			panic("hi, there")
		},
	}
	wrapper := mw(mh)

	request, err := http.NewRequest(http.MethodGet, "hello", nil)
	require.NoError(t, err)
	request.RequestURI = "hi, there RequestURI"
	request.Header.Set("X-Real-IP", "hello,realIP")
	request.Header.Set("X-Request-Id", "someRequestId")
	rw := &httptest.ResponseRecorder{}
	wrapper.ServeHTTP(rw, request)

	require.Len(t, testLog.addedLogs.logged, 3)
	firstDebug := testLog.addedLogs.logged[0]
	firstError := testLog.addedLogs.logged[1]
	secondDebug := testLog.addedLogs.logged[2]

	require.Len(t, firstDebug.fields, 4)
	require.Equal(t, firstDebug.level, "debug")
	require.Equal(t, firstDebug.msg, "started handling request")
	require.Equal(t, firstDebug.fields["method"], "GET")
	require.Equal(t, firstDebug.fields["request"], "hi, there RequestURI")
	require.Equal(t, firstDebug.fields["remote"], "hello,realIP")
	require.Equal(t, firstDebug.fields["request_id"], "someRequestId")

	require.Len(t, firstError.fields, 6)
	require.Equal(t, firstError.level, "error")
	require.Equal(t, firstError.msg, "panic while serving request")
	require.Equal(t, firstError.fields["method"], "GET")
	require.Equal(t, firstError.fields["request"], "hi, there RequestURI")
	require.Equal(t, firstError.fields["remote"], "hello,realIP")
	require.Contains(t, firstError.fields, "stack")
	require.Equal(t, firstError.fields["request_id"], "someRequestId")

	require.Len(t, secondDebug.fields, 6)
	require.Equal(t, secondDebug.level, "debug")
	require.Equal(t, secondDebug.msg, "completed handling request")
	require.Equal(t, secondDebug.fields["method"], "GET")
	require.Equal(t, secondDebug.fields["request"], "hi, there RequestURI")
	require.Equal(t, secondDebug.fields["remote"], "hello,realIP")
	require.Equal(t, secondDebug.fields["status"], 500)
	require.Contains(t, secondDebug.fields, "took")
	require.Equal(t, secondDebug.fields["request_id"], "someRequestId")
}

func TestMW_ok(t *testing.T) {
	testLog := newTestLog()
	mw := NewDefaultMiddleware(testLog)

	mh := &testHandler{
		t: t,
		testHandler: func(t *testing.T, rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusOK)
		},
	}
	wrapper := mw(mh)

	request, err := http.NewRequest(http.MethodGet, "hello", nil)
	require.NoError(t, err)
	request.RequestURI = "hi, there RequestURI"
	request.Header.Set("X-Real-IP", "hello,realIP")
	request.Header.Set("X-Request-Id", "someRequestId")
	rw := &httptest.ResponseRecorder{}
	wrapper.ServeHTTP(rw, request)

	require.Len(t, testLog.addedLogs.logged, 2)
	firstDebug := testLog.addedLogs.logged[0]
	secondDebug := testLog.addedLogs.logged[1]

	require.Len(t, firstDebug.fields, 4)
	require.Equal(t, firstDebug.level, "debug")
	require.Equal(t, firstDebug.msg, "started handling request")
	require.Equal(t, firstDebug.fields["method"], "GET")
	require.Equal(t, firstDebug.fields["request"], "hi, there RequestURI")
	require.Equal(t, firstDebug.fields["remote"], "hello,realIP")
	require.Equal(t, firstDebug.fields["request_id"], "someRequestId")

	require.Len(t, secondDebug.fields, 6)
	require.Equal(t, secondDebug.level, "debug")
	require.Equal(t, secondDebug.msg, "completed handling request")
	require.Equal(t, secondDebug.fields["method"], "GET")
	require.Equal(t, secondDebug.fields["request"], "hi, there RequestURI")
	require.Equal(t, secondDebug.fields["remote"], "hello,realIP")
	require.Equal(t, secondDebug.fields["status"], 200)
	require.Contains(t, secondDebug.fields, "took")
	require.Equal(t, secondDebug.fields["request_id"], "someRequestId")
}

type testHandler struct {
	t           *testing.T
	testHandler func(t *testing.T, rw http.ResponseWriter, r *http.Request)
}

func (t *testHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	t.testHandler(t.t, rw, r)
}

func newTestLog() *testLog {
	return &testLog{
		addedLogs:     &logContainer{},
		currentFields: map[string]interface{}{},
	}
}

type logged struct {
	fields Fields
	level  string
	msg    string
}

type logContainer struct {
	logged []logged
}

type testLog struct {
	currentFields Fields

	addedLogs *logContainer
}

func (t *testLog) WithField(key string, value interface{}) FieldLogger {
	currentFields := make(map[string]interface{})

	for key, value := range t.currentFields {
		currentFields[key] = value
	}
	currentFields[key] = value

	return &testLog{
		addedLogs:     t.addedLogs,
		currentFields: currentFields,
	}
}

func (t *testLog) WithFields(fields Fields) FieldLogger {
	currentFields := make(map[string]interface{})

	for key, value := range t.currentFields {
		currentFields[key] = value
	}
	for key, value := range fields {
		currentFields[key] = value
	}

	return &testLog{
		addedLogs:     t.addedLogs,
		currentFields: currentFields,
	}
}

func (t *testLog) WithError(err error) FieldLogger {
	panic("implement me")
}

func (t *testLog) Debugf(format string, args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Infof(format string, args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Printf(format string, args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Warnf(format string, args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Warningf(format string, args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Errorf(format string, args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Fatalf(format string, args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Panicf(format string, args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Debug(args ...interface{}) {
	t.addedLogs.logged = append(t.addedLogs.logged, logged{
		fields: t.currentFields,
		level:  "debug",
		msg:    args[0].(string),
	})
}

func (t *testLog) Info(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Print(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Warn(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Warning(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Error(args ...interface{}) {
	t.addedLogs.logged = append(t.addedLogs.logged, logged{
		fields: t.currentFields,
		level:  "error",
		msg:    args[0].(string),
	})
}

func (t *testLog) Fatal(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Panic(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Debugln(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Infoln(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Println(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Warnln(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Warningln(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Errorln(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Fatalln(args ...interface{}) {
	panic("implement me")
}

func (t *testLog) Panicln(args ...interface{}) {
	panic("implement me")
}
