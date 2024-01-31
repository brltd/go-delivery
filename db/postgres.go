package db

import (
	"fmt"
	"os"

	"github.com/brltd/delivery/logger"
	"github.com/jmoiron/sqlx"
)

func GetPostgresClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbAddress := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("user=%s password=%s "+
		"host=%s port=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbAddress, dbPort, dbName)

	client, err := sqlx.Open("postgres", dataSource)

	if err != nil {
		logger.Error(fmt.Sprintf("Error initializing database: %+v", err))
		return nil
	}

	return client
}
