package cmd

import (
	v1 "go-example/app/api/v1"
	"io"
	"os"

	// auto connect to sql
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start server at port 5000",
		Long:  `start server at port 5000`,
		Run: func(cmd *cobra.Command, agrs []string) {
			db, err := gorm.Open("sqlite3", "./test.db")
			if err != nil {
				panic("Failed to connect database: " + err.Error())
			}
			defer db.Close()
			fileLog, _ := os.Create("gin.log")
			gin.DefaultWriter = io.MultiWriter(fileLog)
			router := gin.Default()
			apiV1Router := router.Group("/api/v1")
			v1.RegisterRouterAPIV1(apiV1Router, db)
			router.Run(":5000")
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
