package util

import (
	"github.com/sirupsen/logrus"
	"os"
)

const LogTimestampFormat = "2006-01-02T15:04:05.999"

// Log some common log for tests
var Log = &logrus.Logger{
	Level:     logrus.DebugLevel,
	Out:       os.Stdout,
	Formatter: &logrus.JSONFormatter{TimestampFormat: LogTimestampFormat},
}