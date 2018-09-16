// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/bid.proto

package bid

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

type Req struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_bid_f3f8fd6ad45595b3, []int{0}
}
func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (dst *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(dst, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Res struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Res) Reset()         { *m = Res{} }
func (m *Res) String() string { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()    {}
func (*Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_bid_f3f8fd6ad45595b3, []int{1}
}
func (m *Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Res.Unmarshal(m, b)
}
func (m *Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Res.Marshal(b, m, deterministic)
}
func (dst *Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Res.Merge(dst, src)
}
func (m *Res) XXX_Size() int {
	return xxx_messageInfo_Res.Size(m)
}
func (m *Res) XXX_DiscardUnknown() {
	xxx_messageInfo_Res.DiscardUnknown(m)
}

var xxx_messageInfo_Res proto.InternalMessageInfo

func (m *Res) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "Req")
	proto.RegisterType((*Res)(nil), "Res")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	Message(ctx context.Context, opts ...grpc.CallOption) (ChatService_MessageClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Message(ctx context.Context, opts ...grpc.CallOption) (ChatService_MessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/ChatService/Message", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceMessageClient{stream}
	return x, nil
}

type ChatService_MessageClient interface {
	Send(*Req) error
	Recv() (*Res, error)
	grpc.ClientStream
}

type chatServiceMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceMessageClient) Send(m *Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceMessageClient) Recv() (*Res, error) {
	m := new(Res)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	Message(ChatService_MessageServer) error
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Message_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Message(&chatServiceMessageServer{stream})
}

type ChatService_MessageServer interface {
	Send(*Res) error
	Recv() (*Req, error)
	grpc.ServerStream
}

type chatServiceMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceMessageServer) Send(m *Res) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceMessageServer) Recv() (*Req, error) {
	m := new(Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Message",
			Handler:       _ChatService_Message_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/bid.proto",
}

func init() { proto.RegisterFile("pb/bid.proto", fileDescriptor_bid_f3f8fd6ad45595b3) }

var fileDescriptor_bid_f3f8fd6ad45595b3 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0xd2, 0x4f,
	0xca, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x12, 0xe7, 0x62, 0x0e, 0x4a, 0x2d, 0x14,
	0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x21,
	0x12, 0xc5, 0x98, 0x12, 0x46, 0x5a, 0x5c, 0xdc, 0xce, 0x19, 0x89, 0x25, 0xc1, 0xa9, 0x45, 0x65,
	0x99, 0xc9, 0xa9, 0x42, 0xd2, 0x5c, 0xec, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x2c,
	0x7a, 0x41, 0xa9, 0x85, 0x52, 0x20, 0xb2, 0x58, 0x89, 0x41, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d,
	0x6c, 0x89, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x45, 0x9c, 0x87, 0x29, 0x74, 0x00, 0x00, 0x00,
}