// Code generated. DO NOT EDIT.

package echologger

import (
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

func (l Logger) Debug(i ...interface{}) {
	l.msg(zerolog.DebugLevel, i...)
}

func (l Logger) Debugf(f string, i ...interface{}) {
	l.msgf(zerolog.DebugLevel, f, i...)
}

func (l Logger) Debugj(j log.JSON) {
	l.fields(zerolog.DebugLevel, j)
}

func (l Logger) Info(i ...interface{}) {
	l.msg(zerolog.InfoLevel, i...)
}

func (l Logger) Infof(f string, i ...interface{}) {
	l.msgf(zerolog.InfoLevel, f, i...)
}

func (l Logger) Infoj(j log.JSON) {
	l.fields(zerolog.InfoLevel, j)
}

func (l Logger) Warn(i ...interface{}) {
	l.msg(zerolog.WarnLevel, i...)
}

func (l Logger) Warnf(f string, i ...interface{}) {
	l.msgf(zerolog.WarnLevel, f, i...)
}

func (l Logger) Warnj(j log.JSON) {
	l.fields(zerolog.WarnLevel, j)
}

func (l Logger) Error(i ...interface{}) {
	l.msg(zerolog.ErrorLevel, i...)
}

func (l Logger) Errorf(f string, i ...interface{}) {
	l.msgf(zerolog.ErrorLevel, f, i...)
}

func (l Logger) Errorj(j log.JSON) {
	l.fields(zerolog.ErrorLevel, j)
}

func (l Logger) Fatal(i ...interface{}) {
	l.msg(zerolog.FatalLevel, i...)
}

func (l Logger) Fatalf(f string, i ...interface{}) {
	l.msgf(zerolog.FatalLevel, f, i...)
}

func (l Logger) Fatalj(j log.JSON) {
	l.fields(zerolog.FatalLevel, j)
}

func (l Logger) Panic(i ...interface{}) {
	l.msg(zerolog.PanicLevel, i...)
}

func (l Logger) Panicf(f string, i ...interface{}) {
	l.msgf(zerolog.PanicLevel, f, i...)
}

func (l Logger) Panicj(j log.JSON) {
	l.fields(zerolog.PanicLevel, j)
}

func (l Logger) Print(i ...interface{}) {
	l.msg(zerolog.NoLevel, i...)
}

func (l Logger) Printf(f string, i ...interface{}) {
	l.msgf(zerolog.NoLevel, f, i...)
}

func (l Logger) Printj(j log.JSON) {
	l.fields(zerolog.NoLevel, j)
}
