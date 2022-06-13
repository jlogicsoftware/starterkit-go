package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	logger *zerolog.Logger
}

func (l *Logger) Write(p []byte) (n int, err error) {
	return l.logger.Write(p)
}

var output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if os.Getenv("ENV") == "development" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(output)
	}
}

func Info() *zerolog.Event {
	return log.Info()
}

func Warn() *zerolog.Event {
	return log.Warn()
}

func Error() *zerolog.Event {
	if os.Getenv("ENV") == "development" {
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("⛔️ %s", i)
		}
		log := log.Output(output)
		return log.Error()
	}
	return log.Error()
}

func Fatal() *zerolog.Event {
	return log.Fatal()
}
