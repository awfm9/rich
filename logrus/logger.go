package rich

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log func() *logrus.Entry
}

func Log(log func() *logrus.Entry) *Logger {
	return &Logger{log: log}
}

func (l *Logger) Err(err error) *logrus.Entry {
	en := l.log()
	r, ok := err.(*Error)
	if !ok {
		return en.WithError(err)
	}
	for _, f := range r.fs {
		f.Log(en)
	}
	return en.WithError(r.err)
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

type timeField struct {
	time time.Time
}

func (f timeField) String() string {
	return fmt.Sprintf("%s: %s", logrus.FieldKeyTime, time.RFC3339)
}

func (f timeField) Log(en *logrus.Entry) {
	en.WithTime(f.time)
}

type errField struct {
	err error
}

func (f errField) String() string {
	return fmt.Sprintf("%s: %s", logrus.ErrorKey, f.err)
}

func (f errField) Log(en *logrus.Entry) {
	en.WithError(f.err)
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
