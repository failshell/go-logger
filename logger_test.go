package logger

import (
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestInfo(t *testing.T) {
	Info("info test")
}

func TestWarning(t *testing.T) {
	Warning("warning test")
}

func TestDebug(t *testing.T) {
	Debug("debug test")
}

func TestError(t *testing.T) {
	Error("error test")
}
