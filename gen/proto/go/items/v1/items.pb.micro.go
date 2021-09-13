// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: items/v1/items.proto

package itemsv1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Items service

func NewItemsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Items service

type ItemsService interface {
	Create(ctx context.Context, in *CreateItemRequest, opts ...client.CallOption) (*CreateItemResponse, error)
	Filter(ctx context.Context, in *FilterItemRequest, opts ...client.CallOption) (*FilterItemResponse, error)
	List(ctx context.Context, in *ListItemRequest, opts ...client.CallOption) (*ListItemResponse, error)
	Find(ctx context.Context, in *FindItemRequest, opts ...client.CallOption) (*FindItemResponse, error)
	Update(ctx context.Context, in *UpdateItemRequest, opts ...client.CallOption) (*UpdateItemResponse, error)
	Delete(ctx context.Context, in *DeleteItemRequest, opts ...client.CallOption) (*DeleteItemResponse, error)
}

type itemsService struct {
	c    client.Client
	name string
}

func NewItemsService(name string, c client.Client) ItemsService {
	return &itemsService{
		c:    c,
		name: name,
	}
}

func (c *itemsService) Create(ctx context.Context, in *CreateItemRequest, opts ...client.CallOption) (*CreateItemResponse, error) {
	req := c.c.NewRequest(c.name, "Items.Create", in)
	out := new(CreateItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsService) Filter(ctx context.Context, in *FilterItemRequest, opts ...client.CallOption) (*FilterItemResponse, error) {
	req := c.c.NewRequest(c.name, "Items.Filter", in)
	out := new(FilterItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsService) List(ctx context.Context, in *ListItemRequest, opts ...client.CallOption) (*ListItemResponse, error) {
	req := c.c.NewRequest(c.name, "Items.List", in)
	out := new(ListItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsService) Find(ctx context.Context, in *FindItemRequest, opts ...client.CallOption) (*FindItemResponse, error) {
	req := c.c.NewRequest(c.name, "Items.Find", in)
	out := new(FindItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsService) Update(ctx context.Context, in *UpdateItemRequest, opts ...client.CallOption) (*UpdateItemResponse, error) {
	req := c.c.NewRequest(c.name, "Items.Update", in)
	out := new(UpdateItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsService) Delete(ctx context.Context, in *DeleteItemRequest, opts ...client.CallOption) (*DeleteItemResponse, error) {
	req := c.c.NewRequest(c.name, "Items.Delete", in)
	out := new(DeleteItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Items service

type ItemsHandler interface {
	Create(context.Context, *CreateItemRequest, *CreateItemResponse) error
	Filter(context.Context, *FilterItemRequest, *FilterItemResponse) error
	List(context.Context, *ListItemRequest, *ListItemResponse) error
	Find(context.Context, *FindItemRequest, *FindItemResponse) error
	Update(context.Context, *UpdateItemRequest, *UpdateItemResponse) error
	Delete(context.Context, *DeleteItemRequest, *DeleteItemResponse) error
}

func RegisterItemsHandler(s server.Server, hdlr ItemsHandler, opts ...server.HandlerOption) error {
	type items interface {
		Create(ctx context.Context, in *CreateItemRequest, out *CreateItemResponse) error
		Filter(ctx context.Context, in *FilterItemRequest, out *FilterItemResponse) error
		List(ctx context.Context, in *ListItemRequest, out *ListItemResponse) error
		Find(ctx context.Context, in *FindItemRequest, out *FindItemResponse) error
		Update(ctx context.Context, in *UpdateItemRequest, out *UpdateItemResponse) error
		Delete(ctx context.Context, in *DeleteItemRequest, out *DeleteItemResponse) error
	}
	type Items struct {
		items
	}
	h := &itemsHandler{hdlr}
	return s.Handle(s.NewHandler(&Items{h}, opts...))
}

type itemsHandler struct {
	ItemsHandler
}

func (h *itemsHandler) Create(ctx context.Context, in *CreateItemRequest, out *CreateItemResponse) error {
	return h.ItemsHandler.Create(ctx, in, out)
}

func (h *itemsHandler) Filter(ctx context.Context, in *FilterItemRequest, out *FilterItemResponse) error {
	return h.ItemsHandler.Filter(ctx, in, out)
}

func (h *itemsHandler) List(ctx context.Context, in *ListItemRequest, out *ListItemResponse) error {
	return h.ItemsHandler.List(ctx, in, out)
}

func (h *itemsHandler) Find(ctx context.Context, in *FindItemRequest, out *FindItemResponse) error {
	return h.ItemsHandler.Find(ctx, in, out)
}

func (h *itemsHandler) Update(ctx context.Context, in *UpdateItemRequest, out *UpdateItemResponse) error {
	return h.ItemsHandler.Update(ctx, in, out)
}

func (h *itemsHandler) Delete(ctx context.Context, in *DeleteItemRequest, out *DeleteItemResponse) error {
	return h.ItemsHandler.Delete(ctx, in, out)
}