package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

func ConfigureLogging() *logrus.Logger {
	logger := logrus.New()
	// to print logs in json format
	logger.SetFormatter(&logrus.JSONFormatter{})
	// output to stdout
	logger.SetOutput(os.Stdout)
	// setting the level to trace so it will print all levels
	logger.SetLevel(logrus.ErrorLevel)
	return logger
}
