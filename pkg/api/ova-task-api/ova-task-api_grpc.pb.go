// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_task_api

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

// OvaTaskApiClient is the client API for OvaTaskApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OvaTaskApiClient interface {
	CreateTaskV1(ctx context.Context, in *CreateTaskV1Request, opts ...grpc.CallOption) (*CreateTaskV1Response, error)
	DescribeTaskV1(ctx context.Context, in *DescribeTaskV1Request, opts ...grpc.CallOption) (*DescribeTaskV1Response, error)
	ListTasksV1(ctx context.Context, in *ListTasksV1Request, opts ...grpc.CallOption) (*ListTasksV1Response, error)
	RemoveTasksV1(ctx context.Context, in *RemoveTaskV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type ovaTaskApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOvaTaskApiClient(cc grpc.ClientConnInterface) OvaTaskApiClient {
	return &ovaTaskApiClient{cc}
}

func (c *ovaTaskApiClient) CreateTaskV1(ctx context.Context, in *CreateTaskV1Request, opts ...grpc.CallOption) (*CreateTaskV1Response, error) {
	out := new(CreateTaskV1Response)
	err := c.cc.Invoke(ctx, "/ova.task.api.OvaTaskApi/CreateTaskV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaTaskApiClient) DescribeTaskV1(ctx context.Context, in *DescribeTaskV1Request, opts ...grpc.CallOption) (*DescribeTaskV1Response, error) {
	out := new(DescribeTaskV1Response)
	err := c.cc.Invoke(ctx, "/ova.task.api.OvaTaskApi/DescribeTaskV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaTaskApiClient) ListTasksV1(ctx context.Context, in *ListTasksV1Request, opts ...grpc.CallOption) (*ListTasksV1Response, error) {
	out := new(ListTasksV1Response)
	err := c.cc.Invoke(ctx, "/ova.task.api.OvaTaskApi/ListTasksV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaTaskApiClient) RemoveTasksV1(ctx context.Context, in *RemoveTaskV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.task.api.OvaTaskApi/RemoveTasksV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OvaTaskApiServer is the server API for OvaTaskApi service.
// All implementations must embed UnimplementedOvaTaskApiServer
// for forward compatibility
type OvaTaskApiServer interface {
	CreateTaskV1(context.Context, *CreateTaskV1Request) (*CreateTaskV1Response, error)
	DescribeTaskV1(context.Context, *DescribeTaskV1Request) (*DescribeTaskV1Response, error)
	ListTasksV1(context.Context, *ListTasksV1Request) (*ListTasksV1Response, error)
	RemoveTasksV1(context.Context, *RemoveTaskV1Request) (*emptypb.Empty, error)
	mustEmbedUnimplementedOvaTaskApiServer()
}

// UnimplementedOvaTaskApiServer must be embedded to have forward compatible implementations.
type UnimplementedOvaTaskApiServer struct {
}

func (UnimplementedOvaTaskApiServer) CreateTaskV1(context.Context, *CreateTaskV1Request) (*CreateTaskV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskV1 not implemented")
}
func (UnimplementedOvaTaskApiServer) DescribeTaskV1(context.Context, *DescribeTaskV1Request) (*DescribeTaskV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTaskV1 not implemented")
}
func (UnimplementedOvaTaskApiServer) ListTasksV1(context.Context, *ListTasksV1Request) (*ListTasksV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasksV1 not implemented")
}
func (UnimplementedOvaTaskApiServer) RemoveTasksV1(context.Context, *RemoveTaskV1Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTasksV1 not implemented")
}
func (UnimplementedOvaTaskApiServer) mustEmbedUnimplementedOvaTaskApiServer() {}

// UnsafeOvaTaskApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OvaTaskApiServer will
// result in compilation errors.
type UnsafeOvaTaskApiServer interface {
	mustEmbedUnimplementedOvaTaskApiServer()
}

func RegisterOvaTaskApiServer(s grpc.ServiceRegistrar, srv OvaTaskApiServer) {
	s.RegisterService(&OvaTaskApi_ServiceDesc, srv)
}

func _OvaTaskApi_CreateTaskV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaTaskApiServer).CreateTaskV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.task.api.OvaTaskApi/CreateTaskV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaTaskApiServer).CreateTaskV1(ctx, req.(*CreateTaskV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaTaskApi_DescribeTaskV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTaskV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaTaskApiServer).DescribeTaskV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.task.api.OvaTaskApi/DescribeTaskV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaTaskApiServer).DescribeTaskV1(ctx, req.(*DescribeTaskV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaTaskApi_ListTasksV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaTaskApiServer).ListTasksV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.task.api.OvaTaskApi/ListTasksV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaTaskApiServer).ListTasksV1(ctx, req.(*ListTasksV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaTaskApi_RemoveTasksV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTaskV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaTaskApiServer).RemoveTasksV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.task.api.OvaTaskApi/RemoveTasksV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaTaskApiServer).RemoveTasksV1(ctx, req.(*RemoveTaskV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OvaTaskApi_ServiceDesc is the grpc.ServiceDesc for OvaTaskApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OvaTaskApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.task.api.OvaTaskApi",
	HandlerType: (*OvaTaskApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTaskV1",
			Handler:    _OvaTaskApi_CreateTaskV1_Handler,
		},
		{
			MethodName: "DescribeTaskV1",
			Handler:    _OvaTaskApi_DescribeTaskV1_Handler,
		},
		{
			MethodName: "ListTasksV1",
			Handler:    _OvaTaskApi_ListTasksV1_Handler,
		},
		{
			MethodName: "RemoveTasksV1",
			Handler:    _OvaTaskApi_RemoveTasksV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ova-task-api/ova-task-api.proto",
}