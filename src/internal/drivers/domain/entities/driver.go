package entities

type Driver struct {
	IdUser 		int32   `json:"id_user"` 
	Rating 		float32 `json:"rating"`
	Approved 	int    `json:"approved"`
}
func NewDriver(idUser int32, rating float32, approved int) *Driver {
	return &Driver{
		IdUser: idUser,
		Rating: rating,
		Approved: approved,
	}
}