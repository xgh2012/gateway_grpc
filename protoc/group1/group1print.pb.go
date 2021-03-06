// Code generated by protoc-gen-go. DO NOT EDIT.
// source: group1/group1print.proto

package group1

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

type PrintGroup1Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintGroup1Request) Reset()         { *m = PrintGroup1Request{} }
func (m *PrintGroup1Request) String() string { return proto.CompactTextString(m) }
func (*PrintGroup1Request) ProtoMessage()    {}
func (*PrintGroup1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f54894c1759ae9, []int{0}
}

func (m *PrintGroup1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintGroup1Request.Unmarshal(m, b)
}
func (m *PrintGroup1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintGroup1Request.Marshal(b, m, deterministic)
}
func (m *PrintGroup1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintGroup1Request.Merge(m, src)
}
func (m *PrintGroup1Request) XXX_Size() int {
	return xxx_messageInfo_PrintGroup1Request.Size(m)
}
func (m *PrintGroup1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintGroup1Request.DiscardUnknown(m)
}

var xxx_messageInfo_PrintGroup1Request proto.InternalMessageInfo

func (m *PrintGroup1Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PrintGroup1Reply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintGroup1Reply) Reset()         { *m = PrintGroup1Reply{} }
func (m *PrintGroup1Reply) String() string { return proto.CompactTextString(m) }
func (*PrintGroup1Reply) ProtoMessage()    {}
func (*PrintGroup1Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f54894c1759ae9, []int{1}
}

func (m *PrintGroup1Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintGroup1Reply.Unmarshal(m, b)
}
func (m *PrintGroup1Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintGroup1Reply.Marshal(b, m, deterministic)
}
func (m *PrintGroup1Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintGroup1Reply.Merge(m, src)
}
func (m *PrintGroup1Reply) XXX_Size() int {
	return xxx_messageInfo_PrintGroup1Reply.Size(m)
}
func (m *PrintGroup1Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintGroup1Reply.DiscardUnknown(m)
}

var xxx_messageInfo_PrintGroup1Reply proto.InternalMessageInfo

func (m *PrintGroup1Reply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PrintGroup1Request)(nil), "group1.PrintGroup1Request")
	proto.RegisterType((*PrintGroup1Reply)(nil), "group1.PrintGroup1Reply")
}

func init() { proto.RegisterFile("group1/group1print.proto", fileDescriptor_11f54894c1759ae9) }

var fileDescriptor_11f54894c1759ae9 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x2f, 0xca, 0x2f,
	0x2d, 0x30, 0xd4, 0x87, 0x50, 0x05, 0x45, 0x99, 0x79, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x6c, 0x10, 0x21, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c,
	0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x2a, 0x25, 0x0d,
	0x2e, 0xa1, 0x00, 0x90, 0x26, 0x77, 0xb0, 0xe2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b,
	0x49, 0x87, 0x4b, 0x00, 0x45, 0x65, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71,
	0x71, 0x62, 0x3a, 0x4c, 0x29, 0x8c, 0x6b, 0x54, 0xc7, 0x25, 0xe4, 0x9e, 0x58, 0x92, 0x5a, 0x9e,
	0x58, 0x89, 0xa4, 0x49, 0x28, 0x83, 0x8b, 0x1b, 0x99, 0x2b, 0xa5, 0x07, 0x71, 0xa3, 0x1e, 0xa6,
	0x13, 0xa4, 0x24, 0xb0, 0xca, 0x15, 0xe4, 0x54, 0x2a, 0x29, 0x37, 0x5d, 0x7e, 0x32, 0x99, 0x49,
	0x56, 0x49, 0x42, 0xbf, 0xcc, 0x50, 0x3f, 0xb5, 0x22, 0x31, 0xb7, 0x20, 0x27, 0x15, 0xea, 0x7f,
	0x7d, 0x70, 0x00, 0x58, 0x31, 0x6a, 0x25, 0xb1, 0x81, 0xbd, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xd1, 0xfd, 0x04, 0xe7, 0x20, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GatewayPrintGroup1Client is the client API for GatewayPrintGroup1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayPrintGroup1Client interface {
	PrintGroup1(ctx context.Context, in *PrintGroup1Request, opts ...grpc.CallOption) (*PrintGroup1Reply, error)
}

type gatewayPrintGroup1Client struct {
	cc grpc.ClientConnInterface
}

func NewGatewayPrintGroup1Client(cc grpc.ClientConnInterface) GatewayPrintGroup1Client {
	return &gatewayPrintGroup1Client{cc}
}

func (c *gatewayPrintGroup1Client) PrintGroup1(ctx context.Context, in *PrintGroup1Request, opts ...grpc.CallOption) (*PrintGroup1Reply, error) {
	out := new(PrintGroup1Reply)
	err := c.cc.Invoke(ctx, "/group1.GatewayPrintGroup1/PrintGroup1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayPrintGroup1Server is the server API for GatewayPrintGroup1 service.
type GatewayPrintGroup1Server interface {
	PrintGroup1(context.Context, *PrintGroup1Request) (*PrintGroup1Reply, error)
}

// UnimplementedGatewayPrintGroup1Server can be embedded to have forward compatible implementations.
type UnimplementedGatewayPrintGroup1Server struct {
}

func (*UnimplementedGatewayPrintGroup1Server) PrintGroup1(ctx context.Context, req *PrintGroup1Request) (*PrintGroup1Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintGroup1 not implemented")
}

func RegisterGatewayPrintGroup1Server(s *grpc.Server, srv GatewayPrintGroup1Server) {
	s.RegisterService(&_GatewayPrintGroup1_serviceDesc, srv)
}

func _GatewayPrintGroup1_PrintGroup1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintGroup1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayPrintGroup1Server).PrintGroup1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group1.GatewayPrintGroup1/PrintGroup1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayPrintGroup1Server).PrintGroup1(ctx, req.(*PrintGroup1Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayPrintGroup1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "group1.GatewayPrintGroup1",
	HandlerType: (*GatewayPrintGroup1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrintGroup1",
			Handler:    _GatewayPrintGroup1_PrintGroup1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group1/group1print.proto",
}
