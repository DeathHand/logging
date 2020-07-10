package logging

import (
	"fmt"
	"log"
	"os"
)

const (
	LogLevelDebug    int = 1
	LogLevelInfo     int = 2
	LogLevelWarning  int = 3
	LogLevelError    int = 4
	LogLevelCritical int = 5
)

var logLevel = LogLevelDebug

// SetLevel logging level setter
func SetLevel(level int) {
	logLevel = level
}

// Logger main structure
type Logger struct {
	debugLogger    *log.Logger
	infoLogger     *log.Logger
	warningLogger  *log.Logger
	errorLogger    *log.Logger
	criticalLogger *log.Logger
}

// Logger constructor
func New(prefix string) *Logger {
	l := new(Logger)
	l.debugLogger = log.New(os.Stdout, fmt.Sprintf("[%s] DEBUG: ", prefix), log.Ldate|log.Ltime)
	l.infoLogger = log.New(os.Stdout, fmt.Sprintf("[%s] INFO: ", prefix), log.Ldate|log.Ltime)
	l.warningLogger = log.New(os.Stdout, fmt.Sprintf("[%s] WARNING: ", prefix), log.Ldate|log.Ltime)
	l.errorLogger = log.New(os.Stderr, fmt.Sprintf("[%s] ERROR: ", prefix), log.Ldate|log.Ltime)
	l.criticalLogger = log.New(os.Stderr, fmt.Sprintf("[%s] CRITICAL: ", prefix), log.Ldate|log.Ltime)
	return l
}

// Debug print DEBUG message
func (l *Logger) Debug(message string, v ...interface{}) {
	if logLevel > LogLevelDebug {
		return
	}
	l.debugLogger.Printf(message, v...)
}

// Info print INFO message
func (l *Logger) Info(message string, v ...interface{}) {
	if logLevel > LogLevelInfo {
		return
	}
	l.infoLogger.Printf(message, v...)
}

// Warning print WARNING message
func (l *Logger) Warning(message string, v ...interface{}) {
	if logLevel > LogLevelWarning {
		return
	}
	l.warningLogger.Printf(message, v...)
}

// Error print ERROR message
func (l *Logger) Error(message string, v ...interface{}) {
	if logLevel > LogLevelError {
		return
	}
	l.errorLogger.Printf(message, v...)
}

// Critical print CRITICAL message and exit
func (l *Logger) Critical(message string, v ...interface{}) {
	l.criticalLogger.Fatalf(message, v...)
}

// main logger
var main = New("main")

// Debug print DEBUG message
func Debug(message string, v ...interface{}) {
	main.Debug(message, v...)
}

// Info print INFO message
func Info(message string, v ...interface{}) {
	main.Info(message, v...)
}

// Warning print WARNING message
func Warning(message string, v ...interface{}) {
	main.Warning(message, v...)
}

// Error print ERROR message
func Error(message string, v ...interface{}) {
	main.Error(message, v...)
}

// Critical print CRITICAL message and exit
func Critical(message string, v ...interface{}) {
	main.Critical(message, v...)
}
