package logging

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	SetLevel(LogLevelInfo)
	logger := New("test")
	logger.Debug("Debugging test: %s", "test")
	logger.Info("Info test: %s", "test")
	logger.Warning("Warning test: %s", "test")
	logger.Error("Error test: %s", "test")
	logger.Critical("Critical test: %s", "test")
	t.Fail()
}

func TestDebug(t *testing.T) {
	SetLevel(LogLevelInfo)
	Debug("Debugging test: %s", "test")
}

func TestInfo(t *testing.T) {
	Info("Info test: %s", "test")
}

func TestWarning(t *testing.T) {
	Warning("Warning test: %s", "test")
}

func TestError(t *testing.T) {
	Error("Error test: %s", "test")
}

func TestLogger_SetTimeFormat(t *testing.T) {
	SetTimeFormat(time.RFC822Z)
	Error("Error test: %s", "test")
}

func TestCritical(t *testing.T) {
	Critical("Critical test: %s", "test")
	t.Fail()
}
