// Code generated by protoc-gen-go.
// source: improbable/grpcweb/test/test.proto
// DO NOT EDIT!

/*
Package improbable_grpcweb_test is a generated protocol buffer package.

It is generated from these files:
	improbable/grpcweb/test/test.proto

It has these top-level messages:
	PingRequest
	PingResponse
*/
package improbable_grpcweb_test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type PingRequest_FailureType int32

const (
	PingRequest_NONE PingRequest_FailureType = 0
	PingRequest_CODE PingRequest_FailureType = 1
	PingRequest_DROP PingRequest_FailureType = 2
)

var PingRequest_FailureType_name = map[int32]string{
	0: "NONE",
	1: "CODE",
	2: "DROP",
}
var PingRequest_FailureType_value = map[string]int32{
	"NONE": 0,
	"CODE": 1,
	"DROP": 2,
}

func (x PingRequest_FailureType) String() string {
	return proto.EnumName(PingRequest_FailureType_name, int32(x))
}
func (PingRequest_FailureType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type PingRequest struct {
	Value             string                  `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	SleepTimeMs       int32                   `protobuf:"varint,2,opt,name=sleep_time_ms,json=sleepTimeMs" json:"sleep_time_ms,omitempty"`
	ResponseCount     int32                   `protobuf:"varint,3,opt,name=response_count,json=responseCount" json:"response_count,omitempty"`
	ErrorCodeReturned uint32                  `protobuf:"varint,4,opt,name=error_code_returned,json=errorCodeReturned" json:"error_code_returned,omitempty"`
	FailureType       PingRequest_FailureType `protobuf:"varint,5,opt,name=failure_type,json=failureType,enum=improbable.grpcweb.test.PingRequest_FailureType" json:"failure_type,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingRequest) GetSleepTimeMs() int32 {
	if m != nil {
		return m.SleepTimeMs
	}
	return 0
}

func (m *PingRequest) GetResponseCount() int32 {
	if m != nil {
		return m.ResponseCount
	}
	return 0
}

func (m *PingRequest) GetErrorCodeReturned() uint32 {
	if m != nil {
		return m.ErrorCodeReturned
	}
	return 0
}

func (m *PingRequest) GetFailureType() PingRequest_FailureType {
	if m != nil {
		return m.FailureType
	}
	return PingRequest_NONE
}

type PingResponse struct {
	Value   string `protobuf:"bytes,1,opt,name=Value" json:"Value,omitempty"`
	Counter int32  `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingResponse) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "improbable.grpcweb.test.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "improbable.grpcweb.test.PingResponse")
	proto.RegisterEnum("improbable.grpcweb.test.PingRequest_FailureType", PingRequest_FailureType_name, PingRequest_FailureType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	PingEmpty(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) PingEmpty(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/improbable.grpcweb.test.TestService/PingEmpty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/improbable.grpcweb.test.TestService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/improbable.grpcweb.test.TestService/PingError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/improbable.grpcweb.test.TestService/PingList", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_PingListClient interface {
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingListClient struct {
	grpc.ClientStream
}

func (x *testServicePingListClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	PingEmpty(context.Context, *google_protobuf.Empty) (*PingResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingError(context.Context, *PingRequest) (*google_protobuf.Empty, error)
	PingList(*PingRequest, TestService_PingListServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_PingEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.TestService/PingEmpty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingEmpty(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.TestService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.TestService/PingError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingError(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).PingList(m, &testServicePingListServer{stream})
}

type TestService_PingListServer interface {
	Send(*PingResponse) error
	grpc.ServerStream
}

type testServicePingListServer struct {
	grpc.ServerStream
}

func (x *testServicePingListServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.grpcweb.test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingEmpty",
			Handler:    _TestService_PingEmpty_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
		{
			MethodName: "PingError",
			Handler:    _TestService_PingError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingList",
			Handler:       _TestService_PingList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "improbable/grpcweb/test/test.proto",
}

// Client API for FailService service

type FailServiceClient interface {
	NonExistant(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type failServiceClient struct {
	cc *grpc.ClientConn
}

func NewFailServiceClient(cc *grpc.ClientConn) FailServiceClient {
	return &failServiceClient{cc}
}

func (c *failServiceClient) NonExistant(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/improbable.grpcweb.test.FailService/NonExistant", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FailService service

type FailServiceServer interface {
	NonExistant(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterFailServiceServer(s *grpc.Server, srv FailServiceServer) {
	s.RegisterService(&_FailService_serviceDesc, srv)
}

func _FailService_NonExistant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailServiceServer).NonExistant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.grpcweb.test.FailService/NonExistant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailServiceServer).NonExistant(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.grpcweb.test.FailService",
	HandlerType: (*FailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NonExistant",
			Handler:    _FailService_NonExistant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/grpcweb/test/test.proto",
}

func init() { proto.RegisterFile("improbable/grpcweb/test/test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x51, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x43, 0x02, 0xed, 0xb8, 0xa9, 0xc2, 0x82, 0xc0, 0x2a, 0x97, 0xc8, 0xa2, 0x52, 0x24,
	0xa4, 0x4d, 0x15, 0xee, 0x5c, 0x52, 0x73, 0xa1, 0x24, 0x95, 0x1b, 0xb8, 0x80, 0x64, 0xd9, 0xce,
	0xc4, 0x5a, 0x61, 0x7b, 0x97, 0xdd, 0x75, 0x21, 0x7f, 0xc3, 0x87, 0xf1, 0x31, 0x68, 0x77, 0x6d,
	0x35, 0x97, 0x8a, 0x1c, 0x72, 0xb1, 0x76, 0xde, 0xbc, 0x19, 0xcf, 0x7b, 0x0f, 0x42, 0x56, 0x09,
	0xc9, 0xb3, 0x34, 0x2b, 0x71, 0x56, 0x48, 0x91, 0xff, 0xc2, 0x6c, 0xa6, 0x51, 0x69, 0xfb, 0xa1,
	0x42, 0x72, 0xcd, 0xc9, 0xeb, 0x07, 0x0e, 0x6d, 0x39, 0xd4, 0xb4, 0x2f, 0xde, 0x14, 0x9c, 0x17,
	0x25, 0xce, 0x2c, 0x2d, 0x6b, 0xb6, 0x33, 0xac, 0x84, 0xde, 0xb9, 0xa9, 0xf0, 0x4f, 0x1f, 0xfc,
	0x5b, 0x56, 0x17, 0x31, 0xfe, 0x6c, 0x50, 0x69, 0xf2, 0x12, 0x86, 0xf7, 0x69, 0xd9, 0x60, 0xe0,
	0x4d, 0xbc, 0xe9, 0x69, 0xec, 0x0a, 0x12, 0xc2, 0x48, 0x95, 0x88, 0x22, 0xd1, 0xac, 0xc2, 0xa4,
	0x52, 0x41, 0x7f, 0xe2, 0x4d, 0x87, 0xb1, 0x6f, 0xc1, 0x35, 0xab, 0xf0, 0xb3, 0x22, 0x97, 0x70,
	0x2e, 0x51, 0x09, 0x5e, 0x2b, 0x4c, 0x72, 0xde, 0xd4, 0x3a, 0x78, 0x62, 0x49, 0xa3, 0x0e, 0x5d,
	0x18, 0x90, 0x50, 0x78, 0x81, 0x52, 0x72, 0x99, 0xe4, 0x7c, 0x83, 0x89, 0x44, 0xdd, 0xc8, 0x1a,
	0x37, 0xc1, 0x60, 0xe2, 0x4d, 0x47, 0xf1, 0x73, 0xdb, 0x5a, 0xf0, 0x0d, 0xc6, 0x6d, 0x83, 0xdc,
	0xc1, 0xd9, 0x36, 0x65, 0x65, 0x23, 0x31, 0xd1, 0x3b, 0x81, 0xc1, 0x70, 0xe2, 0x4d, 0xcf, 0xe7,
	0x57, 0xf4, 0x11, 0xb5, 0x74, 0x4f, 0x0c, 0xfd, 0xe8, 0x06, 0xd7, 0x3b, 0x81, 0xb1, 0xbf, 0x7d,
	0x28, 0xc2, 0x77, 0xe0, 0xef, 0xf5, 0xc8, 0x09, 0x0c, 0x96, 0xab, 0x65, 0x34, 0xee, 0x99, 0xd7,
	0x62, 0x75, 0x1d, 0x8d, 0x3d, 0xf3, 0xba, 0x8e, 0x57, 0xb7, 0xe3, 0x7e, 0xf8, 0x01, 0xce, 0xdc,
	0x52, 0x27, 0xc3, 0x58, 0xf4, 0x75, 0xdf, 0x22, 0x5b, 0x90, 0x00, 0x9e, 0x59, 0xd5, 0x28, 0x5b,
	0x73, 0xba, 0x72, 0xfe, 0xb7, 0x0f, 0xfe, 0x1a, 0x95, 0xbe, 0x43, 0x79, 0xcf, 0x72, 0x24, 0x37,
	0x70, 0x6a, 0xf6, 0x45, 0x26, 0x05, 0xf2, 0x8a, 0xba, 0x74, 0x68, 0x97, 0x0e, 0xb5, 0xf8, 0xc5,
	0xe5, 0x7f, 0x04, 0xba, 0x5b, 0xc2, 0x1e, 0xf9, 0x02, 0x03, 0x83, 0x90, 0xb7, 0x87, 0x38, 0x72,
	0xf8, 0xda, 0x4f, 0xed, 0x91, 0x26, 0x8f, 0x03, 0x77, 0x3f, 0x22, 0x25, 0xec, 0x91, 0x6f, 0x70,
	0x62, 0x88, 0x37, 0x4c, 0xe9, 0x23, 0xdf, 0x79, 0xe5, 0xcd, 0x7f, 0xb8, 0x2c, 0x3b, 0x77, 0xbf,
	0x83, 0xbf, 0xe4, 0x75, 0xf4, 0x9b, 0x29, 0x9d, 0xd6, 0xc7, 0xfe, 0x5d, 0xf6, 0xd4, 0x6a, 0x7b,
	0xff, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x3c, 0xb8, 0xeb, 0x91, 0x03, 0x00, 0x00,
}
