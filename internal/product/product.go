package product

import (
	"errors"
)

var names map[string]bool
var nextIdentifier int

func init() {
	names = make(map[string]bool)
	nextIdentifier = 1000
}

func New() product {
	nextIdentifier++
	p := product{
		ID: nextIdentifier,
	}
	save(&p)
	return p
}

func (p *product) SetName(name string) error {
	if names[name] {
		return errors.New("cannot have same name twice")
	}
	names[name] = true
	p.Name = name
	save(p)
	return nil
}

func (p *product) GetName() string {
	return p.Name
}

func (p *product) SetPrice(price int) error {
	p.Price = price
	save(p)
	return nil
}

func (p *product) GetPrice() int {
	return p.Price
}

func (p *product) SetStock(stock int) error {
	p.Stock = stock
	save(p)
	return nil
}

func (p *product) GetStock() int {
	return p.Stock
}

func (p *product) GetID() int {
	return p.ID
}

func GetByID(id int) (product, error) {
	return restore(id)
}
