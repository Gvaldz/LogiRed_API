package application

import (
	"logired/src/core"
	"logired/src/internal/users/domain"
	"logired/src/internal/users/domain/entities"
)

type UpdatePassword struct {
	userRepo domain.UserRepository
	hasher   core.PasswordHasher
}

func NewUpdatePassword(userRepo domain.UserRepository, hasher core.PasswordHasher) *UpdatePassword {
	return &UpdatePassword{
		userRepo: userRepo,
		hasher:   hasher,
	}
}

func (uc *UpdatePassword) Execute(id int32, newPassword string) error {
	hashedPassword, err := uc.hasher.Hash(newPassword)
	if err != nil {
		return err
	}

	user := entities.User{
		IdUser:   id,
		Password: hashedPassword,
	}

	return uc.userRepo.UpdatePassword(id, user.Password)
}
