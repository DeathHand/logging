package logging

import (
	"fmt"
	"os"
	"time"
)

const (
	LogLevelDebug   int = 1
	LogLevelInfo    int = 2
	LogLevelWarning int = 3
)

var logLevel = LogLevelDebug
var timeFormat = time.RFC3339Nano

// SetLevel logging level setter
func SetLevel(level int) {
	logLevel = level
}

// SetTimeFormat change time format
func SetTimeFormat(layout string) {
	timeFormat = layout
}

// Logger main structure
type Logger struct {
	prefix string
}

// Logger constructor
func New(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
	}
}

// Debug print DEBUG message
func (l *Logger) Debug(message string, v ...interface{}) {
	if logLevel > LogLevelDebug {
		return
	}
	_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: [%s] DEBUG: %s", time.Now().Format(timeFormat), l.prefix, fmt.Sprintf(message, v...)))
}

// Info print INFO message
func (l *Logger) Info(message string, v ...interface{}) {
	if logLevel > LogLevelInfo {
		return
	}
	_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: [%s] INFO: %s", time.Now().Format(timeFormat), l.prefix, fmt.Sprintf(message, v...)))
}

// Warning print WARNING message
func (l *Logger) Warning(message string, v ...interface{}) {
	if logLevel > LogLevelWarning {
		return
	}
	_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: [%s] WARNING: %s", time.Now().Format(timeFormat), l.prefix, fmt.Sprintf(message, v...)))
}

// Error print ERROR message
func (l *Logger) Error(message string, v ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf("%s: [%s] ERROR: %s", time.Now().Format(timeFormat), l.prefix, fmt.Sprintf(message, v...)))
}

// Critical print CRITICAL message and exit
func (l *Logger) Critical(message string, v ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf("%s: [%s] CRITICAL: %s", time.Now().Format(timeFormat), l.prefix, fmt.Sprintf(message, v...)))
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
