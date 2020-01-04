package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	log logger
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

type VideosLogger interface {
	Printf(format string, v ...interface{})
	Print(v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		OutputPaths: []string{getOutput()},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	log.log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

func GetLogger() VideosLogger {
	return log
}

func getOutput() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(envLogOutput)))
	if output == "" {
		return "stdout"
	}
	return output
}

func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	log.log.Error(msg, tags...)
	log.log.Sync()
}
