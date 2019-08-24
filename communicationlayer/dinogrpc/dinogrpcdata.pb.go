// Code generated by protoc-gen-go.
// source: dinogrpcdata.proto
// DO NOT EDIT!

/*
Package dinogrpc is a generated protocol buffer package.

It is generated from these files:
	dinogrpcdata.proto

It has these top-level messages:
	Animal
	Request
*/
package dinogrpc

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

type Animal struct {
	Id         int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	AnimalType string `protobuf:"bytes,2,opt,name=animal_type,json=animalType" json:"animal_type,omitempty"`
	Nickname   string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	Zone       int32  `protobuf:"varint,4,opt,name=zone" json:"zone,omitempty"`
	Age        int32  `protobuf:"varint,5,opt,name=age" json:"age,omitempty"`
}

func (m *Animal) Reset()                    { *m = Animal{} }
func (m *Animal) String() string            { return proto.CompactTextString(m) }
func (*Animal) ProtoMessage()               {}
func (*Animal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Animal) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Animal) GetAnimalType() string {
	if m != nil {
		return m.AnimalType
	}
	return ""
}

func (m *Animal) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Animal) GetZone() int32 {
	if m != nil {
		return m.Zone
	}
	return 0
}

func (m *Animal) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type Request struct {
	Nickname string `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func init() {
	proto.RegisterType((*Animal)(nil), "dinogrpc.animal")
	proto.RegisterType((*Request)(nil), "dinogrpc.request")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DinoService service

type DinoServiceClient interface {
	GetAnimal(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Animal, error)
	GetAllAnimals(ctx context.Context, in *Request, opts ...grpc.CallOption) (DinoService_GetAllAnimalsClient, error)
}

type dinoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDinoServiceClient(cc *grpc.ClientConn) DinoServiceClient {
	return &dinoServiceClient{cc}
}

func (c *dinoServiceClient) GetAnimal(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Animal, error) {
	out := new(Animal)
	err := grpc.Invoke(ctx, "/dinogrpc.DinoService/GetAnimal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dinoServiceClient) GetAllAnimals(ctx context.Context, in *Request, opts ...grpc.CallOption) (DinoService_GetAllAnimalsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DinoService_serviceDesc.Streams[0], c.cc, "/dinogrpc.DinoService/GetAllAnimals", opts...)
	if err != nil {
		return nil, err
	}
	x := &dinoServiceGetAllAnimalsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DinoService_GetAllAnimalsClient interface {
	Recv() (*Animal, error)
	grpc.ClientStream
}

type dinoServiceGetAllAnimalsClient struct {
	grpc.ClientStream
}

func (x *dinoServiceGetAllAnimalsClient) Recv() (*Animal, error) {
	m := new(Animal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DinoService service

type DinoServiceServer interface {
	GetAnimal(context.Context, *Request) (*Animal, error)
	GetAllAnimals(*Request, DinoService_GetAllAnimalsServer) error
}

func RegisterDinoServiceServer(s *grpc.Server, srv DinoServiceServer) {
	fmt.Println("do we fet ths far")
	s.RegisterService(&_DinoService_serviceDesc, srv)
}

func _DinoService_GetAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DinoServiceServer).GetAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dinogrpc.DinoService/GetAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DinoServiceServer).GetAnimal(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DinoService_GetAllAnimals_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DinoServiceServer).GetAllAnimals(m, &dinoServiceGetAllAnimalsServer{stream})
}

type DinoService_GetAllAnimalsServer interface {
	Send(*Animal) error
	grpc.ServerStream
}

type dinoServiceGetAllAnimalsServer struct {
	grpc.ServerStream
}

func (x *dinoServiceGetAllAnimalsServer) Send(m *Animal) error {
	return x.ServerStream.SendMsg(m)
}

var _DinoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dinogrpc.DinoService",
	HandlerType: (*DinoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnimal",
			Handler:    _DinoService_GetAnimal_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllAnimals",
			Handler:       _DinoService_GetAllAnimals_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dinogrpcdata.proto",
}

func init() { proto.RegisterFile("dinogrpcdata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x49, 0xf7, 0x87, 0xbb, 0xb3, 0x28, 0xeb, 0x9c, 0xc2, 0x5e, 0x5c, 0x0a, 0x42, 0x4f,
	0xa5, 0x28, 0x78, 0x17, 0x04, 0xef, 0xd5, 0xbb, 0xc4, 0x66, 0x28, 0xc1, 0x36, 0x89, 0x69, 0x54,
	0xaa, 0xff, 0xbc, 0x34, 0x69, 0x95, 0xde, 0xf6, 0x36, 0xf9, 0x86, 0xf7, 0x91, 0x37, 0x80, 0x52,
	0x69, 0x53, 0x3b, 0x5b, 0x49, 0xe1, 0x45, 0x6e, 0x9d, 0xf1, 0x06, 0x37, 0x13, 0x4b, 0x7f, 0x60,
	0x2d, 0xb4, 0x6a, 0x45, 0x83, 0x17, 0x90, 0x28, 0xc9, 0xd9, 0x91, 0x65, 0xab, 0x32, 0x51, 0x12,
	0xaf, 0x60, 0x17, 0x37, 0x2f, 0xbe, 0xb7, 0xc4, 0x93, 0x23, 0xcb, 0xb6, 0x25, 0x44, 0xf4, 0xdc,
	0x5b, 0xc2, 0x03, 0x6c, 0xb4, 0xaa, 0xde, 0xb4, 0x68, 0x89, 0x2f, 0xc2, 0xf6, 0xef, 0x8d, 0x08,
	0xcb, 0x6f, 0xa3, 0x89, 0x2f, 0x83, 0x2e, 0xcc, 0xb8, 0x87, 0x85, 0xa8, 0x89, 0xaf, 0x02, 0x1a,
	0xc6, 0xf4, 0x1a, 0xce, 0x1c, 0xbd, 0x7f, 0x50, 0xe7, 0x67, 0x32, 0x36, 0x97, 0xdd, 0x7c, 0xc1,
	0xee, 0x41, 0x69, 0xf3, 0x44, 0xee, 0x53, 0x55, 0x84, 0x05, 0x6c, 0x1f, 0xc9, 0xdf, 0xc7, 0x5f,
	0x5f, 0xe6, 0x53, 0x95, 0x7c, 0x54, 0x1d, 0xf6, 0xff, 0x68, 0xac, 0x76, 0x07, 0xe7, 0x43, 0xa2,
	0x69, 0x62, 0xa8, 0x3b, 0x29, 0x55, 0xb0, 0xd7, 0x75, 0xb8, 0xd6, 0xed, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x41, 0xc2, 0x4c, 0xd9, 0x43, 0x01, 0x00, 0x00,
}
