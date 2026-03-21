package application

import (
	"logired/src/internal/reviews/domain"
	"logired/src/internal/reviews/domain/entities"
)

type GetReviewsByPassanger struct {
	repo domain.IReview
}

func NewGetReviewsByPassanger(repo domain.IReview) *GetReviewsByPassanger {
	return &GetReviewsByPassanger{repo: repo}
}

func (g *GetReviewsByPassanger) Execute(idPassanger int32) ([]entities.Review, error) {
	return g.repo.GetReviewsByPassangerId(idPassanger)
}