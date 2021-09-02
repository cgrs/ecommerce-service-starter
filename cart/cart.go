package cart

import (
	"fmt"
	"sync"

	"github.com/cgrs/ecommerce-service-starter/items"
	"github.com/cgrs/ecommerce-service-starter/utils"
)

type Cart struct {
	ID       string `json:"-"`
	Customer struct {
		Username string
	} `json:"-"`
	Lines []Line `json:"lines"`
}

type Line struct {
	ItemId   int `json:"item_id"`
	Quantity int `json:"qty"`
}

var carts sync.Map

func Count() int {
	return utils.MapCount(carts)
}

func Find(id string) *Cart {
	value, ok := carts.Load(id)
	if !ok {
		return nil
	}
	return value.(*Cart)
}

func Create(id, username string) error {
	cart := Find(id)
	if cart != nil {
		return fmt.Errorf("cart already exists")
	}
	cart = &Cart{
		ID:       id,
		Customer: struct{ Username string }{username},
		Lines:    []Line{},
	}
	carts.Store(id, cart)
	return nil
}

func Delete(id string) error {
	_, ok := carts.LoadAndDelete(id)
	if !ok {
		return fmt.Errorf("cart does not exist")
	}
	return nil
}

func Update(id string, lines []Line) error {
	cart := Find(id)
	if cart == nil {
		return fmt.Errorf("cart does not exist")
	}
	cart.Lines = lines
	carts.Store(id, cart)
	return nil
}

func (c *Cart) GetTotal() float64 {
	t := 0.0
	for _, v := range c.Lines {
		i := items.Find(v.ItemId)
		t += i.UnitPrice * float64(v.Quantity)
	}
	return t
}

func Empty(id string) error {
	return Update(id, []Line{})
}
