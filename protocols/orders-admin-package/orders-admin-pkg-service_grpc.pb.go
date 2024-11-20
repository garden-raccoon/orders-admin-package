// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: orders-admin-pkg-service.proto

package orders_admin_package

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OrderAdminService_GetOrders_FullMethodName    = "/models.OrderAdminService/GetOrders"
	OrderAdminService_CreateOrders_FullMethodName = "/models.OrderAdminService/CreateOrders"
	OrderAdminService_OrderByName_FullMethodName  = "/models.OrderAdminService/OrderByName"
)

// OrderAdminServiceClient is the client API for OrderAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// OrderService is
type OrderAdminServiceClient interface {
	GetOrders(ctx context.Context, in *OrderAdminEmpty, opts ...grpc.CallOption) (*OrdersAdmin, error)
	CreateOrders(ctx context.Context, in *OrdersAdmin, opts ...grpc.CallOption) (*OrderAdminEmpty, error)
	OrderByName(ctx context.Context, in *OrderAdminGetter, opts ...grpc.CallOption) (*OrderAdmin, error)
}

type orderAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderAdminServiceClient(cc grpc.ClientConnInterface) OrderAdminServiceClient {
	return &orderAdminServiceClient{cc}
}

func (c *orderAdminServiceClient) GetOrders(ctx context.Context, in *OrderAdminEmpty, opts ...grpc.CallOption) (*OrdersAdmin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrdersAdmin)
	err := c.cc.Invoke(ctx, OrderAdminService_GetOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAdminServiceClient) CreateOrders(ctx context.Context, in *OrdersAdmin, opts ...grpc.CallOption) (*OrderAdminEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderAdminEmpty)
	err := c.cc.Invoke(ctx, OrderAdminService_CreateOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAdminServiceClient) OrderByName(ctx context.Context, in *OrderAdminGetter, opts ...grpc.CallOption) (*OrderAdmin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderAdmin)
	err := c.cc.Invoke(ctx, OrderAdminService_OrderByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderAdminServiceServer is the server API for OrderAdminService service.
// All implementations must embed UnimplementedOrderAdminServiceServer
// for forward compatibility.
//
// OrderService is
type OrderAdminServiceServer interface {
	GetOrders(context.Context, *OrderAdminEmpty) (*OrdersAdmin, error)
	CreateOrders(context.Context, *OrdersAdmin) (*OrderAdminEmpty, error)
	OrderByName(context.Context, *OrderAdminGetter) (*OrderAdmin, error)
	mustEmbedUnimplementedOrderAdminServiceServer()
}

// UnimplementedOrderAdminServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderAdminServiceServer struct{}

func (UnimplementedOrderAdminServiceServer) GetOrders(context.Context, *OrderAdminEmpty) (*OrdersAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrderAdminServiceServer) CreateOrders(context.Context, *OrdersAdmin) (*OrderAdminEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrders not implemented")
}
func (UnimplementedOrderAdminServiceServer) OrderByName(context.Context, *OrderAdminGetter) (*OrderAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderByName not implemented")
}
func (UnimplementedOrderAdminServiceServer) mustEmbedUnimplementedOrderAdminServiceServer() {}
func (UnimplementedOrderAdminServiceServer) testEmbeddedByValue()                           {}

// UnsafeOrderAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderAdminServiceServer will
// result in compilation errors.
type UnsafeOrderAdminServiceServer interface {
	mustEmbedUnimplementedOrderAdminServiceServer()
}

func RegisterOrderAdminServiceServer(s grpc.ServiceRegistrar, srv OrderAdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderAdminService_ServiceDesc, srv)
}

func _OrderAdminService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderAdminEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAdminServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAdminService_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAdminServiceServer).GetOrders(ctx, req.(*OrderAdminEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAdminService_CreateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAdminServiceServer).CreateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAdminService_CreateOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAdminServiceServer).CreateOrders(ctx, req.(*OrdersAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAdminService_OrderByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderAdminGetter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAdminServiceServer).OrderByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAdminService_OrderByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAdminServiceServer).OrderByName(ctx, req.(*OrderAdminGetter))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderAdminService_ServiceDesc is the grpc.ServiceDesc for OrderAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.OrderAdminService",
	HandlerType: (*OrderAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrders",
			Handler:    _OrderAdminService_GetOrders_Handler,
		},
		{
			MethodName: "CreateOrders",
			Handler:    _OrderAdminService_CreateOrders_Handler,
		},
		{
			MethodName: "OrderByName",
			Handler:    _OrderAdminService_OrderByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders-admin-pkg-service.proto",
}
