package biz

type OrderStatus string

const (
	OrderStatusOpen      OrderStatus = "open"
	OrderStatusClosed    OrderStatus = "closed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID        string
	ShipperID string
	From      string
	To        string
	Price     float64
	Status    OrderStatus
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id string) (*Order, error)
	List() ([]*Order, error)
	Update(order *Order) error
}

type OrderService struct {
	repo OrderRepository
}

func (s *OrderService) CreateOrder(order *Order) error {
	// logika walidacji, reguły biznesowe
	order.Status = OrderStatusOpen
	return s.repo.Create(order)
}
