package domain

import (
	"database/sql"
	"logired/src/internal/drivers/domain/entities"
)

type DriverDetail struct {
	IdUser   int32   `json:"id_user"`
	Rating   float32 `json:"rating"`
	Image    string  `json:"image"`
	Name     string  `json:"name"`
	Lastname string  `json:"lastname"`
	Email    string  `json:"email"`
}

type DriverStats struct {
	DriverID             int32   `json:"driver_id"`
	Rating               float32 `json:"rating"`
	TotalTripsCompleted  int32   `json:"total_trips_completed"`
	TotalTripsCancelled  int32   `json:"total_trips_cancelled"`
	VehicleCount         int32   `json:"vehicle_count"`
	DaysInPlatform       int32   `json:"days_in_platform"`
	ResponseTimeMinutes  float64 `json:"response_time_minutes"`
	CreatedAt            string  `json:"created_at"`
	DocumentsVerified    bool    `json:"documents_verified"`
}

type IDriver interface {
	Exists(userID int32) (bool, error)
	CreateTx(tx *sql.Tx, driver entities.Driver) error
	GetApprovedDrivers() ([]DriverDetail, error)
	GetDriverStats(driverID int32) (*DriverStats, error)
	/*
	GetDriversByCity(city string) ([]DriverDetail, error)
	UpdateCitywork(driverID int32, citywork string) error
	*/
}