package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ABHINAV-JHA-27/scheduleX/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

const create string = `
  CREATE TABLE IF NOT EXISTS schedulex_jobs (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    cron_sequence TEXT NOT NULL,
    command TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
  );`

var DB *sql.DB

func ensureDirectories() {
	path := config.Cfg.Database.Path
	dirPath := filepath.Dir(path)

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	fmt.Printf("Directory '%s' ensured to exist.\n", dirPath)
}

func InitDB() error {
	ensureDirectories()

	db, err := sql.Open("sqlite3", config.Cfg.Database.Path)
	if err != nil {
		return fmt.Errorf("error opening jobs database: %s", err)
	}
	if _, err := db.Exec(create); err != nil {
		return fmt.Errorf("error creating jobs database: %s", err)
	}

	DB = db
	fmt.Println("Database Connected Successfully.")
	return nil
}

func CloseDB() error {
	err := DB.Close()
	if err != nil {
		return fmt.Errorf("error closing database connection: %s", err)
	}
	return nil
}
