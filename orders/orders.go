package orders

import (
	"sync"

	"github.com/cgrs/ecommerce-service-starter/cart"
)

type Order struct {
	Username string
	Number   int
	Lines    []cart.Line
	Status   string
	Total    float64
}

var orders sync.Map

func Store(o *Order) {
	orders.Store(o.Number, o)
}
