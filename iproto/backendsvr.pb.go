// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backendsvr.proto

package iproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type GwError struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GwError) Reset()         { *m = GwError{} }
func (m *GwError) String() string { return proto.CompactTextString(m) }
func (*GwError) ProtoMessage()    {}
func (*GwError) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d30a88daa3bb8f4, []int{0}
}

func (m *GwError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GwError.Unmarshal(m, b)
}
func (m *GwError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GwError.Marshal(b, m, deterministic)
}
func (m *GwError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GwError.Merge(m, src)
}
func (m *GwError) XXX_Size() int {
	return xxx_messageInfo_GwError.Size(m)
}
func (m *GwError) XXX_DiscardUnknown() {
	xxx_messageInfo_GwError.DiscardUnknown(m)
}

var xxx_messageInfo_GwError proto.InternalMessageInfo

func (m *GwError) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GwError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Method1Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Method1Request) Reset()         { *m = Method1Request{} }
func (m *Method1Request) String() string { return proto.CompactTextString(m) }
func (*Method1Request) ProtoMessage()    {}
func (*Method1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d30a88daa3bb8f4, []int{1}
}

func (m *Method1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Method1Request.Unmarshal(m, b)
}
func (m *Method1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Method1Request.Marshal(b, m, deterministic)
}
func (m *Method1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method1Request.Merge(m, src)
}
func (m *Method1Request) XXX_Size() int {
	return xxx_messageInfo_Method1Request.Size(m)
}
func (m *Method1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Method1Request.DiscardUnknown(m)
}

var xxx_messageInfo_Method1Request proto.InternalMessageInfo

type Method1Reply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Method1Reply) Reset()         { *m = Method1Reply{} }
func (m *Method1Reply) String() string { return proto.CompactTextString(m) }
func (*Method1Reply) ProtoMessage()    {}
func (*Method1Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d30a88daa3bb8f4, []int{2}
}

func (m *Method1Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Method1Reply.Unmarshal(m, b)
}
func (m *Method1Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Method1Reply.Marshal(b, m, deterministic)
}
func (m *Method1Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method1Reply.Merge(m, src)
}
func (m *Method1Reply) XXX_Size() int {
	return xxx_messageInfo_Method1Reply.Size(m)
}
func (m *Method1Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Method1Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Method1Reply proto.InternalMessageInfo

type Method2Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Method2Request) Reset()         { *m = Method2Request{} }
func (m *Method2Request) String() string { return proto.CompactTextString(m) }
func (*Method2Request) ProtoMessage()    {}
func (*Method2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d30a88daa3bb8f4, []int{3}
}

func (m *Method2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Method2Request.Unmarshal(m, b)
}
func (m *Method2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Method2Request.Marshal(b, m, deterministic)
}
func (m *Method2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method2Request.Merge(m, src)
}
func (m *Method2Request) XXX_Size() int {
	return xxx_messageInfo_Method2Request.Size(m)
}
func (m *Method2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Method2Request.DiscardUnknown(m)
}

var xxx_messageInfo_Method2Request proto.InternalMessageInfo

type Method2Reply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Method2Reply) Reset()         { *m = Method2Reply{} }
func (m *Method2Reply) String() string { return proto.CompactTextString(m) }
func (*Method2Reply) ProtoMessage()    {}
func (*Method2Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d30a88daa3bb8f4, []int{4}
}

func (m *Method2Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Method2Reply.Unmarshal(m, b)
}
func (m *Method2Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Method2Reply.Marshal(b, m, deterministic)
}
func (m *Method2Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method2Reply.Merge(m, src)
}
func (m *Method2Reply) XXX_Size() int {
	return xxx_messageInfo_Method2Reply.Size(m)
}
func (m *Method2Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Method2Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Method2Reply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GwError)(nil), "iproto.GwError")
	proto.RegisterType((*Method1Request)(nil), "iproto.Method1Request")
	proto.RegisterType((*Method1Reply)(nil), "iproto.Method1Reply")
	proto.RegisterType((*Method2Request)(nil), "iproto.Method2Request")
	proto.RegisterType((*Method2Reply)(nil), "iproto.Method2Reply")
}

func init() { proto.RegisterFile("backendsvr.proto", fileDescriptor_6d30a88daa3bb8f4) }

var fileDescriptor_6d30a88daa3bb8f4 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4a, 0x4c, 0xce,
	0x4e, 0xcd, 0x4b, 0x29, 0x2e, 0x2b, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x04,
	0xd3, 0x4a, 0xe6, 0x5c, 0xec, 0xee, 0xe5, 0xae, 0x45, 0x45, 0xf9, 0x45, 0x42, 0x42, 0x5c, 0x2c,
	0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x60, 0xb6, 0x90, 0x04, 0x17,
	0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c,
	0xab, 0x24, 0xc0, 0xc5, 0xe7, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0x62, 0x18, 0x94, 0x5a, 0x58, 0x9a,
	0x5a, 0x5c, 0xa2, 0xc4, 0xc7, 0xc5, 0x03, 0x17, 0x29, 0xc8, 0xa9, 0x44, 0xa8, 0x30, 0xc2, 0x50,
	0x61, 0x04, 0x56, 0x61, 0xe4, 0xc1, 0xc5, 0x8d, 0x70, 0x98, 0xa1, 0x90, 0x25, 0x17, 0x3b, 0xd4,
	0x00, 0x21, 0x31, 0x3d, 0x88, 0xfb, 0xf4, 0x50, 0xed, 0x90, 0x12, 0xc1, 0x10, 0x07, 0xd9, 0xc4,
	0x80, 0x6a, 0x92, 0x11, 0xc2, 0x24, 0x23, 0x74, 0x93, 0x8c, 0x70, 0x98, 0x64, 0x04, 0x35, 0xc9,
	0x49, 0x95, 0x4b, 0x20, 0x39, 0x3f, 0x57, 0xaf, 0xaa, 0x30, 0x2f, 0xb5, 0x04, 0xaa, 0xc4, 0x09,
	0x1a, 0x58, 0x01, 0x8c, 0x8b, 0x98, 0xa0, 0xcc, 0x24, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0xe4, 0x8b, 0x76, 0x5a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Backendsvr1Client is the client API for Backendsvr1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Backendsvr1Client interface {
	Method1(ctx context.Context, in *Method1Request, opts ...grpc.CallOption) (*Method1Reply, error)
}

type backendsvr1Client struct {
	cc *grpc.ClientConn
}

func NewBackendsvr1Client(cc *grpc.ClientConn) Backendsvr1Client {
	return &backendsvr1Client{cc}
}

func (c *backendsvr1Client) Method1(ctx context.Context, in *Method1Request, opts ...grpc.CallOption) (*Method1Reply, error) {
	out := new(Method1Reply)
	err := c.cc.Invoke(ctx, "/iproto.backendsvr1/Method1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Backendsvr1Server is the server API for Backendsvr1 service.
type Backendsvr1Server interface {
	Method1(context.Context, *Method1Request) (*Method1Reply, error)
}

// UnimplementedBackendsvr1Server can be embedded to have forward compatible implementations.
type UnimplementedBackendsvr1Server struct {
}

func (*UnimplementedBackendsvr1Server) Method1(ctx context.Context, req *Method1Request) (*Method1Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Method1 not implemented")
}

func RegisterBackendsvr1Server(s *grpc.Server, srv Backendsvr1Server) {
	s.RegisterService(&_Backendsvr1_serviceDesc, srv)
}

func _Backendsvr1_Method1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Method1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Backendsvr1Server).Method1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iproto.backendsvr1/Method1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Backendsvr1Server).Method1(ctx, req.(*Method1Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backendsvr1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iproto.backendsvr1",
	HandlerType: (*Backendsvr1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method1",
			Handler:    _Backendsvr1_Method1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backendsvr.proto",
}

// Backendsvr2Client is the client API for Backendsvr2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Backendsvr2Client interface {
	Method2(ctx context.Context, in *Method2Request, opts ...grpc.CallOption) (*Method2Reply, error)
}

type backendsvr2Client struct {
	cc *grpc.ClientConn
}

func NewBackendsvr2Client(cc *grpc.ClientConn) Backendsvr2Client {
	return &backendsvr2Client{cc}
}

func (c *backendsvr2Client) Method2(ctx context.Context, in *Method2Request, opts ...grpc.CallOption) (*Method2Reply, error) {
	out := new(Method2Reply)
	err := c.cc.Invoke(ctx, "/iproto.backendsvr2/Method2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Backendsvr2Server is the server API for Backendsvr2 service.
type Backendsvr2Server interface {
	Method2(context.Context, *Method2Request) (*Method2Reply, error)
}

// UnimplementedBackendsvr2Server can be embedded to have forward compatible implementations.
type UnimplementedBackendsvr2Server struct {
}

func (*UnimplementedBackendsvr2Server) Method2(ctx context.Context, req *Method2Request) (*Method2Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Method2 not implemented")
}

func RegisterBackendsvr2Server(s *grpc.Server, srv Backendsvr2Server) {
	s.RegisterService(&_Backendsvr2_serviceDesc, srv)
}

func _Backendsvr2_Method2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Method2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Backendsvr2Server).Method2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iproto.backendsvr2/Method2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Backendsvr2Server).Method2(ctx, req.(*Method2Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backendsvr2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iproto.backendsvr2",
	HandlerType: (*Backendsvr2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method2",
			Handler:    _Backendsvr2_Method2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backendsvr.proto",
}
