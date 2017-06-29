// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/matcha/pb/touch/touch2.proto

/*
Package touch is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/matcha/pb/touch/touch2.proto

It has these top-level messages:
	Recognizer
	RecognizerList
	ButtonRecognizer
	ButtonEvent
	TapRecognizer
	TapEvent
	PressRecognizer
	PressEvent
*/
package touch

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"
import matcha_layout "github.com/gomatcha/matcha/pb/layout"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventKind int32

const (
	EventKind_EVENT_KIND_POSSIBLE   EventKind = 0
	EventKind_EVENT_KIND_CHANGED    EventKind = 1
	EventKind_EVENT_KIND_FAILED     EventKind = 2
	EventKind_EVENT_KIND_RECOGNIZED EventKind = 3
)

var EventKind_name = map[int32]string{
	0: "EVENT_KIND_POSSIBLE",
	1: "EVENT_KIND_CHANGED",
	2: "EVENT_KIND_FAILED",
	3: "EVENT_KIND_RECOGNIZED",
}
var EventKind_value = map[string]int32{
	"EVENT_KIND_POSSIBLE":   0,
	"EVENT_KIND_CHANGED":    1,
	"EVENT_KIND_FAILED":     2,
	"EVENT_KIND_RECOGNIZED": 3,
}

func (x EventKind) String() string {
	return proto.EnumName(EventKind_name, int32(x))
}
func (EventKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Recognizer struct {
	Id         int64                `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Recognizer *google_protobuf.Any `protobuf:"bytes,3,opt,name=recognizer" json:"recognizer,omitempty"`
}

func (m *Recognizer) Reset()                    { *m = Recognizer{} }
func (m *Recognizer) String() string            { return proto.CompactTextString(m) }
func (*Recognizer) ProtoMessage()               {}
func (*Recognizer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Recognizer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Recognizer) GetRecognizer() *google_protobuf.Any {
	if m != nil {
		return m.Recognizer
	}
	return nil
}

type RecognizerList struct {
	Recognizers []*Recognizer `protobuf:"bytes,1,rep,name=recognizers" json:"recognizers,omitempty"`
}

func (m *RecognizerList) Reset()                    { *m = RecognizerList{} }
func (m *RecognizerList) String() string            { return proto.CompactTextString(m) }
func (*RecognizerList) ProtoMessage()               {}
func (*RecognizerList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RecognizerList) GetRecognizers() []*Recognizer {
	if m != nil {
		return m.Recognizers
	}
	return nil
}

type ButtonRecognizer struct {
	OnEvent       int64 `protobuf:"varint,1,opt,name=onEvent" json:"onEvent,omitempty"`
	IgnoresScroll bool  `protobuf:"varint,2,opt,name=ignoresScroll" json:"ignoresScroll,omitempty"`
}

func (m *ButtonRecognizer) Reset()                    { *m = ButtonRecognizer{} }
func (m *ButtonRecognizer) String() string            { return proto.CompactTextString(m) }
func (*ButtonRecognizer) ProtoMessage()               {}
func (*ButtonRecognizer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ButtonRecognizer) GetOnEvent() int64 {
	if m != nil {
		return m.OnEvent
	}
	return 0
}

func (m *ButtonRecognizer) GetIgnoresScroll() bool {
	if m != nil {
		return m.IgnoresScroll
	}
	return false
}

type ButtonEvent struct {
	Timestamp *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Inside    bool                        `protobuf:"varint,3,opt,name=inside" json:"inside,omitempty"`
	Kind      EventKind                   `protobuf:"varint,4,opt,name=kind,enum=matcha.touch.EventKind" json:"kind,omitempty"`
}

func (m *ButtonEvent) Reset()                    { *m = ButtonEvent{} }
func (m *ButtonEvent) String() string            { return proto.CompactTextString(m) }
func (*ButtonEvent) ProtoMessage()               {}
func (*ButtonEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ButtonEvent) GetTimestamp() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ButtonEvent) GetInside() bool {
	if m != nil {
		return m.Inside
	}
	return false
}

func (m *ButtonEvent) GetKind() EventKind {
	if m != nil {
		return m.Kind
	}
	return EventKind_EVENT_KIND_POSSIBLE
}

type TapRecognizer struct {
	Count          int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	RecognizedFunc int64 `protobuf:"varint,2,opt,name=recognizedFunc" json:"recognizedFunc,omitempty"`
}

func (m *TapRecognizer) Reset()                    { *m = TapRecognizer{} }
func (m *TapRecognizer) String() string            { return proto.CompactTextString(m) }
func (*TapRecognizer) ProtoMessage()               {}
func (*TapRecognizer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TapRecognizer) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *TapRecognizer) GetRecognizedFunc() int64 {
	if m != nil {
		return m.RecognizedFunc
	}
	return 0
}

type TapEvent struct {
	Timestamp *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Position  *matcha_layout.Point        `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
}

func (m *TapEvent) Reset()                    { *m = TapEvent{} }
func (m *TapEvent) String() string            { return proto.CompactTextString(m) }
func (*TapEvent) ProtoMessage()               {}
func (*TapEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TapEvent) GetTimestamp() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *TapEvent) GetPosition() *matcha_layout.Point {
	if m != nil {
		return m.Position
	}
	return nil
}

type PressRecognizer struct {
	MinDuration *google_protobuf1.Duration `protobuf:"bytes,1,opt,name=minDuration" json:"minDuration,omitempty"`
	FuncId      int64                      `protobuf:"varint,2,opt,name=funcId" json:"funcId,omitempty"`
}

func (m *PressRecognizer) Reset()                    { *m = PressRecognizer{} }
func (m *PressRecognizer) String() string            { return proto.CompactTextString(m) }
func (*PressRecognizer) ProtoMessage()               {}
func (*PressRecognizer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PressRecognizer) GetMinDuration() *google_protobuf1.Duration {
	if m != nil {
		return m.MinDuration
	}
	return nil
}

func (m *PressRecognizer) GetFuncId() int64 {
	if m != nil {
		return m.FuncId
	}
	return 0
}

type PressEvent struct {
	Timestamp *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Position  *matcha_layout.Point        `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	Kind      EventKind                   `protobuf:"varint,3,opt,name=kind,enum=matcha.touch.EventKind" json:"kind,omitempty"`
	Duration  *google_protobuf1.Duration  `protobuf:"bytes,4,opt,name=duration" json:"duration,omitempty"`
}

func (m *PressEvent) Reset()                    { *m = PressEvent{} }
func (m *PressEvent) String() string            { return proto.CompactTextString(m) }
func (*PressEvent) ProtoMessage()               {}
func (*PressEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PressEvent) GetTimestamp() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PressEvent) GetPosition() *matcha_layout.Point {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *PressEvent) GetKind() EventKind {
	if m != nil {
		return m.Kind
	}
	return EventKind_EVENT_KIND_POSSIBLE
}

func (m *PressEvent) GetDuration() *google_protobuf1.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterType((*Recognizer)(nil), "matcha.touch.Recognizer")
	proto.RegisterType((*RecognizerList)(nil), "matcha.touch.RecognizerList")
	proto.RegisterType((*ButtonRecognizer)(nil), "matcha.touch.ButtonRecognizer")
	proto.RegisterType((*ButtonEvent)(nil), "matcha.touch.ButtonEvent")
	proto.RegisterType((*TapRecognizer)(nil), "matcha.touch.TapRecognizer")
	proto.RegisterType((*TapEvent)(nil), "matcha.touch.TapEvent")
	proto.RegisterType((*PressRecognizer)(nil), "matcha.touch.PressRecognizer")
	proto.RegisterType((*PressEvent)(nil), "matcha.touch.PressEvent")
	proto.RegisterEnum("matcha.touch.EventKind", EventKind_name, EventKind_value)
}

func init() { proto.RegisterFile("github.com/gomatcha/matcha/pb/touch/touch2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0xad, 0x71, 0x92, 0x92, 0x71, 0xa1, 0x74, 0x9b, 0x04, 0xc3, 0xa1, 0x45, 0x56, 0x55, 0xa1,
	0x56, 0xb2, 0x23, 0xda, 0x4a, 0x55, 0x7b, 0x82, 0xe0, 0xa4, 0x28, 0x84, 0xa0, 0x05, 0xf5, 0x90,
	0x4b, 0x64, 0x6c, 0x03, 0xab, 0xc2, 0xae, 0x65, 0xaf, 0x91, 0xe8, 0x37, 0xf4, 0x2b, 0xfa, 0x5d,
	0xfd, 0x98, 0x8a, 0xf5, 0xda, 0x38, 0x44, 0x8a, 0x2a, 0x55, 0xca, 0x05, 0x34, 0xf3, 0xde, 0xce,
	0xbc, 0x37, 0x33, 0x06, 0x6b, 0x46, 0xf8, 0x3c, 0x9e, 0x98, 0x2e, 0x5b, 0x5a, 0x6c, 0xe5, 0x87,
	0xee, 0x9a, 0x5a, 0x4b, 0x87, 0xbb, 0x73, 0xc7, 0x0a, 0x26, 0x16, 0x67, 0xb1, 0x3b, 0x4f, 0x7e,
	0x5b, 0x66, 0x10, 0x32, 0xce, 0xd0, 0xb3, 0x04, 0x35, 0x45, 0xb2, 0x5e, 0x9b, 0x31, 0x36, 0x5b,
	0xf8, 0x96, 0xc0, 0x26, 0xf1, 0xd4, 0x72, 0xe8, 0x3a, 0x21, 0xd6, 0x5f, 0xed, 0x42, 0x5e, 0x1c,
	0x3a, 0x9c, 0x30, 0x2a, 0xf1, 0xd7, 0xbb, 0x38, 0x27, 0x4b, 0x3f, 0xe2, 0xce, 0x32, 0x90, 0x84,
	0xd3, 0x07, 0xa5, 0x2d, 0x9c, 0x35, 0x8b, 0xb9, 0xfc, 0x4b, 0x5e, 0x18, 0x18, 0x00, 0xfb, 0x2e,
	0x9b, 0x51, 0xf2, 0xd3, 0x0f, 0x51, 0x19, 0x0a, 0xc4, 0xd3, 0x95, 0x86, 0xd2, 0x54, 0x71, 0x81,
	0x78, 0xe8, 0x23, 0x40, 0x98, 0xa1, 0xba, 0xda, 0x50, 0x9a, 0x5a, 0xeb, 0xc8, 0x4c, 0x54, 0x98,
	0xa9, 0x0a, 0xb3, 0x4d, 0xd7, 0x38, 0xc7, 0x33, 0xfa, 0x50, 0xde, 0xd6, 0xec, 0x93, 0x88, 0xa3,
	0x2f, 0xa0, 0x6d, 0xf1, 0x48, 0x57, 0x1a, 0x6a, 0x53, 0x6b, 0xe9, 0x66, 0x7e, 0x2e, 0xe6, 0xf6,
	0x09, 0xce, 0x93, 0x0d, 0x0c, 0x95, 0x4e, 0xcc, 0x39, 0xa3, 0x39, 0x9d, 0x3a, 0x3c, 0x65, 0xd4,
	0x5e, 0xf9, 0x94, 0x4b, 0xb1, 0x69, 0x88, 0xde, 0x40, 0x89, 0xcc, 0x28, 0x0b, 0xfd, 0x68, 0xe4,
	0x86, 0x6c, 0xb1, 0xd0, 0x0b, 0x0d, 0xa5, 0x59, 0xc4, 0x77, 0x93, 0xc6, 0x2f, 0x05, 0xb4, 0xa4,
	0x68, 0xf2, 0xea, 0x33, 0x1c, 0x66, 0xa3, 0x14, 0x15, 0xb5, 0x56, 0xfd, 0x9e, 0xcd, 0x71, 0xca,
	0xc0, 0x5b, 0x32, 0x3a, 0x81, 0x03, 0x42, 0x23, 0xe2, 0xf9, 0x62, 0x3a, 0x45, 0x2c, 0x23, 0xf4,
	0x1e, 0xf6, 0x7e, 0x10, 0xea, 0xe9, 0x7b, 0x0d, 0xa5, 0x59, 0x6e, 0x55, 0xef, 0x5a, 0x15, 0x4d,
	0x2f, 0x09, 0xf5, 0xb0, 0x20, 0x19, 0x57, 0x50, 0x1a, 0x3b, 0x41, 0xce, 0xdf, 0x11, 0xec, 0xbb,
	0x2c, 0xce, 0xdc, 0x25, 0x01, 0x7a, 0x0b, 0xe5, 0x6c, 0x30, 0xde, 0x79, 0x4c, 0x5d, 0x61, 0x4e,
	0xc5, 0x3b, 0x59, 0x63, 0x05, 0xc5, 0xb1, 0x13, 0xfc, 0xaf, 0xb3, 0x53, 0x28, 0x06, 0x2c, 0x22,
	0x9b, 0xf3, 0x13, 0x7d, 0x36, 0x9b, 0x97, 0x2e, 0xe4, 0x05, 0x0d, 0x19, 0xa1, 0x1c, 0x67, 0x2c,
	0x63, 0x0a, 0xcf, 0x87, 0xa1, 0x1f, 0x45, 0x39, 0x23, 0x5f, 0x41, 0x5b, 0x12, 0xda, 0x95, 0x67,
	0x2c, 0x05, 0xd4, 0xee, 0x09, 0x48, 0x09, 0x38, 0xcf, 0xde, 0xcc, 0x76, 0x1a, 0x53, 0xb7, 0xe7,
	0x49, 0x9f, 0x32, 0x32, 0xfe, 0x28, 0x00, 0xa2, 0xd1, 0xa3, 0x5b, 0xcc, 0xd6, 0xaa, 0xfe, 0xc3,
	0x5a, 0xd1, 0x27, 0x28, 0xa6, 0x1f, 0xb0, 0xb8, 0x83, 0x07, 0x9d, 0x67, 0xd4, 0x77, 0x14, 0x0e,
	0xb3, 0x4a, 0xa8, 0x0a, 0x2f, 0xed, 0xef, 0xf6, 0x60, 0x7c, 0x7b, 0xd9, 0x1b, 0x74, 0x6f, 0x87,
	0xd7, 0xa3, 0x51, 0xaf, 0xd3, 0xb7, 0x2b, 0x4f, 0xd0, 0x09, 0xa0, 0x1c, 0x70, 0xf6, 0xad, 0x3d,
	0xb8, 0xb0, 0xbb, 0x15, 0x05, 0x1d, 0xc3, 0x8b, 0x5c, 0xfe, 0xbc, 0xdd, 0xeb, 0xdb, 0xdd, 0x4a,
	0x01, 0xd5, 0xe0, 0x38, 0x97, 0xc6, 0xf6, 0xd9, 0xf5, 0xc5, 0xa0, 0x77, 0x63, 0x77, 0x2b, 0x6a,
	0xa7, 0x7a, 0xb3, 0x2f, 0xf4, 0xff, 0x2e, 0x94, 0xae, 0x84, 0x9d, 0x61, 0x67, 0xbc, 0x89, 0x27,
	0x07, 0x42, 0xe5, 0x87, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x06, 0xdf, 0x72, 0xf1, 0x04,
	0x00, 0x00,
}
