// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conversor.proto

package conversor

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

// Mensagem de Request
type Request struct {
	MoedaDestino         string   `protobuf:"bytes,1,opt,name=MoedaDestino,json=moedaDestino,proto3" json:"MoedaDestino,omitempty"`
	Valor                float32  `protobuf:"fixed32,2,opt,name=Valor,json=valor,proto3" json:"Valor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ab7aa8c077c8dba, []int{0}
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

func (m *Request) GetMoedaDestino() string {
	if m != nil {
		return m.MoedaDestino
	}
	return ""
}

func (m *Request) GetValor() float32 {
	if m != nil {
		return m.Valor
	}
	return 0
}

//Mensagem de resposta
type Reply struct {
	Resultado            float32  `protobuf:"fixed32,1,opt,name=resultado,proto3" json:"resultado,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ab7aa8c077c8dba, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetResultado() float32 {
	if m != nil {
		return m.Resultado
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "conversor.Request")
	proto.RegisterType((*Reply)(nil), "conversor.Reply")
}

func init() { proto.RegisterFile("conversor.proto", fileDescriptor_4ab7aa8c077c8dba) }

var fileDescriptor_4ab7aa8c077c8dba = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0xcf, 0x2b,
	0x4b, 0x2d, 0x2a, 0xce, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x39, 0x73, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x29, 0x71, 0xf1, 0xf8, 0xe6,
	0xa7, 0xa6, 0x24, 0xba, 0xa4, 0x16, 0x97, 0x64, 0xe6, 0xe5, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0xf1, 0xe4, 0x22, 0x89, 0x09, 0x89, 0x70, 0xb1, 0x86, 0x25, 0xe6, 0xe4, 0x17, 0x49, 0x30,
	0x29, 0x30, 0x6a, 0x30, 0x05, 0xb1, 0x96, 0x81, 0x38, 0x4a, 0xaa, 0x5c, 0xac, 0x41, 0xa9, 0x05,
	0x39, 0x95, 0x42, 0x32, 0x5c, 0x9c, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x89, 0x29, 0x10, 0xfd,
	0x4c, 0x41, 0x08, 0x01, 0x23, 0x3b, 0x2e, 0x76, 0xf7, 0xa2, 0xd4, 0xd4, 0x92, 0xd4, 0x22, 0x21,
	0x63, 0x2e, 0xa8, 0x1b, 0x40, 0x1c, 0x21, 0x3d, 0x84, 0x03, 0xa1, 0x8e, 0x91, 0x12, 0x40, 0x11,
	0x2b, 0xc8, 0xa9, 0x54, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xa7, 0xe9, 0x62, 0x78, 0xd0, 0x00, 0x00, 0x00,
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
	Converter(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Converter(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/conversor.Greeter/converter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	Converter(context.Context, *Request) (*Reply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) Converter(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Converter not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Converter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Converter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conversor.Greeter/Converter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Converter(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "conversor.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "converter",
			Handler:    _Greeter_Converter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversor.proto",
}