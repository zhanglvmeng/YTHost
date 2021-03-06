// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package pb

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

// empty message
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// message that only contains a value of string
type StringMsg struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMsg) Reset()         { *m = StringMsg{} }
func (m *StringMsg) String() string { return proto.CompactTextString(m) }
func (*StringMsg) ProtoMessage()    {}
func (*StringMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *StringMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMsg.Unmarshal(m, b)
}
func (m *StringMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMsg.Marshal(b, m, deterministic)
}
func (m *StringMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMsg.Merge(m, src)
}
func (m *StringMsg) XXX_Size() int {
	return xxx_messageInfo_StringMsg.Size(m)
}
func (m *StringMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StringMsg proto.InternalMessageInfo

func (m *StringMsg) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// message that contains a list of string
type StringListMsg struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringListMsg) Reset()         { *m = StringListMsg{} }
func (m *StringListMsg) String() string { return proto.CompactTextString(m) }
func (*StringListMsg) ProtoMessage()    {}
func (*StringListMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

func (m *StringListMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringListMsg.Unmarshal(m, b)
}
func (m *StringListMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringListMsg.Marshal(b, m, deterministic)
}
func (m *StringListMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringListMsg.Merge(m, src)
}
func (m *StringListMsg) XXX_Size() int {
	return xxx_messageInfo_StringListMsg.Size(m)
}
func (m *StringListMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StringListMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StringListMsg proto.InternalMessageInfo

func (m *StringListMsg) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// Connect request message
type ConnectReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addrs                []string `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectReq) Reset()         { *m = ConnectReq{} }
func (m *ConnectReq) String() string { return proto.CompactTextString(m) }
func (*ConnectReq) ProtoMessage()    {}
func (*ConnectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

func (m *ConnectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectReq.Unmarshal(m, b)
}
func (m *ConnectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectReq.Marshal(b, m, deterministic)
}
func (m *ConnectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectReq.Merge(m, src)
}
func (m *ConnectReq) XXX_Size() int {
	return xxx_messageInfo_ConnectReq.Size(m)
}
func (m *ConnectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectReq proto.InternalMessageInfo

func (m *ConnectReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConnectReq) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

// SendMsg request message
type SendMsgReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MsgType              string   `protobuf:"bytes,2,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsgReq) Reset()         { *m = SendMsgReq{} }
func (m *SendMsgReq) String() string { return proto.CompactTextString(m) }
func (*SendMsgReq) ProtoMessage()    {}
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}

func (m *SendMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsgReq.Unmarshal(m, b)
}
func (m *SendMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsgReq.Marshal(b, m, deterministic)
}
func (m *SendMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgReq.Merge(m, src)
}
func (m *SendMsgReq) XXX_Size() int {
	return xxx_messageInfo_SendMsgReq.Size(m)
}
func (m *SendMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgReq proto.InternalMessageInfo

func (m *SendMsgReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SendMsgReq) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *SendMsgReq) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// SendMsg response message
type SendMsgResp struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsgResp) Reset()         { *m = SendMsgResp{} }
func (m *SendMsgResp) String() string { return proto.CompactTextString(m) }
func (*SendMsgResp) ProtoMessage()    {}
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{5}
}

func (m *SendMsgResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsgResp.Unmarshal(m, b)
}
func (m *SendMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsgResp.Marshal(b, m, deterministic)
}
func (m *SendMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgResp.Merge(m, src)
}
func (m *SendMsgResp) XXX_Size() int {
	return xxx_messageInfo_SendMsgResp.Size(m)
}
func (m *SendMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgResp proto.InternalMessageInfo

func (m *SendMsgResp) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*StringMsg)(nil), "pb.StringMsg")
	proto.RegisterType((*StringListMsg)(nil), "pb.StringListMsg")
	proto.RegisterType((*ConnectReq)(nil), "pb.ConnectReq")
	proto.RegisterType((*SendMsgReq)(nil), "pb.SendMsgReq")
	proto.RegisterType((*SendMsgResp)(nil), "pb.SendMsgResp")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xed, 0x6a, 0xab, 0x40,
	0x14, 0x54, 0x83, 0x11, 0x4f, 0xbe, 0x6e, 0x96, 0xcb, 0x45, 0xf2, 0xe7, 0xda, 0x2d, 0x34, 0xfe,
	0x48, 0x2d, 0xd8, 0x27, 0x68, 0x92, 0x82, 0x85, 0x06, 0xc4, 0xb4, 0x0f, 0xa0, 0x71, 0x31, 0x0b,
	0x89, 0xbb, 0x75, 0xb7, 0x05, 0x1f, 0xb4, 0xef, 0x53, 0xfc, 0x48, 0x8c, 0x50, 0x0a, 0xfd, 0x77,
	0xc6, 0x33, 0x33, 0xee, 0x0c, 0x07, 0x06, 0xb2, 0xe0, 0x44, 0xb8, 0x3c, 0x67, 0x92, 0x21, 0x8d,
	0xc7, 0xd8, 0x00, 0xfd, 0xf1, 0xc8, 0x65, 0x81, 0xaf, 0xc0, 0xdc, 0xca, 0x9c, 0x66, 0xe9, 0x46,
	0xa4, 0xe8, 0x2f, 0xe8, 0x1f, 0xd1, 0xe1, 0x9d, 0x58, 0xaa, 0xad, 0x3a, 0x66, 0x58, 0x03, 0x3c,
	0x87, 0x51, 0x4d, 0x79, 0xa6, 0x42, 0x96, 0xb4, 0x7f, 0xd0, 0xaf, 0x36, 0xc2, 0x52, 0xed, 0x9e,
	0x63, 0x86, 0x0d, 0xc2, 0x1e, 0xc0, 0x8a, 0x65, 0x19, 0xd9, 0xc9, 0x90, 0xbc, 0xa1, 0x31, 0x68,
	0x34, 0x69, 0x9c, 0x34, 0x9a, 0x94, 0xe6, 0x51, 0x92, 0xe4, 0xc2, 0xd2, 0x2a, 0x51, 0x0d, 0xb0,
	0x0f, 0xb0, 0x25, 0x59, 0xb2, 0x11, 0xe9, 0x77, 0x1a, 0x0b, 0x8c, 0xa3, 0x48, 0x5f, 0x0a, 0x4e,
	0x2c, 0xad, 0xfa, 0x78, 0x82, 0xe8, 0x0f, 0xf4, 0x8e, 0x22, 0xb5, 0x7a, 0xb6, 0xea, 0x0c, 0xc3,
	0x72, 0xc4, 0xd7, 0x30, 0x38, 0x3b, 0x09, 0xde, 0xcd, 0x32, 0x6c, 0xb2, 0x78, 0x9f, 0x1a, 0x18,
	0x81, 0x17, 0xf8, 0x4c, 0x48, 0x64, 0x83, 0xf6, 0xb4, 0x46, 0xa6, 0xcb, 0x63, 0xb7, 0xea, 0x62,
	0x36, 0x2a, 0xc7, 0x73, 0x1b, 0x58, 0x41, 0x73, 0xd0, 0x1f, 0xca, 0x57, 0x5e, 0x92, 0xa6, 0x2d,
	0xa9, 0xe9, 0x03, 0x2b, 0xe8, 0x06, 0x8c, 0x26, 0x39, 0x1a, 0x97, 0xfb, 0xb6, 0x86, 0x59, 0x2b,
	0xc5, 0x0a, 0x72, 0x00, 0xd6, 0x54, 0x9c, 0xa8, 0xdd, 0xff, 0x75, 0x99, 0x0b, 0x30, 0x9a, 0x34,
	0xb5, 0x63, 0x5b, 0xd2, 0x6c, 0xd2, 0xc1, 0x82, 0x63, 0x05, 0xdd, 0xc2, 0x24, 0x24, 0x29, 0x15,
	0x92, 0xe4, 0x7e, 0x94, 0x25, 0x07, 0x92, 0xff, 0x68, 0x7e, 0x07, 0xd3, 0xd7, 0x2c, 0xff, 0x85,
	0xe0, 0x3f, 0xe8, 0xab, 0x03, 0x13, 0xe4, 0xb2, 0x88, 0x4b, 0xc2, 0x72, 0x01, 0x16, 0x65, 0x6e,
	0xc1, 0xa4, 0x8c, 0x76, 0xfb, 0x88, 0x66, 0x2e, 0xf7, 0xf8, 0x9e, 0x09, 0xe9, 0xf2, 0x78, 0x39,
	0x6c, 0x0a, 0x0f, 0xca, 0xeb, 0xf3, 0xd5, 0x40, 0x8d, 0xfb, 0xd5, 0x21, 0xde, 0x7f, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x3c, 0x00, 0x45, 0x8b, 0x97, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// P2PHostClient is the client API for P2PHost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type P2PHostClient interface {
	ID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringMsg, error)
	Addrs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListMsg, error)
	Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*Empty, error)
	DisConnect(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (*Empty, error)
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
	RegisterHandler(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (*Empty, error)
	UnregisterHandler(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (*Empty, error)
	Close(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type p2PHostClient struct {
	cc *grpc.ClientConn
}

func NewP2PHostClient(cc *grpc.ClientConn) P2PHostClient {
	return &p2PHostClient{cc}
}

func (c *p2PHostClient) ID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringMsg, error) {
	out := new(StringMsg)
	err := c.cc.Invoke(ctx, "/pb.P2PHost/ID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PHostClient) Addrs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListMsg, error) {
	out := new(StringListMsg)
	err := c.cc.Invoke(ctx, "/pb.P2PHost/Addrs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PHostClient) Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.P2PHost/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PHostClient) DisConnect(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.P2PHost/DisConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PHostClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	out := new(SendMsgResp)
	err := c.cc.Invoke(ctx, "/pb.P2PHost/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PHostClient) RegisterHandler(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.P2PHost/RegisterHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PHostClient) UnregisterHandler(ctx context.Context, in *StringMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.P2PHost/UnregisterHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PHostClient) Close(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.P2PHost/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// P2PHostServer is the server API for P2PHost service.
type P2PHostServer interface {
	ID(context.Context, *Empty) (*StringMsg, error)
	Addrs(context.Context, *Empty) (*StringListMsg, error)
	Connect(context.Context, *ConnectReq) (*Empty, error)
	DisConnect(context.Context, *StringMsg) (*Empty, error)
	SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error)
	RegisterHandler(context.Context, *StringMsg) (*Empty, error)
	UnregisterHandler(context.Context, *StringMsg) (*Empty, error)
	Close(context.Context, *Empty) (*Empty, error)
}

// UnimplementedP2PHostServer can be embedded to have forward compatible implementations.
type UnimplementedP2PHostServer struct {
}

func (*UnimplementedP2PHostServer) ID(ctx context.Context, req *Empty) (*StringMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ID not implemented")
}
func (*UnimplementedP2PHostServer) Addrs(ctx context.Context, req *Empty) (*StringListMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Addrs not implemented")
}
func (*UnimplementedP2PHostServer) Connect(ctx context.Context, req *ConnectReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedP2PHostServer) DisConnect(ctx context.Context, req *StringMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisConnect not implemented")
}
func (*UnimplementedP2PHostServer) SendMsg(ctx context.Context, req *SendMsgReq) (*SendMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (*UnimplementedP2PHostServer) RegisterHandler(ctx context.Context, req *StringMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterHandler not implemented")
}
func (*UnimplementedP2PHostServer) UnregisterHandler(ctx context.Context, req *StringMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterHandler not implemented")
}
func (*UnimplementedP2PHostServer) Close(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterP2PHostServer(s *grpc.Server, srv P2PHostServer) {
	s.RegisterService(&_P2PHost_serviceDesc, srv)
}

func _P2PHost_ID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PHostServer).ID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.P2PHost/ID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PHostServer).ID(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PHost_Addrs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PHostServer).Addrs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.P2PHost/Addrs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PHostServer).Addrs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PHost_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PHostServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.P2PHost/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PHostServer).Connect(ctx, req.(*ConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PHost_DisConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PHostServer).DisConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.P2PHost/DisConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PHostServer).DisConnect(ctx, req.(*StringMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PHost_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PHostServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.P2PHost/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PHostServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PHost_RegisterHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PHostServer).RegisterHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.P2PHost/RegisterHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PHostServer).RegisterHandler(ctx, req.(*StringMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PHost_UnregisterHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PHostServer).UnregisterHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.P2PHost/UnregisterHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PHostServer).UnregisterHandler(ctx, req.(*StringMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PHost_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PHostServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.P2PHost/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PHostServer).Close(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _P2PHost_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.P2PHost",
	HandlerType: (*P2PHostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ID",
			Handler:    _P2PHost_ID_Handler,
		},
		{
			MethodName: "Addrs",
			Handler:    _P2PHost_Addrs_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _P2PHost_Connect_Handler,
		},
		{
			MethodName: "DisConnect",
			Handler:    _P2PHost_DisConnect_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _P2PHost_SendMsg_Handler,
		},
		{
			MethodName: "RegisterHandler",
			Handler:    _P2PHost_RegisterHandler_Handler,
		},
		{
			MethodName: "UnregisterHandler",
			Handler:    _P2PHost_UnregisterHandler_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _P2PHost_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types.proto",
}
