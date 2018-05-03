// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello.proto

/*
Package testdata is a generated protocol buffer package.

It is generated from these files:
	proto/hello.proto

It has these top-level messages:
	Pair
*/
package testdata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Pair struct {
	// the comment does not match syntax
	// `foo:"bar"`
	Key   string `foo:"bar" json:"key,omitempty" protobuf:"bytes,1,opt,name=key"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "testdata.Pair")
}

func init() { proto.RegisterFile("proto/hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0xb3, 0x85, 0x38, 0x4a, 0x52, 0x8b, 0x4b,
	0x52, 0x12, 0x4b, 0x12, 0x95, 0xf4, 0xb8, 0x58, 0x02, 0x12, 0x33, 0x8b, 0x84, 0x04, 0xb8, 0x98,
	0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6,
	0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0xb0, 0x18, 0x84, 0x93, 0xc4, 0x06, 0x36, 0xc0, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x13, 0xd2, 0xe1, 0x55, 0x00, 0x00, 0x00,
}
