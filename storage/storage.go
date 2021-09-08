package storage

import "context"

type Storage interface {
	Insert(context.Context, interface{}, interface{}) error
	Range(context.Context, func(key, value interface{}) bool)
	Find(context.Context, interface{}) interface{}
	Delete(context.Context, interface{}) error
	Update(context.Context, interface{}, interface{}) error
}
