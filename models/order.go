package models

import (
	"time"

	"github.com/applied-concurrency-in-go/models/orderstatus"
	"github.com/google/uuid"
)

type Order struct {
	ID        string                  `json:"id,omitempty"`
	Item      Item                    `json:"item"`
	Total     float64                 `json:"total,omitempty"`
	Status    orderstatus.OrderStatus `json:"status,omitempty"`
	Error     string                  `json:"error,omitempty"`
	CreatedAt string                  `json:"createdAt,omitempty"`
}

var ZeroOrder = Order{}

type Item struct {
	ProductID string `json:"productId"`
	Amount    int    `json:"amount"`
}

func NewOrder(item Item) Order {
	const timeFormat = "2006-01-02 15:04:05.000"
	return Order{
		ID:        uuid.New().String(),
		Status:    orderstatus.New,
		CreatedAt: time.Now().Format(timeFormat),
		Item:      item,
	}
}

func (o *Order) Complete() {
	o.Status = orderstatus.Completed
}
