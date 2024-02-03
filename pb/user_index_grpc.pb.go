// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: user_index.proto

package tinder_matching

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

const (
	UserService_CreateUserAndListMatches_FullMethodName = "/user.UserService/CreateUserAndListMatches"
	UserService_DeleteUserById_FullMethodName           = "/user.UserService/DeleteUserById"
	UserService_ListMatchesById_FullMethodName          = "/user.UserService/ListMatchesById"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUserAndListMatches(ctx context.Context, in *CreateUserAndListMatchesRequest, opts ...grpc.CallOption) (*CreateUserAndListMatchesResponse, error)
	DeleteUserById(ctx context.Context, in *DeleteUserByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListMatchesById(ctx context.Context, in *ListMatchesByIdRequest, opts ...grpc.CallOption) (*ListMatchesByIdResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUserAndListMatches(ctx context.Context, in *CreateUserAndListMatchesRequest, opts ...grpc.CallOption) (*CreateUserAndListMatchesResponse, error) {
	out := new(CreateUserAndListMatchesResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUserAndListMatches_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserById(ctx context.Context, in *DeleteUserByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListMatchesById(ctx context.Context, in *ListMatchesByIdRequest, opts ...grpc.CallOption) (*ListMatchesByIdResponse, error) {
	out := new(ListMatchesByIdResponse)
	err := c.cc.Invoke(ctx, UserService_ListMatchesById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUserAndListMatches(context.Context, *CreateUserAndListMatchesRequest) (*CreateUserAndListMatchesResponse, error)
	DeleteUserById(context.Context, *DeleteUserByIdRequest) (*emptypb.Empty, error)
	ListMatchesById(context.Context, *ListMatchesByIdRequest) (*ListMatchesByIdResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUserAndListMatches(context.Context, *CreateUserAndListMatchesRequest) (*CreateUserAndListMatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserAndListMatches not implemented")
}
func (UnimplementedUserServiceServer) DeleteUserById(context.Context, *DeleteUserByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserById not implemented")
}
func (UnimplementedUserServiceServer) ListMatchesById(context.Context, *ListMatchesByIdRequest) (*ListMatchesByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMatchesById not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUserAndListMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserAndListMatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUserAndListMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUserAndListMatches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUserAndListMatches(ctx, req.(*CreateUserAndListMatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUserById(ctx, req.(*DeleteUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListMatchesById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMatchesByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListMatchesById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListMatchesById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListMatchesById(ctx, req.(*ListMatchesByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserAndListMatches",
			Handler:    _UserService_CreateUserAndListMatches_Handler,
		},
		{
			MethodName: "DeleteUserById",
			Handler:    _UserService_DeleteUserById_Handler,
		},
		{
			MethodName: "ListMatchesById",
			Handler:    _UserService_ListMatchesById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_index.proto",
}