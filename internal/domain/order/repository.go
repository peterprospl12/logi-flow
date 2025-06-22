package order

type Repository interface {
	GetByID(id string) (*Order, error)
	Save(o *Order) error
}
