package repo

import (
	"fmt"
	"math"

	"github.com/applied-concurrency-in-go/db"
	"github.com/applied-concurrency-in-go/models/order"
	"github.com/applied-concurrency-in-go/models/orderstatus"
	"github.com/applied-concurrency-in-go/models/product"
)

// repo holds all the dependencies required for repo operations
type repo struct {
	products *db.ProductDB
	orders   *db.OrderDB
}

// Repo is the interface we expose to outside packages
type Repo interface {
	CreateOrder(item order.Item) (*order.Order, error)
	GetAllProducts() []product.Product
	GetOrder(id order.Id) (order.Order, error)
}

// New creates a new Order repo with the correct database dependencies
func New() (Repo, error) {
	p, err := db.NewProductsDB()
	if err != nil {
		return nil, err
	}
	o := repo{
		products: p,
		orders:   db.NewOrdersDB(),
	}
	return &o, nil
}

// GetAllProducts returns all products in the system
func (r *repo) GetAllProducts() []product.Product {
	return r.products.FindAll()
}

// GetProduct returns the given order if one exists
func (r *repo) GetOrder(id order.Id) (order.Order, error) {
	return r.orders.Find(id)
}

// CreateOrder creates a new order for the given item
func (r *repo) CreateOrder(item order.Item) (*order.Order, error) {
	if err := r.validateItem(item); err != nil {
		return nil, err
	}
	order := order.NewOrder(item)
	r.orders.Upsert(order)
	r.processOrders(&order)
	return &order, nil
}

// validateItem runs validations on a given order
func (r *repo) validateItem(item order.Item) error {
	if item.Amount < 1 {
		return fmt.Errorf("order amount must be at least 1:got %d", item.Amount)
	}
	if r.products.Exists(item.ProductID) {
		return fmt.Errorf("product %s does not exist", item.ProductID)
	}
	return nil
}

func (r *repo) processOrders(order *order.Order) {
	r.processOrder(order)
	r.orders.Upsert(*order)
	fmt.Printf("Processing order %s completed\n", order.ID)
}

// processOrder is an internal method which completes or rejects an order
func (r *repo) processOrder(order *order.Order) {
	item := order.Item
	product, err := r.products.Find(item.ProductID)
	if err != nil {
		order.Status = orderstatus.Rejected
		order.Error = err.Error()
		return
	}
	if product.Stock < item.Amount {
		order.Status = orderstatus.Rejected
		order.Error = fmt.Sprintf("not enough stock for product %s:got %d, want %d", item.ProductID, product.Stock, item.Amount)
		return
	}
	remainingStock := product.Stock - item.Amount
	product.Stock = remainingStock
	r.products.Upsert(product)

	total := math.Round(float64(order.Item.Amount)*product.Price*100) / 100
	order.Total = total
	order.Complete()
}
