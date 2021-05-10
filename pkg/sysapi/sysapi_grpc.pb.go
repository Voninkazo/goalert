// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sysapi

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

// SysAPIClient is the client API for SysAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysAPIClient interface {
	AuthSubjects(ctx context.Context, in *AuthSubjectsRequest, opts ...grpc.CallOption) (SysAPI_AuthSubjectsClient, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type sysAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSysAPIClient(cc grpc.ClientConnInterface) SysAPIClient {
	return &sysAPIClient{cc}
}

func (c *sysAPIClient) AuthSubjects(ctx context.Context, in *AuthSubjectsRequest, opts ...grpc.CallOption) (SysAPI_AuthSubjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SysAPI_ServiceDesc.Streams[0], "/goalert.v1.SysAPI/AuthSubjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &sysAPIAuthSubjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SysAPI_AuthSubjectsClient interface {
	Recv() (*AuthSubject, error)
	grpc.ClientStream
}

type sysAPIAuthSubjectsClient struct {
	grpc.ClientStream
}

func (x *sysAPIAuthSubjectsClient) Recv() (*AuthSubject, error) {
	m := new(AuthSubject)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sysAPIClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/goalert.v1.SysAPI/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysAPIServer is the server API for SysAPI service.
// All implementations must embed UnimplementedSysAPIServer
// for forward compatibility
type SysAPIServer interface {
	AuthSubjects(*AuthSubjectsRequest, SysAPI_AuthSubjectsServer) error
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedSysAPIServer()
}

// UnimplementedSysAPIServer must be embedded to have forward compatible implementations.
type UnimplementedSysAPIServer struct {
}

func (UnimplementedSysAPIServer) AuthSubjects(*AuthSubjectsRequest, SysAPI_AuthSubjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method AuthSubjects not implemented")
}
func (UnimplementedSysAPIServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedSysAPIServer) mustEmbedUnimplementedSysAPIServer() {}

// UnsafeSysAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysAPIServer will
// result in compilation errors.
type UnsafeSysAPIServer interface {
	mustEmbedUnimplementedSysAPIServer()
}

func RegisterSysAPIServer(s grpc.ServiceRegistrar, srv SysAPIServer) {
	s.RegisterService(&SysAPI_ServiceDesc, srv)
}

func _SysAPI_AuthSubjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AuthSubjectsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SysAPIServer).AuthSubjects(m, &sysAPIAuthSubjectsServer{stream})
}

type SysAPI_AuthSubjectsServer interface {
	Send(*AuthSubject) error
	grpc.ServerStream
}

type sysAPIAuthSubjectsServer struct {
	grpc.ServerStream
}

func (x *sysAPIAuthSubjectsServer) Send(m *AuthSubject) error {
	return x.ServerStream.SendMsg(m)
}

func _SysAPI_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysAPIServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalert.v1.SysAPI/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysAPIServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SysAPI_ServiceDesc is the grpc.ServiceDesc for SysAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goalert.v1.SysAPI",
	HandlerType: (*SysAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteUser",
			Handler:    _SysAPI_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AuthSubjects",
			Handler:       _SysAPI_AuthSubjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/sysapi/sysapi.proto",
}
