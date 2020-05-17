package rich

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

func Log(log *logrus.Logger) *Logger {
	return &Logger{log}
}

func (l *Logger) WithError(err error) *logrus.Entry {
	r, ok := err.(*Error)
	if !ok {
		return l.log.WithError(err)
	}
	entry := l.log.WithError(r.err)
	for _, f := range r.fs {
		f.Log(entry)
	}
	return entry
}

type field interface {
	fmt.Stringer
	Log(en *logrus.Entry)
}

type fields []field

func (fs fields) String() string {
	ss := make([]string, 0, len(fs))
	for _, f := range fs {
		ss = append(ss, f.String())
	}
	return strings.Join(ss, ", ")
}

type ifaceField struct {
	key   string
	value interface{}
}

func (f ifaceField) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.value)
}
func (f ifaceField) Log(en *logrus.Entry) {
	en.WithField(f.key, f.value)
}

type mapField struct {
	fields logrus.Fields
}

func (f mapField) String() string {
	ss := make([]string, 0, len(f.fields))
	for key, value := range f.fields {
		ss = append(ss, fmt.Sprintf("%s: %v", key, value))
	}
	return strings.Join(ss, ", ")
}

func (f mapField) Log(en *logrus.Entry) {
	en.WithFields(f.fields)
}

type ctxField struct {
	ctx context.Context
}

func (f ctxField) String() string {
	return fmt.Sprintf("context: %v", f.ctx)
}

func (f ctxField) Log(en *logrus.Entry) {
	en.WithContext(f.ctx)
}
