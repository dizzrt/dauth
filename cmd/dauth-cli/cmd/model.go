package cmd

import (
	"fmt"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	"github.com/spf13/cobra"
	"gorm.io/gen"
)

var (
	allTables bool
	tableName string
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
		execute()
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

	modelCmd.Flags().BoolVar(&allTables, "all", false, "Generate models for all tables")
	modelCmd.Flags().StringVarP(&tableName, "table", "t", "", "Table name to generate model")
}

func execute() error {
	if !allTables && tableName == "" {
		fmt.Println("Please specify table name or use --all to generate models for all tables")
		return nil
	}

	conf := conf.GetAppConfig()
	db := foundation.NewDB(conf)

	g := gen.NewGenerator(gen.Config{
		OutPath:          "internal/infra/persistence/core/gen/dao",
		ModelPkgPath:     "internal/infra/persistence/core/gen/model",
		Mode:             gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:    true,
		FieldWithTypeTag: true,
	})

	g.UseDB(db)

	if tableName != "" {
		m := g.GenerateModel(tableName)
		g.ApplyBasic(m)
	} else {
		m := g.GenerateAllTable()
		g.ApplyBasic(m...)
	}

	g.Execute()
	return nil
}
