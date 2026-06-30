package entities

type UserProfile struct {
    IdUser      int32       `json:"iduser"`
    Name        string      `json:"name"`
    Lastname    string      `json:"lastname"`
    Email       string      `json:"email"`
    NumberPhone string      `json:"number_phone" encrypt:"true"`
    UserType    int         `json:"user_type"`
    ImageURL    string      `json:"image_url"`
    DriverInfo  *DriverInfo `json:"driver_info,omitempty"`
}

type DriverInfo struct {
    Rating   float32  `json:"rating"`
    Car      *CarInfo `json:"car,omitempty"`
}

type CarInfo struct {
    IdCar           int32  `json:"id"`
    CarRegistration string `json:"car_registration"`
    Brand           string `json:"brand"`
    Model           string `json:"model"`
    Color           string `json:"color"`
    MaxCapacity     int32  `json:"max_capacity"`
    FrontViewImage  string `json:"front_view_image"`
    BackViewImage   string `json:"back_view_image"`
    LeftViewImage   string `json:"left_view_image"`
    RightViewImage  string `json:"right_view_image"`
    PlatesImage     string `json:"plates_image"`
    SpacesImage     string `json:"spaces_image"`
}

type MyProfile struct {
    IdUser      int32    `json:"iduser"`
    Name        string   `json:"name"`
    Lastname    string   `json:"lastname"`
    Email       string   `json:"email"`
    NumberPhone string   `json:"number_phone" encrypt:"true"`
    Birthdate   string   `json:"birthdate"`
    UserType    int      `json:"user_type"`
    ImageURL    string   `json:"image_url"`
    TotalTrips  int64    `json:"total_trips"`
    Rating      *float32 `json:"rating,omitempty"`
    Car         *CarInfo `json:"car,omitempty"`
}