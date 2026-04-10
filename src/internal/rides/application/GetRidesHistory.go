package application

import (
    "logired/src/internal/rides/domain"
    "logired/src/internal/rides/domain/entities"
)

type GetRidesHistory struct {
    repo domain.IRide
}

func NewGetRidesHistory(repo domain.IRide) *GetRidesHistory {
    return &GetRidesHistory{repo: repo}
}

func (uc *GetRidesHistory) Execute(userID int32, userType int32) ([]entities.Ride, error) {
    return uc.repo.GetRidesHistory(userID, userType)
}