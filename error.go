package rich

import (
	"errors"
	"fmt"
	"net"
	"runtime"
	"time"

	"github.com/rs/zerolog"
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

func (e *Error) Bool(key string, val bool) *Error {
	e.fs = append(e.fs, boolField{key, val})
	return e
}

func (e *Error) Int(key string, val int) *Error {
	e.fs = append(e.fs, intField{key, val})
	return e
}

func (e *Error) Int8(key string, val int8) *Error {
	e.fs = append(e.fs, int8Field{key, val})
	return e
}

func (e *Error) Int16(key string, val int16) *Error {
	e.fs = append(e.fs, int16Field{key, val})
	return e
}

func (e *Error) Int32(key string, val int32) *Error {
	e.fs = append(e.fs, int32Field{key, val})
	return e
}

func (e *Error) Int64(key string, val int64) *Error {
	e.fs = append(e.fs, int64Field{key, val})
	return e
}

func (e *Error) Uint(key string, val uint) *Error {
	e.fs = append(e.fs, uintField{key, val})
	return e
}

func (e *Error) Uint8(key string, val uint8) *Error {
	e.fs = append(e.fs, uint8Field{key, val})
	return e
}

func (e *Error) Uint16(key string, val uint16) *Error {
	e.fs = append(e.fs, uint16Field{key, val})
	return e
}

func (e *Error) Uint32(key string, val uint32) *Error {
	e.fs = append(e.fs, uint32Field{key, val})
	return e
}

func (e *Error) Uint64(key string, val uint64) *Error {
	e.fs = append(e.fs, uint64Field{key, val})
	return e
}

func (e *Error) Float32(key string, val float32) *Error {
	e.fs = append(e.fs, float32Field{key, val})
	return e
}

func (e *Error) Float64(key string, val float64) *Error {
	e.fs = append(e.fs, float64Field{key, val})
	return e
}

func (e *Error) Str(key string, val string) *Error {
	e.fs = append(e.fs, strField{key, val})
	return e
}

func (e *Error) AnErr(key string, val error) *Error {
	e.fs = append(e.fs, errField{key, val})
	return e
}

func (e *Error) Dur(key string, val time.Duration) *Error {
	e.fs = append(e.fs, durField{key, val})
	return e
}

func (e *Error) Time(key string, val time.Time) *Error {
	e.fs = append(e.fs, timeField{key, val})
	return e
}

func (e *Error) Bools(key string, val []bool) *Error {
	e.fs = append(e.fs, boolsField{key, val})
	return e
}

func (e *Error) Ints(key string, val []int) *Error {
	e.fs = append(e.fs, intsField{key, val})
	return e
}

func (e *Error) Ints8(key string, val []int8) *Error {
	e.fs = append(e.fs, ints8Field{key, val})
	return e
}

func (e *Error) Ints16(key string, val []int16) *Error {
	e.fs = append(e.fs, ints16Field{key, val})
	return e
}

func (e *Error) Ints32(key string, val []int32) *Error {
	e.fs = append(e.fs, ints32Field{key, val})
	return e
}

func (e *Error) Ints64(key string, val []int64) *Error {
	e.fs = append(e.fs, ints64Field{key, val})
	return e
}

func (e *Error) Uints(key string, val []uint) *Error {
	e.fs = append(e.fs, uintsField{key, val})
	return e
}

func (e *Error) Uints8(key string, val []uint8) *Error {
	e.fs = append(e.fs, uints8Field{key, val})
	return e
}

func (e *Error) Uints16(key string, val []uint16) *Error {
	e.fs = append(e.fs, uints16Field{key, val})
	return e
}

func (e *Error) Uints32(key string, val []uint32) *Error {
	e.fs = append(e.fs, uints32Field{key, val})
	return e
}

func (e *Error) Uints64(key string, val []uint64) *Error {
	e.fs = append(e.fs, uints64Field{key, val})
	return e
}

func (e *Error) Floats32(key string, val []float32) *Error {
	e.fs = append(e.fs, floats32Field{key, val})
	return e
}

func (e *Error) Floats64(key string, val []float64) *Error {
	e.fs = append(e.fs, floats64Field{key, val})
	return e
}

func (e *Error) Strs(key string, val []string) *Error {
	e.fs = append(e.fs, strsField{key, val})
	return e
}

func (e *Error) Errs(key string, val []error) *Error {
	e.fs = append(e.fs, errsField{key, val})
	return e
}

func (e *Error) Durs(key string, val []time.Duration) *Error {
	e.fs = append(e.fs, dursField{key, val})
	return e
}

func (e *Error) Times(key string, val []time.Time) *Error {
	e.fs = append(e.fs, timesField{key, val})
	return e
}

func (e *Error) Bytes(key string, val []byte) *Error {
	e.fs = append(e.fs, bytesField{key, val})
	return e
}

func (e *Error) Hex(key string, val []byte) *Error {
	e.fs = append(e.fs, hexField{key, val})
	return e
}

func (e *Error) RawJSON(key string, val []byte) *Error {
	e.fs = append(e.fs, jsonField{key, val})
	return e
}

func (e *Error) IPAddr(key string, val net.IP) *Error {
	e.fs = append(e.fs, ipField{key, val})
	return e
}

func (e *Error) IPPrefix(key string, val net.IPNet) *Error {
	e.fs = append(e.fs, prefixField{key, val})
	return e
}

func (e *Error) MACAddr(key string, val net.HardwareAddr) *Error {
	e.fs = append(e.fs, macField{key, val})
	return e
}

func (e *Error) Interface(key string, val interface{}) *Error {
	e.fs = append(e.fs, ifaceField{key, val})
	return e
}

func (e *Error) Timestamp() *Error {
	e.fs = append(e.fs, tsField{zerolog.TimestampFunc()})
	return e
}

func (e *Error) TimeDiff(key string, val1 time.Time, val2 time.Time) *Error {
	e.fs = append(e.fs, diffField{key, val1, val2})
	return e
}

func (e *Error) Fields(val map[string]interface{}) *Error {
	e.fs = append(e.fs, fieldsField{val})
	return e
}

func (e *Error) Array(key string, val zerolog.LogArrayMarshaler) *Error {
	e.fs = append(e.fs, arrayField{key, val})
	return e
}

func (e *Error) Dict(key string, val *zerolog.Event) *Error {
	e.fs = append(e.fs, dictField{key, val})
	return e
}

func (e *Error) Object(key string, val zerolog.LogObjectMarshaler) *Error {
	e.fs = append(e.fs, objectField{key, val})
	return e
}

func (e *Error) EmbedObject(val zerolog.LogObjectMarshaler) *Error {
	e.fs = append(e.fs, embedField{val})
	return e
}

func (e *Error) Caller(val ...int) *Error {
	skip := zerolog.CallerSkipFrameCount
	if len(val) > 0 {
		skip = val[0] + zerolog.CallerSkipFrameCount
	}
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return e
	}
	e.fs = append(e.fs, callerField{file, line})
	return e
}
