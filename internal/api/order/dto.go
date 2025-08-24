package order

import (
	"github.com/google/uuid"
	//domainBid "github.com/peterprospl12/logi-flow/internal/domain/bid"
	domainOrder "github.com/peterprospl12/logi-flow/internal/domain/order"
)

type CreateOrderDTO struct {
	ShipperID    string  `json:"shipper_id"`
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	InitialPrice float64 `json:"initial_price"`
}

type GetOrderDTO struct {
	ID           string  `json:"id"`
	ShipperID    string  `json:"shipper_id"`
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	InitialPrice float64 `json:"initial_price"`
	Status       string  `json:"status"`
	//Bids         []BidDto `json:"bids"`
	//BidsCount    int      `json:"bids_count"`
}

type BidDto struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
}

func ToGetOrderDTO(o *domainOrder.Order) GetOrderDTO {
	return GetOrderDTO{
		ID:           uuidToString(o.ID),
		ShipperID:    o.ShipperID,
		Origin:       o.Origin,
		Destination:  o.Destination,
		InitialPrice: o.InitialPrice,
		Status:       string(o.Status),
		//Bids:         mapBids(o.Bids),
		//BidsCount:    len(o.Bids),
	}
}

//func mapBids(bids []domainBid.Bid) []BidDTO {
//	out := make([]BidDTO, 0, len(bids))
//	for _, b := range bids {
//		out = append(out, BidDTO{
//			ID:        uuidToString(b.ID),
//			UserID:    b.UserID,
//			Amount:    b.Amount,
//			CreatedAt: b.CreatedAt.Format(time.RFC3339),
//		})
//	}
//	return out
//}

func uuidToString(id uuid.UUID) string {
	if id == uuid.Nil {
		return ""
	}
	return id.String()
}
