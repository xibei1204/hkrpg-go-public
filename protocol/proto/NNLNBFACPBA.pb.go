// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NNLNBFACPBA.proto

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

type NNLNBFACPBA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FPLLJLJNJJE uint32 `protobuf:"varint,10,opt,name=FPLLJLJNJJE,proto3" json:"FPLLJLJNJJE,omitempty"`
	// Types that are assignable to CODNGAFHDLP:
	//
	//	*NNLNBFACPBA_FAGLNKBHNME
	//	*NNLNBFACPBA_KIGPEAGPEDA
	//	*NNLNBFACPBA_HNPCPKFPHGB
	//	*NNLNBFACPBA_ILLPHLPAPND
	//	*NNLNBFACPBA_BFCECNEKMMP
	//	*NNLNBFACPBA_CEIDNDMOFFG
	//	*NNLNBFACPBA_GHDMKBABPCB
	//	*NNLNBFACPBA_POJKMDHMFMO
	//	*NNLNBFACPBA_KHBLHCPEFMB
	CODNGAFHDLP isNNLNBFACPBA_CODNGAFHDLP `protobuf_oneof:"CODNGAFHDLP"`
}

func (x *NNLNBFACPBA) Reset() {
	*x = NNLNBFACPBA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NNLNBFACPBA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NNLNBFACPBA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NNLNBFACPBA) ProtoMessage() {}

func (x *NNLNBFACPBA) ProtoReflect() protoreflect.Message {
	mi := &file_NNLNBFACPBA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NNLNBFACPBA.ProtoReflect.Descriptor instead.
func (*NNLNBFACPBA) Descriptor() ([]byte, []int) {
	return file_NNLNBFACPBA_proto_rawDescGZIP(), []int{0}
}

func (x *NNLNBFACPBA) GetFPLLJLJNJJE() uint32 {
	if x != nil {
		return x.FPLLJLJNJJE
	}
	return 0
}

func (m *NNLNBFACPBA) GetCODNGAFHDLP() isNNLNBFACPBA_CODNGAFHDLP {
	if m != nil {
		return m.CODNGAFHDLP
	}
	return nil
}

func (x *NNLNBFACPBA) GetFAGLNKBHNME() *ODAJHHEIHHN {
	if x, ok := x.GetCODNGAFHDLP().(*NNLNBFACPBA_FAGLNKBHNME); ok {
		return x.FAGLNKBHNME
	}
	return nil
}

func (x *NNLNBFACPBA) GetKIGPEAGPEDA() *EEMGCPEIKJE {
	if x, ok := x.GetCODNGAFHDLP().(*NNLNBFACPBA_KIGPEAGPEDA); ok {
		return x.KIGPEAGPEDA
	}
	return nil
}

func (x *NNLNBFACPBA) GetHNPCPKFPHGB() *DOHNNKHFGBP {
	if x, ok := x.GetCODNGAFHDLP().(*NNLNBFACPBA_HNPCPKFPHGB); ok {
		return x.HNPCPKFPHGB
	}
	return nil
}

func (x *NNLNBFACPBA) GetILLPHLPAPND() *LHOBCHCJBHL {
	if x, ok := x.GetCODNGAFHDLP().(*NNLNBFACPBA_ILLPHLPAPND); ok {
		return x.ILLPHLPAPND
	}
	return nil
}

func (x *NNLNBFACPBA) GetBFCECNEKMMP() *GCKHPBFPEMJ {
	if x, ok := x.GetCODNGAFHDLP().(*NNLNBFACPBA_BFCECNEKMMP); ok {
		return x.BFCECNEKMMP
	}
	return nil
}

func (x *NNLNBFACPBA) GetCEIDNDMOFFG() *EBOFBACGCKM {
	if x, ok := x.GetCODNGAFHDLP().(*NNLNBFACPBA_CEIDNDMOFFG); ok {
		return x.CEIDNDMOFFG
	}
	return nil
}

func (x *NNLNBFACPBA) GetGHDMKBABPCB() *BCKPFOLEJAI {
	if x, ok := x.GetCODNGAFHDLP().(*NNLNBFACPBA_GHDMKBABPCB); ok {
		return x.GHDMKBABPCB
	}
	return nil
}

func (x *NNLNBFACPBA) GetPOJKMDHMFMO() *FALMHHAPDGO {
	if x, ok := x.GetCODNGAFHDLP().(*NNLNBFACPBA_POJKMDHMFMO); ok {
		return x.POJKMDHMFMO
	}
	return nil
}

func (x *NNLNBFACPBA) GetKHBLHCPEFMB() *BEFCMBABLLL {
	if x, ok := x.GetCODNGAFHDLP().(*NNLNBFACPBA_KHBLHCPEFMB); ok {
		return x.KHBLHCPEFMB
	}
	return nil
}

type isNNLNBFACPBA_CODNGAFHDLP interface {
	isNNLNBFACPBA_CODNGAFHDLP()
}

type NNLNBFACPBA_FAGLNKBHNME struct {
	FAGLNKBHNME *ODAJHHEIHHN `protobuf:"bytes,6,opt,name=FAGLNKBHNME,proto3,oneof"`
}

type NNLNBFACPBA_KIGPEAGPEDA struct {
	KIGPEAGPEDA *EEMGCPEIKJE `protobuf:"bytes,1,opt,name=KIGPEAGPEDA,proto3,oneof"`
}

type NNLNBFACPBA_HNPCPKFPHGB struct {
	HNPCPKFPHGB *DOHNNKHFGBP `protobuf:"bytes,15,opt,name=HNPCPKFPHGB,proto3,oneof"`
}

type NNLNBFACPBA_ILLPHLPAPND struct {
	ILLPHLPAPND *LHOBCHCJBHL `protobuf:"bytes,14,opt,name=ILLPHLPAPND,proto3,oneof"`
}

type NNLNBFACPBA_BFCECNEKMMP struct {
	BFCECNEKMMP *GCKHPBFPEMJ `protobuf:"bytes,7,opt,name=BFCECNEKMMP,proto3,oneof"`
}

type NNLNBFACPBA_CEIDNDMOFFG struct {
	CEIDNDMOFFG *EBOFBACGCKM `protobuf:"bytes,13,opt,name=CEIDNDMOFFG,proto3,oneof"`
}

type NNLNBFACPBA_GHDMKBABPCB struct {
	GHDMKBABPCB *BCKPFOLEJAI `protobuf:"bytes,4,opt,name=GHDMKBABPCB,proto3,oneof"`
}

type NNLNBFACPBA_POJKMDHMFMO struct {
	POJKMDHMFMO *FALMHHAPDGO `protobuf:"bytes,8,opt,name=POJKMDHMFMO,proto3,oneof"`
}

type NNLNBFACPBA_KHBLHCPEFMB struct {
	KHBLHCPEFMB *BEFCMBABLLL `protobuf:"bytes,5,opt,name=KHBLHCPEFMB,proto3,oneof"`
}

func (*NNLNBFACPBA_FAGLNKBHNME) isNNLNBFACPBA_CODNGAFHDLP() {}

func (*NNLNBFACPBA_KIGPEAGPEDA) isNNLNBFACPBA_CODNGAFHDLP() {}

func (*NNLNBFACPBA_HNPCPKFPHGB) isNNLNBFACPBA_CODNGAFHDLP() {}

func (*NNLNBFACPBA_ILLPHLPAPND) isNNLNBFACPBA_CODNGAFHDLP() {}

func (*NNLNBFACPBA_BFCECNEKMMP) isNNLNBFACPBA_CODNGAFHDLP() {}

func (*NNLNBFACPBA_CEIDNDMOFFG) isNNLNBFACPBA_CODNGAFHDLP() {}

func (*NNLNBFACPBA_GHDMKBABPCB) isNNLNBFACPBA_CODNGAFHDLP() {}

func (*NNLNBFACPBA_POJKMDHMFMO) isNNLNBFACPBA_CODNGAFHDLP() {}

func (*NNLNBFACPBA_KHBLHCPEFMB) isNNLNBFACPBA_CODNGAFHDLP() {}

var File_NNLNBFACPBA_proto protoreflect.FileDescriptor

var file_NNLNBFACPBA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x4e, 0x4c, 0x4e, 0x42, 0x46, 0x41, 0x43, 0x50, 0x42, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x43, 0x4b, 0x48, 0x50, 0x42, 0x46, 0x50, 0x45, 0x4d, 0x4a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x48, 0x4f, 0x42, 0x43, 0x48, 0x43, 0x4a,
	0x42, 0x48, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x44, 0x41, 0x4a, 0x48,
	0x48, 0x45, 0x49, 0x48, 0x48, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x41,
	0x4c, 0x4d, 0x48, 0x48, 0x41, 0x50, 0x44, 0x47, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x44, 0x4f, 0x48, 0x4e, 0x4e, 0x4b, 0x48, 0x46, 0x47, 0x42, 0x50, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x45, 0x45, 0x4d, 0x47, 0x43, 0x50, 0x45, 0x49, 0x4b, 0x4a, 0x45, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x43, 0x4b, 0x50, 0x46, 0x4f, 0x4c, 0x45, 0x4a,
	0x41, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x45, 0x46, 0x43, 0x4d, 0x42,
	0x41, 0x42, 0x4c, 0x4c, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x42, 0x4f,
	0x46, 0x42, 0x41, 0x43, 0x47, 0x43, 0x4b, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80,
	0x04, 0x0a, 0x0b, 0x4e, 0x4e, 0x4c, 0x4e, 0x42, 0x46, 0x41, 0x43, 0x50, 0x42, 0x41, 0x12, 0x20,
	0x0a, 0x0b, 0x46, 0x50, 0x4c, 0x4c, 0x4a, 0x4c, 0x4a, 0x4e, 0x4a, 0x4a, 0x45, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x50, 0x4c, 0x4c, 0x4a, 0x4c, 0x4a, 0x4e, 0x4a, 0x4a, 0x45,
	0x12, 0x30, 0x0a, 0x0b, 0x46, 0x41, 0x47, 0x4c, 0x4e, 0x4b, 0x42, 0x48, 0x4e, 0x4d, 0x45, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x44, 0x41, 0x4a, 0x48, 0x48, 0x45, 0x49,
	0x48, 0x48, 0x4e, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x41, 0x47, 0x4c, 0x4e, 0x4b, 0x42, 0x48, 0x4e,
	0x4d, 0x45, 0x12, 0x30, 0x0a, 0x0b, 0x4b, 0x49, 0x47, 0x50, 0x45, 0x41, 0x47, 0x50, 0x45, 0x44,
	0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x45, 0x4d, 0x47, 0x43, 0x50,
	0x45, 0x49, 0x4b, 0x4a, 0x45, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x49, 0x47, 0x50, 0x45, 0x41, 0x47,
	0x50, 0x45, 0x44, 0x41, 0x12, 0x30, 0x0a, 0x0b, 0x48, 0x4e, 0x50, 0x43, 0x50, 0x4b, 0x46, 0x50,
	0x48, 0x47, 0x42, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4f, 0x48, 0x4e,
	0x4e, 0x4b, 0x48, 0x46, 0x47, 0x42, 0x50, 0x48, 0x00, 0x52, 0x0b, 0x48, 0x4e, 0x50, 0x43, 0x50,
	0x4b, 0x46, 0x50, 0x48, 0x47, 0x42, 0x12, 0x30, 0x0a, 0x0b, 0x49, 0x4c, 0x4c, 0x50, 0x48, 0x4c,
	0x50, 0x41, 0x50, 0x4e, 0x44, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x48,
	0x4f, 0x42, 0x43, 0x48, 0x43, 0x4a, 0x42, 0x48, 0x4c, 0x48, 0x00, 0x52, 0x0b, 0x49, 0x4c, 0x4c,
	0x50, 0x48, 0x4c, 0x50, 0x41, 0x50, 0x4e, 0x44, 0x12, 0x30, 0x0a, 0x0b, 0x42, 0x46, 0x43, 0x45,
	0x43, 0x4e, 0x45, 0x4b, 0x4d, 0x4d, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x47, 0x43, 0x4b, 0x48, 0x50, 0x42, 0x46, 0x50, 0x45, 0x4d, 0x4a, 0x48, 0x00, 0x52, 0x0b, 0x42,
	0x46, 0x43, 0x45, 0x43, 0x4e, 0x45, 0x4b, 0x4d, 0x4d, 0x50, 0x12, 0x30, 0x0a, 0x0b, 0x43, 0x45,
	0x49, 0x44, 0x4e, 0x44, 0x4d, 0x4f, 0x46, 0x46, 0x47, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x45, 0x42, 0x4f, 0x46, 0x42, 0x41, 0x43, 0x47, 0x43, 0x4b, 0x4d, 0x48, 0x00, 0x52,
	0x0b, 0x43, 0x45, 0x49, 0x44, 0x4e, 0x44, 0x4d, 0x4f, 0x46, 0x46, 0x47, 0x12, 0x30, 0x0a, 0x0b,
	0x47, 0x48, 0x44, 0x4d, 0x4b, 0x42, 0x41, 0x42, 0x50, 0x43, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x43, 0x4b, 0x50, 0x46, 0x4f, 0x4c, 0x45, 0x4a, 0x41, 0x49, 0x48,
	0x00, 0x52, 0x0b, 0x47, 0x48, 0x44, 0x4d, 0x4b, 0x42, 0x41, 0x42, 0x50, 0x43, 0x42, 0x12, 0x30,
	0x0a, 0x0b, 0x50, 0x4f, 0x4a, 0x4b, 0x4d, 0x44, 0x48, 0x4d, 0x46, 0x4d, 0x4f, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x41, 0x4c, 0x4d, 0x48, 0x48, 0x41, 0x50, 0x44, 0x47,
	0x4f, 0x48, 0x00, 0x52, 0x0b, 0x50, 0x4f, 0x4a, 0x4b, 0x4d, 0x44, 0x48, 0x4d, 0x46, 0x4d, 0x4f,
	0x12, 0x30, 0x0a, 0x0b, 0x4b, 0x48, 0x42, 0x4c, 0x48, 0x43, 0x50, 0x45, 0x46, 0x4d, 0x42, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x45, 0x46, 0x43, 0x4d, 0x42, 0x41, 0x42,
	0x4c, 0x4c, 0x4c, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x48, 0x42, 0x4c, 0x48, 0x43, 0x50, 0x45, 0x46,
	0x4d, 0x42, 0x42, 0x0d, 0x0a, 0x0b, 0x43, 0x4f, 0x44, 0x4e, 0x47, 0x41, 0x46, 0x48, 0x44, 0x4c,
	0x50, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NNLNBFACPBA_proto_rawDescOnce sync.Once
	file_NNLNBFACPBA_proto_rawDescData = file_NNLNBFACPBA_proto_rawDesc
)

func file_NNLNBFACPBA_proto_rawDescGZIP() []byte {
	file_NNLNBFACPBA_proto_rawDescOnce.Do(func() {
		file_NNLNBFACPBA_proto_rawDescData = protoimpl.X.CompressGZIP(file_NNLNBFACPBA_proto_rawDescData)
	})
	return file_NNLNBFACPBA_proto_rawDescData
}

var file_NNLNBFACPBA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NNLNBFACPBA_proto_goTypes = []interface{}{
	(*NNLNBFACPBA)(nil), // 0: NNLNBFACPBA
	(*ODAJHHEIHHN)(nil), // 1: ODAJHHEIHHN
	(*EEMGCPEIKJE)(nil), // 2: EEMGCPEIKJE
	(*DOHNNKHFGBP)(nil), // 3: DOHNNKHFGBP
	(*LHOBCHCJBHL)(nil), // 4: LHOBCHCJBHL
	(*GCKHPBFPEMJ)(nil), // 5: GCKHPBFPEMJ
	(*EBOFBACGCKM)(nil), // 6: EBOFBACGCKM
	(*BCKPFOLEJAI)(nil), // 7: BCKPFOLEJAI
	(*FALMHHAPDGO)(nil), // 8: FALMHHAPDGO
	(*BEFCMBABLLL)(nil), // 9: BEFCMBABLLL
}
var file_NNLNBFACPBA_proto_depIdxs = []int32{
	1, // 0: NNLNBFACPBA.FAGLNKBHNME:type_name -> ODAJHHEIHHN
	2, // 1: NNLNBFACPBA.KIGPEAGPEDA:type_name -> EEMGCPEIKJE
	3, // 2: NNLNBFACPBA.HNPCPKFPHGB:type_name -> DOHNNKHFGBP
	4, // 3: NNLNBFACPBA.ILLPHLPAPND:type_name -> LHOBCHCJBHL
	5, // 4: NNLNBFACPBA.BFCECNEKMMP:type_name -> GCKHPBFPEMJ
	6, // 5: NNLNBFACPBA.CEIDNDMOFFG:type_name -> EBOFBACGCKM
	7, // 6: NNLNBFACPBA.GHDMKBABPCB:type_name -> BCKPFOLEJAI
	8, // 7: NNLNBFACPBA.POJKMDHMFMO:type_name -> FALMHHAPDGO
	9, // 8: NNLNBFACPBA.KHBLHCPEFMB:type_name -> BEFCMBABLLL
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_NNLNBFACPBA_proto_init() }
func file_NNLNBFACPBA_proto_init() {
	if File_NNLNBFACPBA_proto != nil {
		return
	}
	file_GCKHPBFPEMJ_proto_init()
	file_LHOBCHCJBHL_proto_init()
	file_ODAJHHEIHHN_proto_init()
	file_FALMHHAPDGO_proto_init()
	file_DOHNNKHFGBP_proto_init()
	file_EEMGCPEIKJE_proto_init()
	file_BCKPFOLEJAI_proto_init()
	file_BEFCMBABLLL_proto_init()
	file_EBOFBACGCKM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_NNLNBFACPBA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NNLNBFACPBA); i {
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
	file_NNLNBFACPBA_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NNLNBFACPBA_FAGLNKBHNME)(nil),
		(*NNLNBFACPBA_KIGPEAGPEDA)(nil),
		(*NNLNBFACPBA_HNPCPKFPHGB)(nil),
		(*NNLNBFACPBA_ILLPHLPAPND)(nil),
		(*NNLNBFACPBA_BFCECNEKMMP)(nil),
		(*NNLNBFACPBA_CEIDNDMOFFG)(nil),
		(*NNLNBFACPBA_GHDMKBABPCB)(nil),
		(*NNLNBFACPBA_POJKMDHMFMO)(nil),
		(*NNLNBFACPBA_KHBLHCPEFMB)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_NNLNBFACPBA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NNLNBFACPBA_proto_goTypes,
		DependencyIndexes: file_NNLNBFACPBA_proto_depIdxs,
		MessageInfos:      file_NNLNBFACPBA_proto_msgTypes,
	}.Build()
	File_NNLNBFACPBA_proto = out.File
	file_NNLNBFACPBA_proto_rawDesc = nil
	file_NNLNBFACPBA_proto_goTypes = nil
	file_NNLNBFACPBA_proto_depIdxs = nil
}