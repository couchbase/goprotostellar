// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: couchbase/admin/query/v1/query.proto

package admin_query_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QueryAdminServiceClient is the client API for QueryAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryAdminServiceClient interface {
	GetAllIndexes(ctx context.Context, in *GetAllIndexesRequest, opts ...grpc.CallOption) (*GetAllIndexesResponse, error)
	CreatePrimaryIndex(ctx context.Context, in *CreatePrimaryIndexRequest, opts ...grpc.CallOption) (*CreatePrimaryIndexResponse, error)
	CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*CreateIndexResponse, error)
	DropPrimaryIndex(ctx context.Context, in *DropPrimaryIndexRequest, opts ...grpc.CallOption) (*DropPrimaryIndexResponse, error)
	DropIndex(ctx context.Context, in *DropIndexRequest, opts ...grpc.CallOption) (*DropIndexResponse, error)
	BuildDeferredIndexes(ctx context.Context, in *BuildDeferredIndexesRequest, opts ...grpc.CallOption) (*BuildDeferredIndexesResponse, error)
	WaitForIndexOnline(ctx context.Context, in *WaitForIndexOnlineRequest, opts ...grpc.CallOption) (*WaitForIndexOnlineResponse, error)
}

type queryAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryAdminServiceClient(cc grpc.ClientConnInterface) QueryAdminServiceClient {
	return &queryAdminServiceClient{cc}
}

func (c *queryAdminServiceClient) GetAllIndexes(ctx context.Context, in *GetAllIndexesRequest, opts ...grpc.CallOption) (*GetAllIndexesResponse, error) {
	out := new(GetAllIndexesResponse)
	err := c.cc.Invoke(ctx, "/couchbase.admin.query.v1.QueryAdminService/GetAllIndexes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryAdminServiceClient) CreatePrimaryIndex(ctx context.Context, in *CreatePrimaryIndexRequest, opts ...grpc.CallOption) (*CreatePrimaryIndexResponse, error) {
	out := new(CreatePrimaryIndexResponse)
	err := c.cc.Invoke(ctx, "/couchbase.admin.query.v1.QueryAdminService/CreatePrimaryIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryAdminServiceClient) CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*CreateIndexResponse, error) {
	out := new(CreateIndexResponse)
	err := c.cc.Invoke(ctx, "/couchbase.admin.query.v1.QueryAdminService/CreateIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryAdminServiceClient) DropPrimaryIndex(ctx context.Context, in *DropPrimaryIndexRequest, opts ...grpc.CallOption) (*DropPrimaryIndexResponse, error) {
	out := new(DropPrimaryIndexResponse)
	err := c.cc.Invoke(ctx, "/couchbase.admin.query.v1.QueryAdminService/DropPrimaryIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryAdminServiceClient) DropIndex(ctx context.Context, in *DropIndexRequest, opts ...grpc.CallOption) (*DropIndexResponse, error) {
	out := new(DropIndexResponse)
	err := c.cc.Invoke(ctx, "/couchbase.admin.query.v1.QueryAdminService/DropIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryAdminServiceClient) BuildDeferredIndexes(ctx context.Context, in *BuildDeferredIndexesRequest, opts ...grpc.CallOption) (*BuildDeferredIndexesResponse, error) {
	out := new(BuildDeferredIndexesResponse)
	err := c.cc.Invoke(ctx, "/couchbase.admin.query.v1.QueryAdminService/BuildDeferredIndexes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryAdminServiceClient) WaitForIndexOnline(ctx context.Context, in *WaitForIndexOnlineRequest, opts ...grpc.CallOption) (*WaitForIndexOnlineResponse, error) {
	out := new(WaitForIndexOnlineResponse)
	err := c.cc.Invoke(ctx, "/couchbase.admin.query.v1.QueryAdminService/WaitForIndexOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryAdminServiceServer is the server API for QueryAdminService service.
// All implementations must embed UnimplementedQueryAdminServiceServer
// for forward compatibility
type QueryAdminServiceServer interface {
	GetAllIndexes(context.Context, *GetAllIndexesRequest) (*GetAllIndexesResponse, error)
	CreatePrimaryIndex(context.Context, *CreatePrimaryIndexRequest) (*CreatePrimaryIndexResponse, error)
	CreateIndex(context.Context, *CreateIndexRequest) (*CreateIndexResponse, error)
	DropPrimaryIndex(context.Context, *DropPrimaryIndexRequest) (*DropPrimaryIndexResponse, error)
	DropIndex(context.Context, *DropIndexRequest) (*DropIndexResponse, error)
	BuildDeferredIndexes(context.Context, *BuildDeferredIndexesRequest) (*BuildDeferredIndexesResponse, error)
	WaitForIndexOnline(context.Context, *WaitForIndexOnlineRequest) (*WaitForIndexOnlineResponse, error)
	mustEmbedUnimplementedQueryAdminServiceServer()
}

// UnimplementedQueryAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueryAdminServiceServer struct {
}

func (UnimplementedQueryAdminServiceServer) GetAllIndexes(context.Context, *GetAllIndexesRequest) (*GetAllIndexesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllIndexes not implemented")
}
func (UnimplementedQueryAdminServiceServer) CreatePrimaryIndex(context.Context, *CreatePrimaryIndexRequest) (*CreatePrimaryIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrimaryIndex not implemented")
}
func (UnimplementedQueryAdminServiceServer) CreateIndex(context.Context, *CreateIndexRequest) (*CreateIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIndex not implemented")
}
func (UnimplementedQueryAdminServiceServer) DropPrimaryIndex(context.Context, *DropPrimaryIndexRequest) (*DropPrimaryIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropPrimaryIndex not implemented")
}
func (UnimplementedQueryAdminServiceServer) DropIndex(context.Context, *DropIndexRequest) (*DropIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropIndex not implemented")
}
func (UnimplementedQueryAdminServiceServer) BuildDeferredIndexes(context.Context, *BuildDeferredIndexesRequest) (*BuildDeferredIndexesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildDeferredIndexes not implemented")
}
func (UnimplementedQueryAdminServiceServer) WaitForIndexOnline(context.Context, *WaitForIndexOnlineRequest) (*WaitForIndexOnlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitForIndexOnline not implemented")
}
func (UnimplementedQueryAdminServiceServer) mustEmbedUnimplementedQueryAdminServiceServer() {}

// UnsafeQueryAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryAdminServiceServer will
// result in compilation errors.
type UnsafeQueryAdminServiceServer interface {
	mustEmbedUnimplementedQueryAdminServiceServer()
}

func RegisterQueryAdminServiceServer(s grpc.ServiceRegistrar, srv QueryAdminServiceServer) {
	s.RegisterService(&QueryAdminService_ServiceDesc, srv)
}

func _QueryAdminService_GetAllIndexes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllIndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAdminServiceServer).GetAllIndexes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/couchbase.admin.query.v1.QueryAdminService/GetAllIndexes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAdminServiceServer).GetAllIndexes(ctx, req.(*GetAllIndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryAdminService_CreatePrimaryIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrimaryIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAdminServiceServer).CreatePrimaryIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/couchbase.admin.query.v1.QueryAdminService/CreatePrimaryIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAdminServiceServer).CreatePrimaryIndex(ctx, req.(*CreatePrimaryIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryAdminService_CreateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAdminServiceServer).CreateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/couchbase.admin.query.v1.QueryAdminService/CreateIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAdminServiceServer).CreateIndex(ctx, req.(*CreateIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryAdminService_DropPrimaryIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropPrimaryIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAdminServiceServer).DropPrimaryIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/couchbase.admin.query.v1.QueryAdminService/DropPrimaryIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAdminServiceServer).DropPrimaryIndex(ctx, req.(*DropPrimaryIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryAdminService_DropIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAdminServiceServer).DropIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/couchbase.admin.query.v1.QueryAdminService/DropIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAdminServiceServer).DropIndex(ctx, req.(*DropIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryAdminService_BuildDeferredIndexes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildDeferredIndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAdminServiceServer).BuildDeferredIndexes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/couchbase.admin.query.v1.QueryAdminService/BuildDeferredIndexes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAdminServiceServer).BuildDeferredIndexes(ctx, req.(*BuildDeferredIndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryAdminService_WaitForIndexOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitForIndexOnlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryAdminServiceServer).WaitForIndexOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/couchbase.admin.query.v1.QueryAdminService/WaitForIndexOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryAdminServiceServer).WaitForIndexOnline(ctx, req.(*WaitForIndexOnlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryAdminService_ServiceDesc is the grpc.ServiceDesc for QueryAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "couchbase.admin.query.v1.QueryAdminService",
	HandlerType: (*QueryAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllIndexes",
			Handler:    _QueryAdminService_GetAllIndexes_Handler,
		},
		{
			MethodName: "CreatePrimaryIndex",
			Handler:    _QueryAdminService_CreatePrimaryIndex_Handler,
		},
		{
			MethodName: "CreateIndex",
			Handler:    _QueryAdminService_CreateIndex_Handler,
		},
		{
			MethodName: "DropPrimaryIndex",
			Handler:    _QueryAdminService_DropPrimaryIndex_Handler,
		},
		{
			MethodName: "DropIndex",
			Handler:    _QueryAdminService_DropIndex_Handler,
		},
		{
			MethodName: "BuildDeferredIndexes",
			Handler:    _QueryAdminService_BuildDeferredIndexes_Handler,
		},
		{
			MethodName: "WaitForIndexOnline",
			Handler:    _QueryAdminService_WaitForIndexOnline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "couchbase/admin/query/v1/query.proto",
}
