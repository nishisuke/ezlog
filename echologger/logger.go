package echologger

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

const (
	NoSymbol = "-"
	LevelKey = "level"
)

type Logger struct {
	zerolog zerolog.Logger
	out     io.Writer
	level   log.Lvl
	prefix  string
}

func New(w io.Writer, l log.Lvl) *Logger {
	z := zerolog.New(w).
		Level(zerologLevel(l)).
		With().
		Timestamp().
		Logger()

	return &Logger{
		zerolog: z,
		out:     w,
		level:   l,
	}
}

func eventByLevel(logger zerolog.Logger, level zerolog.Level) *zerolog.Event {
	switch level {
	case zerolog.FatalLevel:
		return logger.Fatal() // WithLevel does not call os.Exit.
	case zerolog.PanicLevel:
		return logger.Panic() // WithLevel does not call panic.
	case zerolog.NoLevel:
		return logger.Log().Str(LevelKey, NoSymbol)
	default:
		return logger.WithLevel(level)
	}
}
