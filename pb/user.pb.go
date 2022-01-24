// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/user.proto

// user 包

package user

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

// 枚举类型
type EnumUserSex int32

const (
	EnumUserSex_SEX_INIT   EnumUserSex = 0
	EnumUserSex_SEX_MALE   EnumUserSex = 1
	EnumUserSex_SEX_FEMALE EnumUserSex = 2
)

var EnumUserSex_name = map[int32]string{
	0: "SEX_INIT",
	1: "SEX_MALE",
	2: "SEX_FEMALE",
}

var EnumUserSex_value = map[string]int32{
	"SEX_INIT":   0,
	"SEX_MALE":   1,
	"SEX_FEMALE": 2,
}

func (x EnumUserSex) String() string {
	return proto.EnumName(EnumUserSex_name, int32(x))
}

func (EnumUserSex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{0}
}

// 用户实体模型
type UserEntity struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserEntity) Reset()         { *m = UserEntity{} }
func (m *UserEntity) String() string { return proto.CompactTextString(m) }
func (*UserEntity) ProtoMessage()    {}
func (*UserEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{0}
}

func (m *UserEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEntity.Unmarshal(m, b)
}
func (m *UserEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEntity.Marshal(b, m, deterministic)
}
func (m *UserEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEntity.Merge(m, src)
}
func (m *UserEntity) XXX_Size() int {
	return xxx_messageInfo_UserEntity.Size(m)
}
func (m *UserEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEntity.DiscardUnknown(m)
}

var xxx_messageInfo_UserEntity proto.InternalMessageInfo

func (m *UserEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserEntity) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

// User 服务的各个接口的请求/响应结构
type UserIndexRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIndexRequest) Reset()         { *m = UserIndexRequest{} }
func (m *UserIndexRequest) String() string { return proto.CompactTextString(m) }
func (*UserIndexRequest) ProtoMessage()    {}
func (*UserIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{1}
}

func (m *UserIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIndexRequest.Unmarshal(m, b)
}
func (m *UserIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIndexRequest.Marshal(b, m, deterministic)
}
func (m *UserIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIndexRequest.Merge(m, src)
}
func (m *UserIndexRequest) XXX_Size() int {
	return xxx_messageInfo_UserIndexRequest.Size(m)
}
func (m *UserIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIndexRequest proto.InternalMessageInfo

func (m *UserIndexRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *UserIndexRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type UserIndexResponse struct {
	Err int32  `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 返回一个 UserEntity 对象的列表数据
	Data                 []*UserEntity `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserIndexResponse) Reset()         { *m = UserIndexResponse{} }
func (m *UserIndexResponse) String() string { return proto.CompactTextString(m) }
func (*UserIndexResponse) ProtoMessage()    {}
func (*UserIndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{2}
}

func (m *UserIndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIndexResponse.Unmarshal(m, b)
}
func (m *UserIndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIndexResponse.Marshal(b, m, deterministic)
}
func (m *UserIndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIndexResponse.Merge(m, src)
}
func (m *UserIndexResponse) XXX_Size() int {
	return xxx_messageInfo_UserIndexResponse.Size(m)
}
func (m *UserIndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserIndexResponse proto.InternalMessageInfo

func (m *UserIndexResponse) GetErr() int32 {
	if m != nil {
		return m.Err
	}
	return 0
}

func (m *UserIndexResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserIndexResponse) GetData() []*UserEntity {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserViewRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserViewRequest) Reset()         { *m = UserViewRequest{} }
func (m *UserViewRequest) String() string { return proto.CompactTextString(m) }
func (*UserViewRequest) ProtoMessage()    {}
func (*UserViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{3}
}

func (m *UserViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserViewRequest.Unmarshal(m, b)
}
func (m *UserViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserViewRequest.Marshal(b, m, deterministic)
}
func (m *UserViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserViewRequest.Merge(m, src)
}
func (m *UserViewRequest) XXX_Size() int {
	return xxx_messageInfo_UserViewRequest.Size(m)
}
func (m *UserViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserViewRequest proto.InternalMessageInfo

func (m *UserViewRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type UserViewResponse struct {
	Err int32  `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 返回一个 UserEntity 对象
	Data                 *UserEntity `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserViewResponse) Reset()         { *m = UserViewResponse{} }
func (m *UserViewResponse) String() string { return proto.CompactTextString(m) }
func (*UserViewResponse) ProtoMessage()    {}
func (*UserViewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{4}
}

func (m *UserViewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserViewResponse.Unmarshal(m, b)
}
func (m *UserViewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserViewResponse.Marshal(b, m, deterministic)
}
func (m *UserViewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserViewResponse.Merge(m, src)
}
func (m *UserViewResponse) XXX_Size() int {
	return xxx_messageInfo_UserViewResponse.Size(m)
}
func (m *UserViewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserViewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserViewResponse proto.InternalMessageInfo

func (m *UserViewResponse) GetErr() int32 {
	if m != nil {
		return m.Err
	}
	return 0
}

func (m *UserViewResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserViewResponse) GetData() *UserEntity {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserPostRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPostRequest) Reset()         { *m = UserPostRequest{} }
func (m *UserPostRequest) String() string { return proto.CompactTextString(m) }
func (*UserPostRequest) ProtoMessage()    {}
func (*UserPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{5}
}

func (m *UserPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPostRequest.Unmarshal(m, b)
}
func (m *UserPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPostRequest.Marshal(b, m, deterministic)
}
func (m *UserPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPostRequest.Merge(m, src)
}
func (m *UserPostRequest) XXX_Size() int {
	return xxx_messageInfo_UserPostRequest.Size(m)
}
func (m *UserPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserPostRequest proto.InternalMessageInfo

func (m *UserPostRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserPostRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserPostRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type UserPostResponse struct {
	Err                  int32    `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPostResponse) Reset()         { *m = UserPostResponse{} }
func (m *UserPostResponse) String() string { return proto.CompactTextString(m) }
func (*UserPostResponse) ProtoMessage()    {}
func (*UserPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{6}
}

func (m *UserPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPostResponse.Unmarshal(m, b)
}
func (m *UserPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPostResponse.Marshal(b, m, deterministic)
}
func (m *UserPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPostResponse.Merge(m, src)
}
func (m *UserPostResponse) XXX_Size() int {
	return xxx_messageInfo_UserPostResponse.Size(m)
}
func (m *UserPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserPostResponse proto.InternalMessageInfo

func (m *UserPostResponse) GetErr() int32 {
	if m != nil {
		return m.Err
	}
	return 0
}

func (m *UserPostResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UserDeleteRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteRequest) Reset()         { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()    {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{7}
}

func (m *UserDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteRequest.Unmarshal(m, b)
}
func (m *UserDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteRequest.Marshal(b, m, deterministic)
}
func (m *UserDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteRequest.Merge(m, src)
}
func (m *UserDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_UserDeleteRequest.Size(m)
}
func (m *UserDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteRequest proto.InternalMessageInfo

func (m *UserDeleteRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type UserDeleteResponse struct {
	Err                  int32    `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteResponse) Reset()         { *m = UserDeleteResponse{} }
func (m *UserDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*UserDeleteResponse) ProtoMessage()    {}
func (*UserDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{8}
}

func (m *UserDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteResponse.Unmarshal(m, b)
}
func (m *UserDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteResponse.Marshal(b, m, deterministic)
}
func (m *UserDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteResponse.Merge(m, src)
}
func (m *UserDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_UserDeleteResponse.Size(m)
}
func (m *UserDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteResponse proto.InternalMessageInfo

func (m *UserDeleteResponse) GetErr() int32 {
	if m != nil {
		return m.Err
	}
	return 0
}

func (m *UserDeleteResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("user.EnumUserSex", EnumUserSex_name, EnumUserSex_value)
	proto.RegisterType((*UserEntity)(nil), "user.UserEntity")
	proto.RegisterType((*UserIndexRequest)(nil), "user.UserIndexRequest")
	proto.RegisterType((*UserIndexResponse)(nil), "user.UserIndexResponse")
	proto.RegisterType((*UserViewRequest)(nil), "user.UserViewRequest")
	proto.RegisterType((*UserViewResponse)(nil), "user.UserViewResponse")
	proto.RegisterType((*UserPostRequest)(nil), "user.UserPostRequest")
	proto.RegisterType((*UserPostResponse)(nil), "user.UserPostResponse")
	proto.RegisterType((*UserDeleteRequest)(nil), "user.UserDeleteRequest")
	proto.RegisterType((*UserDeleteResponse)(nil), "user.UserDeleteResponse")
}

func init() { proto.RegisterFile("pb/user.proto", fileDescriptor_c4f917211a932270) }

var fileDescriptor_c4f917211a932270 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdf, 0x4b, 0xe3, 0x40,
	0x10, 0x6e, 0x9a, 0xdc, 0x91, 0x4c, 0xef, 0x47, 0x6e, 0xe1, 0x7a, 0x21, 0xf7, 0x52, 0xf6, 0xee,
	0xa0, 0xdc, 0x43, 0x85, 0x0a, 0xa2, 0x08, 0x42, 0xd1, 0x08, 0x05, 0x15, 0x49, 0x54, 0x7c, 0x50,
	0x4a, 0x4a, 0x96, 0x12, 0xb0, 0x49, 0xcc, 0x26, 0xb4, 0xf6, 0xd5, 0x7f, 0x5c, 0x66, 0xb7, 0x69,
	0xd6, 0x16, 0xa5, 0xf8, 0x94, 0x6f, 0x26, 0xf3, 0xcd, 0x7e, 0xfb, 0xcd, 0x2c, 0x7c, 0xcd, 0xc6,
	0x3b, 0x25, 0x67, 0x79, 0x2f, 0xcb, 0xd3, 0x22, 0x25, 0x06, 0x62, 0xda, 0x07, 0xb8, 0xe6, 0x2c,
	0xf7, 0x92, 0x22, 0x2e, 0x9e, 0x08, 0x01, 0x23, 0x09, 0xa7, 0xcc, 0xd1, 0x3a, 0x5a, 0xd7, 0xf2,
	0x05, 0x26, 0x36, 0xe8, 0xe1, 0x84, 0x39, 0xcd, 0x8e, 0xd6, 0xfd, 0xe4, 0x23, 0xa4, 0xc7, 0x60,
	0x23, 0x67, 0x98, 0x44, 0x6c, 0xee, 0xb3, 0xc7, 0x92, 0xf1, 0x02, 0x99, 0x19, 0x96, 0x69, 0xa2,
	0x4c, 0x60, 0xf2, 0x1b, 0x2c, 0xfc, 0x8e, 0x78, 0xbc, 0xa8, 0xf8, 0x26, 0x26, 0x82, 0x78, 0xc1,
	0xe8, 0x3d, 0xfc, 0x50, 0x9a, 0xf0, 0x2c, 0x4d, 0xb8, 0x38, 0x8b, 0xe5, 0xf9, 0xb2, 0x09, 0x42,
	0xcc, 0x4c, 0xf9, 0x44, 0xb0, 0x2d, 0x1f, 0x21, 0xf9, 0x0b, 0x46, 0x14, 0x16, 0xa1, 0xa3, 0x77,
	0xf4, 0x6e, 0xab, 0x6f, 0xf7, 0xc4, 0x95, 0xea, 0x3b, 0xf8, 0xe2, 0x2f, 0xfd, 0x03, 0xdf, 0x31,
	0x77, 0x13, 0xb3, 0x59, 0x25, 0xd1, 0x06, 0xbd, 0x8c, 0xa3, 0xaa, 0x79, 0x19, 0x47, 0xf4, 0x4e,
	0x5e, 0x44, 0x16, 0x7d, 0x48, 0x82, 0xf6, 0x8e, 0x84, 0x40, 0x4a, 0xb8, 0x4c, 0x79, 0xa1, 0xb8,
	0xb4, 0xe1, 0xaf, 0x0b, 0x66, 0x16, 0x72, 0x3e, 0x4b, 0xf3, 0x68, 0x79, 0xc6, 0x2a, 0xae, 0xbc,
	0xd7, 0x6b, 0xef, 0xf7, 0xa4, 0x64, 0xd9, 0x74, 0x7b, 0xc9, 0xf4, 0x9f, 0xb4, 0xfb, 0x84, 0x3d,
	0xb0, 0x82, 0xbd, 0xed, 0xc8, 0x3e, 0x10, 0xb5, 0x6c, 0xfb, 0x03, 0xfe, 0x1f, 0x40, 0xcb, 0x4b,
	0xca, 0x29, 0xb2, 0x03, 0x36, 0x27, 0x5f, 0xc0, 0x0c, 0xbc, 0xdb, 0xd1, 0xf0, 0x62, 0x78, 0x65,
	0x37, 0xaa, 0xe8, 0x7c, 0x70, 0xe6, 0xd9, 0x1a, 0xf9, 0x06, 0x80, 0xd1, 0xa9, 0x27, 0xe2, 0x66,
	0xff, 0xb9, 0x09, 0x06, 0xf2, 0xc8, 0x11, 0x58, 0xab, 0x9d, 0x20, 0xed, 0xda, 0x56, 0x75, 0xd3,
	0xdc, 0x5f, 0x1b, 0x79, 0xa9, 0x92, 0x36, 0xc8, 0x21, 0x98, 0xd5, 0x3c, 0xc9, 0xcf, 0xba, 0x4c,
	0x59, 0x02, 0xb7, 0xbd, 0x9e, 0x5e, 0x27, 0xa3, 0xb3, 0x2a, 0x59, 0x19, 0x9f, 0x4a, 0x56, 0x07,
	0x40, 0x1b, 0x64, 0x20, 0x9f, 0x91, 0xf4, 0x8d, 0x28, 0x12, 0x5f, 0x19, 0xee, 0x3a, 0x9b, 0x3f,
	0xaa, 0x16, 0xe3, 0xcf, 0xe2, 0x59, 0xee, 0xbe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x96, 0x51, 0x43,
	0x18, 0xa7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	UserIndex(ctx context.Context, in *UserIndexRequest, opts ...grpc.CallOption) (*UserIndexResponse, error)
	UserView(ctx context.Context, in *UserViewRequest, opts ...grpc.CallOption) (*UserViewResponse, error)
	UserPost(ctx context.Context, in *UserPostRequest, opts ...grpc.CallOption) (*UserPostResponse, error)
	UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserIndex(ctx context.Context, in *UserIndexRequest, opts ...grpc.CallOption) (*UserIndexResponse, error) {
	out := new(UserIndexResponse)
	err := c.cc.Invoke(ctx, "/user.User/UserIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserView(ctx context.Context, in *UserViewRequest, opts ...grpc.CallOption) (*UserViewResponse, error) {
	out := new(UserViewResponse)
	err := c.cc.Invoke(ctx, "/user.User/UserView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserPost(ctx context.Context, in *UserPostRequest, opts ...grpc.CallOption) (*UserPostResponse, error) {
	out := new(UserPostResponse)
	err := c.cc.Invoke(ctx, "/user.User/UserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	out := new(UserDeleteResponse)
	err := c.cc.Invoke(ctx, "/user.User/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	UserIndex(context.Context, *UserIndexRequest) (*UserIndexResponse, error)
	UserView(context.Context, *UserViewRequest) (*UserViewResponse, error)
	UserPost(context.Context, *UserPostRequest) (*UserPostResponse, error)
	UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) UserIndex(ctx context.Context, req *UserIndexRequest) (*UserIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserIndex not implemented")
}
func (*UnimplementedUserServer) UserView(ctx context.Context, req *UserViewRequest) (*UserViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserView not implemented")
}
func (*UnimplementedUserServer) UserPost(ctx context.Context, req *UserPostRequest) (*UserPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserPost not implemented")
}
func (*UnimplementedUserServer) UserDelete(ctx context.Context, req *UserDeleteRequest) (*UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_UserIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserIndex(ctx, req.(*UserIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserView(ctx, req.(*UserViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserPost(ctx, req.(*UserPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserDelete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserIndex",
			Handler:    _User_UserIndex_Handler,
		},
		{
			MethodName: "UserView",
			Handler:    _User_UserView_Handler,
		},
		{
			MethodName: "UserPost",
			Handler:    _User_UserPost_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _User_UserDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/user.proto",
}