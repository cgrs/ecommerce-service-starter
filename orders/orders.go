package orders

import (
	"sync"

	"github.com/cgrs/ecommerce-service-starter/cart"
)

type Order struct {
	Username string      `json:"-"`
	Number   int         `json:"order_id"`
	Lines    []cart.Line `json:"lines"`
	Status   string      `json:"status"`
	Total    float64     `json:"total"`
}

var orders sync.Map

func Store(o *Order) {
	orders.Store(o.Number, o)
}

func FindByCustomer(username string) []*Order {
	result := make([]*Order, 0)
	orders.Range(func(key, value interface{}) bool {
		order := value.(*Order)
		if order.Username == username {
			result = append(result, order)
		}
		return true
	})
	return result
}
