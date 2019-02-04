package daterot

import (
	"log"
	"testing"
	"time"
)

func TestRotate(t *testing.T) {
	BaseFileName = "logs/access.log"
	RotateHour = 1
	// JSONFormat = false
	LogLevel = DebugLevel
	logger, err := Rotate()
	if err != nil {
		log.Fatalf("create rotate file error:%s\n", err)
	}
	for {
		logger.Debug("it is a debug info!")
		time.Sleep(500 * time.Millisecond)
	}
}
