package application

import (
	"logired/src/internal/users/domain"
	"logired/src/internal/users/domain/entities"
)

type GetAllUsers struct {
	userRepo domain.UserRepository
}

func NewGetAllUsers(userRepo domain.UserRepository) *GetAllUsers {
	return &GetAllUsers{userRepo: userRepo}
}

func (lp *GetAllUsers) Execute() ([]entities.User, error) {
	return lp.userRepo.GetAllUsers()
}
