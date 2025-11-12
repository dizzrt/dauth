package common

import (
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/log/zlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(bootstrap *conf.Bootstrap) log.LogWriter {
	logAge, err := time.ParseDuration(bootstrap.Log.MaxAge)
	if err != nil {
		panic(err)
	}

	logger, err := log.NewStdLoggerWriter(bootstrap.Log.File,
		zlog.Symlink(bootstrap.Log.Symlink),
		zlog.Level(zlog.ParseLevel(bootstrap.Log.Level)),
		zlog.MaxAge(logAge),
		zlog.MaxBackups(uint(bootstrap.Log.MaxBackups)),
		zlog.OutputType(zlog.ParseOutputType(bootstrap.Log.OutputType)),
		zlog.ZapOpts(
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
			zap.AddCallerSkip(2),
		),
	)

	if err != nil {
		panic(err)
	}

	return logger
}
