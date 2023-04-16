package handler

import (
	"net/http"
	"time"

	"github.com/ReygaFitra/app-mahasiswa-api/model"
	authDatabase "github.com/ReygaFitra/auth-jwt/database"
	authModel "github.com/ReygaFitra/auth-jwt/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string	`json:"username"`
	jwt.StandardClaims
}

func LoginHandler(ctx *gin.Context) {
	var user authModel.Credential

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	db,err := authDatabase.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var userLogin authModel.Credential
	err = db.QueryRow("SELECT user_name, password FROM credentials WHERE user_name = $1", user.Username).Scan(&userLogin.Username, &userLogin.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized User!"})
		return
	}

	if userLogin.Username != user.Username {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Username!"})
		return
	}

	var getUser model.Student
	err = db.QueryRow("SELECT id, name, age major, student_user_name FROM student WHERE id = $1", getUser.Id).Scan(&getUser.Id, &getUser.Name, &getUser.Age, &getUser.Major, &getUser.StudentUserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error Server"})
		return
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: getUser.StudentUserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
	})

	tokenString,err := token.SignedString(authModel.JwtKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "token generate error!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func MiddlewareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return authModel.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("claims", claims)

		c.Next()
	}
}

func ProfileHandler(ctx *gin.Context) {
	db,err := authDatabase.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

		// ambil username dari JWT token
		claims := ctx.MustGet("claims").(jwt.MapClaims)
		username := claims["username"].(string)
	
		// dapatkan informasi user dari database (dalam hal ini, kita return username)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to profile",
			"username": username,
		})
}