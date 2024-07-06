package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DatabasePostgre struct {
	Conn *sql.DB
}

func ConnectPostgreSQL() (*DatabasePostgre, error) {
	pgDB, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=admin123 dbname=order sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &DatabasePostgre{Conn: pgDB}, nil
}
