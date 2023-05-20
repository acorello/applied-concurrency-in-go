package db

import (
	"fmt"
	"sort"

	"github.com/applied-concurrency-in-go/db/internal/productsmap"
	"github.com/applied-concurrency-in-go/fixtures"
	"github.com/applied-concurrency-in-go/models"
)

type ProductDB struct {
	products productsmap.Map
}

// NewProductsDB creates a new empty product DB
func NewProductsDB() (*ProductDB, error) {
	db := &ProductDB{}
	// load start position
	err := fixtures.ImportProducts(func(productKey string, product models.Product) {
		db.products.Store(productKey, product)
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// TODO: What happens on the calling site if I replace a generic type (eg. `string`) with a tiny-type (et. `type ProductId string`)?

// Checks whether a product with a given id exists
func (p *ProductDB) Exists(id string) bool {
	_, ok := p.products.Load(id)
	return ok
}

// Find returns a given product if one exists
func (p *ProductDB) Find(id string) (models.Product, error) {
	prod, ok := p.products.Load(id)
	if !ok {
		return models.Product{}, fmt.Errorf("no product found for id %s", id)
	}

	return prod, nil
}

// Upsert creates or updates a product in the orders DB
func (p *ProductDB) Upsert(prod models.Product) {
	p.products.Store(prod.ID, prod)
}

// FindAll returns all products in the system
func (p *ProductDB) FindAll() []models.Product {
	var res []models.Product
	p.products.Range(func(_ string, product models.Product) bool {
		res = append(res, product)
		return true
	})
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	return res
}
