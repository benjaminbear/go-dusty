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
	Update               bool     `protobuf:"varint,4,opt,name=update,proto3" json:"update,omitempty"`
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
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x31, 0x25, 0x55, 0x72, 0xea, 0x64, 0xa1, 0xca, 0xaa, 0x18, 0xac, 0x4c, 0x9e, 0x32,
	0xc0, 0xc4, 0x86, 0x10, 0x2b, 0xcb, 0xf1, 0x0b, 0x4c, 0x38, 0x75, 0x89, 0xb1, 0xb9, 0xb3, 0x91,
	0x22, 0xf1, 0xe3, 0x51, 0x4d, 0x49, 0x33, 0x9d, 0xbe, 0x77, 0x77, 0x4f, 0xef, 0x41, 0xc7, 0x69,
	0x1c, 0x12, 0xc7, 0x1c, 0x75, 0x5b, 0xc7, 0x18, 0xa7, 0xfe, 0x09, 0x76, 0x48, 0xc7, 0x32, 0x79,
	0x46, 0x4a, 0xd3, 0xac, 0x0f, 0xd0, 0x06, 0x12, 0xf1, 0x47, 0x12, 0xa3, 0xec, 0xc6, 0x75, 0xb8,
	0xb0, 0xbe, 0x85, 0x86, 0x98, 0x23, 0x9b, 0x6b, 0xab, 0x5c, 0x87, 0x7f, 0xd0, 0xff, 0xc0, 0xee,
	0x8d, 0x72, 0x49, 0x48, 0x5f, 0x85, 0x24, 0x9f, 0x1c, 0x8a, 0x10, 0x7f, 0xfa, 0x40, 0x46, 0xd5,
	0xc3, 0x85, 0xf5, 0x1d, 0x74, 0x92, 0x68, 0x14, 0xa4, 0x14, 0xcf, 0x2e, 0x17, 0xe1, 0xf4, 0xf9,
	0x1d, 0x5e, 0x29, 0x44, 0x9e, 0xcd, 0xc6, 0x2a, 0xd7, 0xe0, 0xc2, 0x7a, 0x0f, 0xdb, 0x92, 0x3e,
	0x7c, 0x26, 0x73, 0x63, 0x95, 0x6b, 0xf1, 0x4c, 0xf7, 0xcf, 0xd0, 0xbc, 0x14, 0xc9, 0xb3, 0x7e,
	0x84, 0xa6, 0xc6, 0xd0, 0xfb, 0xe1, 0xbf, 0xdc, 0xb0, 0xce, 0x75, 0x58, 0xe9, 0xeb, 0xc6, 0xfd,
	0xd5, 0xfb, 0xb6, 0x2e, 0x1e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x56, 0xb9, 0x23, 0xba, 0x21,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DustyClient is the client API for Dusty service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DustyClient interface {
	Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*RegularReply, error)
}

type dustyClient struct {
	cc grpc.ClientConnInterface
}

func NewDustyClient(cc grpc.ClientConnInterface) DustyClient {
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
