package db

import (
	"fmt"

	"github.com/applied-concurrency-in-go/db/internal/ordersmap"
	"github.com/applied-concurrency-in-go/models"
)

type OrderDB struct {
	placedOrders ordersmap.Map
}

// NewOrdersDB creates a new empty order service
func NewOrdersDB() *OrderDB {
	return &OrderDB{
		placedOrders: ordersmap.Map{},
	}
}

// Find order for a given id, if one exists
func (o *OrderDB) Find(id string) (models.Order, error) {
	order, ok := o.placedOrders.Load(id)
	if !ok {
		return models.ZeroOrder, fmt.Errorf("no order found for %s order id", id)
	}

	return order, nil
}

// Upsert creates or updates an order in the orders DB
func (o *OrderDB) Upsert(order models.Order) {
	o.placedOrders.Store(order.ID, order)
}
