// Code generated by protoc-gen-go. DO NOT EDIT.
// source: epl/grpc/stac_service.proto

package grpc // import "github.com/geo-grpc/api/golang/epl/grpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protobuf "github.com/geo-grpc/api/golang/epl/protobuf"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StacServiceClient is the client API for StacService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StacServiceClient interface {
	Search(ctx context.Context, in *protobuf.StacSearchRequest, opts ...grpc.CallOption) (StacService_SearchClient, error)
	Insert(ctx context.Context, in *protobuf.StacMetadata, opts ...grpc.CallOption) (*protobuf.StacUpsertResponse, error)
	Update(ctx context.Context, in *protobuf.StacMetadata, opts ...grpc.CallOption) (*protobuf.StacUpsertResponse, error)
}

type stacServiceClient struct {
	cc *grpc.ClientConn
}

func NewStacServiceClient(cc *grpc.ClientConn) StacServiceClient {
	return &stacServiceClient{cc}
}

func (c *stacServiceClient) Search(ctx context.Context, in *protobuf.StacSearchRequest, opts ...grpc.CallOption) (StacService_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StacService_serviceDesc.Streams[0], "/epl.grpc.StacService/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &stacServiceSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StacService_SearchClient interface {
	Recv() (*protobuf.StacMetadata, error)
	grpc.ClientStream
}

type stacServiceSearchClient struct {
	grpc.ClientStream
}

func (x *stacServiceSearchClient) Recv() (*protobuf.StacMetadata, error) {
	m := new(protobuf.StacMetadata)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stacServiceClient) Insert(ctx context.Context, in *protobuf.StacMetadata, opts ...grpc.CallOption) (*protobuf.StacUpsertResponse, error) {
	out := new(protobuf.StacUpsertResponse)
	err := c.cc.Invoke(ctx, "/epl.grpc.StacService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacServiceClient) Update(ctx context.Context, in *protobuf.StacMetadata, opts ...grpc.CallOption) (*protobuf.StacUpsertResponse, error) {
	out := new(protobuf.StacUpsertResponse)
	err := c.cc.Invoke(ctx, "/epl.grpc.StacService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StacServiceServer is the server API for StacService service.
type StacServiceServer interface {
	Search(*protobuf.StacSearchRequest, StacService_SearchServer) error
	Insert(context.Context, *protobuf.StacMetadata) (*protobuf.StacUpsertResponse, error)
	Update(context.Context, *protobuf.StacMetadata) (*protobuf.StacUpsertResponse, error)
}

func RegisterStacServiceServer(s *grpc.Server, srv StacServiceServer) {
	s.RegisterService(&_StacService_serviceDesc, srv)
}

func _StacService_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(protobuf.StacSearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StacServiceServer).Search(m, &stacServiceSearchServer{stream})
}

type StacService_SearchServer interface {
	Send(*protobuf.StacMetadata) error
	grpc.ServerStream
}

type stacServiceSearchServer struct {
	grpc.ServerStream
}

func (x *stacServiceSearchServer) Send(m *protobuf.StacMetadata) error {
	return x.ServerStream.SendMsg(m)
}

func _StacService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.StacMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/epl.grpc.StacService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacServiceServer).Insert(ctx, req.(*protobuf.StacMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _StacService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.StacMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/epl.grpc.StacService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacServiceServer).Update(ctx, req.(*protobuf.StacMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

var _StacService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "epl.grpc.StacService",
	HandlerType: (*StacServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _StacService_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StacService_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _StacService_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "epl/grpc/stac_service.proto",
}

func init() {
	proto.RegisterFile("epl/grpc/stac_service.proto", fileDescriptor_stac_service_9953d2d253cbe9e0)
}

var fileDescriptor_stac_service_9953d2d253cbe9e0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0x6f, 0x15, 0x16, 0x89, 0x16, 0x92, 0x46, 0x88, 0x85, 0x72, 0x8d, 0x95, 0x89, 0xe8,
	0x1b, 0x5c, 0xe5, 0x15, 0xc2, 0x79, 0xf1, 0x1a, 0x1b, 0x99, 0xcd, 0x8d, 0xb9, 0x85, 0xbd, 0xcd,
	0x98, 0xcc, 0xfa, 0x40, 0x3e, 0x9f, 0x0f, 0x21, 0xc9, 0xb2, 0x20, 0x1c, 0xd8, 0xd8, 0xe6, 0xff,
	0xbe, 0x3f, 0xcc, 0x2f, 0x2e, 0x91, 0x3a, 0xe3, 0x23, 0x39, 0x93, 0x18, 0xdc, 0x5b, 0xc2, 0xf8,
	0xd9, 0x3a, 0xd4, 0x14, 0x03, 0x07, 0x79, 0x82, 0xd4, 0xe9, 0x1c, 0xaa, 0x8b, 0x8c, 0x95, 0xc7,
	0x66, 0x78, 0x2f, 0xe8, 0x88, 0xdc, 0x7f, 0x57, 0xe2, 0xd4, 0x32, 0x38, 0x3b, 0x8a, 0x72, 0x29,
	0x6a, 0x8b, 0x10, 0xdd, 0x4e, 0x5e, 0xe9, 0x6c, 0x4f, 0x8e, 0x1e, 0xa1, 0x9c, 0xac, 0xf1, 0x63,
	0xc0, 0xc4, 0x4a, 0x1d, 0x02, 0x4f, 0xc8, 0xb0, 0x05, 0x86, 0xf9, 0xec, 0xae, 0x92, 0x8f, 0xa2,
	0x5e, 0xf6, 0x09, 0x23, 0xcb, 0x3f, 0x48, 0x75, 0x7d, 0x98, 0x6d, 0x28, 0x5b, 0x6b, 0x4c, 0x14,
	0xfa, 0x84, 0xf3, 0x59, 0x6e, 0xda, 0xd0, 0x16, 0x18, 0xff, 0xdb, 0xb4, 0x78, 0x16, 0x67, 0x2e,
	0xec, 0xf5, 0xb4, 0xcb, 0xe2, 0xfc, 0xd7, 0xed, 0xab, 0x6c, 0xae, 0xaa, 0xd7, 0x1b, 0xdf, 0xf2,
	0x6e, 0x68, 0xb4, 0x0b, 0x7b, 0xe3, 0x31, 0xdc, 0x96, 0x75, 0x81, 0x5a, 0xe3, 0x43, 0x07, 0xbd,
	0x37, 0xd3, 0xe2, 0x5f, 0x47, 0xc7, 0xf6, 0xc5, 0x36, 0x75, 0xf9, 0xf1, 0xe1, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0x49, 0xf2, 0xc8, 0x8a, 0x01, 0x00, 0x00,
}
