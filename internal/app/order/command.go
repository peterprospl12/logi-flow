package order

type CreateOrderCommand struct {
	ShipperID    string  `json:"shipper_id" binding:"required"`
	Origin       string  `json:"origin"      binding:"required"`
	Destination  string  `json:"destination" binding:"required"`
	InitialPrice float64 `json:"initial_price" binding:"required,gt=0"`
}
