package echologger

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/nishisuke/ezlog"
	"github.com/rs/zerolog"
)

func (l Logger) msg(lv zerolog.Level, i ...interface{}) {
	ezlog.EventByLevel(l.Logger.Zerolog, lv).Msg(fmt.Sprint(i...))
}

func (l Logger) msgf(lv zerolog.Level, f string, i ...interface{}) {
	ezlog.EventByLevel(l.Logger.Zerolog, lv).Msgf(f, i...)
}

func (l Logger) fields(lv zerolog.Level, j log.JSON) {
	mp := map[string]interface{}(j)
	ezlog.EventByLevel(l.Logger.Zerolog, lv).Fields(mp).Msg(ezlog.NoSymbol)
}
