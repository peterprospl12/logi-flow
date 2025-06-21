package events

type OrderCreated struct {
	ID        string  `json:"id"`
	ShipperID string  `json:"shipper_id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Price     float64 `json:"price"`
	Status    string  `json:"status"`
}
