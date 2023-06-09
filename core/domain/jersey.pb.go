// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: jersey.proto

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

type Jersey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JerseyId string  `protobuf:"bytes,1,opt,name=jersey_id,json=jerseyId,proto3" json:"jersey_id,omitempty"`
	Team     *Team   `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	Player   *Player `protobuf:"bytes,3,opt,name=player,proto3" json:"player,omitempty"`
	Number   int32   `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Jersey) Reset() {
	*x = Jersey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jersey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jersey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jersey) ProtoMessage() {}

func (x *Jersey) ProtoReflect() protoreflect.Message {
	mi := &file_jersey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jersey.ProtoReflect.Descriptor instead.
func (*Jersey) Descriptor() ([]byte, []int) {
	return file_jersey_proto_rawDescGZIP(), []int{0}
}

func (x *Jersey) GetJerseyId() string {
	if x != nil {
		return x.JerseyId
	}
	return ""
}

func (x *Jersey) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

func (x *Jersey) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *Jersey) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_jersey_proto protoreflect.FileDescriptor

var file_jersey_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6a, 0x65, 0x72, 0x73, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x87, 0x01, 0x0a, 0x06, 0x4a, 0x65, 0x72, 0x73, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6a,
	0x65, 0x72, 0x73, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6a, 0x65, 0x72, 0x73, 0x65, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jersey_proto_rawDescOnce sync.Once
	file_jersey_proto_rawDescData = file_jersey_proto_rawDesc
)

func file_jersey_proto_rawDescGZIP() []byte {
	file_jersey_proto_rawDescOnce.Do(func() {
		file_jersey_proto_rawDescData = protoimpl.X.CompressGZIP(file_jersey_proto_rawDescData)
	})
	return file_jersey_proto_rawDescData
}

var file_jersey_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_jersey_proto_goTypes = []interface{}{
	(*Jersey)(nil), // 0: domain.Jersey
	(*Team)(nil),   // 1: domain.Team
	(*Player)(nil), // 2: domain.Player
}
var file_jersey_proto_depIdxs = []int32{
	1, // 0: domain.Jersey.team:type_name -> domain.Team
	2, // 1: domain.Jersey.player:type_name -> domain.Player
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_jersey_proto_init() }
func file_jersey_proto_init() {
	if File_jersey_proto != nil {
		return
	}
	file_team_proto_init()
	file_player_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_jersey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jersey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jersey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jersey_proto_goTypes,
		DependencyIndexes: file_jersey_proto_depIdxs,
		MessageInfos:      file_jersey_proto_msgTypes,
	}.Build()
	File_jersey_proto = out.File
	file_jersey_proto_rawDesc = nil
	file_jersey_proto_goTypes = nil
	file_jersey_proto_depIdxs = nil
}
