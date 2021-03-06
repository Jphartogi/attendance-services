// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package attendance

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

// AttendanceClient is the client API for Attendance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttendanceClient interface {
	Register(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*RegisterEmployeeResponse, error)
}

type attendanceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttendanceClient(cc grpc.ClientConnInterface) AttendanceClient {
	return &attendanceClient{cc}
}

func (c *attendanceClient) Register(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*RegisterEmployeeResponse, error) {
	out := new(RegisterEmployeeResponse)
	err := c.cc.Invoke(ctx, "/attendance.service.Attendance/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttendanceServer is the server API for Attendance service.
// All implementations must embed UnimplementedAttendanceServer
// for forward compatibility
type AttendanceServer interface {
	Register(context.Context, *Employee) (*RegisterEmployeeResponse, error)
	mustEmbedUnimplementedAttendanceServer()
}

// UnimplementedAttendanceServer must be embedded to have forward compatible implementations.
type UnimplementedAttendanceServer struct {
}

func (UnimplementedAttendanceServer) Register(context.Context, *Employee) (*RegisterEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAttendanceServer) mustEmbedUnimplementedAttendanceServer() {}

// UnsafeAttendanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttendanceServer will
// result in compilation errors.
type UnsafeAttendanceServer interface {
	mustEmbedUnimplementedAttendanceServer()
}

func RegisterAttendanceServer(s grpc.ServiceRegistrar, srv AttendanceServer) {
	s.RegisterService(&Attendance_ServiceDesc, srv)
}

func _Attendance_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attendance.service.Attendance/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServer).Register(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

// Attendance_ServiceDesc is the grpc.ServiceDesc for Attendance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Attendance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attendance.service.Attendance",
	HandlerType: (*AttendanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Attendance_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/attendance/attendance.proto",
}
