// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package com_ta04_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password             int32    `protobuf:"varint,5,opt,name=password,proto3" json:"password,omitempty"`
	PrimeNumber          int32    `protobuf:"varint,6,opt,name=prime_number,json=primeNumber,proto3" json:"prime_number,omitempty"`
	GeneratorValue       int32    `protobuf:"varint,7,opt,name=generator_value,json=generatorValue,proto3" json:"generator_value,omitempty"`
	EmailAddress         string   `protobuf:"bytes,8,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,9,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DateOfBirth          string   `protobuf:"bytes,10,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Address              string   `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
	Role                 string   `protobuf:"bytes,12,opt,name=role,proto3" json:"role,omitempty"`
	Status               string   `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty"`
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

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
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

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() int32 {
	if m != nil {
		return m.Password
	}
	return 0
}

func (m *User) GetPrimeNumber() int32 {
	if m != nil {
		return m.PrimeNumber
	}
	return 0
}

func (m *User) GetGeneratorValue() int32 {
	if m != nil {
		return m.GeneratorValue
	}
	return 0
}

func (m *User) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *User) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetAllUsersRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllUsersRequest) Reset()         { *m = GetAllUsersRequest{} }
func (m *GetAllUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllUsersRequest) ProtoMessage()    {}
func (*GetAllUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetAllUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUsersRequest.Unmarshal(m, b)
}
func (m *GetAllUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUsersRequest.Marshal(b, m, deterministic)
}
func (m *GetAllUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUsersRequest.Merge(m, src)
}
func (m *GetAllUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllUsersRequest.Size(m)
}
func (m *GetAllUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUsersRequest proto.InternalMessageInfo

func (m *GetAllUsersRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *GetAllUsersRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *GetAllUsersRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetOneUserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	WithCredentials      bool     `protobuf:"varint,3,opt,name=with_credentials,json=withCredentials,proto3" json:"with_credentials,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOneUserRequest) Reset()         { *m = GetOneUserRequest{} }
func (m *GetOneUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetOneUserRequest) ProtoMessage()    {}
func (*GetOneUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *GetOneUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneUserRequest.Unmarshal(m, b)
}
func (m *GetOneUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneUserRequest.Marshal(b, m, deterministic)
}
func (m *GetOneUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneUserRequest.Merge(m, src)
}
func (m *GetOneUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetOneUserRequest.Size(m)
}
func (m *GetOneUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneUserRequest proto.InternalMessageInfo

func (m *GetOneUserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetOneUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetOneUserRequest) GetWithCredentials() bool {
	if m != nil {
		return m.WithCredentials
	}
	return false
}

type Response struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Error                *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "com.ta04.srv.user.User")
	proto.RegisterType((*GetAllUsersRequest)(nil), "com.ta04.srv.user.GetAllUsersRequest")
	proto.RegisterType((*GetOneUserRequest)(nil), "com.ta04.srv.user.GetOneUserRequest")
	proto.RegisterType((*Response)(nil), "com.ta04.srv.user.Response")
	proto.RegisterType((*Error)(nil), "com.ta04.srv.user.Error")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xed, 0x6e, 0xd3, 0x30,
	0x14, 0xa5, 0x69, 0xb3, 0x35, 0x37, 0xed, 0xc6, 0x2c, 0x04, 0xd6, 0x26, 0xa4, 0x12, 0x40, 0x14,
	0x21, 0x22, 0x54, 0xe0, 0x01, 0xc6, 0x84, 0xfa, 0x6f, 0x13, 0x99, 0xb6, 0xbf, 0x91, 0xdb, 0xdc,
	0xae, 0x41, 0x49, 0x9c, 0xd9, 0x4e, 0x27, 0xde, 0x80, 0x27, 0xe0, 0x95, 0x78, 0x2d, 0xe4, 0x9b,
	0xa6, 0xfb, 0x0a, 0x1f, 0x12, 0xff, 0x7c, 0xcf, 0x39, 0x3a, 0xc7, 0xb9, 0xc7, 0x0a, 0x40, 0xa5,
	0x51, 0x85, 0xa5, 0x92, 0x46, 0xb2, 0xbd, 0xb9, 0xcc, 0x43, 0x23, 0xde, 0x7d, 0x08, 0xb5, 0x5a,
	0x85, 0x96, 0x08, 0xbe, 0x77, 0xa1, 0x77, 0xa6, 0x51, 0xb1, 0x1d, 0x70, 0xd2, 0x84, 0x77, 0x46,
	0x9d, 0xb1, 0x1b, 0x39, 0x69, 0xc2, 0x9e, 0x02, 0x2c, 0x52, 0xa5, 0x4d, 0x5c, 0x88, 0x1c, 0xb9,
	0x33, 0xea, 0x8c, 0xbd, 0xc8, 0x23, 0xe4, 0x58, 0xe4, 0xc8, 0x0e, 0xc0, 0xcb, 0x44, 0xc3, 0x76,
	0x89, 0xed, 0x5b, 0x80, 0xc8, 0x7d, 0xe8, 0x5b, 0x73, 0xe2, 0x7a, 0x35, 0xd7, 0xcc, 0x96, 0x2b,
	0x85, 0xd6, 0x57, 0x52, 0x25, 0xdc, 0xa5, 0xb4, 0xcd, 0xcc, 0x9e, 0xc1, 0xa0, 0x54, 0x69, 0x8e,
	0x71, 0x51, 0xe5, 0x33, 0x54, 0x7c, 0x8b, 0x78, 0x9f, 0xb0, 0x63, 0x82, 0xd8, 0x2b, 0xd8, 0xbd,
	0xc0, 0x02, 0x95, 0x30, 0x52, 0xc5, 0x2b, 0x91, 0x55, 0xc8, 0xb7, 0x49, 0xb5, 0xb3, 0x81, 0xcf,
	0x2d, 0xca, 0x9e, 0xc3, 0x10, 0x73, 0x91, 0x66, 0xb1, 0x48, 0x12, 0x85, 0x5a, 0xf3, 0x3e, 0x5d,
	0x64, 0x40, 0xe0, 0x61, 0x8d, 0x51, 0xe0, 0x52, 0x16, 0x9b, 0x40, 0x8f, 0x34, 0x3e, 0x61, 0xeb,
	0xc0, 0x00, 0x86, 0x89, 0x30, 0x18, 0xcb, 0x45, 0x3c, 0x4b, 0x95, 0x59, 0x72, 0xa8, 0x35, 0x16,
	0x3c, 0x59, 0x7c, 0xb2, 0x10, 0xe3, 0xb0, 0xdd, 0xa4, 0xf8, 0xc4, 0x36, 0x23, 0x63, 0xd0, 0x53,
	0x32, 0x43, 0x3e, 0x20, 0x98, 0xce, 0xec, 0x31, 0x6c, 0x69, 0x23, 0x4c, 0xa5, 0xf9, 0x90, 0xd0,
	0xf5, 0x14, 0x9c, 0x03, 0x9b, 0xa2, 0x39, 0xcc, 0x32, 0xdb, 0x87, 0x8e, 0xf0, 0xb2, 0x42, 0x6d,
	0xd8, 0x23, 0x70, 0x2f, 0x2b, 0x54, 0xdf, 0xa8, 0x1a, 0x2f, 0xaa, 0x87, 0x8d, 0xaf, 0xd3, 0xea,
	0xdb, 0xbd, 0xe5, 0xfb, 0x15, 0xf6, 0xa6, 0x68, 0x4e, 0x0a, 0xb4, 0xbe, 0x8d, 0xed, 0xdd, 0xba,
	0x6f, 0x56, 0xe6, 0xdc, 0xa9, 0xec, 0x35, 0x3c, 0xbc, 0x4a, 0xcd, 0x32, 0x9e, 0x2b, 0x4c, 0xb0,
	0x30, 0xa9, 0xc8, 0xea, 0x88, 0x7e, 0xb4, 0x6b, 0xf1, 0xa3, 0x6b, 0x38, 0xf8, 0xd1, 0x81, 0x7e,
	0x84, 0xba, 0x94, 0x85, 0x46, 0xf6, 0x16, 0x5c, 0xeb, 0xa1, 0x79, 0x67, 0xd4, 0x1d, 0xfb, 0x93,
	0x27, 0xe1, 0xbd, 0xe7, 0x17, 0xd2, 0x95, 0x6a, 0x15, 0x7b, 0x03, 0x3d, 0x7b, 0xa0, 0xf8, 0x3f,
	0xa8, 0x49, 0xc4, 0x42, 0x70, 0x51, 0x29, 0xa9, 0xe8, 0x22, 0xfe, 0x84, 0xb7, 0xa8, 0x3f, 0x5b,
	0x3e, 0xaa, 0x65, 0xc1, 0x47, 0x70, 0x69, 0xb6, 0x9b, 0x9b, 0xcb, 0x04, 0xd7, 0x9f, 0x4e, 0x67,
	0xdb, 0x5f, 0x8e, 0x5a, 0x8b, 0x8b, 0xe6, 0xdb, 0x9b, 0x71, 0xf2, 0xd3, 0x01, 0xdf, 0xa6, 0x9e,
	0xa2, 0x5a, 0xa5, 0x73, 0x64, 0xa7, 0xe0, 0xdf, 0xe8, 0x88, 0xbd, 0x6c, 0x89, 0xbd, 0xdf, 0xe1,
	0xfe, 0x41, 0x8b, 0xac, 0xd9, 0x52, 0xf0, 0x80, 0x7d, 0x01, 0xb8, 0x2e, 0x88, 0xbd, 0x68, 0xf7,
	0xbc, 0xdd, 0xdf, 0xdf, 0x2c, 0xa7, 0x30, 0x3c, 0x52, 0x68, 0x9f, 0xe8, 0xda, 0xf5, 0x77, 0xeb,
	0xfc, 0x07, 0xa3, 0xb3, 0x32, 0xf9, 0x7f, 0xa3, 0xd9, 0x16, 0xfd, 0x82, 0xde, 0xff, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x64, 0xd8, 0xfa, 0xd5, 0x90, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...client.CallOption) (*Response, error)
	GetOneUser(ctx context.Context, in *GetOneUserRequest, opts ...client.CallOption) (*Response, error)
	CreateOneUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	UpdateOneUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.ta04.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAllUsers", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetOneUser(ctx context.Context, in *GetOneUserRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetOneUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateOneUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.CreateOneUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateOneUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.UpdateOneUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	GetAllUsers(context.Context, *GetAllUsersRequest, *Response) error
	GetOneUser(context.Context, *GetOneUserRequest, *Response) error
	CreateOneUser(context.Context, *User, *Response) error
	UpdateOneUser(context.Context, *User, *Response) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, out *Response) error {
	return h.UserServiceHandler.GetAllUsers(ctx, in, out)
}

func (h *UserService) GetOneUser(ctx context.Context, in *GetOneUserRequest, out *Response) error {
	return h.UserServiceHandler.GetOneUser(ctx, in, out)
}

func (h *UserService) CreateOneUser(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.CreateOneUser(ctx, in, out)
}

func (h *UserService) UpdateOneUser(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.UpdateOneUser(ctx, in, out)
}
