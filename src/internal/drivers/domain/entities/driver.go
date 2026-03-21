package entities

type Driver struct {
	IdUser int32   `json:"id_user"` 
	Rating float32 `json:"rating"`
	Image  string  `json:"image"`
}
func NewDriver(idUser int32, rating float32, image string) *Driver {
	return &Driver{
		IdUser: idUser,
		Rating: rating,
		Image:  image,
	}
}