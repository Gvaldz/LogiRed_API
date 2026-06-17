package application

import (
	"fmt"
	"logired/src/core"
	authDomain "logired/src/core/services/auth/domain"
	"logired/src/internal/users/domain"
)

type UpdatePassword struct {
	userRepo domain.UserRepository
	authRepo authDomain.AuthRepository
	hasher   core.PasswordHasher
}

func NewUpdatePassword(
	userRepo domain.UserRepository,
	authRepo authDomain.AuthRepository,
	hasher core.PasswordHasher,
) *UpdatePassword {
	return &UpdatePassword{
		userRepo: userRepo,
		authRepo: authRepo,
		hasher:   hasher,
	}
}

func (uc *UpdatePassword) Execute(id int32, oldPassword string, newPassword string) error {
	user, err := uc.authRepo.FindUserByID(id)
	if err != nil {
		return fmt.Errorf("usuario no encontrado o credenciales inválidas")
	}

	err = uc.hasher.Compare(user.Password, oldPassword)
	if err != nil {
		return fmt.Errorf("la contraseña actual es incorrecta")
	}

	hashedPassword, err := uc.hasher.Hash(newPassword)
	if err != nil {
		return err
	}

	return uc.userRepo.UpdatePassword(user.IdUser, hashedPassword)
}
