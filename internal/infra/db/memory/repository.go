package memory

import (
	"context"
	"errors"
	"github.com/google/uuid"
	appOrder "github.com/peterprospl12/logi-flow/internal/app/order"
	"sync"

	"github.com/peterprospl12/logi-flow/internal/domain/order"
)

type Store struct {
	mu     sync.RWMutex
	orders map[string]*order.Order
}

func NewStore() *Store {
	return &Store{
		orders: make(map[string]*order.Order),
	}
}

type orderRepo struct {
	st *Store
}

func NewOrderRepository(st *Store) appOrder.Repository {
	return &orderRepo{st: st}
}

func (r *orderRepo) Save(ctx context.Context, o *order.Order) error {
	if o == nil {
		return errors.New("nil order")
	}

	if o.ID == uuid.Nil {
		return errors.New("empty id")
	}

	r.st.mu.Lock()
	defer r.st.mu.Unlock()
	r.st.orders[o.ID.String()] = o
	return nil
}

func (r *orderRepo) List(ctx context.Context) ([]*order.Order, error) {
	r.st.mu.RLock()
	defer r.st.mu.RUnlock()

	orders := make([]*order.Order, 0, len(r.st.orders))
	for _, o := range r.st.orders {
		orders = append(orders, o)
	}
	return orders, nil
}
