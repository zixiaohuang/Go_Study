// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: scene.proto

package api

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

type GetRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId int64 `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *GetRoleReq) Reset() {
	*x = GetRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleReq) ProtoMessage() {}

func (x *GetRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_scene_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleReq.ProtoReflect.Descriptor instead.
func (*GetRoleReq) Descriptor() ([]byte, []int) {
	return file_scene_proto_rawDescGZIP(), []int{0}
}

func (x *GetRoleReq) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type GetRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GetRoleReply_Role `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetRoleReply) Reset() {
	*x = GetRoleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleReply) ProtoMessage() {}

func (x *GetRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_scene_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleReply.ProtoReflect.Descriptor instead.
func (*GetRoleReply) Descriptor() ([]byte, []int) {
	return file_scene_proto_rawDescGZIP(), []int{1}
}

func (x *GetRoleReply) GetItems() []*GetRoleReply_Role {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetRoleReply_Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId   int64   `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	BloodVal float32 `protobuf:"fixed32,2,opt,name=blood_val,json=bloodVal,proto3" json:"blood_val,omitempty"`
}

func (x *GetRoleReply_Role) Reset() {
	*x = GetRoleReply_Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleReply_Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleReply_Role) ProtoMessage() {}

func (x *GetRoleReply_Role) ProtoReflect() protoreflect.Message {
	mi := &file_scene_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleReply_Role.ProtoReflect.Descriptor instead.
func (*GetRoleReply_Role) Descriptor() ([]byte, []int) {
	return file_scene_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetRoleReply_Role) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *GetRoleReply_Role) GetBloodVal() float32 {
	if x != nil {
		return x.BloodVal
	}
	return 0
}

var File_scene_proto protoreflect.FileDescriptor

var file_scene_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x25, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x3c, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x56, 0x61, 0x6c, 0x32, 0x4c, 0x0a, 0x05, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x19,
	0x2e, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x61, 0x70, 0x69, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scene_proto_rawDescOnce sync.Once
	file_scene_proto_rawDescData = file_scene_proto_rawDesc
)

func file_scene_proto_rawDescGZIP() []byte {
	file_scene_proto_rawDescOnce.Do(func() {
		file_scene_proto_rawDescData = protoimpl.X.CompressGZIP(file_scene_proto_rawDescData)
	})
	return file_scene_proto_rawDescData
}

var file_scene_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_scene_proto_goTypes = []interface{}{
	(*GetRoleReq)(nil),        // 0: scene.service.GetRoleReq
	(*GetRoleReply)(nil),      // 1: scene.service.GetRoleReply
	(*GetRoleReply_Role)(nil), // 2: scene.service.GetRoleReply.Role
}
var file_scene_proto_depIdxs = []int32{
	2, // 0: scene.service.GetRoleReply.items:type_name -> scene.service.GetRoleReply.Role
	0, // 1: scene.service.Scene.GetRole:input_type -> scene.service.GetRoleReq
	1, // 2: scene.service.Scene.GetRole:output_type -> scene.service.GetRoleReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scene_proto_init() }
func file_scene_proto_init() {
	if File_scene_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scene_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleReq); i {
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
		file_scene_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleReply); i {
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
		file_scene_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleReply_Role); i {
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
			RawDescriptor: file_scene_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scene_proto_goTypes,
		DependencyIndexes: file_scene_proto_depIdxs,
		MessageInfos:      file_scene_proto_msgTypes,
	}.Build()
	File_scene_proto = out.File
	file_scene_proto_rawDesc = nil
	file_scene_proto_goTypes = nil
	file_scene_proto_depIdxs = nil
}
