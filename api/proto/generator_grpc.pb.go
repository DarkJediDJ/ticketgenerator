// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ticketgenerator

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

// TicketGeneratorClient is the client API for TicketGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketGeneratorClient interface {
	GetTicket(ctx context.Context, in *TicketRequset, opts ...grpc.CallOption) (*LinkReply, error)
}

type ticketGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketGeneratorClient(cc grpc.ClientConnInterface) TicketGeneratorClient {
	return &ticketGeneratorClient{cc}
}

func (c *ticketGeneratorClient) GetTicket(ctx context.Context, in *TicketRequset, opts ...grpc.CallOption) (*LinkReply, error) {
	out := new(LinkReply)
	err := c.cc.Invoke(ctx, "/transactions.TicketGenerator/GetTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketGeneratorServer is the server API for TicketGenerator service.
// All implementations must embed UnimplementedTicketGeneratorServer
// for forward compatibility
type TicketGeneratorServer interface {
	GetTicket(context.Context, *TicketRequset) (*LinkReply, error)
	mustEmbedUnimplementedTicketGeneratorServer()
}

// UnimplementedTicketGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedTicketGeneratorServer struct {
}

func (UnimplementedTicketGeneratorServer) GetTicket(context.Context, *TicketRequset) (*LinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedTicketGeneratorServer) mustEmbedUnimplementedTicketGeneratorServer() {}

// UnsafeTicketGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketGeneratorServer will
// result in compilation errors.
type UnsafeTicketGeneratorServer interface {
	mustEmbedUnimplementedTicketGeneratorServer()
}

func RegisterTicketGeneratorServer(s grpc.ServiceRegistrar, srv TicketGeneratorServer) {
	s.RegisterService(&TicketGenerator_ServiceDesc, srv)
}

func _TicketGenerator_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketGeneratorServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TicketGenerator/GetTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketGeneratorServer).GetTicket(ctx, req.(*TicketRequset))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketGenerator_ServiceDesc is the grpc.ServiceDesc for TicketGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transactions.TicketGenerator",
	HandlerType: (*TicketGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTicket",
			Handler:    _TicketGenerator_GetTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/generator.proto",
}
