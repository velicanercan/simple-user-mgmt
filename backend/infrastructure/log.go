package infrastructure

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitializeLogger() error {
	logrus.SetLevel(logrus.InfoLevel)

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logFile, err := os.OpenFile("./log/simple-user-mgmt.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(logFile)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}

	return nil
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
