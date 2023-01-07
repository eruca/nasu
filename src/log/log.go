package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type log struct {
	Logger *logrus.Logger
}

var Log log

func Init(path string) {
	targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Printf("[Nasu-Log] Fail to open log file, filename: %s, err: %s\n", path, err.Error())
	}
	writers := []io.Writer{
		targetFile,
		os.Stderr,
	}
	textFormatter := new(logrus.TextFormatter)
	textFormatter.DisableColors = false
	textFormatter.TimestampFormat = "2006-01-02 15:04:05"
	textFormatter.FullTimestamp = true
	logger := logrus.New()
	logger.SetFormatter(textFormatter)
	logger.Out = io.MultiWriter(writers...)
	logger.SetLevel(logrus.DebugLevel)
	Log = log{Logger: logger}
}

func (l *log) GetLogger() *logrus.Logger {
	return l.Logger
}

func (l *log) Info(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *log) Warn(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *log) Error(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

func (l *log) Fatal(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}