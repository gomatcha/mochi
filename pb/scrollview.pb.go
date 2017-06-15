// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/mochi/pb/scrollview.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScrollView struct {
	ScrollEnabled                  bool `protobuf:"varint,1,opt,name=scrollEnabled" json:"scrollEnabled,omitempty"`
	ShowsHorizontalScrollIndicator bool `protobuf:"varint,2,opt,name=showsHorizontalScrollIndicator" json:"showsHorizontalScrollIndicator,omitempty"`
	ShowsVerticalScrollIndicator   bool `protobuf:"varint,3,opt,name=showsVerticalScrollIndicator" json:"showsVerticalScrollIndicator,omitempty"`
}

func (m *ScrollView) Reset()                    { *m = ScrollView{} }
func (m *ScrollView) String() string            { return proto.CompactTextString(m) }
func (*ScrollView) ProtoMessage()               {}
func (*ScrollView) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ScrollView) GetScrollEnabled() bool {
	if m != nil {
		return m.ScrollEnabled
	}
	return false
}

func (m *ScrollView) GetShowsHorizontalScrollIndicator() bool {
	if m != nil {
		return m.ShowsHorizontalScrollIndicator
	}
	return false
}

func (m *ScrollView) GetShowsVerticalScrollIndicator() bool {
	if m != nil {
		return m.ShowsVerticalScrollIndicator
	}
	return false
}

func init() {
	proto.RegisterType((*ScrollView)(nil), "paint.ScrollView")
}

func init() { proto.RegisterFile("github.com/overcyn/mochi/pb/scrollview.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2f, 0x4b, 0x2d, 0x4a, 0xae, 0xcc, 0xd3, 0xcf,
	0xcd, 0x4f, 0xce, 0xc8, 0xd4, 0x2f, 0x48, 0xd2, 0x2f, 0x4e, 0x2e, 0xca, 0xcf, 0xc9, 0x29, 0xcb,
	0x4c, 0x2d, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x48, 0xcc, 0xcc, 0x2b, 0x51,
	0xda, 0xc7, 0xc8, 0xc5, 0x15, 0x0c, 0x96, 0x0b, 0xcb, 0x4c, 0x2d, 0x17, 0x52, 0xe1, 0xe2, 0x85,
	0xa8, 0x74, 0xcd, 0x4b, 0x4c, 0xca, 0x49, 0x4d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x42,
	0x15, 0x14, 0x72, 0xe3, 0x92, 0x2b, 0xce, 0xc8, 0x2f, 0x2f, 0xf6, 0xc8, 0x2f, 0xca, 0xac, 0xca,
	0xcf, 0x2b, 0x49, 0xcc, 0x81, 0x18, 0xe1, 0x99, 0x97, 0x92, 0x99, 0x9c, 0x58, 0x92, 0x5f, 0x24,
	0xc1, 0x04, 0xd6, 0x46, 0x40, 0x95, 0x90, 0x13, 0x97, 0x0c, 0x58, 0x45, 0x58, 0x6a, 0x51, 0x49,
	0x66, 0x32, 0xa6, 0x29, 0xcc, 0x60, 0x53, 0xf0, 0xaa, 0x71, 0xe2, 0x8b, 0x62, 0x2a, 0x48, 0x5a,
	0xc4, 0xc4, 0xee, 0x0b, 0xf2, 0x69, 0x80, 0x53, 0x12, 0x1b, 0xd8, 0x7b, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbe, 0x60, 0xa9, 0xaa, 0x0e, 0x01, 0x00, 0x00,
}
