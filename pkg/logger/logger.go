package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func New() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{}) // JSON format for structured logs
	logger.SetOutput(os.Stdout)                  // Log to stdout
	logger.SetLevel(logrus.InfoLevel)            // Default log level

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		multiWriter := io.MultiWriter(os.Stdout, file)
		logger.SetOutput(multiWriter)
	} else {
		logger.Warn("Failed to log to file, using default stdout")
	}

	return logger
}
