// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerLoginScRsp.proto

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

type PlayerLoginScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode           uint32           `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ServerTimestampMs uint64           `protobuf:"varint,8,opt,name=server_timestamp_ms,json=serverTimestampMs,proto3" json:"server_timestamp_ms,omitempty"`
	LoginRandom       uint64           `protobuf:"varint,15,opt,name=login_random,json=loginRandom,proto3" json:"login_random,omitempty"`
	BasicInfo         *PlayerBasicInfo `protobuf:"bytes,7,opt,name=basic_info,json=basicInfo,proto3" json:"basic_info,omitempty"`
	Stamina           uint32           `protobuf:"varint,2,opt,name=stamina,proto3" json:"stamina,omitempty"`
	CurTimezone       int32            `protobuf:"zigzag32,13,opt,name=cur_timezone,json=curTimezone,proto3" json:"cur_timezone,omitempty"`
}

func (x *PlayerLoginScRsp) Reset() {
	*x = PlayerLoginScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerLoginScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginScRsp) ProtoMessage() {}

func (x *PlayerLoginScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerLoginScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginScRsp.ProtoReflect.Descriptor instead.
func (*PlayerLoginScRsp) Descriptor() ([]byte, []int) {
	return file_PlayerLoginScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLoginScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PlayerLoginScRsp) GetServerTimestampMs() uint64 {
	if x != nil {
		return x.ServerTimestampMs
	}
	return 0
}

func (x *PlayerLoginScRsp) GetLoginRandom() uint64 {
	if x != nil {
		return x.LoginRandom
	}
	return 0
}

func (x *PlayerLoginScRsp) GetBasicInfo() *PlayerBasicInfo {
	if x != nil {
		return x.BasicInfo
	}
	return nil
}

func (x *PlayerLoginScRsp) GetStamina() uint32 {
	if x != nil {
		return x.Stamina
	}
	return 0
}

func (x *PlayerLoginScRsp) GetCurTimezone() int32 {
	if x != nil {
		return x.CurTimezone
	}
	return 0
}

var File_PlayerLoginScRsp_proto protoreflect.FileDescriptor

var file_PlayerLoginScRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xed, 0x01, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x12, 0x2f, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x62, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x75, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_PlayerLoginScRsp_proto_rawDescOnce sync.Once
	file_PlayerLoginScRsp_proto_rawDescData = file_PlayerLoginScRsp_proto_rawDesc
)

func file_PlayerLoginScRsp_proto_rawDescGZIP() []byte {
	file_PlayerLoginScRsp_proto_rawDescOnce.Do(func() {
		file_PlayerLoginScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerLoginScRsp_proto_rawDescData)
	})
	return file_PlayerLoginScRsp_proto_rawDescData
}

var file_PlayerLoginScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerLoginScRsp_proto_goTypes = []interface{}{
	(*PlayerLoginScRsp)(nil), // 0: PlayerLoginScRsp
	(*PlayerBasicInfo)(nil),  // 1: PlayerBasicInfo
}
var file_PlayerLoginScRsp_proto_depIdxs = []int32{
	1, // 0: PlayerLoginScRsp.basic_info:type_name -> PlayerBasicInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerLoginScRsp_proto_init() }
func file_PlayerLoginScRsp_proto_init() {
	if File_PlayerLoginScRsp_proto != nil {
		return
	}
	file_PlayerBasicInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerLoginScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginScRsp); i {
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
			RawDescriptor: file_PlayerLoginScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerLoginScRsp_proto_goTypes,
		DependencyIndexes: file_PlayerLoginScRsp_proto_depIdxs,
		MessageInfos:      file_PlayerLoginScRsp_proto_msgTypes,
	}.Build()
	File_PlayerLoginScRsp_proto = out.File
	file_PlayerLoginScRsp_proto_rawDesc = nil
	file_PlayerLoginScRsp_proto_goTypes = nil
	file_PlayerLoginScRsp_proto_depIdxs = nil
}
