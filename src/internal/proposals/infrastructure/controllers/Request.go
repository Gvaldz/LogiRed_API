package controllers

type AcceptProposalRequest struct {
    IdStatus int32 `json:"idstatus" binding:"required" example:"1"`
}

type CreateProposalRequest struct {
    Price   float32 `json:"price" binding:"required" example:"150.5"`
    IdRide  int32   `json:"id_ride" binding:"required" example:"1"`
    Comment string  `json:"comment" example:"Puedo llevar"`
    IdCar   int32   `json:"idcar" binding:"required" example:"3"`
}