package application

import (
    "logired/src/internal/users/domain"
    "logired/src/internal/users/domain/entities"
)

type GetUserProfile struct {
    userRepo domain.UserRepository
}

func NewGetUserProfile(userRepo domain.UserRepository) *GetUserProfile {
    return &GetUserProfile{userRepo: userRepo}
}

func (uc *GetUserProfile) Execute(id int32) (entities.UserProfile, error) {
    return uc.userRepo.GetUserProfileByID(id)
}