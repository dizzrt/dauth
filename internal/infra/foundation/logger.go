package foundation

import (
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/log/zlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(ac *conf.AppConfig) log.LogWriter {
	logAge, err := time.ParseDuration(ac.Log.MaxAge)
	if err != nil {
		panic(err)
	}

	logger, err := log.NewStdLoggerWriter(ac.Log.File,
		zlog.Symlink(ac.Log.Symlink),
		zlog.Level(zlog.ParseLevel(ac.Log.Level)),
		zlog.MaxAge(logAge),
		zlog.MaxBackups(uint(ac.Log.MaxBackups)),
		zlog.OutputType(zlog.ParseOutputType(ac.Log.OutputType)),
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
