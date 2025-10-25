package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/iqbalnzls/sistem-manajemen-armada/pkg/config"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config *config.DatabaseConfig) *Database {
	fmt.Println("Try NewDatabase ...")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=disable search_path=%s",
		config.Username,
		config.Password,
		config.Name,
		config.Host,
		config.Port,
		config.Schema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	psql, err := db.DB()
	if err != nil {
		panic(err)
	}

	psql.SetMaxIdleConns(config.MaxIdleConnections)
	psql.SetMaxOpenConns(config.MaxOpenConnections)

	if config.DebugMode {
		db = db.Debug()
	}

	return &Database{
		db,
	}
}
