package domain

type Order struct {
	UserID         int     `json:"Id"`
	BookID         []int64 `json:"Bookid"`
	AmountRequired []int64 `json:"AmountRequired"`
	Total          int64   `json:"Total"`
}

type OrderRepository interface {
	SaveOrder(*Order) error
}

type OrderInteractor interface {
	CreateOrder(*Order) (*Order, error)
}
