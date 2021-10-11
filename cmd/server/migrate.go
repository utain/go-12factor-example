package main

import (
	"go-example/internal/config"
	"go-example/internal/entities"
	"go-example/internal/log"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	migrateCMD = &cobra.Command{
		Use:   "migrate",
		Short: "migrate db schema and seed data",
		Run:   migrateCMDRunner,
	}
)

func init() {
	rootCmd.AddCommand(migrateCMD)
}

func migrateCMDRunner(cmd *cobra.Command, agrs []string) {
	log.Info("Start migrate")
	db, err := gorm.Open(postgres.Open(config.Default.Database.URL))
	if err != nil {
		log.Fatalf("Failed to connect database[%v]: %w", config.Parse().Database.URL, err)
	}
	defer func() {
		if dbSql, err := db.DB(); err != nil {
			dbSql.Close()
		}
	}()
	entities.AutoMigrate(db)
}
