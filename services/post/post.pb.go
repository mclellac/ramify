// Code generated by protoc-gen-go.
// source: post.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	post.proto

It has these top-level messages:
	Response
	Request
	Post
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Response struct {
	Message string `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=Url" json:"Url,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=Error" json:"Error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Post struct {
	Id      int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Created int32  `protobuf:"varint,2,opt,name=created" json:"created,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Article string `protobuf:"bytes,4,opt,name=article" json:"article,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto1.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Post) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetArticle() string {
	if m != nil {
		return m.Article
	}
	return ""
}

func init() {
	proto1.RegisterType((*Response)(nil), "proto.Response")
	proto1.RegisterType((*Request)(nil), "proto.Request")
	proto1.RegisterType((*Post)(nil), "proto.Post")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	Add(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *Request, opts ...grpc.CallOption) (Service_ListClient, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Add(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Service/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Delete(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Service/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) List(ctx context.Context, in *Request, opts ...grpc.CallOption) (Service_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Service_serviceDesc.Streams[0], c.cc, "/proto.Service/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_ListClient interface {
	Recv() (*Post, error)
	grpc.ClientStream
}

type serviceListClient struct {
	grpc.ClientStream
}

func (x *serviceListClient) Recv() (*Post, error) {
	m := new(Post)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Service service

type ServiceServer interface {
	Add(context.Context, *Post) (*Response, error)
	Delete(context.Context, *Post) (*Response, error)
	List(*Request, Service_ListServer) error
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Add(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Delete(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).List(m, &serviceListServer{stream})
}

type Service_ListServer interface {
	Send(*Post) error
	grpc.ServerStream
}

type serviceListServer struct {
	grpc.ServerStream
}

func (x *serviceListServer) Send(m *Post) error {
	return x.ServerStream.SendMsg(m)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Service_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Service_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Service_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "post.proto",
}

func init() { proto1.RegisterFile("post.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xcf, 0x4a, 0x04, 0x31,
	0x0c, 0x87, 0x77, 0xfe, 0xed, 0xb8, 0x11, 0x54, 0x82, 0x87, 0xb2, 0x27, 0x29, 0x88, 0x7b, 0x5a,
	0x44, 0x9f, 0x40, 0xd0, 0xdb, 0x0a, 0x52, 0xf1, 0xee, 0x38, 0x0d, 0x52, 0x18, 0xec, 0x98, 0x46,
	0x5f, 0xc0, 0x17, 0x97, 0x66, 0x76, 0xc4, 0xdb, 0x9e, 0xda, 0x2f, 0x4d, 0xbf, 0xe4, 0x07, 0x30,
	0xc6, 0x24, 0xdb, 0x91, 0xa3, 0x44, 0x6c, 0xf4, 0xb0, 0x3b, 0x38, 0x72, 0x94, 0xc6, 0xf8, 0x91,
	0x08, 0x0d, 0xb4, 0x8f, 0x94, 0x52, 0xf7, 0x4e, 0xa6, 0xb8, 0x28, 0x36, 0x2b, 0x37, 0x23, 0x9e,
	0x41, 0xf5, 0xc2, 0x83, 0x29, 0xb5, 0x9a, 0xaf, 0x78, 0x0e, 0xcd, 0x03, 0x73, 0x64, 0x53, 0x69,
	0x6d, 0x02, 0xbb, 0x82, 0xd6, 0xd1, 0xe7, 0x17, 0x25, 0xb1, 0xaf, 0x50, 0x3f, 0xc5, 0x24, 0x78,
	0x02, 0x65, 0xf0, 0xea, 0xab, 0x5c, 0x19, 0x7c, 0x1e, 0xd2, 0x33, 0x75, 0x42, 0x5e, 0x75, 0x8d,
	0x9b, 0x31, 0x2b, 0x25, 0xc8, 0x40, 0xb3, 0x52, 0x21, 0xf7, 0x77, 0x2c, 0xa1, 0x1f, 0xc8, 0xd4,
	0xd3, 0x52, 0x7b, 0xbc, 0xf9, 0x29, 0xa0, 0x7d, 0x26, 0xfe, 0x0e, 0x3d, 0xe1, 0x25, 0x54, 0x77,
	0xde, 0xe3, 0xf1, 0x14, 0x6e, 0x9b, 0x27, 0xaf, 0x4f, 0xf7, 0x30, 0xe7, 0xb3, 0x0b, 0xdc, 0xc0,
	0xf2, 0x9e, 0x06, 0x12, 0x3a, 0xd8, 0x79, 0x05, 0xf5, 0x2e, 0xe4, 0xf5, 0xff, 0x9e, 0x34, 0xd6,
	0xfa, 0xff, 0x3f, 0xbb, 0xb8, 0x2e, 0xde, 0x96, 0xca, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0x53, 0xdc, 0x39, 0x5c, 0x01, 0x00, 0x00,
}