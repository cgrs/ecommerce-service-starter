package storage

import "context"

type Storage interface {
	Insert(ctx context.Context, key, value interface{}) error
	Find(ctx context.Context, key interface{}) interface{}
	Update(ctx context.Context, key, value interface{}) error
	FetchAll(ctx context.Context) map[interface{}]interface{}
	Delete(ctx context.Context, key interface{}) error
}
