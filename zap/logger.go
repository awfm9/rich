package rich

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
	fs []zap.Field
}

func Log(log *zap.Logger, err error) *Logger {
	fs := []zap.Field{zap.Error(err)}
	r, ok := err.(*Error)
	if ok {
		fs = append(fs, r.fs...)
	}
	return &Logger{log, fs}
}

func (l *Logger) DPanic(msg string, fields ...zap.Field) {
	l.Logger.DPanic(msg, append(l.fs, fields...)...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, append(l.fs, fields...)...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, append(l.fs, fields...)...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, append(l.fs, fields...)...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, append(l.fs, fields...)...)
}

type fields []zap.Field

func (fs fields) String() string {
	ss := make([]string, 0, len(fs))
	for _, f := range fs {
		ss = append(ss, fmt.Sprintf("%s: %v", f.Key, f.Interface))
	}
	return strings.Join(ss, ", ")
}
