# Minor edit
package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/yourusername/dbmigrate/config"
	"github.com/yourusername/dbmigrate/database"
	"github.com/yourusername/dbmigrate/migration"
)

func main() {
	configPath := flag.String("config", "config.json", "path to configuration file")
	flag.Parse()

	// Load configuration from file
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate database
	if err := migration.Migrate(db, cfg.Migrations); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database migration complete")
}