// Code generated by protoc-gen-go. DO NOT EDIT.
// source: startrek.proto

package startrek

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	nullint "github.com/roockie-nebula/protobuf/types/nullint"
	nullstring "github.com/roockie-nebula/protobuf/types/nullstring"
	timestamp "github.com/roockie-nebula/protobuf/types/timestamp"
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

type StarfleetShip struct {
	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NoOfPassengers   *nullint.NullInt       `protobuf:"bytes,2,opt,name=no_of_passengers,json=noOfPassengers,proto3" json:"no_of_passengers,omitempty"`
	MissionStatement *nullstring.NullString `protobuf:"bytes,3,opt,name=mission_statement,json=missionStatement,proto3" json:"mission_statement,omitempty"`
	// use a different db column name for the departure time
	// `db:"we_are_leaving_at"`
	DepartureTime        *timestamp.Timestamp `protobuf:"bytes,4,opt,name=departure_time,json=departureTime,proto3" json:"departure_time,omitempty" db:"we_are_leaving_at"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StarfleetShip) Reset()         { *m = StarfleetShip{} }
func (m *StarfleetShip) String() string { return proto.CompactTextString(m) }
func (*StarfleetShip) ProtoMessage()    {}
func (*StarfleetShip) Descriptor() ([]byte, []int) {
	return fileDescriptor_f50c8cdd5334cdc8, []int{0}
}

func (m *StarfleetShip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StarfleetShip.Unmarshal(m, b)
}
func (m *StarfleetShip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StarfleetShip.Marshal(b, m, deterministic)
}
func (m *StarfleetShip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StarfleetShip.Merge(m, src)
}
func (m *StarfleetShip) XXX_Size() int {
	return xxx_messageInfo_StarfleetShip.Size(m)
}
func (m *StarfleetShip) XXX_DiscardUnknown() {
	xxx_messageInfo_StarfleetShip.DiscardUnknown(m)
}

var xxx_messageInfo_StarfleetShip proto.InternalMessageInfo

func (m *StarfleetShip) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StarfleetShip) GetNoOfPassengers() *nullint.NullInt {
	if m != nil {
		return m.NoOfPassengers
	}
	return nil
}

func (m *StarfleetShip) GetMissionStatement() *nullstring.NullString {
	if m != nil {
		return m.MissionStatement
	}
	return nil
}

func (m *StarfleetShip) GetDepartureTime() *timestamp.Timestamp {
	if m != nil {
		return m.DepartureTime
	}
	return nil
}

func init() {
	proto.RegisterType((*StarfleetShip)(nil), "startrek.StarfleetShip")
}

func init() { proto.RegisterFile("startrek.proto", fileDescriptor_f50c8cdd5334cdc8) }

var fileDescriptor_f50c8cdd5334cdc8 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0xa8, 0x10, 0x18, 0xb5, 0x2a, 0x9e, 0x42, 0x61, 0xa8, 0x10, 0x82, 0x4e, 0x89,
	0x04, 0xff, 0x00, 0xa6, 0x2e, 0x80, 0x12, 0xf6, 0xc8, 0x85, 0x4b, 0xb1, 0x6a, 0x9f, 0xad, 0xbb,
	0xcb, 0xc0, 0x2f, 0x67, 0x45, 0x75, 0x93, 0x2e, 0x7c, 0x6c, 0xfe, 0x78, 0x9e, 0xf7, 0x74, 0xaf,
	0x9a, 0xb0, 0x18, 0x12, 0x82, 0x4d, 0x11, 0x29, 0x48, 0xd0, 0xc7, 0xc3, 0x7d, 0x76, 0x43, 0x21,
	0xbc, 0x6d, 0x2c, 0x94, 0xe9, 0x63, 0xd5, 0xb5, 0xa5, 0x7c, 0x46, 0xe0, 0x52, 0xac, 0x07, 0x16,
	0xe3, 0xe3, 0xce, 0x98, 0xdd, 0xfe, 0xc1, 0x61, 0xe7, 0x1c, 0x0b, 0x59, 0x5c, 0xf7, 0xe0, 0xf5,
	0x3f, 0xa0, 0x45, 0xd9, 0x51, 0x57, 0x5f, 0x99, 0x1a, 0xd7, 0x62, 0xa8, 0x75, 0x00, 0x52, 0x7f,
	0xd8, 0xa8, 0xb5, 0x1a, 0xa1, 0xf1, 0x90, 0x67, 0xf3, 0x6c, 0x71, 0x52, 0xa5, 0xb3, 0x7e, 0x54,
	0x53, 0x0c, 0x4d, 0x68, 0x9b, 0x68, 0x98, 0x01, 0xd7, 0x40, 0x9c, 0x1f, 0xcc, 0xb3, 0xc5, 0xe9,
	0xdd, 0x79, 0xd1, 0x8f, 0x29, 0x86, 0x31, 0xc5, 0x53, 0xe7, 0xdc, 0x12, 0xa5, 0x9a, 0x60, 0x78,
	0x6e, 0x5f, 0xf6, 0x82, 0x5e, 0xaa, 0x33, 0x6f, 0x99, 0x6d, 0xc0, 0x86, 0xc5, 0x08, 0x78, 0x40,
	0xc9, 0x0f, 0x53, 0xca, 0xe5, 0xef, 0x29, 0x75, 0xda, 0xa7, 0x9a, 0xf6, 0x5a, 0x3d, 0x58, 0xfa,
	0x41, 0x4d, 0xde, 0x21, 0x1a, 0x92, 0x8e, 0xa0, 0xd9, 0x36, 0x94, 0x8f, 0x52, 0xce, 0xc5, 0xcf,
	0x9c, 0xd7, 0xa1, 0xbf, 0x6a, 0xbc, 0x57, 0xb6, 0x6f, 0xab, 0xa3, 0x84, 0xdc, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x31, 0x6f, 0xe4, 0x9e, 0x93, 0x01, 0x00, 0x00,
}
