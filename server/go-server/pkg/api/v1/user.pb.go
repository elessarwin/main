// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type User struct {
	// message body
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	DisplayName          string   `protobuf:"bytes,5,opt,name=displayName,proto3" json:"displayName,omitempty"`
	AvatarURL            string   `protobuf:"bytes,6,opt,name=avatarURL,proto3" json:"avatarURL,omitempty"`
	ChatgroupIds         []string `protobuf:"bytes,7,rep,name=chatgroupIds,proto3" json:"chatgroupIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetAvatarURL() string {
	if m != nil {
		return m.AvatarURL
	}
	return ""
}

func (m *User) GetChatgroupIds() []string {
	if m != nil {
		return m.ChatgroupIds
	}
	return nil
}

type UserSignUp struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSignUp) Reset()         { *m = UserSignUp{} }
func (m *UserSignUp) String() string { return proto.CompactTextString(m) }
func (*UserSignUp) ProtoMessage()    {}
func (*UserSignUp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserSignUp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSignUp.Unmarshal(m, b)
}
func (m *UserSignUp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSignUp.Marshal(b, m, deterministic)
}
func (m *UserSignUp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSignUp.Merge(m, src)
}
func (m *UserSignUp) XXX_Size() int {
	return xxx_messageInfo_UserSignUp.Size(m)
}
func (m *UserSignUp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSignUp.DiscardUnknown(m)
}

var xxx_messageInfo_UserSignUp proto.InternalMessageInfo

func (m *UserSignUp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserSignUp) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type UserSignIn struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSignIn) Reset()         { *m = UserSignIn{} }
func (m *UserSignIn) String() string { return proto.CompactTextString(m) }
func (*UserSignIn) ProtoMessage()    {}
func (*UserSignIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserSignIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSignIn.Unmarshal(m, b)
}
func (m *UserSignIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSignIn.Marshal(b, m, deterministic)
}
func (m *UserSignIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSignIn.Merge(m, src)
}
func (m *UserSignIn) XXX_Size() int {
	return xxx_messageInfo_UserSignIn.Size(m)
}
func (m *UserSignIn) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSignIn.DiscardUnknown(m)
}

var xxx_messageInfo_UserSignIn proto.InternalMessageInfo

func (m *UserSignIn) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "v1.User")
	proto.RegisterType((*UserSignUp)(nil), "v1.UserSignUp")
	proto.RegisterType((*UserSignIn)(nil), "v1.UserSignIn")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x9b, 0xf4, 0x87, 0x66, 0x8a, 0xba, 0xb0, 0x10, 0x8a, 0x02, 0x8b, 0xc8, 0x62, 0xd1,
	0x95, 0xab, 0x96, 0x2b, 0xc0, 0x22, 0x12, 0x62, 0x11, 0x94, 0x03, 0xb8, 0x8d, 0x6b, 0x2c, 0x25,
	0xb1, 0x65, 0x3b, 0x91, 0x7a, 0x45, 0x4e, 0x85, 0x62, 0xb7, 0x0d, 0x15, 0x88, 0xe5, 0x7c, 0xcf,
	0x6f, 0xe6, 0xf9, 0x01, 0xb4, 0x86, 0x69, 0xa2, 0xb4, 0xb4, 0x12, 0x85, 0xdd, 0x26, 0x79, 0xe0,
	0x52, 0xf2, 0x8a, 0xad, 0x1d, 0xd9, 0xb5, 0x87, 0x35, 0xab, 0x95, 0x3d, 0xfa, 0x07, 0xf8, 0x2b,
	0x80, 0x49, 0x61, 0x98, 0x46, 0x4b, 0x08, 0x45, 0x19, 0x07, 0x69, 0xb0, 0x8a, 0xf2, 0x50, 0x94,
	0xe8, 0x11, 0xa2, 0x83, 0xd0, 0xc6, 0xbe, 0xd3, 0x9a, 0xc5, 0xa1, 0xc3, 0x03, 0x40, 0x09, 0xcc,
	0x2b, 0x7a, 0x12, 0xc7, 0x4e, 0xbc, 0xcc, 0xe8, 0x0e, 0xa6, 0xac, 0xa6, 0xa2, 0x8a, 0x27, 0x4e,
	0xf0, 0x03, 0x4a, 0x61, 0x51, 0x0a, 0xa3, 0x2a, 0x7a, 0x74, 0xa6, 0xa9, 0xd3, 0x7e, 0xa2, 0xfe,
	0x22, 0xed, 0xa8, 0xa5, 0xba, 0xc8, 0xdf, 0xe2, 0x99, 0xbf, 0x78, 0x01, 0x08, 0xc3, 0xed, 0xfe,
	0x93, 0x5a, 0xae, 0x65, 0xab, 0xb2, 0xd2, 0xc4, 0x37, 0xe9, 0x78, 0x15, 0xe5, 0x57, 0x0c, 0xbf,
	0x00, 0xf4, 0x7f, 0xf9, 0x10, 0xbc, 0x29, 0xd4, 0x90, 0x23, 0xf8, 0x27, 0x47, 0xf8, 0x2b, 0x07,
	0xc6, 0xc3, 0x96, 0xac, 0xf9, 0x7b, 0xcb, 0x96, 0xc3, 0xc2, 0xbd, 0x61, 0xba, 0x13, 0x7b, 0x86,
	0xb6, 0x30, 0x3b, 0x1d, 0x5d, 0x92, 0x6e, 0x43, 0x86, 0x10, 0xc9, 0x3d, 0xf1, 0xed, 0x93, 0x73,
	0xfb, 0xe4, 0xb5, 0x6f, 0x1f, 0x8f, 0xd0, 0x93, 0xf7, 0x64, 0xcd, 0xb5, 0x27, 0x6b, 0x92, 0xf9,
	0x79, 0xc6, 0xa3, 0xdd, 0xcc, 0xf9, 0x9e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x4e, 0x92,
	0x91, 0xd5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// SignUp register a new user
	SignUp(ctx context.Context, in *UserSignUp, opts ...grpc.CallOption) (*empty.Empty, error)
	// SignIn signin user
	SignIn(ctx context.Context, in *UserSignIn, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) SignUp(ctx context.Context, in *UserSignUp, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.UserService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignIn(ctx context.Context, in *UserSignIn, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/v1.UserService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// SignUp register a new user
	SignUp(context.Context, *UserSignUp) (*empty.Empty, error)
	// SignIn signin user
	SignIn(context.Context, *UserSignIn) (*User, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) SignUp(ctx context.Context, req *UserSignUp) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedUserServiceServer) SignIn(ctx context.Context, req *UserSignIn) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignUp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignUp(ctx, req.(*UserSignUp))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignIn(ctx, req.(*UserSignIn))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _UserService_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _UserService_SignIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
