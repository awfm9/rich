package rich

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Error struct {
	err error
	fs  fields
}

func Errorf(format string, a ...interface{}) *Error {
	err := fmt.Errorf(format, a...)
	w := errors.Unwrap(err)
	var fs fields
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

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) WithField(key string, value string) *Error {
	e.fs = append(e.fs, ifaceField{key, value})
	return e
}

func (e *Error) WithFields(fields logrus.Fields) *Error {
	e.fs = append(e.fs, mapField{fields})
	return e
}

func (e *Error) WithContext(ctx context.Context) *Error {
	e.fs = append(e.fs, ctxField{ctx})
	return e
}
