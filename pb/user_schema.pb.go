// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: user_schema.proto

package tinder_matching

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Height              uint32     `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Gender              UserGender `protobuf:"varint,4,opt,name=gender,proto3,enum=user.UserGender" json:"gender,omitempty"`
	RemainNumberOfDates uint32     `protobuf:"varint,5,opt,name=remain_number_of_dates,json=remainNumberOfDates,proto3" json:"remain_number_of_dates,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_schema_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *User) GetGender() UserGender {
	if x != nil {
		return x.Gender
	}
	return UserGender_USER_GENDER_MALE
}

func (x *User) GetRemainNumberOfDates() uint32 {
	if x != nil {
		return x.RemainNumberOfDates
	}
	return 0
}

var File_user_schema_proto protoreflect.FileDescriptor

var file_user_schema_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0x92, 0x41, 0x09, 0x32, 0x07, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x32, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x3a, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x32, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x20, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x16, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x1b, 0x92, 0x41, 0x18,
	0x32, 0x16, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20,
	0x6f, 0x66, 0x20, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x13, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x44, 0x61, 0x74, 0x65, 0x73, 0x42, 0x25, 0x5a,
	0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x79, 0x31,
	0x31, 0x32, 0x32, 0x39, 0x2f, 0x74, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_schema_proto_rawDescOnce sync.Once
	file_user_schema_proto_rawDescData = file_user_schema_proto_rawDesc
)

func file_user_schema_proto_rawDescGZIP() []byte {
	file_user_schema_proto_rawDescOnce.Do(func() {
		file_user_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_schema_proto_rawDescData)
	})
	return file_user_schema_proto_rawDescData
}

var file_user_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_schema_proto_goTypes = []interface{}{
	(*User)(nil),    // 0: user.User
	(UserGender)(0), // 1: user.UserGender
}
var file_user_schema_proto_depIdxs = []int32{
	1, // 0: user.User.gender:type_name -> user.UserGender
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_schema_proto_init() }
func file_user_schema_proto_init() {
	if File_user_schema_proto != nil {
		return
	}
	file_user_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_user_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_schema_proto_goTypes,
		DependencyIndexes: file_user_schema_proto_depIdxs,
		MessageInfos:      file_user_schema_proto_msgTypes,
	}.Build()
	File_user_schema_proto = out.File
	file_user_schema_proto_rawDesc = nil
	file_user_schema_proto_goTypes = nil
	file_user_schema_proto_depIdxs = nil
}
