package migration

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/yourusername/dbmigrate/database"
)

// Config represents the migration configuration
type Config struct {
	Version string `json:"version"`
	Script  string `json:"script"`
}

// Migrate migrates the database
func Migrate(db *sql.DB, migrations []Config) error {
	for _, migration := range migrations {
		// Load migration script from file
		scriptPath := filepath.Join("migrations", migration.Script)
		script, err := ioutil.ReadFile(scriptPath)
		if err != nil {
			return err
		}

		// Execute migration script
		if _, err := db.Exec(string(script)); err != nil {
			return err
		}

		log.Printf("Applied migration %s", migration.Version)
	}

	return nil
}