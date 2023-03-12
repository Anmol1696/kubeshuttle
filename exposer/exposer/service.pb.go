// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: exposer/service.proto

package exposer

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseNodeID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,proto3" json:"node_id,omitempty"`
}

func (x *ResponseNodeID) Reset() {
	*x = ResponseNodeID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exposer_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseNodeID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseNodeID) ProtoMessage() {}

func (x *ResponseNodeID) ProtoReflect() protoreflect.Message {
	mi := &file_exposer_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseNodeID.ProtoReflect.Descriptor instead.
func (*ResponseNodeID) Descriptor() ([]byte, []int) {
	return file_exposer_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseNodeID) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type ResponsePubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ResponsePubKey) Reset() {
	*x = ResponsePubKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exposer_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePubKey) ProtoMessage() {}

func (x *ResponsePubKey) ProtoReflect() protoreflect.Message {
	mi := &file_exposer_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePubKey.ProtoReflect.Descriptor instead.
func (*ResponsePubKey) Descriptor() ([]byte, []int) {
	return file_exposer_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResponsePubKey) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResponsePubKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ResponseFileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseFileData) Reset() {
	*x = ResponseFileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exposer_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFileData) ProtoMessage() {}

func (x *ResponseFileData) ProtoReflect() protoreflect.Message {
	mi := &file_exposer_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFileData.ProtoReflect.Descriptor instead.
func (*ResponseFileData) Descriptor() ([]byte, []int) {
	return file_exposer_service_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseFileData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_exposer_service_proto protoreflect.FileDescriptor

var file_exposer_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x2f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x26, 0x0a, 0x10, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x32, 0x8f, 0x03, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12,
	0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x12,
	0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x12,
	0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e,
	0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x0d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x58, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x4b, 0x65, 0x79, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x69, 0x76, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4b, 0x65,
	0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x70, 0x72, 0x69, 0x76,
	0x5f, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x7e, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x6e, 0x6d, 0x6f, 0x6c, 0x31, 0x36, 0x39, 0x36, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x45, 0x58,
	0x58, 0xaa, 0x02, 0x07, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0xca, 0x02, 0x07, 0x45, 0x78,
	0x70, 0x6f, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x13, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x45, 0x78,
	0x70, 0x6f, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exposer_service_proto_rawDescOnce sync.Once
	file_exposer_service_proto_rawDescData = file_exposer_service_proto_rawDesc
)

func file_exposer_service_proto_rawDescGZIP() []byte {
	file_exposer_service_proto_rawDescOnce.Do(func() {
		file_exposer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_exposer_service_proto_rawDescData)
	})
	return file_exposer_service_proto_rawDescData
}

var file_exposer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_exposer_service_proto_goTypes = []interface{}{
	(*ResponseNodeID)(nil),   // 0: exposer.ResponseNodeID
	(*ResponsePubKey)(nil),   // 1: exposer.ResponsePubKey
	(*ResponseFileData)(nil), // 2: exposer.ResponseFileData
	(*emptypb.Empty)(nil),    // 3: google.protobuf.Empty
	(*GenesisState)(nil),     // 4: GenesisState
	(*Keys)(nil),             // 5: exposer.Keys
	(*PrivValidatorKey)(nil), // 6: exposer.PrivValidatorKey
}
var file_exposer_service_proto_depIdxs = []int32{
	3, // 0: exposer.Exposer.GetNodeID:input_type -> google.protobuf.Empty
	3, // 1: exposer.Exposer.GetPubKey:input_type -> google.protobuf.Empty
	3, // 2: exposer.Exposer.GetGenesisFile:input_type -> google.protobuf.Empty
	3, // 3: exposer.Exposer.GetKeys:input_type -> google.protobuf.Empty
	3, // 4: exposer.Exposer.GetPrivKeysFile:input_type -> google.protobuf.Empty
	0, // 5: exposer.Exposer.GetNodeID:output_type -> exposer.ResponseNodeID
	1, // 6: exposer.Exposer.GetPubKey:output_type -> exposer.ResponsePubKey
	4, // 7: exposer.Exposer.GetGenesisFile:output_type -> GenesisState
	5, // 8: exposer.Exposer.GetKeys:output_type -> exposer.Keys
	6, // 9: exposer.Exposer.GetPrivKeysFile:output_type -> exposer.PrivValidatorKey
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_exposer_service_proto_init() }
func file_exposer_service_proto_init() {
	if File_exposer_service_proto != nil {
		return
	}
	file_exposer_mnemonic_proto_init()
	file_exposer_genesis_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_exposer_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseNodeID); i {
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
		file_exposer_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePubKey); i {
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
		file_exposer_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFileData); i {
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
			RawDescriptor: file_exposer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exposer_service_proto_goTypes,
		DependencyIndexes: file_exposer_service_proto_depIdxs,
		MessageInfos:      file_exposer_service_proto_msgTypes,
	}.Build()
	File_exposer_service_proto = out.File
	file_exposer_service_proto_rawDesc = nil
	file_exposer_service_proto_goTypes = nil
	file_exposer_service_proto_depIdxs = nil
}