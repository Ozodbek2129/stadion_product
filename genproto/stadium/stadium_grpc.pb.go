// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: stadion_protos/stadium/stadium.proto

package stadium

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
	StadiumService_CreateStadium_FullMethodName           = "/stadium.StadiumService/CreateStadium"
	StadiumService_UpdateStadium_FullMethodName           = "/stadium.StadiumService/UpdateStadium"
	StadiumService_GetStadium_FullMethodName              = "/stadium.StadiumService/GetStadium"
	StadiumService_GetStadiums_FullMethodName             = "/stadium.StadiumService/GetStadiums"
	StadiumService_DeleteStadium_FullMethodName           = "/stadium.StadiumService/DeleteStadium"
	StadiumService_CreateOrderStadium_FullMethodName      = "/stadium.StadiumService/CreateOrderStadium"
	StadiumService_GetOrderStadiums_FullMethodName        = "/stadium.StadiumService/GetOrderStadiums"
	StadiumService_GetOrderStadium_FullMethodName         = "/stadium.StadiumService/GetOrderStadium"
	StadiumService_UpdateOrderStadium_FullMethodName      = "/stadium.StadiumService/UpdateOrderStadium"
	StadiumService_DeleteOrderStadium_FullMethodName      = "/stadium.StadiumService/DeleteOrderStadium"
	StadiumService_GetDeletedOrderStadiums_FullMethodName = "/stadium.StadiumService/GetDeletedOrderStadiums"
)

// StadiumServiceClient is the client API for StadiumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StadiumServiceClient interface {
	CreateStadium(ctx context.Context, in *CreateStadiumRequest, opts ...grpc.CallOption) (*CreateStadiumResponse, error)
	UpdateStadium(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetStadium(ctx context.Context, in *GetStadiumRequest, opts ...grpc.CallOption) (*GetStadiumResponse, error)
	GetStadiums(ctx context.Context, in *GetStadiumsRequest, opts ...grpc.CallOption) (*GetStadiumsResponse, error)
	DeleteStadium(ctx context.Context, in *DeleteStadiumRequest, opts ...grpc.CallOption) (*DeleteStadiumResponse, error)
	CreateOrderStadium(ctx context.Context, in *CreateOrderStadiumRequest, opts ...grpc.CallOption) (*CreateOrderStadiumResponse, error)
	GetOrderStadiums(ctx context.Context, in *GetOrderStadiumsRequest, opts ...grpc.CallOption) (*GetOrderStadiumsResponse, error)
	GetOrderStadium(ctx context.Context, in *GetOrderStadiumRequest, opts ...grpc.CallOption) (*GetOrderStadiumResponse, error)
	UpdateOrderStadium(ctx context.Context, in *UpdateOrderStadiumRequest, opts ...grpc.CallOption) (*UpdateOrderStadiumResponse, error)
	DeleteOrderStadium(ctx context.Context, in *DeleteOrderStadiumRequest, opts ...grpc.CallOption) (*DeleteOrderStadiumResponse, error)
	GetDeletedOrderStadiums(ctx context.Context, in *GetDeletedOrderStadiumsRequest, opts ...grpc.CallOption) (*GetDeletedOrderStadiumsResponse, error)
}

type stadiumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStadiumServiceClient(cc grpc.ClientConnInterface) StadiumServiceClient {
	return &stadiumServiceClient{cc}
}

func (c *stadiumServiceClient) CreateStadium(ctx context.Context, in *CreateStadiumRequest, opts ...grpc.CallOption) (*CreateStadiumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStadiumResponse)
	err := c.cc.Invoke(ctx, StadiumService_CreateStadium_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) UpdateStadium(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, StadiumService_UpdateStadium_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) GetStadium(ctx context.Context, in *GetStadiumRequest, opts ...grpc.CallOption) (*GetStadiumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStadiumResponse)
	err := c.cc.Invoke(ctx, StadiumService_GetStadium_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) GetStadiums(ctx context.Context, in *GetStadiumsRequest, opts ...grpc.CallOption) (*GetStadiumsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStadiumsResponse)
	err := c.cc.Invoke(ctx, StadiumService_GetStadiums_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) DeleteStadium(ctx context.Context, in *DeleteStadiumRequest, opts ...grpc.CallOption) (*DeleteStadiumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteStadiumResponse)
	err := c.cc.Invoke(ctx, StadiumService_DeleteStadium_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) CreateOrderStadium(ctx context.Context, in *CreateOrderStadiumRequest, opts ...grpc.CallOption) (*CreateOrderStadiumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrderStadiumResponse)
	err := c.cc.Invoke(ctx, StadiumService_CreateOrderStadium_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) GetOrderStadiums(ctx context.Context, in *GetOrderStadiumsRequest, opts ...grpc.CallOption) (*GetOrderStadiumsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderStadiumsResponse)
	err := c.cc.Invoke(ctx, StadiumService_GetOrderStadiums_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) GetOrderStadium(ctx context.Context, in *GetOrderStadiumRequest, opts ...grpc.CallOption) (*GetOrderStadiumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderStadiumResponse)
	err := c.cc.Invoke(ctx, StadiumService_GetOrderStadium_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) UpdateOrderStadium(ctx context.Context, in *UpdateOrderStadiumRequest, opts ...grpc.CallOption) (*UpdateOrderStadiumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateOrderStadiumResponse)
	err := c.cc.Invoke(ctx, StadiumService_UpdateOrderStadium_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) DeleteOrderStadium(ctx context.Context, in *DeleteOrderStadiumRequest, opts ...grpc.CallOption) (*DeleteOrderStadiumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteOrderStadiumResponse)
	err := c.cc.Invoke(ctx, StadiumService_DeleteOrderStadium_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stadiumServiceClient) GetDeletedOrderStadiums(ctx context.Context, in *GetDeletedOrderStadiumsRequest, opts ...grpc.CallOption) (*GetDeletedOrderStadiumsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDeletedOrderStadiumsResponse)
	err := c.cc.Invoke(ctx, StadiumService_GetDeletedOrderStadiums_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StadiumServiceServer is the server API for StadiumService service.
// All implementations must embed UnimplementedStadiumServiceServer
// for forward compatibility.
type StadiumServiceServer interface {
	CreateStadium(context.Context, *CreateStadiumRequest) (*CreateStadiumResponse, error)
	UpdateStadium(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetStadium(context.Context, *GetStadiumRequest) (*GetStadiumResponse, error)
	GetStadiums(context.Context, *GetStadiumsRequest) (*GetStadiumsResponse, error)
	DeleteStadium(context.Context, *DeleteStadiumRequest) (*DeleteStadiumResponse, error)
	CreateOrderStadium(context.Context, *CreateOrderStadiumRequest) (*CreateOrderStadiumResponse, error)
	GetOrderStadiums(context.Context, *GetOrderStadiumsRequest) (*GetOrderStadiumsResponse, error)
	GetOrderStadium(context.Context, *GetOrderStadiumRequest) (*GetOrderStadiumResponse, error)
	UpdateOrderStadium(context.Context, *UpdateOrderStadiumRequest) (*UpdateOrderStadiumResponse, error)
	DeleteOrderStadium(context.Context, *DeleteOrderStadiumRequest) (*DeleteOrderStadiumResponse, error)
	GetDeletedOrderStadiums(context.Context, *GetDeletedOrderStadiumsRequest) (*GetDeletedOrderStadiumsResponse, error)
	mustEmbedUnimplementedStadiumServiceServer()
}

// UnimplementedStadiumServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStadiumServiceServer struct{}

func (UnimplementedStadiumServiceServer) CreateStadium(context.Context, *CreateStadiumRequest) (*CreateStadiumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStadium not implemented")
}
func (UnimplementedStadiumServiceServer) UpdateStadium(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStadium not implemented")
}
func (UnimplementedStadiumServiceServer) GetStadium(context.Context, *GetStadiumRequest) (*GetStadiumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStadium not implemented")
}
func (UnimplementedStadiumServiceServer) GetStadiums(context.Context, *GetStadiumsRequest) (*GetStadiumsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStadiums not implemented")
}
func (UnimplementedStadiumServiceServer) DeleteStadium(context.Context, *DeleteStadiumRequest) (*DeleteStadiumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStadium not implemented")
}
func (UnimplementedStadiumServiceServer) CreateOrderStadium(context.Context, *CreateOrderStadiumRequest) (*CreateOrderStadiumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderStadium not implemented")
}
func (UnimplementedStadiumServiceServer) GetOrderStadiums(context.Context, *GetOrderStadiumsRequest) (*GetOrderStadiumsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStadiums not implemented")
}
func (UnimplementedStadiumServiceServer) GetOrderStadium(context.Context, *GetOrderStadiumRequest) (*GetOrderStadiumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStadium not implemented")
}
func (UnimplementedStadiumServiceServer) UpdateOrderStadium(context.Context, *UpdateOrderStadiumRequest) (*UpdateOrderStadiumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStadium not implemented")
}
func (UnimplementedStadiumServiceServer) DeleteOrderStadium(context.Context, *DeleteOrderStadiumRequest) (*DeleteOrderStadiumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderStadium not implemented")
}
func (UnimplementedStadiumServiceServer) GetDeletedOrderStadiums(context.Context, *GetDeletedOrderStadiumsRequest) (*GetDeletedOrderStadiumsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeletedOrderStadiums not implemented")
}
func (UnimplementedStadiumServiceServer) mustEmbedUnimplementedStadiumServiceServer() {}
func (UnimplementedStadiumServiceServer) testEmbeddedByValue()                        {}

// UnsafeStadiumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StadiumServiceServer will
// result in compilation errors.
type UnsafeStadiumServiceServer interface {
	mustEmbedUnimplementedStadiumServiceServer()
}

func RegisterStadiumServiceServer(s grpc.ServiceRegistrar, srv StadiumServiceServer) {
	// If the following call pancis, it indicates UnimplementedStadiumServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StadiumService_ServiceDesc, srv)
}

func _StadiumService_CreateStadium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStadiumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).CreateStadium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_CreateStadium_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).CreateStadium(ctx, req.(*CreateStadiumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_UpdateStadium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).UpdateStadium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_UpdateStadium_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).UpdateStadium(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_GetStadium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStadiumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).GetStadium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_GetStadium_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).GetStadium(ctx, req.(*GetStadiumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_GetStadiums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStadiumsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).GetStadiums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_GetStadiums_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).GetStadiums(ctx, req.(*GetStadiumsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_DeleteStadium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStadiumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).DeleteStadium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_DeleteStadium_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).DeleteStadium(ctx, req.(*DeleteStadiumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_CreateOrderStadium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderStadiumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).CreateOrderStadium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_CreateOrderStadium_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).CreateOrderStadium(ctx, req.(*CreateOrderStadiumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_GetOrderStadiums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderStadiumsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).GetOrderStadiums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_GetOrderStadiums_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).GetOrderStadiums(ctx, req.(*GetOrderStadiumsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_GetOrderStadium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderStadiumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).GetOrderStadium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_GetOrderStadium_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).GetOrderStadium(ctx, req.(*GetOrderStadiumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_UpdateOrderStadium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStadiumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).UpdateOrderStadium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_UpdateOrderStadium_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).UpdateOrderStadium(ctx, req.(*UpdateOrderStadiumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_DeleteOrderStadium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderStadiumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).DeleteOrderStadium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_DeleteOrderStadium_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).DeleteOrderStadium(ctx, req.(*DeleteOrderStadiumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StadiumService_GetDeletedOrderStadiums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeletedOrderStadiumsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StadiumServiceServer).GetDeletedOrderStadiums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StadiumService_GetDeletedOrderStadiums_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StadiumServiceServer).GetDeletedOrderStadiums(ctx, req.(*GetDeletedOrderStadiumsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StadiumService_ServiceDesc is the grpc.ServiceDesc for StadiumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StadiumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stadium.StadiumService",
	HandlerType: (*StadiumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStadium",
			Handler:    _StadiumService_CreateStadium_Handler,
		},
		{
			MethodName: "UpdateStadium",
			Handler:    _StadiumService_UpdateStadium_Handler,
		},
		{
			MethodName: "GetStadium",
			Handler:    _StadiumService_GetStadium_Handler,
		},
		{
			MethodName: "GetStadiums",
			Handler:    _StadiumService_GetStadiums_Handler,
		},
		{
			MethodName: "DeleteStadium",
			Handler:    _StadiumService_DeleteStadium_Handler,
		},
		{
			MethodName: "CreateOrderStadium",
			Handler:    _StadiumService_CreateOrderStadium_Handler,
		},
		{
			MethodName: "GetOrderStadiums",
			Handler:    _StadiumService_GetOrderStadiums_Handler,
		},
		{
			MethodName: "GetOrderStadium",
			Handler:    _StadiumService_GetOrderStadium_Handler,
		},
		{
			MethodName: "UpdateOrderStadium",
			Handler:    _StadiumService_UpdateOrderStadium_Handler,
		},
		{
			MethodName: "DeleteOrderStadium",
			Handler:    _StadiumService_DeleteOrderStadium_Handler,
		},
		{
			MethodName: "GetDeletedOrderStadiums",
			Handler:    _StadiumService_GetDeletedOrderStadiums_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stadion_protos/stadium/stadium.proto",
}
