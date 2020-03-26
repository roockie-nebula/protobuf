// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nullfloat.proto

package nullfloat

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

// NullFloat ...
type NullFloat struct {
	// Represents the actual Float
	Float float64 `protobuf:"fixed64,1,opt,name=float,proto3" json:"float,omitempty"`
	// is set to true if the float is null
	IsNotNull            bool     `protobuf:"varint,2,opt,name=is_not_null,json=isNotNull,proto3" json:"is_not_null,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullFloat) Reset()         { *m = NullFloat{} }
func (m *NullFloat) String() string { return proto.CompactTextString(m) }
func (*NullFloat) ProtoMessage()    {}
func (*NullFloat) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa0aaeca416962dd, []int{0}
}

func (m *NullFloat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullFloat.Unmarshal(m, b)
}
func (m *NullFloat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullFloat.Marshal(b, m, deterministic)
}
func (m *NullFloat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullFloat.Merge(m, src)
}
func (m *NullFloat) XXX_Size() int {
	return xxx_messageInfo_NullFloat.Size(m)
}
func (m *NullFloat) XXX_DiscardUnknown() {
	xxx_messageInfo_NullFloat.DiscardUnknown(m)
}

var xxx_messageInfo_NullFloat proto.InternalMessageInfo

func (m *NullFloat) GetFloat() float64 {
	if m != nil {
		return m.Float
	}
	return 0
}

func (m *NullFloat) GetIsNotNull() bool {
	if m != nil {
		return m.IsNotNull
	}
	return false
}

func init() {
	proto.RegisterType((*NullFloat)(nil), "dkfbasel.protobuf.NullFloat")
}

func init() { proto.RegisterFile("nullfloat.proto", fileDescriptor_aa0aaeca416962dd) }

var fileDescriptor_aa0aaeca416962dd = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2b, 0xcd, 0xc9,
	0x49, 0xcb, 0xc9, 0x4f, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xc9, 0x4e,
	0x4b, 0x4a, 0x2c, 0x4e, 0xcd, 0x81, 0xf0, 0x93, 0x4a, 0xd3, 0x94, 0x1c, 0xb9, 0x38, 0xfd, 0x4a,
	0x73, 0x72, 0xdc, 0x40, 0xaa, 0x84, 0x44, 0xb8, 0x58, 0xc1, 0xca, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x18, 0x83, 0x20, 0x1c, 0x21, 0x39, 0x2e, 0xee, 0xcc, 0xe2, 0xf8, 0xbc, 0xfc, 0x92, 0x78, 0x90,
	0x79, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x9c, 0x99, 0xc5, 0x7e, 0xf9, 0x25, 0x20, 0xad,
	0x4e, 0x7a, 0x51, 0x3a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x30,
	0x2b, 0xf4, 0x61, 0x56, 0xe8, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0xc3, 0xdd, 0x92, 0xc4, 0x06,
	0x96, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x91, 0xc2, 0x1a, 0xdb, 0x9f, 0x00, 0x00, 0x00,
}