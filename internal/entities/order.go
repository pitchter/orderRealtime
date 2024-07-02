package entities

type Order struct {
    ID     int       `json:"id"`
    Items  []MenuItem `json:"items"`
    Total  float64   `json:"total"`
    Status string    `json:"status"`
}