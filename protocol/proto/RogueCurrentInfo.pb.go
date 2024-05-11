// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueCurrentInfo.proto

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

type RogueCurrentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueAeonInfo    *RogueAeon                `protobuf:"bytes,3,opt,name=rogue_aeon_info,json=rogueAeonInfo,proto3" json:"rogue_aeon_info,omitempty"`
	RogueMiracleInfo *RogueMiracleInfo         `protobuf:"bytes,2,opt,name=rogue_miracle_info,json=rogueMiracleInfo,proto3" json:"rogue_miracle_info,omitempty"`
	RogueAvatarInfo  *RogueAvatarInfo          `protobuf:"bytes,6,opt,name=rogue_avatar_info,json=rogueAvatarInfo,proto3" json:"rogue_avatar_info,omitempty"`
	Status           RogueStatus               `protobuf:"varint,7,opt,name=status,proto3,enum=RogueStatus" json:"status,omitempty"`
	RoomMap          *RogueMapInfo             `protobuf:"bytes,9,opt,name=room_map,json=roomMap,proto3" json:"room_map,omitempty"`
	PendingAction    *RogueCommonPendingAction `protobuf:"bytes,8,opt,name=pending_action,json=pendingAction,proto3" json:"pending_action,omitempty"`
	RogueVirtualItem *RogueVirtualItemInfo     `protobuf:"bytes,10,opt,name=rogue_virtual_item,json=rogueVirtualItem,proto3" json:"rogue_virtual_item,omitempty"`
	RogueBuffInfo    *RogueBuffInfo            `protobuf:"bytes,14,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
}

func (x *RogueCurrentInfo) Reset() {
	*x = RogueCurrentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueCurrentInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueCurrentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueCurrentInfo) ProtoMessage() {}

func (x *RogueCurrentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueCurrentInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueCurrentInfo.ProtoReflect.Descriptor instead.
func (*RogueCurrentInfo) Descriptor() ([]byte, []int) {
	return file_RogueCurrentInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueCurrentInfo) GetRogueAeonInfo() *RogueAeon {
	if x != nil {
		return x.RogueAeonInfo
	}
	return nil
}

func (x *RogueCurrentInfo) GetRogueMiracleInfo() *RogueMiracleInfo {
	if x != nil {
		return x.RogueMiracleInfo
	}
	return nil
}

func (x *RogueCurrentInfo) GetRogueAvatarInfo() *RogueAvatarInfo {
	if x != nil {
		return x.RogueAvatarInfo
	}
	return nil
}

func (x *RogueCurrentInfo) GetStatus() RogueStatus {
	if x != nil {
		return x.Status
	}
	return RogueStatus_ROGUE_STATUS_NONE
}

func (x *RogueCurrentInfo) GetRoomMap() *RogueMapInfo {
	if x != nil {
		return x.RoomMap
	}
	return nil
}

func (x *RogueCurrentInfo) GetPendingAction() *RogueCommonPendingAction {
	if x != nil {
		return x.PendingAction
	}
	return nil
}

func (x *RogueCurrentInfo) GetRogueVirtualItem() *RogueVirtualItemInfo {
	if x != nil {
		return x.RogueVirtualItem
	}
	return nil
}

func (x *RogueCurrentInfo) GetRogueBuffInfo() *RogueBuffInfo {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

var File_RogueCurrentInfo_proto protoreflect.FileDescriptor

var file_RogueCurrentInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41,
	0x65, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42,
	0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd4, 0x03, 0x0a, 0x10, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x61,
	0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x12, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x5f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x4d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x11, 0x72, 0x6f,
	0x67, 0x75, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28,
	0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x4d, 0x61, 0x70, 0x12, 0x40, 0x0a, 0x0e, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x12, 0x72, 0x6f,
	0x67, 0x75, 0x65, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x36, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x42,
	0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueCurrentInfo_proto_rawDescOnce sync.Once
	file_RogueCurrentInfo_proto_rawDescData = file_RogueCurrentInfo_proto_rawDesc
)

func file_RogueCurrentInfo_proto_rawDescGZIP() []byte {
	file_RogueCurrentInfo_proto_rawDescOnce.Do(func() {
		file_RogueCurrentInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueCurrentInfo_proto_rawDescData)
	})
	return file_RogueCurrentInfo_proto_rawDescData
}

var file_RogueCurrentInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueCurrentInfo_proto_goTypes = []interface{}{
	(*RogueCurrentInfo)(nil),         // 0: RogueCurrentInfo
	(*RogueAeon)(nil),                // 1: RogueAeon
	(*RogueMiracleInfo)(nil),         // 2: RogueMiracleInfo
	(*RogueAvatarInfo)(nil),          // 3: RogueAvatarInfo
	(RogueStatus)(0),                 // 4: RogueStatus
	(*RogueMapInfo)(nil),             // 5: RogueMapInfo
	(*RogueCommonPendingAction)(nil), // 6: RogueCommonPendingAction
	(*RogueVirtualItemInfo)(nil),     // 7: RogueVirtualItemInfo
	(*RogueBuffInfo)(nil),            // 8: RogueBuffInfo
}
var file_RogueCurrentInfo_proto_depIdxs = []int32{
	1, // 0: RogueCurrentInfo.rogue_aeon_info:type_name -> RogueAeon
	2, // 1: RogueCurrentInfo.rogue_miracle_info:type_name -> RogueMiracleInfo
	3, // 2: RogueCurrentInfo.rogue_avatar_info:type_name -> RogueAvatarInfo
	4, // 3: RogueCurrentInfo.status:type_name -> RogueStatus
	5, // 4: RogueCurrentInfo.room_map:type_name -> RogueMapInfo
	6, // 5: RogueCurrentInfo.pending_action:type_name -> RogueCommonPendingAction
	7, // 6: RogueCurrentInfo.rogue_virtual_item:type_name -> RogueVirtualItemInfo
	8, // 7: RogueCurrentInfo.rogue_buff_info:type_name -> RogueBuffInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_RogueCurrentInfo_proto_init() }
func file_RogueCurrentInfo_proto_init() {
	if File_RogueCurrentInfo_proto != nil {
		return
	}
	file_RogueAeon_proto_init()
	file_RogueAvatarInfo_proto_init()
	file_RogueMiracleInfo_proto_init()
	file_RogueBuffInfo_proto_init()
	file_RogueCommonPendingAction_proto_init()
	file_RogueVirtualItemInfo_proto_init()
	file_RogueStatus_proto_init()
	file_RogueMapInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueCurrentInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueCurrentInfo); i {
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
			RawDescriptor: file_RogueCurrentInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueCurrentInfo_proto_goTypes,
		DependencyIndexes: file_RogueCurrentInfo_proto_depIdxs,
		MessageInfos:      file_RogueCurrentInfo_proto_msgTypes,
	}.Build()
	File_RogueCurrentInfo_proto = out.File
	file_RogueCurrentInfo_proto_rawDesc = nil
	file_RogueCurrentInfo_proto_goTypes = nil
	file_RogueCurrentInfo_proto_depIdxs = nil
}
