package db

import (
	"fmt"

	"github.com/applied-concurrency-in-go/db/internal/ordersmap"
	"github.com/applied-concurrency-in-go/models/order"
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
func (I *OrderDB) Find(id string) (order.Order, error) {
	o, ok := I.placedOrders.Load(id)
	if !ok {
		return order.ZeroOrder, fmt.Errorf("no order found for %s order id", id)
	}

	return o, nil
}

// Upsert creates or updates an order in the orders DB
func (o *OrderDB) Upsert(order order.Order) {
	o.placedOrders.Store(order.ID, order)
}
