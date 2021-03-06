// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/acm/common.proto

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

type ResultCode int32

const (
	ResultCode_Success     ResultCode = 0
	ResultCode_MaxLimit    ResultCode = 1
	ResultCode_Repeated    ResultCode = 2
	ResultCode_NotExisted  ResultCode = 3
	ResultCode_DBException ResultCode = 4
	ResultCode_Empty       ResultCode = 5
)

var ResultCode_name = map[int32]string{
	0: "Success",
	1: "MaxLimit",
	2: "Repeated",
	3: "NotExisted",
	4: "DBException",
	5: "Empty",
}

var ResultCode_value = map[string]int32{
	"Success":     0,
	"MaxLimit":    1,
	"Repeated":    2,
	"NotExisted":  3,
	"DBException": 4,
	"Empty":       5,
}

func (x ResultCode) String() string {
	return proto.EnumName(ResultCode_name, int32(x))
}

func (ResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e0fbd99cee7f5b8, []int{0}
}

type RequestInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e0fbd99cee7f5b8, []int{0}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyStatus struct {
	Code                 ResultCode `protobuf:"varint,1,opt,name=code,proto3,enum=omo.msp.acm.ResultCode" json:"code,omitempty"`
	Msg                  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReplyStatus) Reset()         { *m = ReplyStatus{} }
func (m *ReplyStatus) String() string { return proto.CompactTextString(m) }
func (*ReplyStatus) ProtoMessage()    {}
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e0fbd99cee7f5b8, []int{1}
}

func (m *ReplyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatus.Unmarshal(m, b)
}
func (m *ReplyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatus.Marshal(b, m, deterministic)
}
func (m *ReplyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatus.Merge(m, src)
}
func (m *ReplyStatus) XXX_Size() int {
	return xxx_messageInfo_ReplyStatus.Size(m)
}
func (m *ReplyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatus proto.InternalMessageInfo

func (m *ReplyStatus) GetCode() ResultCode {
	if m != nil {
		return m.Code
	}
	return ResultCode_Success
}

func (m *ReplyStatus) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ReplyInfo struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyInfo) Reset()         { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()    {}
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e0fbd99cee7f5b8, []int{2}
}

func (m *ReplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInfo.Unmarshal(m, b)
}
func (m *ReplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInfo.Merge(m, src)
}
func (m *ReplyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyInfo.Size(m)
}
func (m *ReplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInfo proto.InternalMessageInfo

func (m *ReplyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("omo.msp.acm.ResultCode", ResultCode_name, ResultCode_value)
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.acm.RequestInfo")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.acm.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.acm.ReplyInfo")
}

func init() {
	proto.RegisterFile("proto/acm/common.proto", fileDescriptor_5e0fbd99cee7f5b8)
}

var fileDescriptor_5e0fbd99cee7f5b8 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbf, 0xe9, 0xaf, 0x6f, 0x33, 0x91, 0xba, 0x2c, 0x68, 0x83, 0x27, 0xc9, 0x49, 0x14,
	0x52, 0xa9, 0xc7, 0xde, 0xaa, 0x39, 0x08, 0x55, 0x61, 0x7b, 0xf3, 0x52, 0xd6, 0xcd, 0x2a, 0x81,
	0x4e, 0x66, 0xcd, 0x6e, 0x20, 0xfd, 0xef, 0x65, 0x97, 0x12, 0x45, 0xf0, 0x36, 0xef, 0x31, 0xef,
	0x7d, 0x86, 0x81, 0x73, 0xd3, 0x90, 0xa3, 0x85, 0x54, 0xb8, 0x50, 0x84, 0x48, 0x75, 0x1e, 0x0c,
	0x9e, 0x10, 0x52, 0x8e, 0xd6, 0xe4, 0x52, 0x61, 0xb6, 0x82, 0x44, 0xe8, 0xcf, 0x56, 0x5b, 0xf7,
	0x58, 0xbf, 0x13, 0x67, 0x30, 0x6c, 0xab, 0x32, 0x8d, 0x2e, 0xa3, 0xab, 0x58, 0xf8, 0x91, 0x5f,
	0xc0, 0x94, 0x8c, 0x6e, 0xa4, 0xa3, 0x26, 0x1d, 0x04, 0xbb, 0xd7, 0xd9, 0xc6, 0x87, 0xcd, 0xfe,
	0xb0, 0x75, 0xd2, 0xb5, 0x96, 0xdf, 0xc0, 0x48, 0x51, 0xa9, 0x43, 0x7a, 0xb6, 0x9c, 0xe7, 0x3f,
	0x38, 0xb9, 0xd0, 0xb6, 0xdd, 0xbb, 0x7b, 0x2a, 0xb5, 0x08, 0x4b, 0x9e, 0x84, 0xf6, 0xe3, 0x58,
	0xe9, 0xc7, 0xec, 0x05, 0xe2, 0xd0, 0xf6, 0xc7, 0x21, 0xb7, 0x30, 0xb1, 0x81, 0x13, 0x32, 0xc9,
	0x32, 0xfd, 0xd5, 0xdf, 0xdf, 0x21, 0x8e, 0x7b, 0xd7, 0x12, 0xe0, 0x1b, 0xcb, 0x13, 0xf8, 0xbf,
	0x6d, 0x95, 0xd2, 0xd6, 0xb2, 0x7f, 0xfc, 0x04, 0xa6, 0x4f, 0xb2, 0xdb, 0x54, 0x58, 0x39, 0x16,
	0x79, 0x25, 0xb4, 0xd1, 0xd2, 0xe9, 0x92, 0x0d, 0xf8, 0x0c, 0xe0, 0x99, 0x5c, 0xd1, 0x55, 0xd6,
	0xeb, 0x21, 0x3f, 0x85, 0xe4, 0x61, 0x5d, 0x74, 0x4a, 0x1b, 0x57, 0x51, 0xcd, 0x46, 0x3c, 0x86,
	0x71, 0x81, 0xc6, 0x1d, 0xd8, 0x78, 0x3d, 0x7f, 0x3d, 0xeb, 0xbf, 0xbc, 0x22, 0xa4, 0x1d, 0x5a,
	0xb3, 0x93, 0x0a, 0xdf, 0x26, 0xc1, 0xbe, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xf2, 0xa2,
	0x99, 0x85, 0x01, 0x00, 0x00,
}
