// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indicator.proto

/*
Package indicator is a generated protocol buffer package.

It is generated from these files:
	indicator.proto

It has these top-level messages:
	Request
	Response
*/
package indicator

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

type Request struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Label  string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	Icon   string `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	Active bool   `protobuf:"varint,4,opt,name=active" json:"active,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Request) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Request) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type Response struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "indicator.Request")
	proto.RegisterType((*Response)(nil), "indicator.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Shepherd service

type ShepherdClient interface {
	Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type shepherdClient struct {
	cc *grpc.ClientConn
}

func NewShepherdClient(cc *grpc.ClientConn) ShepherdClient {
	return &shepherdClient{cc}
}

func (c *shepherdClient) Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/indicator.Shepherd/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Shepherd service

type ShepherdServer interface {
	Update(context.Context, *Request) (*Response, error)
}

func RegisterShepherdServer(s *grpc.Server, srv ShepherdServer) {
	s.RegisterService(&_Shepherd_serviceDesc, srv)
}

func _Shepherd_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShepherdServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicator.Shepherd/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShepherdServer).Update(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shepherd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "indicator.Shepherd",
	HandlerType: (*ShepherdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Shepherd_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indicator.proto",
}

func init() { proto.RegisterFile("indicator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xbd, 0x0e, 0x82, 0x30,
	0x14, 0x85, 0xe5, 0x47, 0x84, 0x3b, 0xa8, 0xb9, 0x1a, 0xd3, 0x18, 0x07, 0xc2, 0xc4, 0xc4, 0x20,
	0x0f, 0xe0, 0x3b, 0xd4, 0x38, 0x39, 0x15, 0x7a, 0x13, 0x9a, 0x10, 0x5a, 0x4b, 0xf5, 0xf9, 0x8d,
	0x85, 0x68, 0xdc, 0xce, 0xf9, 0x86, 0x93, 0xef, 0xc0, 0x46, 0x0d, 0x52, 0xb5, 0xc2, 0x69, 0x5b,
	0x19, 0xab, 0x9d, 0xc6, 0xec, 0x0b, 0x8a, 0x3b, 0xac, 0x38, 0x3d, 0x9e, 0x34, 0x3a, 0x5c, 0x43,
	0xa8, 0x24, 0x0b, 0xf2, 0xa0, 0xcc, 0x78, 0xa8, 0x24, 0xee, 0x61, 0xd9, 0x8b, 0x86, 0x7a, 0x16,
	0x7a, 0x34, 0x15, 0x44, 0x88, 0x55, 0xab, 0x07, 0x16, 0x79, 0xe8, 0x33, 0x1e, 0x20, 0x11, 0xad,
	0x53, 0x2f, 0x62, 0x71, 0x1e, 0x94, 0x29, 0x9f, 0x5b, 0x71, 0x82, 0x94, 0xd3, 0x68, 0xf4, 0x30,
	0x12, 0x6e, 0x21, 0x22, 0x6b, 0xe7, 0xf9, 0x4f, 0x3c, 0x5f, 0x20, 0xbd, 0x76, 0x64, 0x3a, 0xb2,
	0x12, 0x6b, 0x48, 0x6e, 0x46, 0x0a, 0x47, 0x88, 0xd5, 0xcf, 0x76, 0x36, 0x3b, 0xee, 0xfe, 0xd8,
	0x34, 0x58, 0x2c, 0x9a, 0xc4, 0xbf, 0xa9, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0xe0, 0xa1,
	0x77, 0xe0, 0x00, 0x00, 0x00,
}
