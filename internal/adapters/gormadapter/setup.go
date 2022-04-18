package gormadapter

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormConfig struct {
	// Connection url
	Url string
	// SetMaxOpenConns sets the maximum number of open connections to the database (min: 10).
	MaxPoolOpen int
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool (min: 5).
	MaxPoolIdle int
}

func Connect(conf GormConfig) (db *gorm.DB, err error) {
	sqlDB, err := sql.Open("pgx", conf.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect datasource: %w", err)
	}

	if conf.MaxPoolOpen < 10 {
		conf.MaxPoolOpen = 10
	}
	if conf.MaxPoolIdle < 5 {
		conf.MaxPoolIdle = 5
	}
	sqlDB.SetMaxOpenConns(conf.MaxPoolOpen)
	sqlDB.SetMaxIdleConns(conf.MaxPoolIdle)

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to initial driver: %w", err)
	}

	return db, err
}

func Close(db *gorm.DB) (err error) {
	if db == nil {
		return nil
	}
	sql, err := db.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}
