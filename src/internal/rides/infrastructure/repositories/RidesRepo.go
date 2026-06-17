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

func nullInt32ToPtr(n sql.NullInt32) *int32 {
	if !n.Valid {
		return nil
	}
	return &n.Int32
}

func (r *RideRepo) CreateRide(ride entities.Ride) error {
	query := `INSERT INTO rides (
			idclient, 
			origincity, 
			origin_address, 
			origin_lat, 
			origin_lng, 
			destination_address, 
			destination_lat, 
			destination_lng,
			distance_km,
			date, 
			hour, 
			aproxweight, 
			description, 
			idridestatus) 
	        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	_, err := r.db.Exec(query,
		ride.IdClient,
		ride.OriginCity,
		ride.OriginAddress,
		ride.OriginLat,
		ride.OriginLng,
		ride.DestinationAddres,
		ride.DestinationLat,
		ride.DestinationLng,
		ride.DistanceKm,
		ride.Date,
		ride.Hour,
		ride.AproxWeight,
		ride.Description,
		ride.IdStatus)
	if err != nil {
		return fmt.Errorf("error al crear viaje: %w", err)
	}
	log.Println("[RideRepo] Viaje creado correctamente")
	return nil
}

func (r *RideRepo) GetRidesByClientId(idClient int32) ([]entities.Ride, error) {
	query := `SELECT 
			idride, 
			idclient, 
			iddriver,
			origincity, 
			origin_address, 
			origin_lat, 
			origin_lng, 
			destination_address, 
			destination_lat, 
			destination_lng, 
			distance_km,
			date, 
			hour, 
			aproxweight, 
			description, 
			idridestatus
	        FROM rides WHERE idclient = $1`
	rows, err := r.db.Query(query, idClient)
	if err != nil {
		return nil, fmt.Errorf("error al obtener viajes por cliente: %w", err)
	}
	defer rows.Close()

	var rides []entities.Ride
	for rows.Next() {
		var rd entities.Ride
		var iddriver sql.NullInt32
		if err := rows.Scan(
			&rd.IdRide,
			&rd.IdClient,
			&iddriver,
			&rd.OriginCity,
			&rd.OriginAddress,
			&rd.OriginLat,
			&rd.OriginLng,
			&rd.DestinationAddres,
			&rd.DestinationLat,
			&rd.DestinationLng,
			&rd.DistanceKm,
			&rd.Date,
			&rd.Hour,
			&rd.AproxWeight,
			&rd.Description,
			&rd.IdStatus); err != nil {
			return nil, fmt.Errorf("error al escanear viaje: %w", err)
		}
		rd.IdDriver = nullInt32ToPtr(iddriver)
		rides = append(rides, rd)
	}
	return rides, nil
}

func (r *RideRepo) GetRideById(idRide int32) (entities.Ride, error) {
	var ride entities.Ride
	var iddriver sql.NullInt32
	query := `SELECT 
			idride, 
			idclient, 
			iddriver,
			origincity, 
			origin_address, 
			origin_lat, 
			origin_lng, 
			destination_address, 
			destination_lat, 
			destination_lng, 
			distance_km,
			date, 
			hour, 
			aproxweight, 
			description, 
			idridestatus
	        FROM rides WHERE idride = $1`
	err := r.db.QueryRow(query, idRide).Scan(
			&ride.IdRide,
			&ride.IdClient,
			&iddriver,
			&ride.OriginCity,
			&ride.OriginAddress,
			&ride.OriginLat,
			&ride.OriginLng,
			&ride.DestinationAddres,
			&ride.DestinationLat,
			&ride.DestinationLng,
			&ride.DistanceKm,
			&ride.Date,
			&ride.Hour,
			&ride.AproxWeight,
			&ride.Description,
			&ride.IdStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return ride, fmt.Errorf("viaje no encontrado")
		}
		return ride, fmt.Errorf("error al obtener viaje: %w", err)
	}
	ride.IdDriver = nullInt32ToPtr(iddriver)
	return ride, nil
}

func (r *RideRepo) GetRidesByDriverId(idDriver int32) ([]entities.Ride, error) {
	query := `
		SELECT 
			idride, 
			idclient,
			iddriver,
			origincity, 
			origin_address, 
			origin_lat, 
			origin_lng, 
			destination_address, 
			destination_lat, 
			destination_lng, 
			distance_km,
			date, 
			hour, 
			aproxweight,
			description, 
			idridestatus
			FROM rides 
			WHERE iddriver = $1
	`
	rows, err := r.db.Query(query, idDriver)
	if err != nil {
		return nil, fmt.Errorf("error al obtener viajes por conductor: %w", err)
	}
	defer rows.Close()

	var rides []entities.Ride
	for rows.Next() {
		var rd entities.Ride
		var iddriver sql.NullInt32
		if err := rows.Scan(
			&rd.IdRide,
			&rd.IdClient,
			&iddriver,
			&rd.OriginCity,
			&rd.OriginAddress,
			&rd.OriginLat,
			&rd.OriginLng,
			&rd.DestinationAddres,
			&rd.DestinationLat,
			&rd.DestinationLng,
			&rd.DistanceKm,
			&rd.Date,
			&rd.Hour,
			&rd.AproxWeight,
			&rd.Description,
			&rd.IdStatus); err != nil {
			return nil, fmt.Errorf("error al escanear viaje: %w", err)
		}
		rd.IdDriver = nullInt32ToPtr(iddriver)
		rides = append(rides, rd)
	}
	return rides, nil
}

func (r *RideRepo) GetRidesByCity(city string) ([]entities.Ride, error) {
	query := `SELECT 
			idride, 
			idclient, 
			iddriver,
			origincity, 
			origin_address, 
			origin_lat, 
			origin_lng, 
			destination_address, 
			destination_lat, 
			destination_lng, 
			distance_km,
			date, 
			hour, 
			aproxweight,
			description, 
			idridestatus 
            FROM rides 
            WHERE origincity LIKE $1 AND idridestatus = 6 `
	rows, err := r.db.Query(query, "%"+city+"%")
	if err != nil {
		return nil, fmt.Errorf("error al obtener viajes por ciudad: %w", err)
	}
	defer rows.Close()

	var rides []entities.Ride
	for rows.Next() {
		var rd entities.Ride
		var iddriver sql.NullInt32
		if err := rows.Scan(
			&rd.IdRide,
			&rd.IdClient,
			&iddriver,
			&rd.OriginCity,
			&rd.OriginAddress,
			&rd.OriginLat,
			&rd.OriginLng,
			&rd.DestinationAddres,
			&rd.DestinationLat,
			&rd.DestinationLng,
			&rd.DistanceKm,
			&rd.Date,
			&rd.Hour,
			&rd.AproxWeight,
			&rd.Description,
			&rd.IdStatus); err != nil {
			return nil, fmt.Errorf("error al escanear viaje: %w", err)
		}
		rd.IdDriver = nullInt32ToPtr(iddriver)
		rides = append(rides, rd)
	}
	return rides, nil
}

func (r *RideRepo) UpdateRideStatus(idride int32, idstatus int32) error {
	query := `UPDATE rides 
	          SET idridestatus = $1
	          WHERE idride = $2`
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

func (r *RideRepo) AssignDriver(idRide int32, idDriver int32) error {
	query := "UPDATE rides SET iddriver = $1, idridestatus = 1 WHERE idride = $2"
	result, err := r.db.Exec(query, idDriver, idRide)
	if err != nil {
		return fmt.Errorf("error al asignar conductor: %w", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("viaje no encontrado")
	}
	return nil
}

func (r *RideRepo) GetRidesHistory(userID int32, userType int32) ([]entities.Ride, error) {
    var query string
    var args []interface{}

    if userType == 1 {
        query = `
			SELECT
			idride, 
			idclient, 
			iddriver,
			origincity, 
			origin_address, 
			origin_lat, 
			origin_lng, 
			destination_address, 
			destination_lat, 
			destination_lng, 
			distance_km,
			date, 
			hour, 
			aproxweight,
			description, 
			idridestatus 
            FROM rides 
            WHERE idclient = $1 AND idridestatus IN (4, 5)
            ORDER BY date DESC
        `
        args = append(args, userID)
    } else {
        query = `
            SELECT 
			idride, 
			idclient, 
			iddriver,
			origincity, 
			origin_address, 
			origin_lat, 
			origin_lng, 
			destination_address, 
			destination_lat, 
			destination_lng, 
			distance_km,
			date, 
			hour, 
			aproxweight,
			description, 
			idridestatus 
            FROM rides 
            WHERE iddriver = $1 AND idridestatus IN (4, 5)
            ORDER BY date DESC
        `
        args = append(args, userID)
    }

    rows, err := r.db.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("error al obtener historial: %w", err)
    }
    defer rows.Close()

    var rides []entities.Ride
    for rows.Next() {
        var ride entities.Ride
        var idDriver sql.NullInt32

        if err := rows.Scan(
			&ride.IdRide,
			&ride.IdClient,
			&idDriver,
			&ride.OriginCity,
			&ride.OriginAddress,
			&ride.OriginLat,
			&ride.OriginLng,
			&ride.DestinationAddres,
			&ride.DestinationLat,
			&ride.DestinationLng,
			&ride.DistanceKm,
			&ride.Date,
			&ride.Hour,
			&ride.AproxWeight,
			&ride.Description,
            &ride.IdStatus,
        ); err != nil {
            return nil, fmt.Errorf("error al escanear ride: %w", err)
        }

		if idDriver.Valid {
			ride.IdDriver = &idDriver.Int32
		}

        rides = append(rides, ride)
    }

    return rides, nil
}

func (r *RideRepo) GetRidesByStatus(userID int32, userType int32, idstatus int32) ([]entities.Ride, error) {
    var query string
    var args []interface{}

    if userType == 1 {
        query = `
            SELECT 
			idride, 
			idclient, 
			iddriver, 
			origincity, 
           	origin_address, 
		   	origin_lat, 
		   	origin_lng, 
            destination_address, 
			destination_lat, 
			destination_lng, 
            distance_km, 
			date, 
			hour, 
			aproxweight, 
			description, 
			idridestatus 
            FROM rides 
            WHERE idclient = $1 AND idridestatus = $2
            ORDER BY date DESC
        `
        args = append(args, userID, idstatus) 
    } else {
        query = `
            SELECT
			idride, 
			idclient, 
			iddriver, 
			origincity, 
           	origin_address, 
		   	origin_lat, 
		   	origin_lng, 
            destination_address, 
			destination_lat, 
			destination_lng, 
            distance_km, 
			date, 
			hour, 
			aproxweight, 
			description, 
			idridestatus 
            FROM rides 
            WHERE iddriver = $1 AND idridestatus = $2
            ORDER BY date DESC
        `
        args = append(args, userID, idstatus) 
    }

    rows, err := r.db.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("error al obtener historial: %w", err)
    }
    defer rows.Close()

    var rides []entities.Ride
    for rows.Next() {
        var ride entities.Ride
        var idDriver sql.NullInt32

        if err := rows.Scan(
			&ride.IdRide,
			&ride.IdClient,
			&idDriver,
			&ride.OriginCity,
			&ride.OriginAddress,
			&ride.OriginLat,
			&ride.OriginLng,
			&ride.DestinationAddres,
			&ride.DestinationLat,
			&ride.DestinationLng,
			&ride.DistanceKm,
			&ride.Date,
			&ride.Hour,
			&ride.AproxWeight,
			&ride.Description,
            &ride.IdStatus,
        ); err != nil {
            return nil, fmt.Errorf("error al escanear ride: %w", err)
        }

		if idDriver.Valid {
			ride.IdDriver = &idDriver.Int32
		}

        rides = append(rides, ride)
    }

    return rides, nil
}