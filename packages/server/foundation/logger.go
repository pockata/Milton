package foundation

import (
	"log/slog"
	"os"
)

type StandardLogger struct {
	log *slog.Logger
}

func NewStandardLogger(name string) StandardLogger {
	return StandardLogger{
		log: slog.New(slog.NewTextHandler(os.Stdout, nil)).With("service", name),
	}
}

func (sl StandardLogger) Info(msg string, args ...any) {
	sl.log.Info(msg, args...)
}

func (sl StandardLogger) Error(msg string, args ...any) {
	sl.log.Error(msg, args...)
}
