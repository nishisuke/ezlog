package warn

import "github.com/rs/zerolog/log"

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Warnf(f string, i ...interface{}) {
	log.Warn().Msgf(f, i...)
}
