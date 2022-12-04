// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protocol/student.proto

package protocol

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

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	GetStudentByID(ctx context.Context, in *SearchByID, opts ...grpc.CallOption) (*Student, error)
	GetStudentsByName(ctx context.Context, in *SearchByName, opts ...grpc.CallOption) (StudentService_GetStudentsByNameClient, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) GetStudentByID(ctx context.Context, in *SearchByID, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/StudentService/GetStudentByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudentsByName(ctx context.Context, in *SearchByName, opts ...grpc.CallOption) (StudentService_GetStudentsByNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &StudentService_ServiceDesc.Streams[0], "/StudentService/GetStudentsByName", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentServiceGetStudentsByNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StudentService_GetStudentsByNameClient interface {
	Recv() (*Student, error)
	grpc.ClientStream
}

type studentServiceGetStudentsByNameClient struct {
	grpc.ClientStream
}

func (x *studentServiceGetStudentsByNameClient) Recv() (*Student, error) {
	m := new(Student)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations should embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	GetStudentByID(context.Context, *SearchByID) (*Student, error)
	GetStudentsByName(*SearchByName, StudentService_GetStudentsByNameServer) error
}

// UnimplementedStudentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) GetStudentByID(context.Context, *SearchByID) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentByID not implemented")
}
func (UnimplementedStudentServiceServer) GetStudentsByName(*SearchByName, StudentService_GetStudentsByNameServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStudentsByName not implemented")
}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_GetStudentByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudentByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StudentService/GetStudentByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudentByID(ctx, req.(*SearchByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudentsByName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchByName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentServiceServer).GetStudentsByName(m, &studentServiceGetStudentsByNameServer{stream})
}

type StudentService_GetStudentsByNameServer interface {
	Send(*Student) error
	grpc.ServerStream
}

type studentServiceGetStudentsByNameServer struct {
	grpc.ServerStream
}

func (x *studentServiceGetStudentsByNameServer) Send(m *Student) error {
	return x.ServerStream.SendMsg(m)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudentByID",
			Handler:    _StudentService_GetStudentByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStudentsByName",
			Handler:       _StudentService_GetStudentsByName_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol/student.proto",
}
