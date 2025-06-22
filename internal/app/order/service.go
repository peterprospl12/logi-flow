package order

type Service interface {
	CreateOrder(cmd CreateOrderDTO) error
}
