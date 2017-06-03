// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/mochi/pb/view/tabnavigator/tabnavigator.proto

/*
Package tabnavigator is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/mochi/pb/view/tabnavigator/tabnavigator.proto

It has these top-level messages:
	Tab
	TabNavigator
*/
package tabnavigator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import imageview "github.com/overcyn/mochi/pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Tab struct {
	Id           int64            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title        string           `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Icon         *imageview.Image `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	SelectedIcon *imageview.Image `protobuf:"bytes,4,opt,name=selectedIcon" json:"selectedIcon,omitempty"`
	Badge        string           `protobuf:"bytes,5,opt,name=badge" json:"badge,omitempty"`
}

func (m *Tab) Reset()                    { *m = Tab{} }
func (m *Tab) String() string            { return proto.CompactTextString(m) }
func (*Tab) ProtoMessage()               {}
func (*Tab) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Tab) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tab) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Tab) GetIcon() *imageview.Image {
	if m != nil {
		return m.Icon
	}
	return nil
}

func (m *Tab) GetSelectedIcon() *imageview.Image {
	if m != nil {
		return m.SelectedIcon
	}
	return nil
}

func (m *Tab) GetBadge() string {
	if m != nil {
		return m.Badge
	}
	return ""
}

type TabNavigator struct {
	Tabs []*Tab `protobuf:"bytes,1,rep,name=tabs" json:"tabs,omitempty"`
}

func (m *TabNavigator) Reset()                    { *m = TabNavigator{} }
func (m *TabNavigator) String() string            { return proto.CompactTextString(m) }
func (*TabNavigator) ProtoMessage()               {}
func (*TabNavigator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TabNavigator) GetTabs() []*Tab {
	if m != nil {
		return m.Tabs
	}
	return nil
}

func init() {
	proto.RegisterType((*Tab)(nil), "tabnavigator.Tab")
	proto.RegisterType((*TabNavigator)(nil), "tabnavigator.TabNavigator")
}

func init() {
	proto.RegisterFile("github.com/overcyn/mochi/pb/view/tabnavigator/tabnavigator.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x49, 0xdb, 0x09, 0x66, 0x45, 0x34, 0x7a, 0x08, 0x9e, 0xca, 0x98, 0x50, 0x10, 0x1a,
	0x98, 0x7a, 0x97, 0xdd, 0x76, 0x50, 0xa4, 0xf4, 0xe4, 0xed, 0xbd, 0x34, 0x74, 0x0f, 0xd6, 0x66,
	0x74, 0xb1, 0xe2, 0x87, 0xf1, 0xe2, 0x27, 0x95, 0xb4, 0x88, 0xcb, 0x41, 0x6f, 0xf9, 0xe7, 0xfd,
	0xf3, 0x7b, 0xe4, 0xc7, 0x1f, 0x1b, 0x72, 0xdb, 0x37, 0x2c, 0xb4, 0x6d, 0x95, 0x1d, 0x4c, 0xaf,
	0x3f, 0x3a, 0xd5, 0x5a, 0xbd, 0x25, 0xb5, 0x47, 0x35, 0x90, 0x79, 0x57, 0x0e, 0xb0, 0x83, 0x81,
	0x1a, 0x70, 0xb6, 0x0f, 0x42, 0xb1, 0xef, 0xad, 0xb3, 0x22, 0x3d, 0xbe, 0xbb, 0xbe, 0xfd, 0x8f,
	0x47, 0x2d, 0x34, 0xc6, 0x43, 0xa7, 0xa7, 0x8b, 0x4f, 0xc6, 0xe3, 0x0a, 0x50, 0x9c, 0xf1, 0x88,
	0x6a, 0xc9, 0x32, 0x96, 0xc7, 0x65, 0x44, 0xb5, 0xb8, 0xe2, 0x33, 0x47, 0x6e, 0x67, 0x64, 0x94,
	0xb1, 0xfc, 0xb4, 0x9c, 0x82, 0x58, 0xf2, 0x84, 0xb4, 0xed, 0x64, 0x9c, 0xb1, 0x7c, 0xbe, 0x3a,
	0x2f, 0x7e, 0x69, 0x1b, 0x7f, 0x2a, 0xc7, 0xa9, 0xb8, 0xe7, 0xe9, 0xc1, 0xec, 0x8c, 0x76, 0xa6,
	0xde, 0xf8, 0x76, 0xf2, 0x47, 0x3b, 0x68, 0xf9, 0x8d, 0x08, 0x75, 0x63, 0xe4, 0x6c, 0xda, 0x38,
	0x86, 0xc5, 0x03, 0x4f, 0x2b, 0xc0, 0xe7, 0x9f, 0xcf, 0x89, 0x1b, 0x9e, 0x38, 0xc0, 0x83, 0x64,
	0x59, 0x9c, 0xcf, 0x57, 0x17, 0x45, 0x60, 0xa3, 0x02, 0x2c, 0xc7, 0xf1, 0x7a, 0xf9, 0x1a, 0x38,
	0xf9, 0x8a, 0x2e, 0x9f, 0xbc, 0x81, 0x97, 0xf5, 0x31, 0x0c, 0x4f, 0x46, 0x07, 0x77, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0x07, 0x9f, 0xa2, 0x82, 0x01, 0x00, 0x00,
}