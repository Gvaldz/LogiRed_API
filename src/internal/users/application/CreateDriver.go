package application

import (
	"fmt"
	"logired/src/core"
	driverDomain "logired/src/internal/drivers/domain"
	driverEntities "logired/src/internal/drivers/domain/entities"
	userDomain "logired/src/internal/users/domain"
	userEntities "logired/src/internal/users/domain/entities"
)

type RegisterDriver struct {
	userRepo   userDomain.UserRepository
	driverRepo driverDomain.IDriver
	hasher     core.PasswordHasher
}

func NewRegisterDriver(
	userRepo userDomain.UserRepository,
	driverRepo driverDomain.IDriver,
	hasher core.PasswordHasher,
) *RegisterDriver {
	return &RegisterDriver{
		userRepo:   userRepo,
		driverRepo: driverRepo,
		hasher:     hasher,
	}
}

func (uc *RegisterDriver) Execute(input userEntities.RegisterDriverInput) (userEntities.User, error) {
	tx, err := uc.userRepo.BeginTx()
	if err != nil {
		return userEntities.User{}, fmt.Errorf("error al iniciar transacción: %w", err)
	}

	hashed, err := uc.hasher.Hash(input.User.Password)
	if err != nil {
		tx.Rollback()
		return userEntities.User{}, err
	}
	input.User.Password = hashed

	createdUser, err := uc.userRepo.CreateUserTx(tx, input.User)
	if err != nil {
		tx.Rollback() 
		return userEntities.User{}, err
	}

	driver := driverEntities.Driver{
		IdUser:   createdUser.IdUser,
		Citywork: input.Citywork,
		Rating:   0,
	}
	if err := uc.driverRepo.CreateTx(tx, driver); err != nil {
		tx.Rollback()
		return userEntities.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return userEntities.User{}, fmt.Errorf("error al confirmar transacción: %w", err)
	}

	return createdUser, nil
}
