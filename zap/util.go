package rich

import (
	"errors"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func fs(err error) []zap.Field {
	var r *Error
	if errors.As(err, &r) {
		return append([]zap.Field{zap.Error(r.err)}, r.fields...)
	}
	var s *Sugared
	if errors.As(err, &s) {
		return append([]zap.Field{zap.Error(s.err)}, sweeten(s.args)...)
	}
	return []zap.Field{zap.Error(err)}
}

func as(err error) []interface{} {
	var r *Error
	if errors.As(err, &r) {
		return append([]interface{}{zap.Error(r.err)}, flatten(r.fields)...)
	}
	var s *Sugared
	if errors.As(err, &s) {
		return append([]interface{}{zap.Error(s.err)}, s.args...)
	}
	return []interface{}{zap.Error(s.err)}
}

func flatten(fields []zap.Field) []interface{} {
	args := make([]interface{}, 0, 2*len(fields))
	for _, field := range fields {
		args = append(args, field.Key, field.Interface)
	}
	return args
}

func sweeten(args []interface{}) []zap.Field {
	if len(args) == 0 {
		return nil
	}
	fields := make([]zap.Field, 0, len(args))
	var invalid invalidPairs
	for i := 0; i < len(args); {
		if f, ok := args[i].(zap.Field); ok {
			fields = append(fields, f)
			i++
			continue
		}
		if i == len(args)-1 {
			fields = append(fields, zap.Any("ignored", args[i]))
			break
		}
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, zap.Any(keyStr, val))
		}
		i += 2
	}
	if len(invalid) > 0 {
		fields = append(fields, zap.Array("invalid", invalid))
	}
	return fields
}

type invalidPair struct {
	position   int
	key, value interface{}
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	zap.Any("key", p.key).AddTo(enc)
	zap.Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}
