package logger

import (
	"log"
	"os"
)

var fatalPrefix string = "fatal - "
var debugPrefix string = "debug - "
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

// Debug logger debug logs
func Debug(format string, v ...interface{}) {
	logger := NewLogger(debugPrefix)
	logger.Printf(format, v...)
}

// Fatal logger fatal logs and exit
func Fatal(format string, v ...interface{}) {
	logger := NewLogger(fatalPrefix)
	logger.Fatalf(format, v...)
}

// FatalInfo logger fatal logs
func FatalInfo(format string, v ...interface{}) {
	logger := NewLogger(fatalPrefix)
	logger.Printf(format, v...)
}
