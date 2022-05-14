package ezlog

import (
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/nishisuke/ezlog/echologger"
	"github.com/rs/zerolog"
)

type (
	Config struct {
		io.Writer
		log.Lvl
		TimeFieldFormat string
	}
)

var DefaultConfig = Config{
	Writer:          os.Stdout,
	Lvl:             log.WARN,
	TimeFieldFormat: zerolog.TimeFormatUnix,
}

func Prepare(e *echo.Echo) {
	conf := DefaultConfig
	zerolog.TimeFieldFormat = conf.TimeFieldFormat
	e.Logger = echologger.New(conf.Writer, conf.Lvl)
}

func PrepareWithConfig(e *echo.Echo, conf *Config) {
	setDefaultIfZeroValue(conf)
	zerolog.TimeFieldFormat = conf.TimeFieldFormat
	e.Logger = echologger.New(conf.Writer, conf.Lvl)
}

func setDefaultIfZeroValue(conf *Config) {
	if conf.Writer == nil {
		conf.Writer = DefaultConfig.Writer
	}

	if conf.Lvl == 0 {
		conf.Lvl = DefaultConfig.Lvl
	}

	if conf.TimeFieldFormat == "" {
		conf.TimeFieldFormat = DefaultConfig.TimeFieldFormat
	}
}
