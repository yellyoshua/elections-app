package logger

import (
	"log"
	"os"
)

var FatalPrefix string = "fatal - "
var PanicPrefix string = "panic - "
var DebugPrefix string = "debug - "
var InfoPrefix string = "info - "

// NewLogger Initialize new log.Logger
func NewLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.LstdFlags)
}

// Info logger info logs
func Info(format string, v ...interface{}) {
	logger := NewLogger(InfoPrefix)
	logger.Printf(format, v...)
}

// Debug logger debug logs
func Debug(format string, v ...interface{}) {
	logger := NewLogger(DebugPrefix)
	logger.Printf(format, v...)
}

// Fatal logger fatal logs and exit
func Fatal(format string, v ...interface{}) {
	logger := NewLogger(FatalPrefix)
	logger.Fatalf(format, v...)
}

// Panic logger log and exit
func Panic(format string, v ...interface{}) {
	logger := NewLogger(PanicPrefix)
	logger.Panicf(format, v...)
}

// Error logger fatal logs
func Error(format string, v ...interface{}) {
	logger := NewLogger(FatalPrefix)
	logger.Printf(format, v...)
}
