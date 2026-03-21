package domain

import "logired/src/internal/users/domain/entities"

type UserRepository interface {
	CreateUser(entities.User) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetUserByID(IdUsuario int32) (entities.User, error)
	UpdateUser(IdUsuario int32, user entities.User) error
	UpdatePassword(IdUsuario int32, password string) error
	DeleteUser(IdUsuario int32) error
}
