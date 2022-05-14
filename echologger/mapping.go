package echologger

import (
	"github.com/labstack/gommon/log"
	"github.com/nishisuke/ezlog/internal/warn"
	"github.com/rs/zerolog"
)

func zerologLevel(l log.Lvl) zerolog.Level {
	switch l {
	case log.DEBUG:
		return zerolog.DebugLevel
	case log.INFO:
		return zerolog.InfoLevel
	case log.WARN:
		return zerolog.WarnLevel
	case log.ERROR:
		return zerolog.ErrorLevel
	case log.OFF: // OFF logs only fatal and panic
		return zerolog.FatalLevel
	}

	warn.Warnf("Unhandled gommon level: %s", l)
	return zerolog.TraceLevel
}

func gommonLevel(l zerolog.Level) log.Lvl {
	switch l {
	case zerolog.DebugLevel, zerolog.TraceLevel:
		return log.DEBUG
	case zerolog.InfoLevel:
		return log.INFO
	case zerolog.WarnLevel:
		return log.WARN
	case zerolog.ErrorLevel:
		return log.ERROR
	case zerolog.FatalLevel:
		return log.OFF
	case zerolog.PanicLevel, zerolog.NoLevel, zerolog.Disabled:
		warn.Warn("No compatible level in gommon")
		return log.OFF
	}

	warn.Warnf("Unhandled zerolog level: %s", l)
	return log.OFF
}
