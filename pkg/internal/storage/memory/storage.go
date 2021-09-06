package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/cgrs/ecommerce-service-starter/pkg/internal/storage"
)

type memoryStorage struct {
	storage sync.Map
}

func NewStorage() storage.Storage {
	return &memoryStorage{}
}

func (s *memoryStorage) Insert(ctx context.Context, key, value interface{}) error {
	if s.exists(ctx, key) {
		return fmt.Errorf(`item with key "%v" already exists in storage`, key)
	}
	s.storage.Store(key, value)
	return nil
}

func (s *memoryStorage) Find(ctx context.Context, key interface{}) interface{} {
	value, _ := s.storage.Load(key)
	return value
}

func (s *memoryStorage) Update(ctx context.Context, key, value interface{}) error {
	if !s.exists(ctx, key) {
		return fmt.Errorf(`item with key "%v" does not exist in storage`, key)
	}
	s.storage.Store(key, value)
	return nil
}

func (s *memoryStorage) FetchAll(ctx context.Context) map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	s.storage.Range(func(key, value interface{}) bool {
		result[key] = value
		return true
	})
	return result
}

func (s *memoryStorage) Delete(ctx context.Context, key interface{}) error {
	if !s.exists(ctx, key) {
		return fmt.Errorf(`item with key "%v" does not exist in storage`, key)
	}
	s.storage.Delete(key)
	return nil
}

func (s *memoryStorage) exists(ctx context.Context, key interface{}) bool {
	_, found := s.storage.Load(key)
	return found
}
