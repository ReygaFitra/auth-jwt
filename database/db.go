package authDatabase

import (
	"database/sql"
	"fmt"
	"log"

	authUtils "github.com/ReygaFitra/auth-jwt/utils"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dbHost := authUtils.DotEnv("DB_HOST")
	dbPort := authUtils.DotEnv("DB_PORT")
	dbUser := authUtils.DotEnv("DB_USER")
	dbPassword := authUtils.DotEnv("DB_PASSWORD")
	dbName := authUtils.DotEnv("DB_NAME")
	sslMode := authUtils.DotEnv("SSL_MODE")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}