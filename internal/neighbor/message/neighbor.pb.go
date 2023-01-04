// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: internal/neighbor/message/neighbor.proto

package message

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

type Neighbor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChassisId        []byte   `protobuf:"bytes,1,opt,name=chassis_id,json=chassisId,proto3" json:"chassis_id,omitempty"`
	ChassisName      string   `protobuf:"bytes,2,opt,name=chassis_name,json=chassisName,proto3" json:"chassis_name,omitempty"`
	ChassisMac       []byte   `protobuf:"bytes,3,opt,name=chassis_mac,json=chassisMac,proto3" json:"chassis_mac,omitempty"`
	ChassisInterface string   `protobuf:"bytes,4,opt,name=chassis_interface,json=chassisInterface,proto3" json:"chassis_interface,omitempty"`
	MgmtAddress      [][]byte `protobuf:"bytes,12,rep,name=mgmt_address,json=mgmtAddress,proto3" json:"mgmt_address,omitempty"`
}

func (x *Neighbor) Reset() {
	*x = Neighbor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_neighbor_message_neighbor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Neighbor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Neighbor) ProtoMessage() {}

func (x *Neighbor) ProtoReflect() protoreflect.Message {
	mi := &file_internal_neighbor_message_neighbor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Neighbor.ProtoReflect.Descriptor instead.
func (*Neighbor) Descriptor() ([]byte, []int) {
	return file_internal_neighbor_message_neighbor_proto_rawDescGZIP(), []int{0}
}

func (x *Neighbor) GetChassisId() []byte {
	if x != nil {
		return x.ChassisId
	}
	return nil
}

func (x *Neighbor) GetChassisName() string {
	if x != nil {
		return x.ChassisName
	}
	return ""
}

func (x *Neighbor) GetChassisMac() []byte {
	if x != nil {
		return x.ChassisMac
	}
	return nil
}

func (x *Neighbor) GetChassisInterface() string {
	if x != nil {
		return x.ChassisInterface
	}
	return ""
}

func (x *Neighbor) GetMgmtAddress() [][]byte {
	if x != nil {
		return x.MgmtAddress
	}
	return nil
}

var File_internal_neighbor_message_neighbor_proto protoreflect.FileDescriptor

var file_internal_neighbor_message_neighbor_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6e, 0x65, 0x69, 0x67, 0x68,
	0x62, 0x6f, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x65, 0x69, 0x67,
	0x68, 0x62, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x08, 0x4e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x73, 0x73, 0x69, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x63, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x4d, 0x61, 0x63, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x68,
	0x61, 0x73, 0x73, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x67, 0x6d, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x6d,
	0x67, 0x6d, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x1b, 0x5a, 0x19, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_neighbor_message_neighbor_proto_rawDescOnce sync.Once
	file_internal_neighbor_message_neighbor_proto_rawDescData = file_internal_neighbor_message_neighbor_proto_rawDesc
)

func file_internal_neighbor_message_neighbor_proto_rawDescGZIP() []byte {
	file_internal_neighbor_message_neighbor_proto_rawDescOnce.Do(func() {
		file_internal_neighbor_message_neighbor_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_neighbor_message_neighbor_proto_rawDescData)
	})
	return file_internal_neighbor_message_neighbor_proto_rawDescData
}

var file_internal_neighbor_message_neighbor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_neighbor_message_neighbor_proto_goTypes = []interface{}{
	(*Neighbor)(nil), // 0: Neighbor
}
var file_internal_neighbor_message_neighbor_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_neighbor_message_neighbor_proto_init() }
func file_internal_neighbor_message_neighbor_proto_init() {
	if File_internal_neighbor_message_neighbor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_neighbor_message_neighbor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Neighbor); i {
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
			RawDescriptor: file_internal_neighbor_message_neighbor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_neighbor_message_neighbor_proto_goTypes,
		DependencyIndexes: file_internal_neighbor_message_neighbor_proto_depIdxs,
		MessageInfos:      file_internal_neighbor_message_neighbor_proto_msgTypes,
	}.Build()
	File_internal_neighbor_message_neighbor_proto = out.File
	file_internal_neighbor_message_neighbor_proto_rawDesc = nil
	file_internal_neighbor_message_neighbor_proto_goTypes = nil
	file_internal_neighbor_message_neighbor_proto_depIdxs = nil
}
