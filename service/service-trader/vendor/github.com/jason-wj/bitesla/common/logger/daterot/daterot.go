package daterot

import (
	"errors"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

//日志按日期进行分割，结合使用了file-rotatelogs和logrus两个库

// Rotate for Rotate log file
func Rotate() (*logrus.Logger, error) {
	// check file name. if use relative path, the link name work wrong!
	// if !strings.HasPrefix(BaseFileName, "/") {
	// 	logrus.Fatal("BaseFileName should use absolute path!\n")
	// }
	// check and mkdir dest log path
	baseName := filepath.Dir(BaseFileName)
	if _, err := os.Stat(baseName); os.IsNotExist(err) {
		os.Mkdir(baseName, 0755)
		logrus.Println(time.Now().Format(fmtStr), "mkdir log file path:", baseName)
	}

	// tansfer relative BaseFileName to absolute path
	if !filepath.IsAbs(BaseFileName) {
		absPath, err := filepath.Abs(BaseFileName)
		if err != nil {
			return nil, errors.New("get absolute path error")
		}
		BaseFileName = absPath
	}

	// rotate writer
	rotateLog, err := rotatelogs.New(
		BaseFileName+".%Y%m%d-%H%M",
		rotatelogs.WithLinkName(BaseLinkName),
		rotatelogs.WithMaxAge(24*MaxAgeDays*time.Hour),
		rotatelogs.WithRotationTime(RotateHour*time.Hour),
	)
	if err != nil {
		logrus.Printf("failed to create rotatelogs: %s\n", err)
		return nil, nil
	}
	// logger level
	logger := logrus.New()
	logger.SetLevel(LogLevel)
	if JSONFormat {
		logger.Formatter = &logrus.JSONFormatter{}
	} else {
		logger.Formatter = &logrus.TextFormatter{}
	}

	// dynamic modify log level by signal
	if EnableDynamic {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGUSR1, syscall.SIGUSR2)
		go sigModLoglevel(sig, logger)
	}

	// logger interface wrapper
	logger.Out = rotateLog
	return logger, nil
}

// accept signal to modify log level dynamicall
// USR1 to turn down log level from Panic to Debug
// USR2 to turn - up log level from Debug to Panic
func sigModLoglevel(sigCh chan os.Signal, logger *logrus.Logger) {
	for {
		select {
		case sig := <-sigCh:
			modBySig(sig, logger)
		}
	}
}

func modBySig(sig os.Signal, logger *logrus.Logger) {
	if sig == syscall.SIGUSR1 {
		level := logger.Level
		if level != PanicLevel {
			logger.SetLevel(level - 1)
		}
		logrus.Println(time.Now().Format(fmtStr), "Raise log level to:", logger.Level)

	} else if sig == syscall.SIGUSR2 {
		level := logger.Level
		if level != DebugLevel {
			logger.SetLevel(level + 1)
		}
		logrus.Println(time.Now().Format(fmtStr), "Reduce log level to:", logger.Level)
	} else {
		logrus.Println(time.Now().Format(fmtStr), "receive unknown signal:", sig)
	}
}
