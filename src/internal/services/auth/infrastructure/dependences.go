package infrastructure

import (
	"database/sql"
	"logired/src/core"
	"logired/src/internal/services/auth/application"
	"logired/src/internal/services/auth/infrastructure/controllers"
	users_domain "logired/src/internal/users/domain"
)

type AuthDependencies struct {
	DB       *sql.DB
	Hasher   *core.BcryptHasher
	UserRepo users_domain.UserRepository
}

func NewAuthDependencies(db *sql.DB, hasher *core.BcryptHasher, userRepo users_domain.UserRepository) *AuthDependencies {
	return &AuthDependencies{
		DB:       db,
		Hasher:   hasher,
		UserRepo: userRepo,
	}
}

func (d *AuthDependencies) GetRoutes() *AuthRoutes {
	authRepo := NewAuthRepository(d.DB)
	tokenService := core.NewJWTService()

	loginUC := application.NewLogin(
		authRepo,
		d.UserRepo,
		tokenService,
		d.Hasher,
	)

	loginController := controllers.NewLoginController(loginUC)

	return NewAuthRoutes(loginController)
}
