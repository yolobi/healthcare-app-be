package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log Logger

const FileName = "./logs/api_v1.log"

func init() {
	lw := logrus.New()
	lw.SetFormatter(&logrus.JSONFormatter{})
	lw.SetOutput(os.Stdout)
	Log = &loggerWrapper{
		lw: lw,
	}
	file, err := os.OpenFile(FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	lw.Out = file
}

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	WithFields(args map[string]interface{}) Logger
}

type loggerWrapper struct {
	lw    *logrus.Logger
	entry *logrus.Entry
}

func (logger *loggerWrapper) Info(args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Info(args...)
}

func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Infof(format, args...)
}

func (logger *loggerWrapper) Warn(args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Warn(args...)
}

func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Warnf(format, args...)
}

func (logger *loggerWrapper) Debug(args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Debug(args...)
}

func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Debugf(format, args...)
}

func (logger *loggerWrapper) Error(args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Error(args...)
}

func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Errorf(format, args...)
}

func (logger *loggerWrapper) Fatal(args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Fatal(args...)
}

func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
	if logger.entry != nil {
		logger.entry.Error(args...)
		logger.entry = nil
		return
	}
	logger.lw.Fatalf(format, args...)
}

func (logger *loggerWrapper) WithFields(args map[string]interface{}) Logger {
	entry := logger.lw.WithFields(args)
	return &loggerWrapper{
		lw:    logger.lw,
		entry: entry,
	}
}
