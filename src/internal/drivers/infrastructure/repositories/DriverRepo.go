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

func (r *DriverRepo) CreateTx(tx *sql.Tx, driver entities.Driver) error {
	query := "INSERT INTO drivers (iduser, rating, citywork) VALUES (?, ?, ?)"
	_, err := tx.Exec(query, driver.IdUser, driver.Rating, driver.Citywork)
	if err != nil {
		return fmt.Errorf("error al crear conductor: %w", err)
	}
	return nil
}

func (r *DriverRepo) UpdateCitywork(idUser int32, citywork string) error {
    query := "UPDATE drivers SET citywork = ? WHERE iduser = ?"
    _, err := r.db.Exec(query, citywork, idUser)
    return err
}

func (r *DriverRepo) GetDriversByCity(city string) ([]domain.DriverDetail, error) {
    query := `
        SELECT d.iduser, d.rating, u.name, u.lastname, u.email
        FROM drivers d
        INNER JOIN users u ON d.iduser = u.iduser
        WHERE d.citywork LIKE ?
    `
    rows, err := r.db.Query(query, "%"+city+"%")
    if err != nil {
        return nil, fmt.Errorf("error al obtener conductores: %w", err)
    }
    defer rows.Close()

    var drivers []domain.DriverDetail
    for rows.Next() {
        var d domain.DriverDetail
        if err := rows.Scan(&d.IdUser, &d.Rating, &d.Name, &d.Lastname, &d.Email); err != nil {
            return nil, fmt.Errorf("error al escanear conductor: %w", err)
        }
        drivers = append(drivers, d)
    }
    return drivers, nil
}