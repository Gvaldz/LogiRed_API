package entities

type Ride struct{
	IdRide 				int32  	`json:"id"`
	IdClient 			int32  	`json:"id_client"`
	IdDriver 			*int32 	`json:"id_driver,omitempty"`
	OriginCity  		string	`json:"origin_city"`
	OriginAddress 		string 	`json:"origin_address"`
	OriginLat	 		float64 `json:"origin_lat"`
	OriginLng 			float64 `json:"origin_lng"`
	DestinationAddres 	string 	`json:"destination_address"`
	DestinationLat 		float64 `json:"destination_lat"`
	DestinationLng 		float64 `json:"destination_lng"`
	DistanceKm 			float64 `json:"distance_km,omitempty"`
	Date 				string 	`json:"date"`
	Hour 				string 	`json:"hour"`
	AproxWeight 		float64 `json:"approx_weight"`
	Description 		string 	`json:"description"`
	IdStatus 			int32	`json:"idstatus"`
}

func newRide(
	idRide int32, 
	idClient int32, 
	idDriver *int32, 
	originCity string, 
	originAddress string, 
	originLat float64, 
	originLng float64, 
	destinationAddress string, 
	destinationLat float64, 
	destinationLng float64, 
	distanceKm float64,
	date string, 
	hour string, 
	aproxweight float64, 
	description string, 
	idstatus int32) *Ride {
	return &Ride{
		IdRide: idRide,
		IdClient: idClient,
		IdDriver: idDriver,
		OriginCity: originCity,
		OriginAddress: originAddress,
		OriginLat: originLat,
		OriginLng: originLng,
		DestinationAddres: destinationAddress,
		DestinationLat: destinationLat,
		DestinationLng: destinationLng,
		DistanceKm: distanceKm,
		Date: date,
		Hour: hour,
		AproxWeight: aproxweight,
		Description: description,
		IdStatus: idstatus,
	}
}