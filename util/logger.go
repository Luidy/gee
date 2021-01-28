package util

import (
	"strings"
	"time"

	"github.com/juju/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2020-01-01 00:00:00"))
}

func levelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// InitLog ...
func InitLog(logLv string) (*zap.SugaredLogger, error) {
	logConfig := zap.NewDevelopmentConfig()
	logConfig.Encoding = "json"

	if checkLogLevel(logLv, "debug") {
		logConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else if checkLogLevel(logLv, "info") {
		logConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	} else if checkLogLevel(logLv, "error") {
		logConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = timeEncoder
	encodeConfig.EncodeLevel = levelEncoder
	logConfig.EncoderConfig = encodeConfig

	logger, err := logConfig.Build()
	if err != nil {
		return nil, errors.Annotatef(err, "InitLog")
	}
	return logger.Sugar(), nil
}

func checkLogLevel(lv, target string) bool {
	return strings.Contains(strings.ToLower(lv), target)
}
