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

func ListItems() []*Item {
	res := make([]*Item, 0)
	items.Range(func(key, value interface{}) bool {
		i, _ := value.(*Item)
		res = append(res, i)
		return true
	})
	return res
}