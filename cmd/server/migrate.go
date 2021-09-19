package main

import (
	"go-example/internal/config"
	"go-example/internal/entities"
	"go-example/internal/log"

	// auto connect to sql

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
)

var (
	migrateCMD = &cobra.Command{
		Use:   "migrate",
		Short: "migrate db schema and seed data",
		Run:   migrateCMDRunner,
	}
	seedData bool
)

func init() {
	rootCmd.AddCommand(migrateCMD)
	migrateCMD.PersistentFlags().BoolVarP(&seedData, "seed", "s", false, "seed data (default: false)")
}

func migrateCMDRunner(cmd *cobra.Command, agrs []string) {
	log.Info("Start migrate")
	db, err := gorm.Open("postgres", config.Parse().Database.URL)
	if err != nil {
		log.Fatalf("Failed to connect database[%v]: %w", config.Parse().Database.URL, err)
	}
	defer db.Close()
	entities.AutoMigrate(db)
	if seedData {
		entities.SeedData(db)
	}
}
