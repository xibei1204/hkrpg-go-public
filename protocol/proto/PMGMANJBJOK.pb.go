// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PMGMANJBJOK.proto

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

type PMGMANJBJOK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EJMJIKCDIMM uint32 `protobuf:"varint,6,opt,name=EJMJIKCDIMM,proto3" json:"EJMJIKCDIMM,omitempty"`
	JFLFDIJPNEO uint32 `protobuf:"varint,14,opt,name=JFLFDIJPNEO,proto3" json:"JFLFDIJPNEO,omitempty"`
	ILADMEOHHEC uint32 `protobuf:"varint,3,opt,name=ILADMEOHHEC,proto3" json:"ILADMEOHHEC,omitempty"`
	DCIOBLHLICO uint32 `protobuf:"varint,1,opt,name=DCIOBLHLICO,proto3" json:"DCIOBLHLICO,omitempty"`
	KLLEONMNLDI uint32 `protobuf:"varint,15,opt,name=KLLEONMNLDI,proto3" json:"KLLEONMNLDI,omitempty"`
}

func (x *PMGMANJBJOK) Reset() {
	*x = PMGMANJBJOK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PMGMANJBJOK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PMGMANJBJOK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PMGMANJBJOK) ProtoMessage() {}

func (x *PMGMANJBJOK) ProtoReflect() protoreflect.Message {
	mi := &file_PMGMANJBJOK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PMGMANJBJOK.ProtoReflect.Descriptor instead.
func (*PMGMANJBJOK) Descriptor() ([]byte, []int) {
	return file_PMGMANJBJOK_proto_rawDescGZIP(), []int{0}
}

func (x *PMGMANJBJOK) GetEJMJIKCDIMM() uint32 {
	if x != nil {
		return x.EJMJIKCDIMM
	}
	return 0
}

func (x *PMGMANJBJOK) GetJFLFDIJPNEO() uint32 {
	if x != nil {
		return x.JFLFDIJPNEO
	}
	return 0
}

func (x *PMGMANJBJOK) GetILADMEOHHEC() uint32 {
	if x != nil {
		return x.ILADMEOHHEC
	}
	return 0
}

func (x *PMGMANJBJOK) GetDCIOBLHLICO() uint32 {
	if x != nil {
		return x.DCIOBLHLICO
	}
	return 0
}

func (x *PMGMANJBJOK) GetKLLEONMNLDI() uint32 {
	if x != nil {
		return x.KLLEONMNLDI
	}
	return 0
}

var File_PMGMANJBJOK_proto protoreflect.FileDescriptor

var file_PMGMANJBJOK_proto_rawDesc = []byte{
	0x0a, 0x11, 0x50, 0x4d, 0x47, 0x4d, 0x41, 0x4e, 0x4a, 0x42, 0x4a, 0x4f, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x50, 0x4d, 0x47, 0x4d, 0x41, 0x4e, 0x4a, 0x42,
	0x4a, 0x4f, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x4a, 0x4d, 0x4a, 0x49, 0x4b, 0x43, 0x44, 0x49,
	0x4d, 0x4d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x4a, 0x4d, 0x4a, 0x49, 0x4b,
	0x43, 0x44, 0x49, 0x4d, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x46, 0x4c, 0x46, 0x44, 0x49, 0x4a,
	0x50, 0x4e, 0x45, 0x4f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x46, 0x4c, 0x46,
	0x44, 0x49, 0x4a, 0x50, 0x4e, 0x45, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4c, 0x41, 0x44, 0x4d,
	0x45, 0x4f, 0x48, 0x48, 0x45, 0x43, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4c,
	0x41, 0x44, 0x4d, 0x45, 0x4f, 0x48, 0x48, 0x45, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x43, 0x49,
	0x4f, 0x42, 0x4c, 0x48, 0x4c, 0x49, 0x43, 0x4f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x44, 0x43, 0x49, 0x4f, 0x42, 0x4c, 0x48, 0x4c, 0x49, 0x43, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4b,
	0x4c, 0x4c, 0x45, 0x4f, 0x4e, 0x4d, 0x4e, 0x4c, 0x44, 0x49, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4b, 0x4c, 0x4c, 0x45, 0x4f, 0x4e, 0x4d, 0x4e, 0x4c, 0x44, 0x49, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_PMGMANJBJOK_proto_rawDescOnce sync.Once
	file_PMGMANJBJOK_proto_rawDescData = file_PMGMANJBJOK_proto_rawDesc
)

func file_PMGMANJBJOK_proto_rawDescGZIP() []byte {
	file_PMGMANJBJOK_proto_rawDescOnce.Do(func() {
		file_PMGMANJBJOK_proto_rawDescData = protoimpl.X.CompressGZIP(file_PMGMANJBJOK_proto_rawDescData)
	})
	return file_PMGMANJBJOK_proto_rawDescData
}

var file_PMGMANJBJOK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PMGMANJBJOK_proto_goTypes = []interface{}{
	(*PMGMANJBJOK)(nil), // 0: PMGMANJBJOK
}
var file_PMGMANJBJOK_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PMGMANJBJOK_proto_init() }
func file_PMGMANJBJOK_proto_init() {
	if File_PMGMANJBJOK_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PMGMANJBJOK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PMGMANJBJOK); i {
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
			RawDescriptor: file_PMGMANJBJOK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PMGMANJBJOK_proto_goTypes,
		DependencyIndexes: file_PMGMANJBJOK_proto_depIdxs,
		MessageInfos:      file_PMGMANJBJOK_proto_msgTypes,
	}.Build()
	File_PMGMANJBJOK_proto = out.File
	file_PMGMANJBJOK_proto_rawDesc = nil
	file_PMGMANJBJOK_proto_goTypes = nil
	file_PMGMANJBJOK_proto_depIdxs = nil
}