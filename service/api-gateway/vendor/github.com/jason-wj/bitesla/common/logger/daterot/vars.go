package daterot

import (
	"time"

	"github.com/sirupsen/logrus"
)

// log level
const (
	// PanicLevel panic
	PanicLevel = logrus.PanicLevel
	// FatalLevel fatal
	FatalLevel = logrus.FatalLevel
	// ErrorLevel error
	ErrorLevel = logrus.ErrorLevel
	// WarnLevel warning
	WarnLevel = logrus.WarnLevel
	// InfoLevel info
	InfoLevel = logrus.InfoLevel
	// DebugLevel debug
	DebugLevel = logrus.DebugLevel

	fmtStr = "2006-01-02 15:04:05"
)

// RotLogger used for user to define a global logger
// because of logrus.Logger is struct,
// `Logger logrus.Logger` var is not a type
// encapsulute logrus.Logger as a struct for daterot package
type RotLogger struct {
	Logger *logrus.Logger
}

var (
	// LoggerPtr is struct for user
	//LoggerPtr = new(RotLogger)
	//LoggerPtr = RotLogger{}

	// BaseFileName : base file name of log
	BaseFileName = "./logs/current.log"
	// BaseLinkName : base link file name of log
	BaseLinkName = BaseFileName
	// MaxAgeDays : max days before log file to be purged in file system
	MaxAgeDays time.Duration = 7
	// RotateHour : hour for each rotate
	RotateHour time.Duration = 24
	// LogLevel : log level
	LogLevel = DebugLevel

	// EnableDynamic : enable dynamic modify log level by os Signal
	EnableDynamic = true
	// JSONFormat : use text format by default, true for json
	JSONFormat = false
)
