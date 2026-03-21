package infrastructure

import (
	"database/sql"
	"fmt"
	"logired/src/internal/services/auth/domain"
	user "logired/src/internal/users/domain/entities"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(DB *sql.DB) domain.AuthRepository {
	return &AuthRepository{DB: DB}
}

func (r *AuthRepository) FindUserByEmail(email string) (user.User, error) {
	var u user.User
	query := "SELECT iduser, email, password, usertype FROM users WHERE email = ?"

	err := r.DB.QueryRow(query, email).Scan(&u.IdUser, &u.Email, &u.Password, &u.UserType)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *AuthRepository) UpdateLastLogin(userID int32) error {
	query := "UPDATE users SET ultimo_login = NOW() WHERE iduser = ?"
	_, err := r.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error updating last login: %w", err)
	}
	return nil
}

func (r *AuthRepository) FindUserByID(userID int32) (user.User, error) {
	var user user.User
	query := `SELECT iduser, email, password, usertype FROM users WHERE iduser = ?`
	err := r.DB.QueryRow(query, userID).Scan(&user.IdUser, &user.Email, &user.Password, &user.UserType)
	return user, err
}
