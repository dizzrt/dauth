package cmd

import (
	"os"

	"github.com/dizzrt/ellie/registry"
	"github.com/dizzrt/ellie/transport/grpc"
	"github.com/dizzrt/ellie/transport/http"
	"go.opentelemetry.io/otel/trace"

	"github.com/dizzrt/ellie"
	"github.com/dizzrt/ellie/log"
	"github.com/spf13/cobra"
)

// go build -ldflags "-X main.version=x.y.z"
var (
	service string = "dauth"
	version string = "dev"
	id, _          = os.Hostname()
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start this service",
	Run: func(cmd *cobra.Command, args []string) {
		app, cleanup, err := wireApp()
		if err != nil {
			panic(err)
		}

		defer cleanup()
		if err := app.Run(); err != nil {
			panic(err)
		}
	},
}

func newApp(logger log.LogWriter, tracer trace.TracerProvider, registrar registry.Registrar, gs *grpc.Server, hs *http.Server) (*ellie.App, func(), error) {
	app := ellie.New(
		ellie.ID(id),
		ellie.Name(service),
		ellie.Version(version),
		ellie.Metadata(map[string]string{}),
		ellie.Logger(logger),
		ellie.Tracer(tracer),
		ellie.Registrar(registrar),
		ellie.Server(gs, hs),
	)

	cleanup := func() {
		// do some cleanup
	}

	return app, cleanup, nil
}
