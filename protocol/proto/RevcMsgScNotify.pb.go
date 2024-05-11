// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RevcMsgScNotify.proto

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

type RevcMsgScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToUid    uint32   `protobuf:"varint,3,opt,name=to_uid,json=toUid,proto3" json:"to_uid,omitempty"`
	Emote    uint32   `protobuf:"varint,2,opt,name=emote,proto3" json:"emote,omitempty"`
	MsgType  MsgType  `protobuf:"varint,15,opt,name=msg_type,json=msgType,proto3,enum=MsgType" json:"msg_type,omitempty"`
	FromUid  uint32   `protobuf:"varint,11,opt,name=from_uid,json=fromUid,proto3" json:"from_uid,omitempty"`
	Text     string   `protobuf:"bytes,9,opt,name=text,proto3" json:"text,omitempty"`
	ChatType ChatType `protobuf:"varint,6,opt,name=chat_type,json=chatType,proto3,enum=ChatType" json:"chat_type,omitempty"`
}

func (x *RevcMsgScNotify) Reset() {
	*x = RevcMsgScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RevcMsgScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevcMsgScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevcMsgScNotify) ProtoMessage() {}

func (x *RevcMsgScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RevcMsgScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevcMsgScNotify.ProtoReflect.Descriptor instead.
func (*RevcMsgScNotify) Descriptor() ([]byte, []int) {
	return file_RevcMsgScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RevcMsgScNotify) GetToUid() uint32 {
	if x != nil {
		return x.ToUid
	}
	return 0
}

func (x *RevcMsgScNotify) GetEmote() uint32 {
	if x != nil {
		return x.Emote
	}
	return 0
}

func (x *RevcMsgScNotify) GetMsgType() MsgType {
	if x != nil {
		return x.MsgType
	}
	return MsgType_MSG_TYPE_NONE
}

func (x *RevcMsgScNotify) GetFromUid() uint32 {
	if x != nil {
		return x.FromUid
	}
	return 0
}

func (x *RevcMsgScNotify) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *RevcMsgScNotify) GetChatType() ChatType {
	if x != nil {
		return x.ChatType
	}
	return ChatType_CHAT_TYPE_NONE
}

var File_RevcMsgScNotify_proto protoreflect.FileDescriptor

var file_RevcMsgScNotify_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x65, 0x76, 0x63, 0x4d, 0x73, 0x67, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x76, 0x63, 0x4d,
	0x73, 0x67, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x4d, 0x73, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x66, 0x72, 0x6f, 0x6d, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RevcMsgScNotify_proto_rawDescOnce sync.Once
	file_RevcMsgScNotify_proto_rawDescData = file_RevcMsgScNotify_proto_rawDesc
)

func file_RevcMsgScNotify_proto_rawDescGZIP() []byte {
	file_RevcMsgScNotify_proto_rawDescOnce.Do(func() {
		file_RevcMsgScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RevcMsgScNotify_proto_rawDescData)
	})
	return file_RevcMsgScNotify_proto_rawDescData
}

var file_RevcMsgScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RevcMsgScNotify_proto_goTypes = []interface{}{
	(*RevcMsgScNotify)(nil), // 0: RevcMsgScNotify
	(MsgType)(0),            // 1: MsgType
	(ChatType)(0),           // 2: ChatType
}
var file_RevcMsgScNotify_proto_depIdxs = []int32{
	1, // 0: RevcMsgScNotify.msg_type:type_name -> MsgType
	2, // 1: RevcMsgScNotify.chat_type:type_name -> ChatType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RevcMsgScNotify_proto_init() }
func file_RevcMsgScNotify_proto_init() {
	if File_RevcMsgScNotify_proto != nil {
		return
	}
	file_ChatType_proto_init()
	file_MsgType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RevcMsgScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevcMsgScNotify); i {
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
			RawDescriptor: file_RevcMsgScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RevcMsgScNotify_proto_goTypes,
		DependencyIndexes: file_RevcMsgScNotify_proto_depIdxs,
		MessageInfos:      file_RevcMsgScNotify_proto_msgTypes,
	}.Build()
	File_RevcMsgScNotify_proto = out.File
	file_RevcMsgScNotify_proto_rawDesc = nil
	file_RevcMsgScNotify_proto_goTypes = nil
	file_RevcMsgScNotify_proto_depIdxs = nil
}
