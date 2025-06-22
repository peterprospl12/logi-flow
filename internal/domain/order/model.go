package order

import (
	"github.com/google/uuid"
	"github.com/peterprospl12/logi-flow/internal/domain/bid"
)

type Order struct {
	ID           uuid.UUID `json:"id"`
	ShipperID    string    `json:"shipper_id"`
	Origin       string    `json:"from"`
	Destination  string    `json:"destination"`
	InitialPrice float64   `json:"price"`
	Status       Status    `json:"status"`
	//TransportType TransportType `json:"transport_type"`
	Bids []bid.Bid `json:"bids"`
}

func New(shipperID, origin string, destination string, initialPrice float64) (*Order, error) {
	switch {
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
		ID:           uuid.New(),
		ShipperID:    shipperID,
		Origin:       origin,
		Destination:  destination,
		InitialPrice: initialPrice,
		Status:       Created,
		//TransportType: transportType,
		Bids: []bid.Bid{},
	}, nil
}
