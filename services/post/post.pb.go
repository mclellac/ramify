// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post.proto

/*
Package post is a generated protocol buffer package.

It is generated from these files:
	post.proto

It has these top-level messages:
	Response
	Request
	Content
*/
package post

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

type Response struct {
	Message string `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=Url" json:"Url,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=Error" json:"Error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
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
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Content struct {
	Id      int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Created int32  `protobuf:"varint,2,opt,name=created" json:"created,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Article string `protobuf:"bytes,4,opt,name=article" json:"article,omitempty"`
}

func (m *Content) Reset()                    { *m = Content{} }
func (m *Content) String() string            { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()               {}
func (*Content) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Content) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Content) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Content) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Content) GetArticle() string {
	if m != nil {
		return m.Article
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "post.Response")
	proto.RegisterType((*Request)(nil), "post.Request")
	proto.RegisterType((*Content)(nil), "post.Content")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Post service

type PostClient interface {
	Add(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *Request, opts ...grpc.CallOption) (Post_ListClient, error)
}

type postClient struct {
	cc *grpc.ClientConn
}

func NewPostClient(cc *grpc.ClientConn) PostClient {
	return &postClient{cc}
}

func (c *postClient) Add(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/post.Post/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) Delete(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/post.Post/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) List(ctx context.Context, in *Request, opts ...grpc.CallOption) (Post_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Post_serviceDesc.Streams[0], c.cc, "/post.Post/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &postListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Post_ListClient interface {
	Recv() (*Content, error)
	grpc.ClientStream
}

type postListClient struct {
	grpc.ClientStream
}

func (x *postListClient) Recv() (*Content, error) {
	m := new(Content)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Post service

type PostServer interface {
	Add(context.Context, *Content) (*Response, error)
	Delete(context.Context, *Content) (*Response, error)
	List(*Request, Post_ListServer) error
}

func RegisterPostServer(s *grpc.Server, srv PostServer) {
	s.RegisterService(&_Post_serviceDesc, srv)
}

func _Post_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.Post/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Add(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.Post/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Delete(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostServer).List(m, &postListServer{stream})
}

type Post_ListServer interface {
	Send(*Content) error
	grpc.ServerStream
}

type postListServer struct {
	grpc.ServerStream
}

func (x *postListServer) Send(m *Content) error {
	return x.ServerStream.SendMsg(m)
}

var _Post_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post.Post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Post_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Post_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Post_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "post.proto",
}

func init() { proto.RegisterFile("post.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xbb, 0x7f, 0xda, 0xb5, 0x03, 0x16, 0x19, 0x3c, 0x84, 0x9e, 0x24, 0x07, 0xa9, 0x97,
	0x22, 0xfa, 0x04, 0xa2, 0xde, 0x2a, 0x48, 0xc0, 0x07, 0x58, 0x37, 0x83, 0x04, 0xc2, 0x66, 0xcd,
	0x8c, 0x4f, 0xe0, 0x8b, 0x4b, 0xb2, 0x9b, 0x83, 0xa7, 0xde, 0xe6, 0x37, 0x7c, 0xf9, 0xbe, 0x7c,
	0x03, 0x30, 0x05, 0x96, 0xe3, 0x14, 0x83, 0x04, 0x6c, 0xd3, 0xac, 0x4f, 0x70, 0x61, 0x88, 0xa7,
	0x30, 0x32, 0xa1, 0x82, 0xee, 0x8d, 0x98, 0xfb, 0x2f, 0x52, 0xd5, 0x4d, 0x75, 0xd8, 0x9a, 0x82,
	0x78, 0x05, 0xcd, 0x47, 0xf4, 0xaa, 0xce, 0xdb, 0x34, 0xe2, 0x35, 0xac, 0x5f, 0x63, 0x0c, 0x51,
	0x35, 0x79, 0x37, 0x83, 0xde, 0x42, 0x67, 0xe8, 0xfb, 0x87, 0x58, 0xf4, 0x00, 0xdd, 0x73, 0x18,
	0x85, 0x46, 0xc1, 0x1d, 0xd4, 0xce, 0x66, 0xcb, 0xc6, 0xd4, 0xce, 0xa6, 0x9c, 0x21, 0x52, 0x2f,
	0x64, 0xb3, 0xe3, 0xda, 0x14, 0x4c, 0xae, 0xe2, 0xc4, 0x53, 0x71, 0xcd, 0x90, 0xf4, 0x7d, 0x14,
	0x37, 0x78, 0x52, 0xed, 0xfc, 0xaf, 0x05, 0x1f, 0x7e, 0x2b, 0x68, 0xdf, 0x03, 0x0b, 0xde, 0x42,
	0xf3, 0x64, 0x2d, 0x5e, 0x1e, 0x73, 0xc1, 0x25, 0x78, 0xbf, 0x9b, 0xb1, 0x14, 0xd4, 0x2b, 0xbc,
	0x83, 0xcd, 0x0b, 0x79, 0x12, 0x3a, 0x2f, 0x3d, 0x40, 0x7b, 0x72, 0x2c, 0x45, 0xb8, 0xf4, 0xda,
	0xff, 0x7f, 0xa7, 0x57, 0xf7, 0xd5, 0xe7, 0x26, 0x1f, 0xf4, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x79, 0xe2, 0xb8, 0x52, 0x5e, 0x01, 0x00, 0x00,
}
