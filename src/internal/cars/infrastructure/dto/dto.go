package dto

type CreateCarRequest struct {
    CarRegistration string `form:"car_registration" binding:"required"`
    Brand           string `form:"brand" binding:"required"`
    Model           string `form:"model" binding:"required"`
    Color           string `form:"color" binding:"required"`
    MaxCapacity     int32  `form:"max_capacity" binding:"required"`
}

type CarResponse struct {
    IdCar           int32  `json:"idcar"`
    IdDriver        int32  `json:"iduser"`
    CarRegistration string `json:"car_registration"`
    Brand           string `json:"brand"`
    Model           string `json:"model"`
    Color           string `json:"color"`
    MaxCapacity     int32  `json:"max_capacity"`
    FrontViewImage  string `json:"frontview_image"`
    BackViewImage   string `json:"backview_image"`
    PlatesImage     string `json:"plates_image"`
    SpacesImage     string `json:"space_image"`
}