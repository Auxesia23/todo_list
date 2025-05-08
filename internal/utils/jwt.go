package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/Auxesia23/todo_list/internal/models"
	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte

func GenerateToken(user *models.User) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	jwtSecret = []byte(secretKey)
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	secretKey := os.Getenv("SECRET_KEY")
	jwtSecret = []byte(secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
