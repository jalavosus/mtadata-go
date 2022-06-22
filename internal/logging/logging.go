package logging

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger returns a configured zap.Logger instance.
func NewLogger() *zap.Logger {
	l, _ := zap.NewProduction(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return l
}

// WithLogger is an instantiated fx.WithLogger option
var WithLogger = fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: logger}
})
