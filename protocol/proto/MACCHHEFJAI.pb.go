// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MACCHHEFJAI.proto

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

type MACCHHEFJAI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AJCGHMHMCAO []uint32                  `protobuf:"varint,14,rep,packed,name=AJCGHMHMCAO,proto3" json:"AJCGHMHMCAO,omitempty"`
	NCFICCIEHEH MonopolySocialEventStatus `protobuf:"varint,2,opt,name=NCFICCIEHEH,proto3,enum=MonopolySocialEventStatus" json:"NCFICCIEHEH,omitempty"`
}

func (x *MACCHHEFJAI) Reset() {
	*x = MACCHHEFJAI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MACCHHEFJAI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MACCHHEFJAI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MACCHHEFJAI) ProtoMessage() {}

func (x *MACCHHEFJAI) ProtoReflect() protoreflect.Message {
	mi := &file_MACCHHEFJAI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MACCHHEFJAI.ProtoReflect.Descriptor instead.
func (*MACCHHEFJAI) Descriptor() ([]byte, []int) {
	return file_MACCHHEFJAI_proto_rawDescGZIP(), []int{0}
}

func (x *MACCHHEFJAI) GetAJCGHMHMCAO() []uint32 {
	if x != nil {
		return x.AJCGHMHMCAO
	}
	return nil
}

func (x *MACCHHEFJAI) GetNCFICCIEHEH() MonopolySocialEventStatus {
	if x != nil {
		return x.NCFICCIEHEH
	}
	return MonopolySocialEventStatus_MONOPOLY_SOCIAL_EVENT_STATUS_NONE
}

var File_MACCHHEFJAI_proto protoreflect.FileDescriptor

var file_MACCHHEFJAI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x41, 0x43, 0x43, 0x48, 0x48, 0x45, 0x46, 0x4a, 0x41, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x53, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0b, 0x4d, 0x41, 0x43, 0x43, 0x48, 0x48, 0x45, 0x46,
	0x4a, 0x41, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x4a, 0x43, 0x47, 0x48, 0x4d, 0x48, 0x4d, 0x43,
	0x41, 0x4f, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4a, 0x43, 0x47, 0x48, 0x4d,
	0x48, 0x4d, 0x43, 0x41, 0x4f, 0x12, 0x3c, 0x0a, 0x0b, 0x4e, 0x43, 0x46, 0x49, 0x43, 0x43, 0x49,
	0x45, 0x48, 0x45, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x4d, 0x6f, 0x6e,
	0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x4e, 0x43, 0x46, 0x49, 0x43, 0x43, 0x49, 0x45,
	0x48, 0x45, 0x48, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MACCHHEFJAI_proto_rawDescOnce sync.Once
	file_MACCHHEFJAI_proto_rawDescData = file_MACCHHEFJAI_proto_rawDesc
)

func file_MACCHHEFJAI_proto_rawDescGZIP() []byte {
	file_MACCHHEFJAI_proto_rawDescOnce.Do(func() {
		file_MACCHHEFJAI_proto_rawDescData = protoimpl.X.CompressGZIP(file_MACCHHEFJAI_proto_rawDescData)
	})
	return file_MACCHHEFJAI_proto_rawDescData
}

var file_MACCHHEFJAI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MACCHHEFJAI_proto_goTypes = []interface{}{
	(*MACCHHEFJAI)(nil),            // 0: MACCHHEFJAI
	(MonopolySocialEventStatus)(0), // 1: MonopolySocialEventStatus
}
var file_MACCHHEFJAI_proto_depIdxs = []int32{
	1, // 0: MACCHHEFJAI.NCFICCIEHEH:type_name -> MonopolySocialEventStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MACCHHEFJAI_proto_init() }
func file_MACCHHEFJAI_proto_init() {
	if File_MACCHHEFJAI_proto != nil {
		return
	}
	file_MonopolySocialEventStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MACCHHEFJAI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MACCHHEFJAI); i {
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
			RawDescriptor: file_MACCHHEFJAI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MACCHHEFJAI_proto_goTypes,
		DependencyIndexes: file_MACCHHEFJAI_proto_depIdxs,
		MessageInfos:      file_MACCHHEFJAI_proto_msgTypes,
	}.Build()
	File_MACCHHEFJAI_proto = out.File
	file_MACCHHEFJAI_proto_rawDesc = nil
	file_MACCHHEFJAI_proto_goTypes = nil
	file_MACCHHEFJAI_proto_depIdxs = nil
}