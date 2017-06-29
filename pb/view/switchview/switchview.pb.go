// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/matcha/pb/view/switchview/switchview.proto

/*
Package switchview is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/matcha/pb/view/switchview/switchview.proto

It has these top-level messages:
	View
	Event
*/
package switchview

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

type View struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *View) Reset()                    { *m = View{} }
func (m *View) String() string            { return proto.CompactTextString(m) }
func (*View) ProtoMessage()               {}
func (*View) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *View) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type Event struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*View)(nil), "switchview.View")
	proto.RegisterType((*Event)(nil), "switchview.Event")
}

func init() {
	proto.RegisterFile("github.com/gomatcha/matcha/pb/view/switchview/switchview.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2f, 0x4b, 0x2d, 0x4a, 0xae, 0xcc, 0xd3, 0xcf,
	0x4d, 0x2c, 0x49, 0xce, 0x48, 0xd4, 0x2f, 0x48, 0xd2, 0x2f, 0xcb, 0x4c, 0x2d, 0xd7, 0x2f, 0x2e,
	0xcf, 0x2c, 0x49, 0xce, 0x40, 0x63, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x21, 0x44,
	0x94, 0x64, 0xb8, 0x58, 0xc2, 0x32, 0x53, 0xcb, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0x20, 0x1c, 0x25, 0x59, 0x2e, 0x56, 0xd7, 0xb2,
	0xd4, 0xbc, 0x12, 0xec, 0xd2, 0x4e, 0x8a, 0x51, 0x48, 0x46, 0x2d, 0x62, 0x12, 0xf2, 0x05, 0xbb,
	0x21, 0xc0, 0x29, 0x18, 0x2c, 0x08, 0x32, 0x37, 0x89, 0x0d, 0x6c, 0xa5, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x8c, 0x92, 0x38, 0xc5, 0xb3, 0x00, 0x00, 0x00,
}
