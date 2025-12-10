package util

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	return token, claims, nil
}

// Extract user ID from JWT claims
func GetUserIDFromToken(claims jwt.MapClaims) (string, error) {
	sub, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid sub claim")
	}
	return sub, nil
}

// Extract role ID
func GetRoleFromToken(claims jwt.MapClaims) (string, error) {
	role, ok := claims["role"].(string)
	if !ok {
		return "", errors.New("invalid role claim")
	}
	return role, nil
}