// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api/rbac/v1/rbac.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_v1_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_v1_rbac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_rbac_v1_rbac_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_v1_rbac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_v1_rbac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_api_rbac_v1_rbac_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MeRequest) Reset() {
	*x = MeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_v1_rbac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeRequest) ProtoMessage() {}

func (x *MeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_v1_rbac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeRequest.ProtoReflect.Descriptor instead.
func (*MeRequest) Descriptor() ([]byte, []int) {
	return file_api_rbac_v1_rbac_proto_rawDescGZIP(), []int{2}
}

type MeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
}

func (x *MeReply) Reset() {
	*x = MeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_v1_rbac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeReply) ProtoMessage() {}

func (x *MeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_v1_rbac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeReply.ProtoReflect.Descriptor instead.
func (*MeReply) Descriptor() ([]byte, []int) {
	return file_api_rbac_v1_rbac_proto_rawDescGZIP(), []int{3}
}

func (x *MeReply) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

var File_api_rbac_v1_rbac_proto protoreflect.FileDescriptor

var file_api_rbac_v1_rbac_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x62,
	0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62,
	0x61, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x26, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x0b, 0x0a, 0x09, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a,
	0x07, 0x4d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x32, 0x93, 0x01, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x4d, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x72, 0x62, 0x61,
	0x63, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x3e, 0x0a, 0x02, 0x4d, 0x65,
	0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x4d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63,
	0x2e, 0x4d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x6d, 0x65, 0x42, 0x3f, 0x0a, 0x16, 0x64, 0x65,
	0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61,
	0x63, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x52, 0x42, 0x41, 0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x16, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_rbac_v1_rbac_proto_rawDescOnce sync.Once
	file_api_rbac_v1_rbac_proto_rawDescData = file_api_rbac_v1_rbac_proto_rawDesc
)

func file_api_rbac_v1_rbac_proto_rawDescGZIP() []byte {
	file_api_rbac_v1_rbac_proto_rawDescOnce.Do(func() {
		file_api_rbac_v1_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rbac_v1_rbac_proto_rawDescData)
	})
	return file_api_rbac_v1_rbac_proto_rawDescData
}

var file_api_rbac_v1_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_rbac_v1_rbac_proto_goTypes = []interface{}{
	(*LoginRequest)(nil), // 0: api.rbac.LoginRequest
	(*LoginReply)(nil),   // 1: api.rbac.LoginReply
	(*MeRequest)(nil),    // 2: api.rbac.MeRequest
	(*MeReply)(nil),      // 3: api.rbac.MeReply
}
var file_api_rbac_v1_rbac_proto_depIdxs = []int32{
	0, // 0: api.rbac.V1.Login:input_type -> api.rbac.LoginRequest
	2, // 1: api.rbac.V1.Me:input_type -> api.rbac.MeRequest
	1, // 2: api.rbac.V1.Login:output_type -> api.rbac.LoginReply
	3, // 3: api.rbac.V1.Me:output_type -> api.rbac.MeReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_rbac_v1_rbac_proto_init() }
func file_api_rbac_v1_rbac_proto_init() {
	if File_api_rbac_v1_rbac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rbac_v1_rbac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_api_rbac_v1_rbac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_api_rbac_v1_rbac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeRequest); i {
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
		file_api_rbac_v1_rbac_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeReply); i {
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
			RawDescriptor: file_api_rbac_v1_rbac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_rbac_v1_rbac_proto_goTypes,
		DependencyIndexes: file_api_rbac_v1_rbac_proto_depIdxs,
		MessageInfos:      file_api_rbac_v1_rbac_proto_msgTypes,
	}.Build()
	File_api_rbac_v1_rbac_proto = out.File
	file_api_rbac_v1_rbac_proto_rawDesc = nil
	file_api_rbac_v1_rbac_proto_goTypes = nil
	file_api_rbac_v1_rbac_proto_depIdxs = nil
}
