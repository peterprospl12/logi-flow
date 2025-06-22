package bid

type Bid struct {
	ID      string  `json:"id"`
	OrderID string  `json:"order_id"`
	Price   float64 `json:"price"`
}
