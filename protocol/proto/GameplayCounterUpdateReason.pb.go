// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GameplayCounterUpdateReason.proto

package proto

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

type GameplayCounterUpdateReason int32

const (
	GameplayCounterUpdateReason_GAMEPLAY_COUNTER_UPDATE_REASON_NONE       GameplayCounterUpdateReason = 0
	GameplayCounterUpdateReason_GAMEPLAY_COUNTER_UPDATE_REASON_ACTIVATE   GameplayCounterUpdateReason = 1
	GameplayCounterUpdateReason_GAMEPLAY_COUNTER_UPDATE_REASON_DEACTIVATE GameplayCounterUpdateReason = 2
	GameplayCounterUpdateReason_GAMEPLAY_COUNTER_UPDATE_REASON_CHANGE     GameplayCounterUpdateReason = 3
)

// Enum value maps for GameplayCounterUpdateReason.
var (
	GameplayCounterUpdateReason_name = map[int32]string{
		0: "GAMEPLAY_COUNTER_UPDATE_REASON_NONE",
		1: "GAMEPLAY_COUNTER_UPDATE_REASON_ACTIVATE",
		2: "GAMEPLAY_COUNTER_UPDATE_REASON_DEACTIVATE",
		3: "GAMEPLAY_COUNTER_UPDATE_REASON_CHANGE",
	}
	GameplayCounterUpdateReason_value = map[string]int32{
		"GAMEPLAY_COUNTER_UPDATE_REASON_NONE":       0,
		"GAMEPLAY_COUNTER_UPDATE_REASON_ACTIVATE":   1,
		"GAMEPLAY_COUNTER_UPDATE_REASON_DEACTIVATE": 2,
		"GAMEPLAY_COUNTER_UPDATE_REASON_CHANGE":     3,
	}
)

func (x GameplayCounterUpdateReason) Enum() *GameplayCounterUpdateReason {
	p := new(GameplayCounterUpdateReason)
	*p = x
	return p
}

func (x GameplayCounterUpdateReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameplayCounterUpdateReason) Descriptor() protoreflect.EnumDescriptor {
	return file_GameplayCounterUpdateReason_proto_enumTypes[0].Descriptor()
}

func (GameplayCounterUpdateReason) Type() protoreflect.EnumType {
	return &file_GameplayCounterUpdateReason_proto_enumTypes[0]
}

func (x GameplayCounterUpdateReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameplayCounterUpdateReason.Descriptor instead.
func (GameplayCounterUpdateReason) EnumDescriptor() ([]byte, []int) {
	return file_GameplayCounterUpdateReason_proto_rawDescGZIP(), []int{0}
}

var File_GameplayCounterUpdateReason_proto protoreflect.FileDescriptor

var file_GameplayCounterUpdateReason_proto_rawDesc = []byte{
	0x0a, 0x21, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xcd, 0x01, 0x0a, 0x1b, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x23, 0x47, 0x41, 0x4d, 0x45, 0x50, 0x4c, 0x41, 0x59, 0x5f,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27,
	0x47, 0x41, 0x4d, 0x45, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x2d, 0x0a, 0x29, 0x47, 0x41, 0x4d,
	0x45, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x29, 0x0a, 0x25, 0x47, 0x41, 0x4d, 0x45,
	0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GameplayCounterUpdateReason_proto_rawDescOnce sync.Once
	file_GameplayCounterUpdateReason_proto_rawDescData = file_GameplayCounterUpdateReason_proto_rawDesc
)

func file_GameplayCounterUpdateReason_proto_rawDescGZIP() []byte {
	file_GameplayCounterUpdateReason_proto_rawDescOnce.Do(func() {
		file_GameplayCounterUpdateReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_GameplayCounterUpdateReason_proto_rawDescData)
	})
	return file_GameplayCounterUpdateReason_proto_rawDescData
}

var file_GameplayCounterUpdateReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GameplayCounterUpdateReason_proto_goTypes = []interface{}{
	(GameplayCounterUpdateReason)(0), // 0: GameplayCounterUpdateReason
}
var file_GameplayCounterUpdateReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GameplayCounterUpdateReason_proto_init() }
func file_GameplayCounterUpdateReason_proto_init() {
	if File_GameplayCounterUpdateReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GameplayCounterUpdateReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GameplayCounterUpdateReason_proto_goTypes,
		DependencyIndexes: file_GameplayCounterUpdateReason_proto_depIdxs,
		EnumInfos:         file_GameplayCounterUpdateReason_proto_enumTypes,
	}.Build()
	File_GameplayCounterUpdateReason_proto = out.File
	file_GameplayCounterUpdateReason_proto_rawDesc = nil
	file_GameplayCounterUpdateReason_proto_goTypes = nil
	file_GameplayCounterUpdateReason_proto_depIdxs = nil
}