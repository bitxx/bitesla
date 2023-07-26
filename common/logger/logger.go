package logger

import (
	"github.com/bitxx/bitesla/common/logger/daterot"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var logger *logrus.Logger

// 封装logrus.Fields
type Fields logrus.Fields

func Init(isDebug bool, filename string, level logrus.Level, enableDynamic bool, isJsonFormat bool, maxAgeDays time.Duration) {

	if level != 0 {
		daterot.LogLevel = level
	}

	if maxAgeDays <= 0 {
		maxAgeDays = 7
	}
	daterot.EnableDynamic = enableDynamic
	daterot.JSONFormat = isJsonFormat
	daterot.BaseFileName = filename
	daterot.MaxAgeDays = maxAgeDays

	logger, _ = daterot.Rotate()
	// PanicLevel:0, FatalLevel:1, ErrorLevel:2, WarnLevel:3, InfoLevel:4, DebugLevel:5
	if isDebug {
		//debug状态，则控制台打印
		switch level {
		case 5:
			logger.SetOutput(os.Stderr)
		}
	}

}

// Debug
func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		logger.WithFields(logrus.Fields{}).Debug(args)
	}
}

// 带有field的Debug
func DebugWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.DebugLevel {
		logger.WithFields(logrus.Fields(f)).Debug(l)
	}
}

// Info
func Info(args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		logger.WithFields(logrus.Fields{}).Info(args...)
	}
}

// 带有field的Info
func InfoWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.InfoLevel {
		logger.WithFields(logrus.Fields(f)).Info(l)
	}
}

// Warn
func Warn(args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		logger.WithFields(logrus.Fields{}).Warn(args...)
	}
}

// 带有Field的Warn
func WarnWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.WarnLevel {
		logger.WithFields(logrus.Fields(f)).Warn(l)
	}
}

// Error
func Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		logger.WithFields(logrus.Fields{}).Error(args...)
	}
}

// 带有Fields的Error
func ErrorWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.ErrorLevel {
		logger.WithFields(logrus.Fields(f)).Error(l)
	}
}

// Fatal
func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		logger.WithFields(logrus.Fields{}).Fatal(args...)
	}
}

// 带有Field的Fatal
func FatalWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.FatalLevel {
		logger.WithFields(logrus.Fields(f)).Fatal(l)
	}
}

// Panic
func Panic(args ...interface{}) {
	if logger.Level >= logrus.PanicLevel {
		logger.WithFields(logrus.Fields{}).Panic(args...)
	}
}

// 带有Field的Panic
func PanicWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.PanicLevel {
		logger.WithFields(logrus.Fields(f)).Panic(l)
	}
}
