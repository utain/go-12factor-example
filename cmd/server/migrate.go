package main

import (
	"go-example/internal/config"
	"go-example/internal/models"
	"go-example/log"

	// auto connect to sql

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
)

var (
	migrateCMD = &cobra.Command{
		Use:   "migrate",
		Short: "automated migrate data",
		Long:  `automated migrate database schema and initial data`,
		Run:   migrateCMDRunner,
	}
	initialData bool
)

func init() {
	rootCmd.AddCommand(migrateCMD)
	migrateCMD.PersistentFlags().BoolVarP(&initialData, "data", "d", false, "initial data also (default: false)")
}

func migrateCMDRunner(cmd *cobra.Command, agrs []string) {
	log.Info("Start migrate")
	db, err := gorm.Open(config.Viper().GetString("database.type"), config.Viper().GetString("database.url"))
	if err != nil {
		log.Fatalf("Failed to connect database[%v]: %w", config.Get("database.type"), err)
	}
	defer db.Close()
	models.AutoMigrate(db)
	if initialData {
		models.InitData(db)
	}
}
