// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ordersv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrdersServiceClient is the client API for OrdersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersServiceClient interface {
	Create(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
	List(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Find(ctx context.Context, in *FindOrderRequest, opts ...grpc.CallOption) (*Order, error)
	Update(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*Order, error)
	Delete(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type ordersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersServiceClient(cc grpc.ClientConnInterface) OrdersServiceClient {
	return &ordersServiceClient{cc}
}

func (c *ordersServiceClient) Create(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/orders.v1.OrdersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) List(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/orders.v1.OrdersService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) Find(ctx context.Context, in *FindOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/orders.v1.OrdersService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) Update(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/orders.v1.OrdersService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) Delete(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/orders.v1.OrdersService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServiceServer is the server API for OrdersService service.
// All implementations must embed UnimplementedOrdersServiceServer
// for forward compatibility
type OrdersServiceServer interface {
	Create(context.Context, *CreateOrderRequest) (*Order, error)
	List(context.Context, *ListOrderRequest) (*ListResponse, error)
	Find(context.Context, *FindOrderRequest) (*Order, error)
	Update(context.Context, *UpdateOrderRequest) (*Order, error)
	Delete(context.Context, *DeleteOrderRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrdersServiceServer()
}

// UnimplementedOrdersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersServiceServer struct {
}

func (UnimplementedOrdersServiceServer) Create(context.Context, *CreateOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrdersServiceServer) List(context.Context, *ListOrderRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrdersServiceServer) Find(context.Context, *FindOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedOrdersServiceServer) Update(context.Context, *UpdateOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrdersServiceServer) Delete(context.Context, *DeleteOrderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOrdersServiceServer) mustEmbedUnimplementedOrdersServiceServer() {}

// UnsafeOrdersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServiceServer will
// result in compilation errors.
type UnsafeOrdersServiceServer interface {
	mustEmbedUnimplementedOrdersServiceServer()
}

func RegisterOrdersServiceServer(s grpc.ServiceRegistrar, srv OrdersServiceServer) {
	s.RegisterService(&OrdersService_ServiceDesc, srv)
}

func _OrdersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.v1.OrdersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).Create(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.v1.OrdersService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).List(ctx, req.(*ListOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.v1.OrdersService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).Find(ctx, req.(*FindOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.v1.OrdersService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).Update(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.v1.OrdersService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).Delete(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrdersService_ServiceDesc is the grpc.ServiceDesc for OrdersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.v1.OrdersService",
	HandlerType: (*OrdersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrdersService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _OrdersService_List_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _OrdersService_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrdersService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrdersService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/v1/orders.proto",
}
