package controllers

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	
	auth struct {
		Email    string `gorm:"type:varchar(255)" json:"email" binding:"required"`
		Password string `gorm:"type:varchar(255)" json:"password" binding:"required"`
	}

	Claims struct {
		Email string `json:"email"`
		Names string `json:"names"`
		jwt.StandardClaims
	}
)

var (
	credentials = auth{}

	jwtKey = []byte("redmouse_2022")

	authUser = models.User{}

	WrongCredential = gin.H{"Error": "Wrong credentials"}
)

func Login(context *gin.Context) {

	if err := context.BindJSON(&credentials); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	result := db.Conect().Where(&models.User{Email: credentials.Email}).First(&authUser)
	if result.RowsAffected == 0 {
		context.JSON(200, WrongCredential)
		return
	}

	if result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}

	isValidPassword := models.CheckPasswordHash(credentials.Password, authUser.Password)

	if isValidPassword {
		token, expirationTime, err := generateToken(authUser)
		if err != nil {
			context.JSON(200, gin.H{"Error": "Something went wrong"})
			return
		}
		context.JSON(200, gin.H{"token": token, "expiration_time": expirationTime})
		return
	}

	context.JSON(200, WrongCredential)
}

func generateToken(user models.User) (string, int64, error) {

	expirationTime := time.Now().Add(72 * time.Hour)

	claims := &Claims{
		Email: user.Email,
		Names: user.Names,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime.Unix(), nil
}
