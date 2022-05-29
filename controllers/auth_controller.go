package controllers

import (
	// "bytes"
	// "encoding/json"
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

	wrongCredential = gin.H{"Error": "Wrong credentials"}

	tokenInvalid = errors.New("Invalid token")
)

func Login(context *gin.Context) {

	if err := context.BindJSON(&credentials); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		RespondWithError(context, 500, err.Error())
	}

	result := db.Conect().Where(&models.User{Email: credentials.Email}).First(&authUser)
	if result.RowsAffected == 0 {
		context.JSON(401, wrongCredential)
		return
	}

	if result.Error != nil {
		context.JSON(401, gin.H{"Error": result.Error})
		return
	}

	isValidPassword := models.CheckPasswordHash(credentials.Password, authUser.Password)

	if isValidPassword {
		token, expirationTime, err := generateToken(authUser)
		if err != nil {
			context.JSON(500, gin.H{"Error": "Something went wrong"})
			return
		}
		context.JSON(200, gin.H{"token": token, "expires_at": expirationTime})
		return
	}

	context.JSON(401, wrongCredential)
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

func checkValidToken(token string) (models.User, error) {

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return models.User{}, err
		}
	}

	if !tkn.Valid {
		return models.User{}, tokenInvalid
	}

	result := db.Conect().Where(&models.User{Email: claims.Email}).First(&authUser)
	if result.RowsAffected == 0 {
		return models.User{}, tokenInvalid
	}

	if result.Error != nil {
		return models.User{}, tokenInvalid
	}

	if claims.ExpiresAt <= time.Now().Unix() {
		return models.User{}, tokenInvalid
	}

	return authUser, nil
}

func AuthMiddleware() gin.HandlerFunc {

	return func(context *gin.Context) {

		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.JSON(401, gin.H{"Error": "Not authorized"})
			context.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]

		if tokenString == "" {
			context.JSON(401, gin.H{"Error": "Not authorized"})
			context.Abort()
			return
		}

		_, err := checkValidToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"Error": err.Error()})
			context.Abort()
			return
		}

		context.Next()

		// reqBodyBytes := new(bytes.Buffer)
		// json.NewEncoder(reqBodyBytes).Encode(gin.H{"user": authUser})

		// reqBodyBytes.Bytes()

		// context.Data(200, "application/json", reqBodyBytes.Bytes())

	}
}
