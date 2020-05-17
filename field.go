package rich

import (
	"fmt"
	"net"
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
	return fmt.Sprintf("%s: %s", f.key, f.val.Format(zerolog.TimeFieldFormat))
}

func (f timeField) Log(ev *zerolog.Event) {
	ev.Time(f.key, f.val)
}

type boolsField struct {
	key string
	val []bool
}

func (f boolsField) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f boolsField) Log(ev *zerolog.Event) {
	ev.Bools(f.key, f.val)
}

type intsField struct {
	key string
	val []int
}

func (f intsField) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f intsField) Log(ev *zerolog.Event) {
	ev.Ints(f.key, f.val)
}

type ints8Field struct {
	key string
	val []int8
}

func (f ints8Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f ints8Field) Log(ev *zerolog.Event) {
	ev.Ints8(f.key, f.val)
}

type ints16Field struct {
	key string
	val []int16
}

func (f ints16Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f ints16Field) Log(ev *zerolog.Event) {
	ev.Ints16(f.key, f.val)
}

type ints32Field struct {
	key string
	val []int32
}

func (f ints32Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f ints32Field) Log(ev *zerolog.Event) {
	ev.Ints32(f.key, f.val)
}

type ints64Field struct {
	key string
	val []int64
}

func (f ints64Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f ints64Field) Log(ev *zerolog.Event) {
	ev.Ints64(f.key, f.val)
}

type uintsField struct {
	key string
	val []uint
}

func (f uintsField) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f uintsField) Log(ev *zerolog.Event) {
	ev.Uints(f.key, f.val)
}

type uints8Field struct {
	key string
	val []uint8
}

func (f uints8Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f uints8Field) Log(ev *zerolog.Event) {
	ev.Uints8(f.key, f.val)
}

type uints16Field struct {
	key string
	val []uint16
}

func (f uints16Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f uints16Field) Log(ev *zerolog.Event) {
	ev.Uints16(f.key, f.val)
}

type uints32Field struct {
	key string
	val []uint32
}

func (f uints32Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f uints32Field) Log(ev *zerolog.Event) {
	ev.Uints32(f.key, f.val)
}

type uints64Field struct {
	key string
	val []uint64
}

func (f uints64Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f uints64Field) Log(ev *zerolog.Event) {
	ev.Uints64(f.key, f.val)
}

type floats32Field struct {
	key string
	val []float32
}

func (f floats32Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f floats32Field) Log(ev *zerolog.Event) {
	ev.Floats32(f.key, f.val)
}

type floats64Field struct {
	key string
	val []float64
}

func (f floats64Field) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f floats64Field) Log(ev *zerolog.Event) {
	ev.Floats64(f.key, f.val)
}

type strsField struct {
	key string
	val []string
}

func (f strsField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f strsField) Log(ev *zerolog.Event) {
	ev.Strs(f.key, f.val)
}

type errsField struct {
	key string
	err []error
}

func (f errsField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.err)
}

func (f errsField) Log(ev *zerolog.Event) {
	ev.Errs(f.key, f.err)
}

type dursField struct {
	key string
	val []time.Duration
}

func (f dursField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f dursField) Log(ev *zerolog.Event) {
	ev.Durs(f.key, f.val)
}

type timesField struct {
	key string
	val []time.Time
}

func (f timesField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f timesField) Log(ev *zerolog.Event) {
	ev.Times(f.key, f.val)
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

type ipField struct {
	key string
	val net.IP
}

func (f ipField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f ipField) Log(ev *zerolog.Event) {
	ev.IPAddr(f.key, f.val)
}

type prefixField struct {
	key string
	val net.IPNet
}

func (f prefixField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f prefixField) Log(ev *zerolog.Event) {
	ev.IPPrefix(f.key, f.val)
}

type macField struct {
	key string
	val net.HardwareAddr
}

func (f macField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f macField) Log(ev *zerolog.Event) {
	ev.MACAddr(f.key, f.val)
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

type tsField struct {
	val time.Time
}

func (f tsField) String() string {
	return fmt.Sprintf("%s: %s", zerolog.TimestampFieldName, f.val.Format(zerolog.TimeFieldFormat))
}

func (f tsField) Log(ev *zerolog.Event) {
	ev.Time(zerolog.TimestampFieldName, f.val)
}

type diffField struct {
	key  string
	val1 time.Time
	val2 time.Time
}

func (f diffField) String() string {
	return fmt.Sprintf("%s: %s -> %s", f.key, f.val1.Format(zerolog.TimeFieldFormat), f.val2.Format(zerolog.TimeFieldFormat))
}

func (f diffField) Log(ev *zerolog.Event) {
	ev.TimeDiff(f.key, f.val1, f.val2)
}

type fieldsField struct {
	val map[string]interface{}
}

func (f fieldsField) String() string {
	ss := make([]string, 0, len(f.val))
	for k, v := range f.val {
		ss = append(ss, fmt.Sprintf("%s: %s", k, v))
	}
	return strings.Join(ss, ", ")
}

func (f fieldsField) Log(ev *zerolog.Event) {
	ev.Fields(f.val)
}

type arrayField struct {
	key string
	val zerolog.LogArrayMarshaler
}

func (f arrayField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f arrayField) Log(ev *zerolog.Event) {
	ev.Array(f.key, f.val)
}

type dictField struct {
	key string
	val *zerolog.Event
}

func (f dictField) String() string {
	return fmt.Sprintf("%s: %v", f.key, f.val)
}

func (f dictField) Log(ev *zerolog.Event) {
	ev.Dict(f.key, f.val)
}

type objectField struct {
	key string
	val zerolog.LogObjectMarshaler
}

func (f objectField) String() string {
	return fmt.Sprintf("%s: %s", f.key, f.val)
}

func (f objectField) Log(ev *zerolog.Event) {
	ev.Object(f.key, f.val)
}

type embedField struct {
	val zerolog.LogObjectMarshaler
}

func (f embedField) String() string {
	return fmt.Sprintf("%s", f.val)
}

func (f embedField) Log(ev *zerolog.Event) {
	ev.EmbedObject(f.val)
}

type callerField struct {
	file string
	line int
}

func (f callerField) String() string {
	return fmt.Sprintf("%s: %s", zerolog.CallerFieldName, zerolog.CallerMarshalFunc(f.file, f.line))
}

func (f callerField) Log(ev *zerolog.Event) {
	ev.Str(zerolog.CallerFieldName, zerolog.CallerMarshalFunc(f.file, f.line))
}
