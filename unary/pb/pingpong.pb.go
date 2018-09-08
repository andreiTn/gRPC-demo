// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/pingpong.proto

package pingpong

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

type PPRequest struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PPRequest) Reset()         { *m = PPRequest{} }
func (m *PPRequest) String() string { return proto.CompactTextString(m) }
func (*PPRequest) ProtoMessage()    {}
func (*PPRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pingpong_733e60cc9076bdcf, []int{0}
}
func (m *PPRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PPRequest.Unmarshal(m, b)
}
func (m *PPRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PPRequest.Marshal(b, m, deterministic)
}
func (dst *PPRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PPRequest.Merge(dst, src)
}
func (m *PPRequest) XXX_Size() int {
	return xxx_messageInfo_PPRequest.Size(m)
}
func (m *PPRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PPRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PPRequest proto.InternalMessageInfo

func (m *PPRequest) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type PPResponse struct {
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PPResponse) Reset()         { *m = PPResponse{} }
func (m *PPResponse) String() string { return proto.CompactTextString(m) }
func (*PPResponse) ProtoMessage()    {}
func (*PPResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pingpong_733e60cc9076bdcf, []int{1}
}
func (m *PPResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PPResponse.Unmarshal(m, b)
}
func (m *PPResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PPResponse.Marshal(b, m, deterministic)
}
func (dst *PPResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PPResponse.Merge(dst, src)
}
func (m *PPResponse) XXX_Size() int {
	return xxx_messageInfo_PPResponse.Size(m)
}
func (m *PPResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PPResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PPResponse proto.InternalMessageInfo

func (m *PPResponse) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*PPRequest)(nil), "PPRequest")
	proto.RegisterType((*PPResponse)(nil), "PPResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PPServiceClient is the client API for PPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PPServiceClient interface {
	PingPong(ctx context.Context, in *PPRequest, opts ...grpc.CallOption) (*PPResponse, error)
}

type pPServiceClient struct {
	cc *grpc.ClientConn
}

func NewPPServiceClient(cc *grpc.ClientConn) PPServiceClient {
	return &pPServiceClient{cc}
}

func (c *pPServiceClient) PingPong(ctx context.Context, in *PPRequest, opts ...grpc.CallOption) (*PPResponse, error) {
	out := new(PPResponse)
	err := c.cc.Invoke(ctx, "/PPService/PingPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PPServiceServer is the server API for PPService service.
type PPServiceServer interface {
	PingPong(context.Context, *PPRequest) (*PPResponse, error)
}

func RegisterPPServiceServer(s *grpc.Server, srv PPServiceServer) {
	s.RegisterService(&_PPService_serviceDesc, srv)
}

func _PPService_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PPServiceServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PPService/PingPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PPServiceServer).PingPong(ctx, req.(*PPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PPService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PPService",
	HandlerType: (*PPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingPong",
			Handler:    _PPService_PingPong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pingpong.proto",
}

func init() { proto.RegisterFile("pb/pingpong.proto", fileDescriptor_pingpong_733e60cc9076bdcf) }

var fileDescriptor_pingpong_733e60cc9076bdcf = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x48, 0xd2, 0x2f,
	0xc8, 0xcc, 0x4b, 0x2f, 0xc8, 0xcf, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7,
	0xe2, 0x0c, 0x08, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x01, 0x49,
	0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x0a, 0x5c, 0x5c, 0x20, 0x05, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x60, 0x15, 0xf9, 0x48, 0x2a, 0xf2, 0xf3, 0xd2, 0x8d, 0x8c, 0x40,
	0x46, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xa9, 0x72, 0x71, 0x04, 0x64, 0xe6, 0xa5,
	0x07, 0xe4, 0xe7, 0xa5, 0x0b, 0x71, 0xe9, 0xc1, 0x8d, 0x96, 0xe2, 0xd6, 0x43, 0x98, 0xa2, 0xc4,
	0x90, 0xc4, 0x06, 0xb6, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xd6, 0x01, 0x2a, 0x92,
	0x00, 0x00, 0x00,
}
