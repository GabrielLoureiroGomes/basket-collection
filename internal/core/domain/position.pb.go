// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: position.proto

package domain

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Position int32

const (
	Position_POINT_GUARD    Position = 0
	Position_SHOOTING_GUARD Position = 1
	Position_SMALL_FORWARD  Position = 2
	Position_POWER_FORWARD  Position = 3
	Position_CENTER         Position = 4
)

// Enum value maps for Position.
var (
	Position_name = map[int32]string{
		0: "POINT_GUARD",
		1: "SHOOTING_GUARD",
		2: "SMALL_FORWARD",
		3: "POWER_FORWARD",
		4: "CENTER",
	}
	Position_value = map[string]int32{
		"POINT_GUARD":    0,
		"SHOOTING_GUARD": 1,
		"SMALL_FORWARD":  2,
		"POWER_FORWARD":  3,
		"CENTER":         4,
	}
)

func (x Position) Enum() *Position {
	p := new(Position)
	*p = x
	return p
}

func (x Position) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Position) Descriptor() protoreflect.EnumDescriptor {
	return file_position_proto_enumTypes[0].Descriptor()
}

func (Position) Type() protoreflect.EnumType {
	return &file_position_proto_enumTypes[0]
}

func (x Position) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Position.Descriptor instead.
func (Position) EnumDescriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{0}
}

var File_position_proto protoreflect.FileDescriptor

var file_position_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2a, 0x61, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x47, 0x55,
	0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x48, 0x4f, 0x4f, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x47, 0x55, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x4d, 0x41,
	0x4c, 0x4c, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d,
	0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x04, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_position_proto_rawDescOnce sync.Once
	file_position_proto_rawDescData = file_position_proto_rawDesc
)

func file_position_proto_rawDescGZIP() []byte {
	file_position_proto_rawDescOnce.Do(func() {
		file_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_position_proto_rawDescData)
	})
	return file_position_proto_rawDescData
}

var file_position_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_position_proto_goTypes = []interface{}{
	(Position)(0), // 0: domain.Position
}
var file_position_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_position_proto_init() }
func file_position_proto_init() {
	if File_position_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_position_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_position_proto_goTypes,
		DependencyIndexes: file_position_proto_depIdxs,
		EnumInfos:         file_position_proto_enumTypes,
	}.Build()
	File_position_proto = out.File
	file_position_proto_rawDesc = nil
	file_position_proto_goTypes = nil
	file_position_proto_depIdxs = nil
}