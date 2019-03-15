// Code generated by protoc-gen-go. DO NOT EDIT.
// source: epl/grpc/stac_operators.proto

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
	Insert(ctx context.Context, in *protobuf.StacMetadata, opts ...grpc.CallOption) (*protobuf.StacInsertResponse, error)
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

func (c *stacServiceClient) Insert(ctx context.Context, in *protobuf.StacMetadata, opts ...grpc.CallOption) (*protobuf.StacInsertResponse, error) {
	out := new(protobuf.StacInsertResponse)
	err := c.cc.Invoke(ctx, "/epl.grpc.StacService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StacServiceServer is the server API for StacService service.
type StacServiceServer interface {
	Search(*protobuf.StacSearchRequest, StacService_SearchServer) error
	Insert(context.Context, *protobuf.StacMetadata) (*protobuf.StacInsertResponse, error)
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

var _StacService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "epl.grpc.StacService",
	HandlerType: (*StacServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _StacService_Insert_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _StacService_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "epl/grpc/stac_operators.proto",
}

func init() {
	proto.RegisterFile("epl/grpc/stac_operators.proto", fileDescriptor_stac_operators_be6dc47f891ce110)
}

var fileDescriptor_stac_operators_be6dc47f891ce110 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xb7, 0x0a, 0x45, 0xa2, 0x07, 0xc9, 0x45, 0x08, 0x88, 0xb2, 0x17, 0x4f, 0x26, 0xa2,
	0x6f, 0xb0, 0x27, 0xf7, 0x20, 0xac, 0x1b, 0x4f, 0x5e, 0x64, 0x1a, 0xc7, 0x6c, 0xa1, 0xdb, 0x89,
	0xc9, 0xd4, 0x07, 0xea, 0x93, 0x4a, 0x52, 0x0a, 0x42, 0xc1, 0xeb, 0xfc, 0xdf, 0xc7, 0xcc, 0xfc,
	0xe2, 0x1a, 0x43, 0x67, 0x7c, 0x0c, 0xce, 0x24, 0x06, 0xf7, 0x41, 0x01, 0x23, 0x30, 0xc5, 0xa4,
	0x43, 0x24, 0x26, 0x79, 0x86, 0xa1, 0xd3, 0x39, 0x56, 0x57, 0x19, 0x2c, 0xc3, 0x66, 0xf8, 0x2a,
	0xf0, 0x84, 0x3c, 0x8e, 0x95, 0x38, 0xb7, 0x0c, 0xce, 0x62, 0xfc, 0x69, 0x1d, 0xca, 0xad, 0xa8,
	0x2d, 0x42, 0x74, 0x07, 0x79, 0xa3, 0xb3, 0x3d, 0x3b, 0x7a, 0x82, 0x72, 0xb2, 0xc7, 0xef, 0x01,
	0x13, 0x2b, 0xb5, 0x04, 0x5e, 0x90, 0xe1, 0x13, 0x18, 0xd6, 0xab, 0x87, 0x4a, 0x3e, 0x8b, 0x7a,
	0xdb, 0x27, 0x8c, 0x2c, 0xff, 0x21, 0xd5, 0xed, 0x32, 0x9b, 0xac, 0x3d, 0xa6, 0x40, 0x7d, 0xc2,
	0xf5, 0x6a, 0xf3, 0x2a, 0x2e, 0x1c, 0x1d, 0xf5, 0xfc, 0xcd, 0xe6, 0xf2, 0xcf, 0xc5, 0xbb, 0x6c,
	0xee, 0xaa, 0xf7, 0x3b, 0xdf, 0xf2, 0x61, 0x68, 0xb4, 0xa3, 0xa3, 0xf1, 0x48, 0xf7, 0xa5, 0x15,
	0x08, 0xad, 0xf1, 0xd4, 0x41, 0xef, 0xcd, 0xdc, 0xd4, 0x78, 0x72, 0x6a, 0xdf, 0x6c, 0x53, 0x97,
	0x8d, 0x4f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x9a, 0x25, 0x29, 0x42, 0x01, 0x00, 0x00,
}
