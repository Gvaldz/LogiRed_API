package domain

import "logired/src/internal/drivers/domain/entities"

type DriverDetail struct {
	IdUser   int32   `json:"id_user"`
	Rating   float32 `json:"rating"`
	Image    string  `json:"image"`
	Name     string  `json:"name"`
	Lastname string  `json:"lastname"`
	Email    string  `json:"email"`
}

type IDriver interface {
	Create(driver entities.Driver) error
	GetByUserID(userID int32) (*DriverDetail, error)
	GetByID(driverID int32) (*DriverDetail, error) 
	GetAll() ([]DriverDetail, error)
	Update(driver entities.Driver) error
	Delete(userID int32) error
	Exists(userID int32) (bool, error)
}