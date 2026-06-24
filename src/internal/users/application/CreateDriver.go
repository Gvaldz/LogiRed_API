package application

import (
    "fmt"
    "logired/src/core"
    driverDomain "logired/src/internal/drivers/domain"
    driverEntities "logired/src/internal/drivers/domain/entities"
    userDomain "logired/src/internal/users/domain"
    userEntities "logired/src/internal/users/domain/entities"
    
    carDomain "logired/src/internal/cars/domain"
    carEntities "logired/src/internal/cars/domain/entities"

    documentDomain "logired/src/internal/documents/domain"
    documentEntities "logired/src/internal/documents/domain/entities"
)

type RegisterDriverInput struct {
    User      userEntities.User
    Car       carEntities.Car
    Documents []documentEntities.Document 
}

type RegisterDriver struct {
    userRepo     userDomain.UserRepository
    driverRepo   driverDomain.IDriver
    carRepo      carDomain.ICar
    documentRepo documentDomain.IDocument 
    hasher       core.PasswordHasher
}

func NewRegisterDriver(
    userRepo userDomain.UserRepository,
    driverRepo driverDomain.IDriver,
    carRepo carDomain.ICar,
    documentRepo documentDomain.IDocument,
    hasher core.PasswordHasher,
) *RegisterDriver {
    return &RegisterDriver{
        userRepo:     userRepo,
        driverRepo:   driverRepo,
        carRepo:      carRepo,
        documentRepo: documentRepo,
        hasher:       hasher,
    }
}

func (uc *RegisterDriver) Execute(input RegisterDriverInput) (userEntities.User, error) {
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
        IdUser: createdUser.IdUser,
        Rating: 0,
    }
    if err := uc.driverRepo.CreateTx(tx, driver); err != nil {
        tx.Rollback()
        return userEntities.User{}, err
    }

    input.Car.IdDriver = createdUser.IdUser
    if err := uc.carRepo.CreateCarTx(tx, input.Car); err != nil {
        tx.Rollback()
        return userEntities.User{}, fmt.Errorf("error al registrar el vehículo: %w", err)
    }

    for _, doc := range input.Documents {
        doc.IdDriver = createdUser.IdUser 
        if err := uc.documentRepo.CreateDocumentTx(tx, doc); err != nil {
            tx.Rollback()
            return userEntities.User{}, fmt.Errorf("error al registrar documento %s: %w", doc.IdDocumentType, err)
        }
    }

    if err := tx.Commit(); err != nil {
        return userEntities.User{}, fmt.Errorf("error al confirmar transacción: %w", err)
    }

    return createdUser, nil
}