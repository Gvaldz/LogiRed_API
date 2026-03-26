package entities

type UserProfile struct {
    IdUser      int32  `json:"iduser"`
    Name        string `json:"name"`
    Lastname    string `json:"lastname"`
    Email       string `json:"email"`
    NumberPhone string `json:"number_phone"`
    Birthdate   string `json:"birthdate"`
    UserType    int    `json:"user_type"`
    ImageURL    string `json:"image_url"`

	DriverInfo  *DriverInfo `json:"driver_info,omitempty"`
}

type DriverInfo struct {
    Citywork string  `json:"citywork"`
    Rating   float32 `json:"rating"`
}