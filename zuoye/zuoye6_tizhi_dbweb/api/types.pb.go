// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: types.proto

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

type PersonalShow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"column:id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"column:id"`
	// @gotags: gorm:"column:person_id"
	PersonID int64 `protobuf:"varint,2,opt,name=personID,proto3" json:"personID,omitempty" gorm:"column:person_id"`
	// @gotags: gorm:"column:person_name"
	Name            string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" gorm:"column:person_name"`
	ShowDescription string `protobuf:"bytes,5,opt,name=showDescription,proto3" json:"showDescription,omitempty"`
	// @gotags: gorm:"column:by_time_weight"
	Weight float32 `protobuf:"fixed32,6,opt,name=weight,proto3" json:"weight,omitempty" gorm:"column:by_time_weight"`
	// @gotags: gorm:"column:by_time_tall"
	Tall float32 `protobuf:"fixed32,7,opt,name=tall,proto3" json:"tall,omitempty" gorm:"column:by_time_tall"`
	// @gotags: gorm:"column:by_time_age"
	Age int64 `protobuf:"varint,8,opt,name=age,proto3" json:"age,omitempty" gorm:"column:by_time_age"`
	// @gotags: gorm:"column:visiable"
	Visiable bool `protobuf:"varint,9,opt,name=visiable,proto3" json:"visiable,omitempty" gorm:"column:visiable"`
}

func (x *PersonalShow) Reset() {
	*x = PersonalShow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalShow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalShow) ProtoMessage() {}

func (x *PersonalShow) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalShow.ProtoReflect.Descriptor instead.
func (*PersonalShow) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *PersonalShow) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PersonalShow) GetPersonID() int64 {
	if x != nil {
		return x.PersonID
	}
	return 0
}

func (x *PersonalShow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonalShow) GetShowDescription() string {
	if x != nil {
		return x.ShowDescription
	}
	return ""
}

func (x *PersonalShow) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *PersonalShow) GetTall() float32 {
	if x != nil {
		return x.Tall
	}
	return 0
}

func (x *PersonalShow) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *PersonalShow) GetVisiable() bool {
	if x != nil {
		return x.Visiable
	}
	return false
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x22, 0xd2, 0x01, 0x0a, 0x0c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x53,
	0x68, 0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x77, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x68,
	0x6f, 0x77, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x6c, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76,
	0x69, 0x73, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76,
	0x69, 0x73, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_proto_goTypes = []interface{}{
	(*PersonalShow)(nil), // 0: api.PersonalShow
}
var file_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonalShow); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
