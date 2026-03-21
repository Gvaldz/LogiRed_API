package application

import (
	"logired/src/internal/reviews/domain"
	"logired/src/internal/reviews/domain/entities"
)

type UpdateReview struct {
	repo domain.IReview
}

func NewUpdateReview(repo domain.IReview) *UpdateReview {
	return &UpdateReview{repo: repo}
}

func (ur *UpdateReview) Execute(review entities.Review) error {
	return ur.repo.UpdateReview(review)
}