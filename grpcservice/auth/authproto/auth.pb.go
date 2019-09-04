// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package authproto

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

type AuthReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthReq) Reset()         { *m = AuthReq{} }
func (m *AuthReq) String() string { return proto.CompactTextString(m) }
func (*AuthReq) ProtoMessage()    {}
func (*AuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *AuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReq.Unmarshal(m, b)
}
func (m *AuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReq.Marshal(b, m, deterministic)
}
func (m *AuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReq.Merge(m, src)
}
func (m *AuthReq) XXX_Size() int {
	return xxx_messageInfo_AuthReq.Size(m)
}
func (m *AuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReq proto.InternalMessageInfo

func (m *AuthReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AuthRsp struct {
	Errorid              int32    `protobuf:"varint,1,opt,name=errorid,proto3" json:"errorid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Userid               int32    `protobuf:"varint,3,opt,name=userid,proto3" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRsp) Reset()         { *m = AuthRsp{} }
func (m *AuthRsp) String() string { return proto.CompactTextString(m) }
func (*AuthRsp) ProtoMessage()    {}
func (*AuthRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRsp.Unmarshal(m, b)
}
func (m *AuthRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRsp.Marshal(b, m, deterministic)
}
func (m *AuthRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRsp.Merge(m, src)
}
func (m *AuthRsp) XXX_Size() int {
	return xxx_messageInfo_AuthRsp.Size(m)
}
func (m *AuthRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRsp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRsp proto.InternalMessageInfo

func (m *AuthRsp) GetErrorid() int32 {
	if m != nil {
		return m.Errorid
	}
	return 0
}

func (m *AuthRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthRsp) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthReq)(nil), "authproto.AuthReq")
	proto.RegisterType((*AuthRsp)(nil), "authproto.AuthRsp")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xb1, 0xc1, 0x4c, 0x25, 0x59, 0x2e, 0x76,
	0xc7, 0xd2, 0x92, 0x8c, 0xa0, 0xd4, 0x42, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x1f, 0x2a, 0x5d, 0x5c, 0x20, 0x24, 0xc1,
	0xc5, 0x9e, 0x5a, 0x54, 0x94, 0x5f, 0x94, 0x99, 0x02, 0x56, 0xc1, 0x1a, 0x04, 0xe3, 0xc2, 0x35,
	0x32, 0x21, 0x34, 0x0a, 0x89, 0x71, 0xb1, 0x95, 0x16, 0xa7, 0x82, 0x14, 0x33, 0x83, 0x15, 0x43,
	0x79, 0x46, 0x25, 0x5c, 0xdc, 0x20, 0x03, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x0c,
	0xb8, 0x58, 0x40, 0x5c, 0x21, 0x21, 0x3d, 0xb8, 0x93, 0xf4, 0xa0, 0xee, 0x91, 0xc2, 0x10, 0x2b,
	0x2e, 0x50, 0x62, 0x10, 0x32, 0x86, 0xb8, 0xc8, 0x31, 0x25, 0x85, 0x78, 0x4d, 0x49, 0x6c, 0x60,
	0x01, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xcb, 0xcd, 0x51, 0x05, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRsp, error)
	AuthAdd(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRsp, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRsp, error) {
	out := new(AuthRsp)
	err := c.cc.Invoke(ctx, "/authproto.AuthService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthAdd(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRsp, error) {
	out := new(AuthRsp)
	err := c.cc.Invoke(ctx, "/authproto.AuthService/AuthAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Auth(context.Context, *AuthReq) (*AuthRsp, error)
	AuthAdd(context.Context, *AuthReq) (*AuthRsp, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Auth(ctx context.Context, req *AuthReq) (*AuthRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedAuthServiceServer) AuthAdd(ctx context.Context, req *AuthReq) (*AuthRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthAdd not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Auth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthService/AuthAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthAdd(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authproto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _AuthService_Auth_Handler,
		},
		{
			MethodName: "AuthAdd",
			Handler:    _AuthService_AuthAdd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}