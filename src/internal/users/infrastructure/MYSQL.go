package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	users "logired/src/internal/users/domain"
	user "logired/src/internal/users/domain/entities"
	"strings"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) users.UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) CreateUser(u user.User) (user.User, error) {

	query := "INSERT INTO users (name, lastname, birthdate, numberphone, email, password, usertype, image_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.DB.Exec(query, u.Name, u.Lastname, u.Birthdate, u.NumberPhone,u.Email, u.Password, u.UserType, u.ImageURL)
	if err != nil {
		return user.User{}, fmt.Errorf("error al crear usuario: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user.User{}, fmt.Errorf("error al obtener ID: %w", err)
	}

	return user.User{
		IdUser:   	int32(id),
		Name:     		u.Name,
		Lastname: 		u.Lastname,
		Email:    		u.Email,
		UserType: 		u.UserType,
		NumberPhone: 	u.NumberPhone,
		Birthdate: 		u.Birthdate,
		ImageURL: 		u.ImageURL,
	}, nil
}

func (r *UserRepository) GetAllUsers() ([]user.User, error) {
	query := "SELECT iduser, name, lastname, email, usertype FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %w", err)
	}
	defer rows.Close()

	var usersList []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.IdUser, &u.Name, &u.Lastname, &u.Email, &u.UserType); err != nil {
			return nil, fmt.Errorf("error al escanear user: %w", err)
		}
		usersList = append(usersList, u)
	}

	return usersList, nil
}

func (r *UserRepository) GetUserByID(iduser int32) (user.User, error) {
	if r.DB == nil {
		return user.User{}, fmt.Errorf("database connection is nil")
	}

	var u user.User
	query := "SELECT iduser, name, lastname, birthdate, numberphone, email, usertype, image_url FROM users WHERE iduser = ?"

	err := r.DB.QueryRow(query, iduser).Scan(&u.IdUser, &u.Name, &u.Lastname, &u.Birthdate, &u.NumberPhone, &u.Email, &u.UserType, &u.ImageURL)

	if err != nil {
		return u, fmt.Errorf("error al obtener usuario: %w", err)
	}
	return u, nil
}

func (r *UserRepository) GetUserByEmail(email string) (user.User, error) {
	var u user.User
	query := "SELECT iduser, name, lastname, email, password, usertype FROM users WHERE email = ?"

	err := r.DB.QueryRow(query, email).Scan(
		&u.IdUser,
		&u.Name,
		&u.Lastname,
		&u.Password,
		&u.Password,
		&u.UserType,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("usuario no encontrado")
		}
		return u, fmt.Errorf("error al obtener usuario por email: %w", err)
	}
	return u, nil
}

func (r *UserRepository) UpdateUser(id int32, user user.User) error {
    setClauses := []string{}
    args := []interface{}{}

    if user.Name != ""        { setClauses = append(setClauses, "name = ?");         args = append(args, user.Name) }
    if user.Lastname != ""    { setClauses = append(setClauses, "lastname = ?");      args = append(args, user.Lastname) }
    if user.Email != ""       { setClauses = append(setClauses, "email = ?");         args = append(args, user.Email) }
    if user.NumberPhone != "" { setClauses = append(setClauses, "number_phone = ?");  args = append(args, user.NumberPhone) }
    if user.Birthdate != ""   { setClauses = append(setClauses, "birthdate = ?");     args = append(args, user.Birthdate) }
    if user.ImageURL != ""    { setClauses = append(setClauses, "image_url = ?");     args = append(args, user.ImageURL) }

    if len(setClauses) == 0 {
        return errors.New("no hay campos para actualizar")
    }

    args = append(args, id)
    query := fmt.Sprintf("UPDATE users SET %s WHERE iduser = ?", strings.Join(setClauses, ", "))
    _, err := r.DB.Exec(query, args...)
    return err
}

func (r *UserRepository) UpdatePassword(id int32, newHashedPassword string) error {
	query := "UPDATE users SET password = ? WHERE iduser = ?"
	result, err := r.DB.Exec(query, newHashedPassword, id)
	if err != nil {
		return fmt.Errorf("error al actualizar contraseña: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar actualización de contraseña: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	return nil
}

func (r *UserRepository) DeleteUser(id int32) error {
	query := "DELETE FROM users WHERE iduser = ?"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar eliminación: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	return nil
}

func (r *UserRepository) CreateUserTx(tx *sql.Tx, u user.User) (user.User, error) {
	query := "INSERT INTO users (name, lastname, birthdate, numberphone, email, password, usertype, image_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.Exec(query, u.Name, u.Lastname, u.Birthdate, u.NumberPhone, u.Email, u.Password, u.UserType, u.ImageURL)
	if err != nil {
		return user.User{}, fmt.Errorf("error al crear usuario: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user.User{}, fmt.Errorf("error al obtener ID: %w", err)
	}

	return user.User{
		IdUser:      int32(id),
		Name:        u.Name,
		Lastname:    u.Lastname,
		Email:       u.Email,
		UserType:    u.UserType,
		NumberPhone: u.NumberPhone,
		Birthdate:   u.Birthdate,
		ImageURL:    u.ImageURL,
	}, nil
}

func (r *UserRepository) BeginTx() (*sql.Tx, error) {
	return r.DB.Begin()
}