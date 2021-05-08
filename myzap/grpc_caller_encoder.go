package myzap

import (
	"go.uber.org/zap/zapcore"
)

func GrpcCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.TrimmedPath())
}
