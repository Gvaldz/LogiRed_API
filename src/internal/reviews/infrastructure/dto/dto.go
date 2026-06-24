package dto

type CreateReviewRequest struct {
    Review   string  `json:"review" binding:"required"`
    Rating   float32 `json:"rating" binding:"required,min=1,max=5"`
    IdDriver int32   `json:"iduser" binding:"required"`
}

type ReviewResponse struct {
    IdReview    int32   `json:"idreview"`
    Review      string  `json:"review"`
    Rating      float32 `json:"rating"`
    IdDriver    int32   `json:"iddriver"`
    IdPassanger int32   `json:"idpassanger"`
}

type UpdateReviewRequest struct {
    Review string  `json:"review"`
    Rating float32 `json:"rating"`
}