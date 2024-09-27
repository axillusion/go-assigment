package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is an interface that other loggers must implement
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Fatal(args ...interface{})
	WithFields(fields logrus.Fields) *logrus.Entry
}

// LogrusLogger is a wrapper around logrus.Logger that implements the Logger interface
type LogrusLogger struct {
	logger *logrus.Logger
}

// NewLogger creates and returns a new LogrusLogger
func NewLogger() Logger {
	log := logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.JSONFormatter{})
	return &LogrusLogger{logger: log}
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *LogrusLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.logger.WithFields(fields)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}
