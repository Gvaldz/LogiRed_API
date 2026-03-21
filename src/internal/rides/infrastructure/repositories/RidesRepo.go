package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"logired/src/internal/rides/domain/entities"
)

type RideRepo struct {
	db *sql.DB
}

func NewRideRepo(db *sql.DB) *RideRepo {
	return &RideRepo{db: db}
}

func (r *RideRepo) CreateRide(ride entities.Ride) error {
	query := `INSERT INTO rides (idclient, date, hour, origin, destination, description) 
	          VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, ride.IdClient, ride.Date, ride.Hour, ride.Origin, ride.Destination, ride.Description)
	if err != nil {
		return fmt.Errorf("error al crear viaje: %w", err)
	}
	log.Println("[RideRepo] Viaje creado correctamente")
	return nil
}

func (r *RideRepo) CancelRide(idRide int32, idClient int32) error {
	query := `DELETE FROM rides WHERE idride = ? AND idclient = ?`
	result, err := r.db.Exec(query, idRide, idClient)
	if err != nil {
		return fmt.Errorf("error al cancelar viaje: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("viaje no encontrado o no tienes permiso para cancelarlo")
	}
	log.Println("[RideRepo] Viaje cancelado correctamente")
	return nil
}

func (r *RideRepo) GetRidesByClientId(idClient int32) ([]entities.Ride, error) {
	query := `SELECT idride, idclient, date, hour, origin, destination, description 
	          FROM rides WHERE idclient = ?`
	rows, err := r.db.Query(query, idClient)
	if err != nil {
		return nil, fmt.Errorf("error al obtener viajes por cliente: %w", err)
	}
	defer rows.Close()

	var rides []entities.Ride
	for rows.Next() {
		var rd entities.Ride
		if err := rows.Scan(&rd.IdRide, &rd.IdClient, &rd.Date, &rd.Hour, &rd.Origin, &rd.Destination, &rd.Description); err != nil {
			return nil, fmt.Errorf("error al escanear viaje: %w", err)
		}
		rides = append(rides, rd)
	}
	return rides, nil
}

func (r *RideRepo) GetRideById(idRide int32) (entities.Ride, error) {
	var ride entities.Ride
	query := `SELECT idride, idclient, date, hour, origin, destination, description 
	          FROM rides WHERE idride = ?`
	err := r.db.QueryRow(query, idRide).Scan(&ride.IdRide, &ride.IdClient, &ride.Date, &ride.Hour, &ride.Origin, &ride.Destination, &ride.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return ride, fmt.Errorf("viaje no encontrado")
		}
		return ride, fmt.Errorf("error al obtener viaje: %w", err)
	}
	return ride, nil
}

func (r *RideRepo) GetRidesByDriverId(idDriver int32) ([]entities.Ride, error) {
	query := `
		SELECT r.idride, r.idclient, r.date, r.hour, r.origin, r.destination, r.description
		FROM rides r
		WHERE p.iddriver = ? AND p.accepted = true
	`
	rows, err := r.db.Query(query, idDriver)
	if err != nil {
		return nil, fmt.Errorf("error al obtener viajes por conductor: %w", err)
	}
	defer rows.Close()

	var rides []entities.Ride
	for rows.Next() {
		var rd entities.Ride
		if err := rows.Scan(&rd.IdRide, &rd.IdClient, &rd.Date, &rd.Hour, &rd.Origin, &rd.Destination, &rd.Description); err != nil {
			return nil, fmt.Errorf("error al escanear viaje: %w", err)
		}
		rides = append(rides, rd)
	}
	return rides, nil
}