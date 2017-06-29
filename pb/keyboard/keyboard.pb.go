// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/matcha/pb/keyboard/keyboard.proto

/*
Package keyboard is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/matcha/pb/keyboard/keyboard.proto

It has these top-level messages:
*/
package keyboard

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

type Type int32

const (
	Type_DEFAULT_TYPE            Type = 0
	Type_NUMBER_TYPE             Type = 1
	Type_NUMBER_PUNCTUATION_TYPE Type = 2
	Type_DECIMAL_TYPE            Type = 3
	Type_PHONE_TYPE              Type = 4
	Type_ASCII_TYPE              Type = 5
	Type_EMAIL_TYPE              Type = 6
	Type_URL_TYPE                Type = 7
	Type_WEB_SEARCH_TYPE         Type = 8
	Type_NAME_PHONE_TYPE         Type = 9
)

var Type_name = map[int32]string{
	0: "DEFAULT_TYPE",
	1: "NUMBER_TYPE",
	2: "NUMBER_PUNCTUATION_TYPE",
	3: "DECIMAL_TYPE",
	4: "PHONE_TYPE",
	5: "ASCII_TYPE",
	6: "EMAIL_TYPE",
	7: "URL_TYPE",
	8: "WEB_SEARCH_TYPE",
	9: "NAME_PHONE_TYPE",
}
var Type_value = map[string]int32{
	"DEFAULT_TYPE":            0,
	"NUMBER_TYPE":             1,
	"NUMBER_PUNCTUATION_TYPE": 2,
	"DECIMAL_TYPE":            3,
	"PHONE_TYPE":              4,
	"ASCII_TYPE":              5,
	"EMAIL_TYPE":              6,
	"URL_TYPE":                7,
	"WEB_SEARCH_TYPE":         8,
	"NAME_PHONE_TYPE":         9,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Appearance int32

const (
	Appearance_DEFAULT_APPEARANCE Appearance = 0
	Appearance_LIGHT_APPEARANCE   Appearance = 1
	Appearance_DARK_APPEARANCE    Appearance = 2
)

var Appearance_name = map[int32]string{
	0: "DEFAULT_APPEARANCE",
	1: "LIGHT_APPEARANCE",
	2: "DARK_APPEARANCE",
}
var Appearance_value = map[string]int32{
	"DEFAULT_APPEARANCE": 0,
	"LIGHT_APPEARANCE":   1,
	"DARK_APPEARANCE":    2,
}

func (x Appearance) String() string {
	return proto.EnumName(Appearance_name, int32(x))
}
func (Appearance) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ReturnType int32

const (
	ReturnType_DEFAULT_RETURN_TYPE        ReturnType = 0
	ReturnType_GO_RETURN_TYPE             ReturnType = 1
	ReturnType_GOOGLE_RETURN_TYPE         ReturnType = 2
	ReturnType_JOIN_RETURN_TYPE           ReturnType = 3
	ReturnType_NEXT_RETURN_TYPE           ReturnType = 4
	ReturnType_ROUTE_RETURN_TYPE          ReturnType = 5
	ReturnType_SEARCH_RETURN_TYPE         ReturnType = 6
	ReturnType_SEND_RETURN_TYPE           ReturnType = 7
	ReturnType_YAHOO_RETURN_TYPE          ReturnType = 8
	ReturnType_DONE_RETURN_TYPE           ReturnType = 9
	ReturnType_EMERGENCY_CALL_RETURN_TYPE ReturnType = 10
	ReturnType_CONTINUE_RETURN_TYPE       ReturnType = 11
)

var ReturnType_name = map[int32]string{
	0:  "DEFAULT_RETURN_TYPE",
	1:  "GO_RETURN_TYPE",
	2:  "GOOGLE_RETURN_TYPE",
	3:  "JOIN_RETURN_TYPE",
	4:  "NEXT_RETURN_TYPE",
	5:  "ROUTE_RETURN_TYPE",
	6:  "SEARCH_RETURN_TYPE",
	7:  "SEND_RETURN_TYPE",
	8:  "YAHOO_RETURN_TYPE",
	9:  "DONE_RETURN_TYPE",
	10: "EMERGENCY_CALL_RETURN_TYPE",
	11: "CONTINUE_RETURN_TYPE",
}
var ReturnType_value = map[string]int32{
	"DEFAULT_RETURN_TYPE":        0,
	"GO_RETURN_TYPE":             1,
	"GOOGLE_RETURN_TYPE":         2,
	"JOIN_RETURN_TYPE":           3,
	"NEXT_RETURN_TYPE":           4,
	"ROUTE_RETURN_TYPE":          5,
	"SEARCH_RETURN_TYPE":         6,
	"SEND_RETURN_TYPE":           7,
	"YAHOO_RETURN_TYPE":          8,
	"DONE_RETURN_TYPE":           9,
	"EMERGENCY_CALL_RETURN_TYPE": 10,
	"CONTINUE_RETURN_TYPE":       11,
}

func (x ReturnType) String() string {
	return proto.EnumName(ReturnType_name, int32(x))
}
func (ReturnType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterEnum("matcha.keyboard.Type", Type_name, Type_value)
	proto.RegisterEnum("matcha.keyboard.Appearance", Appearance_name, Appearance_value)
	proto.RegisterEnum("matcha.keyboard.ReturnType", ReturnType_name, ReturnType_value)
}

func init() {
	proto.RegisterFile("github.com/gomatcha/matcha/pb/keyboard/keyboard.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x69, 0xd6, 0x75, 0xdd, 0xd9, 0xb4, 0x16, 0x6f, 0x30, 0x04, 0x82, 0x07, 0xe8, 0x45,
	0x7b, 0x01, 0x2f, 0xe0, 0xba, 0x87, 0xd4, 0x2c, 0xb1, 0x23, 0xd7, 0x16, 0x94, 0x9b, 0x2a, 0x09,
	0x11, 0x43, 0x68, 0x4d, 0x14, 0xa5, 0x48, 0x7d, 0x1d, 0xee, 0x78, 0x0b, 0x1e, 0x0d, 0xb9, 0x4e,
	0x50, 0x7c, 0x97, 0xff, 0xf3, 0xf1, 0xe7, 0x9c, 0xa3, 0x03, 0x1f, 0xbe, 0xff, 0x68, 0x1e, 0x0f,
	0xd9, 0x3c, 0x2f, 0x9f, 0x16, 0xe5, 0xaf, 0xa2, 0xce, 0x8f, 0xfb, 0xc5, 0x53, 0xda, 0xe4, 0x8f,
	0xe9, 0xa2, 0xca, 0x16, 0x3f, 0x8b, 0x63, 0x56, 0xa6, 0xf5, 0xb7, 0xff, 0x1f, 0xf3, 0xaa, 0x2e,
	0x9b, 0x92, 0x4c, 0x5c, 0xc9, 0xbc, 0xc3, 0xb3, 0xbf, 0x03, 0x18, 0xea, 0x63, 0x55, 0x90, 0x29,
	0x5c, 0xaf, 0xf0, 0x23, 0x35, 0x91, 0xde, 0xe9, 0x6d, 0x82, 0xd3, 0x67, 0x64, 0x02, 0x57, 0xc2,
	0xc4, 0x4b, 0x54, 0x0e, 0x0c, 0xc8, 0x1b, 0xb8, 0x6f, 0x41, 0x62, 0x04, 0xd3, 0x86, 0x6a, 0x2e,
	0x85, 0x3b, 0x0c, 0xdc, 0x7d, 0xc6, 0x63, 0x1a, 0x39, 0x72, 0x46, 0x6e, 0x00, 0x92, 0xb5, 0x14,
	0xe8, 0xf2, 0xd0, 0x66, 0xba, 0x61, 0x9c, 0xbb, 0x7c, 0x6e, 0x33, 0xc6, 0x94, 0xb7, 0xf5, 0x23,
	0x72, 0x0d, 0x63, 0xa3, 0xda, 0x74, 0x41, 0x6e, 0x61, 0xf2, 0x19, 0x97, 0xbb, 0x0d, 0x52, 0xc5,
	0xd6, 0x0e, 0x8e, 0x2d, 0x14, 0x34, 0xc6, 0x5d, 0xcf, 0x7b, 0x39, 0x93, 0x00, 0xb4, 0xaa, 0x8a,
	0xb4, 0x4e, 0xf7, 0x79, 0x41, 0x5e, 0x02, 0xe9, 0xfa, 0xa0, 0x49, 0x82, 0x54, 0x51, 0xc1, 0x6c,
	0x37, 0x77, 0x30, 0x8d, 0x78, 0xb8, 0xf6, 0xe8, 0xc0, 0x0a, 0x57, 0x54, 0x3d, 0xf4, 0x61, 0x30,
	0xfb, 0x13, 0x00, 0xa8, 0xa2, 0x39, 0xd4, 0xfb, 0xd3, 0x64, 0xee, 0xe1, 0xb6, 0x33, 0x2a, 0xd4,
	0x46, 0x89, 0x6e, 0x40, 0x04, 0x6e, 0x42, 0xe9, 0xb1, 0x81, 0x7d, 0x3e, 0x94, 0x32, 0x8c, 0xd0,
	0xe3, 0x81, 0x7d, 0xfe, 0x93, 0xe4, 0xc2, 0xa3, 0x67, 0x96, 0x0a, 0xfc, 0xe2, 0x7b, 0x87, 0xe4,
	0x05, 0x3c, 0x57, 0xd2, 0x68, 0x5f, 0x71, 0x6e, 0xd5, 0xed, 0x34, 0xfa, 0x7c, 0x64, 0x25, 0x1b,
	0x14, 0x2b, 0x8f, 0x5e, 0x58, 0xc9, 0x96, 0xae, 0xa5, 0xff, 0x7f, 0x63, 0x5b, 0xbc, 0xb2, 0xb3,
	0xeb, 0xd3, 0x4b, 0xf2, 0x0e, 0x5e, 0x63, 0x8c, 0x2a, 0x44, 0xc1, 0xb6, 0x3b, 0x46, 0xa3, 0xc8,
	0x3b, 0x07, 0xf2, 0x0a, 0xee, 0x98, 0x14, 0x9a, 0x0b, 0xe3, 0xdf, 0xbc, 0x5a, 0xbe, 0xfd, 0x3a,
	0xee, 0x76, 0xe9, 0x77, 0x30, 0x8d, 0x4f, 0xdb, 0xf5, 0xd0, 0x82, 0x64, 0x99, 0x8d, 0x4e, 0x6b,
	0xf7, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xde, 0x61, 0x94, 0xae, 0x02, 0x00, 0x00,
}
