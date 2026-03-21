package application

import (
	"logired/src/internal/reviews/domain"
	"logired/src/internal/reviews/domain/entities"
)

type GetReviewsByDriverIdPublic struct {
	repo domain.IReview
}

func NewGetReviewsByDriverIdPublic(repo domain.IReview) *GetReviewsByDriverIdPublic {
	return &GetReviewsByDriverIdPublic{repo: repo}
}

func (g *GetReviewsByDriverIdPublic) Execute(idDriver int32) ([]entities.Review, error) {
	return g.repo.GetReviewsByDriverIdPublic(idDriver)
}