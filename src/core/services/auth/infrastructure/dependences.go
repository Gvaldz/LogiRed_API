package infrastructure

import (
	"database/sql"
	"logired/src/core"
	"logired/src/core/services/auth/application"
	"logired/src/core/services/auth/infrastructure/controllers"
	users_domain "logired/src/internal/users/domain"
)

type AuthDependencies struct {
	DB              *sql.DB
	Hasher          *core.BcryptHasher
	UserRepo        users_domain.UserRepository
	CredentialsFile string
}

func NewAuthDependencies(
	db *sql.DB,
	hasher *core.BcryptHasher,
	userRepo users_domain.UserRepository,
	credentialsFile string,
) *AuthDependencies {
	return &AuthDependencies{
		DB:              db,
		Hasher:          hasher,
		UserRepo:        userRepo,
		CredentialsFile: credentialsFile,
	}
}

func (d *AuthDependencies) GetRoutes() *AuthRoutes {
	authRepo := NewAuthRepository(d.DB)
	tokenService := core.NewJWTService()

	loginUC := application.NewLogin(authRepo, d.UserRepo, tokenService, d.Hasher)

	verifier := core.NewFirebaseVerifier(d.CredentialsFile)
	googleLoginUC := application.NewGoogleLogin(verifier, authRepo, d.UserRepo, tokenService)

	loginController := controllers.NewLoginController(loginUC)
	googleLoginController := controllers.NewGoogleLoginController(googleLoginUC)

	return NewAuthRoutes(loginController, googleLoginController)
}
