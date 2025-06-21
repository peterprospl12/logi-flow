package commands

type CreateOrder struct {
	ShipperID string  `json:"shipper_id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Price     float64 `json:"price"`
}
