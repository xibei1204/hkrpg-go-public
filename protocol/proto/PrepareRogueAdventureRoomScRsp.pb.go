// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PrepareRogueAdventureRoomScRsp.proto

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

type PrepareRogueAdventureRoomScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OHPALFLEFKI *GFBHIMKAMEJ `protobuf:"bytes,9,opt,name=OHPALFLEFKI,proto3" json:"OHPALFLEFKI,omitempty"`
	Retcode     uint32       `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *PrepareRogueAdventureRoomScRsp) Reset() {
	*x = PrepareRogueAdventureRoomScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PrepareRogueAdventureRoomScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareRogueAdventureRoomScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareRogueAdventureRoomScRsp) ProtoMessage() {}

func (x *PrepareRogueAdventureRoomScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PrepareRogueAdventureRoomScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareRogueAdventureRoomScRsp.ProtoReflect.Descriptor instead.
func (*PrepareRogueAdventureRoomScRsp) Descriptor() ([]byte, []int) {
	return file_PrepareRogueAdventureRoomScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PrepareRogueAdventureRoomScRsp) GetOHPALFLEFKI() *GFBHIMKAMEJ {
	if x != nil {
		return x.OHPALFLEFKI
	}
	return nil
}

func (x *PrepareRogueAdventureRoomScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_PrepareRogueAdventureRoomScRsp_proto protoreflect.FileDescriptor

var file_PrepareRogueAdventureRoomScRsp_proto_rawDesc = []byte{
	0x0a, 0x24, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x64,
	0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x46, 0x42, 0x48, 0x49, 0x4d, 0x4b, 0x41,
	0x4d, 0x45, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x1e, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x4f,
	0x48, 0x50, 0x41, 0x4c, 0x46, 0x4c, 0x45, 0x46, 0x4b, 0x49, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x47, 0x46, 0x42, 0x48, 0x49, 0x4d, 0x4b, 0x41, 0x4d, 0x45, 0x4a, 0x52, 0x0b,
	0x4f, 0x48, 0x50, 0x41, 0x4c, 0x46, 0x4c, 0x45, 0x46, 0x4b, 0x49, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PrepareRogueAdventureRoomScRsp_proto_rawDescOnce sync.Once
	file_PrepareRogueAdventureRoomScRsp_proto_rawDescData = file_PrepareRogueAdventureRoomScRsp_proto_rawDesc
)

func file_PrepareRogueAdventureRoomScRsp_proto_rawDescGZIP() []byte {
	file_PrepareRogueAdventureRoomScRsp_proto_rawDescOnce.Do(func() {
		file_PrepareRogueAdventureRoomScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PrepareRogueAdventureRoomScRsp_proto_rawDescData)
	})
	return file_PrepareRogueAdventureRoomScRsp_proto_rawDescData
}

var file_PrepareRogueAdventureRoomScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PrepareRogueAdventureRoomScRsp_proto_goTypes = []interface{}{
	(*PrepareRogueAdventureRoomScRsp)(nil), // 0: PrepareRogueAdventureRoomScRsp
	(*GFBHIMKAMEJ)(nil),                    // 1: GFBHIMKAMEJ
}
var file_PrepareRogueAdventureRoomScRsp_proto_depIdxs = []int32{
	1, // 0: PrepareRogueAdventureRoomScRsp.OHPALFLEFKI:type_name -> GFBHIMKAMEJ
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PrepareRogueAdventureRoomScRsp_proto_init() }
func file_PrepareRogueAdventureRoomScRsp_proto_init() {
	if File_PrepareRogueAdventureRoomScRsp_proto != nil {
		return
	}
	file_GFBHIMKAMEJ_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PrepareRogueAdventureRoomScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareRogueAdventureRoomScRsp); i {
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
			RawDescriptor: file_PrepareRogueAdventureRoomScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PrepareRogueAdventureRoomScRsp_proto_goTypes,
		DependencyIndexes: file_PrepareRogueAdventureRoomScRsp_proto_depIdxs,
		MessageInfos:      file_PrepareRogueAdventureRoomScRsp_proto_msgTypes,
	}.Build()
	File_PrepareRogueAdventureRoomScRsp_proto = out.File
	file_PrepareRogueAdventureRoomScRsp_proto_rawDesc = nil
	file_PrepareRogueAdventureRoomScRsp_proto_goTypes = nil
	file_PrepareRogueAdventureRoomScRsp_proto_depIdxs = nil
}