package application

import (
	"context"
	"fmt"
	"logired/src/core"
	auth "logired/src/core/services/auth/domain"
	userRepo "logired/src/internal/users/domain"
)

type GoogleLogin struct {
	verifier     *core.FirebaseVerifier
	authRepo     auth.AuthRepository
	userRepo     userRepo.UserRepository
	tokenService auth.TokenService
}

func NewGoogleLogin(
	verifier *core.FirebaseVerifier,
	authRepo auth.AuthRepository,
	userRepo userRepo.UserRepository,
	tokenService auth.TokenService,
) *GoogleLogin {
	return &GoogleLogin{
		verifier:     verifier,
		authRepo:     authRepo,
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (uc *GoogleLogin) Execute(ctx context.Context, idToken string) (auth.Token, error) {
	firebaseToken, err := uc.verifier.VerifyIDToken(ctx, idToken)
	if err != nil {
		return auth.Token{}, err
	}

	email, ok := firebaseToken.Claims["email"].(string)
	if !ok || email == "" {
		return auth.Token{}, fmt.Errorf("no se pudo obtener el email del token de Google")
	}

	user, err := uc.authRepo.FindUserByEmail(email)
	if err != nil {
		return auth.Token{}, fmt.Errorf("no existe una cuenta registrada con este correo")
	}

	approved := false
	if user.UserType == 2 {
		approved, err = uc.authRepo.FindDriverApproved(user.IdUser)
		if err != nil {
			return auth.Token{}, fmt.Errorf("error al obtener datos del conductor")
		}
	}

	token, err := uc.tokenService.GenerateToken(user.IdUser, user.Email, user.UserType, approved)
	if err != nil {
		return auth.Token{}, fmt.Errorf("fallo al generar token")
	}

	go func() {
		_ = uc.authRepo.UpdateLastLogin(user.IdUser)
	}()

	return token, nil
}
