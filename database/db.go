package database

import (
	"database/sql"
	"fmt"
	"log"

	authController "github.com/ReygaFitra/auth-jwt/controller"
	authUtils "github.com/ReygaFitra/auth-jwt/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func ConnectDB() {
	dbHost := authUtils.DotEnv("DB_HOST")
	dbPort := authUtils.DotEnv("DB_PORT")
	dbUser := authUtils.DotEnv("DB_USER")
	dbPassword := authUtils.DotEnv("DB_PASSWORD")
	dbName := authUtils.DotEnv("DB_NAME")
	sslMode := authUtils.DotEnv("SSL_MODE")
	serverPort := authUtils.DotEnv("SERVER_PORT")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := gin.Default()

	// login routes
	router.POST("/auth/login", authController.Login)
	// register routes
	// router.POST("/auth/register",register)

	// // students routes
	studentRouter := router.Group("/api/v1/students/")
	studentRouter.Use(authController.AuthMiddleware())

	// studentRouter.GET("", getall)
	// studentRouter.GET("/:id", getbyid)
	// studentRouter.PUT("", update)
	// studentRouter.DELETE("/:id", delete)


	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}