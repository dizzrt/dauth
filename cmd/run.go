package cmd

import (
	"context"
	"os"
	"time"

	"github.com/dizzrt/ellie/transport/grpc"
	"github.com/dizzrt/ellie/transport/http"

	"github.com/dizzrt/ellie"
	"github.com/dizzrt/ellie/log"
	"github.com/spf13/cobra"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name    string = "dauth"
	Version string = "dev"
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

func newApp(logger log.LogWriter, gs *grpc.Server, hs *http.Server) *ellie.App {
	// ctx := context.Background()
	// cli, err := api.NewClient(&api.Config{
	// 	Address: "dev-lan.dauth.com:8500",
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// dis := consul.New(cli)

	return ellie.New(
		ellie.ID(id),
		ellie.Name(Name),
		ellie.Version(Version),
		ellie.Metadata(map[string]string{}),
		ellie.Logger(logger),
		ellie.Server(gs, hs),
		// ellie.Registrar(dis),
	)
}
