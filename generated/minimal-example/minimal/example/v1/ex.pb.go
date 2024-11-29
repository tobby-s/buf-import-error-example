// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: minimal/example/v1/ex.proto

package examplev1

import (
	_ "github.com/tobby-s/minimal-example/google/api"
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

type OptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []*EntityRequest `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *OptionsRequest) Reset() {
	*x = OptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minimal_example_v1_ex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionsRequest) ProtoMessage() {}

func (x *OptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_minimal_example_v1_ex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionsRequest.ProtoReflect.Descriptor instead.
func (*OptionsRequest) Descriptor() ([]byte, []int) {
	return file_minimal_example_v1_ex_proto_rawDescGZIP(), []int{0}
}

func (x *OptionsRequest) GetEntities() []*EntityRequest {
	if x != nil {
		return x.Entities
	}
	return nil
}

type EntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EntityRequest) Reset() {
	*x = EntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minimal_example_v1_ex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityRequest) ProtoMessage() {}

func (x *EntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_minimal_example_v1_ex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityRequest.ProtoReflect.Descriptor instead.
func (*EntityRequest) Descriptor() ([]byte, []int) {
	return file_minimal_example_v1_ex_proto_rawDescGZIP(), []int{1}
}

func (x *EntityRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OptionsResponse) Reset() {
	*x = OptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minimal_example_v1_ex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionsResponse) ProtoMessage() {}

func (x *OptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_minimal_example_v1_ex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionsResponse.ProtoReflect.Descriptor instead.
func (*OptionsResponse) Descriptor() ([]byte, []int) {
	return file_minimal_example_v1_ex_proto_rawDescGZIP(), []int{2}
}

func (x *OptionsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_minimal_example_v1_ex_proto protoreflect.FileDescriptor

var file_minimal_example_v1_ex_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4f, 0x0a, 0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3d, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x22, 0x1f, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x21, 0x0a, 0x0f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x72, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x67, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0xcc, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x07, 0x45, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x62, 0x62, 0x79,
	0x2d, 0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x4d, 0x45, 0x58, 0xaa, 0x02, 0x12, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x4d, 0x69, 0x6e,
	0x69, 0x6d, 0x61, 0x6c, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1e, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x14, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_minimal_example_v1_ex_proto_rawDescOnce sync.Once
	file_minimal_example_v1_ex_proto_rawDescData = file_minimal_example_v1_ex_proto_rawDesc
)

func file_minimal_example_v1_ex_proto_rawDescGZIP() []byte {
	file_minimal_example_v1_ex_proto_rawDescOnce.Do(func() {
		file_minimal_example_v1_ex_proto_rawDescData = protoimpl.X.CompressGZIP(file_minimal_example_v1_ex_proto_rawDescData)
	})
	return file_minimal_example_v1_ex_proto_rawDescData
}

var file_minimal_example_v1_ex_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_minimal_example_v1_ex_proto_goTypes = []interface{}{
	(*OptionsRequest)(nil),  // 0: minimal.example.v1.OptionsRequest
	(*EntityRequest)(nil),   // 1: minimal.example.v1.EntityRequest
	(*OptionsResponse)(nil), // 2: minimal.example.v1.OptionsResponse
}
var file_minimal_example_v1_ex_proto_depIdxs = []int32{
	1, // 0: minimal.example.v1.OptionsRequest.entities:type_name -> minimal.example.v1.EntityRequest
	0, // 1: minimal.example.v1.Service.Options:input_type -> minimal.example.v1.OptionsRequest
	2, // 2: minimal.example.v1.Service.Options:output_type -> minimal.example.v1.OptionsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_minimal_example_v1_ex_proto_init() }
func file_minimal_example_v1_ex_proto_init() {
	if File_minimal_example_v1_ex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_minimal_example_v1_ex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionsRequest); i {
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
		file_minimal_example_v1_ex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityRequest); i {
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
		file_minimal_example_v1_ex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionsResponse); i {
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
			RawDescriptor: file_minimal_example_v1_ex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_minimal_example_v1_ex_proto_goTypes,
		DependencyIndexes: file_minimal_example_v1_ex_proto_depIdxs,
		MessageInfos:      file_minimal_example_v1_ex_proto_msgTypes,
	}.Build()
	File_minimal_example_v1_ex_proto = out.File
	file_minimal_example_v1_ex_proto_rawDesc = nil
	file_minimal_example_v1_ex_proto_goTypes = nil
	file_minimal_example_v1_ex_proto_depIdxs = nil
}
