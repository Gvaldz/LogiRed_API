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
	query := `INSERT INTO rides (idclient, origin, origincity, destination, date, hour, aproxweight, description, idridestatus) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, ride.IdClient, ride.Origin, ride.OriginCity, ride.Destination, ride.Date, ride.Hour, ride.AproxWeight, ride.Description, ride.IdStatus)
	if err != nil {
		return fmt.Errorf("error al crear viaje: %w", err)
	}
	log.Println("[RideRepo] Viaje creado correctamente")
	return nil
}

func (r *RideRepo) GetRidesByClientId(idClient int32) ([]entities.Ride, error) {
	query := `SELECT idride, idclient, date, hour, origin, destination, description, aproxweight, idridestatus 
	          FROM rides WHERE idclient = ?`
	rows, err := r.db.Query(query, idClient)
	if err != nil {
		return nil, fmt.Errorf("error al obtener viajes por cliente: %w", err)
	}
	defer rows.Close()

	var rides []entities.Ride
	for rows.Next() {
		var rd entities.Ride
		if err := rows.Scan(&rd.IdRide, &rd.IdClient, &rd.Date, &rd.Hour, &rd.Origin, &rd.Destination, &rd.Description, &rd.AproxWeight, &rd.IdStatus); err != nil {
			return nil, fmt.Errorf("error al escanear viaje: %w", err)
		}
		rides = append(rides, rd)
	}
	return rides, nil
}

func (r *RideRepo) GetRideById(idRide int32) (entities.Ride, error) {
	var ride entities.Ride
	query := `SELECT idride, idclient, date, hour, origin, destination, description, aproxweight, idridestatus 
	          FROM rides WHERE idride = ?`
	err := r.db.QueryRow(query, idRide).Scan(&ride.IdRide, &ride.IdClient, &ride.Date, &ride.Hour, &ride.Origin, &ride.Destination, &ride.Description, &ride.AproxWeight, &ride.IdStatus)
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
		SELECT r.idride, r.idclient, r.date, r.hour, r.origin, r.destination, r.description, r.aproxweight, r.idridestatus
		FROM rides r
		WHERE p.iddriver = ? AND p.idproposalstatus = 1
	`
	rows, err := r.db.Query(query, idDriver)
	if err != nil {
		return nil, fmt.Errorf("error al obtener viajes por conductor: %w", err)
	}
	defer rows.Close()

	var rides []entities.Ride
	for rows.Next() {
		var rd entities.Ride
		if err := rows.Scan(&rd.IdRide, &rd.IdClient, &rd.Date, &rd.Hour, &rd.Origin, &rd.Destination, &rd.Description, &rd.AproxWeight, &rd.IdStatus); err != nil {
			return nil, fmt.Errorf("error al escanear viaje: %w", err)
		}
		rides = append(rides, rd)
	}
	return rides, nil
}

func (r *RideRepo) GetRidesByCity(city string) ([]entities.Ride, error) {
    query := `SELECT idride, idclient, date, hour, origin, destination, description, aproxweight, idridestatus 
              FROM rides 
              WHERE origincity LIKE ? AND idridestatus = 6`
    rows, err := r.db.Query(query, "%"+city+"%")
    if err != nil {
        return nil, fmt.Errorf("error al obtener viajes por ciudad: %w", err)
    }
    defer rows.Close()

    var rides []entities.Ride
    for rows.Next() {
        var r entities.Ride
        if err := rows.Scan(&r.IdRide, &r.IdClient, &r.Date, &r.Hour, &r.Origin, &r.Destination, &r.Description, &r.AproxWeight, &r.IdStatus); err != nil {
            return nil, fmt.Errorf("error al escanear viaje: %w", err)
        }
        rides = append(rides, r)
    }
    return rides, nil
}


func (r *RideRepo) UpdateRideStatus(idride int32, idstatus int32) error {
	query := `UPDATE rides 
	          SET idridestatus = ?
	          WHERE idride = ?`
	result, err := r.db.Exec(query, idstatus, idride)
	if err != nil {
		return fmt.Errorf("error al actualizar estado del viaje: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("viaje no encontrado o no tienes permiso para editarla")
	}
	log.Println("[RideRepo] Viaje actualizado correctamente")
	return nil
}