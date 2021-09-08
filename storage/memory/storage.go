package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/cgrs/ecommerce-service-starter/storage"
)

type memoryStorage struct {
	m sync.Map
}

func New() storage.Storage {
	return &memoryStorage{}
}

func (m *memoryStorage) Insert(ctx context.Context, key, value interface{}) error {
	if m.exists(ctx, key) {
		return fmt.Errorf(`key "%v" already exists`, key)
	}
	m.m.Store(key, value)
	return nil
}

func (m *memoryStorage) exists(ctx context.Context, key interface{}) bool {
	_, ok := m.m.Load(key)
	return ok
}

func (m *memoryStorage) Range(ctx context.Context, fn func(key, value interface{}) bool) {
	m.m.Range(fn)
}

func (m *memoryStorage) Find(ctx context.Context, key interface{}) interface{} {
	value, _ := m.m.Load(key)
	return value
}

func (m *memoryStorage) Update(ctx context.Context, key, value interface{}) error {
	if !m.exists(ctx, key) {
		return fmt.Errorf(`key "%v" not found`, key)
	}
	m.m.Store(key, value)
	return nil
}

func (m *memoryStorage) Delete(ctx context.Context, key interface{}) error {
	if !m.exists(ctx, key) {
		return fmt.Errorf(`key "%v" not found`, key)
	}
	m.m.Delete(key)
	return nil
}
