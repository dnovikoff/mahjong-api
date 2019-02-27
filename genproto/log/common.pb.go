// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log/common.proto

package log

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

// Represents score changing.
type Changes struct {
	Changes []int64 `protobuf:"zigzag64,1,rep,packed,name=changes,proto3" json:"changes,omitempty"`
	// After changes applied
	Results              []int64  `protobuf:"zigzag64,2,rep,packed,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Changes) Reset()         { *m = Changes{} }
func (m *Changes) String() string { return proto.CompactTextString(m) }
func (*Changes) ProtoMessage()    {}
func (*Changes) Descriptor() ([]byte, []int) {
	return fileDescriptor_72958bde97e673bf, []int{0}
}

func (m *Changes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Changes.Unmarshal(m, b)
}
func (m *Changes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Changes.Marshal(b, m, deterministic)
}
func (m *Changes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Changes.Merge(m, src)
}
func (m *Changes) XXX_Size() int {
	return xxx_messageInfo_Changes.Size(m)
}
func (m *Changes) XXX_DiscardUnknown() {
	xxx_messageInfo_Changes.DiscardUnknown(m)
}

var xxx_messageInfo_Changes proto.InternalMessageInfo

func (m *Changes) GetChanges() []int64 {
	if m != nil {
		return m.Changes
	}
	return nil
}

func (m *Changes) GetResults() []int64 {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*Changes)(nil), "log.Changes")
}

func init() { proto.RegisterFile("log/common.proto", fileDescriptor_72958bde97e673bf) }

var fileDescriptor_72958bde97e673bf = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xc9, 0x4f, 0xd7,
	0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0xc9,
	0x4f, 0x57, 0xb2, 0xe5, 0x62, 0x77, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x2d, 0x16, 0x92, 0xe0, 0x62,
	0x4f, 0x86, 0x30, 0x25, 0x18, 0x15, 0x98, 0x35, 0x84, 0x82, 0x60, 0x5c, 0x90, 0x4c, 0x51, 0x6a,
	0x71, 0x69, 0x4e, 0x49, 0xb1, 0x04, 0x13, 0x44, 0x06, 0xca, 0x75, 0xd2, 0x8f, 0xd2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xc9, 0xcb, 0x2f, 0xcb, 0xcc, 0xce,
	0x4f, 0x4b, 0xd3, 0xcf, 0x4d, 0xcc, 0xc8, 0xca, 0xcf, 0x4b, 0xd7, 0x4d, 0x2c, 0xc8, 0xd4, 0x4f,
	0x4f, 0xcd, 0x03, 0x5b, 0xa8, 0x9f, 0x93, 0x9f, 0x9e, 0xc4, 0x06, 0x66, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x2e, 0xb2, 0xf2, 0x46, 0x8f, 0x00, 0x00, 0x00,
}
