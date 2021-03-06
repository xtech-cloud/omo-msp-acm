// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/acm/menu.proto

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

type MenuInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Path                 string   `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
	Method               string   `protobuf:"bytes,8,opt,name=method,proto3" json:"method,omitempty"`
	Creator              string   `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,10,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MenuInfo) Reset()         { *m = MenuInfo{} }
func (m *MenuInfo) String() string { return proto.CompactTextString(m) }
func (*MenuInfo) ProtoMessage()    {}
func (*MenuInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f186ad4ef5b8c0, []int{0}
}

func (m *MenuInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MenuInfo.Unmarshal(m, b)
}
func (m *MenuInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MenuInfo.Marshal(b, m, deterministic)
}
func (m *MenuInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MenuInfo.Merge(m, src)
}
func (m *MenuInfo) XXX_Size() int {
	return xxx_messageInfo_MenuInfo.Size(m)
}
func (m *MenuInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MenuInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MenuInfo proto.InternalMessageInfo

func (m *MenuInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MenuInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MenuInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *MenuInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *MenuInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MenuInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MenuInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MenuInfo) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *MenuInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MenuInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqMenuAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Method               string   `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMenuAdd) Reset()         { *m = ReqMenuAdd{} }
func (m *ReqMenuAdd) String() string { return proto.CompactTextString(m) }
func (*ReqMenuAdd) ProtoMessage()    {}
func (*ReqMenuAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f186ad4ef5b8c0, []int{1}
}

func (m *ReqMenuAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMenuAdd.Unmarshal(m, b)
}
func (m *ReqMenuAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMenuAdd.Marshal(b, m, deterministic)
}
func (m *ReqMenuAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMenuAdd.Merge(m, src)
}
func (m *ReqMenuAdd) XXX_Size() int {
	return xxx_messageInfo_ReqMenuAdd.Size(m)
}
func (m *ReqMenuAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMenuAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMenuAdd proto.InternalMessageInfo

func (m *ReqMenuAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqMenuAdd) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ReqMenuAdd) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ReqMenuAdd) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ReqMenuAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyMenuInfo struct {
	Info                 *MenuInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyMenuInfo) Reset()         { *m = ReplyMenuInfo{} }
func (m *ReplyMenuInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyMenuInfo) ProtoMessage()    {}
func (*ReplyMenuInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f186ad4ef5b8c0, []int{2}
}

func (m *ReplyMenuInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMenuInfo.Unmarshal(m, b)
}
func (m *ReplyMenuInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMenuInfo.Marshal(b, m, deterministic)
}
func (m *ReplyMenuInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMenuInfo.Merge(m, src)
}
func (m *ReplyMenuInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyMenuInfo.Size(m)
}
func (m *ReplyMenuInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMenuInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMenuInfo proto.InternalMessageInfo

func (m *ReplyMenuInfo) GetInfo() *MenuInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyMenuInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyMenuList struct {
	Number               int32        `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*MenuInfo  `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyMenuList) Reset()         { *m = ReplyMenuList{} }
func (m *ReplyMenuList) String() string { return proto.CompactTextString(m) }
func (*ReplyMenuList) ProtoMessage()    {}
func (*ReplyMenuList) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f186ad4ef5b8c0, []int{3}
}

func (m *ReplyMenuList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMenuList.Unmarshal(m, b)
}
func (m *ReplyMenuList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMenuList.Marshal(b, m, deterministic)
}
func (m *ReplyMenuList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMenuList.Merge(m, src)
}
func (m *ReplyMenuList) XXX_Size() int {
	return xxx_messageInfo_ReplyMenuList.Size(m)
}
func (m *ReplyMenuList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMenuList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMenuList proto.InternalMessageInfo

func (m *ReplyMenuList) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyMenuList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyMenuList) GetList() []*MenuInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqMenuUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Path                 string   `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Method               string   `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMenuUpdate) Reset()         { *m = ReqMenuUpdate{} }
func (m *ReqMenuUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqMenuUpdate) ProtoMessage()    {}
func (*ReqMenuUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f186ad4ef5b8c0, []int{4}
}

func (m *ReqMenuUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMenuUpdate.Unmarshal(m, b)
}
func (m *ReqMenuUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMenuUpdate.Marshal(b, m, deterministic)
}
func (m *ReqMenuUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMenuUpdate.Merge(m, src)
}
func (m *ReqMenuUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqMenuUpdate.Size(m)
}
func (m *ReqMenuUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMenuUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMenuUpdate proto.InternalMessageInfo

func (m *ReqMenuUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqMenuUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqMenuUpdate) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ReqMenuUpdate) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ReqMenuUpdate) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ReqMenuUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func init() {
	proto.RegisterType((*MenuInfo)(nil), "omo.msp.acm.MenuInfo")
	proto.RegisterType((*ReqMenuAdd)(nil), "omo.msp.acm.ReqMenuAdd")
	proto.RegisterType((*ReplyMenuInfo)(nil), "omo.msp.acm.ReplyMenuInfo")
	proto.RegisterType((*ReplyMenuList)(nil), "omo.msp.acm.ReplyMenuList")
	proto.RegisterType((*ReqMenuUpdate)(nil), "omo.msp.acm.ReqMenuUpdate")
}

func init() {
	proto.RegisterFile("proto/acm/menu.proto", fileDescriptor_97f186ad4ef5b8c0)
}

var fileDescriptor_97f186ad4ef5b8c0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x8e, 0x7e, 0xac, 0xd8, 0x63, 0x5a, 0xca, 0xd2, 0x38, 0x8b, 0x4e, 0x46, 0x27, 0xf7, 0xa2,
	0x14, 0xf7, 0xd8, 0x43, 0xea, 0x1c, 0x5a, 0x0a, 0x2d, 0x85, 0x0d, 0xbd, 0xf4, 0x12, 0x14, 0xed,
	0x84, 0x08, 0xb4, 0x5a, 0x45, 0x5a, 0x05, 0x02, 0x3d, 0xf6, 0x05, 0xfa, 0x7c, 0x7d, 0x83, 0x3e,
	0x45, 0xd9, 0x91, 0xe5, 0x58, 0x8e, 0xea, 0xe0, 0xdb, 0xce, 0x7c, 0x1f, 0x33, 0xdf, 0xcc, 0x37,
	0x2c, 0xbc, 0x2e, 0x2b, 0x6d, 0xf4, 0x59, 0x92, 0xaa, 0x33, 0x85, 0x45, 0x13, 0x53, 0xc8, 0xa6,
	0x5a, 0xe9, 0x58, 0xd5, 0x65, 0x9c, 0xa4, 0x2a, 0x9c, 0x3d, 0x52, 0x52, 0xad, 0x94, 0x2e, 0x5a,
	0x52, 0xf4, 0xd7, 0x81, 0xf1, 0x57, 0x2c, 0x9a, 0xcf, 0xc5, 0x8d, 0x66, 0xaf, 0xc0, 0x6b, 0x32,
	0xc9, 0x9d, 0xb9, 0xb3, 0x98, 0x08, 0xfb, 0x64, 0x2f, 0xc1, 0xcd, 0x24, 0x77, 0xe7, 0xce, 0xc2,
	0x17, 0x6e, 0x26, 0x19, 0x87, 0xe3, 0xb4, 0xc2, 0xc4, 0xa0, 0xe4, 0xde, 0xdc, 0x59, 0x78, 0xa2,
	0x0b, 0x2d, 0xd2, 0x94, 0x92, 0x10, 0xbf, 0x45, 0xd6, 0x21, 0x63, 0xe0, 0x17, 0x89, 0x42, 0x3e,
	0xa2, 0xb2, 0xf4, 0xb6, 0x39, 0xf3, 0x50, 0x22, 0x0f, 0xda, 0x9c, 0x7d, 0xdb, 0x5c, 0x99, 0x98,
	0x5b, 0x7e, 0xdc, 0xe6, 0xec, 0x9b, 0xcd, 0x20, 0x50, 0x68, 0x6e, 0xb5, 0xe4, 0x63, 0xca, 0xae,
	0xa3, 0x8d, 0x0e, 0x5d, 0xf1, 0x09, 0x01, 0x5d, 0xc8, 0x42, 0x18, 0xeb, 0x12, 0x2b, 0x82, 0x80,
	0xa0, 0x4d, 0x1c, 0xfd, 0x04, 0x10, 0x78, 0x67, 0xc7, 0x5d, 0xc9, 0x47, 0x5d, 0xce, 0x80, 0x2e,
	0x77, 0x40, 0x97, 0x37, 0xa8, 0xcb, 0xef, 0xe9, 0xda, 0xee, 0x3e, 0xda, 0xe9, 0x9e, 0xc3, 0x0b,
	0x81, 0x65, 0xfe, 0xb0, 0x59, 0xf7, 0x1b, 0xf0, 0xb3, 0xe2, 0x46, 0x93, 0x80, 0xe9, 0xf2, 0x24,
	0xde, 0xf2, 0x2b, 0xee, 0x48, 0x82, 0x28, 0xec, 0x2d, 0x04, 0xb5, 0x49, 0x4c, 0x53, 0x93, 0xb2,
	0xe9, 0x92, 0xf7, 0xc8, 0x54, 0xf6, 0x92, 0x70, 0xb1, 0xe6, 0x45, 0xbf, 0x9c, 0xad, 0x76, 0x5f,
	0xb2, 0xda, 0x58, 0xcd, 0x45, 0xa3, 0xae, 0xb1, 0xa2, 0x86, 0x23, 0xb1, 0x8e, 0x0e, 0xaf, 0x6d,
	0x85, 0xe7, 0x59, 0x6d, 0xb8, 0x37, 0xf7, 0xf6, 0x08, 0xb7, 0x94, 0xe8, 0x37, 0xc9, 0xa0, 0x9d,
	0x7f, 0xa7, 0x7b, 0x18, 0x38, 0xb2, 0xce, 0x08, 0x77, 0xc0, 0x08, 0x6f, 0xc0, 0x08, 0x7f, 0xd0,
	0x88, 0xd1, 0x7f, 0x8d, 0x08, 0xfa, 0x46, 0x2c, 0xff, 0xb8, 0x30, 0xb5, 0x82, 0x2e, 0xb1, 0xba,
	0xcf, 0x52, 0x64, 0xe7, 0x10, 0xac, 0xa4, 0xfc, 0x56, 0x20, 0x3b, 0xdd, 0x19, 0xbd, 0xbb, 0x95,
	0x30, 0x7c, 0xba, 0x93, 0x6e, 0xd0, 0xe8, 0x88, 0x7d, 0x80, 0xe0, 0x13, 0x1a, 0x5b, 0x60, 0x77,
	0x77, 0x77, 0x0d, 0xd6, 0xc6, 0xb2, 0x9e, 0xa9, 0x70, 0x0e, 0x13, 0x81, 0x4a, 0xdf, 0xe3, 0xfe,
	0x22, 0xb3, 0xa7, 0x45, 0x7a, 0x12, 0x56, 0x79, 0x7e, 0xb8, 0x04, 0x7b, 0x1c, 0xd1, 0x11, 0xfb,
	0x08, 0xd0, 0x3a, 0x74, 0x91, 0xd4, 0xc8, 0xc2, 0xa1, 0x4d, 0xb4, 0xf8, 0xfe, 0x51, 0x2e, 0x4e,
	0x7f, 0x9c, 0x6c, 0xfe, 0x9a, 0xf7, 0x5a, 0xe9, 0x2b, 0x55, 0x97, 0x57, 0x49, 0xaa, 0xae, 0x03,
	0x4a, 0xbf, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xdb, 0x37, 0x62, 0xae, 0x04, 0x00, 0x00,
}
