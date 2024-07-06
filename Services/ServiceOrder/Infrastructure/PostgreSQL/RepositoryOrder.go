package postgresql

import (
	domain "grpc_project/Services/ServiceOrder/Domain"

	"github.com/lib/pq"
)

type OrderRepository struct {
	database *DatabasePostgre
}

func NewRepositoryOrder(db *DatabasePostgre) domain.OrderRepository {
	return &OrderRepository{database: db}
}

func (orderr *OrderRepository) SaveOrder(or *domain.Order) error {
	_, err := orderr.database.Conn.Exec("INSERT INTO userorder(userid,bookid,amount) VALUES ($1,$2,$3)", or.UserID, pq.Array(or.BookID), pq.Array(or.AmountRequired))
	if err != nil {
		return err
	}
	return nil
}
