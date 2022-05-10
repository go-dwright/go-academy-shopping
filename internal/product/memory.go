package product

import "errors"

// map of products id to a product pointer
var products map[int]*product

func init() {
	products = make(map[int]*product)
}

func save(c *product) error {
	if c.ID == 0 {
		return errors.New("id cannot be 0")
	}
	products[c.ID] = c
	return nil
}

func restore(id int) (product, error) {
	c, exists := products[id]
	if !exists {
		return product{}, errors.New("product with that id does not exist")
	}
	return *c, nil
}
