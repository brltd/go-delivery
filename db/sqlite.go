package db

import (
	"fmt"
	"os"

	"github.com/brltd/delivery/logger"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func GetSQLiteClient() *sqlx.DB {
	dbPath := os.Getenv("SQL_LITE_DB_PATH")

	if dbPath == "" {
		logger.Error("SQL_LITE_DB_PATH environment variable is not set")
		return nil
	}

	dataSource := fmt.Sprintf("%s?_foreign_keys=on", dbPath) // Enable foreign key support
	client, err := sqlx.Open("sqlite", dataSource)

	if err != nil {
		logger.Error(fmt.Sprintf("Error initializing SQLite database: %+v", err))
		return nil
	}

	return client
}
