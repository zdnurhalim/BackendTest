// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passager.proto

/*
Package passager is a generated protocol buffer package.

It is generated from these files:
	passager.proto

It has these top-level messages:
	RequestResponse
	RequestUser
	RequestFilter
	AcceptUser
	AcceptResponse
	RequestLocation
	ResponseLocation
*/
package passager

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type RequestResponse struct {
	From    string `protobuf:"bytes,1,opt,name=From" json:"From,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *RequestResponse) Reset()                    { *m = RequestResponse{} }
func (m *RequestResponse) String() string            { return proto.CompactTextString(m) }
func (*RequestResponse) ProtoMessage()               {}
func (*RequestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RequestResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RequestUser struct {
	From   string  `protobuf:"bytes,1,opt,name=From" json:"From,omitempty"`
	Lat    float32 `protobuf:"fixed32,2,opt,name=Lat" json:"Lat,omitempty"`
	Lon    float32 `protobuf:"fixed32,3,opt,name=Lon" json:"Lon,omitempty"`
	Status string  `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
}

func (m *RequestUser) Reset()                    { *m = RequestUser{} }
func (m *RequestUser) String() string            { return proto.CompactTextString(m) }
func (*RequestUser) ProtoMessage()               {}
func (*RequestUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RequestUser) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RequestUser) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *RequestUser) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func (m *RequestUser) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RequestFilter struct {
	UserKey string `protobuf:"bytes,1,opt,name=UserKey" json:"UserKey,omitempty"`
}

func (m *RequestFilter) Reset()                    { *m = RequestFilter{} }
func (m *RequestFilter) String() string            { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()               {}
func (*RequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RequestFilter) GetUserKey() string {
	if m != nil {
		return m.UserKey
	}
	return ""
}

type AcceptUser struct {
	From string  `protobuf:"bytes,1,opt,name=From" json:"From,omitempty"`
	To   string  `protobuf:"bytes,2,opt,name=To" json:"To,omitempty"`
	Lat  float32 `protobuf:"fixed32,3,opt,name=Lat" json:"Lat,omitempty"`
	Lon  float32 `protobuf:"fixed32,4,opt,name=Lon" json:"Lon,omitempty"`
}

func (m *AcceptUser) Reset()                    { *m = AcceptUser{} }
func (m *AcceptUser) String() string            { return proto.CompactTextString(m) }
func (*AcceptUser) ProtoMessage()               {}
func (*AcceptUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AcceptUser) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *AcceptUser) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *AcceptUser) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *AcceptUser) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type AcceptResponse struct {
	From    string `protobuf:"bytes,1,opt,name=From" json:"From,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *AcceptResponse) Reset()                    { *m = AcceptResponse{} }
func (m *AcceptResponse) String() string            { return proto.CompactTextString(m) }
func (*AcceptResponse) ProtoMessage()               {}
func (*AcceptResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AcceptResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *AcceptResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RequestLocation struct {
	From string  `protobuf:"bytes,1,opt,name=From" json:"From,omitempty"`
	To   string  `protobuf:"bytes,2,opt,name=To" json:"To,omitempty"`
	Lat  float32 `protobuf:"fixed32,3,opt,name=Lat" json:"Lat,omitempty"`
	Lon  float32 `protobuf:"fixed32,4,opt,name=Lon" json:"Lon,omitempty"`
}

func (m *RequestLocation) Reset()                    { *m = RequestLocation{} }
func (m *RequestLocation) String() string            { return proto.CompactTextString(m) }
func (*RequestLocation) ProtoMessage()               {}
func (*RequestLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RequestLocation) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RequestLocation) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *RequestLocation) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *RequestLocation) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type ResponseLocation struct {
	From    string `protobuf:"bytes,1,opt,name=From" json:"From,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *ResponseLocation) Reset()                    { *m = ResponseLocation{} }
func (m *ResponseLocation) String() string            { return proto.CompactTextString(m) }
func (*ResponseLocation) ProtoMessage()               {}
func (*ResponseLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ResponseLocation) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ResponseLocation) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*RequestResponse)(nil), "passager.RequestResponse")
	proto.RegisterType((*RequestUser)(nil), "passager.RequestUser")
	proto.RegisterType((*RequestFilter)(nil), "passager.RequestFilter")
	proto.RegisterType((*AcceptUser)(nil), "passager.AcceptUser")
	proto.RegisterType((*AcceptResponse)(nil), "passager.AcceptResponse")
	proto.RegisterType((*RequestLocation)(nil), "passager.RequestLocation")
	proto.RegisterType((*ResponseLocation)(nil), "passager.ResponseLocation")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Passager service

type PassagerClient interface {
	SendRequest(ctx context.Context, in *RequestUser, opts ...grpc.CallOption) (*RequestResponse, error)
	GetRequestStatus(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Passager_GetRequestStatusClient, error)
	ReceiveRequest(ctx context.Context, in *AcceptUser, opts ...grpc.CallOption) (*AcceptResponse, error)
	GetLocation(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Passager_GetLocationClient, error)
	ReceiveLocation(ctx context.Context, in *RequestLocation, opts ...grpc.CallOption) (*ResponseLocation, error)
}

type passagerClient struct {
	cc *grpc.ClientConn
}

func NewPassagerClient(cc *grpc.ClientConn) PassagerClient {
	return &passagerClient{cc}
}

func (c *passagerClient) SendRequest(ctx context.Context, in *RequestUser, opts ...grpc.CallOption) (*RequestResponse, error) {
	out := new(RequestResponse)
	err := grpc.Invoke(ctx, "/passager.passager/SendRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passagerClient) GetRequestStatus(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Passager_GetRequestStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Passager_serviceDesc.Streams[0], c.cc, "/passager.passager/GetRequestStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &passagerGetRequestStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Passager_GetRequestStatusClient interface {
	Recv() (*AcceptUser, error)
	grpc.ClientStream
}

type passagerGetRequestStatusClient struct {
	grpc.ClientStream
}

func (x *passagerGetRequestStatusClient) Recv() (*AcceptUser, error) {
	m := new(AcceptUser)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *passagerClient) ReceiveRequest(ctx context.Context, in *AcceptUser, opts ...grpc.CallOption) (*AcceptResponse, error) {
	out := new(AcceptResponse)
	err := grpc.Invoke(ctx, "/passager.passager/ReceiveRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passagerClient) GetLocation(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Passager_GetLocationClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Passager_serviceDesc.Streams[1], c.cc, "/passager.passager/getLocation", opts...)
	if err != nil {
		return nil, err
	}
	x := &passagerGetLocationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Passager_GetLocationClient interface {
	Recv() (*RequestLocation, error)
	grpc.ClientStream
}

type passagerGetLocationClient struct {
	grpc.ClientStream
}

func (x *passagerGetLocationClient) Recv() (*RequestLocation, error) {
	m := new(RequestLocation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *passagerClient) ReceiveLocation(ctx context.Context, in *RequestLocation, opts ...grpc.CallOption) (*ResponseLocation, error) {
	out := new(ResponseLocation)
	err := grpc.Invoke(ctx, "/passager.passager/ReceiveLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Passager service

type PassagerServer interface {
	SendRequest(context.Context, *RequestUser) (*RequestResponse, error)
	GetRequestStatus(*RequestFilter, Passager_GetRequestStatusServer) error
	ReceiveRequest(context.Context, *AcceptUser) (*AcceptResponse, error)
	GetLocation(*RequestFilter, Passager_GetLocationServer) error
	ReceiveLocation(context.Context, *RequestLocation) (*ResponseLocation, error)
}

func RegisterPassagerServer(s *grpc.Server, srv PassagerServer) {
	s.RegisterService(&_Passager_serviceDesc, srv)
}

func _Passager_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassagerServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passager.passager/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassagerServer).SendRequest(ctx, req.(*RequestUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passager_GetRequestStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PassagerServer).GetRequestStatus(m, &passagerGetRequestStatusServer{stream})
}

type Passager_GetRequestStatusServer interface {
	Send(*AcceptUser) error
	grpc.ServerStream
}

type passagerGetRequestStatusServer struct {
	grpc.ServerStream
}

func (x *passagerGetRequestStatusServer) Send(m *AcceptUser) error {
	return x.ServerStream.SendMsg(m)
}

func _Passager_ReceiveRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassagerServer).ReceiveRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passager.passager/ReceiveRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassagerServer).ReceiveRequest(ctx, req.(*AcceptUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passager_GetLocation_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PassagerServer).GetLocation(m, &passagerGetLocationServer{stream})
}

type Passager_GetLocationServer interface {
	Send(*RequestLocation) error
	grpc.ServerStream
}

type passagerGetLocationServer struct {
	grpc.ServerStream
}

func (x *passagerGetLocationServer) Send(m *RequestLocation) error {
	return x.ServerStream.SendMsg(m)
}

func _Passager_ReceiveLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassagerServer).ReceiveLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passager.passager/ReceiveLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassagerServer).ReceiveLocation(ctx, req.(*RequestLocation))
	}
	return interceptor(ctx, in, info, handler)
}

var _Passager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passager.passager",
	HandlerType: (*PassagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRequest",
			Handler:    _Passager_SendRequest_Handler,
		},
		{
			MethodName: "ReceiveRequest",
			Handler:    _Passager_ReceiveRequest_Handler,
		},
		{
			MethodName: "ReceiveLocation",
			Handler:    _Passager_ReceiveLocation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRequestStatus",
			Handler:       _Passager_GetRequestStatus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "getLocation",
			Handler:       _Passager_GetLocation_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "passager.proto",
}

func init() { proto.RegisterFile("passager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x9b, 0xb4, 0x2a, 0xed, 0x55, 0x84, 0xca, 0xe2, 0x8f, 0xc9, 0x54, 0x79, 0x2a, 0x4b,
	0x85, 0x60, 0x07, 0x8a, 0x44, 0x91, 0xa0, 0x93, 0x5b, 0x06, 0x06, 0x86, 0x60, 0x4e, 0x55, 0x25,
	0x88, 0x43, 0xec, 0x20, 0xf1, 0x99, 0xf9, 0x12, 0x28, 0xae, 0x9d, 0xb4, 0xa4, 0x61, 0x40, 0x6c,
	0xbe, 0x73, 0xde, 0xf3, 0x2f, 0xf7, 0x6c, 0x08, 0x92, 0x48, 0xa9, 0x68, 0x81, 0xe9, 0x28, 0x49,
	0xa5, 0x96, 0xa4, 0xe3, 0x6a, 0x76, 0x09, 0x7b, 0x1c, 0xdf, 0x33, 0x54, 0x9a, 0xa3, 0x4a, 0x64,
	0xac, 0x90, 0x10, 0x68, 0x4d, 0x52, 0xf9, 0x46, 0xbd, 0x81, 0x37, 0xec, 0x72, 0xb3, 0x26, 0x14,
	0x76, 0x54, 0x26, 0x04, 0x2a, 0x45, 0xfd, 0x81, 0x37, 0xec, 0x70, 0x57, 0xb2, 0x27, 0xe8, 0x59,
	0x83, 0x07, 0x85, 0xe9, 0x56, 0x71, 0x1f, 0x9a, 0xd3, 0x48, 0x1b, 0xa1, 0xcf, 0xf3, 0xa5, 0xe9,
	0xc8, 0x98, 0x36, 0x6d, 0x47, 0xc6, 0xe4, 0x10, 0xda, 0x4a, 0x47, 0x3a, 0x53, 0xb4, 0x65, 0x94,
	0xb6, 0x62, 0x27, 0xb0, 0x6b, 0xed, 0x27, 0xcb, 0x57, 0x8d, 0x69, 0x4e, 0x92, 0x1f, 0x74, 0x8f,
	0x9f, 0xf6, 0x0c, 0x57, 0xb2, 0x39, 0xc0, 0x58, 0x08, 0x4c, 0xea, 0x41, 0x02, 0xf0, 0xe7, 0xd2,
	0x70, 0x74, 0xb9, 0x3f, 0x97, 0x0e, 0xac, 0x59, 0x01, 0x6b, 0x15, 0x60, 0xec, 0x02, 0x82, 0x95,
	0xeb, 0x1f, 0xe7, 0xf3, 0x58, 0x0c, 0x78, 0x2a, 0x45, 0xa4, 0x97, 0x32, 0xfe, 0x37, 0xb4, 0x2b,
	0xe8, 0x3b, 0xa8, 0x5f, 0xbd, 0x6b, 0xe1, 0xce, 0xbe, 0x7c, 0x28, 0xae, 0x02, 0x19, 0x43, 0x6f,
	0x86, 0xf1, 0x8b, 0xa5, 0x25, 0x07, 0xa3, 0xe2, 0xd2, 0xac, 0x05, 0x1c, 0x1e, 0x57, 0xda, 0x8e,
	0x81, 0x35, 0xc8, 0x0d, 0xf4, 0x6f, 0x51, 0xdb, 0xfe, 0xcc, 0x24, 0x48, 0x8e, 0x2a, 0x82, 0x55,
	0x92, 0xe1, 0x7e, 0xb9, 0x51, 0xe6, 0xc6, 0x1a, 0xa7, 0x1e, 0xb9, 0x86, 0x80, 0xa3, 0xc0, 0xe5,
	0x07, 0x3a, 0x98, 0xad, 0xdf, 0x86, 0xf4, 0x67, 0x77, 0x03, 0xa5, 0xb7, 0xc0, 0x72, 0xe6, 0xb5,
	0x14, 0xd5, 0xff, 0x71, 0x1a, 0x83, 0x72, 0x97, 0xc7, 0x67, 0x50, 0x0a, 0xab, 0x7a, 0x45, 0x18,
	0xae, 0x6f, 0x6d, 0x26, 0xc3, 0x1a, 0xcf, 0x6d, 0xf3, 0xf8, 0xce, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x44, 0x11, 0x66, 0xb5, 0x8e, 0x03, 0x00, 0x00,
}
