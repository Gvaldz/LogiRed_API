package application

import (
	"logired/src/core"
	"logired/src/internal/users/domain"
	"logired/src/internal/users/domain/entities"
)

type CreateUser struct {
	userRepo domain.UserRepository
	hasher   core.PasswordHasher
}

func NewCreateUser(userRepo domain.UserRepository, hasher core.PasswordHasher) *CreateUser {
	return &CreateUser{
		userRepo: userRepo,
		hasher:   hasher,
	}
}

func (uc *CreateUser) Execute(user entities.User) (entities.User, error) {
	hashedPassword, err := uc.hasher.Hash(user.Password)
	if err != nil {
		return entities.User{}, err
	}

	user.Password = hashedPassword

	createdUser, err := uc.userRepo.CreateUser(user)
	if err != nil {
		return entities.User{}, err
	}

	return createdUser, nil
}
