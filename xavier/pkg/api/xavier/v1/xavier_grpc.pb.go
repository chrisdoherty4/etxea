// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package xavierv1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BindingServiceClient is the client API for BindingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BindingServiceClient interface {
}

type bindingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBindingServiceClient(cc grpc.ClientConnInterface) BindingServiceClient {
	return &bindingServiceClient{cc}
}

// BindingServiceServer is the server API for BindingService service.
// All implementations should embed UnimplementedBindingServiceServer
// for forward compatibility
type BindingServiceServer interface {
}

// UnimplementedBindingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBindingServiceServer struct {
}

// UnsafeBindingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BindingServiceServer will
// result in compilation errors.
type UnsafeBindingServiceServer interface {
	mustEmbedUnimplementedBindingServiceServer()
}

func RegisterBindingServiceServer(s grpc.ServiceRegistrar, srv BindingServiceServer) {
	s.RegisterService(&BindingService_ServiceDesc, srv)
}

// BindingService_ServiceDesc is the grpc.ServiceDesc for BindingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BindingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xavier.v1.BindingService",
	HandlerType: (*BindingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "xavier/v1/xavier.proto",
}

// FlagServiceClient is the client API for FlagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlagServiceClient interface {
}

type flagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlagServiceClient(cc grpc.ClientConnInterface) FlagServiceClient {
	return &flagServiceClient{cc}
}

// FlagServiceServer is the server API for FlagService service.
// All implementations should embed UnimplementedFlagServiceServer
// for forward compatibility
type FlagServiceServer interface {
}

// UnimplementedFlagServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFlagServiceServer struct {
}

// UnsafeFlagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlagServiceServer will
// result in compilation errors.
type UnsafeFlagServiceServer interface {
	mustEmbedUnimplementedFlagServiceServer()
}

func RegisterFlagServiceServer(s grpc.ServiceRegistrar, srv FlagServiceServer) {
	s.RegisterService(&FlagService_ServiceDesc, srv)
}

// FlagService_ServiceDesc is the grpc.ServiceDesc for FlagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xavier.v1.FlagService",
	HandlerType: (*FlagServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "xavier/v1/xavier.proto",
}