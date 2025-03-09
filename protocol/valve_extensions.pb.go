// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: valve_extensions.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_valve_extensions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "protocol.boxed_type",
		Tag:           "bytes,50001,opt,name=boxed_type",
		Filename:      "valve_extensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50002,
		Name:          "protocol.additional_includes",
		Tag:           "bytes,50002,opt,name=additional_includes",
		Filename:      "valve_extensions.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string boxed_type = 50001;
	E_BoxedType = &file_valve_extensions_proto_extTypes[0]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional string additional_includes = 50002;
	E_AdditionalIncludes = &file_valve_extensions_proto_extTypes[1]
)

var File_valve_extensions_proto protoreflect.FileDescriptor

var file_valve_extensions_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x76, 0x61, 0x6c, 0x76, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3e, 0x0a, 0x0a, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x78, 0x65, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x3a, 0x4f, 0x0a, 0x13, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x86, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c,
})

var file_valve_extensions_proto_goTypes = []any{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
	(*descriptorpb.FileOptions)(nil),  // 1: google.protobuf.FileOptions
}
var file_valve_extensions_proto_depIdxs = []int32{
	0, // 0: protocol.boxed_type:extendee -> google.protobuf.FieldOptions
	1, // 1: protocol.additional_includes:extendee -> google.protobuf.FileOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_valve_extensions_proto_init() }
func file_valve_extensions_proto_init() {
	if File_valve_extensions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_valve_extensions_proto_rawDesc), len(file_valve_extensions_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_valve_extensions_proto_goTypes,
		DependencyIndexes: file_valve_extensions_proto_depIdxs,
		ExtensionInfos:    file_valve_extensions_proto_extTypes,
	}.Build()
	File_valve_extensions_proto = out.File
	file_valve_extensions_proto_goTypes = nil
	file_valve_extensions_proto_depIdxs = nil
}
