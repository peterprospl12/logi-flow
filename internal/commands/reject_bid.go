package commands

type RejectBid struct {
	BidID   string `json:"bid_id"`
	OrderID string `json:"order_id"`
	Reason  string `json:"reason"`
}
