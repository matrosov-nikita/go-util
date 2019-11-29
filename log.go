package util

import (
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

func LogWithError(log logrus.FieldLogger, err error, msg string) {
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
func LogStartStop(log logrus.FieldLogger, version interface{}, appName string) func() {
	entry := log.WithFields(logrus.Fields{
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
