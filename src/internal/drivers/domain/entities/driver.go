package entities

type Driver struct {
	IdUser 		int32   `json:"id_user"` 
	Rating 		float32 `json:"rating"`
	Citywork 	string  `json:"citywork"`
}
func NewDriver(idUser int32, rating float32, citywork string) *Driver {
	return &Driver{
		IdUser: idUser,
		Rating: rating,
		Citywork: citywork,
	}
}