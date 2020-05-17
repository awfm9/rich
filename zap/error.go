package rich

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type Error struct {
	err    error
	fields fields
}

func Errorf(format string, a ...interface{}) *Error {
	err := fmt.Errorf(format, a...)
	w := errors.Unwrap(err)
	var fields fields
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
	return fmt.Sprintf("%s (%s)", e.err, e.fields)
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
	args args
}

func (s *Sugared) Error() string {
	return fmt.Sprintf("%s (%s)", s.err, s.args)
}

func (s *Sugared) With(args ...interface{}) *Sugared {
	s.args = append(s.args, args...)
	return s
}
