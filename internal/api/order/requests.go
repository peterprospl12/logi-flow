package order

import appOrder "github.com/peterprospl12/logi-flow/internal/app/order"

type CreateOrderRequest struct {
	ShipperID    string  `json:"shipper_id" binding:"required"`
	Origin       string  `json:"origin"      binding:"required"`
	Destination  string  `json:"destination" binding:"required"`
	InitialPrice float64 `json:"initial_price" binding:"required,gt=0"`
}

func (r CreateOrderRequest) ToCommand() appOrder.CreateOrderCommand {
	return appOrder.CreateOrderCommand{
		ShipperID:    r.ShipperID,
		Origin:       r.Origin,
		Destination:  r.Destination,
		InitialPrice: r.InitialPrice,
	}
}
