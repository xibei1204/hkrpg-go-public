// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LineupInfo.proto

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

type LineupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaneId         uint32          `protobuf:"varint,11,opt,name=plane_id,json=planeId,proto3" json:"plane_id,omitempty"`
	AvatarList      []*LineupAvatar `protobuf:"bytes,15,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
	Name            string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Index           uint32          `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	LeaderSlot      uint32          `protobuf:"varint,7,opt,name=leader_slot,json=leaderSlot,proto3" json:"leader_slot,omitempty"`
	ExtraLineupType ExtraLineupType `protobuf:"varint,13,opt,name=extra_lineup_type,json=extraLineupType,proto3,enum=ExtraLineupType" json:"extra_lineup_type,omitempty"`
	IsVirtual       bool            `protobuf:"varint,8,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual,omitempty"`
	MaxMp           uint32          `protobuf:"varint,6,opt,name=max_mp,json=maxMp,proto3" json:"max_mp,omitempty"`
	Mp              uint32          `protobuf:"varint,9,opt,name=mp,proto3" json:"mp,omitempty"`
}

func (x *LineupInfo) Reset() {
	*x = LineupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LineupInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineupInfo) ProtoMessage() {}

func (x *LineupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_LineupInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineupInfo.ProtoReflect.Descriptor instead.
func (*LineupInfo) Descriptor() ([]byte, []int) {
	return file_LineupInfo_proto_rawDescGZIP(), []int{0}
}

func (x *LineupInfo) GetPlaneId() uint32 {
	if x != nil {
		return x.PlaneId
	}
	return 0
}

func (x *LineupInfo) GetAvatarList() []*LineupAvatar {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

func (x *LineupInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LineupInfo) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LineupInfo) GetLeaderSlot() uint32 {
	if x != nil {
		return x.LeaderSlot
	}
	return 0
}

func (x *LineupInfo) GetExtraLineupType() ExtraLineupType {
	if x != nil {
		return x.ExtraLineupType
	}
	return ExtraLineupType_LINEUP_NONE
}

func (x *LineupInfo) GetIsVirtual() bool {
	if x != nil {
		return x.IsVirtual
	}
	return false
}

func (x *LineupInfo) GetMaxMp() uint32 {
	if x != nil {
		return x.MaxMp
	}
	return 0
}

func (x *LineupInfo) GetMp() uint32 {
	if x != nil {
		return x.Mp
	}
	return 0
}

var File_LineupInfo_proto protoreflect.FileDescriptor

var file_LineupInfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02,
	0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4c,
	0x69, 0x6e, 0x65, 0x75, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x0a, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x6c,
	0x6f, 0x74, 0x12, 0x3c, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x12,
	0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6d, 0x61, 0x78, 0x4d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x6d, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LineupInfo_proto_rawDescOnce sync.Once
	file_LineupInfo_proto_rawDescData = file_LineupInfo_proto_rawDesc
)

func file_LineupInfo_proto_rawDescGZIP() []byte {
	file_LineupInfo_proto_rawDescOnce.Do(func() {
		file_LineupInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_LineupInfo_proto_rawDescData)
	})
	return file_LineupInfo_proto_rawDescData
}

var file_LineupInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LineupInfo_proto_goTypes = []interface{}{
	(*LineupInfo)(nil),   // 0: LineupInfo
	(*LineupAvatar)(nil), // 1: LineupAvatar
	(ExtraLineupType)(0), // 2: ExtraLineupType
}
var file_LineupInfo_proto_depIdxs = []int32{
	1, // 0: LineupInfo.avatar_list:type_name -> LineupAvatar
	2, // 1: LineupInfo.extra_lineup_type:type_name -> ExtraLineupType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_LineupInfo_proto_init() }
func file_LineupInfo_proto_init() {
	if File_LineupInfo_proto != nil {
		return
	}
	file_ExtraLineupType_proto_init()
	file_LineupAvatar_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LineupInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineupInfo); i {
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
			RawDescriptor: file_LineupInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LineupInfo_proto_goTypes,
		DependencyIndexes: file_LineupInfo_proto_depIdxs,
		MessageInfos:      file_LineupInfo_proto_msgTypes,
	}.Build()
	File_LineupInfo_proto = out.File
	file_LineupInfo_proto_rawDesc = nil
	file_LineupInfo_proto_goTypes = nil
	file_LineupInfo_proto_depIdxs = nil
}
