// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.5
// source: server.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Service_Login_FullMethodName                 = "/Service/Login"
	Service_GetTaskListAll_FullMethodName        = "/Service/GetTaskListAll"
	Service_GetTaskListOne_FullMethodName        = "/Service/GetTaskListOne"
	Service_ImportToTaskListTable_FullMethodName = "/Service/ImportToTaskListTable"
	Service_ImportXLSToPatchTable_FullMethodName = "/Service/ImportXLSToPatchTable"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义服务
type ServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	GetTaskListAll(ctx context.Context, in *GetTaskListAllRequest, opts ...grpc.CallOption) (*GetTaskListAllReply, error)
	GetTaskListOne(ctx context.Context, in *GetTaskListOneRequest, opts ...grpc.CallOption) (*GetTaskListOneReply, error)
	ImportToTaskListTable(ctx context.Context, in *ImportToTaskListRequest, opts ...grpc.CallOption) (*ImportToTaskListReply, error)
	ImportXLSToPatchTable(ctx context.Context, in *ImportXLSToPatchRequest, opts ...grpc.CallOption) (*ImportXLSToPatchReply, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Service_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTaskListAll(ctx context.Context, in *GetTaskListAllRequest, opts ...grpc.CallOption) (*GetTaskListAllReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskListAllReply)
	err := c.cc.Invoke(ctx, Service_GetTaskListAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTaskListOne(ctx context.Context, in *GetTaskListOneRequest, opts ...grpc.CallOption) (*GetTaskListOneReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskListOneReply)
	err := c.cc.Invoke(ctx, Service_GetTaskListOne_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ImportToTaskListTable(ctx context.Context, in *ImportToTaskListRequest, opts ...grpc.CallOption) (*ImportToTaskListReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImportToTaskListReply)
	err := c.cc.Invoke(ctx, Service_ImportToTaskListTable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ImportXLSToPatchTable(ctx context.Context, in *ImportXLSToPatchRequest, opts ...grpc.CallOption) (*ImportXLSToPatchReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImportXLSToPatchReply)
	err := c.cc.Invoke(ctx, Service_ImportXLSToPatchTable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
//
// 定义服务
type ServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	GetTaskListAll(context.Context, *GetTaskListAllRequest) (*GetTaskListAllReply, error)
	GetTaskListOne(context.Context, *GetTaskListOneRequest) (*GetTaskListOneReply, error)
	ImportToTaskListTable(context.Context, *ImportToTaskListRequest) (*ImportToTaskListReply, error)
	ImportXLSToPatchTable(context.Context, *ImportXLSToPatchRequest) (*ImportXLSToPatchReply, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedServiceServer) GetTaskListAll(context.Context, *GetTaskListAllRequest) (*GetTaskListAllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskListAll not implemented")
}
func (UnimplementedServiceServer) GetTaskListOne(context.Context, *GetTaskListOneRequest) (*GetTaskListOneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskListOne not implemented")
}
func (UnimplementedServiceServer) ImportToTaskListTable(context.Context, *ImportToTaskListRequest) (*ImportToTaskListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportToTaskListTable not implemented")
}
func (UnimplementedServiceServer) ImportXLSToPatchTable(context.Context, *ImportXLSToPatchRequest) (*ImportXLSToPatchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportXLSToPatchTable not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTaskListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskListAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTaskListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetTaskListAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTaskListAll(ctx, req.(*GetTaskListAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTaskListOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskListOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTaskListOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetTaskListOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTaskListOne(ctx, req.(*GetTaskListOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ImportToTaskListTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportToTaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ImportToTaskListTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ImportToTaskListTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ImportToTaskListTable(ctx, req.(*ImportToTaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ImportXLSToPatchTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportXLSToPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ImportXLSToPatchTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ImportXLSToPatchTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ImportXLSToPatchTable(ctx, req.(*ImportXLSToPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Service_Login_Handler,
		},
		{
			MethodName: "GetTaskListAll",
			Handler:    _Service_GetTaskListAll_Handler,
		},
		{
			MethodName: "GetTaskListOne",
			Handler:    _Service_GetTaskListOne_Handler,
		},
		{
			MethodName: "ImportToTaskListTable",
			Handler:    _Service_ImportToTaskListTable_Handler,
		},
		{
			MethodName: "ImportXLSToPatchTable",
			Handler:    _Service_ImportXLSToPatchTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
