package items

import (
	"fmt"
	"sync"

	"github.com/cgrs/ecommerce-service-starter/utils"
)

type Item struct {
	ID          int     `json:"-"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"price"`
}

var items sync.Map
var updateLock sync.Mutex

func Count() int {
	return utils.MapCount(items)
}

func Add(i *Item) error {
	_, ok := items.LoadOrStore(i.ID, i)
	if ok {
		return fmt.Errorf("item already exists: %s", i.Name)
	}
	return nil
}

func Update(i *Item) error {
	updateLock.Lock()
	defer updateLock.Unlock()
	value, ok := items.LoadAndDelete(i.ID)
	if !ok {
		return fmt.Errorf("item not found for update: %s", i.Name)
	}
	item, ok := value.(*Item)
	if !ok {
		return fmt.Errorf("unexpected conversion error")
	}
	item.Name = i.Name
	item.Description = i.Description
	item.UnitPrice = i.UnitPrice
	items.Store(item.ID, item)
	return nil
}

func Find(id int) *Item {
	value, ok := items.Load(id)
	if !ok {
		return nil
	}
	return value.(*Item)
}

func List() []*Item {
	itms := make([]*Item, Count())
	ix := 0
	items.Range(func(key, value interface{}) bool {
		itms[ix] = value.(*Item)
		ix++
		return true
	})
	return itms
}
