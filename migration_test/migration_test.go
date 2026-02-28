# Minor edit
# Minor edit
package migration

import (
	"database/sql"
	"testing"

	"github.com/yourusername/dbmigrate/database"
	_ "github.com/lib/pq"
)

func TestMigrate(t *testing.T) {
	// Create a test database
	db, err := database.Connect(database.Config{
		Dialect: "postgres",
		Username: "postgres",
		Password: "password",
		Host:     "localhost",
		Port:     5432,
		Name:     "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Create a test migration
	migrations := []Config{
		{
			Version: "001",
			Script:  "001_create_table.sql",
		},
		{
			Version: "002",
			Script:  "002_add_column.sql",
		},
	}

	// Migrate the database
	if err := Migrate(db, migrations); err != nil {
		t.Fatal(err)
	}

	// Verify that the migrations were applied
	var version string
	if err := db.QueryRow("SELECT version FROM migrations").Scan(&version); err != nil {
		t.Fatal(err)
	}

	if version != "002" {
		t.Errorf("expected version 002, got %s", version)
	}
}