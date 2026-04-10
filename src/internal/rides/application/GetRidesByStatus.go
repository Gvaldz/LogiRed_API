package application

import (
    "logired/src/internal/rides/domain"
    "logired/src/internal/rides/domain/entities"
)

type GetRidesByStatus struct {
    repo domain.IRide
}

func NewGetRidesByStatus(repo domain.IRide) *GetRidesByStatus {
    return &GetRidesByStatus{repo: repo}
}

func (uc *GetRidesByStatus) Execute(userID int32, userType int32, idstatus int32) ([]entities.Ride, error) {
    return uc.repo.GetRidesByStatus(userID, userType, idstatus)
}