// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMonopolyInfoScRsp.proto

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

type GetMonopolyInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FIKKMEAGPOP   *DKACFEEHFMC       `protobuf:"bytes,7,opt,name=FIKKMEAGPOP,proto3" json:"FIKKMEAGPOP,omitempty"`
	LDONALFKGCP   *MACCHHEFJAI       `protobuf:"bytes,8,opt,name=LDONALFKGCP,proto3" json:"LDONALFKGCP,omitempty"`
	HHNEAAMFPBL   *CCNLHIIDNKM       `protobuf:"bytes,1,opt,name=HHNEAAMFPBL,proto3" json:"HHNEAAMFPBL,omitempty"`
	CPBHKDKBNKH   []uint32           `protobuf:"varint,4,rep,packed,name=CPBHKDKBNKH,proto3" json:"CPBHKDKBNKH,omitempty"`
	RogueBuffInfo *MonopolyBuffInfo  `protobuf:"bytes,10,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	MapInfo       *MonopolyMapInfo   `protobuf:"bytes,12,opt,name=map_info,json=mapInfo,proto3" json:"map_info,omitempty"`
	Retcode       uint32             `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MMECFEOPAMA   *ACAJPHPNKAE       `protobuf:"bytes,15,opt,name=MMECFEOPAMA,proto3" json:"MMECFEOPAMA,omitempty"`
	PCJOKBGOEKJ   *EMKPPMIIENK       `protobuf:"bytes,2,opt,name=PCJOKBGOEKJ,proto3" json:"PCJOKBGOEKJ,omitempty"`
	Stt           *MonopolyEventInfo `protobuf:"bytes,5,opt,name=stt,proto3" json:"stt,omitempty"`
	HGHFOEOHIMD   *OLDDDDEDBIK       `protobuf:"bytes,3,opt,name=HGHFOEOHIMD,proto3" json:"HGHFOEOHIMD,omitempty"`
	Report        *MonopolyReport    `protobuf:"bytes,6,opt,name=report,proto3" json:"report,omitempty"`
	GKLPHPDHCEK   *NEJKDENGLKO       `protobuf:"bytes,13,opt,name=GKLPHPDHCEK,proto3" json:"GKLPHPDHCEK,omitempty"`
}

func (x *GetMonopolyInfoScRsp) Reset() {
	*x = GetMonopolyInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMonopolyInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMonopolyInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMonopolyInfoScRsp) ProtoMessage() {}

func (x *GetMonopolyInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMonopolyInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMonopolyInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetMonopolyInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetMonopolyInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMonopolyInfoScRsp) GetFIKKMEAGPOP() *DKACFEEHFMC {
	if x != nil {
		return x.FIKKMEAGPOP
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetLDONALFKGCP() *MACCHHEFJAI {
	if x != nil {
		return x.LDONALFKGCP
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetHHNEAAMFPBL() *CCNLHIIDNKM {
	if x != nil {
		return x.HHNEAAMFPBL
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetCPBHKDKBNKH() []uint32 {
	if x != nil {
		return x.CPBHKDKBNKH
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetRogueBuffInfo() *MonopolyBuffInfo {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetMapInfo() *MonopolyMapInfo {
	if x != nil {
		return x.MapInfo
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMonopolyInfoScRsp) GetMMECFEOPAMA() *ACAJPHPNKAE {
	if x != nil {
		return x.MMECFEOPAMA
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetPCJOKBGOEKJ() *EMKPPMIIENK {
	if x != nil {
		return x.PCJOKBGOEKJ
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetStt() *MonopolyEventInfo {
	if x != nil {
		return x.Stt
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetHGHFOEOHIMD() *OLDDDDEDBIK {
	if x != nil {
		return x.HGHFOEOHIMD
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetReport() *MonopolyReport {
	if x != nil {
		return x.Report
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetGKLPHPDHCEK() *NEJKDENGLKO {
	if x != nil {
		return x.GKLPHPDHCEK
	}
	return nil
}

var File_GetMonopolyInfoScRsp_proto protoreflect.FileDescriptor

var file_GetMonopolyInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x43,
	0x41, 0x4a, 0x50, 0x48, 0x50, 0x4e, 0x4b, 0x41, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4e, 0x45, 0x4a, 0x4b, 0x44, 0x45, 0x4e, 0x47, 0x4c, 0x4b, 0x4f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x42, 0x75, 0x66, 0x66,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4d, 0x4b, 0x50,
	0x50, 0x4d, 0x49, 0x49, 0x45, 0x4e, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d,
	0x41, 0x43, 0x43, 0x48, 0x48, 0x45, 0x46, 0x4a, 0x41, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x4f, 0x4c, 0x44, 0x44, 0x44, 0x44, 0x45, 0x44, 0x42, 0x49, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4b, 0x41, 0x43, 0x46, 0x45, 0x45, 0x48, 0x46, 0x4d, 0x43,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x43, 0x4e, 0x4c, 0x48, 0x49, 0x49, 0x44, 0x4e,
	0x4b, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f,
	0x6c, 0x79, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd9, 0x04, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x49, 0x4b, 0x4b,
	0x4d, 0x45, 0x41, 0x47, 0x50, 0x4f, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x44, 0x4b, 0x41, 0x43, 0x46, 0x45, 0x45, 0x48, 0x46, 0x4d, 0x43, 0x52, 0x0b, 0x46, 0x49, 0x4b,
	0x4b, 0x4d, 0x45, 0x41, 0x47, 0x50, 0x4f, 0x50, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x44, 0x4f, 0x4e,
	0x41, 0x4c, 0x46, 0x4b, 0x47, 0x43, 0x50, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4d, 0x41, 0x43, 0x43, 0x48, 0x48, 0x45, 0x46, 0x4a, 0x41, 0x49, 0x52, 0x0b, 0x4c, 0x44, 0x4f,
	0x4e, 0x41, 0x4c, 0x46, 0x4b, 0x47, 0x43, 0x50, 0x12, 0x2e, 0x0a, 0x0b, 0x48, 0x48, 0x4e, 0x45,
	0x41, 0x41, 0x4d, 0x46, 0x50, 0x42, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x43, 0x43, 0x4e, 0x4c, 0x48, 0x49, 0x49, 0x44, 0x4e, 0x4b, 0x4d, 0x52, 0x0b, 0x48, 0x48, 0x4e,
	0x45, 0x41, 0x41, 0x4d, 0x46, 0x50, 0x42, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x50, 0x42, 0x48,
	0x4b, 0x44, 0x4b, 0x42, 0x4e, 0x4b, 0x48, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x43,
	0x50, 0x42, 0x48, 0x4b, 0x44, 0x4b, 0x42, 0x4e, 0x4b, 0x48, 0x12, 0x39, 0x0a, 0x0f, 0x72, 0x6f,
	0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x42, 0x75,
	0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66,
	0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f,
	0x6c, 0x79, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b,
	0x4d, 0x4d, 0x45, 0x43, 0x46, 0x45, 0x4f, 0x50, 0x41, 0x4d, 0x41, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x43, 0x41, 0x4a, 0x50, 0x48, 0x50, 0x4e, 0x4b, 0x41, 0x45, 0x52,
	0x0b, 0x4d, 0x4d, 0x45, 0x43, 0x46, 0x45, 0x4f, 0x50, 0x41, 0x4d, 0x41, 0x12, 0x2e, 0x0a, 0x0b,
	0x50, 0x43, 0x4a, 0x4f, 0x4b, 0x42, 0x47, 0x4f, 0x45, 0x4b, 0x4a, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x4d, 0x4b, 0x50, 0x50, 0x4d, 0x49, 0x49, 0x45, 0x4e, 0x4b, 0x52,
	0x0b, 0x50, 0x43, 0x4a, 0x4f, 0x4b, 0x42, 0x47, 0x4f, 0x45, 0x4b, 0x4a, 0x12, 0x24, 0x0a, 0x03,
	0x73, 0x74, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4d, 0x6f, 0x6e, 0x6f,
	0x70, 0x6f, 0x6c, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x73,
	0x74, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x48, 0x47, 0x48, 0x46, 0x4f, 0x45, 0x4f, 0x48, 0x49, 0x4d,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4c, 0x44, 0x44, 0x44, 0x44,
	0x45, 0x44, 0x42, 0x49, 0x4b, 0x52, 0x0b, 0x48, 0x47, 0x48, 0x46, 0x4f, 0x45, 0x4f, 0x48, 0x49,
	0x4d, 0x44, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x47,
	0x4b, 0x4c, 0x50, 0x48, 0x50, 0x44, 0x48, 0x43, 0x45, 0x4b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4e, 0x45, 0x4a, 0x4b, 0x44, 0x45, 0x4e, 0x47, 0x4c, 0x4b, 0x4f, 0x52, 0x0b,
	0x47, 0x4b, 0x4c, 0x50, 0x48, 0x50, 0x44, 0x48, 0x43, 0x45, 0x4b, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetMonopolyInfoScRsp_proto_rawDescOnce sync.Once
	file_GetMonopolyInfoScRsp_proto_rawDescData = file_GetMonopolyInfoScRsp_proto_rawDesc
)

func file_GetMonopolyInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetMonopolyInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetMonopolyInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMonopolyInfoScRsp_proto_rawDescData)
	})
	return file_GetMonopolyInfoScRsp_proto_rawDescData
}

var file_GetMonopolyInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMonopolyInfoScRsp_proto_goTypes = []interface{}{
	(*GetMonopolyInfoScRsp)(nil), // 0: GetMonopolyInfoScRsp
	(*DKACFEEHFMC)(nil),          // 1: DKACFEEHFMC
	(*MACCHHEFJAI)(nil),          // 2: MACCHHEFJAI
	(*CCNLHIIDNKM)(nil),          // 3: CCNLHIIDNKM
	(*MonopolyBuffInfo)(nil),     // 4: MonopolyBuffInfo
	(*MonopolyMapInfo)(nil),      // 5: MonopolyMapInfo
	(*ACAJPHPNKAE)(nil),          // 6: ACAJPHPNKAE
	(*EMKPPMIIENK)(nil),          // 7: EMKPPMIIENK
	(*MonopolyEventInfo)(nil),    // 8: MonopolyEventInfo
	(*OLDDDDEDBIK)(nil),          // 9: OLDDDDEDBIK
	(*MonopolyReport)(nil),       // 10: MonopolyReport
	(*NEJKDENGLKO)(nil),          // 11: NEJKDENGLKO
}
var file_GetMonopolyInfoScRsp_proto_depIdxs = []int32{
	1,  // 0: GetMonopolyInfoScRsp.FIKKMEAGPOP:type_name -> DKACFEEHFMC
	2,  // 1: GetMonopolyInfoScRsp.LDONALFKGCP:type_name -> MACCHHEFJAI
	3,  // 2: GetMonopolyInfoScRsp.HHNEAAMFPBL:type_name -> CCNLHIIDNKM
	4,  // 3: GetMonopolyInfoScRsp.rogue_buff_info:type_name -> MonopolyBuffInfo
	5,  // 4: GetMonopolyInfoScRsp.map_info:type_name -> MonopolyMapInfo
	6,  // 5: GetMonopolyInfoScRsp.MMECFEOPAMA:type_name -> ACAJPHPNKAE
	7,  // 6: GetMonopolyInfoScRsp.PCJOKBGOEKJ:type_name -> EMKPPMIIENK
	8,  // 7: GetMonopolyInfoScRsp.stt:type_name -> MonopolyEventInfo
	9,  // 8: GetMonopolyInfoScRsp.HGHFOEOHIMD:type_name -> OLDDDDEDBIK
	10, // 9: GetMonopolyInfoScRsp.report:type_name -> MonopolyReport
	11, // 10: GetMonopolyInfoScRsp.GKLPHPDHCEK:type_name -> NEJKDENGLKO
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_GetMonopolyInfoScRsp_proto_init() }
func file_GetMonopolyInfoScRsp_proto_init() {
	if File_GetMonopolyInfoScRsp_proto != nil {
		return
	}
	file_ACAJPHPNKAE_proto_init()
	file_NEJKDENGLKO_proto_init()
	file_MonopolyBuffInfo_proto_init()
	file_EMKPPMIIENK_proto_init()
	file_MACCHHEFJAI_proto_init()
	file_OLDDDDEDBIK_proto_init()
	file_DKACFEEHFMC_proto_init()
	file_MonopolyEventInfo_proto_init()
	file_MonopolyReport_proto_init()
	file_CCNLHIIDNKM_proto_init()
	file_MonopolyMapInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMonopolyInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMonopolyInfoScRsp); i {
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
			RawDescriptor: file_GetMonopolyInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMonopolyInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetMonopolyInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetMonopolyInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetMonopolyInfoScRsp_proto = out.File
	file_GetMonopolyInfoScRsp_proto_rawDesc = nil
	file_GetMonopolyInfoScRsp_proto_goTypes = nil
	file_GetMonopolyInfoScRsp_proto_depIdxs = nil
}