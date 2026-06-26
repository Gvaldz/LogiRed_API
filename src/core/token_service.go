package core

import (
	"errors"
	"logired/src/core/services/auth/domain"
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

func (s *JWTService) GenerateToken(userID int32, email string, usertype int, approved bool) (domain.Token, error) {
    expiresAt := time.Now().Add(100 * 365 * 24 * time.Hour).Unix()

    claims := jwt.MapClaims{
        "user_id":  userID,
        "email":    email,
        "exp":      expiresAt,
        "usertype": usertype,
    }

    if usertype == 2 {
        claims["approved"] = approved
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

// ExtractUserType - Extrae el tipo de usuario del token
func (s *JWTService) ExtractUserType(tokenString string) (int, error) {
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
		if usertype, exists := claims["usertype"]; exists {
			return int(usertype.(float64)), nil
		}
		return 0, errors.New("usertype not found in token")
	}

	return 0, errors.New("invalid token")
}

// ExtractApproved - Extrae el estado de aprobación del token (solo para conductores)
func (s *JWTService) ExtractApproved(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if approved, exists := claims["approved"]; exists {
			return approved.(bool), nil
		}
		// Si no existe "approved" en el token, es un cliente (no conductor)
		return false, nil
	}

	return false, errors.New("invalid token")
}
