// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PropExtraInfo.proto

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

type PropExtraInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Info:
	//
	//	*PropExtraInfo_RogueInfo
	Info isPropExtraInfo_Info `protobuf_oneof:"info"`
}

func (x *PropExtraInfo) Reset() {
	*x = PropExtraInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PropExtraInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropExtraInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropExtraInfo) ProtoMessage() {}

func (x *PropExtraInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PropExtraInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropExtraInfo.ProtoReflect.Descriptor instead.
func (*PropExtraInfo) Descriptor() ([]byte, []int) {
	return file_PropExtraInfo_proto_rawDescGZIP(), []int{0}
}

func (m *PropExtraInfo) GetInfo() isPropExtraInfo_Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (x *PropExtraInfo) GetRogueInfo() *PropRogueInfo {
	if x, ok := x.GetInfo().(*PropExtraInfo_RogueInfo); ok {
		return x.RogueInfo
	}
	return nil
}

type isPropExtraInfo_Info interface {
	isPropExtraInfo_Info()
}

type PropExtraInfo_RogueInfo struct {
	RogueInfo *PropRogueInfo `protobuf:"bytes,11,opt,name=rogue_info,json=rogueInfo,proto3,oneof"`
}

func (*PropExtraInfo_RogueInfo) isPropExtraInfo_Info() {}

var File_PropExtraInfo_proto protoreflect.FileDescriptor

var file_PropExtraInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0d, 0x50, 0x72,
	0x6f, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x0a, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x09, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PropExtraInfo_proto_rawDescOnce sync.Once
	file_PropExtraInfo_proto_rawDescData = file_PropExtraInfo_proto_rawDesc
)

func file_PropExtraInfo_proto_rawDescGZIP() []byte {
	file_PropExtraInfo_proto_rawDescOnce.Do(func() {
		file_PropExtraInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PropExtraInfo_proto_rawDescData)
	})
	return file_PropExtraInfo_proto_rawDescData
}

var file_PropExtraInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PropExtraInfo_proto_goTypes = []interface{}{
	(*PropExtraInfo)(nil), // 0: PropExtraInfo
	(*PropRogueInfo)(nil), // 1: PropRogueInfo
}
var file_PropExtraInfo_proto_depIdxs = []int32{
	1, // 0: PropExtraInfo.rogue_info:type_name -> PropRogueInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PropExtraInfo_proto_init() }
func file_PropExtraInfo_proto_init() {
	if File_PropExtraInfo_proto != nil {
		return
	}
	file_PropRogueInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PropExtraInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropExtraInfo); i {
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
	file_PropExtraInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PropExtraInfo_RogueInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PropExtraInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PropExtraInfo_proto_goTypes,
		DependencyIndexes: file_PropExtraInfo_proto_depIdxs,
		MessageInfos:      file_PropExtraInfo_proto_msgTypes,
	}.Build()
	File_PropExtraInfo_proto = out.File
	file_PropExtraInfo_proto_rawDesc = nil
	file_PropExtraInfo_proto_goTypes = nil
	file_PropExtraInfo_proto_depIdxs = nil
}
