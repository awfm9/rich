package rich

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type field interface {
	fmt.Stringer
	Log(ev *zerolog.Event)
}

type fields []field

func (fs fields) String() string {
	ss := make([]string, 0, len(fs))
	for _, f := range fs {
		ss = append(ss, f.String())
	}
	return strings.Join(ss, ", ")
}

type boolField struct {
	key string
	val bool
}

func (f boolField) String() string {
	return fmt.Sprintf("%s: %t", f.key, f.val)
}

func (f boolField) Log(ev *zerolog.Event) {
	ev.Bool(f.key, f.val)
}

type intField struct {
	key string
	val int
}

func (f intField) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f intField) Log(ev *zerolog.Event) {
	ev.Int(f.key, f.val)
}

type int8Field struct {
	key string
	val int8
}

func (f int8Field) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f int8Field) Log(ev *zerolog.Event) {
	ev.Int8(f.key, f.val)
}

type int16Field struct {
	key string
	val int16
}

func (f int16Field) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f int16Field) Log(ev *zerolog.Event) {
	ev.Int16(f.key, f.val)
}

type int32Field struct {
	key string
	val int32
}

func (f int32Field) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f int32Field) Log(ev *zerolog.Event) {
	ev.Int32(f.key, f.val)
}

type int64Field struct {
	key string
	val int64
}

func (f int64Field) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f int64Field) Log(ev *zerolog.Event) {
	ev.Int64(f.key, f.val)
}

type uintField struct {
	key string
	val uint
}

func (f uintField) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f uintField) Log(ev *zerolog.Event) {
	ev.Uint(f.key, f.val)
}

type uint8Field struct {
	key string
	val uint8
}

func (f uint8Field) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f uint8Field) Log(ev *zerolog.Event) {
	ev.Uint8(f.key, f.val)
}

type uint16Field struct {
	key string
	val uint16
}

func (f uint16Field) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f uint16Field) Log(ev *zerolog.Event) {
	ev.Uint16(f.key, f.val)
}

type uint32Field struct {
	key string
	val uint32
}

func (f uint32Field) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f uint32Field) Log(ev *zerolog.Event) {
	ev.Uint32(f.key, f.val)
}

type uint64Field struct {
	key string
	val uint64
}

func (f uint64Field) String() string {
	return fmt.Sprintf("%s: %d", f.key, f.val)
}

func (f uint64Field) Log(ev *zerolog.Event) {
	ev.Uint64(f.key, f.val)
}

type float32Field struct {
	key string
	val float32
}

func (f float32Field) String() string {
	return fmt.Sprintf("%s: %f", f.key, f.val)
}

func (f float32Field) Log(ev *zerolog.Event) {
	ev.Float32(f.key, f.val)
}

type float64Field struct {
	key string
	val float64
}

func (f float64Field) String() string {
	return fmt.Sprintf("%s: %f", f.key, f.val)
}

func (f float64Field) Log(ev *zerolog.Event) {
	ev.Float64(f.key, f.val)
}

type bytesField struct {
	key string
	val []byte
}

func (f bytesField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f bytesField) Log(ev *zerolog.Event) {
	ev.Bytes(f.key, f.val)
}

type hexField struct {
	key string
	val []byte
}

func (f hexField) String() string {
	return fmt.Sprintf("%s: %x", f.key, f.val)
}

func (f hexField) Log(ev *zerolog.Event) {
	ev.Hex(f.key, f.val)
}

type strField struct {
	key string
	val string
}

func (f strField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f strField) Log(ev *zerolog.Event) {
	ev.Str(f.key, f.val)
}

type errField struct {
	key string
	err error
}

func (f errField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.err)
}

func (f errField) Log(ev *zerolog.Event) {
	ev.AnErr(f.key, f.err)
}

type durField struct {
	key string
	val time.Duration
}

func (f durField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f durField) Log(ev *zerolog.Event) {
	ev.Dur(f.key, f.val)
}

type timeField struct {
	key string
	val time.Time
}

func (f timeField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f timeField) Log(ev *zerolog.Event) {
	ev.Time(f.key, f.val)
}

type jsonField struct {
	key string
	val []byte
}

func (f jsonField) String() string {
	return fmt.Sprintf("%s: %s", f.key, string(f.val))
}

func (f jsonField) Log(ev *zerolog.Event) {
	ev.RawJSON(f.key, f.val)
}

type ifaceField struct {
	key string
	val interface{}
}

func (f ifaceField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f ifaceField) Log(ev *zerolog.Event) {
	ev.Interface(f.key, f.val)
}
