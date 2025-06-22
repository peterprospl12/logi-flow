package order

import (
	"github.com/peterprospl12/logi-flow/internal/domain/bid"
)

type Order struct {
	ID            string        `json:"id"`
	ShipperID     string        `json:"shipper_id"`
	Origin        string        `json:"from"`
	Destination   string        `json:"destination"`
	InitialPrice  float64       `json:"price"`
	Status        Status        `json:"status"`
	TransportType TransportType `json:"transport_type"`
	Bids          []bid.Bid     `json:"bids"`
}

func New(id, shipperID, origin, destination string, initialPrice float64, transportType TransportType) (*Order, error) {
	switch {
	case id == "":
		return nil, ErrEmptyID
	case shipperID == "":
		return nil, ErrEmptyShipperID
	case origin == "":
		return nil, ErrEmptyOrigin
	case destination == "":
		return nil, ErrEmptyDestination
	case initialPrice <= 0:
		return nil, ErrInvalidPrice
	}

	return &Order{
		ID:            id,
		ShipperID:     shipperID,
		Origin:        origin,
		Destination:   destination,
		InitialPrice:  initialPrice,
		Status:        Created,
		TransportType: transportType,
		Bids:          []bid.Bid{},
	}, nil
}
