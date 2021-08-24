package customers

import (
	"fmt"
	"sync"

	"github.com/cgrs/ecommerce-service-starter/utils"
)

type Customer struct {
	Username string
	Password string
	Email    string
	Orders   []Order
	Admin    bool
}

var customers sync.Map

func Login(username, password string) error {
	value, ok := customers.Load(username)
	if !ok {
		return fmt.Errorf("user not found: %s", username)
	}
	customer, ok := value.(*Customer)
	if !ok {
		return fmt.Errorf("unexpected conversion error")
	}
	if password != customer.Password {
		return fmt.Errorf("incorrect password for user %s", username)
	}
	return nil
}

func Register(username, password, email string, admin bool) error {
	_, ok := customers.Load(username)
	if ok {
		return fmt.Errorf("username already used: %s", username)
	}
	customers.Store(username, &Customer{username, password, email, []Order{}, admin})
	return nil
}

func Find(username string) (*Customer, bool) {
	value, ok := customers.Load(username)
	return value.(*Customer), ok
}

func CustomersCount() int {
	return utils.MapCount(customers)
}
