// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store-service.proto

/*
Package store_service is a generated protocol buffer package.

It is generated from these files:
	store-service.proto

It has these top-level messages:
	StoreRequest
	StoreResponse
*/
package store_service

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

type StoreRequest struct {
	Crawled   string   `protobuf:"bytes,1,opt,name=crawled" json:"crawled,omitempty"`
	Collected []string `protobuf:"bytes,2,rep,name=collected" json:"collected,omitempty"`
}

func (m *StoreRequest) Reset()                    { *m = StoreRequest{} }
func (m *StoreRequest) String() string            { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()               {}
func (*StoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StoreRequest) GetCrawled() string {
	if m != nil {
		return m.Crawled
	}
	return ""
}

func (m *StoreRequest) GetCollected() []string {
	if m != nil {
		return m.Collected
	}
	return nil
}

type StoreResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *StoreResponse) Reset()                    { *m = StoreResponse{} }
func (m *StoreResponse) String() string            { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()               {}
func (*StoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StoreResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*StoreRequest)(nil), "StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "StoreResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StoreService service

type StoreServiceClient interface {
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
}

type storeServiceClient struct {
	cc *grpc.ClientConn
}

func NewStoreServiceClient(cc *grpc.ClientConn) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := grpc.Invoke(ctx, "/StoreService/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StoreService service

type StoreServiceServer interface {
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StoreService/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _StoreService_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store-service.proto",
}

func init() { proto.RegisterFile("store-service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0xc9, 0x2f,
	0x4a, 0xd5, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57,
	0x72, 0xe3, 0xe2, 0x09, 0x06, 0x09, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70,
	0xb1, 0x27, 0x17, 0x25, 0x96, 0xe7, 0xa4, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1,
	0xb8, 0x42, 0x32, 0x5c, 0x9c, 0xc9, 0xf9, 0x39, 0x39, 0xa9, 0xc9, 0x25, 0xa9, 0x29, 0x12, 0x4c,
	0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x25, 0x4d, 0x2e, 0x5e, 0xa8, 0x39, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0x20, 0x83, 0x8a, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0xc1, 0x06, 0x71, 0x04,
	0xc1, 0xb8, 0x46, 0x66, 0x50, 0x2b, 0x83, 0x21, 0x0e, 0x11, 0x52, 0xe3, 0x62, 0x05, 0xf3, 0x85,
	0x78, 0xf5, 0x90, 0x9d, 0x22, 0xc5, 0xa7, 0x87, 0x62, 0x62, 0x12, 0x1b, 0xd8, 0xc5, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x23, 0xc9, 0x1c, 0xc8, 0x00, 0x00, 0x00,
}
