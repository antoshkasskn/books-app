package logger

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"log"
	"os"
)

const (
	keyLog      = "_key_log_"
	logFilePath = "./app.log"
)

var defaultLogger *logrus.Logger

func InitLogger() {
	defaultLogger = logrus.New()
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		file, err = os.Create(logFilePath)
		if err != nil {
			log.Fatal(err)
		}
	}
	defaultLogger.Out = file
}

func FromContext(ctx context.Context) *logrus.Logger {
	if logger, ok := ctx.Value(keyLog).(*logrus.Logger); ok {
		return logger
	}
	return defaultLogger
}

func Printf(format string, v ...interface{}) {
	defaultLogger.Printf(format, v...)
}

func Println(v ...interface{}) {
	defaultLogger.Println(v...)
}

func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}
