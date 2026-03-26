package infrastructure

import (
	"database/sql"
	"logired/src/core"
	"logired/src/internal/users/application"
	"logired/src/internal/users/infrastructure/controllers"
	drivers "logired/src/internal/drivers/domain"
	driverApp "logired/src/internal/drivers/application"
	entities_auth "logired/src/internal/services/auth/domain"
	entities_users "logired/src/internal/users/domain"
	middleware "logired/src/server/middleware"
)

type UserDependencies struct {
	DB           *sql.DB
	Hasher       *core.BcryptHasher
	UserRepo     entities_users.UserRepository
	DriverRepo   drivers.IDriver
	AuthRepo     entities_auth.AuthRepository
	TokenService *core.JWTService
}

func NewUserDependencies(
	db *sql.DB,
	hasher *core.BcryptHasher,
	tokenService *core.JWTService,
	authRepo entities_auth.AuthRepository,
	userRepo entities_users.UserRepository,
	driverRepo drivers.IDriver,
) *UserDependencies {
	return &UserDependencies{
		DB:           db,
		Hasher:       hasher,
		TokenService: tokenService,
		AuthRepo:     authRepo,
		UserRepo:     userRepo,
		DriverRepo:   driverRepo,
	}
}

func (d *UserDependencies) GetRoutes() *UserRoutes {
	createUserUseCase        := application.NewCreateUser(d.UserRepo, d.Hasher)
	getAllUserUseCase         := application.NewGetAllUsers(d.UserRepo)
	getUserUseCase            := application.NewGetUserByID(d.UserRepo)
	getUserProfileUseCase     := application.NewGetUserProfile(d.UserRepo)
	updateUserUseCase         := application.NewUpdateUser(d.UserRepo)
	updateDriverProfileUseCase := driverApp.NewUpdateDriverProfile(d.DriverRepo) 
	updatePasswordUseCase     := application.NewUpdatePassword(d.UserRepo, d.Hasher)
	deleteUserUseCase         := application.NewDeleteUser(d.UserRepo)
	createDriverUseCase       := application.NewRegisterDriver(d.UserRepo, d.DriverRepo, d.Hasher)

	createUserController      := controllers.NewCreateUserController(createUserUseCase, createDriverUseCase)
	getUsersController         := controllers.NewGetAllUsersController(getAllUserUseCase)
	getUserController          := controllers.NewGetByUserIDController(getUserUseCase)
	getUserProfileController   := controllers.NewGetUserProfileController(getUserProfileUseCase)   
	updateUserController       := controllers.NewUpdateUserController(updateUserUseCase, updateDriverProfileUseCase) 
	updatePasswordController   := controllers.NewUpdatePasswordController(updatePasswordUseCase)
	deleteUserController       := controllers.NewDeleteUserController(deleteUserUseCase)

	authMiddleware := middleware.AuthMiddleware(d.TokenService, d.UserRepo)

	return NewUserRoutes(
		createUserController,
		getUsersController,
		getUserController,
		getUserProfileController,  
		updateUserController,
		updatePasswordController,
		deleteUserController,
		authMiddleware,
	)
}