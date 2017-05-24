// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layout.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		layout.proto
		view.proto

	It has these top-level messages:
		Point
		Rect
		Insets
		Guide
		Node
		Root
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptorLayout, []int{0} }

func (m *Point) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Rect struct {
	Min *Point `protobuf:"bytes,1,opt,name=min" json:"min,omitempty"`
	Max *Point `protobuf:"bytes,2,opt,name=max" json:"max,omitempty"`
}

func (m *Rect) Reset()                    { *m = Rect{} }
func (m *Rect) String() string            { return proto.CompactTextString(m) }
func (*Rect) ProtoMessage()               {}
func (*Rect) Descriptor() ([]byte, []int) { return fileDescriptorLayout, []int{1} }

func (m *Rect) GetMin() *Point {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *Rect) GetMax() *Point {
	if m != nil {
		return m.Max
	}
	return nil
}

type Insets struct {
	Top    float64 `protobuf:"fixed64,1,opt,name=top,proto3" json:"top,omitempty"`
	Left   float64 `protobuf:"fixed64,2,opt,name=left,proto3" json:"left,omitempty"`
	Bottom float64 `protobuf:"fixed64,3,opt,name=bottom,proto3" json:"bottom,omitempty"`
	Right  float64 `protobuf:"fixed64,4,opt,name=right,proto3" json:"right,omitempty"`
}

func (m *Insets) Reset()                    { *m = Insets{} }
func (m *Insets) String() string            { return proto.CompactTextString(m) }
func (*Insets) ProtoMessage()               {}
func (*Insets) Descriptor() ([]byte, []int) { return fileDescriptorLayout, []int{2} }

func (m *Insets) GetTop() float64 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *Insets) GetLeft() float64 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *Insets) GetBottom() float64 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *Insets) GetRight() float64 {
	if m != nil {
		return m.Right
	}
	return 0
}

type Guide struct {
	Frame  *Rect   `protobuf:"bytes,1,opt,name=frame" json:"frame,omitempty"`
	Insets *Insets `protobuf:"bytes,2,opt,name=insets" json:"insets,omitempty"`
	ZIndex int64   `protobuf:"varint,3,opt,name=zIndex,proto3" json:"zIndex,omitempty"`
}

func (m *Guide) Reset()                    { *m = Guide{} }
func (m *Guide) String() string            { return proto.CompactTextString(m) }
func (*Guide) ProtoMessage()               {}
func (*Guide) Descriptor() ([]byte, []int) { return fileDescriptorLayout, []int{3} }

func (m *Guide) GetFrame() *Rect {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *Guide) GetInsets() *Insets {
	if m != nil {
		return m.Insets
	}
	return nil
}

func (m *Guide) GetZIndex() int64 {
	if m != nil {
		return m.ZIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "layout.Point")
	proto.RegisterType((*Rect)(nil), "layout.Rect")
	proto.RegisterType((*Insets)(nil), "layout.Insets")
	proto.RegisterType((*Guide)(nil), "layout.Guide")
}
func (m *Point) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Point) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.X != 0 {
		dAtA[i] = 0x9
		i++
		i = encodeFixed64Layout(dAtA, i, uint64(math.Float64bits(float64(m.X))))
	}
	if m.Y != 0 {
		dAtA[i] = 0x11
		i++
		i = encodeFixed64Layout(dAtA, i, uint64(math.Float64bits(float64(m.Y))))
	}
	return i, nil
}

func (m *Rect) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rect) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Min != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLayout(dAtA, i, uint64(m.Min.Size()))
		n1, err := m.Min.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Max != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLayout(dAtA, i, uint64(m.Max.Size()))
		n2, err := m.Max.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Insets) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Insets) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Top != 0 {
		dAtA[i] = 0x9
		i++
		i = encodeFixed64Layout(dAtA, i, uint64(math.Float64bits(float64(m.Top))))
	}
	if m.Left != 0 {
		dAtA[i] = 0x11
		i++
		i = encodeFixed64Layout(dAtA, i, uint64(math.Float64bits(float64(m.Left))))
	}
	if m.Bottom != 0 {
		dAtA[i] = 0x19
		i++
		i = encodeFixed64Layout(dAtA, i, uint64(math.Float64bits(float64(m.Bottom))))
	}
	if m.Right != 0 {
		dAtA[i] = 0x21
		i++
		i = encodeFixed64Layout(dAtA, i, uint64(math.Float64bits(float64(m.Right))))
	}
	return i, nil
}

func (m *Guide) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Guide) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Frame != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLayout(dAtA, i, uint64(m.Frame.Size()))
		n3, err := m.Frame.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Insets != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLayout(dAtA, i, uint64(m.Insets.Size()))
		n4, err := m.Insets.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.ZIndex != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLayout(dAtA, i, uint64(m.ZIndex))
	}
	return i, nil
}

func encodeFixed64Layout(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Layout(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLayout(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Point) Size() (n int) {
	var l int
	_ = l
	if m.X != 0 {
		n += 9
	}
	if m.Y != 0 {
		n += 9
	}
	return n
}

func (m *Rect) Size() (n int) {
	var l int
	_ = l
	if m.Min != nil {
		l = m.Min.Size()
		n += 1 + l + sovLayout(uint64(l))
	}
	if m.Max != nil {
		l = m.Max.Size()
		n += 1 + l + sovLayout(uint64(l))
	}
	return n
}

func (m *Insets) Size() (n int) {
	var l int
	_ = l
	if m.Top != 0 {
		n += 9
	}
	if m.Left != 0 {
		n += 9
	}
	if m.Bottom != 0 {
		n += 9
	}
	if m.Right != 0 {
		n += 9
	}
	return n
}

func (m *Guide) Size() (n int) {
	var l int
	_ = l
	if m.Frame != nil {
		l = m.Frame.Size()
		n += 1 + l + sovLayout(uint64(l))
	}
	if m.Insets != nil {
		l = m.Insets.Size()
		n += 1 + l + sovLayout(uint64(l))
	}
	if m.ZIndex != 0 {
		n += 1 + sovLayout(uint64(m.ZIndex))
	}
	return n
}

func sovLayout(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLayout(x uint64) (n int) {
	return sovLayout(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Point) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLayout
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Point: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Point: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.X = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Y = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipLayout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLayout
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Rect) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLayout
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Rect: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rect: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLayout
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Min == nil {
				m.Min = &Point{}
			}
			if err := m.Min.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLayout
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Max == nil {
				m.Max = &Point{}
			}
			if err := m.Max.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLayout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLayout
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Insets) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLayout
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Insets: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Insets: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Top", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Top = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Left = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bottom", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Bottom = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Right = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipLayout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLayout
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Guide) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLayout
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Guide: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Guide: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frame", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLayout
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Frame == nil {
				m.Frame = &Rect{}
			}
			if err := m.Frame.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Insets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLayout
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Insets == nil {
				m.Insets = &Insets{}
			}
			if err := m.Insets.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZIndex", wireType)
			}
			m.ZIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ZIndex |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLayout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLayout
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLayout(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLayout
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLayout
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLayout
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLayout
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLayout
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLayout(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLayout = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLayout   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("layout.proto", fileDescriptorLayout) }

var fileDescriptorLayout = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x71, 0xf3, 0x67, 0x38, 0x02, 0x42, 0xa7, 0x0a, 0x65, 0x0a, 0xc8, 0x48, 0x88, 0xa9,
	0x03, 0xbc, 0x01, 0x0b, 0x74, 0x43, 0x1e, 0x11, 0x4b, 0x42, 0x5d, 0xb0, 0x68, 0xe2, 0x28, 0xbd,
	0x4a, 0x0e, 0x4f, 0xc2, 0x23, 0x31, 0xf2, 0x08, 0x28, 0xbc, 0x08, 0xf2, 0xd9, 0xdd, 0xba, 0xdd,
	0x77, 0xf7, 0xbb, 0xdc, 0x17, 0x43, 0xb1, 0xa9, 0x47, 0xbb, 0xa3, 0x45, 0x3f, 0x58, 0xb2, 0x98,
	0x07, 0x92, 0x57, 0x90, 0x3d, 0x59, 0xd3, 0x11, 0x16, 0x20, 0x5c, 0x29, 0x2e, 0xc5, 0x8d, 0x50,
	0xc2, 0x79, 0x1a, 0xcb, 0x59, 0xa0, 0x51, 0x3e, 0x42, 0xaa, 0xf4, 0x2b, 0xe1, 0x05, 0x24, 0xad,
	0xe9, 0x38, 0x75, 0x7c, 0x7b, 0xb2, 0x88, 0x1f, 0xe4, 0x7d, 0xe5, 0x27, 0x1c, 0xa8, 0x1d, 0x2f,
	0x1e, 0x08, 0xd4, 0x4e, 0xbe, 0x40, 0xbe, 0xec, 0xb6, 0x9a, 0xb6, 0x78, 0x06, 0x09, 0xd9, 0x3e,
	0x5e, 0xf4, 0x25, 0x22, 0xa4, 0x1b, 0xbd, 0xa6, 0x78, 0x96, 0x6b, 0x3c, 0x87, 0xbc, 0xb1, 0x44,
	0xb6, 0x2d, 0x13, 0xee, 0x46, 0xc2, 0x39, 0x64, 0x83, 0x79, 0x7b, 0xa7, 0x32, 0xe5, 0x76, 0x00,
	0xf9, 0x01, 0xd9, 0xc3, 0xce, 0xac, 0x34, 0x4a, 0xc8, 0xd6, 0x43, 0xdd, 0xea, 0xa8, 0x5a, 0xec,
	0x4d, 0xfc, 0x5f, 0xa8, 0x30, 0xc2, 0x6b, 0xc8, 0x0d, 0xab, 0x44, 0xdd, 0xd3, 0x7d, 0x28, 0x08,
	0xaa, 0x38, 0xf5, 0x0a, 0x9f, 0xcb, 0x6e, 0xa5, 0x1d, 0x2b, 0x24, 0x2a, 0xd2, 0xfd, 0xfc, 0x7b,
	0xaa, 0xc4, 0xcf, 0x54, 0x89, 0xdf, 0xa9, 0x12, 0x5f, 0x7f, 0xd5, 0xd1, 0xf3, 0xac, 0x6f, 0x9a,
	0x9c, 0x9f, 0xf7, 0xee, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xea, 0x77, 0x99, 0xd8, 0x6e, 0x01, 0x00,
	0x00,
}
