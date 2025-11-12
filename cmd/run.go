package cmd

import (
	"context"
	"os"
	"time"

	"github.com/dizzrt/ellie/registry"
	"github.com/dizzrt/ellie/transport/grpc"
	"github.com/dizzrt/ellie/transport/http"

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
		wa, cleanup, err := wireApp()
		if err != nil {
			panic(err)
		}
		defer cleanup()

		if wa.TP != nil {
			defer func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				wa.TP.Shutdown(ctx)
			}()
		}

		if err := wa.App.Run(); err != nil {
			panic(err)
		}
	},
}

func newApp(logger log.LogWriter, registrar registry.Registrar, gs *grpc.Server, hs *http.Server) *ellie.App {
	return ellie.New(
		ellie.ID(id),
		ellie.Name(service),
		ellie.Version(version),
		ellie.Metadata(map[string]string{}),
		ellie.Logger(logger),
		ellie.Registrar(registrar),
		ellie.Server(gs, hs),
	)
}
