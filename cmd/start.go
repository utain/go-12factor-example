package cmd

import (
	v1 "go-example/app/services/api/v1"
	// auto connect to sql
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start server",
		Long:  `start server at port 5000`,
		Run: func(cmd *cobra.Command, agrs []string) {
			db, err := gorm.Open("sqlite3", "test.db")
			if err != nil {
				panic("failed to connect database")
			}
			defer db.Close()
			r := gin.Default()
			v1.Migrate(db)
			service := v1.GetService(db)
			r.GET("/user", service.GetAllUser)
			r.GET("/user/:id", service.GetUser)
			r.Run(":5000")
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
