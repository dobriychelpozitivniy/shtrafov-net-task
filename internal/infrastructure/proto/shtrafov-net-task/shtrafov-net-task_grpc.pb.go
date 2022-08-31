// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/proto/shtrafov-net-task/shtrafov-net-task.proto

package grpcService

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

// INNServiceClient is the client API for INNService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type INNServiceClient interface {
	GetOrganisationInfo(ctx context.Context, in *INNRequest, opts ...grpc.CallOption) (*INNResponse, error)
}

type iNNServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewINNServiceClient(cc grpc.ClientConnInterface) INNServiceClient {
	return &iNNServiceClient{cc}
}

func (c *iNNServiceClient) GetOrganisationInfo(ctx context.Context, in *INNRequest, opts ...grpc.CallOption) (*INNResponse, error) {
	out := new(INNResponse)
	err := c.cc.Invoke(ctx, "/INNService/GetOrganisationInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// INNServiceServer is the server API for INNService service.
// All implementations must embed UnimplementedINNServiceServer
// for forward compatibility
type INNServiceServer interface {
	GetOrganisationInfo(context.Context, *INNRequest) (*INNResponse, error)
	mustEmbedUnimplementedINNServiceServer()
}

// UnimplementedINNServiceServer must be embedded to have forward compatible implementations.
type UnimplementedINNServiceServer struct {
}

func (UnimplementedINNServiceServer) GetOrganisationInfo(context.Context, *INNRequest) (*INNResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganisationInfo not implemented")
}
func (UnimplementedINNServiceServer) mustEmbedUnimplementedINNServiceServer() {}

// UnsafeINNServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to INNServiceServer will
// result in compilation errors.
type UnsafeINNServiceServer interface {
	mustEmbedUnimplementedINNServiceServer()
}

func RegisterINNServiceServer(s grpc.ServiceRegistrar, srv INNServiceServer) {
	s.RegisterService(&INNService_ServiceDesc, srv)
}

func _INNService_GetOrganisationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(INNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INNServiceServer).GetOrganisationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/INNService/GetOrganisationInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INNServiceServer).GetOrganisationInfo(ctx, req.(*INNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// INNService_ServiceDesc is the grpc.ServiceDesc for INNService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var INNService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "INNService",
	HandlerType: (*INNServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrganisationInfo",
			Handler:    _INNService_GetOrganisationInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/shtrafov-net-task/shtrafov-net-task.proto",
}