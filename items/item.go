package items

import (
	"fmt"
	"sync"
)

type Item struct {
	Name string
	Description string
	Price float64
}

var items sync.Map

func AddItem(i *Item) (*Item, error) {
	if i == nil {
		return nil, fmt.Errorf("item is empty")
	}
	if _, ok := items.Load(i.Name); ok {
		return nil, fmt.Errorf("item exists with same name")
	}
	items.Store(i.Name, i)
	return i, nil
}

func FindItem(name string) *Item {
	v, _ := items.Load(name);
	item, _ := v.(*Item);
	return item;
}