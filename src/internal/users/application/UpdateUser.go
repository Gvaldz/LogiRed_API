package application

import (
	"logired/src/internal/users/domain"
	"logired/src/internal/users/domain/entities"
)

type UpdateUser struct {
	userRepo domain.UserRepository
}

func NewUpdateUser(userRepo domain.UserRepository) *UpdateUser {
	return &UpdateUser{userRepo: userRepo}
}

func (uc *UpdateUser) Execute(id int32, user entities.User) error {
	user.Password = ""
	return uc.userRepo.UpdateUser(id, user)
}
