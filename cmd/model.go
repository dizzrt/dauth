package cmd

import (
	"fmt"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	"github.com/spf13/cobra"
	"gorm.io/gen"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := execute(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	genCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func execute() error {
	conf := conf.GetAppConfig()
	db := foundation.NewDB(conf)

	g := gen.NewGenerator(gen.Config{
		OutPath:          "internal/infra/repo/core/gen/dao",
		ModelPkgPath:     "internal/infra/repo/core/gen/model",
		Mode:             gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:    true,
		FieldWithTypeTag: true,
	})

	g.UseDB(db)
	m := g.GenerateAllTable()
	// m := g.GenerateModel("identity_users")
	g.ApplyBasic(m...)

	g.Execute()

	return nil
}
