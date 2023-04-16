package authDelivery

import (
	"log"

	authController "github.com/ReygaFitra/auth-jwt/controller"
	authDatabase "github.com/ReygaFitra/auth-jwt/database"
	authRepository "github.com/ReygaFitra/auth-jwt/repository"
	authUsecase "github.com/ReygaFitra/auth-jwt/usecase"
	authUtils "github.com/ReygaFitra/auth-jwt/utils"
	"github.com/gin-gonic/gin"
)

func Router() {
	db, err := authDatabase.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	serverPort := authUtils.DotEnv("SERVER_PORT")
	router := gin.Default()

	authRepo := authRepository.NewAuthRepo(db)
	authUsecase := authUsecase.NewAuthUsecase(authRepo)
	authCtrl := authController.NewAuthController(authUsecase)

	// login routes
	router.POST("/auth/login", authCtrl.Login)
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
