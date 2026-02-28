package database

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

// Config represents the database configuration
type Config struct {
	Dialect string `json:"dialect"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
}

// Connect connects to the database
func Connect(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Name)

	db, err := sql.Open(cfg.Dialect, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}