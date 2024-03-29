// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: couchbase/routing/v1/routing.proto

package routing_v1

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

// RoutingServiceClient is the client API for RoutingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoutingServiceClient interface {
	WatchRouting(ctx context.Context, in *WatchRoutingRequest, opts ...grpc.CallOption) (RoutingService_WatchRoutingClient, error)
}

type routingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutingServiceClient(cc grpc.ClientConnInterface) RoutingServiceClient {
	return &routingServiceClient{cc}
}

func (c *routingServiceClient) WatchRouting(ctx context.Context, in *WatchRoutingRequest, opts ...grpc.CallOption) (RoutingService_WatchRoutingClient, error) {
	stream, err := c.cc.NewStream(ctx, &RoutingService_ServiceDesc.Streams[0], "/couchbase.routing.v1.RoutingService/WatchRouting", opts...)
	if err != nil {
		return nil, err
	}
	x := &routingServiceWatchRoutingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutingService_WatchRoutingClient interface {
	Recv() (*WatchRoutingResponse, error)
	grpc.ClientStream
}

type routingServiceWatchRoutingClient struct {
	grpc.ClientStream
}

func (x *routingServiceWatchRoutingClient) Recv() (*WatchRoutingResponse, error) {
	m := new(WatchRoutingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoutingServiceServer is the server API for RoutingService service.
// All implementations must embed UnimplementedRoutingServiceServer
// for forward compatibility
type RoutingServiceServer interface {
	WatchRouting(*WatchRoutingRequest, RoutingService_WatchRoutingServer) error
	mustEmbedUnimplementedRoutingServiceServer()
}

// UnimplementedRoutingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoutingServiceServer struct {
}

func (UnimplementedRoutingServiceServer) WatchRouting(*WatchRoutingRequest, RoutingService_WatchRoutingServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchRouting not implemented")
}
func (UnimplementedRoutingServiceServer) mustEmbedUnimplementedRoutingServiceServer() {}

// UnsafeRoutingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutingServiceServer will
// result in compilation errors.
type UnsafeRoutingServiceServer interface {
	mustEmbedUnimplementedRoutingServiceServer()
}

func RegisterRoutingServiceServer(s grpc.ServiceRegistrar, srv RoutingServiceServer) {
	s.RegisterService(&RoutingService_ServiceDesc, srv)
}

func _RoutingService_WatchRouting_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRoutingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutingServiceServer).WatchRouting(m, &routingServiceWatchRoutingServer{stream})
}

type RoutingService_WatchRoutingServer interface {
	Send(*WatchRoutingResponse) error
	grpc.ServerStream
}

type routingServiceWatchRoutingServer struct {
	grpc.ServerStream
}

func (x *routingServiceWatchRoutingServer) Send(m *WatchRoutingResponse) error {
	return x.ServerStream.SendMsg(m)
}

// RoutingService_ServiceDesc is the grpc.ServiceDesc for RoutingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoutingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "couchbase.routing.v1.RoutingService",
	HandlerType: (*RoutingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchRouting",
			Handler:       _RoutingService_WatchRouting_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "couchbase/routing/v1/routing.proto",
}
