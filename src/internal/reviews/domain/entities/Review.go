package entities

type Review struct {
	IdReview    int32   `json:"id"`
	Review      string  `json:"review"`
	Rating      float32 `json:"rating"`
	IdDriver    int32   `json:"iduser"`
	IdPassanger int32   `json:"id_passanger"`
}

func newReview(idReview int32, review string, rating float32, idDriver int32, idPassanger int32) *Review {
	return &Review{
		IdReview:    idReview,
		Review:      review,
		Rating:      rating,
		IdDriver:    idDriver,
		IdPassanger: idPassanger,
	}
}
