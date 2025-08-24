package order

import (
	"context"
	domain "github.com/peterprospl12/logi-flow/internal/domain/order"
)

type Repository interface {
	Save(ctx context.Context, o *domain.Order) error
	List(ctx context.Context) ([]*domain.Order, error)
}

type Service interface {
	CreateOrder(ctx context.Context, cmd CreateOrderCommand) (*domain.Order, error)
	GetOrders(ctx context.Context) ([]*domain.Order, error)
}
