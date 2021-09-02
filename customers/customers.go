package customers

import (
	"fmt"
	"sync"
	"time"

	"github.com/cgrs/ecommerce-service-starter/cart"
	"github.com/cgrs/ecommerce-service-starter/orders"
	"github.com/cgrs/ecommerce-service-starter/utils"
)

type Customer struct {
	Username string
	Password string
	Email    string
	Orders   []string
	Admin    bool
}

var customers sync.Map
var updateLock sync.Mutex

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
	customers.Store(username, &Customer{username, password, email, []string{}, admin})
	return nil
}

func Find(username string) *Customer {
	value, ok := customers.Load(username)
	if !ok {
		return nil
	}
	return value.(*Customer)
}

func Count() int {
	return utils.MapCount(customers)
}

func FindBySession(sessionId string) *Customer {
	c := cart.Find(sessionId)
	if c == nil {
		return nil
	}
	cust := Find(c.Customer.Username)
	return cust
}

func DoCheckout(cartId string) (*orders.Order, error) {
	c := cart.Find(cartId)
	if c == nil {
		return nil, fmt.Errorf("user does not have a cart")
	}
	o := &orders.Order{
		Username: c.Customer.Username,
		Number:   int(time.Now().Unix()),
		Lines:    c.Lines,
		Status:   "completed",
		Total:    c.GetTotal(),
	}
	orders.Store(o)

	customer := Find(c.Customer.Username)
	customer.Orders = append(customer.Orders, fmt.Sprint(o.Number))
	if err := cart.Empty(cartId); err != nil {
		return nil, err
	}
	return o, Update(customer)
}

func Update(c *Customer) error {
	updateLock.Lock()
	defer updateLock.Unlock()
	value, ok := customers.LoadAndDelete(c.Username)
	if !ok {
		return fmt.Errorf("customer not found for update: %s", c.Username)
	}
	customer, ok := value.(*Customer)
	if !ok {
		return fmt.Errorf("unexpected conversion error")
	}
	customer.Username = c.Username
	customer.Admin = c.Admin
	customer.Email = c.Email
	customer.Password = c.Password
	customer.Orders = c.Orders
	customers.Store(customer.Username, customer)
	return nil
}

func OrdersCount() int {
	result := 0
	customers.Range(func(key, value interface{}) bool {
		c := value.(*Customer)
		result += len(c.Orders)
		return true
	})
	return result
}
