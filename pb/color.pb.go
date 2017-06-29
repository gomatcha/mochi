// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/matcha/pb/color.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/matcha/pb/color.proto
	github.com/overcyn/matcha/pb/image.proto

It has these top-level messages:
	Color
	Image
	ImageProperties
	ImageOrResource
*/
package pb

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

type Color struct {
	Red   uint32 `protobuf:"varint,1,opt,name=red" json:"red,omitempty"`
	Blue  uint32 `protobuf:"varint,2,opt,name=blue" json:"blue,omitempty"`
	Green uint32 `protobuf:"varint,3,opt,name=green" json:"green,omitempty"`
	Alpha uint32 `protobuf:"varint,4,opt,name=alpha" json:"alpha,omitempty"`
}

func (m *Color) Reset()                    { *m = Color{} }
func (m *Color) String() string            { return proto.CompactTextString(m) }
func (*Color) ProtoMessage()               {}
func (*Color) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Color) GetRed() uint32 {
	if m != nil {
		return m.Red
	}
	return 0
}

func (m *Color) GetBlue() uint32 {
	if m != nil {
		return m.Blue
	}
	return 0
}

func (m *Color) GetGreen() uint32 {
	if m != nil {
		return m.Green
	}
	return 0
}

func (m *Color) GetAlpha() uint32 {
	if m != nil {
		return m.Alpha
	}
	return 0
}

func init() {
	proto.RegisterType((*Color)(nil), "matcha.Color")
}

func init() { proto.RegisterFile("github.com/gomatcha/matcha/pb/color.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2f, 0x4b, 0x2d, 0x4a, 0xae, 0xcc, 0xd3, 0xcf,
	0x4d, 0x2c, 0x49, 0xce, 0x48, 0xd4, 0x2f, 0x48, 0xd2, 0x4f, 0xce, 0xcf, 0xc9, 0x2f, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x08, 0x2b, 0x45, 0x72, 0xb1, 0x3a, 0x83, 0x84, 0x85,
	0x04, 0xb8, 0x98, 0x8b, 0x52, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x40, 0x4c, 0x21,
	0x21, 0x2e, 0x96, 0xa4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0xb0, 0x10, 0x98, 0x2d, 0x24, 0xc2, 0xc5,
	0x9a, 0x5e, 0x94, 0x9a, 0x9a, 0x27, 0xc1, 0x0c, 0x16, 0x84, 0x70, 0x40, 0xa2, 0x89, 0x39, 0x05,
	0x19, 0x89, 0x12, 0x2c, 0x10, 0x51, 0x30, 0xc7, 0x89, 0x3f, 0x8a, 0xa9, 0x20, 0x69, 0x11, 0x13,
	0x87, 0x2f, 0xd8, 0xa6, 0x00, 0xa7, 0x24, 0x36, 0xb0, 0xd5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcd, 0x38, 0x98, 0x3b, 0xa6, 0x00, 0x00, 0x00,
}
