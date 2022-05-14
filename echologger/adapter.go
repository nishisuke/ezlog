package echologger

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

func (l Logger) msg(lv zerolog.Level, i ...interface{}) {
	eventByLevel(l.zerolog, lv).Msg(fmt.Sprint(i...))
}

func (l Logger) msgf(lv zerolog.Level, f string, i ...interface{}) {
	eventByLevel(l.zerolog, lv).Msgf(f, i...)
}

func (l Logger) fields(lv zerolog.Level, j log.JSON) {
	mp := map[string]interface{}(j)
	eventByLevel(l.zerolog, lv).Fields(mp).Msg(NoSymbol)
}
