// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetChessRogueBuffEnhanceInfoScRsp.proto

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

type GetChessRogueBuffEnhanceInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode         uint32                     `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BuffEnhanceInfo *ChessRogueBuffEnhanceInfo `protobuf:"bytes,9,opt,name=buff_enhance_info,json=buffEnhanceInfo,proto3" json:"buff_enhance_info,omitempty"`
}

func (x *GetChessRogueBuffEnhanceInfoScRsp) Reset() {
	*x = GetChessRogueBuffEnhanceInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetChessRogueBuffEnhanceInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChessRogueBuffEnhanceInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChessRogueBuffEnhanceInfoScRsp) ProtoMessage() {}

func (x *GetChessRogueBuffEnhanceInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetChessRogueBuffEnhanceInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChessRogueBuffEnhanceInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetChessRogueBuffEnhanceInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetChessRogueBuffEnhanceInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetChessRogueBuffEnhanceInfoScRsp) GetBuffEnhanceInfo() *ChessRogueBuffEnhanceInfo {
	if x != nil {
		return x.BuffEnhanceInfo
	}
	return nil
}

var File_GetChessRogueBuffEnhanceInfoScRsp_proto protoreflect.FileDescriptor

var file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x27, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42,
	0x75, 0x66, 0x66, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x21, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66,
	0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x62, 0x75,
	0x66, 0x66, 0x5f, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0f, 0x62, 0x75, 0x66, 0x66, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDescOnce sync.Once
	file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDescData = file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDesc
)

func file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDescData)
	})
	return file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDescData
}

var file_GetChessRogueBuffEnhanceInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetChessRogueBuffEnhanceInfoScRsp_proto_goTypes = []interface{}{
	(*GetChessRogueBuffEnhanceInfoScRsp)(nil), // 0: GetChessRogueBuffEnhanceInfoScRsp
	(*ChessRogueBuffEnhanceInfo)(nil),         // 1: ChessRogueBuffEnhanceInfo
}
var file_GetChessRogueBuffEnhanceInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetChessRogueBuffEnhanceInfoScRsp.buff_enhance_info:type_name -> ChessRogueBuffEnhanceInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetChessRogueBuffEnhanceInfoScRsp_proto_init() }
func file_GetChessRogueBuffEnhanceInfoScRsp_proto_init() {
	if File_GetChessRogueBuffEnhanceInfoScRsp_proto != nil {
		return
	}
	file_ChessRogueBuffEnhanceInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetChessRogueBuffEnhanceInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChessRogueBuffEnhanceInfoScRsp); i {
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
			RawDescriptor: file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetChessRogueBuffEnhanceInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetChessRogueBuffEnhanceInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetChessRogueBuffEnhanceInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetChessRogueBuffEnhanceInfoScRsp_proto = out.File
	file_GetChessRogueBuffEnhanceInfoScRsp_proto_rawDesc = nil
	file_GetChessRogueBuffEnhanceInfoScRsp_proto_goTypes = nil
	file_GetChessRogueBuffEnhanceInfoScRsp_proto_depIdxs = nil
}