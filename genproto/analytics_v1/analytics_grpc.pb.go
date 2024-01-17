// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: couchbase/analytics/v1/analytics.proto

package analytics_v1

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

// AnalyticsServiceClient is the client API for AnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsServiceClient interface {
	AnalyticsQuery(ctx context.Context, in *AnalyticsQueryRequest, opts ...grpc.CallOption) (AnalyticsService_AnalyticsQueryClient, error)
}

type analyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsServiceClient(cc grpc.ClientConnInterface) AnalyticsServiceClient {
	return &analyticsServiceClient{cc}
}

func (c *analyticsServiceClient) AnalyticsQuery(ctx context.Context, in *AnalyticsQueryRequest, opts ...grpc.CallOption) (AnalyticsService_AnalyticsQueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &AnalyticsService_ServiceDesc.Streams[0], "/couchbase.analytics.v1.AnalyticsService/AnalyticsQuery", opts...)
	if err != nil {
		return nil, err
	}
	x := &analyticsServiceAnalyticsQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AnalyticsService_AnalyticsQueryClient interface {
	Recv() (*AnalyticsQueryResponse, error)
	grpc.ClientStream
}

type analyticsServiceAnalyticsQueryClient struct {
	grpc.ClientStream
}

func (x *analyticsServiceAnalyticsQueryClient) Recv() (*AnalyticsQueryResponse, error) {
	m := new(AnalyticsQueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnalyticsServiceServer is the server API for AnalyticsService service.
// All implementations must embed UnimplementedAnalyticsServiceServer
// for forward compatibility
type AnalyticsServiceServer interface {
	AnalyticsQuery(*AnalyticsQueryRequest, AnalyticsService_AnalyticsQueryServer) error
	mustEmbedUnimplementedAnalyticsServiceServer()
}

// UnimplementedAnalyticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServiceServer struct {
}

func (UnimplementedAnalyticsServiceServer) AnalyticsQuery(*AnalyticsQueryRequest, AnalyticsService_AnalyticsQueryServer) error {
	return status.Errorf(codes.Unimplemented, "method AnalyticsQuery not implemented")
}
func (UnimplementedAnalyticsServiceServer) mustEmbedUnimplementedAnalyticsServiceServer() {}

// UnsafeAnalyticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsServiceServer will
// result in compilation errors.
type UnsafeAnalyticsServiceServer interface {
	mustEmbedUnimplementedAnalyticsServiceServer()
}

func RegisterAnalyticsServiceServer(s grpc.ServiceRegistrar, srv AnalyticsServiceServer) {
	s.RegisterService(&AnalyticsService_ServiceDesc, srv)
}

func _AnalyticsService_AnalyticsQuery_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AnalyticsQueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnalyticsServiceServer).AnalyticsQuery(m, &analyticsServiceAnalyticsQueryServer{stream})
}

type AnalyticsService_AnalyticsQueryServer interface {
	Send(*AnalyticsQueryResponse) error
	grpc.ServerStream
}

type analyticsServiceAnalyticsQueryServer struct {
	grpc.ServerStream
}

func (x *analyticsServiceAnalyticsQueryServer) Send(m *AnalyticsQueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// AnalyticsService_ServiceDesc is the grpc.ServiceDesc for AnalyticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "couchbase.analytics.v1.AnalyticsService",
	HandlerType: (*AnalyticsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AnalyticsQuery",
			Handler:       _AnalyticsService_AnalyticsQuery_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "couchbase/analytics/v1/analytics.proto",
}
