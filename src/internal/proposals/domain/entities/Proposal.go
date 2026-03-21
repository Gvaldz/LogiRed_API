package entities

type Proposal struct {
	IdProposal int32   `json:"id"`
	Price      float32 `json:"price"`
	IdDriver   int32   `json:"iduser"`
	IdRide     int32   `json:"id_ride"`
	Accepted   bool    `json:"accepted"`
}

func newProposal(idProposal int32, price float32, idDriver int32, idRide int32, accepted bool) *Proposal {
	return &Proposal{
		IdProposal: idProposal,
		Price:      price,
		IdDriver:   idDriver,
		IdRide:     idRide,
		Accepted:   accepted,
	}
}
