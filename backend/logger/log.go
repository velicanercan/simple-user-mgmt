package logger

import (
	"github.com/sirupsen/logrus"
)

func InitializeLogger() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("Logger initialized, using default stderr for logging...")
}

// Log() logs stuff
func Log(level, message string, additionalInfo ...string) {
	switch level {
	case "info":
		logrus.Info(message, additionalInfo)
	case "error":
		logrus.Error(message, additionalInfo)
	case "panic":
		logrus.Panic(message, additionalInfo)
	default:
		logrus.Info(message, additionalInfo)
	}
}
