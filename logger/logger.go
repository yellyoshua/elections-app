package logger

import (
	"log"
	"os"
)

var databasePrefix string = "[DATABASE] - "
var fatalDatabasePrefix string = "[FATAL DATABASE] - "
var serverPrefix string = "[SERVER] - "
var fatalServerPrefix string = "[FATAL SERVER] - "

// NewLogger Initialize new log.Logger
func NewLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.LstdFlags)
}

// Database logger for database
func Database(format string, v ...interface{}) {
	logger := NewLogger(databasePrefix)
	logger.Printf(format, v...)
}

// DatabaseFatal fatal logger for database
func DatabaseFatal(format string, v ...interface{}) {
	logger := NewLogger(fatalDatabasePrefix)
	logger.Fatalf(format, v...)
}

// Server logger for server actions and methods
func Server(format string, v ...interface{}) {
	logger := NewLogger(serverPrefix)
	logger.Printf(format, v...)
}

// ServerFatal fatal logger for server actions and methods
func ServerFatal(format string, v ...interface{}) {
	logger := NewLogger(fatalServerPrefix)
	logger.Fatalf(format, v...)
}
