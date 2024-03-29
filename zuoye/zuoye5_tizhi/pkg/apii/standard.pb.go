// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: standard.proto

package apii

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sex    string  `protobuf:"bytes,2,opt,name=sex,proto3" json:"sex,omitempty"`
	Age    int64   `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Tall   float32 `protobuf:"fixed32,4,opt,name=tall,proto3" json:"tall,omitempty"`
	Weight float32 `protobuf:"fixed32,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Fatr   float32 `protobuf:"fixed32,6,opt,name=fatr,proto3" json:"fatr,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_standard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_standard_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *Person) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetTall() float32 {
	if x != nil {
		return x.Tall
	}
	return 0
}

func (x *Person) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Person) GetFatr() float32 {
	if x != nil {
		return x.Fatr
	}
	return 0
}

type Persons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Person `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // items 是字段
}

func (x *Persons) Reset() {
	*x = Persons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Persons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Persons) ProtoMessage() {}

func (x *Persons) ProtoReflect() protoreflect.Message {
	mi := &file_standard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Persons.ProtoReflect.Descriptor instead.
func (*Persons) Descriptor() ([]byte, []int) {
	return file_standard_proto_rawDescGZIP(), []int{1}
}

func (x *Persons) GetItems() []*Person {
	if x != nil {
		return x.Items
	}
	return nil
}

type RankItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fatr float32 `protobuf:"fixed32,2,opt,name=fatr,proto3" json:"fatr,omitempty"`
	Rank int64   `protobuf:"varint,3,opt,name=ranks,proto3" json:"ranks,omitempty"`
}

func (x *RankItem) Reset() {
	*x = RankItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankItem) ProtoMessage() {}

func (x *RankItem) ProtoReflect() protoreflect.Message {
	mi := &file_standard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankItem.ProtoReflect.Descriptor instead.
func (*RankItem) Descriptor() ([]byte, []int) {
	return file_standard_proto_rawDescGZIP(), []int{2}
}

func (x *RankItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RankItem) GetFatr() float32 {
	if x != nil {
		return x.Fatr
	}
	return 0
}

func (x *RankItem) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type Rank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemsS []*RankItem `protobuf:"bytes,1,rep,name=itemsS,proto3" json:"itemsS,omitempty"`
}

func (x *Rank) Reset() {
	*x = Rank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rank) ProtoMessage() {}

func (x *Rank) ProtoReflect() protoreflect.Message {
	mi := &file_standard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rank.ProtoReflect.Descriptor instead.
func (*Rank) Descriptor() ([]byte, []int) {
	return file_standard_proto_rawDescGZIP(), []int{3}
}

func (x *Rank) GetItemsS() []*RankItem {
	if x != nil {
		return x.ItemsS
	}
	return nil
}

var File_standard_proto protoreflect.FileDescriptor

var file_standard_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x70, 0x69, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x6c,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x74, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x66, 0x61, 0x74, 0x72, 0x22, 0x2d, 0x0a, 0x07, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x46, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x74, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x66, 0x61, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x22, 0x2e, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x26, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x53, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x53,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_standard_proto_rawDescOnce sync.Once
	file_standard_proto_rawDescData = file_standard_proto_rawDesc
)

func file_standard_proto_rawDescGZIP() []byte {
	file_standard_proto_rawDescOnce.Do(func() {
		file_standard_proto_rawDescData = protoimpl.X.CompressGZIP(file_standard_proto_rawDescData)
	})
	return file_standard_proto_rawDescData
}

var file_standard_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_standard_proto_goTypes = []interface{}{
	(*Person)(nil),   // 0: apiss.Person
	(*Persons)(nil),  // 1: apiss.Persons
	(*RankItem)(nil), // 2: apiss.RankItem
	(*Rank)(nil),     // 3: apiss.Rank
}
var file_standard_proto_depIdxs = []int32{
	0, // 0: apiss.Persons.items:type_name -> apiss.Person
	2, // 1: apiss.Rank.itemsS:type_name -> apiss.RankItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_standard_proto_init() }
func file_standard_proto_init() {
	if File_standard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_standard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_standard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Persons); i {
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
		file_standard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankItem); i {
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
		file_standard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rank); i {
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
			RawDescriptor: file_standard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_standard_proto_goTypes,
		DependencyIndexes: file_standard_proto_depIdxs,
		MessageInfos:      file_standard_proto_msgTypes,
	}.Build()
	File_standard_proto = out.File
	file_standard_proto_rawDesc = nil
	file_standard_proto_goTypes = nil
	file_standard_proto_depIdxs = nil
}
