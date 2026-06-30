package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"logired/src/core/security"
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

	encryptedU, err := security.PrepareForWrite(u)
	if err != nil {
		return user.User{}, fmt.Errorf("error al cifrar usuario: %w", err)
	}

	query := "INSERT INTO users (name, lastname, birthdate, numberphone, email, password, usertype, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING iduser"

	var id int32
	err = r.DB.QueryRow(query, encryptedU.Name, encryptedU.Lastname, encryptedU.Birthdate, encryptedU.NumberPhone, encryptedU.Email, encryptedU.Password, encryptedU.UserType, encryptedU.ImageURL).Scan(&id)
	if err != nil {
		return user.User{}, fmt.Errorf("error al crear usuario: %w", err)
	}

	return user.User{
		IdUser:     id,
		Name:       u.Name,
		Lastname:   u.Lastname,
		Email:      u.Email,
		UserType:   u.UserType,
		NumberPhone: u.NumberPhone,
		Birthdate:  u.Birthdate,
		ImageURL:   u.ImageURL,
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
		decryptedU, err := security.MaterializeFromRead(u)
		if err != nil {
			return nil, fmt.Errorf("error al descifrar user: %w", err)
		}
		usersList = append(usersList, decryptedU)
	}

	return usersList, nil
}

func (r *UserRepository) GetUserByID(iduser int32) (user.User, error) {
	if r.DB == nil {
		return user.User{}, fmt.Errorf("database connection is nil")
	}

	var u user.User
	query := "SELECT iduser, name, lastname, birthdate, numberphone, email, usertype, image_url FROM users WHERE iduser = $1"

	err := r.DB.QueryRow(query, iduser).Scan(&u.IdUser, &u.Name, &u.Lastname, &u.Birthdate, &u.NumberPhone, &u.Email, &u.UserType, &u.ImageURL)

	if err != nil {
		return u, fmt.Errorf("error al obtener usuario: %w", err)
	}
	
	decryptedU, err := security.MaterializeFromRead(u)
	if err != nil {
		return u, fmt.Errorf("error al descifrar usuario: %w", err)
	}
	return decryptedU, nil
}

func (r *UserRepository) GetUserByEmail(email string) (user.User, error) {
	var u user.User
	query := "SELECT iduser, name, lastname, email, password, usertype FROM users WHERE email = $1"

	err := r.DB.QueryRow(query, email).Scan(
		&u.IdUser,
		&u.Name,
		&u.Lastname,
		&u.Email,      
		&u.Password,
		&u.UserType,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("usuario no encontrado")
		}
		return u, fmt.Errorf("error al obtener usuario por email: %w", err)
	}
	
	decryptedU, err := security.MaterializeFromRead(u)
	if err != nil {
		return u, fmt.Errorf("error al descifrar usuario: %w", err)
	}
	return decryptedU, nil
}

func (r *UserRepository) UpdateUser(id int32, u user.User) error {
    encryptedU, err := security.PrepareForWrite(u)
    if err != nil {
        return fmt.Errorf("error al cifrar usuario para update: %w", err)
    }
    u = encryptedU

    setClauses := []string{}
    args := []interface{}{}
    argCounter := 1

    if u.Name != "" {
        setClauses = append(setClauses, fmt.Sprintf("name = $%d", argCounter))
        args = append(args, u.Name)
        argCounter++
    }
    if u.Lastname != "" {
        setClauses = append(setClauses, fmt.Sprintf("lastname = $%d", argCounter))
        args = append(args, u.Lastname)
        argCounter++
    }
    if u.Email != "" {
        setClauses = append(setClauses, fmt.Sprintf("email = $%d", argCounter))
        args = append(args, u.Email)
        argCounter++
    }
    if u.NumberPhone != "" {
        setClauses = append(setClauses, fmt.Sprintf("numberphone = $%d", argCounter))
        args = append(args, u.NumberPhone)
        argCounter++
    }
    if u.Birthdate != "" {
        setClauses = append(setClauses, fmt.Sprintf("birthdate = $%d", argCounter))
        args = append(args, u.Birthdate)
        argCounter++
    }
    if u.ImageURL != "" {
        setClauses = append(setClauses, fmt.Sprintf("image_url = $%d", argCounter))
        args = append(args, u.ImageURL)
        argCounter++
    }

    if len(setClauses) == 0 {
        return errors.New("no hay campos para actualizar")
    }

    var exists int
    err = r.DB.QueryRow("SELECT COUNT(*) FROM users WHERE iduser = $1", id).Scan(&exists)
    if err != nil {
        return fmt.Errorf("error al verificar usuario: %w", err)
    }
    if exists == 0 {
        return fmt.Errorf("usuario no encontrado")
    }

    args = append(args, id)
    query := fmt.Sprintf("UPDATE users SET %s WHERE iduser = $%d", strings.Join(setClauses, ", "), argCounter)

    result, err := r.DB.Exec(query, args...)
    if err != nil {
        return fmt.Errorf("error al actualizar usuario: %w", err)
    }

    rows, _ := result.RowsAffected()
    fmt.Printf(">>> RowsAffected: %d\n", rows)

    return nil
}

func (r *UserRepository) RessetPassword(email string, newHashedPassword string) error {
    query := "UPDATE users SET password = $1 WHERE email = $2"
    
    result, err := r.DB.Exec(query, newHashedPassword, email)
    if err != nil {
        return fmt.Errorf("error al actualizar contraseña por email: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error al verificar filas afectadas: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no se encontró ningún usuario con el email: %s", email)
    }

    return nil
}

func (r *UserRepository) UpdatePassword(id int32, newHashedPassword string) error {
	query := "UPDATE users SET password = $1 WHERE iduser = $2"
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
	query := "DELETE FROM users WHERE iduser = $1"
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
	encryptedU, err := security.PrepareForWrite(u)
	if err != nil {
		return user.User{}, fmt.Errorf("error al cifrar usuario tx: %w", err)
	}

	query := "INSERT INTO users (name, lastname, birthdate, numberphone, email, password, usertype, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING iduser"

	var id int32
	err = tx.QueryRow(query, encryptedU.Name, encryptedU.Lastname, encryptedU.Birthdate, encryptedU.NumberPhone, encryptedU.Email, encryptedU.Password, encryptedU.UserType, encryptedU.ImageURL).Scan(&id)
	if err != nil {
		return user.User{}, fmt.Errorf("error al crear usuario: %w", err)
	}

	return user.User{
		IdUser:      id,
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

func (r *UserRepository) GetUserProfileByID(id int32) (user.UserProfile, error) {
    query := `
        SELECT
            u.iduser, u.name, u.lastname, u.email, u.numberphone,
            u.usertype, u.image_url,
            d.rating,
            c.idcar, c.car_registration, c.brand, c.model, c.color, c.max_capacity,
            c.frontview_image, c.backview_image, c.leftview_image, c.rightview_image,
            c.plates_image, c.space_image
        FROM users u
        LEFT JOIN drivers d ON u.iduser = d.iduser
        LEFT JOIN cars c ON u.iduser = c.iduser
        WHERE u.iduser = $1
    `

    var profile user.UserProfile
    var rating          sql.NullFloat64
    var idCar           sql.NullInt32
    var carRegistration sql.NullString
    var brand           sql.NullString
    var model           sql.NullString
    var color           sql.NullString
    var maxCapacity     sql.NullInt32
    var frontView       sql.NullString
    var backView        sql.NullString
    var leftView        sql.NullString
    var rightView       sql.NullString
    var platesImg       sql.NullString
    var spacesImg       sql.NullString

    err := r.DB.QueryRow(query, id).Scan(
        &profile.IdUser,
        &profile.Name,
        &profile.Lastname,
        &profile.Email,
        &profile.NumberPhone,
        &profile.UserType,
        &profile.ImageURL,
        &rating,
        &idCar,
        &carRegistration,
        &brand,
        &model,
        &color,
        &maxCapacity,
        &frontView,
        &backView,
        &leftView,
        &rightView,
        &platesImg,
        &spacesImg,
    )
    if err != nil {
        return user.UserProfile{}, fmt.Errorf("usuario no encontrado: %w", err)
    }

    // 1. Ahora validamos con 'rating.Valid' para saber si existen datos de conductor
    if rating.Valid {
        driverInfo := &user.DriverInfo{
            // 2. Eliminamos el campo Citywork de aquí
            Rating: float32(rating.Float64),
        }
        
        if idCar.Valid {
            driverInfo.Car = &user.CarInfo{
                IdCar:           idCar.Int32,
                CarRegistration: carRegistration.String,
                Brand:           brand.String,
                Model:           model.String,
                Color:           color.String,
                MaxCapacity:     maxCapacity.Int32,
                FrontViewImage:  frontView.String,
                BackViewImage:   backView.String,
                LeftViewImage:   leftView.String,
                RightViewImage:  rightView.String,
                PlatesImage:     platesImg.String,
                SpacesImage:     spacesImg.String,
            }
        }
        profile.DriverInfo = driverInfo
    }

	decryptedProfile, err := security.MaterializeFromRead(profile)
	if err != nil {
		return user.UserProfile{}, fmt.Errorf("error al descifrar perfil: %w", err)
	}
    return decryptedProfile, nil
}

func (r *UserRepository) GetMyProfileByID(id int32) (user.MyProfile, error) {
    query := `
        SELECT
            u.iduser, u.name, u.lastname, u.email, u.numberphone,
            u.birthdate, u.usertype, u.image_url,
            d.rating,
            (SELECT COUNT(*) FROM rides WHERE idclient = u.iduser) AS client_trips,
            (SELECT COUNT(*) FROM rides WHERE iddriver = u.iduser) AS driver_trips,
            c.idcar, c.car_registration, c.brand, c.model, c.color, c.max_capacity,
            c.frontview_image, c.backview_image, c.leftview_image, c.rightview_image,
            c.plates_image, c.space_image
        FROM users u
        LEFT JOIN drivers d ON u.iduser = d.iduser
        LEFT JOIN cars c ON u.iduser = c.iduser
        WHERE u.iduser = $1
    `

    var profile user.MyProfile
    var rating          sql.NullFloat64
    var clientTrips     int64
    var driverTrips     int64
    var idCar           sql.NullInt32
    var carRegistration sql.NullString
    var brand           sql.NullString
    var model           sql.NullString
    var color           sql.NullString
    var maxCapacity     sql.NullInt32
    var frontView       sql.NullString
    var backView        sql.NullString
    var leftView        sql.NullString
    var rightView       sql.NullString
    var platesImg       sql.NullString
    var spacesImg       sql.NullString

    err := r.DB.QueryRow(query, id).Scan(
        &profile.IdUser,
        &profile.Name,
        &profile.Lastname,
        &profile.Email,
        &profile.NumberPhone,
        &profile.Birthdate,
        &profile.UserType,
        &profile.ImageURL,
        &rating,
        &clientTrips,
        &driverTrips,
        &idCar,
        &carRegistration,
        &brand,
        &model,
        &color,
        &maxCapacity,
        &frontView,
        &backView,
        &leftView,
        &rightView,
        &platesImg,
        &spacesImg,
    )
    if err != nil {
        return user.MyProfile{}, fmt.Errorf("usuario no encontrado: %w", err)
    }

    if profile.UserType == 2 {
        profile.TotalTrips = driverTrips
        if rating.Valid {
            r32 := float32(rating.Float64)
            profile.Rating = &r32
        }
        if idCar.Valid {
            profile.Car = &user.CarInfo{
                IdCar:           idCar.Int32,
                CarRegistration: carRegistration.String,
                Brand:           brand.String,
                Model:           model.String,
                Color:           color.String,
                MaxCapacity:     maxCapacity.Int32,
                FrontViewImage:  frontView.String,
                BackViewImage:   backView.String,
                LeftViewImage:   leftView.String,
                RightViewImage:  rightView.String,
                PlatesImage:     platesImg.String,
                SpacesImage:     spacesImg.String,
            }
        }
    } else {
        profile.TotalTrips = clientTrips
    }

	decryptedProfile, err := security.MaterializeFromRead(profile)
	if err != nil {
		return user.MyProfile{}, fmt.Errorf("error al descifrar mi perfil: %w", err)
	}
    return decryptedProfile, nil
}