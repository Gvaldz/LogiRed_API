package entities

type Proposal struct {
	IdProposal int32   `json:"id"`
	Price      float32 `json:"price"`
	Comment    string  `json:"comment"`
	IdDriver   int32   `json:"iduser"`
	IdRide     int32   `json:"id_ride"`
	IdStatus   int32   `json:"idstatus"`
}

func newProposal(idProposal int32, price float32, idDriver int32, idRide int32, idstatus int32) *Proposal {
	return &Proposal{
		IdProposal: idProposal,
		Price:      price,
		IdDriver:   idDriver,
		IdRide:     idRide,
		IdStatus:   idstatus,
	}
}
