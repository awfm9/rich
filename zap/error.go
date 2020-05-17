package rich

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type Error struct {
	err error
	fs  fields
}

func Errorf(format string, a ...interface{}) *Error {
	err := fmt.Errorf(format, a...)
	w := errors.Unwrap(err)
	var fs []zap.Field
	var r *Error
	if errors.As(w, &r) {
		fs = r.fs
	}
	return &Error{
		err: err,
		fs:  fs,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s (%s)", e.err, e.fs)
}

func (e *Error) Add(fields ...zap.Field) *Error {
	e.fs = append(e.fs, fields...)
	return e
}
