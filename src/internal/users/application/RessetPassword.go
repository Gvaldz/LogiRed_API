package application

import (
	"fmt"
	"logired/src/core"
	authDomain "logired/src/core/services/auth/domain"
	"logired/src/internal/users/domain"
)

type RessetPassword struct {
	userRepo domain.UserRepository
	authRepo authDomain.AuthRepository
	hasher   core.PasswordHasher
}

func NewRessetPassword(
	userRepo domain.UserRepository,
	authRepo authDomain.AuthRepository,
	hasher core.PasswordHasher,
) *RessetPassword {
	return &RessetPassword{
		userRepo: userRepo,
		authRepo: authRepo,
		hasher:   hasher,
	}
}

func (uc *RessetPassword) Execute(email, newPassword string) error {
	_, err := uc.authRepo.FindUserByEmail(email)
	if err != nil {
		return fmt.Errorf("el correo electrónico no está registrado")
	}

	hashedPassword, err := uc.hasher.Hash(newPassword)
	if err != nil {
		return err
	}

	return uc.userRepo.RessetPassword(email, hashedPassword)
}
