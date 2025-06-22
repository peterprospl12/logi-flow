package order

type CreateOrderDTO struct {
	ID        string  `json:"id"`
	ShipperID string  `json:"shipper_id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Price     float64 `json:"price"`
}

type AcceptBidDTO struct {
	OrderID string
	BidID   string
}
