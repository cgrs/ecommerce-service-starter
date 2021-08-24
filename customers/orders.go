package customers

import (
	"sync"

	"github.com/cgrs/ecommerce-service-starter/items"
	"github.com/cgrs/ecommerce-service-starter/utils"
)

type Order struct {
	Username string
	Number   int
	Items    []items.Item
	Status   string
	Total    float64
}

var orders sync.Map

func OrdersCount() int {
	return utils.MapCount(orders)
}
