package domain

import "logired/src/internal/reviews/domain/entities"

type IReview interface {
	CreateReview(review entities.Review) error
	GetReviewsByDriverId(idDriver int32) ([]entities.Review, error)
	GetReviewsByPassangerId(idPassanger int32) ([]entities.Review, error)
	GetReviewsByDriverIdPublic(idDriver int32)([]entities.Review, error)
	UpdateReview(review entities.Review) error
}