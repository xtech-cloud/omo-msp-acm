// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/acm/role.proto

package omo_msp_acm

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RoleInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	Creator              string   `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Menus                []string `protobuf:"bytes,9,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleInfo) Reset()         { *m = RoleInfo{} }
func (m *RoleInfo) String() string { return proto.CompactTextString(m) }
func (*RoleInfo) ProtoMessage()    {}
func (*RoleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_599166315131034b, []int{0}
}

func (m *RoleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleInfo.Unmarshal(m, b)
}
func (m *RoleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleInfo.Marshal(b, m, deterministic)
}
func (m *RoleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleInfo.Merge(m, src)
}
func (m *RoleInfo) XXX_Size() int {
	return xxx_messageInfo_RoleInfo.Size(m)
}
func (m *RoleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoleInfo proto.InternalMessageInfo

func (m *RoleInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RoleInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RoleInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RoleInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *RoleInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *RoleInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RoleInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RoleInfo) GetMenus() []string {
	if m != nil {
		return m.Menus
	}
	return nil
}

type ReqRoleAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Menus                []string `protobuf:"bytes,4,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRoleAdd) Reset()         { *m = ReqRoleAdd{} }
func (m *ReqRoleAdd) String() string { return proto.CompactTextString(m) }
func (*ReqRoleAdd) ProtoMessage()    {}
func (*ReqRoleAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_599166315131034b, []int{1}
}

func (m *ReqRoleAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRoleAdd.Unmarshal(m, b)
}
func (m *ReqRoleAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRoleAdd.Marshal(b, m, deterministic)
}
func (m *ReqRoleAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRoleAdd.Merge(m, src)
}
func (m *ReqRoleAdd) XXX_Size() int {
	return xxx_messageInfo_ReqRoleAdd.Size(m)
}
func (m *ReqRoleAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRoleAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRoleAdd proto.InternalMessageInfo

func (m *ReqRoleAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqRoleAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqRoleAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRoleAdd) GetMenus() []string {
	if m != nil {
		return m.Menus
	}
	return nil
}

type ReplyRoleInfo struct {
	Info                 *RoleInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyRoleInfo) Reset()         { *m = ReplyRoleInfo{} }
func (m *ReplyRoleInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyRoleInfo) ProtoMessage()    {}
func (*ReplyRoleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_599166315131034b, []int{2}
}

func (m *ReplyRoleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRoleInfo.Unmarshal(m, b)
}
func (m *ReplyRoleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRoleInfo.Marshal(b, m, deterministic)
}
func (m *ReplyRoleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRoleInfo.Merge(m, src)
}
func (m *ReplyRoleInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyRoleInfo.Size(m)
}
func (m *ReplyRoleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRoleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRoleInfo proto.InternalMessageInfo

func (m *ReplyRoleInfo) GetInfo() *RoleInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyRoleInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyRoleList struct {
	Number               int32        `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*RoleInfo  `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyRoleList) Reset()         { *m = ReplyRoleList{} }
func (m *ReplyRoleList) String() string { return proto.CompactTextString(m) }
func (*ReplyRoleList) ProtoMessage()    {}
func (*ReplyRoleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_599166315131034b, []int{3}
}

func (m *ReplyRoleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRoleList.Unmarshal(m, b)
}
func (m *ReplyRoleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRoleList.Marshal(b, m, deterministic)
}
func (m *ReplyRoleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRoleList.Merge(m, src)
}
func (m *ReplyRoleList) XXX_Size() int {
	return xxx_messageInfo_ReplyRoleList.Size(m)
}
func (m *ReplyRoleList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRoleList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRoleList proto.InternalMessageInfo

func (m *ReplyRoleList) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyRoleList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyRoleList) GetList() []*RoleInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqRoleUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Menus                []string `protobuf:"bytes,5,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRoleUpdate) Reset()         { *m = ReqRoleUpdate{} }
func (m *ReqRoleUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqRoleUpdate) ProtoMessage()    {}
func (*ReqRoleUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_599166315131034b, []int{4}
}

func (m *ReqRoleUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRoleUpdate.Unmarshal(m, b)
}
func (m *ReqRoleUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRoleUpdate.Marshal(b, m, deterministic)
}
func (m *ReqRoleUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRoleUpdate.Merge(m, src)
}
func (m *ReqRoleUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqRoleUpdate.Size(m)
}
func (m *ReqRoleUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRoleUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRoleUpdate proto.InternalMessageInfo

func (m *ReqRoleUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqRoleUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqRoleUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqRoleUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRoleUpdate) GetMenus() []string {
	if m != nil {
		return m.Menus
	}
	return nil
}

type ReqRoleMenus struct {
	Role                 string   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Menus                []string `protobuf:"bytes,3,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRoleMenus) Reset()         { *m = ReqRoleMenus{} }
func (m *ReqRoleMenus) String() string { return proto.CompactTextString(m) }
func (*ReqRoleMenus) ProtoMessage()    {}
func (*ReqRoleMenus) Descriptor() ([]byte, []int) {
	return fileDescriptor_599166315131034b, []int{5}
}

func (m *ReqRoleMenus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRoleMenus.Unmarshal(m, b)
}
func (m *ReqRoleMenus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRoleMenus.Marshal(b, m, deterministic)
}
func (m *ReqRoleMenus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRoleMenus.Merge(m, src)
}
func (m *ReqRoleMenus) XXX_Size() int {
	return xxx_messageInfo_ReqRoleMenus.Size(m)
}
func (m *ReqRoleMenus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRoleMenus.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRoleMenus proto.InternalMessageInfo

func (m *ReqRoleMenus) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *ReqRoleMenus) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRoleMenus) GetMenus() []string {
	if m != nil {
		return m.Menus
	}
	return nil
}

type ReplyRoleMenu struct {
	Role                 string       `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Menus                []string     `protobuf:"bytes,3,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyRoleMenu) Reset()         { *m = ReplyRoleMenu{} }
func (m *ReplyRoleMenu) String() string { return proto.CompactTextString(m) }
func (*ReplyRoleMenu) ProtoMessage()    {}
func (*ReplyRoleMenu) Descriptor() ([]byte, []int) {
	return fileDescriptor_599166315131034b, []int{6}
}

func (m *ReplyRoleMenu) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRoleMenu.Unmarshal(m, b)
}
func (m *ReplyRoleMenu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRoleMenu.Marshal(b, m, deterministic)
}
func (m *ReplyRoleMenu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRoleMenu.Merge(m, src)
}
func (m *ReplyRoleMenu) XXX_Size() int {
	return xxx_messageInfo_ReplyRoleMenu.Size(m)
}
func (m *ReplyRoleMenu) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRoleMenu.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRoleMenu proto.InternalMessageInfo

func (m *ReplyRoleMenu) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *ReplyRoleMenu) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyRoleMenu) GetMenus() []string {
	if m != nil {
		return m.Menus
	}
	return nil
}

func init() {
	proto.RegisterType((*RoleInfo)(nil), "omo.msp.acm.RoleInfo")
	proto.RegisterType((*ReqRoleAdd)(nil), "omo.msp.acm.ReqRoleAdd")
	proto.RegisterType((*ReplyRoleInfo)(nil), "omo.msp.acm.ReplyRoleInfo")
	proto.RegisterType((*ReplyRoleList)(nil), "omo.msp.acm.ReplyRoleList")
	proto.RegisterType((*ReqRoleUpdate)(nil), "omo.msp.acm.ReqRoleUpdate")
	proto.RegisterType((*ReqRoleMenus)(nil), "omo.msp.acm.ReqRoleMenus")
	proto.RegisterType((*ReplyRoleMenu)(nil), "omo.msp.acm.ReplyRoleMenu")
}

func init() {
	proto.RegisterFile("proto/acm/role.proto", fileDescriptor_599166315131034b)
}

var fileDescriptor_599166315131034b = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x6d, 0xe2, 0x34, 0x6b, 0xbf, 0x0e, 0x84, 0xac, 0xad, 0x33, 0x3d, 0x55, 0x39, 0x95, 0x4b,
	0x87, 0xca, 0x91, 0xc3, 0xe8, 0x24, 0x98, 0x90, 0x40, 0x48, 0x2e, 0x5c, 0xb8, 0x4c, 0x69, 0xf2,
	0x4d, 0x0a, 0x8b, 0xe3, 0xcc, 0x71, 0x26, 0x71, 0xe2, 0xc2, 0x8f, 0xe4, 0x3f, 0xf0, 0x27, 0x90,
	0xdd, 0xa4, 0x4b, 0xb7, 0x64, 0x53, 0xc5, 0xcd, 0x2f, 0xef, 0xd3, 0x7b, 0xcf, 0xcf, 0x9f, 0x02,
	0x47, 0xb9, 0x92, 0x5a, 0x9e, 0x86, 0x91, 0x38, 0x55, 0x32, 0xc5, 0xb9, 0x85, 0x74, 0x24, 0x85,
	0x9c, 0x8b, 0x22, 0x9f, 0x87, 0x91, 0x98, 0x8c, 0xef, 0x46, 0x22, 0x29, 0x84, 0xcc, 0x36, 0x43,
	0xc1, 0x1f, 0x07, 0x06, 0x5c, 0xa6, 0xf8, 0x31, 0xbb, 0x92, 0xf4, 0x05, 0x90, 0x32, 0x89, 0x99,
	0x33, 0x75, 0x66, 0x43, 0x6e, 0x8e, 0xf4, 0x39, 0xb8, 0x49, 0xcc, 0xdc, 0xa9, 0x33, 0xf3, 0xb8,
	0x9b, 0xc4, 0x94, 0xc1, 0x41, 0xa4, 0x30, 0xd4, 0x18, 0x33, 0x32, 0x75, 0x66, 0x84, 0xd7, 0xd0,
	0x30, 0x65, 0x1e, 0x5b, 0xc6, 0xdb, 0x30, 0x15, 0xa4, 0x14, 0xbc, 0x2c, 0x14, 0xc8, 0xfa, 0x56,
	0xd6, 0x9e, 0xe9, 0x18, 0x7c, 0x85, 0x22, 0x54, 0xd7, 0xcc, 0xb7, 0x5f, 0x2b, 0xb4, 0xd5, 0x97,
	0x8a, 0x1d, 0x58, 0xa2, 0x86, 0x74, 0x02, 0x03, 0x99, 0xa3, 0xb2, 0xd4, 0xc0, 0x52, 0x5b, 0x4c,
	0x8f, 0xa0, 0x2f, 0x30, 0x2b, 0x0b, 0x36, 0x9c, 0x92, 0xd9, 0x90, 0x6f, 0x40, 0xf0, 0x03, 0x80,
	0xe3, 0x8d, 0xb9, 0xdc, 0x32, 0xbe, 0x4b, 0xe1, 0xb4, 0xa6, 0x70, 0x77, 0x52, 0x34, 0xbd, 0x48,
	0x97, 0x97, 0xd7, 0xf4, 0x4a, 0xe1, 0x19, 0xc7, 0x3c, 0xfd, 0xb9, 0xad, 0xf2, 0x15, 0x78, 0x49,
	0x76, 0x25, 0xad, 0xdd, 0x68, 0x71, 0x3c, 0x6f, 0xbc, 0xc5, 0xbc, 0x1e, 0xe2, 0x76, 0x84, 0xbe,
	0x06, 0xbf, 0xd0, 0xa1, 0x2e, 0x0b, 0x9b, 0x62, 0xb4, 0x60, 0xbb, 0xc3, 0x46, 0x76, 0x65, 0x79,
	0x5e, 0xcd, 0x05, 0xbf, 0x9d, 0x86, 0xdd, 0xa7, 0xa4, 0xd0, 0xe6, 0x26, 0x59, 0x29, 0xd6, 0xa8,
	0xac, 0x61, 0x9f, 0x57, 0x68, 0x7f, 0x6d, 0x13, 0x3c, 0x4d, 0x0a, 0xcd, 0xc8, 0x94, 0x3c, 0x12,
	0xdc, 0x8c, 0x04, 0xbf, 0x4c, 0x0a, 0x5b, 0xf0, 0x37, 0xfb, 0xd4, 0x2d, 0xfb, 0x53, 0xb7, 0xee,
	0xb6, 0xb6, 0x4e, 0x3a, 0x5b, 0xf7, 0xba, 0x5a, 0xef, 0x37, 0x5b, 0xff, 0x0a, 0x87, 0x55, 0x80,
	0xcf, 0x06, 0x1b, 0x37, 0xb3, 0xff, 0xf5, 0x1b, 0x9b, 0xf3, 0x8e, 0xaa, 0xdb, 0xa5, 0x4a, 0x9a,
	0xaa, 0xd7, 0x8d, 0x72, 0x8d, 0x6e, 0xab, 0xec, 0xfe, 0xc5, 0xb6, 0x9a, 0x2d, 0xfe, 0x12, 0x18,
	0x19, 0xa3, 0x15, 0xaa, 0xdb, 0x24, 0x42, 0x7a, 0x06, 0xfe, 0x32, 0x8e, 0xbf, 0x64, 0x48, 0x4f,
	0xee, 0x29, 0xd6, 0x9b, 0x3c, 0x99, 0x3c, 0xb4, 0xaa, 0x1f, 0x26, 0xe8, 0xd1, 0x77, 0xe0, 0x5f,
	0xa0, 0x36, 0x02, 0xf7, 0x23, 0xdd, 0x94, 0x58, 0x68, 0x33, 0xf5, 0x84, 0xc2, 0x19, 0x0c, 0x39,
	0x0a, 0x79, 0x8b, 0x8f, 0x8b, 0x8c, 0x1f, 0x8a, 0xec, 0x44, 0x58, 0xa6, 0xe9, 0xfe, 0x11, 0xcc,
	0x32, 0x07, 0x3d, 0xfa, 0x01, 0x60, 0xb3, 0x52, 0xe7, 0x61, 0x81, 0x74, 0xd2, 0xd6, 0xc4, 0x86,
	0x7f, 0xe2, 0x2a, 0xef, 0x01, 0x96, 0x79, 0x8e, 0x59, 0x6c, 0xdf, 0xf1, 0x65, 0x9b, 0x8e, 0xdd,
	0x9c, 0x2e, 0x19, 0x43, 0x06, 0x3d, 0x7a, 0x01, 0x87, 0xab, 0x72, 0xad, 0x55, 0x18, 0xe9, 0xff,
	0x12, 0x3a, 0x3f, 0xf9, 0x7e, 0xbc, 0xfd, 0x0f, 0xbf, 0x95, 0x42, 0x5e, 0x8a, 0x22, 0xbf, 0x0c,
	0x23, 0xb1, 0xf6, 0xed, 0xe7, 0x37, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x54, 0x37, 0x30,
	0xca, 0x05, 0x00, 0x00,
}
