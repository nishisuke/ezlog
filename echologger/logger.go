package echologger

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/nishisuke/ezlog"
)

type Logger struct {
	*ezlog.Logger
	out    io.Writer
	level  log.Lvl
	prefix string
}

func New(w io.Writer, l log.Lvl) *Logger {
	return &Logger{
		Logger: ezlog.New(w, zerologLevel(l)),
		out:    w,
		level:  l,
	}
}
