package entities

type Ride struct{
	IdRide 		int32  `json:"id"`
	IdClient 	int32  `json:"id_client"`
	Date 		string `json:"date"`
	Hour 		string `json:"hour"`
	Origin 		string `json:"origin"`
	Destination string `json:"destination"`
	Description string `json:"description"`
}

func newRide(idRide int32, idClient int32, date string, hour string, origin string, destination string, description string) *Ride {
	return &Ride{
		IdRide: idRide,
		IdClient: idClient,
		Date: date,
		Hour: hour,
		Origin: origin,
		Destination: destination,
		Description: description,
	}
}