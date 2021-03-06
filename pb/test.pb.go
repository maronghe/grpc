// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package pb

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

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

type HelloRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type HelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "pb.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "pb.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloService service

type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SayBye(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/pb.HelloService/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) SayBye(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/pb.HelloService/SayBye", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloService service

type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	SayBye(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_SayBye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayBye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloService/SayBye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayBye(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
		{
			MethodName: "SayBye",
			Handler:    _HelloService_SayBye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe2, 0xe2, 0xf1, 0x48,
	0xcd, 0xc9, 0xc9, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0x2d,
	0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95,
	0x34, 0xb9, 0x78, 0xa1, 0x6a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x73,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x61, 0x6a, 0x61, 0x5c, 0xa3, 0x22, 0xa8, 0xb1, 0xc1, 0xa9, 0x45,
	0x65, 0x99, 0xc9, 0xa9, 0x42, 0x86, 0x5c, 0x1c, 0xc1, 0x89, 0x95, 0x60, 0x21, 0x21, 0x01, 0xbd,
	0x82, 0x24, 0x3d, 0x64, 0x4b, 0xa5, 0x04, 0x91, 0x44, 0x20, 0x46, 0x2b, 0x31, 0x08, 0xe9, 0x73,
	0xb1, 0x05, 0x27, 0x56, 0x3a, 0x55, 0xa6, 0x12, 0xa9, 0x21, 0x89, 0x0d, 0xec, 0x2b, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xc1, 0x61, 0x2d, 0xe3, 0x00, 0x00, 0x00,
}
