package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger initializes a custom logger with consistent JSON formatting
func InitLogger(environment string, logLevel zapcore.Level) {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel), // Dynamic log level
		Development: environment == "development",   // Enable development mode for debugging
		Encoding:    "json",                         // Force JSON encoding
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",                   // Key for timestamps
			LevelKey:       "level",                       // Key for log level
			NameKey:        "logger",                      // Key for logger name
			CallerKey:      "caller",                      // Key for caller info
			MessageKey:     "msg",                         // Key for the message
			StacktraceKey:  "stacktrace",                  // Key for stack traces
			LineEnding:     zapcore.DefaultLineEnding,     // Default line ending
			EncodeLevel:    zapcore.CapitalLevelEncoder,   // INFO, ERROR, etc.
			EncodeTime:     zapcore.ISO8601TimeEncoder,    // Human-readable timestamps
			EncodeDuration: zapcore.StringDurationEncoder, // e.g., "1.2s"
			EncodeCaller:   zapcore.ShortCallerEncoder,    // Short caller paths
		},
		OutputPaths:      []string{"stdout"}, // Log to console
		ErrorOutputPaths: []string{"stderr"}, // Log errors to stderr
	}

	// Build the logger
	var err error
	Log, err = config.Build()
	if err != nil {
		panic(err)
	}
}

// SyncLogger flushes any buffered log entries
func SyncLogger() {
	if Log != nil {
		_ = Log.Sync() // Ignore errors as it's not critical
	}
}
