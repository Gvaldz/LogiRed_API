package application

import (
	"fmt"
	"logired/src/core"
	"logired/src/internal/users/domain"
	authDomain "logired/src/internal/services/auth/domain"
)

type RessetPassword struct {
    userRepo domain.UserRepository      
    authRepo authDomain.AuthRepository  
    hasher   core.PasswordHasher
}

func NewRessetPassword(
    userRepo domain.UserRepository, 
    authRepo authDomain.AuthRepository, 
    hasher core.PasswordHasher,
) *RessetPassword {
    return &RessetPassword{
        userRepo: userRepo,
        authRepo: authRepo,
        hasher:   hasher,
    }
}

func (uc *RessetPassword) Execute(email, newPassword string) error {
    user, err := uc.authRepo.FindUserByEmail(email)
    if err != nil {
        return fmt.Errorf("usuario no encontrado o credenciales inválidas")
    }

    hashedPassword, err := uc.hasher.Hash(newPassword)
    if err != nil {
        return err
    }

    return uc.userRepo.RessetPassword(user.IdUser, hashedPassword)
}