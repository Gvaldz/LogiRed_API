package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"logired/src/internal/cars/domain/entities"
	"strings"
)

type CarRepo struct {
	db *sql.DB
}

func NewCarRepo(db *sql.DB) *CarRepo {
	return &CarRepo{db: db}
}

func (r *CarRepo) CreateCar(car entities.Car) error {
	query := "INSERT INTO cars (iduser, car_registration, brand, model, color, max_capacity, frontview_image, backview_image, plates_image, space_image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, car.IdDriver, car.CarRegistration, car.Brand, car.Model, car.Color, car.MaxCapacity, car.FrontViewImage, car.BackViewImage, car.PlatesImage, car.SpacesImage)
	if err != nil {
		return fmt.Errorf("error al crear car: %w", err)
	}
	return nil
}

func (r *CarRepo) UpdateCar(car entities.Car) error {
	setClauses := []string{}
	args := []interface{}{}

	if car.CarRegistration != "" { setClauses = append(setClauses, "car_registration = ?"); args = append(args, car.CarRegistration) }
	if car.Brand != ""           { setClauses = append(setClauses, "brand = ?");            args = append(args, car.Brand) }
	if car.Model != ""           { setClauses = append(setClauses, "model = ?");            args = append(args, car.Model) }
	if car.Color != ""           { setClauses = append(setClauses, "color = ?");            args = append(args, car.Color) }
	if car.MaxCapacity != 0      { setClauses = append(setClauses, "max_capacity = ?");     args = append(args, car.MaxCapacity) }
	if car.FrontViewImage != ""  { setClauses = append(setClauses, "frontview_image = ?");  args = append(args, car.FrontViewImage) }
	if car.BackViewImage != ""   { setClauses = append(setClauses, "backview_image = ?");   args = append(args, car.BackViewImage) }
	if car.PlatesImage != ""     { setClauses = append(setClauses, "plates_image = ?");     args = append(args, car.PlatesImage) }
	if car.SpacesImage != ""     { setClauses = append(setClauses, "space_image = ?");      args = append(args, car.SpacesImage) }

	if len(setClauses) == 0 {
		return fmt.Errorf("no hay campos para actualizar")
	}

	args = append(args, car.IdCar, car.IdDriver)
	query := fmt.Sprintf(
		"UPDATE cars SET %s WHERE idcar = ? AND iduser = ?",
		strings.Join(setClauses, ", "),
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
	query := "SELECT idcar, iduser, car_registration, brand, model, color, max_capacity, frontview_image, backview_image, plates_image, space_image FROM cars WHERE iduser = ?"
	rows, err := r.db.Query(query, idDriver)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cars: %w", err)
	}
	defer rows.Close()

	var cars []entities.Car
	for rows.Next() {
		var car entities.Car
		if err := rows.Scan(&car.IdCar, &car.IdDriver, &car.CarRegistration, &car.Brand, &car.Model, &car.Color, &car.MaxCapacity, &car.FrontViewImage, &car.BackViewImage, &car.PlatesImage, &car.SpacesImage); err != nil {
			return nil, fmt.Errorf("error al escanear car: %w", err)
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (r *CarRepo) GetCarById(idCar int32, idDriver int32) (entities.Car, error) {
	var car entities.Car
	query := "SELECT idcar, iduser, car_registration, brand, model, color, max_capacity, frontview_image, backview_image, plates_image, space_image FROM cars WHERE idcar = ? AND iduser = ?"
	err := r.db.QueryRow(query, idCar, idDriver).Scan(&car.IdCar, &car.IdDriver, &car.CarRegistration, &car.Brand, &car.Model, &car.Color, &car.MaxCapacity, &car.FrontViewImage, &car.BackViewImage, &car.PlatesImage, &car.SpacesImage)
	if err != nil {
		return car, fmt.Errorf("car no encontrado o acceso denegado: %w", err)
	}
	return car, nil
}

func (r *CarRepo) DeleteCar(idCar int32, idDriver int32) error {
	query := "DELETE FROM cars WHERE idcar = ? AND iduser = ?"
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
