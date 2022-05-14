package echologger

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/nishisuke/ezlog/internal/warn"
)

const (
	prefixKey = "prefix"
)

func (l Logger) Output() io.Writer {
	return l.out
}

func (l *Logger) SetOutput(w io.Writer) {
	l.out = w
	l.Logger.SetZerolog(l.Logger.Zerolog.Output(w))
}

func (l Logger) Level() log.Lvl {
	return l.level
}

func (l *Logger) SetLevel(lvl log.Lvl) {
	l.level = lvl
	l.Logger.SetZerolog(l.Logger.Zerolog.Level(zerologLevel(lvl)))
}

func (l Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) SetPrefix(p string) {
	l.prefix = p
	l.Logger.SetZerolog(l.Logger.Zerolog.With().Str(prefixKey, p).Logger())
}

func (l Logger) SetHeader(h string) {
	warn.Warn("Not implemented SetHeader")
}
