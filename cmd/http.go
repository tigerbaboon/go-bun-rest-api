package cmd

import (
	"app/modules"
	"app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var http = &cobra.Command{
	Use:     "http",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		godotenv.Load()

		mod := modules.Get()
		routes.Api(r.Group("/api"), mod)

		r.Run(":8080")
	},
}

func init() {
	rootCmd.AddCommand(http)
}
