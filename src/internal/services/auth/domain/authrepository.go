package domain

import (
	user "logired/src/internal/users/domain/entities"
)

type AuthRepository interface {
	FindUserByEmail(email string) (user.User, error)
	UpdateLastLogin(userID int32) error
}
