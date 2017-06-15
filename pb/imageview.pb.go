// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/mochi/pb/imageview.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import env "github.com/overcyn/mochi/pb/env"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResizeMode int32

const (
	ResizeMode_FIT     ResizeMode = 0
	ResizeMode_FILL    ResizeMode = 1
	ResizeMode_STRETCH ResizeMode = 2
	ResizeMode_CENTER  ResizeMode = 3
)

var ResizeMode_name = map[int32]string{
	0: "FIT",
	1: "FILL",
	2: "STRETCH",
	3: "CENTER",
}
var ResizeMode_value = map[string]int32{
	"FIT":     0,
	"FILL":    1,
	"STRETCH": 2,
	"CENTER":  3,
}

func (x ResizeMode) String() string {
	return proto.EnumName(ResizeMode_name, int32(x))
}
func (ResizeMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Image struct {
	Width  int64  `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height int64  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Stride int64  `protobuf:"varint,4,opt,name=stride" json:"stride,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Image) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetStride() int64 {
	if m != nil {
		return m.Stride
	}
	return 0
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ImageView struct {
	Image      *Image        `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Resource   *env.Resource `protobuf:"bytes,4,opt,name=resource" json:"resource,omitempty"`
	ResizeMode ResizeMode    `protobuf:"varint,2,opt,name=resizeMode,enum=imageview.ResizeMode" json:"resizeMode,omitempty"`
	Tint       *Color        `protobuf:"bytes,3,opt,name=tint" json:"tint,omitempty"`
}

func (m *ImageView) Reset()                    { *m = ImageView{} }
func (m *ImageView) String() string            { return proto.CompactTextString(m) }
func (*ImageView) ProtoMessage()               {}
func (*ImageView) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ImageView) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImageView) GetResource() *env.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ImageView) GetResizeMode() ResizeMode {
	if m != nil {
		return m.ResizeMode
	}
	return ResizeMode_FIT
}

func (m *ImageView) GetTint() *Color {
	if m != nil {
		return m.Tint
	}
	return nil
}

func init() {
	proto.RegisterType((*Image)(nil), "imageview.Image")
	proto.RegisterType((*ImageView)(nil), "imageview.ImageView")
	proto.RegisterEnum("imageview.ResizeMode", ResizeMode_name, ResizeMode_value)
}

func init() { proto.RegisterFile("github.com/overcyn/mochi/pb/imageview.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x34, 0x1f, 0xfd, 0x7a, 0xa9, 0x25, 0x2c, 0x2a, 0xa1, 0xa7, 0xd2, 0x83, 0x56, 0x85, 0x0d,
	0x44, 0x04, 0xcf, 0x2d, 0x2d, 0x16, 0x5a, 0x91, 0x35, 0x78, 0xf0, 0x96, 0x8f, 0x47, 0xb3, 0x60,
	0xb3, 0x65, 0xbb, 0x4d, 0xd1, 0x9f, 0xe3, 0x9f, 0xf0, 0xef, 0x49, 0x36, 0xb5, 0xed, 0xa9, 0xb7,
	0x7d, 0x33, 0xf3, 0x76, 0xe6, 0x0d, 0xdc, 0x2f, 0xb8, 0xca, 0x36, 0x31, 0x4d, 0xc4, 0xd2, 0x17,
	0x05, 0xca, 0xe4, 0x2b, 0xf7, 0x97, 0x22, 0xc9, 0xb8, 0xbf, 0x8a, 0x7d, 0xbe, 0x8c, 0x16, 0x58,
	0x70, 0xdc, 0xd2, 0x95, 0x14, 0x4a, 0x90, 0xd6, 0x1e, 0xe8, 0xde, 0x9c, 0xda, 0x4b, 0xc4, 0xa7,
	0x90, 0xd5, 0x4e, 0x97, 0x9e, 0x12, 0x62, 0x5e, 0xf8, 0x12, 0xd7, 0x62, 0x23, 0x13, 0xac, 0xf4,
	0xfd, 0x08, 0x6a, 0xd3, 0xd2, 0x85, 0x5c, 0x40, 0x6d, 0xcb, 0x53, 0x95, 0x79, 0x46, 0xcf, 0x18,
	0x58, 0xac, 0x1a, 0xc8, 0x15, 0xd4, 0x33, 0xe4, 0x8b, 0x4c, 0x79, 0xa6, 0x86, 0x77, 0x53, 0x89,
	0xaf, 0x95, 0xe4, 0x29, 0x7a, 0x76, 0x85, 0x57, 0x13, 0x21, 0x60, 0xa7, 0x91, 0x8a, 0x3c, 0xab,
	0x67, 0x0c, 0xda, 0x4c, 0xbf, 0xfb, 0xbf, 0x06, 0xb4, 0xb4, 0xc7, 0x3b, 0xc7, 0x2d, 0xb9, 0x86,
	0x9a, 0x3e, 0x4b, 0xfb, 0x38, 0x81, 0x4b, 0x0f, 0x57, 0x6b, 0x11, 0xab, 0x68, 0x72, 0x0b, 0xcd,
	0xff, 0xa8, 0xda, 0xc3, 0x09, 0xce, 0x29, 0xe6, 0x05, 0x65, 0x3b, 0x90, 0xed, 0x69, 0xf2, 0x08,
	0x20, 0x71, 0xcd, 0xbf, 0x71, 0x2e, 0x52, 0xd4, 0x41, 0x3b, 0xc1, 0xe5, 0xd1, 0xbf, 0x6c, 0x4f,
	0xb2, 0x23, 0x21, 0xe9, 0x81, 0xad, 0x78, 0xae, 0x74, 0x56, 0x27, 0x68, 0x53, 0x5d, 0x13, 0x1d,
	0x95, 0x65, 0x32, 0xcd, 0xdc, 0x3d, 0x01, 0x1c, 0x76, 0x49, 0x03, 0xac, 0xc9, 0x34, 0x74, 0xcf,
	0x48, 0x13, 0xec, 0xc9, 0x74, 0x36, 0x73, 0x0d, 0xe2, 0x40, 0xe3, 0x2d, 0x64, 0xe3, 0x70, 0xf4,
	0xec, 0x9a, 0x04, 0xa0, 0x3e, 0x1a, 0xbf, 0x84, 0x63, 0xe6, 0x5a, 0xc3, 0xce, 0x87, 0xb9, 0x8a,
	0x7f, 0xcc, 0xc6, 0xbc, 0xfc, 0xf3, 0x75, 0x18, 0xd7, 0x75, 0xdb, 0x0f, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xbc, 0x49, 0x1b, 0x66, 0x00, 0x02, 0x00, 0x00,
}
