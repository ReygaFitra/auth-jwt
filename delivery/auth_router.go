package authDelivery

import (
	"log"

	authController "github.com/ReygaFitra/auth-jwt/controller"
	authUtils "github.com/ReygaFitra/auth-jwt/utils"
	"github.com/gin-gonic/gin"
	// _ "github.com/lib/pq"
)


func AuthRouter() {
	serverPort := authUtils.DotEnv("SERVER_PORT")

	router := gin.Default()

	// login routes
	router.POST("/auth/login", authController.Login)
	// register routes
	// router.POST("/auth/register",register)

	// // students routes
	// studentRouter := router.Group("/api/v1/students/")
	// studentRouter.Use(authController.AuthMiddleware())

	// studentRouter.GET("", getall)
	// studentRouter.GET("/:id", getbyid)
	// studentRouter.PUT("", update)
	// studentRouter.DELETE("/:id", delete)


	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}