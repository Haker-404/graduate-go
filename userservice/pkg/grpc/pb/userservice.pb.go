// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userservice.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type LoginReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Resp                 *Resp    `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{1}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *LoginReply) GetResp() *Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

type GetUserListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListRequest) Reset()         { *m = GetUserListRequest{} }
func (m *GetUserListRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserListRequest) ProtoMessage()    {}
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{2}
}

func (m *GetUserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListRequest.Unmarshal(m, b)
}
func (m *GetUserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListRequest.Marshal(b, m, deterministic)
}
func (m *GetUserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListRequest.Merge(m, src)
}
func (m *GetUserListRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserListRequest.Size(m)
}
func (m *GetUserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListRequest proto.InternalMessageInfo

type GetUserListReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	User                 []*User  `protobuf:"bytes,2,rep,name=user,proto3" json:"user,omitempty"`
	Resp                 *Resp    `protobuf:"bytes,3,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListReply) Reset()         { *m = GetUserListReply{} }
func (m *GetUserListReply) String() string { return proto.CompactTextString(m) }
func (*GetUserListReply) ProtoMessage()    {}
func (*GetUserListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{3}
}

func (m *GetUserListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListReply.Unmarshal(m, b)
}
func (m *GetUserListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListReply.Marshal(b, m, deterministic)
}
func (m *GetUserListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListReply.Merge(m, src)
}
func (m *GetUserListReply) XXX_Size() int {
	return xxx_messageInfo_GetUserListReply.Size(m)
}
func (m *GetUserListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListReply proto.InternalMessageInfo

func (m *GetUserListReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetUserListReply) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUserListReply) GetResp() *Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

type GetUserInfoRequest struct {
	Seq                  string   `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoRequest) Reset()         { *m = GetUserInfoRequest{} }
func (m *GetUserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoRequest) ProtoMessage()    {}
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{4}
}

func (m *GetUserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoRequest.Unmarshal(m, b)
}
func (m *GetUserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoRequest.Merge(m, src)
}
func (m *GetUserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoRequest.Size(m)
}
func (m *GetUserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoRequest proto.InternalMessageInfo

func (m *GetUserInfoRequest) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

type GetUserInfoReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Resp                 *Resp    `protobuf:"bytes,3,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoReply) Reset()         { *m = GetUserInfoReply{} }
func (m *GetUserInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoReply) ProtoMessage()    {}
func (*GetUserInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{5}
}

func (m *GetUserInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoReply.Unmarshal(m, b)
}
func (m *GetUserInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoReply.Marshal(b, m, deterministic)
}
func (m *GetUserInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoReply.Merge(m, src)
}
func (m *GetUserInfoReply) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoReply.Size(m)
}
func (m *GetUserInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoReply proto.InternalMessageInfo

func (m *GetUserInfoReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetUserInfoReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUserInfoReply) GetResp() *Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

type Resp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Succ                 string   `protobuf:"bytes,2,opt,name=succ,proto3" json:"succ,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{6}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Resp) GetSucc() string {
	if m != nil {
		return m.Succ
	}
	return ""
}

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Seq                  string   `protobuf:"bytes,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Photo                []byte   `protobuf:"bytes,3,opt,name=photo,proto3" json:"photo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{7}
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

func (m *User) GetPhoto() []byte {
	if m != nil {
		return m.Photo
	}
	return nil
}

type SignRequest struct {
	Seq                  string   `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{8}
}

func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRequest.Unmarshal(m, b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
}
func (m *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(m, src)
}
func (m *SignRequest) XXX_Size() int {
	return xxx_messageInfo_SignRequest.Size(m)
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

func (m *SignRequest) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

type SignReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Resp                 *Resp    `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignReply) Reset()         { *m = SignReply{} }
func (m *SignReply) String() string { return proto.CompactTextString(m) }
func (*SignReply) ProtoMessage()    {}
func (*SignReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{9}
}

func (m *SignReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignReply.Unmarshal(m, b)
}
func (m *SignReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignReply.Marshal(b, m, deterministic)
}
func (m *SignReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignReply.Merge(m, src)
}
func (m *SignReply) XXX_Size() int {
	return xxx_messageInfo_SignReply.Size(m)
}
func (m *SignReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SignReply.DiscardUnknown(m)
}

var xxx_messageInfo_SignReply proto.InternalMessageInfo

func (m *SignReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SignReply) GetResp() *Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "pb.LoginReply")
	proto.RegisterType((*GetUserListRequest)(nil), "pb.GetUserListRequest")
	proto.RegisterType((*GetUserListReply)(nil), "pb.GetUserListReply")
	proto.RegisterType((*GetUserInfoRequest)(nil), "pb.GetUserInfoRequest")
	proto.RegisterType((*GetUserInfoReply)(nil), "pb.GetUserInfoReply")
	proto.RegisterType((*Resp)(nil), "pb.Resp")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*SignRequest)(nil), "pb.SignRequest")
	proto.RegisterType((*SignReply)(nil), "pb.SignReply")
}

func init() { proto.RegisterFile("userservice.proto", fileDescriptor_68a7ca558839fd2b) }

var fileDescriptor_68a7ca558839fd2b = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x8a, 0x13, 0x41,
	0x10, 0x26, 0x93, 0xc9, 0xb2, 0x5b, 0x13, 0x75, 0x2c, 0x42, 0x08, 0xc3, 0x82, 0x4b, 0x3f, 0x88,
	0x88, 0x64, 0x60, 0x7d, 0x5b, 0x7c, 0xd9, 0x7d, 0x11, 0x61, 0x41, 0x98, 0x65, 0x0f, 0x30, 0x89,
	0x9d, 0xb1, 0x21, 0xe9, 0xee, 0x4c, 0x75, 0x94, 0x45, 0x7c, 0xf1, 0x0a, 0x5e, 0xc5, 0x9b, 0x78,
	0x05, 0x0f, 0x22, 0xd5, 0xf3, 0x63, 0x87, 0x28, 0xea, 0xbe, 0x75, 0x7d, 0x53, 0xf5, 0xd5, 0x57,
	0x55, 0xdf, 0xc0, 0xe3, 0x1d, 0xc9, 0x9a, 0x64, 0xfd, 0x41, 0x2d, 0xe5, 0xdc, 0xd6, 0xc6, 0x19,
	0x8c, 0xec, 0x22, 0x3b, 0xad, 0x8c, 0xa9, 0xd6, 0x32, 0x2f, 0xad, 0xca, 0x4b, 0xad, 0x8d, 0x2b,
	0x9d, 0x32, 0x9a, 0x9a, 0x0c, 0xf1, 0x0a, 0xc6, 0xd7, 0xa6, 0x52, 0xba, 0x90, 0xdb, 0x9d, 0x24,
	0x87, 0x19, 0x1c, 0x33, 0x8d, 0x2e, 0x37, 0x72, 0x36, 0x38, 0x1b, 0x3c, 0x3b, 0x29, 0xfa, 0x18,
	0x53, 0x18, 0xda, 0x8f, 0xef, 0x66, 0x91, 0x87, 0xf9, 0x29, 0xae, 0x00, 0xda, 0x6a, 0xbb, 0xbe,
	0xc3, 0x29, 0x1c, 0x91, 0x2b, 0xdd, 0x8e, 0x7c, 0xe5, 0xa8, 0x68, 0x23, 0x3c, 0x85, 0xb8, 0x96,
	0x64, 0x7d, 0x61, 0x72, 0x7e, 0x3c, 0xb7, 0x8b, 0x79, 0x21, 0xc9, 0x16, 0x1e, 0x15, 0x13, 0xc0,
	0xd7, 0xd2, 0xdd, 0x92, 0xac, 0xaf, 0x15, 0xb9, 0x56, 0x87, 0x58, 0x41, 0xba, 0x87, 0xfe, 0x85,
	0x9f, 0x35, 0xce, 0xa2, 0xb3, 0x61, 0xc7, 0xcf, 0x85, 0x85, 0x47, 0xfb, 0xee, 0xc3, 0xdf, 0x76,
	0x7f, 0xda, 0x77, 0x7f, 0xa3, 0x57, 0xa6, 0xdb, 0x42, 0x0a, 0x43, 0x92, 0xdb, 0x76, 0x01, 0xfc,
	0x0c, 0xf4, 0x34, 0x79, 0xff, 0xa6, 0x67, 0xf0, 0xdf, 0x7a, 0x5e, 0x40, 0xcc, 0x11, 0x2b, 0xd8,
	0x50, 0xd5, 0x29, 0xd8, 0x50, 0x85, 0x08, 0x31, 0xed, 0x96, 0xcb, 0x76, 0xfd, 0xfe, 0x2d, 0xae,
	0x20, 0x66, 0x66, 0xfe, 0x16, 0x5c, 0x2c, 0xee, 0xae, 0xc5, 0x33, 0x44, 0xfd, 0x0c, 0x38, 0x81,
	0x91, 0x7d, 0x6f, 0x9c, 0xf1, 0xad, 0xc7, 0x45, 0x13, 0x88, 0x27, 0x90, 0xdc, 0xa8, 0x4a, 0xff,
	0x79, 0xf4, 0x4b, 0x38, 0x69, 0x12, 0xee, 0x7d, 0xe3, 0xf3, 0x6f, 0x11, 0x24, 0x2c, 0xf4, 0xa6,
	0x71, 0x27, 0x5e, 0xc2, 0xc8, 0xfb, 0x06, 0x53, 0x4e, 0x0c, 0x0d, 0x98, 0x3d, 0x0c, 0x10, 0xbb,
	0xbe, 0x13, 0xd3, 0x2f, 0xdf, 0x7f, 0x7c, 0x8d, 0x52, 0x91, 0xe4, 0xbc, 0xbd, 0x7c, 0xcd, 0x5f,
	0x2e, 0x06, 0xcf, 0xf1, 0x2d, 0x24, 0x81, 0x41, 0x70, 0xca, 0x65, 0x87, 0x3e, 0xca, 0x26, 0x07,
	0x38, 0x93, 0xa2, 0x27, 0x1d, 0x23, 0xb4, 0xa4, 0xcc, 0x70, 0xdb, 0x13, 0xf2, 0x85, 0xf7, 0x08,
	0x03, 0x6b, 0xec, 0x11, 0xf6, 0x56, 0x10, 0x33, 0x4f, 0x88, 0x98, 0x36, 0x84, 0x4a, 0xaf, 0x4c,
	0xfe, 0x89, 0xe4, 0xf6, 0x33, 0x5e, 0x40, 0xcc, 0xdb, 0xc3, 0x47, 0x5c, 0x17, 0x2c, 0x3a, 0x7b,
	0xf0, 0x0b, 0x08, 0x24, 0x89, 0x56, 0x12, 0xa9, 0x4a, 0x2f, 0x8e, 0xfc, 0x3f, 0xfa, 0xf2, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x88, 0x68, 0x35, 0xda, 0x03, 0x00, 0x00,
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
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListReply, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error)
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignReply, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/pb.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListReply, error) {
	out := new(GetUserListReply)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error) {
	out := new(GetUserInfoReply)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignReply, error) {
	out := new(SignReply)
	err := c.cc.Invoke(ctx, "/pb.UserService/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListReply, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoReply, error)
	Sign(context.Context, *SignRequest) (*SignReply, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServiceServer) GetUserList(ctx context.Context, req *GetUserListRequest) (*GetUserListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (*UnimplementedUserServiceServer) GetUserInfo(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (*UnimplementedUserServiceServer) Sign(ctx context.Context, req *SignRequest) (*SignReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _UserService_GetUserList_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _UserService_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userservice.proto",
}