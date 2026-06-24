package dto

type CreateRideRequest struct {
    OriginCity         string  `json:"origin_city"`
    OriginAddress      string  `json:"origin_address"`
    OriginLat          float64 `json:"origin_lat"`
    OriginLng          float64 `json:"origin_lng"`
    DestinationAddress string  `json:"destination_address"`
    DestinationLat     float64 `json:"destination_lat"`
    DestinationLng     float64 `json:"destination_lng"`
    DistanceKm         float64 `json:"distance_km"`
    Date               string  `json:"date"`
    Hour               string  `json:"hour"`
    AproxWeight        float64 `json:"approx_weight"`
    Description        string  `json:"description"`
}

type RideResponse struct {
    IdRide             int32   `json:"idride"`
    IdClient           int32   `json:"idclient"`
    IdDriver           *int32  `json:"iddriver,omitempty"`
    OriginCity         string  `json:"origincity"`
    OriginAddress      string  `json:"origin_address"`
    OriginLat          float64 `json:"origin_lat"`
    OriginLng          float64 `json:"origin_lng"`
    DestinationAddress string  `json:"destination_address"`
    DestinationLat     float64 `json:"destination_lat"`
    DestinationLng     float64 `json:"destination_lng"`
    DistanceKm         float64 `json:"distance_km"`
    Date               string  `json:"date"`
    Hour               string  `json:"hour"`
    AproxWeight        float64 `json:"aproxweight"`
    Description        string  `json:"description"`
    IdStatus           int32   `json:"idridestatus"`
}

type CreateRideResponse struct {
    Message string       `json:"message"`
    Ride    RideResponse `json:"ride"`
}