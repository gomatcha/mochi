// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/matcha/pb/view/slider/slider.proto

/*
Package slider is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/matcha/pb/view/slider/slider.proto

It has these top-level messages:
	View
	Event
*/
package slider

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
	Value    float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	MaxValue float64 `protobuf:"fixed64,2,opt,name=maxValue" json:"maxValue,omitempty"`
	MinValue float64 `protobuf:"fixed64,3,opt,name=minValue" json:"minValue,omitempty"`
	Enabled  bool    `protobuf:"varint,4,opt,name=enabled" json:"enabled,omitempty"`
}

func (m *View) Reset()                    { *m = View{} }
func (m *View) String() string            { return proto.CompactTextString(m) }
func (*View) ProtoMessage()               {}
func (*View) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *View) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *View) GetMaxValue() float64 {
	if m != nil {
		return m.MaxValue
	}
	return 0
}

func (m *View) GetMinValue() float64 {
	if m != nil {
		return m.MinValue
	}
	return 0
}

func (m *View) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type Event struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*View)(nil), "slider.View")
	proto.RegisterType((*Event)(nil), "slider.Event")
}

func init() {
	proto.RegisterFile("github.com/gomatcha/matcha/pb/view/slider/slider.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2f, 0x4b, 0x2d, 0x4a, 0xae, 0xcc, 0xd3, 0xcf,
	0x4d, 0x2c, 0x49, 0xce, 0x48, 0xd4, 0x2f, 0x48, 0xd2, 0x2f, 0xcb, 0x4c, 0x2d, 0xd7, 0x2f, 0xce,
	0xc9, 0x4c, 0x49, 0x2d, 0x82, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e,
	0x52, 0x1e, 0x17, 0x4b, 0x58, 0x66, 0x6a, 0xb9, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69,
	0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x84, 0x23, 0x24, 0xc5, 0xc5, 0x91, 0x9b, 0x58,
	0x11, 0x06, 0x96, 0x60, 0x02, 0x4b, 0xc0, 0xf9, 0x60, 0xb9, 0xcc, 0x3c, 0x88, 0x1c, 0x33, 0x54,
	0x0e, 0xca, 0x17, 0x92, 0xe0, 0x62, 0x4f, 0xcd, 0x4b, 0x4c, 0xca, 0x49, 0x4d, 0x91, 0x60, 0x51,
	0x60, 0xd4, 0xe0, 0x08, 0x82, 0x71, 0x95, 0x64, 0xb9, 0x58, 0x5d, 0xcb, 0x52, 0xf3, 0x4a, 0xb0,
	0x5b, 0xe8, 0x24, 0x19, 0x05, 0x75, 0xd8, 0x22, 0x26, 0x3e, 0x5f, 0xb0, 0x2f, 0x82, 0xc1, 0xdc,
	0x80, 0xa4, 0x24, 0x36, 0xb0, 0xc3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x86, 0xb9,
	0xa4, 0xf1, 0x00, 0x00, 0x00,
}
