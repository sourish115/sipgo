package sip

import (
	"log/slog"
	"os"
)

var (
	defLogger *slog.Logger
)

// SetDefaultLogger sets default logger that will be used withing sip package
// Must be called before any usage of library
func SetDefaultLogger(l *slog.Logger) {
	defLogger = l
}

func DefaultLogger() *slog.Logger {
	if defLogger != nil {
		return defLogger
	}
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	defLogger = slog.New(slog.NewTextHandler(os.Stdout, opts))
	return defLogger
}
