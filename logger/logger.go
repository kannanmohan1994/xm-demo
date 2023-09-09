package logger

import (
	"context"
	"xm/config"
	"xm/consts"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugaredLogger *zap.SugaredLogger

func InitLogger(cfg *config.Config) {
	logger, err := getLoggerbyEnv(cfg.Environment)
	if err != nil {
		panic(err)
	}

	sugaredLogger = logger.Sugar()
}

func getLoggerbyEnv(env string) (logger *zap.Logger, err error) {
	option := zap.AddCallerSkip(1)

	if env == string(consts.PRODUCTION) {
		return zap.NewProduction(option)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return config.Build(option)
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	newLogger := sugaredLogger

	if ctx != nil {
		if ctxRqId, ok := ctx.Value(consts.CorrelationID).(string); ok {
			newLogger = newLogger.With(zap.String(string(consts.CorrelationID), ctxRqId))
		}
	}

	return newLogger
}

func Errorw(ctx context.Context, message string, args ...interface{}) {
	args = appendRequestIDIntoArgs(ctx, args)
	sugaredLogger.Errorw(message, args...)
}

func Infow(ctx context.Context, message string, args ...interface{}) {
	args = appendRequestIDIntoArgs(ctx, args)
	sugaredLogger.Infow(message, args...)
}

func Criticalw(ctx context.Context, criticalError string, message string, args ...interface{}) {
	args = append(args, "CRITICAL_ERROR")
	args = append(args, criticalError)
	args = appendRequestIDIntoArgs(ctx, args)
	sugaredLogger.Errorw(message, args...)
}

func appendRequestIDIntoArgs(ctx context.Context, args []interface{}) []interface{} {
	ridValue, ok := ctx.Value(consts.CorrelationID).(string)
	if !ok {
		return args
	}
	args = append(args, consts.CorrelationID)
	args = append(args, ridValue)
	return args
}

func Errorf(message string, args ...interface{}) {
	sugaredLogger.Errorf(message, args...)
}

func Error(args ...interface{}) {
	sugaredLogger.Error(args...)
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
