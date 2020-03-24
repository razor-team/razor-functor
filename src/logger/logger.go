package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

var debug bool = false

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

//Log write log msgs to std out
func Log(s string, v ...interface{}) {
	if debug {
		log.Debug().Msgf(s, v...)
	}
}

//Debug enable debug logging output
func Debug(v interface{}) {
	debug = cast.ToBool(v)
}
