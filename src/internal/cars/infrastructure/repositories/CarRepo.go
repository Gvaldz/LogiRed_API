package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"logired/src/internal/cars/domain/entities"
	"strings"
	"errors"
)

type CarRepo struct {
	db *sql.DB
}

func NewCarRepo(db *sql.DB) *CarRepo {
	return &CarRepo{db: db}
}

func (r *CarRepo) CreateCar(car entities.Car) error {
	query := "INSERT INTO cars (iduser, car_registration, brand, model, color, max_capacity, frontview_image, backview_image, leftview_image, rightview_image, plates_image, space_image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	_, err := r.db.Exec(query, car.IdDriver, car.CarRegistration, car.Brand, car.Model, car.Color, car.MaxCapacity, car.FrontViewImage, car.BackViewImage, car.LeftViewImage, car.RightViewImage, car.PlatesImage, car.SpacesImage)
	if err != nil {
		return fmt.Errorf("error al crear car: %w", err)
	}
	return nil
}

func (r *CarRepo) UpdateCar(car entities.Car) error {
	setClauses := []string{}
	args := []interface{}{}
	argCounter := 1

	if car.CarRegistration != "" {
		setClauses = append(setClauses, fmt.Sprintf("car_registration = $%d", argCounter))
		args = append(args, car.CarRegistration)
		argCounter++
	}
	if car.Brand != "" {
		setClauses = append(setClauses, fmt.Sprintf("brand = $%d", argCounter))
		args = append(args, car.Brand)
		argCounter++
	}
	if car.Model != "" {
		setClauses = append(setClauses, fmt.Sprintf("model = $%d", argCounter))
		args = append(args, car.Model)
		argCounter++
	}
	if car.Color != "" {
		setClauses = append(setClauses, fmt.Sprintf("color = $%d", argCounter))
		args = append(args, car.Color)
		argCounter++
	}
	if car.MaxCapacity != 0 {
		setClauses = append(setClauses, fmt.Sprintf("max_capacity = $%d", argCounter))
		args = append(args, car.MaxCapacity)
		argCounter++
	}
	if car.FrontViewImage != "" {
		setClauses = append(setClauses, fmt.Sprintf("frontview_image = $%d", argCounter))
		args = append(args, car.FrontViewImage)
		argCounter++
	}
	if car.BackViewImage != "" {
		setClauses = append(setClauses, fmt.Sprintf("backview_image = $%d", argCounter))
		args = append(args, car.BackViewImage)
		argCounter++
	}
	if car.LeftViewImage != "" {
		setClauses = append(setClauses, fmt.Sprintf("leftview_image = $%d", argCounter))
		args = append(args, car.LeftViewImage)
		argCounter++
	}
	if car.RightViewImage != "" {
		setClauses = append(setClauses, fmt.Sprintf("rightview_image = $%d", argCounter))
		args = append(args, car.RightViewImage)
		argCounter++
	}
	if car.PlatesImage != "" {
		setClauses = append(setClauses, fmt.Sprintf("plates_image = $%d", argCounter))
		args = append(args, car.PlatesImage)
		argCounter++
	}
	if car.SpacesImage != "" {
		setClauses = append(setClauses, fmt.Sprintf("space_image = $%d", argCounter))
		args = append(args, car.SpacesImage)
		argCounter++
	}

	if len(setClauses) == 0 {
		return fmt.Errorf("no hay campos para actualizar")
	}

	args = append(args, car.IdCar, car.IdDriver)
	query := fmt.Sprintf(
		"UPDATE cars SET %s WHERE idcar = $%d AND iduser = $%d",
		strings.Join(setClauses, ", "),
		argCounter, argCounter+1,
	)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al actualizar car: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("car no encontrado o no tienes permiso para editarlo")
	}
	return nil
}

func (r *CarRepo) GetCarsByDriverId(idDriver int32) ([]entities.Car, error) {
	query := "SELECT idcar, iduser, car_registration, brand, model, color, max_capacity, frontview_image, backview_image, leftview_image, rightview_image, plates_image, space_image FROM cars WHERE iduser = $1"
	rows, err := r.db.Query(query, idDriver)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cars: %w", err)
	}
	defer rows.Close()

	var cars []entities.Car
	for rows.Next() {
		var car entities.Car
		if err := rows.Scan(&car.IdCar, &car.IdDriver, &car.CarRegistration, &car.Brand, &car.Model, &car.Color, &car.MaxCapacity, &car.FrontViewImage, &car.BackViewImage, &car.LeftViewImage, &car.RightViewImage, &car.PlatesImage, &car.SpacesImage); err != nil {
			return nil, fmt.Errorf("error al escanear car: %w", err)
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (r *CarRepo) GetCarById(idCar int32) (entities.Car, error) {
	var car entities.Car
	query := "SELECT idcar, iduser, car_registration, brand, model, color, max_capacity, frontview_image, backview_image, leftview_image, rightview_image, plates_image, space_image FROM cars WHERE idcar = $1"
	err := r.db.QueryRow(query, idCar).Scan(&car.IdCar, &car.IdDriver, &car.CarRegistration, &car.Brand, &car.Model, &car.Color, &car.MaxCapacity, &car.FrontViewImage, &car.BackViewImage, &car.LeftViewImage, &car.RightViewImage, &car.PlatesImage, &car.SpacesImage)
	if err != nil {
		return car, fmt.Errorf("car no encontrado o acceso denegado: %w", err)
	}
	return car, nil
}

func (r *CarRepo) DeleteCar(idCar int32, idDriver int32) error {
	query := "DELETE FROM cars WHERE idcar = $1 AND iduser = $2"
	result, err := r.db.Exec(query, idCar, idDriver)
	if err != nil {
		return fmt.Errorf("error al eliminar car: %w", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("car no encontrado o no tienes permiso para eliminarlo")
	}
	log.Println("[CarRepo] - car eliminado correctamente")
	return nil
}


func (repo *CarRepo) CreateCarTx(tx interface{}, car entities.Car) error {
	sqlTx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("error crítico: la transacción provista no es del tipo *sql.Tx")
	}

	query := `
		INSERT INTO cars (
			iduser, car_registration, brand, model, color, 
			max_capacity, frontview_image, backview_image, leftview_image, rightview_image, plates_image, space_image
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := sqlTx.Exec(query,
		car.IdDriver,
		car.CarRegistration,
		car.Brand,
		car.Model,
		car.Color,
		car.MaxCapacity,
		car.FrontViewImage,
		car.BackViewImage,
		car.LeftViewImage,
		car.RightViewImage,
		car.PlatesImage,
		car.SpacesImage,
	)

	if err != nil {
		return fmt.Errorf("error al guardar el vehículo en la transacción: %w", err)
	}

	return nil
}