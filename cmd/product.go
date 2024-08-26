package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var product = &cobra.Command{
	Use:     "migrate-product",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating product")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.Product)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model product up success")
	},
}

func init() {
	rootCmd.AddCommand(product)
}
