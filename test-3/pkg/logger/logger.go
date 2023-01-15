package logger

import (
	"io"
	"os"
	"path"
	"strings"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// type Interface interface {
// 	Debug(message interface{}, args ...interface{})
// 	Info(message string, args ...interface{})
// 	Warn(message string, args ...interface{})
// 	Error(message interface{}, args ...interface{})
// 	Fatal(message interface{}, args ...interface{})
// }

type Logger struct {
	*zerolog.Logger
}

// var _ Interface = (*Logger)(nil)

func New(config Config) *Logger {
	var (
		logWriters []io.Writer
		logLevel   zerolog.Level
	)

	// If config.ConsoleLoggingEnabled is true then set io writers for console log
	if config.ConsoleLoggingEnabled {
		logWriters = append(logWriters, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	// If config.FileLoggingEnabled is true then set io writers for file log
	if config.FileLoggingEnabled {
		logWriters = append(logWriters, newRollingFile(config))
	}

	// Set zerolog's log level
	switch strings.ToLower(config.Level) {
	case "panic":
		logLevel = zerolog.PanicLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "debug":
		logLevel = zerolog.DebugLevel
	case "trace":
		logLevel = zerolog.TraceLevel
	default:
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	// Set io writers for console & file logging
	mw := io.MultiWriter(logWriters...)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJson).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("logging configured")

	return &Logger{
		&logger,
	}
}

func newRollingFile(config Config) io.Writer {
	// file logger use lumberjack library for rotate capabilities
	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
	}
}
