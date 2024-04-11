package adapters

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"time"
)

type StandardLogger struct {
	log *slog.Logger
}

func NewStandardLogger(name string) StandardLogger {
	lopts := &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: replace,
	}

	return StandardLogger{
		log: slog.New(slog.NewTextHandler(os.Stdout, lopts)).With("service", name),
	}
}

func (sl StandardLogger) Info(msg string, args ...any) {
	sl.customLog(slog.LevelInfo, msg, args...)
}

func (sl StandardLogger) Error(msg string, args ...any) {
	sl.customLog(slog.LevelError, msg, args...)
}

func (sl StandardLogger) customLog(level slog.Level, msg string, args ...any) {
	if !sl.log.Enabled(context.Background(), level) {
		return
	}

	// Skip [Callers, this func, LevelCaller, Controller's Response Helper]
	var pcs [1]uintptr
	runtime.Callers(4, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(args...)

	_ = sl.log.Handler().Handle(context.Background(), r)
}

func replace(groups []string, a slog.Attr) slog.Attr {
	// Remove the project directory from the source's filename.
	if a.Key == slog.SourceKey {
		source := a.Value.Any().(*slog.Source)
		pwd, _ := os.Getwd()
		path, _ := strings.CutPrefix(source.File, pwd)
		source.File = strings.TrimLeft(path, "/")
	}

	return a
}
