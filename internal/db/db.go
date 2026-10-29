package db

import (
	"errors"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/yourusername/forge-network/internal/models"
	"github.com/yourusername/forge-network/internal/config"
)

// DB wraps a gorm DB instance.
type DB struct {
	*gorm.DB
}

// New connects to the configured database. Supports Postgres DSN or sqlite file.
func New(cfg *config.Config) (*DB, error) {
	dsn := cfg.DatabaseURL
	var dialector gorm.Dialector
	if strings.HasPrefix(dsn, "host=") || strings.Contains(dsn, "postgres") {
		dialector = postgres.Open(dsn)
	} else {
		// treat as sqlite path
		dialector = sqlite.Open(dsn)
	}

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{gormDB}, nil
}

// AutoMigrate runs model migrations for demo purposes.
func (db *DB) AutoMigrate() error {
	if db == nil || db.DB == nil {
		return errors.New("nil database")
	}
	return db.DB.AutoMigrate(&models.Wallet{}, &models.Transaction{})
}
