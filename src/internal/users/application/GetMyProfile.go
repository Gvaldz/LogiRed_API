package application

import (
    "logired/src/internal/users/domain"
    "logired/src/internal/users/domain/entities"
)

type GetMyProfile struct {
    userRepo domain.UserRepository
}

func NewGetMyProfile(userRepo domain.UserRepository) *GetMyProfile {
    return &GetMyProfile{userRepo: userRepo}
}

func (uc *GetMyProfile) Execute(id int32) (entities.MyProfile, error) {
    return uc.userRepo.GetMyProfileByID(id)
}
