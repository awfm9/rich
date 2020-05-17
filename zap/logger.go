package rich

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type Logger struct {
	log *zap.Logger
}

func Log(log *zap.Logger) *Logger {
	return &Logger{log}
}

func (l *Logger) With(fields ...zap.Field) *zap.Logger {
	return l.log.With(expand(fields)...)
}

type SugaredLogger struct {
	log *zap.SugaredLogger
}

func Sugar(log *zap.SugaredLogger) *SugaredLogger {
	return &SugaredLogger{log}
}

func (s *SugaredLogger) With(args ...interface{}) *zap.SugaredLogger {
	return s.log.With(flatten(expand(sweeten(args)))...)
}

type fields []zap.Field

func (fs fields) String() string {
	ss := make([]string, 0, len(fs))
	for _, f := range fs {
		ss = append(ss, fmt.Sprintf("%s: %v", f.Key, f.Interface))
	}
	return strings.Join(ss, ", ")
}

type args []interface{}

func (as args) String() string {
	return fields(sweeten(as)).String()
}
