package cmd

import (
	"context"
	"os"
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/log/zlog"
	"github.com/dizzrt/ellie/middleware/tracing"
	"github.com/dizzrt/ellie/transport/grpc"
	"github.com/dizzrt/ellie/transport/http"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dizzrt/ellie"
	"github.com/dizzrt/ellie/config"
	"github.com/dizzrt/ellie/log"
	"github.com/spf13/cobra"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     string
	Version  string
	flagConf string
	id, _    = os.Hostname()
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start this service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// load configs
		conf := config.NewStdViperConfig()
		if err := conf.Load(); err != nil {
			panic(err)
		}

		bootstrap, err := buildBootstrap(conf)
		if err != nil {
			panic(err)
		}

		// init logger
		logger := initLogger(bootstrap)

		// init trace provider
		tp := initTrace(ctx, bootstrap)
		defer func() {
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			tp.Shutdown(ctx)
		}()

		// init app
		app, cleanup, err := wireApp(bootstrap, logger)
		if err != nil {
			panic(err)
		}
		defer cleanup()

		if err := app.Run(); err != nil {
			panic(err)
		}
	},
}

func initTrace(ctx context.Context, _ *conf.Bootstrap) *trace.TracerProvider {
	// TODO read from conf
	tp, err := tracing.Initialize(ctx,
		tracing.ServiceName("dauth"),
		tracing.ServiceVersion("dev"),
		tracing.Endpoint("192.168.124.10:4317"),
		tracing.EndpointType(tracing.EndpointType_GRPC),
		tracing.Insecure(true),
		tracing.Metadata(map[string]string{
			"ip":  "127.0.0.1",
			"env": "dev",
		}),
	)

	if err != nil {
		panic(err)
	}

	return tp
}

func initLogger(bootstrap *conf.Bootstrap) log.LogWriter {
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

func buildBootstrap(c config.Config) (*conf.Bootstrap, error) {
	var bootstrap conf.Bootstrap
	if err := c.UnmarshalKey("server", &bootstrap.Server); err != nil {
		return nil, err
	}

	if err := c.UnmarshalKey("log", &bootstrap.Log); err != nil {
		return nil, err
	}

	return &bootstrap, nil
}

func newApp(logger log.LogWriter, gs *grpc.Server, hs *http.Server) *ellie.App {
	return ellie.New(
		ellie.ID(id),
		ellie.Name(Name),
		ellie.Version(Version),
		ellie.Metadata(map[string]string{}),
		ellie.Logger(logger),
		ellie.Server(gs, hs),
	)
}
