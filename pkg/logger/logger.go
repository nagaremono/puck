package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger struct {
	*zerolog.Logger
}

func New(level zerolog.Level) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(level)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout})

	return &Logger{
		&logger,
	}
}
