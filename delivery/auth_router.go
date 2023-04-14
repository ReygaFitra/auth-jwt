package authDelivery

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func AuthRouter() {
	
	serverPort := utils.DotEnv("SERVER_PORT")

	router := gin.Default()

	// login routes
	router.POST("/auth/login", controller.Login)
	// register routes
	router.POST("/api/v1/students",register)

	// students routes
	studentRouter := router.Group("/api/v1/students/profile")
	studentRouter.Use(controller.AuthMiddleware())

	studentRouter.GET("", getall)
	studentRouter.GET("/:id", getbyid)
	studentRouter.PUT("", update)
	studentRouter.DELETE("/:id", delete)


	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}