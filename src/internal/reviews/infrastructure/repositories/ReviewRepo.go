package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"logired/src/internal/reviews/domain/entities"
)

type ReviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (r *ReviewRepo) CreateReview(review entities.Review) error {
	query := `INSERT INTO reviews (review, rating, iddriver, idpassanger) 
	          VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, review.Review, review.Rating, review.IdDriver, review.IdPassanger)
	if err != nil {
		return fmt.Errorf("error al crear reseña: %w", err)
	}
	log.Println("[ReviewRepo] Reseña creada correctamente")
	return nil
}

func (r *ReviewRepo) GetReviewsByDriverId(idDriver int32) ([]entities.Review, error) {
	query := `SELECT idreview, review, rating, iddriver, idpassanger 
	          FROM reviews WHERE iddriver = ?`
	rows, err := r.db.Query(query, idDriver)
	if err != nil {
		return nil, fmt.Errorf("error al obtener reseñas por conductor: %w", err)
	}
	defer rows.Close()

	var reviews []entities.Review
	for rows.Next() {
		var rv entities.Review
		if err := rows.Scan(&rv.IdReview, &rv.Review, &rv.Rating, &rv.IdDriver, &rv.IdPassanger); err != nil {
			return nil, fmt.Errorf("error al escanear reseña: %w", err)
		}
		reviews = append(reviews, rv)
	}
	return reviews, nil
}

func (r *ReviewRepo) GetReviewsByDriverIdPublic(idDriver int32) ([]entities.Review, error) {
	query := `SELECT idreview, review, rating, iddriver, idpassanger 
	          FROM reviews WHERE iddriver = ?`
	rows, err := r.db.Query(query, idDriver)
	if err != nil {
		return nil, fmt.Errorf("error al obtener reseñas por conductor: %w", err)
	}
	defer rows.Close()

	var reviews []entities.Review
	for rows.Next() {
		var rv entities.Review
		if err := rows.Scan(&rv.IdReview, &rv.Review, &rv.Rating, &rv.IdDriver, &rv.IdPassanger); err != nil {
			return nil, fmt.Errorf("error al escanear reseña: %w", err)
		}
		reviews = append(reviews, rv)
	}
	return reviews, nil
}

func (r *ReviewRepo) GetReviewsByPassangerId(idPassanger int32) ([]entities.Review, error) {
	query := `SELECT idreview, review, rating, iddriver, idpassanger 
	          FROM reviews WHERE idpassanger = ?`
	rows, err := r.db.Query(query, idPassanger)
	if err != nil {
		return nil, fmt.Errorf("error al obtener reseñas por pasajero: %w", err)
	}
	defer rows.Close()

	var reviews []entities.Review
	for rows.Next() {
		var rv entities.Review
		if err := rows.Scan(&rv.IdReview, &rv.Review, &rv.Rating, &rv.IdDriver, &rv.IdPassanger); err != nil {
			return nil, fmt.Errorf("error al escanear reseña: %w", err)
		}
		reviews = append(reviews, rv)
	}
	return reviews, nil
}

func (r *ReviewRepo) UpdateReview(review entities.Review) error {
	query := `UPDATE reviews 
	          SET review = ?, rating = ? 
	          WHERE idreview = ? AND idpassanger = ?`
	result, err := r.db.Exec(query, review.Review, review.Rating, review.IdReview, review.IdPassanger)
	if err != nil {
		return fmt.Errorf("error al actualizar reseña: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("reseña no encontrada o no tienes permiso para editarla")
	}
	log.Println("[ReviewRepo] Reseña actualizada correctamente")
	return nil
}