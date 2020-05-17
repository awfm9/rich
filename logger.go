package rich

import (
	"github.com/rs/zerolog"
)

type Logger struct {
	ev *zerolog.Event
}

func Log(entry func() *zerolog.Event) Logger {
	return Logger{ev: entry()}
}

func (l Logger) Err(err error) *zerolog.Event {
	r, ok := err.(*Error)
	if !ok {
		l.ev.Err(err)
	}
	for _, f := range r.fs {
		f.Log(l.ev)
	}
	return l.ev.Err(r.err)
}
