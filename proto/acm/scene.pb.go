// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/acm/scene.proto

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

type SceneLink struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Scene                string   `protobuf:"bytes,8,opt,name=scene,proto3" json:"scene,omitempty"`
	Type                 uint32   `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	Status               uint32   `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	Links                []string `protobuf:"bytes,11,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SceneLink) Reset()         { *m = SceneLink{} }
func (m *SceneLink) String() string { return proto.CompactTextString(m) }
func (*SceneLink) ProtoMessage()    {}
func (*SceneLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_121d39c1cf6c1956, []int{0}
}

func (m *SceneLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SceneLink.Unmarshal(m, b)
}
func (m *SceneLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SceneLink.Marshal(b, m, deterministic)
}
func (m *SceneLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SceneLink.Merge(m, src)
}
func (m *SceneLink) XXX_Size() int {
	return xxx_messageInfo_SceneLink.Size(m)
}
func (m *SceneLink) XXX_DiscardUnknown() {
	xxx_messageInfo_SceneLink.DiscardUnknown(m)
}

var xxx_messageInfo_SceneLink proto.InternalMessageInfo

func (m *SceneLink) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SceneLink) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SceneLink) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *SceneLink) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *SceneLink) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SceneLink) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *SceneLink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SceneLink) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *SceneLink) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SceneLink) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SceneLink) GetLinks() []string {
	if m != nil {
		return m.Links
	}
	return nil
}

type ReqSceneAdd struct {
	Scene                string   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Type                 uint32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Links                []string `protobuf:"bytes,4,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSceneAdd) Reset()         { *m = ReqSceneAdd{} }
func (m *ReqSceneAdd) String() string { return proto.CompactTextString(m) }
func (*ReqSceneAdd) ProtoMessage()    {}
func (*ReqSceneAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_121d39c1cf6c1956, []int{1}
}

func (m *ReqSceneAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneAdd.Unmarshal(m, b)
}
func (m *ReqSceneAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneAdd.Marshal(b, m, deterministic)
}
func (m *ReqSceneAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneAdd.Merge(m, src)
}
func (m *ReqSceneAdd) XXX_Size() int {
	return xxx_messageInfo_ReqSceneAdd.Size(m)
}
func (m *ReqSceneAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneAdd proto.InternalMessageInfo

func (m *ReqSceneAdd) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqSceneAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqSceneAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqSceneAdd) GetLinks() []string {
	if m != nil {
		return m.Links
	}
	return nil
}

type ReplySceneLink struct {
	Info                 *SceneLink   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySceneLink) Reset()         { *m = ReplySceneLink{} }
func (m *ReplySceneLink) String() string { return proto.CompactTextString(m) }
func (*ReplySceneLink) ProtoMessage()    {}
func (*ReplySceneLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_121d39c1cf6c1956, []int{2}
}

func (m *ReplySceneLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySceneLink.Unmarshal(m, b)
}
func (m *ReplySceneLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySceneLink.Marshal(b, m, deterministic)
}
func (m *ReplySceneLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySceneLink.Merge(m, src)
}
func (m *ReplySceneLink) XXX_Size() int {
	return xxx_messageInfo_ReplySceneLink.Size(m)
}
func (m *ReplySceneLink) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySceneLink.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySceneLink proto.InternalMessageInfo

func (m *ReplySceneLink) GetInfo() *SceneLink {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplySceneLink) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqSceneStatus struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               uint32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSceneStatus) Reset()         { *m = ReqSceneStatus{} }
func (m *ReqSceneStatus) String() string { return proto.CompactTextString(m) }
func (*ReqSceneStatus) ProtoMessage()    {}
func (*ReqSceneStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_121d39c1cf6c1956, []int{3}
}

func (m *ReqSceneStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneStatus.Unmarshal(m, b)
}
func (m *ReqSceneStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneStatus.Marshal(b, m, deterministic)
}
func (m *ReqSceneStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneStatus.Merge(m, src)
}
func (m *ReqSceneStatus) XXX_Size() int {
	return xxx_messageInfo_ReqSceneStatus.Size(m)
}
func (m *ReqSceneStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneStatus proto.InternalMessageInfo

func (m *ReqSceneStatus) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSceneStatus) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReqSceneStatus) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqSceneStatus) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ReplySceneList struct {
	Total                uint64       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageNow              uint32       `protobuf:"varint,2,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	PageMax              uint32       `protobuf:"varint,3,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	List                 []*SceneLink `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySceneList) Reset()         { *m = ReplySceneList{} }
func (m *ReplySceneList) String() string { return proto.CompactTextString(m) }
func (*ReplySceneList) ProtoMessage()    {}
func (*ReplySceneList) Descriptor() ([]byte, []int) {
	return fileDescriptor_121d39c1cf6c1956, []int{4}
}

func (m *ReplySceneList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySceneList.Unmarshal(m, b)
}
func (m *ReplySceneList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySceneList.Marshal(b, m, deterministic)
}
func (m *ReplySceneList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySceneList.Merge(m, src)
}
func (m *ReplySceneList) XXX_Size() int {
	return xxx_messageInfo_ReplySceneList.Size(m)
}
func (m *ReplySceneList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySceneList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySceneList proto.InternalMessageInfo

func (m *ReplySceneList) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplySceneList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplySceneList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplySceneList) GetList() []*SceneLink {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*SceneLink)(nil), "omo.msp.acm.SceneLink")
	proto.RegisterType((*ReqSceneAdd)(nil), "omo.msp.acm.ReqSceneAdd")
	proto.RegisterType((*ReplySceneLink)(nil), "omo.msp.acm.ReplySceneLink")
	proto.RegisterType((*ReqSceneStatus)(nil), "omo.msp.acm.ReqSceneStatus")
	proto.RegisterType((*ReplySceneList)(nil), "omo.msp.acm.ReplySceneList")
}

func init() {
	proto.RegisterFile("proto/acm/scene.proto", fileDescriptor_121d39c1cf6c1956)
}

var fileDescriptor_121d39c1cf6c1956 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x14, 0xec, 0x7e, 0x34, 0x69, 0x9c, 0xb6, 0x42, 0x16, 0x04, 0x2b, 0xbd, 0x44, 0x7b, 0x8a, 0x38,
	0xa4, 0xa8, 0x1c, 0x39, 0xa0, 0x14, 0xa1, 0x0a, 0x89, 0x00, 0x32, 0xe2, 0xc2, 0xa5, 0x32, 0xeb,
	0xd7, 0x6a, 0xd5, 0xb5, 0xbd, 0xac, 0x9d, 0x7e, 0x9c, 0x91, 0xf8, 0x27, 0xfc, 0x4f, 0xe4, 0xe7,
	0xec, 0x26, 0x29, 0xab, 0xf4, 0xe6, 0xf1, 0x3c, 0xcf, 0x78, 0xc6, 0x9b, 0x90, 0x17, 0x55, 0x6d,
	0x9c, 0x39, 0x15, 0xb9, 0x3a, 0xb5, 0x39, 0x68, 0x98, 0x21, 0xa6, 0x43, 0xa3, 0xcc, 0x4c, 0xd9,
	0x6a, 0x26, 0x72, 0x35, 0x1e, 0xad, 0x67, 0x72, 0xa3, 0x94, 0xd1, 0x61, 0x28, 0xfb, 0x1d, 0x93,
	0xc1, 0x37, 0x7f, 0xe8, 0x53, 0xa1, 0x6f, 0xe8, 0x33, 0x92, 0x2c, 0x0b, 0xc9, 0xa2, 0x49, 0x34,
	0x1d, 0x70, 0xbf, 0xa4, 0xc7, 0x24, 0x2e, 0x24, 0x8b, 0x27, 0xd1, 0x34, 0xe5, 0x71, 0x21, 0x29,
	0x23, 0xfd, 0xbc, 0x06, 0xe1, 0x40, 0xb2, 0x64, 0x12, 0x4d, 0x13, 0xde, 0x40, 0xcf, 0x2c, 0x2b,
	0x89, 0x4c, 0x1a, 0x98, 0x15, 0x6c, 0xcf, 0x98, 0x9a, 0xed, 0xa3, 0x72, 0x03, 0xe9, 0x98, 0x1c,
	0x98, 0x0a, 0x6a, 0xa4, 0x7a, 0x48, 0xb5, 0x98, 0x52, 0x92, 0x6a, 0xa1, 0x80, 0xf5, 0x71, 0x1f,
	0xd7, 0xf4, 0x39, 0xd9, 0xc7, 0x84, 0xec, 0x00, 0x37, 0x03, 0xf0, 0x93, 0xee, 0xa1, 0x02, 0x36,
	0x98, 0x44, 0xd3, 0x23, 0x8e, 0x6b, 0x3a, 0x22, 0x3d, 0xeb, 0x84, 0x5b, 0x5a, 0x46, 0x70, 0x77,
	0x85, 0xbc, 0x42, 0x59, 0xe8, 0x1b, 0xcb, 0x86, 0x93, 0xc4, 0x2b, 0x20, 0xc8, 0x0a, 0x32, 0xe4,
	0xf0, 0x0b, 0x7b, 0x98, 0x4b, 0xb9, 0xb6, 0x89, 0xba, 0x6c, 0xe2, 0x0d, 0x9b, 0xcd, 0x00, 0xc9,
	0xa3, 0x00, 0xad, 0x55, 0xba, 0x69, 0xa5, 0xc9, 0x31, 0x87, 0xaa, 0x7c, 0x58, 0x97, 0xfe, 0x8a,
	0xa4, 0x85, 0xbe, 0x32, 0x68, 0x36, 0x3c, 0x1b, 0xcd, 0x36, 0x9e, 0x6d, 0xd6, 0x4e, 0x71, 0x9c,
	0xa1, 0xaf, 0xdb, 0x58, 0x31, 0x4e, 0xb3, 0xad, 0xe9, 0x20, 0x8c, 0x7c, 0x13, 0x38, 0x2b, 0xbd,
	0x5f, 0x88, 0x16, 0x98, 0x8e, 0x47, 0x1e, 0x6d, 0xa9, 0xae, 0xcb, 0x7a, 0x22, 0x9d, 0xb9, 0xd3,
	0x50, 0xe3, 0x63, 0x0f, 0x78, 0x00, 0xd9, 0x9f, 0x68, 0x3b, 0x9e, 0x75, 0x7e, 0xd0, 0x19, 0x27,
	0x4a, 0x34, 0x4c, 0x79, 0x00, 0xfe, 0x9b, 0xa8, 0xc4, 0x35, 0x7c, 0x36, 0x77, 0x2b, 0xcf, 0x06,
	0x36, 0xcc, 0x42, 0xdc, 0xa3, 0xe7, 0x8a, 0x59, 0x88, 0x7b, 0x5f, 0x54, 0x59, 0x58, 0x87, 0x7d,
	0xee, 0x28, 0xca, 0xcf, 0x9c, 0xfd, 0x4d, 0xc8, 0x61, 0x08, 0x0d, 0xf5, 0x6d, 0x91, 0x03, 0x9d,
	0x93, 0xde, 0x5c, 0xca, 0x2f, 0x1a, 0xe8, 0xe3, 0xce, 0xda, 0x77, 0x1f, 0x9f, 0x74, 0xb4, 0xd9,
	0xe8, 0x66, 0x7b, 0x5e, 0xe2, 0x02, 0x5c, 0xa7, 0xc4, 0x12, 0xac, 0xfb, 0xa8, 0xaf, 0xcc, 0x53,
	0x12, 0xef, 0xc8, 0x80, 0x83, 0x32, 0xb7, 0xb0, 0x5b, 0x65, 0xf4, 0xbf, 0x8a, 0xdf, 0xcf, 0xf6,
	0xe8, 0x39, 0xe9, 0x5f, 0x80, 0xc3, 0x62, 0x3b, 0x8f, 0x7f, 0x15, 0xd7, 0xb0, 0xe3, 0x12, 0xd6,
	0x65, 0x7b, 0xf4, 0x03, 0x39, 0xfc, 0x8e, 0x3f, 0xcd, 0xd5, 0x07, 0x71, 0xd2, 0x59, 0x48, 0x20,
	0x77, 0x5c, 0xe5, 0x3d, 0x39, 0x0a, 0x32, 0x0b, 0x23, 0x97, 0x25, 0xd8, 0xee, 0x0b, 0x79, 0xcf,
	0x2e, 0x91, 0x70, 0x97, 0xf3, 0x97, 0x3f, 0xd6, 0xff, 0x5e, 0x6f, 0x8d, 0x32, 0x97, 0xca, 0x56,
	0x97, 0x22, 0x57, 0x3f, 0x7b, 0xb8, 0xfd, 0xe6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xd2,
	0x46, 0x69, 0xdd, 0x04, 0x00, 0x00,
}
