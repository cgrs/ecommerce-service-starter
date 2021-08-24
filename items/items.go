package items

import (
	"fmt"
	"sync"

	"github.com/cgrs/ecommerce-service-starter/utils"
)

type Item struct {
	Name        string
	Description string
	UnitPrice   float64
}

var items sync.Map
var updateLock sync.Mutex

func Count() int {
	return utils.MapCount(items)
}

func Add(i *Item) error {
	_, ok := items.LoadOrStore(i.Name, i)
	if ok {
		return fmt.Errorf("item already exists: %s", i.Name)
	}
	return nil
}

func Update(oldName string, i *Item) error {
	updateLock.Lock()
	defer updateLock.Unlock()
	value, ok := items.LoadAndDelete(oldName)
	if !ok {
		return fmt.Errorf("item not found for update: %s", oldName)
	}
	item, ok := value.(*Item)
	if !ok {
		return fmt.Errorf("unexpected conversion error")
	}
	if oldName != i.Name {
		item.Name = i.Name
	}
	item.Description = i.Description
	item.UnitPrice = i.UnitPrice
	items.Store(item.Name, item)
	return nil
}
