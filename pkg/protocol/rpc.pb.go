// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package protocol

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

type RegularReply struct {
	Messages             []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegularReply) Reset()         { *m = RegularReply{} }
func (m *RegularReply) String() string { return proto.CompactTextString(m) }
func (*RegularReply) ProtoMessage()    {}
func (*RegularReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *RegularReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegularReply.Unmarshal(m, b)
}
func (m *RegularReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegularReply.Marshal(b, m, deterministic)
}
func (m *RegularReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegularReply.Merge(m, src)
}
func (m *RegularReply) XXX_Size() int {
	return xxx_messageInfo_RegularReply.Size(m)
}
func (m *RegularReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegularReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegularReply proto.InternalMessageInfo

func (m *RegularReply) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *RegularReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SetupRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	SpecsRepo            string   `protobuf:"bytes,2,opt,name=specsRepo,proto3" json:"specsRepo,omitempty"`
	VmMemory             int32    `protobuf:"varint,3,opt,name=vmMemory,proto3" json:"vmMemory,omitempty"`
	SshKeyPath           string   `protobuf:"bytes,4,opt,name=sshKeyPath,proto3" json:"sshKeyPath,omitempty"`
	Update               bool     `protobuf:"varint,5,opt,name=update,proto3" json:"update,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetupRequest) Reset()         { *m = SetupRequest{} }
func (m *SetupRequest) String() string { return proto.CompactTextString(m) }
func (*SetupRequest) ProtoMessage()    {}
func (*SetupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *SetupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetupRequest.Unmarshal(m, b)
}
func (m *SetupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetupRequest.Marshal(b, m, deterministic)
}
func (m *SetupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetupRequest.Merge(m, src)
}
func (m *SetupRequest) XXX_Size() int {
	return xxx_messageInfo_SetupRequest.Size(m)
}
func (m *SetupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetupRequest proto.InternalMessageInfo

func (m *SetupRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SetupRequest) GetSpecsRepo() string {
	if m != nil {
		return m.SpecsRepo
	}
	return ""
}

func (m *SetupRequest) GetVmMemory() int32 {
	if m != nil {
		return m.VmMemory
	}
	return 0
}

func (m *SetupRequest) GetSshKeyPath() string {
	if m != nil {
		return m.SshKeyPath
	}
	return ""
}

func (m *SetupRequest) GetUpdate() bool {
	if m != nil {
		return m.Update
	}
	return false
}

func init() {
	proto.RegisterType((*RegularReply)(nil), "protocol.RegularReply")
	proto.RegisterType((*SetupRequest)(nil), "protocol.SetupRequest")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x8d, 0x6b, 0x96, 0x76, 0xd8, 0x53, 0x90, 0x25, 0x2c, 0x22, 0xa5, 0xa7, 0x9e, 0x7a,
	0xd0, 0x93, 0x37, 0x11, 0x6f, 0x22, 0xc8, 0xf8, 0x0b, 0x62, 0x1d, 0x76, 0x0f, 0xad, 0x89, 0x33,
	0x89, 0xd0, 0xff, 0xe2, 0x8f, 0x95, 0x8d, 0x6b, 0xcd, 0x29, 0x7c, 0x2f, 0xf3, 0x1e, 0x6f, 0x06,
	0x6a, 0x0e, 0x43, 0x1f, 0xd8, 0x47, 0x6f, 0xaa, 0xfc, 0x0c, 0x7e, 0x6c, 0xef, 0x61, 0x83, 0xb4,
	0x4f, 0xa3, 0x63, 0xa4, 0x30, 0xce, 0x66, 0x07, 0xd5, 0x44, 0x22, 0x6e, 0x4f, 0x62, 0x55, 0xb3,
	0xea, 0x6a, 0x5c, 0xd8, 0x5c, 0x82, 0x26, 0x66, 0xcf, 0xf6, 0xbc, 0x51, 0x5d, 0x8d, 0xbf, 0xd0,
	0x7e, 0x2b, 0xd8, 0xbc, 0x52, 0x4c, 0x01, 0xe9, 0x33, 0x91, 0xc4, 0x63, 0x44, 0x12, 0xe2, 0x0f,
	0x37, 0x91, 0x55, 0x79, 0x72, 0x61, 0x73, 0x05, 0xb5, 0x04, 0x1a, 0x04, 0x29, 0xf8, 0x53, 0xcc,
	0xbf, 0x70, 0x74, 0x7e, 0x4d, 0xcf, 0x34, 0x79, 0x9e, 0xed, 0xaa, 0x51, 0x9d, 0xc6, 0x85, 0xcd,
	0x35, 0x80, 0xc8, 0xe1, 0x89, 0xe6, 0x17, 0x17, 0x0f, 0xf6, 0x22, 0x5b, 0x0b, 0xc5, 0x6c, 0x61,
	0x9d, 0xc2, 0xbb, 0x8b, 0x64, 0x75, 0xa3, 0xba, 0x0a, 0x4f, 0x74, 0xf3, 0x00, 0xfa, 0x31, 0x49,
	0x9c, 0xcd, 0x1d, 0xe8, 0x5c, 0xd3, 0x6c, 0xfb, 0xbf, 0xed, 0xfb, 0xb2, 0xf7, 0xae, 0xd0, 0xcb,
	0x93, 0xb4, 0x67, 0x6f, 0xeb, 0xfc, 0x71, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x64, 0x25, 0x40,
	0x83, 0x42, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DustyClient is the client API for Dusty service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DustyClient interface {
	Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*RegularReply, error)
}

type dustyClient struct {
	cc *grpc.ClientConn
}

func NewDustyClient(cc *grpc.ClientConn) DustyClient {
	return &dustyClient{cc}
}

func (c *dustyClient) Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*RegularReply, error) {
	out := new(RegularReply)
	err := c.cc.Invoke(ctx, "/protocol.Dusty/Setup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DustyServer is the server API for Dusty service.
type DustyServer interface {
	Setup(context.Context, *SetupRequest) (*RegularReply, error)
}

// UnimplementedDustyServer can be embedded to have forward compatible implementations.
type UnimplementedDustyServer struct {
}

func (*UnimplementedDustyServer) Setup(ctx context.Context, req *SetupRequest) (*RegularReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setup not implemented")
}

func RegisterDustyServer(s *grpc.Server, srv DustyServer) {
	s.RegisterService(&_Dusty_serviceDesc, srv)
}

func _Dusty_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DustyServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Dusty/Setup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DustyServer).Setup(ctx, req.(*SetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dusty_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Dusty",
	HandlerType: (*DustyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Setup",
			Handler:    _Dusty_Setup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
