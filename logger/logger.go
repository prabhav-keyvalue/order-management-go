package logger

import (
	"github.com/prabhav-keyvalue/order-management-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugaredLogger *zap.SugaredLogger

func InitLogger(env config.Environment) (*zap.SugaredLogger, error) {
	logger, err := GetLoggerByEnv(env)
	if err != nil {
		return nil, err
	}

	sugaredLogger = logger.Sugar()
	return sugaredLogger, nil
}

func GetLoggerByEnv(env config.Environment) (logger *zap.Logger, err error) {
	option := zap.AddCallerSkip(1)

	if env == config.Production {
		return zap.NewProduction(option)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return config.Build(option)

}

func Errorw(message string, args ...interface{}) {
	sugaredLogger.Errorw(message, args...)
}

func Errorf(message string, args ...interface{}) {
	sugaredLogger.Errorf(message, args...)
}

func Error(args ...interface{}) {
	sugaredLogger.Error(args...)
}

func Infow(message string, args ...interface{}) {
	sugaredLogger.Infow(message, args...)
}

func Infof(message string, args ...interface{}) {
	sugaredLogger.Infof(message, args...)
}

func Info(args ...interface{}) {
	sugaredLogger.Info(args...)
}

func Warnw(message string, args ...interface{}) {
	sugaredLogger.Warnw(message, args...)
}

func Warnf(message string, args ...interface{}) {
	sugaredLogger.Warnf(message, args...)
}

func Warn(args ...interface{}) {
	sugaredLogger.Warn(args...)
}

func Debugw(message string, args ...interface{}) {
	sugaredLogger.Debugw(message, args...)
}

func Debugf(message string, args ...interface{}) {
	sugaredLogger.Debugf(message, args...)
}

func Debug(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

func Fatalf(message string, args ...interface{}) {
	sugaredLogger.Fatalf(message, args...)
}

func Fatal(args ...interface{}) {
	sugaredLogger.Fatal(args...)
}
