package main

import (
	"fmt"
	v1 "go-example/internal/api/v1"
	"go-example/internal/config"
	"go-example/internal/models"
	"go-example/log"
	"strings"

	// auto connect to sql
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
)

var (
	configPath string
	startCmd   = &cobra.Command{
		Use:   "start",
		Short: "start server",
		Long:  `start server, default port is 5000`,
		Run:   startServer,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is $HOME/.serverd/default.yaml)")
	startCmd.PersistentFlags().Int("port", 5000, "Port to run Application server on")
	config.Viper().BindPFlag("port", startCmd.PersistentFlags().Lookup("port"))
}

func initConfig() {
	config.Viper().SafeWriteConfig()
	config.Viper().WriteConfigAs("$HOME/.serverd/.config")
	if len(configPath) != 0 {
		config.Viper().SetConfigFile(configPath)
	} else {
		config.Viper().AddConfigPath("$HOME/.serverd/")
		config.Viper().AddConfigPath("./config")
		config.Viper().SetConfigName("default")
	}
	config.Viper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.Viper().AutomaticEnv()
	if err := config.Viper().ReadInConfig(); err != nil {
		log.Fatalf("Using config file [%s]: %v", config.Viper().ConfigFileUsed(), err)
	}
	log.Info("Config paths:", config.Viper().ConfigFileUsed())
	log.Info("DBConnection:", len(config.Viper().GetString("database.url")))
}

func startServer(cmd *cobra.Command, agrs []string) {
	log.Info("Start http-server")
	db, err := gorm.Open("postgres", config.AllConf().Database.URL)
	if err != nil {
		log.Fatalf("Failed to connect database: %w", err)
	}
	defer db.Close()
	go models.AutoMigrate(db)
	router := gin.Default()
	pprof.Register(router, "monitor/pprof")
	apiV1Router := router.Group("/api/v1")
	v1.RegisterRouterAPIV1(apiV1Router, db)
	router.Run(fmt.Sprintf(":%d", config.AllConf().Port))
}
