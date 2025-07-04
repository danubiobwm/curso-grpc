// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: src/proto/department-service.proto

package department

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DepartmentService_ListPerson_FullMethodName = "/department.DepartmentService/ListPerson"
)

// DepartmentServiceClient is the client API for DepartmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepartmentServiceClient interface {
	ListPerson(ctx context.Context, in *ListPersonRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListPersonResponse], error)
}

type departmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentServiceClient(cc grpc.ClientConnInterface) DepartmentServiceClient {
	return &departmentServiceClient{cc}
}

func (c *departmentServiceClient) ListPerson(ctx context.Context, in *ListPersonRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListPersonResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DepartmentService_ServiceDesc.Streams[0], DepartmentService_ListPerson_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListPersonRequest, ListPersonResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DepartmentService_ListPersonClient = grpc.ServerStreamingClient[ListPersonResponse]

// DepartmentServiceServer is the server API for DepartmentService service.
// All implementations must embed UnimplementedDepartmentServiceServer
// for forward compatibility.
type DepartmentServiceServer interface {
	ListPerson(*ListPersonRequest, grpc.ServerStreamingServer[ListPersonResponse]) error
	mustEmbedUnimplementedDepartmentServiceServer()
}

// UnimplementedDepartmentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDepartmentServiceServer struct{}

func (UnimplementedDepartmentServiceServer) ListPerson(*ListPersonRequest, grpc.ServerStreamingServer[ListPersonResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ListPerson not implemented")
}
func (UnimplementedDepartmentServiceServer) mustEmbedUnimplementedDepartmentServiceServer() {}
func (UnimplementedDepartmentServiceServer) testEmbeddedByValue()                           {}

// UnsafeDepartmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepartmentServiceServer will
// result in compilation errors.
type UnsafeDepartmentServiceServer interface {
	mustEmbedUnimplementedDepartmentServiceServer()
}

func RegisterDepartmentServiceServer(s grpc.ServiceRegistrar, srv DepartmentServiceServer) {
	// If the following call pancis, it indicates UnimplementedDepartmentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DepartmentService_ServiceDesc, srv)
}

func _DepartmentService_ListPerson_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPersonRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DepartmentServiceServer).ListPerson(m, &grpc.GenericServerStream[ListPersonRequest, ListPersonResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DepartmentService_ListPersonServer = grpc.ServerStreamingServer[ListPersonResponse]

// DepartmentService_ServiceDesc is the grpc.ServiceDesc for DepartmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepartmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "department.DepartmentService",
	HandlerType: (*DepartmentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPerson",
			Handler:       _DepartmentService_ListPerson_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/proto/department-service.proto",
}
