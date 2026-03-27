package application

import "logired/src/internal/rides/domain"

type UpdateRideStatus struct {
	repo domain.IRide
}

func NewUpdateRideStatus(repo domain.IRide) *UpdateRideStatus {
	return &UpdateRideStatus{repo: repo}
}

func (urs *UpdateRideStatus) Execute(idRide int32, idStatus int32) error {
	return urs.repo.UpdateRideStatus(idRide, idStatus)
}