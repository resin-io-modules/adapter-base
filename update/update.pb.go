// Code generated by protoc-gen-go. DO NOT EDIT.
// source: update.proto

/*
Package update is a generated protocol buffer package.

It is generated from these files:
	update.proto

It has these top-level messages:
	StartRequest
	StartResponse
	StatusRequest
	StatusResponse
*/
package update

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StatusResponse_State int32

const (
	StatusResponse__         StatusResponse_State = 0
	StatusResponse_STARTED   StatusResponse_State = 1
	StatusResponse_COMPLETED StatusResponse_State = 2
	StatusResponse_CANCELLED StatusResponse_State = 3
	StatusResponse_FAILED    StatusResponse_State = 4
	StatusResponse_TIMED_OUT StatusResponse_State = 5
)

var StatusResponse_State_name = map[int32]string{
	0: "_",
	1: "STARTED",
	2: "COMPLETED",
	3: "CANCELLED",
	4: "FAILED",
	5: "TIMED_OUT",
}
var StatusResponse_State_value = map[string]int32{
	"_":         0,
	"STARTED":   1,
	"COMPLETED": 2,
	"CANCELLED": 3,
	"FAILED":    4,
	"TIMED_OUT": 5,
}

func (x StatusResponse_State) String() string {
	return proto.EnumName(StatusResponse_State_name, int32(x))
}
func (StatusResponse_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type StartRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	Timeout int64  `protobuf:"varint,3,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StartRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *StartRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type StartResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StartResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StatusRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StatusResponse struct {
	Id           string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	StartRequest *StartRequest        `protobuf:"bytes,2,opt,name=startRequest" json:"startRequest,omitempty"`
	State        StatusResponse_State `protobuf:"varint,3,opt,name=state,enum=update.StatusResponse_State" json:"state,omitempty"`
	Progress     int32                `protobuf:"varint,4,opt,name=progress" json:"progress,omitempty"`
	Message      string               `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	Started      int64                `protobuf:"varint,6,opt,name=started" json:"started,omitempty"`
	Completed    int64                `protobuf:"varint,7,opt,name=completed" json:"completed,omitempty"`
	Duration     int64                `protobuf:"varint,8,opt,name=duration" json:"duration,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatusResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StatusResponse) GetStartRequest() *StartRequest {
	if m != nil {
		return m.StartRequest
	}
	return nil
}

func (m *StatusResponse) GetState() StatusResponse_State {
	if m != nil {
		return m.State
	}
	return StatusResponse__
}

func (m *StatusResponse) GetProgress() int32 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func (m *StatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StatusResponse) GetStarted() int64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *StatusResponse) GetCompleted() int64 {
	if m != nil {
		return m.Completed
	}
	return 0
}

func (m *StatusResponse) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterType((*StartRequest)(nil), "update.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "update.StartResponse")
	proto.RegisterType((*StatusRequest)(nil), "update.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "update.StatusResponse")
	proto.RegisterEnum("update.StatusResponse_State", StatusResponse_State_name, StatusResponse_State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Update service

type UpdateClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Cancel(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type updateClient struct {
	cc *grpc.ClientConn
}

func NewUpdateClient(cc *grpc.ClientConn) UpdateClient {
	return &updateClient{cc}
}

func (c *updateClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/update.Update/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/update.Update/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) Cancel(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/update.Update/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Update service

type UpdateServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	Cancel(context.Context, *StatusRequest) (*StatusResponse, error)
}

func RegisterUpdateServer(s *grpc.Server, srv UpdateServer) {
	s.RegisterService(&_Update_serviceDesc, srv)
}

func _Update_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update.Update/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update.Update/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update.Update/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).Cancel(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Update_serviceDesc = grpc.ServiceDesc{
	ServiceName: "update.Update",
	HandlerType: (*UpdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Update_Start_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Update_Status_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Update_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "update.proto",
}

func init() { proto.RegisterFile("update.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6a, 0xdb, 0x40,
	0x14, 0x85, 0x3b, 0x72, 0x24, 0xc7, 0x37, 0x8a, 0x31, 0x97, 0xa4, 0x08, 0x63, 0xa8, 0xd1, 0xca,
	0x64, 0x11, 0x83, 0xbb, 0x29, 0xdd, 0x19, 0x5b, 0xa5, 0x01, 0xbb, 0x29, 0x8a, 0xb2, 0x2b, 0x84,
	0x69, 0x66, 0x30, 0x02, 0x5b, 0xa3, 0x6a, 0x46, 0x8b, 0x52, 0xba, 0xe9, 0x2b, 0xf4, 0xd1, 0xfa,
	0x0a, 0x5d, 0x97, 0x3e, 0x42, 0x99, 0x1f, 0xa5, 0xb2, 0xf1, 0x2a, 0x3b, 0x9d, 0x39, 0x47, 0xe7,
	0xde, 0xf9, 0x84, 0x20, 0xac, 0x4b, 0x46, 0x15, 0xbf, 0x2e, 0x2b, 0xa1, 0x04, 0x06, 0x56, 0x0d,
	0x47, 0x1b, 0x21, 0x36, 0x5b, 0x3e, 0xa5, 0x65, 0x3e, 0xa5, 0x45, 0x21, 0x14, 0x55, 0xb9, 0x28,
	0xa4, 0x4d, 0xc5, 0x9f, 0x20, 0xbc, 0x53, 0xb4, 0x52, 0x29, 0xff, 0x52, 0x73, 0xa9, 0x30, 0x82,
	0x2e, 0x65, 0xac, 0xe2, 0x52, 0x46, 0x64, 0x4c, 0x26, 0xbd, 0xb4, 0x91, 0xda, 0x29, 0xe9, 0xd7,
	0xad, 0xa0, 0x2c, 0xf2, 0xac, 0xe3, 0xa4, 0x76, 0x54, 0xbe, 0xe3, 0xa2, 0x56, 0x51, 0x67, 0x4c,
	0x26, 0x9d, 0xb4, 0x91, 0xf1, 0x2b, 0x38, 0x77, 0xed, 0xb2, 0x14, 0x85, 0xe4, 0xd8, 0x07, 0x2f,
	0x67, 0xae, 0xd9, 0xcb, 0x99, 0x0b, 0xa8, 0x5a, 0x36, 0xf3, 0x0f, 0x03, 0x7f, 0x3c, 0xe8, 0x37,
	0x89, 0xe3, 0x1d, 0xf8, 0x06, 0x42, 0xd9, 0xba, 0x82, 0xd9, 0xee, 0x6c, 0x76, 0x71, 0xed, 0x68,
	0xb4, 0xaf, 0x97, 0xee, 0x25, 0x71, 0x06, 0xbe, 0x54, 0x54, 0x71, 0xb3, 0x76, 0x7f, 0x36, 0x6a,
	0xbd, 0xd2, 0x1a, 0x68, 0x24, 0x4f, 0x6d, 0x14, 0x87, 0x70, 0x5a, 0x56, 0x62, 0x63, 0x08, 0x9d,
	0x8c, 0xc9, 0xc4, 0x4f, 0x9f, 0xb4, 0x06, 0xb1, 0xe3, 0x52, 0xd2, 0x0d, 0x8f, 0x7c, 0x8b, 0xc8,
	0x49, 0xed, 0x98, 0xc9, 0x9c, 0x45, 0x81, 0x45, 0xe4, 0x24, 0x8e, 0xa0, 0xf7, 0x28, 0x76, 0xe5,
	0x96, 0x6b, 0xaf, 0x6b, 0xbc, 0xff, 0x07, 0x7a, 0x1a, 0xab, 0x2b, 0xf3, 0xc5, 0xa2, 0x53, 0x63,
	0x3e, 0xe9, 0x38, 0x03, 0xdf, 0x6c, 0x86, 0x3e, 0x90, 0x87, 0xc1, 0x0b, 0x3c, 0x83, 0xee, 0x5d,
	0x36, 0x4f, 0xb3, 0x64, 0x39, 0x20, 0x78, 0x0e, 0xbd, 0xc5, 0xed, 0xfa, 0xe3, 0x2a, 0xd1, 0xd2,
	0x33, 0x72, 0xfe, 0x61, 0x91, 0xac, 0x56, 0xc9, 0x72, 0xd0, 0x41, 0x80, 0xe0, 0xdd, 0xfc, 0x46,
	0x3f, 0x9f, 0x68, 0x2b, 0xbb, 0x59, 0x27, 0xcb, 0x87, 0xdb, 0xfb, 0x6c, 0xe0, 0xcf, 0xfe, 0x12,
	0x08, 0xee, 0x0d, 0x06, 0x7c, 0x6f, 0x06, 0x54, 0x0a, 0x8f, 0xb2, 0x1c, 0x5e, 0x1e, 0x9c, 0x5a,
	0x5a, 0x31, 0xfe, 0xf8, 0xf5, 0xfb, 0xa7, 0x17, 0xc6, 0xdd, 0xa9, 0xb5, 0xdf, 0x92, 0x2b, 0x5c,
	0x43, 0x60, 0x99, 0xe2, 0xe5, 0x21, 0x63, 0xdb, 0xf5, 0xf2, 0x38, 0xfa, 0xf8, 0xc2, 0x94, 0xf5,
	0x31, 0x74, 0x65, 0xd3, 0x6f, 0x39, 0xfb, 0xae, 0xeb, 0x16, 0xb4, 0x78, 0xe4, 0xdb, 0x67, 0xd6,
	0x5d, 0xed, 0xd5, 0x7d, 0x0e, 0xcc, 0xaf, 0xf0, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2,
	0xba, 0x60, 0x63, 0x40, 0x03, 0x00, 0x00,
}
