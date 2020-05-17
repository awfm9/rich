package rich

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
)

type Error struct {
	err error
	fs  fields
}

func (e Error) Error() string {
	return fmt.Sprintf("%s (%s)", e.err, e.fs)
}

func Errorf(format string, a ...interface{}) Error {
	return Error{err: fmt.Errorf(format, a...)}
}

func (e Error) Log(ev *zerolog.Event) *zerolog.Event {
	for _, f := range e.fs {
		f.Log(ev)
	}
	return ev.Err(e.err)
}

func (e Error) Bool(key string, val bool) Error {
	e.fs = append(e.fs, boolField{key, val})
	return e
}

func (e Error) Int(key string, val int) Error {
	e.fs = append(e.fs, intField{key, val})
	return e
}

func (e Error) Int8(key string, val int8) Error {
	e.fs = append(e.fs, int8Field{key, val})
	return e
}

func (e Error) Int16(key string, val int16) Error {
	e.fs = append(e.fs, int16Field{key, val})
	return e
}

func (e Error) Int32(key string, val int32) Error {
	e.fs = append(e.fs, int32Field{key, val})
	return e
}

func (e Error) Int64(key string, val int64) Error {
	e.fs = append(e.fs, int64Field{key, val})
	return e
}

func (e Error) Uint(key string, val uint) Error {
	e.fs = append(e.fs, uintField{key, val})
	return e
}

func (e Error) Uint8(key string, val uint8) Error {
	e.fs = append(e.fs, uint8Field{key, val})
	return e
}

func (e Error) Uint16(key string, val uint16) Error {
	e.fs = append(e.fs, uint16Field{key, val})
	return e
}

func (e Error) Uint32(key string, val uint32) Error {
	e.fs = append(e.fs, uint32Field{key, val})
	return e
}

func (e Error) Uint64(key string, val uint64) Error {
	e.fs = append(e.fs, uint64Field{key, val})
	return e
}

func (e Error) Float32(key string, val float32) Error {
	e.fs = append(e.fs, float32Field{key, val})
	return e
}

func (e Error) Float64(key string, val float64) Error {
	e.fs = append(e.fs, float64Field{key, val})
	return e
}

func (e Error) Bytes(key string, val []byte) Error {
	e.fs = append(e.fs, bytesField{key, val})
	return e
}

func (e Error) Hex(key string, val []byte) Error {
	e.fs = append(e.fs, hexField{key, val})
	return e
}

func (e Error) Str(key string, val string) Error {
	e.fs = append(e.fs, strField{key, val})
	return e
}

func (e Error) AnErr(key string, val error) Error {
	e.fs = append(e.fs, errField{key, val})
	return e
}

func (e Error) Dur(key string, val time.Duration) Error {
	e.fs = append(e.fs, durField{key, val})
	return e
}

func (e Error) Time(key string, val time.Time) Error {
	e.fs = append(e.fs, timeField{key, val})
	return e
}

func (e Error) RawJSON(key string, val []byte) Error {
	e.fs = append(e.fs, jsonField{key, val})
	return e
}

func (e Error) Interface(key string, val interface{}) Error {
	e.fs = append(e.fs, ifaceField{key, val})
	return e
}
