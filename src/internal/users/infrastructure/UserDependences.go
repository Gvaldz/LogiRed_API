package infrastructure

import (
    "database/sql"
    "logired/src/core"
    storageService      "logired/src/core/services/storage"
    entities_auth       "logired/src/core/services/auth/domain"
    middleware          "logired/src/server/middleware"
    
    entities_users      "logired/src/internal/users/domain"
    drivers             "logired/src/internal/drivers/domain"
    cars                "logired/src/internal/cars/domain"
    documents           "logired/src/internal/documents/domain"
    
    "logired/src/internal/users/application"
    "logired/src/internal/users/infrastructure/controllers"
)

type UserDependencies struct {
    DB             *sql.DB
    Hasher         *core.BcryptHasher
    UserRepo       entities_users.UserRepository
    DriverRepo     drivers.IDriver
    CarRepo        cars.ICar             
    DocumentRepo   documents.IDocument   
    AuthRepo       entities_auth.AuthRepository
    TokenService   *core.JWTService
    StorageService storageService.StorageService
}

func NewUserDependencies(
    db             *sql.DB,
    hasher         *core.BcryptHasher,
    tokenService   *core.JWTService,
    authRepo       entities_auth.AuthRepository,
    userRepo       entities_users.UserRepository,
    driverRepo     drivers.IDriver,
    carRepo        cars.ICar,             
    documentRepo   documents.IDocument,   
    storageService storageService.StorageService,
) *UserDependencies {
    return &UserDependencies{
        DB:             db,
        Hasher:         hasher,
        TokenService:   tokenService,
        AuthRepo:       authRepo,
        UserRepo:       userRepo,
        DriverRepo:     driverRepo,
        CarRepo:        carRepo,         
        DocumentRepo:   documentRepo,    
        StorageService: storageService,
    }
}

func (d *UserDependencies) GetRoutes() *UserRoutes {
    createUserUseCase := application.NewCreateUser(d.UserRepo, d.Hasher)
    getAllUserUseCase := application.NewGetAllUsers(d.UserRepo)
    getUserUseCase := application.NewGetUserByID(d.UserRepo)
    getUserProfileUseCase := application.NewGetUserProfile(d.UserRepo)
    updateUserUseCase := application.NewUpdateUser(d.UserRepo)
    updatePasswordUseCase := application.NewUpdatePassword(d.UserRepo, d.AuthRepo, d.Hasher)
    ressetPasswordUseCase := application.NewRessetPassword(d.UserRepo, d.AuthRepo, d.Hasher)
    deleteUserUseCase := application.NewDeleteUser(d.UserRepo)
    
    // AQUÍ le pasamos los nuevos repositorios al caso de uso
    createDriverUseCase := application.NewRegisterDriver(d.UserRepo, d.DriverRepo, d.CarRepo, d.DocumentRepo, d.Hasher)

    createUserController := controllers.NewCreateUserController(createUserUseCase, createDriverUseCase, d.StorageService)
    getUsersController := controllers.NewGetAllUsersController(getAllUserUseCase)
    getUserController := controllers.NewGetByUserIDController(getUserUseCase)
    getUserProfileController := controllers.NewGetUserProfileController(getUserProfileUseCase)
    updateUserController := controllers.NewUpdateUserController(updateUserUseCase, d.StorageService)
    updatePasswordController := controllers.NewUpdatePasswordController(updatePasswordUseCase)
    ressetPasswordController := controllers.NewRessetPasswordController(ressetPasswordUseCase)
    deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)

    authMiddleware := middleware.AuthMiddleware(d.TokenService, d.UserRepo)

    return NewUserRoutes(
        createUserController,
        getUsersController,
        getUserController,
        getUserProfileController,
        updateUserController,
        updatePasswordController,
        ressetPasswordController,
        deleteUserController,
        authMiddleware,
    )
}