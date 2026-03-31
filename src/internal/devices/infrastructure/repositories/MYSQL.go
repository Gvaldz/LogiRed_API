package repositories

import (
	"database/sql"
	"fmt"
)

type DeviceRepo struct {
	db *sql.DB
}

func NewDeviceRepo(db *sql.DB) *DeviceRepo {
	return &DeviceRepo{db: db}
}

func (r *DeviceRepo) SaveToken(userID int32, token, deviceName string) error {
	query := `
		INSERT INTO user_devices (iduser, fcm_token, device_name)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
			iduser      = VALUES(iduser),
			device_name = VALUES(device_name),
			updated_at  = CURRENT_TIMESTAMP
	`
	_, err := r.db.Exec(query, userID, token, deviceName)
	if err != nil {
		return fmt.Errorf("error al guardar token: %w", err)
	}
	return nil
}

func (r *DeviceRepo) GetTokensByUser(userID int32) ([]string, error) {
	query := "SELECT fcm_token FROM user_devices WHERE iduser = ?"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener tokens: %w", err)
	}
	defer rows.Close()

	var tokens []string
	for rows.Next() {
		var token string
		if err := rows.Scan(&token); err != nil {
			return nil, fmt.Errorf("error al escanear token: %w", err)
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (r *DeviceRepo) DeleteToken(token string) error {
	_, err := r.db.Exec("DELETE FROM user_devices WHERE fcm_token = ?", token)
	if err != nil {
		return fmt.Errorf("error al eliminar token: %w", err)
	}
	return nil
}