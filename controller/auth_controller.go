package authController

import (
	"net/http"
	"time"

	authModel "github.com/ReygaFitra/auth-jwt/model"
	authUsecase "github.com/ReygaFitra/auth-jwt/usecase"
	authUtils "github.com/ReygaFitra/auth-jwt/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUsecase authUsecase.AuthUsecase
}

func (c *AuthController) Login(ctx *gin.Context) {
	var user authModel.Credential

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if user.Username == "secretkey" && user.Password == "password" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

		tokenString, err :=token.SignedString(authModel.JwtKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed generate token!"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unregistered student!"})
	}

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}
	res := c.authUsecase.SignIn(&user)
	ctx.JSON(http.StatusCreated, res)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var newUser authModel.Credential

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if newUser.Username == "secretkey" && newUser.Password == "password" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = newUser.Username
		claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

		tokenString, err :=token.SignedString(authModel.JwtKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed generate token!"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unregistered student!"})
	}

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.authUsecase.SignUp(&newUser)
	ctx.JSON(http.StatusCreated, res)
}

func (c *AuthController) Profile(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome Student",
		"username": username,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	secretKey := authUtils.DotEnv("SECRET_KEY")
	var jwtKey = []byte(secretKey)

	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}

func NewAuthController(u authUsecase.AuthUsecase) *AuthController {
	controller := AuthController{
		authUsecase: u,
	}

	return &controller
}
