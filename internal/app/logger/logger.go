package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type BaseLogger struct {
	lg zerolog.Logger
}

func NewLogger() *BaseLogger {
	return &BaseLogger{
		lg: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func (bl *BaseLogger) Log(msg string, requestId string) {
	if requestId == "" {
		bl.lg.Info().Msg(msg)
	} else {
		bl.lg.Info().Str("request-id", requestId).Msg(msg)
	}
}

func (bl *BaseLogger) LogError(msg string, err error, requestId string) {
	if requestId == "" {
		bl.lg.Err(err).Msg(msg)
	} else {
		bl.lg.Err(err).Str("request-id", requestId).Msg(msg)
	}
}

func (bl *BaseLogger) LogFatal(msg string) {
	bl.lg.Fatal().Msg(msg)
}
