package cmd

import (
	v1 "go-example/app/services/api/v1"
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
			db, err := gorm.Open("sqlite3", "test.db")
			if err != nil {
				panic("failed to connect database")
			}
			defer db.Close()
			f, _ := os.Create("gin.log")
			gin.DefaultWriter = io.MultiWriter(f)
			router := gin.Default()
			apiV1 := router.Group("/api/v1")
			{
				service := v1.NewUserService(db, true)
				apiV1.GET("/users", service.GetAllUser)
				apiV1.GET("/users/:id", service.GetUser)

				prodServ := v1.NewProductService(db, true)
				apiV1.GET("/products", prodServ.FindAll)
				apiV1.GET("/products/:id", prodServ.GetProduct)
				apiV1.DELETE("/products/:id", prodServ.DeleteProduct)
			}
			router.Run(":5000")
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
