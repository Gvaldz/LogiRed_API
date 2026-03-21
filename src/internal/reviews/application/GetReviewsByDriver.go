package application

import (
	"logired/src/internal/reviews/domain"
	"logired/src/internal/reviews/domain/entities"
)

type GetReviewsByDriver struct {
	repo domain.IReview
}

func NewGetReviewsByDriver(repo domain.IReview) *GetReviewsByDriver {
	return &GetReviewsByDriver{repo: repo}
}

func (g *GetReviewsByDriver) Execute(idDriver int32) ([]entities.Review, error) {
	return g.repo.GetReviewsByDriverId(idDriver)
}