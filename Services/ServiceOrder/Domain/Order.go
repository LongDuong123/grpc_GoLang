package domain

type Order struct {
	UserID         int
	BookID         []int64
	AmountRequired []int64
	Total          int64
}

type OrderRepository interface {
	SaveOrder(*Order) error
}

type OrderInteractor interface {
	CreateOrder(*Order) (*Order, error)
}
