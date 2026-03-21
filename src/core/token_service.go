package core

import (
	"errors"
	"logired/src/internal/services/auth/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey []byte
}

func NewJWTService() *JWTService {
	key := []byte(os.Getenv("JWT_SECRET"))
	if len(key) == 0 {
		panic("JWT_SECRET not set")
	}
	return &JWTService{secretKey: key}
}

func (s *JWTService) GenerateToken(userID int32, email string, usertype int) (domain.Token, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	claims := jwt.MapClaims{
		"user_id": 	 userID,
		"email":   	 email,
		"exp":    	 expiresAt,
		"usertype":  usertype,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return domain.Token{}, err
	}

	return domain.Token{
		Token:     tokenString,
		ExpiresAt: expiresAt,
	}, nil
}

func (s *JWTService) ValidateToken(tokenString string) (int32, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int32(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, errors.New("invalid token")
}
