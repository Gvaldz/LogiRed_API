package entities

type Car struct {
	IdCar           	int32  `json:"id"`
	IdDriver        	int32  `json:"iduser"`
	CarRegistration 	string `json:"car_registration"`
	Brand           	string `json:"brand"`
	Model           	string `json:"model"`
	Color           	string `json:"color"`
	MaxCapacity     	int32  `json:"max_capacity"`
	FrontViewImage  	string `json:"front_view_image"`
	BackViewImage   	string `json:"back_view_image"`
	LeftViewImage   	string `json:"left_view_image"`
	RightViewImage  	string `json:"right_view_image"`
	PlatesImage      	string `json:"plates_image"`
	SpacesImage      	string `json:"spaces_image"`
}

func newCar(
	idCar int32,
	idDriver int32,
	carRegistration string,
	brand string,
	model string,
	color string,
	maxCapacity int32,
	frontviewImage string,
	backviewImage string,
	leftviewImage string,
	rightviewImage string,
	platesImage string,
	spacesImage string,
	) *Car {
	return &Car{
		IdCar:           idCar,
		IdDriver:        idDriver,
		CarRegistration: carRegistration,
		Brand:           brand,
		Model:           model,
		Color:           color,
		MaxCapacity:     maxCapacity,
		FrontViewImage:  frontviewImage,
		BackViewImage:   backviewImage,
		LeftViewImage:   leftviewImage,
		RightViewImage:  rightviewImage,
		PlatesImage:     platesImage,
		SpacesImage:     spacesImage,
	}
}
