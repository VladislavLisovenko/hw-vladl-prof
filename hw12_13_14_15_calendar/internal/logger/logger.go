package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	log *zerolog.Logger
}

func New(level string) *Logger {
	switch strings.ToLower(level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logger := zerolog.New(os.Stdout)
	return &Logger{log: &logger}
}

func (l *Logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *Logger) Error(msg string) {
	l.log.Error().Msg(msg)
}

func (l *Logger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

func (l *Logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}
