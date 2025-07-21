package helper

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GenerateToken(userID uint) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func GetUserFromToken(c echo.Context) (uint, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("authorization header not found")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id not found in token")
	}

	return uint(userIDFloat), nil
}

func GetUserIDFromToken(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)         // Tokenni kontekstdan olamiz
	claims := user.Claims.(jwt.MapClaims)      // Tokenni claims (ma'lumotlar) qismiga ajratamiz
	userID := int(claims["user_id"].(float64)) // `user_id` ni o‘qib olamiz

	return userID, nil
}
