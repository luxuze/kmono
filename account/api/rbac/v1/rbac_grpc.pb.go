// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// V1Client is the client API for V1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type V1Client interface {
	// Login Account Login
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// Me Get User Info
	Me(ctx context.Context, in *MeRequest, opts ...grpc.CallOption) (*MeReply, error)
}

type v1Client struct {
	cc grpc.ClientConnInterface
}

func NewV1Client(cc grpc.ClientConnInterface) V1Client {
	return &v1Client{cc}
}

func (c *v1Client) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/api.rbac.V1/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1Client) Me(ctx context.Context, in *MeRequest, opts ...grpc.CallOption) (*MeReply, error) {
	out := new(MeReply)
	err := c.cc.Invoke(ctx, "/api.rbac.V1/Me", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V1Server is the server API for V1 service.
// All implementations must embed UnimplementedV1Server
// for forward compatibility
type V1Server interface {
	// Login Account Login
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// Me Get User Info
	Me(context.Context, *MeRequest) (*MeReply, error)
	mustEmbedUnimplementedV1Server()
}

// UnimplementedV1Server must be embedded to have forward compatible implementations.
type UnimplementedV1Server struct {
}

func (UnimplementedV1Server) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedV1Server) Me(context.Context, *MeRequest) (*MeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Me not implemented")
}
func (UnimplementedV1Server) mustEmbedUnimplementedV1Server() {}

// UnsafeV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to V1Server will
// result in compilation errors.
type UnsafeV1Server interface {
	mustEmbedUnimplementedV1Server()
}

func RegisterV1Server(s grpc.ServiceRegistrar, srv V1Server) {
	s.RegisterService(&V1_ServiceDesc, srv)
}

func _V1_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.rbac.V1/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1Server).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1_Me_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1Server).Me(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.rbac.V1/Me",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1Server).Me(ctx, req.(*MeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// V1_ServiceDesc is the grpc.ServiceDesc for V1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var V1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rbac.V1",
	HandlerType: (*V1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _V1_Login_Handler,
		},
		{
			MethodName: "Me",
			Handler:    _V1_Me_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/rbac/v1/rbac.proto",
}
