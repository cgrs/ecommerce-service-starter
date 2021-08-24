package cart

import (
	"sync"

	"github.com/cgrs/ecommerce-service-starter/utils"
)

type Cart struct {
}

var carts sync.Map

func Count() int {
	return utils.MapCount(carts)
}
