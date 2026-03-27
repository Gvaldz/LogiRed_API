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

type IDriver interface {
	GetDriversByCity(city string) ([]DriverDetail, error)
	UpdateCitywork(driverID int32, citywork string) error
	Exists(userID int32) (bool, error)
	CreateTx(tx *sql.Tx, driver entities.Driver) error
}