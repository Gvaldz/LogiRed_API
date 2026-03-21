package infrastructure

import (
	"database/sql"
	"logired/src/core"
	"logired/src/internal/users/application"
	"logired/src/internal/users/infrastructure/controllers"

	entities_auth "logired/src/internal/services/auth/domain"
	entities_users "logired/src/internal/users/domain"
	middleware "logired/src/server/middleware"
)

type UserDependencies struct {
	DB           *sql.DB
	Hasher       *core.BcryptHasher
	UserRepo     entities_users.UserRepository
	AuthRepo     entities_auth.AuthRepository
	TokenService *core.JWTService
}

func NewUserDependencies(db *sql.DB, hasher *core.BcryptHasher, tokenService *core.JWTService, authRepo entities_auth.AuthRepository, userRepo entities_users.UserRepository) *UserDependencies {

	return &UserDependencies{
		DB:           db,
		Hasher:       hasher,
		TokenService: tokenService,
		AuthRepo:     authRepo,
		UserRepo:     userRepo,
	}
}

func (d *UserDependencies) GetRoutes() *UserRoutes {
	createUserUseCase := application.NewCreateUser(d.UserRepo, d.Hasher)
	getAllUserUseCase := application.NewGetAllUsers(d.UserRepo)
	getUserUseCase := application.NewGetUserByID(d.UserRepo)
	updateUserUseCase := application.NewUpdateUser(d.UserRepo)
	updatePasswordUseCase := application.NewUpdatePassword(d.UserRepo, d.Hasher)
	deleteUserUseCase := application.NewDeleteUser(d.UserRepo)

	createUserController := controllers.NewCreateUserController(createUserUseCase)
	getUsersController := controllers.NewGetAllUsersController(getAllUserUseCase)
	getUserController := controllers.NewGetByUserIDController(getUserUseCase)
	updateUserController := controllers.NewUpdateUserController(updateUserUseCase)
	updatePasswordController := controllers.NewUpdatePasswordController(updatePasswordUseCase)
	deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)

	authMiddleware := middleware.AuthMiddleware(d.TokenService, d.UserRepo)

	return NewUserRoutes(
		createUserController,
		getUsersController,
		getUserController,
		updateUserController,
		updatePasswordController,
		deleteUserController,
		authMiddleware,
	)
}
