// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: EventService.proto

package specification

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

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	Prices(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*SymbolPrices, error)
	SymbolPrice(ctx context.Context, in *SymbolPriceRequest, opts ...grpc.CallOption) (*SymbolPrices, error)
	TickerPrices(ctx context.Context, in *DurationSeconds, opts ...grpc.CallOption) (EventService_TickerPricesClient, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Prices(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*SymbolPrices, error) {
	out := new(SymbolPrices)
	err := c.cc.Invoke(ctx, "/event.EventService/SymbolPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) SymbolPrice(ctx context.Context, in *SymbolPriceRequest, opts ...grpc.CallOption) (*SymbolPrices, error) {
	out := new(SymbolPrices)
	err := c.cc.Invoke(ctx, "/event.EventService/SymbolPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) TickerPrices(ctx context.Context, in *DurationSeconds, opts ...grpc.CallOption) (EventService_TickerPricesClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[0], "/event.EventService/TickerPrices", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceTickerPricesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventService_TickerPricesClient interface {
	Recv() (*SymbolPrices, error)
	grpc.ClientStream
}

type eventServiceTickerPricesClient struct {
	grpc.ClientStream
}

func (x *eventServiceTickerPricesClient) Recv() (*SymbolPrices, error) {
	m := new(SymbolPrices)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	Prices(context.Context, *EmptyRequest) (*SymbolPrices, error)
	SymbolPrice(context.Context, *SymbolPriceRequest) (*SymbolPrices, error)
	TickerPrices(*DurationSeconds, EventService_TickerPricesServer) error
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) Prices(context.Context, *EmptyRequest) (*SymbolPrices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SymbolPrices not implemented")
}
func (UnimplementedEventServiceServer) SymbolPrice(context.Context, *SymbolPriceRequest) (*SymbolPrices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SymbolPrice not implemented")
}
func (UnimplementedEventServiceServer) TickerPrices(*DurationSeconds, EventService_TickerPricesServer) error {
	return status.Errorf(codes.Unimplemented, "method TickerPrices not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_Prices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Prices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventService/SymbolPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Prices(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_SymbolPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymbolPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SymbolPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventService/SymbolPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SymbolPrice(ctx, req.(*SymbolPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_TickerPrices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DurationSeconds)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventServiceServer).TickerPrices(m, &eventServiceTickerPricesServer{stream})
}

type EventService_TickerPricesServer interface {
	Send(*SymbolPrices) error
	grpc.ServerStream
}

type eventServiceTickerPricesServer struct {
	grpc.ServerStream
}

func (x *eventServiceTickerPricesServer) Send(m *SymbolPrices) error {
	return x.ServerStream.SendMsg(m)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SymbolPrices",
			Handler:    _EventService_Prices_Handler,
		},
		{
			MethodName: "SymbolPrice",
			Handler:    _EventService_SymbolPrice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TickerPrices",
			Handler:       _EventService_TickerPrices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "EventService.proto",
}
