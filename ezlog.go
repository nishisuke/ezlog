package ezlog

import (
	"io"

	"github.com/rs/zerolog"
)

const (
	NoSymbol = "-"
	LevelKey = "level"
)

type Logger struct {
	Zerolog zerolog.Logger
}

func New(w io.Writer, l zerolog.Level) *Logger {
	return &Logger{
		Zerolog: newZerolog(w, l),
	}
}

func newZerolog(w io.Writer, l zerolog.Level) zerolog.Logger {
	return zerolog.New(w).
		Level(l).
		With().
		Timestamp().
		Logger()
}

func (l *Logger) SetZerolog(logger zerolog.Logger) {
	l.Zerolog = logger
}
func Prepare() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func EventByLevel(logger zerolog.Logger, level zerolog.Level) *zerolog.Event {
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
