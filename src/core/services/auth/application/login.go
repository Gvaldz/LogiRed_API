package application

import (
	"errors"
	"logired/src/core"
	auth "logired/src/core/services/auth/domain"
	user_repo "logired/src/internal/users/domain"
	user "logired/src/internal/users/domain/entities"
)

type Login struct {
	authRepo     auth.AuthRepository
	userRepo     user_repo.UserRepository
	tokenService auth.TokenService
	hasher       core.PasswordHasher
}

func NewLogin(
	authRepo auth.AuthRepository,
	userRepo user_repo.UserRepository,
	tokenService auth.TokenService,
	hasher core.PasswordHasher,
) *Login {
	return &Login{
		authRepo:     authRepo,
		userRepo:     userRepo,
		tokenService: tokenService,
		hasher:       hasher,
	}
}

func (uc *Login) Execute(credentials user.User) (auth.Token, error) {
	user, err := uc.authRepo.FindUserByEmail(credentials.Email)
	if err != nil {
		return auth.Token{}, errors.New("datos incorrectos")
	}

	if err := uc.hasher.Compare(user.Password, credentials.Password); err != nil {
		return auth.Token{}, errors.New("datos incorrectos")
	}

	citywork := ""

	if user.UserType == 2 {
		citywork, err = uc.authRepo.FindDriverCityWorkByUserID(user.IdUser)
		if err != nil {
			return auth.Token{}, errors.New("error al obtener datos del conductor")
		}
	}

	token, err := uc.tokenService.GenerateToken(user.IdUser, user.Email, user.UserType, citywork)
	if err != nil {
		return auth.Token{}, errors.New("fallo en generar token")
	}

	go func() {
		_ = uc.authRepo.UpdateLastLogin(user.IdUser)
	}()

	return token, nil
}
