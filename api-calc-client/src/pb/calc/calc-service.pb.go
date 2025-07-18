// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: src/proto/calc-service.proto

package calc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         int32                  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Index         int32                  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Input) Reset() {
	*x = Input{}
	mi := &file_src_proto_calc_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_calc_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_src_proto_calc_service_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Input) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Quantity      int32                  `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Aver          float64                `protobuf:"fixed64,2,opt,name=aver,proto3" json:"aver,omitempty"`
	Total         int32                  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Output) Reset() {
	*x = Output{}
	mi := &file_src_proto_calc_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_calc_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_src_proto_calc_service_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Output) GetAver() float64 {
	if x != nil {
		return x.Aver
	}
	return 0
}

func (x *Output) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_src_proto_calc_service_proto protoreflect.FileDescriptor

const file_src_proto_calc_service_proto_rawDesc = "" +
	"\n" +
	"\x1csrc/proto/calc-service.proto\x12\x04calc\"3\n" +
	"\x05Input\x12\x14\n" +
	"\x05value\x18\x01 \x01(\x05R\x05value\x12\x14\n" +
	"\x05index\x18\x02 \x01(\x05R\x05index\"N\n" +
	"\x06Output\x12\x1a\n" +
	"\bquantity\x18\x01 \x01(\x05R\bquantity\x12\x12\n" +
	"\x04aver\x18\x02 \x01(\x01R\x04aver\x12\x14\n" +
	"\x05total\x18\x03 \x01(\x05R\x05total22\n" +
	"\vCalcService\x12#\n" +
	"\x04Calc\x12\v.calc.Input\x1a\f.calc.Output(\x01B\x0fZ\r./src/pb/calcb\x06proto3"

var (
	file_src_proto_calc_service_proto_rawDescOnce sync.Once
	file_src_proto_calc_service_proto_rawDescData []byte
)

func file_src_proto_calc_service_proto_rawDescGZIP() []byte {
	file_src_proto_calc_service_proto_rawDescOnce.Do(func() {
		file_src_proto_calc_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_src_proto_calc_service_proto_rawDesc), len(file_src_proto_calc_service_proto_rawDesc)))
	})
	return file_src_proto_calc_service_proto_rawDescData
}

var file_src_proto_calc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_proto_calc_service_proto_goTypes = []any{
	(*Input)(nil),  // 0: calc.Input
	(*Output)(nil), // 1: calc.Output
}
var file_src_proto_calc_service_proto_depIdxs = []int32{
	0, // 0: calc.CalcService.Calc:input_type -> calc.Input
	1, // 1: calc.CalcService.Calc:output_type -> calc.Output
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_proto_calc_service_proto_init() }
func file_src_proto_calc_service_proto_init() {
	if File_src_proto_calc_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_src_proto_calc_service_proto_rawDesc), len(file_src_proto_calc_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_calc_service_proto_goTypes,
		DependencyIndexes: file_src_proto_calc_service_proto_depIdxs,
		MessageInfos:      file_src_proto_calc_service_proto_msgTypes,
	}.Build()
	File_src_proto_calc_service_proto = out.File
	file_src_proto_calc_service_proto_goTypes = nil
	file_src_proto_calc_service_proto_depIdxs = nil
}
