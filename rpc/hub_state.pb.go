// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.21.12
// source: hub_state.proto

package rpc

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

type HubState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  uint32 last_eth_block = 1; // Deprecated
	LastFnameProof uint64 `protobuf:"varint,2,opt,name=last_fname_proof,json=lastFnameProof,proto3" json:"last_fname_proof,omitempty"`
	LastL2Block    uint64 `protobuf:"varint,3,opt,name=last_l2_block,json=lastL2Block,proto3" json:"last_l2_block,omitempty"` //  bool syncEvents = 4; // Deprecated
}

func (x *HubState) Reset() {
	*x = HubState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HubState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HubState) ProtoMessage() {}

func (x *HubState) ProtoReflect() protoreflect.Message {
	mi := &file_hub_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HubState.ProtoReflect.Descriptor instead.
func (*HubState) Descriptor() ([]byte, []int) {
	return file_hub_state_proto_rawDescGZIP(), []int{0}
}

func (x *HubState) GetLastFnameProof() uint64 {
	if x != nil {
		return x.LastFnameProof
	}
	return 0
}

func (x *HubState) GetLastL2Block() uint64 {
	if x != nil {
		return x.LastL2Block
	}
	return 0
}

var File_hub_state_proto protoreflect.FileDescriptor

var file_hub_state_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x58, 0x0a, 0x08, 0x48, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x6e, 0x61,
	0x6d, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6c, 0x32, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x4c, 0x32, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x74,
	0x69, 0x63, 0x6b, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x61, 0x72, 0x63, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hub_state_proto_rawDescOnce sync.Once
	file_hub_state_proto_rawDescData = file_hub_state_proto_rawDesc
)

func file_hub_state_proto_rawDescGZIP() []byte {
	file_hub_state_proto_rawDescOnce.Do(func() {
		file_hub_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_hub_state_proto_rawDescData)
	})
	return file_hub_state_proto_rawDescData
}

var file_hub_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hub_state_proto_goTypes = []interface{}{
	(*HubState)(nil), // 0: HubState
}
var file_hub_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hub_state_proto_init() }
func file_hub_state_proto_init() {
	if File_hub_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hub_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HubState); i {
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
			RawDescriptor: file_hub_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hub_state_proto_goTypes,
		DependencyIndexes: file_hub_state_proto_depIdxs,
		MessageInfos:      file_hub_state_proto_msgTypes,
	}.Build()
	File_hub_state_proto = out.File
	file_hub_state_proto_rawDesc = nil
	file_hub_state_proto_goTypes = nil
	file_hub_state_proto_depIdxs = nil
}
