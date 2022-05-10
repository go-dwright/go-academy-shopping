package customer

import "errors"

// map of customersMap id to a customer pointer
var customersMap map[int]*customer
var customersList []*customer
var mapList bool

func init() {
	customersMap = make(map[int]*customer)
	customersList = make([]*customer, 0)
	mapList = true
}

func save(c *customer) error {
	if mapList {
		return saveMap(c)
	}
	return saveList(c)
}

func restore(id int) (customer, error) {
	if mapList {
		return restoreMap(id)
	}
	return restoreList(id)
}

func saveMap(c *customer) error {
	if c.ID == 0 {
		return errors.New("id cannot be 0")
	}
	customersMap[c.ID] = c
	return nil
}

func restoreMap(id int) (customer, error) {
	c, exists := customersMap[id]
	if !exists {
		return customer{}, errors.New("customer with that id does not exist")
	}
	return *c, nil
}

func saveList(cust *customer) error {
	if cust.ID == 0 {
		return errors.New("id cannot be 0")
	}
	for i, c := range customersList {
		if c.ID == cust.ID {
			customersList[i] = cust
			return nil
		}
	}
	customersList = append(customersList, cust)
	return nil
}

func restoreList(id int) (customer, error) {
	for _, c := range customersList {
		if c.ID == id {
			return *c, nil
		}
	}
	return customer{}, errors.New("customer with that id does not exist")
}
