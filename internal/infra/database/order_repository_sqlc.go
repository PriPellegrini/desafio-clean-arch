package database

import (
	"context"
	"database/sql"

	"github.com/PriPellegrini/desafio-clean-arch/internal/entity"
	"github.com/PriPellegrini/desafio-clean-arch/internal/infra/database/sqlc"
)

type OrderRepositorySQLC struct {
	db *sqlc.Queries
}

func NewOrderRepositorySQLC(db *sql.DB) *OrderRepositorySQLC {
	return &OrderRepositorySQLC{
		db: sqlc.New(db),
	}
}

func (r *OrderRepositorySQLC) Save(order *entity.Order) error {
	ctx := context.Background()
	return r.db.CreateOrder(ctx, sqlc.CreateOrderParams{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	})
}

func (r *OrderRepositorySQLC) GetTotal() (int, error) {
	ctx := context.Background()
	orders, err := r.db.ListOrders(ctx)
	if err != nil {
		return 0, err
	}
	return len(orders), nil
}
