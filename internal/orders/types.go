package orders

import (
	"context"
	repo "github.com/mohammedkhalf/Ecommerce-API/internal/adapters/postgresql/sqlc"
)

type OrderItem struct {
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customer_id"`
	Items      []OrderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}
