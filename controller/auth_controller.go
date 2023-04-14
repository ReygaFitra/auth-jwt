package authController

import (
	"net/http"
	"time"

	authUtils "github.com/ReygaFitra/auth-jwt/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func Login(ctx *gin.Context) {
	secretKey := authUtils.DotEnv("SECRET_KEY")
	var jwtKey = []byte(secretKey)
	
	var user authModel.Credential

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if user.Username == secretKey {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

		tokenString, err :=token.SignedString(jwtKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed generate token!"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unregistered student!"})
	}
}

func Profile(ctx *gin.Context) {
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

