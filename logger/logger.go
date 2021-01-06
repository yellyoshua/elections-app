package logger

import (
	"log"
	"os"
)

var fatalPrefix string = "fatal - "
var infoPrefix string = "info - "

// NewLogger Initialize new log.Logger
func NewLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.LstdFlags)
}

// Info logger info logs
func Info(format string, v ...interface{}) {
	logger := NewLogger(infoPrefix)
	logger.Printf(format, v...)
}

// Fatal logger fatal logs
func Fatal(format string, v ...interface{}) {
	logger := NewLogger(fatalPrefix)
	logger.Fatalf(format, v...)
}
