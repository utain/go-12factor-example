package cmd

import (
	"fmt"
	v1 "go-example/app/api/v1"
	"go-example/app/config"
	"io"
	"os"
	"strconv"

	// auto connect to sql
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var (
	configPath string
	startCmd   = &cobra.Command{
		Use:   "start",
		Short: "start server",
		Long:  `start server, default port is 5000`,
		Run: func(cmd *cobra.Command, agrs []string) {
			conf, _ := config.Read(configPath)
			fmt.Println("[Start server with config]:", conf)
			db, err := gorm.Open(conf.Database.Type, conf.Database.URL)
			if err != nil {
				panic(fmt.Errorf("Failed to connect database: %w", err))
			}
			defer db.Close()
			fileLog, _ := os.Create(conf.Logging.Path)
			gin.DefaultWriter = io.MultiWriter(fileLog)
			router := gin.Default()
			apiV1Router := router.Group("/api/v1")
			v1.RegisterRouterAPIV1(apiV1Router, db)
			router.Run(":" + strconv.Itoa(conf.Port))
		},
	}
)

func init() {
	startCmd.Flags().Int("port", 5000, "Port to run Application server on")
	startCmd.Flags().StringVar(&configPath, "config", "", "Path to config file for App server")
	config.Viper().BindPFlag("port", startCmd.Flags().Lookup("port"))
	config.Viper().BindPFlag("config", startCmd.Flags().Lookup("config"))
	rootCmd.AddCommand(startCmd)
}
