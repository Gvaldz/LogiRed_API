package application

import (
	"logired/src/internal/users/domain"
	"logired/src/internal/users/domain/entities"
)

type GetUserByID struct {
	repo domain.UserRepository
}

func NewGetUserByID(repo domain.UserRepository) *GetUserByID {
	return &GetUserByID{repo: repo}
}

func (cp *GetUserByID) Execute(IDuser int32) (entities.User, error) {
	return cp.repo.GetUserByID(IDuser)
}
