package util

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const LogTimestampFormat = "2006-01-02T15:04:05.999"

// Log some common log for tests
var Log = &logrus.Logger{
	Level:     logrus.DebugLevel,
	Out:       os.Stdout,
	Formatter: &logrus.JSONFormatter{TimestampFormat: LogTimestampFormat},
}

func LogWithError(log FieldLogger, err error, msg string) {
	if err != nil {
		log.WithError(err).Error(msg)
		return
	}
	log.Debug(msg)
}

// LogStartStop logs start and finish of application
// correct way to call is the following:
//
// defer util.LogStartStop(config.Log, config.Version, "arp server") ()
func LogStartStop(log FieldLogger, version interface{}, appName string) func() {
	entry := log.WithFields(Fields{
		"version": version,
		"app":     appName,
	})
	entry.Debug("application starting")

	return func() {
		err := recover()
		if err != nil {
			entry.WithField("err", err).Error("exited with error")
			panic(err)
		} else {
			entry.Debug("exit with ok")
		}
		time.Sleep(time.Millisecond)
	}
}

type Fields map[string]interface{}

// The FieldLogger interface generalizes the Entry and Logger types
type FieldLogger interface {
	WithField(key string, value interface{}) FieldLogger
	WithFields(fields Fields) FieldLogger
	WithError(err error) FieldLogger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	// IsDebugEnabled() bool
	// IsInfoEnabled() bool
	// IsWarnEnabled() bool
	// IsErrorEnabled() bool
	// IsFatalEnabled() bool
	// IsPanicEnabled() bool
}

// GetJSON returns json by input object
func GetJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		Log.Fatal(err)
	}
	return string(data)
}

// PrintJSON prints object in json to console
func PrintJSON(v interface{}) {
	fmt.Println(GetJSON(v))
}
