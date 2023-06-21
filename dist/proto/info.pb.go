// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.21.5
// source: dist/proto/info.proto

package proto

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

type BenchSmall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Key    []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *BenchSmall) Reset() {
	*x = BenchSmall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchSmall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchSmall) ProtoMessage() {}

func (x *BenchSmall) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchSmall.ProtoReflect.Descriptor instead.
func (*BenchSmall) Descriptor() ([]byte, []int) {
	return file_dist_proto_info_proto_rawDescGZIP(), []int{0}
}

func (x *BenchSmall) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *BenchSmall) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type BenchMedium struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age    int64   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Height float32 `protobuf:"fixed32,3,opt,name=height,proto3" json:"height,omitempty"`
	Weight float64 `protobuf:"fixed64,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Alive  bool    `protobuf:"varint,5,opt,name=alive,proto3" json:"alive,omitempty"`
	Desc   []byte  `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *BenchMedium) Reset() {
	*x = BenchMedium{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchMedium) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchMedium) ProtoMessage() {}

func (x *BenchMedium) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchMedium.ProtoReflect.Descriptor instead.
func (*BenchMedium) Descriptor() ([]byte, []int) {
	return file_dist_proto_info_proto_rawDescGZIP(), []int{1}
}

func (x *BenchMedium) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BenchMedium) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *BenchMedium) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BenchMedium) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *BenchMedium) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

func (x *BenchMedium) GetDesc() []byte {
	if x != nil {
		return x.Desc
	}
	return nil
}

type BenchLarge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age      int64   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Height   float32 `protobuf:"fixed32,3,opt,name=height,proto3" json:"height,omitempty"`
	Weight   float64 `protobuf:"fixed64,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Alive    bool    `protobuf:"varint,5,opt,name=alive,proto3" json:"alive,omitempty"`
	Desc     []byte  `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc,omitempty"`
	Nickname string  `protobuf:"bytes,7,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Num      int64   `protobuf:"varint,8,opt,name=num,proto3" json:"num,omitempty"`
	Flt      float32 `protobuf:"fixed32,9,opt,name=flt,proto3" json:"flt,omitempty"`
	Dbl      float64 `protobuf:"fixed64,10,opt,name=dbl,proto3" json:"dbl,omitempty"`
	Tru      bool    `protobuf:"varint,11,opt,name=tru,proto3" json:"tru,omitempty"`
	Data     []byte  `protobuf:"bytes,12,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BenchLarge) Reset() {
	*x = BenchLarge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchLarge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchLarge) ProtoMessage() {}

func (x *BenchLarge) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchLarge.ProtoReflect.Descriptor instead.
func (*BenchLarge) Descriptor() ([]byte, []int) {
	return file_dist_proto_info_proto_rawDescGZIP(), []int{2}
}

func (x *BenchLarge) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BenchLarge) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *BenchLarge) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BenchLarge) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *BenchLarge) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

func (x *BenchLarge) GetDesc() []byte {
	if x != nil {
		return x.Desc
	}
	return nil
}

func (x *BenchLarge) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *BenchLarge) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *BenchLarge) GetFlt() float32 {
	if x != nil {
		return x.Flt
	}
	return 0
}

func (x *BenchLarge) GetDbl() float64 {
	if x != nil {
		return x.Dbl
	}
	return 0
}

func (x *BenchLarge) GetTru() bool {
	if x != nil {
		return x.Tru
	}
	return false
}

func (x *BenchLarge) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_dist_proto_info_proto protoreflect.FileDescriptor

var file_dist_proto_info_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x36, 0x0a,
	0x0a, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x4d,
	0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x84, 0x02, 0x0a, 0x0a, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x4c,
	0x61, 0x72, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6c, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x66, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x62, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x64, 0x62, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x72, 0x75, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x72, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x74, 0x65, 0x73,
	0x68, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dist_proto_info_proto_rawDescOnce sync.Once
	file_dist_proto_info_proto_rawDescData = file_dist_proto_info_proto_rawDesc
)

func file_dist_proto_info_proto_rawDescGZIP() []byte {
	file_dist_proto_info_proto_rawDescOnce.Do(func() {
		file_dist_proto_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_dist_proto_info_proto_rawDescData)
	})
	return file_dist_proto_info_proto_rawDescData
}

var file_dist_proto_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dist_proto_info_proto_goTypes = []interface{}{
	(*BenchSmall)(nil),  // 0: info.BenchSmall
	(*BenchMedium)(nil), // 1: info.BenchMedium
	(*BenchLarge)(nil),  // 2: info.BenchLarge
}
var file_dist_proto_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dist_proto_info_proto_init() }
func file_dist_proto_info_proto_init() {
	if File_dist_proto_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dist_proto_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchSmall); i {
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
		file_dist_proto_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchMedium); i {
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
		file_dist_proto_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchLarge); i {
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
			RawDescriptor: file_dist_proto_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dist_proto_info_proto_goTypes,
		DependencyIndexes: file_dist_proto_info_proto_depIdxs,
		MessageInfos:      file_dist_proto_info_proto_msgTypes,
	}.Build()
	File_dist_proto_info_proto = out.File
	file_dist_proto_info_proto_rawDesc = nil
	file_dist_proto_info_proto_goTypes = nil
	file_dist_proto_info_proto_depIdxs = nil
}
