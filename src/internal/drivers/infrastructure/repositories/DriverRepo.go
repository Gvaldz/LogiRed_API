package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"logired/src/internal/drivers/domain"
	"logired/src/internal/drivers/domain/entities"
)

type DriverRepo struct {
	db *sql.DB
}

func NewDriverRepo(db *sql.DB) *DriverRepo {
	return &DriverRepo{db: db}
}

func (r *DriverRepo) Create(driver entities.Driver) error {
	query := `INSERT INTO drivers (iduser, rating, image) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, driver.IdUser, driver.Rating, driver.Image)
	if err != nil {
		return fmt.Errorf("error al crear conductor: %w", err)
	}
	log.Println("[DriverRepo] Conductor creado correctamente")
	return nil
}

func (r *DriverRepo) GetByUserID(userID int32) (*domain.DriverDetail, error) {
	query := `
		SELECT d.iduser, d.rating, d.image, u.name, u.lastname, u.email
		FROM drivers d
		INNER JOIN users u ON d.iduser = u.idusers
		WHERE d.iduser = ?
	`
	row := r.db.QueryRow(query, userID)
	var detail domain.DriverDetail
	err := row.Scan(&detail.IdUser, &detail.Rating, &detail.Image, &detail.Name, &detail.Lastname, &detail.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("conductor no encontrado")
		}
		return nil, fmt.Errorf("error al obtener conductor: %w", err)
	}
	return &detail, nil
}

func (r *DriverRepo) GetByID(driverID int32) (*domain.DriverDetail, error) {
	// Es el mismo método que GetByUserID, ya que el ID del conductor es el iduser
	return r.GetByUserID(driverID)
}

func (r *DriverRepo) GetAll() ([]domain.DriverDetail, error) {
	query := `
		SELECT d.iduser, d.rating, d.image, u.name, u.lastname, u.email
		FROM drivers d
		INNER JOIN users u ON d.iduser = u.idusers
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener conductores: %w", err)
	}
	defer rows.Close()

	var drivers []domain.DriverDetail
	for rows.Next() {
		var d domain.DriverDetail
		if err := rows.Scan(&d.IdUser, &d.Rating, &d.Image, &d.Name, &d.Lastname, &d.Email); err != nil {
			return nil, fmt.Errorf("error al escanear conductor: %w", err)
		}
		drivers = append(drivers, d)
	}
	return drivers, nil
}

func (r *DriverRepo) Update(driver entities.Driver) error {
	query := `UPDATE drivers SET rating = ?, image = ? WHERE iduser = ?`
	result, err := r.db.Exec(query, driver.Rating, driver.Image, driver.IdUser)
	if err != nil {
		return fmt.Errorf("error al actualizar conductor: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("conductor no encontrado")
	}
	log.Println("[DriverRepo] Conductor actualizado correctamente")
	return nil
}

func (r *DriverRepo) Delete(userID int32) error {
	query := `DELETE FROM drivers WHERE iduser = ?`
	result, err := r.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error al eliminar conductor: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("conductor no encontrado")
	}
	log.Println("[DriverRepo] Conductor eliminado correctamente")
	return nil
}

func (r *DriverRepo) Exists(userID int32) (bool, error) {
	query := `SELECT COUNT(*) FROM drivers WHERE iduser = ?`
	var count int
	err := r.db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error al verificar existencia: %w", err)
	}
	return count > 0, nil
}