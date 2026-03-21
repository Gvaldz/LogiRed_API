package application

import (
	"logired/src/internal/reviews/domain"
	"logired/src/internal/reviews/domain/entities"
)

type CreateReview struct {
	repo domain.IReview
}

func NewCreateReview(repo domain.IReview) *CreateReview {
	return &CreateReview{repo: repo}
}

func (cr *CreateReview) Execute(review entities.Review) error {
	return cr.repo.CreateReview(review)
}