// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/field_behavior.proto

package annotations

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// An indicator of the behavior of a given field (for example, that a field
// is required in requests, or given as output but ignored as input).
// This **does not** change the behavior in protocol buffers itself; it only
// denotes the behavior and may affect how API tooling handles the field.
//
// Note: This enum **may** receive new values in the future.
type FieldBehavior int32

const (
	// Conventional default for enums. Do not use this.
	FieldBehavior_FIELD_BEHAVIOR_UNSPECIFIED FieldBehavior = 0
	// Specifically denotes a field as optional.
	// While all fields in protocol buffers are optional, this may be specified
	// for emphasis if appropriate.
	FieldBehavior_OPTIONAL FieldBehavior = 1
	// Denotes a field as required.
	// This indicates that the field **must** be provided as part of the request,
	// and failure to do so will cause an error (usually `INVALID_ARGUMENT`).
	FieldBehavior_REQUIRED FieldBehavior = 2
	// Denotes a field as output only.
	// This indicates that the field is provided in responses, but including the
	// field in a request does nothing (the server *must* ignore it and
	// *must not* throw an error as a result of the field's presence).
	FieldBehavior_OUTPUT_ONLY FieldBehavior = 3
	// Denotes a field as input only.
	// This indicates that the field is provided in requests, and the
	// corresponding field is not included in output.
	FieldBehavior_INPUT_ONLY FieldBehavior = 4
	// Denotes a field as immutable.
	// This indicates that the field may be set once in a request to create a
	// resource, but may not be changed thereafter.
	FieldBehavior_IMMUTABLE FieldBehavior = 5
)

var FieldBehavior_name = map[int32]string{
	0: "FIELD_BEHAVIOR_UNSPECIFIED",
	1: "OPTIONAL",
	2: "REQUIRED",
	3: "OUTPUT_ONLY",
	4: "INPUT_ONLY",
	5: "IMMUTABLE",
}

var FieldBehavior_value = map[string]int32{
	"FIELD_BEHAVIOR_UNSPECIFIED": 0,
	"OPTIONAL":                   1,
	"REQUIRED":                   2,
	"OUTPUT_ONLY":                3,
	"INPUT_ONLY":                 4,
	"IMMUTABLE":                  5,
}

func (x FieldBehavior) String() string {
	return proto.EnumName(FieldBehavior_name, int32(x))
}

func (FieldBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4648f18fd5079967, []int{0}
}

var E_FieldBehavior = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: ([]FieldBehavior)(nil),
	Field:         1052,
	Name:          "google.api.field_behavior",
	Tag:           "varint,1052,rep,name=field_behavior,enum=google.api.FieldBehavior",
	Filename:      "google/api/field_behavior.proto",
}

func init() {
	proto.RegisterEnum("google.api.FieldBehavior", FieldBehavior_name, FieldBehavior_value)
	proto.RegisterExtension(E_FieldBehavior)
}

func init() {
	proto.RegisterFile("google/api/field_behavior.proto", fileDescriptor_4648f18fd5079967)
}

var fileDescriptor_4648f18fd5079967 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4f, 0xb3, 0x30,
	0x1c, 0xc7, 0x9f, 0xfd, 0x79, 0xcc, 0xac, 0x0e, 0x49, 0x4f, 0xba, 0x44, 0xdd, 0xd1, 0x78, 0x28,
	0x89, 0xde, 0xf4, 0x04, 0xae, 0xd3, 0x26, 0x8c, 0x56, 0x04, 0x13, 0xbd, 0x60, 0xb7, 0xb1, 0xda,
	0x64, 0xd2, 0x06, 0xd0, 0x8b, 0x6f, 0xc5, 0x93, 0xaf, 0xd4, 0xd0, 0x31, 0x85, 0x5b, 0xbf, 0xf9,
	0x7d, 0xfa, 0xeb, 0xe7, 0x5b, 0x70, 0x2a, 0x94, 0x12, 0xeb, 0xd4, 0xe1, 0x5a, 0x3a, 0x2b, 0x99,
	0xae, 0x97, 0xc9, 0x3c, 0x7d, 0xe5, 0x1f, 0x52, 0xe5, 0x48, 0xe7, 0xaa, 0x54, 0x10, 0x6c, 0x00,
	0xc4, 0xb5, 0x1c, 0x8d, 0x6b, 0xd8, 0x4c, 0xe6, 0xef, 0x2b, 0x67, 0x99, 0x16, 0x8b, 0x5c, 0xea,
	0x72, 0x4b, 0x9f, 0x7f, 0x82, 0xe1, 0xb4, 0xda, 0xe2, 0xd5, 0x4b, 0xe0, 0x09, 0x18, 0x4d, 0x09,
	0xf6, 0x27, 0x89, 0x87, 0xef, 0xdc, 0x47, 0x42, 0xc3, 0x24, 0x0e, 0x1e, 0x18, 0xbe, 0x21, 0x53,
	0x82, 0x27, 0xf6, 0x3f, 0xb8, 0x0f, 0x06, 0x94, 0x45, 0x84, 0x06, 0xae, 0x6f, 0x77, 0xaa, 0x14,
	0xe2, 0xfb, 0x98, 0x84, 0x78, 0x62, 0x77, 0xe1, 0x01, 0xd8, 0xa3, 0x71, 0xc4, 0xe2, 0x28, 0xa1,
	0x81, 0xff, 0x64, 0xf7, 0xa0, 0x05, 0x00, 0x09, 0x7e, 0x73, 0x1f, 0x0e, 0xc1, 0x2e, 0x99, 0xcd,
	0xe2, 0xc8, 0xf5, 0x7c, 0x6c, 0xff, 0xbf, 0x7a, 0x01, 0x56, 0xbb, 0x02, 0x3c, 0x46, 0xb5, 0xfd,
	0xd6, 0x18, 0x19, 0x3b, 0xaa, 0x4b, 0xa9, 0xb2, 0xe2, 0xf0, 0x6b, 0x30, 0xee, 0x9d, 0x59, 0x17,
	0x47, 0xe8, 0xaf, 0x23, 0x6a, 0xe9, 0x87, 0xc3, 0x55, 0x33, 0x7a, 0x1a, 0x58, 0x0b, 0xf5, 0xd6,
	0xc0, 0x3d, 0xd8, 0xe2, 0x59, 0xf5, 0x0c, 0xeb, 0x3c, 0xbb, 0x35, 0x21, 0xd4, 0x9a, 0x67, 0x02,
	0xa9, 0x5c, 0x38, 0x22, 0xcd, 0x8c, 0x84, 0xb3, 0x19, 0x71, 0x2d, 0x0b, 0xf3, 0xe9, 0x3c, 0xcb,
	0x54, 0xc9, 0x8d, 0xcf, 0x75, 0xe3, 0xfc, 0xdd, 0xed, 0xdf, 0xba, 0x8c, 0xcc, 0x77, 0xcc, 0xa5,
	0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x94, 0x57, 0x94, 0xa8, 0x01, 0x00, 0x00,
}
