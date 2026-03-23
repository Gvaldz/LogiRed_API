package entities

type Ride struct{
	IdRide 		int32  	`json:"id"`
	IdClient 	int32  	`json:"id_client"`
	Origin 		string 	`json:"origin"`
	Destination string 	`json:"destination"`
	Date 		string 	`json:"date"`
	Hour 		string 	`json:"hour"`
	AproxWeight float64 `json:"aprow_weight"`
	Description string 	`json:"description"`
	IdStatus 	int32	`json:"idstatus"`
}

func newRide(idRide int32, idClient int32, origin string, destination string, date string, hour string, aproxweight float64, description string, idstatus int32) *Ride {
	return &Ride{
		IdRide: idRide,
		IdClient: idClient,
		Origin: origin,
		Destination: destination,
		Date: date,
		Hour: hour,
		AproxWeight: aproxweight,
		Description: description,
		IdStatus: idstatus,
	}
}