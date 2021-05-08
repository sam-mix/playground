package myzap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/natefinch/lumberjack"
)

type ZapLogger struct {
	Logger *zap.Logger
}

func NewLogger() *ZapLogger {
	return NewLoggerWithEncoder(zapcore.ISO8601TimeEncoder, GrpcCallerEncoder)
}

func NewLoggerWithEncoder(timeEncoder zapcore.TimeEncoder, callerEncoder zapcore.CallerEncoder) *ZapLogger {
	level := zap.ErrorLevel
	var syncers []zapcore.WriteSyncer
	logger := &lumberjack.Logger{
		Filename:   "/tmp/playground/log.log",
		MaxSize:    100,
		MaxBackups: 500,
		MaxAge:     10000000,
		Compress:   false,
	}
	syncers = append(syncers, zapcore.AddSync(logger))
	syncers = append(syncers, zapcore.AddSync(os.Stdout))
	ws := zapcore.NewMultiWriteSyncer(syncers...)

	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeTime = timeEncoder
	encoderConf.EncodeCaller = callerEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConf)

	return &ZapLogger{
		Logger: zap.New(
			zapcore.NewCore(encoder, ws, zap.NewAtomicLevelAt(level)),
			zap.AddCaller(),
			zap.AddCallerSkip(2),
		),
	}
}
