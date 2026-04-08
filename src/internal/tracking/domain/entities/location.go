package entities

type Location struct {
    Lat       float64 `json:"lat"`
    Lng       float64 `json:"lng"`
    Timestamp int64   `json:"timestamp"`
}