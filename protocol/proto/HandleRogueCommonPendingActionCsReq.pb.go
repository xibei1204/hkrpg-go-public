// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HandleRogueCommonPendingActionCsReq.proto

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

type HandleRogueCommonPendingActionCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PendingAction:
	//
	//	*HandleRogueCommonPendingActionCsReq_BuffSelectResult
	//	*HandleRogueCommonPendingActionCsReq_RollBuff
	//	*HandleRogueCommonPendingActionCsReq_MiracleSelectResult
	//	*HandleRogueCommonPendingActionCsReq_BonusSelectResult
	PendingAction isHandleRogueCommonPendingActionCsReq_PendingAction `protobuf_oneof:"pending_action"`
}

func (x *HandleRogueCommonPendingActionCsReq) Reset() {
	*x = HandleRogueCommonPendingActionCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleRogueCommonPendingActionCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleRogueCommonPendingActionCsReq) ProtoMessage() {}

func (x *HandleRogueCommonPendingActionCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleRogueCommonPendingActionCsReq.ProtoReflect.Descriptor instead.
func (*HandleRogueCommonPendingActionCsReq) Descriptor() ([]byte, []int) {
	return file_HandleRogueCommonPendingActionCsReq_proto_rawDescGZIP(), []int{0}
}

func (m *HandleRogueCommonPendingActionCsReq) GetPendingAction() isHandleRogueCommonPendingActionCsReq_PendingAction {
	if m != nil {
		return m.PendingAction
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetBuffSelectResult() *RogueCommonBuffSelectResult {
	if x, ok := x.GetPendingAction().(*HandleRogueCommonPendingActionCsReq_BuffSelectResult); ok {
		return x.BuffSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetRollBuff() *RogueBuffRollInfo {
	if x, ok := x.GetPendingAction().(*HandleRogueCommonPendingActionCsReq_RollBuff); ok {
		return x.RollBuff
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetMiracleSelectResult() *RogueMiracleSelectResult {
	if x, ok := x.GetPendingAction().(*HandleRogueCommonPendingActionCsReq_MiracleSelectResult); ok {
		return x.MiracleSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetBonusSelectResult() *RogueBonusSelectResult {
	if x, ok := x.GetPendingAction().(*HandleRogueCommonPendingActionCsReq_BonusSelectResult); ok {
		return x.BonusSelectResult
	}
	return nil
}

type isHandleRogueCommonPendingActionCsReq_PendingAction interface {
	isHandleRogueCommonPendingActionCsReq_PendingAction()
}

type HandleRogueCommonPendingActionCsReq_BuffSelectResult struct {
	BuffSelectResult *RogueCommonBuffSelectResult `protobuf:"bytes,1586,opt,name=buff_select_result,json=buffSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_RollBuff struct {
	RollBuff *RogueBuffRollInfo `protobuf:"bytes,308,opt,name=roll_buff,json=rollBuff,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_MiracleSelectResult struct {
	MiracleSelectResult *RogueMiracleSelectResult `protobuf:"bytes,1209,opt,name=miracle_select_result,json=miracleSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_BonusSelectResult struct {
	BonusSelectResult *RogueBonusSelectResult `protobuf:"bytes,1156,opt,name=bonus_select_result,json=bonusSelectResult,proto3,oneof"`
}

func (*HandleRogueCommonPendingActionCsReq_BuffSelectResult) isHandleRogueCommonPendingActionCsReq_PendingAction() {
}

func (*HandleRogueCommonPendingActionCsReq_RollBuff) isHandleRogueCommonPendingActionCsReq_PendingAction() {
}

func (*HandleRogueCommonPendingActionCsReq_MiracleSelectResult) isHandleRogueCommonPendingActionCsReq_PendingAction() {
}

func (*HandleRogueCommonPendingActionCsReq_BonusSelectResult) isHandleRogueCommonPendingActionCsReq_PendingAction() {
}

var File_HandleRogueCommonPendingActionCsReq_proto protoreflect.FileDescriptor

var file_HandleRogueCommonPendingActionCsReq_proto_rawDesc = []byte{
	0x0a, 0x29, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x6f, 0x6c, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f,
	0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x23, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x4d, 0x0a,
	0x12, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0xb2, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x10, 0x62, 0x75, 0x66, 0x66,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x09,
	0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x18, 0xb4, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x6f, 0x6c, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x6c, 0x42, 0x75, 0x66, 0x66,
	0x12, 0x50, 0x0a, 0x15, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xb9, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x13, 0x6d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x4a, 0x0a, 0x13, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x84, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x11, 0x62, 0x6f, 0x6e,
	0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x10,
	0x0a, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HandleRogueCommonPendingActionCsReq_proto_rawDescOnce sync.Once
	file_HandleRogueCommonPendingActionCsReq_proto_rawDescData = file_HandleRogueCommonPendingActionCsReq_proto_rawDesc
)

func file_HandleRogueCommonPendingActionCsReq_proto_rawDescGZIP() []byte {
	file_HandleRogueCommonPendingActionCsReq_proto_rawDescOnce.Do(func() {
		file_HandleRogueCommonPendingActionCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HandleRogueCommonPendingActionCsReq_proto_rawDescData)
	})
	return file_HandleRogueCommonPendingActionCsReq_proto_rawDescData
}

var file_HandleRogueCommonPendingActionCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HandleRogueCommonPendingActionCsReq_proto_goTypes = []interface{}{
	(*HandleRogueCommonPendingActionCsReq)(nil), // 0: HandleRogueCommonPendingActionCsReq
	(*RogueCommonBuffSelectResult)(nil),         // 1: RogueCommonBuffSelectResult
	(*RogueBuffRollInfo)(nil),                   // 2: RogueBuffRollInfo
	(*RogueMiracleSelectResult)(nil),            // 3: RogueMiracleSelectResult
	(*RogueBonusSelectResult)(nil),              // 4: RogueBonusSelectResult
}
var file_HandleRogueCommonPendingActionCsReq_proto_depIdxs = []int32{
	1, // 0: HandleRogueCommonPendingActionCsReq.buff_select_result:type_name -> RogueCommonBuffSelectResult
	2, // 1: HandleRogueCommonPendingActionCsReq.roll_buff:type_name -> RogueBuffRollInfo
	3, // 2: HandleRogueCommonPendingActionCsReq.miracle_select_result:type_name -> RogueMiracleSelectResult
	4, // 3: HandleRogueCommonPendingActionCsReq.bonus_select_result:type_name -> RogueBonusSelectResult
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_HandleRogueCommonPendingActionCsReq_proto_init() }
func file_HandleRogueCommonPendingActionCsReq_proto_init() {
	if File_HandleRogueCommonPendingActionCsReq_proto != nil {
		return
	}
	file_RogueCommonBuffSelectResult_proto_init()
	file_RogueBuffRollInfo_proto_init()
	file_RogueMiracleSelectResult_proto_init()
	file_RogueBonusSelectResult_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleRogueCommonPendingActionCsReq); i {
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
	file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HandleRogueCommonPendingActionCsReq_BuffSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_RollBuff)(nil),
		(*HandleRogueCommonPendingActionCsReq_MiracleSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_BonusSelectResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HandleRogueCommonPendingActionCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HandleRogueCommonPendingActionCsReq_proto_goTypes,
		DependencyIndexes: file_HandleRogueCommonPendingActionCsReq_proto_depIdxs,
		MessageInfos:      file_HandleRogueCommonPendingActionCsReq_proto_msgTypes,
	}.Build()
	File_HandleRogueCommonPendingActionCsReq_proto = out.File
	file_HandleRogueCommonPendingActionCsReq_proto_rawDesc = nil
	file_HandleRogueCommonPendingActionCsReq_proto_goTypes = nil
	file_HandleRogueCommonPendingActionCsReq_proto_depIdxs = nil
}
