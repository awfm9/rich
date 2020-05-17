package rich

import (
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type Error struct {
	err    error
	fields []zap.Field
}

func Errorf(format string, a ...interface{}) *Error {
	err := fmt.Errorf(format, a...)
	w := errors.Unwrap(err)
	var fields []zap.Field
	var e *Error
	if errors.As(w, &e) {
		fields = e.fields
	}
	var s *Sugared
	if errors.As(w, &s) {
		fields = sweeten(s.args)
	}
	return &Error{
		err:    err,
		fields: fields,
	}
}

func (e *Error) Error() string {
	ss := make([]string, 0, len(e.fields))
	for _, field := range e.fields {
		ss = append(ss, fmt.Sprintf("%s: %v", field.Key, field.Interface))
	}
	msg := strings.Join(ss, ", ")
	return fmt.Sprintf("%s (%s)", e.err, msg)
}

func (e *Error) With(fields ...zap.Field) *Error {
	e.fields = append(e.fields, fields...)
	return e
}

func (e *Error) Sugar() *Sugared {
	return &Sugared{
		err:  e.err,
		args: flatten(e.fields),
	}
}

type Sugared struct {
	err  error
	args []interface{}
}

func (s *Sugared) Error() string {
	fields := sweeten(s.args)
	ss := make([]string, 0, len(fields))
	for _, field := range fields {
		ss = append(ss, fmt.Sprintf("%s: %v", field.Key, field.Interface))
	}
	msg := strings.Join(ss, ", ")
	return fmt.Sprintf("%s (%s)", s.err, msg)
}

func (s *Sugared) With(args ...interface{}) *Sugared {
	s.args = append(s.args, args...)
	return s
}
