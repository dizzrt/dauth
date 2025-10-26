package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/log/zlog"
	"github.com/dizzrt/ellie/transport/grpc"
	"github.com/dizzrt/ellie/transport/http"
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
		conf := config.NewStdViperConfig()
		if err := conf.Load(); err != nil {
			panic(err)
		}

		bootstrap, err := buildBootstrap(conf)
		if err != nil {
			panic(err)
		}

		fmt.Println(bootstrap.Log.File)

		logger, err := log.NewStdLoggerWriter(bootstrap.Log.File,
			zlog.Symlink(bootstrap.Log.Symlink),
			zlog.Level(zlog.ParseLevel(bootstrap.Log.Level)),
			zlog.MaxAge(time.Duration(bootstrap.Log.MaxAge)*time.Second),
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
