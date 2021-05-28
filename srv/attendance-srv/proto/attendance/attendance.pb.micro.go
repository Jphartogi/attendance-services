// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/attendance/attendance.proto

package attendance

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Attendance service

func NewAttendanceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Attendance service

type AttendanceService interface {
	Register(ctx context.Context, in *Employee, opts ...client.CallOption) (*RegisterEmployeeResponse, error)
}

type attendanceService struct {
	c    client.Client
	name string
}

func NewAttendanceService(name string, c client.Client) AttendanceService {
	return &attendanceService{
		c:    c,
		name: name,
	}
}

func (c *attendanceService) Register(ctx context.Context, in *Employee, opts ...client.CallOption) (*RegisterEmployeeResponse, error) {
	req := c.c.NewRequest(c.name, "Attendance.Register", in)
	out := new(RegisterEmployeeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Attendance service

type AttendanceHandler interface {
	Register(context.Context, *Employee, *RegisterEmployeeResponse) error
}

func RegisterAttendanceHandler(s server.Server, hdlr AttendanceHandler, opts ...server.HandlerOption) error {
	type attendance interface {
		Register(ctx context.Context, in *Employee, out *RegisterEmployeeResponse) error
	}
	type Attendance struct {
		attendance
	}
	h := &attendanceHandler{hdlr}
	return s.Handle(s.NewHandler(&Attendance{h}, opts...))
}

type attendanceHandler struct {
	AttendanceHandler
}

func (h *attendanceHandler) Register(ctx context.Context, in *Employee, out *RegisterEmployeeResponse) error {
	return h.AttendanceHandler.Register(ctx, in, out)
}
