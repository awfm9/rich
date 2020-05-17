package rich

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.Logger
}

func Log(log *zap.Logger) *Logger {
	return &Logger{log}
}

func (l *Logger) With(fields ...zap.Field) *zap.Logger {
	if len(fields) > 0 && fields[0].Type == zapcore.ErrorType {
		fields = append(fs(fields[0].Interface.(error)), fields[1:]...)
	}
	return l.log.With(fields...)
}

type SugaredLogger struct {
	log *zap.SugaredLogger
}

func Sugar(log *zap.SugaredLogger) *SugaredLogger {
	return &SugaredLogger{log}
}

func (s *SugaredLogger) With(args ...interface{}) *zap.SugaredLogger {
	if len(args) == 0 {
		return s.log.With(args...)
	}
	field, isField := args[0].(zap.Field)
	if isField && field.Type == zapcore.ErrorType {
		args = append(as(field.Interface.(error)), args[1:]...)
	}
	key, isKey := args[0].(string)
	if isKey && key == "error" && len(args) > 1 {
		err, isErr := args[1].(error)
		if isErr {
			args = append(as(err), args[2:]...)
		}
	}
	return s.log.With(args...)
}
