package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger
func InitLogger(environment string, logLevel zapcore.Level) {
	config := zap.Config{
               Level:       zap.NewAtomicLevelAt(logLevel),
               Development: environment == "development",
               Encoding:    "json",
               EncoderConfig: zapcore.EncoderConfig{
                       TimeKey:        "timestamp",
                       LevelKey:       "level",
                       NameKey:        "logger",
                       CallerKey:      "caller",
                       MessageKey:     "msg",
                       StacktraceKey:  "stacktrace",
                       LineEnding:     zapcore.DefaultLineEnding,
                       EncodeLevel:    zapcore.CapitalLevelEncoder,
                       EncodeTime:     zapcore.ISO8601TimeEncoder,
                       EncodeDuration: zapcore.StringDurationEncoder,
                       EncodeCaller:   zapcore.ShortCallerEncoder,
               },
               OutputPaths:      []string{"stdout"},
               ErrorOutputPaths: []string{"stderr"},
       }
       var err error
       Log, err = config.Build()
       if err != nil {
               panic(err)
       }
}
func SyncLogger() {
       if Log != nil {
               _ = Log.Sync()
       }
}
