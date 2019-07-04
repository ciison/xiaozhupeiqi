// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

package go_micro_srv_greeter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.greeter.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.greeter.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.greeter.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.greeter.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.greeter.Pong")
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor_e585294ab3f34af5) }

var fileDescriptor_e585294ab3f34af5 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3f, 0x4b, 0x03, 0x41,
	0x14, 0xc4, 0x73, 0xf9, 0xef, 0x03, 0x25, 0xae, 0x41, 0xf4, 0x4c, 0x82, 0x6c, 0xa1, 0xd1, 0xe2,
	0x2e, 0x68, 0x67, 0x6b, 0x91, 0x34, 0x82, 0x9c, 0xb5, 0xc5, 0x1a, 0x1e, 0xeb, 0xe1, 0xdd, 0xbe,
	0xb8, 0xbb, 0x11, 0x6c, 0xfd, 0x0a, 0x7e, 0x34, 0x7b, 0x2b, 0x3f, 0x88, 0xdc, 0xde, 0x06, 0x44,
	0xee, 0xb0, 0xb0, 0xdb, 0xc7, 0x6f, 0x66, 0x18, 0x86, 0x85, 0x6d, 0xa9, 0x11, 0x2d, 0xea, 0x68,
	0xa5, 0xc9, 0x12, 0x1b, 0x4a, 0x8a, 0xf2, 0x74, 0xa9, 0x29, 0x32, 0xfa, 0x25, 0xf2, 0x2c, 0x1c,
	0x49, 0x22, 0x99, 0x61, 0x2c, 0x56, 0x69, 0x2c, 0x94, 0x22, 0x2b, 0x6c, 0x4a, 0xca, 0x94, 0x1e,
	0x7e, 0x04, 0xbd, 0x1b, 0x34, 0x46, 0x48, 0x64, 0x03, 0x68, 0x19, 0xf1, 0x7a, 0x10, 0x1c, 0x07,
	0xd3, 0xad, 0xa4, 0x78, 0xf2, 0x31, 0xf4, 0x12, 0x7c, 0x5e, 0xa3, 0xb1, 0x8c, 0x41, 0x5b, 0x89,
	0x1c, 0x3d, 0x75, 0x6f, 0x3e, 0x82, 0x7e, 0x82, 0x66, 0x45, 0xca, 0x38, 0x73, 0x6e, 0xe4, 0xc6,
	0x9c, 0x1b, 0xc9, 0xa7, 0x30, 0xb8, 0xb3, 0x1a, 0x45, 0x9e, 0x2a, 0xb9, 0x49, 0x19, 0x42, 0x67,
	0x49, 0x6b, 0x65, 0x9d, 0xae, 0x95, 0x94, 0x07, 0x3f, 0x83, 0xdd, 0x1f, 0x4a, 0x1f, 0x58, 0x2d,
	0x9d, 0x40, 0xfb, 0x36, 0x55, 0x92, 0xed, 0x43, 0xd7, 0x58, 0x4d, 0x4f, 0xe8, 0xb1, 0xbf, 0x1c,
	0xa7, 0x7a, 0x7e, 0xf1, 0xd9, 0x84, 0xde, 0xbc, 0x1c, 0x86, 0xcd, 0xa1, 0x7d, 0x2d, 0xb2, 0x8c,
	0x8d, 0xa3, 0xaa, 0xdd, 0x22, 0xdf, 0x39, 0x9c, 0xd4, 0xe1, 0xb2, 0x28, 0x6f, 0x30, 0x01, 0x9d,
	0x05, 0x66, 0x19, 0xfd, 0x37, 0xe9, 0xf0, 0xed, 0xe3, 0xeb, 0xbd, 0xb9, 0xc7, 0x77, 0x62, 0x8f,
	0xe2, 0xc7, 0x22, 0xf6, 0x2a, 0x38, 0x67, 0xf7, 0xd0, 0x2d, 0x27, 0x62, 0x27, 0xd5, 0x21, 0xbf,
	0xa7, 0x0e, 0x4f, 0xff, 0xd4, 0x6d, 0xfa, 0xcf, 0x02, 0xb6, 0x80, 0x7e, 0x31, 0xab, 0x9b, 0x2e,
	0xac, 0x36, 0x16, 0x3c, 0xac, 0x63, 0xa4, 0x24, 0x6f, 0x4c, 0x83, 0x59, 0xf0, 0xd0, 0x75, 0xdf,
	0xea, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x89, 0x12, 0x27, 0x2d, 0x9b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Greeter_StreamClient, error)
	PingPong(ctx context.Context, opts ...grpc.CallOption) (Greeter_PingPongClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.greeter.Greeter/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.greeter.Greeter/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Greeter_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/go.micro.srv.greeter.Greeter/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_StreamClient interface {
	Recv() (*StreamingResponse, error)
	grpc.ClientStream
}

type greeterStreamClient struct {
	grpc.ClientStream
}

func (x *greeterStreamClient) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) PingPong(ctx context.Context, opts ...grpc.CallOption) (Greeter_PingPongClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[1], "/go.micro.srv.greeter.Greeter/PingPong", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterPingPongClient{stream}
	return x, nil
}

type Greeter_PingPongClient interface {
	Send(*Ping) error
	Recv() (*Pong, error)
	grpc.ClientStream
}

type greeterPingPongClient struct {
	grpc.ClientStream
}

func (x *greeterPingPongClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterPingPongClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	Call(context.Context, *Request) (*Response, error)
	Hello(context.Context, *Request) (*Response, error)
	Stream(*StreamingRequest, Greeter_StreamServer) error
	PingPong(Greeter_PingPongServer) error
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) Call(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedGreeterServer) Hello(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedGreeterServer) Stream(req *StreamingRequest, srv Greeter_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedGreeterServer) PingPong(srv Greeter_PingPongServer) error {
	return status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.greeter.Greeter/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.greeter.Greeter/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).Stream(m, &greeterStreamServer{stream})
}

type Greeter_StreamServer interface {
	Send(*StreamingResponse) error
	grpc.ServerStream
}

type greeterStreamServer struct {
	grpc.ServerStream
}

func (x *greeterStreamServer) Send(m *StreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_PingPong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).PingPong(&greeterPingPongServer{stream})
}

type Greeter_PingPongServer interface {
	Send(*Pong) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type greeterPingPongServer struct {
	grpc.ServerStream
}

func (x *greeterPingPongServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterPingPongServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.greeter.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Greeter_Call_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _Greeter_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Greeter_Stream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingPong",
			Handler:       _Greeter_PingPong_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greeter.proto",
}
