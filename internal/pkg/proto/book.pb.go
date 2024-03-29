// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.21.5
// source: internal/pkg/proto/book.proto

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

type BookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isbn   int32  `protobuf:"varint,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BookRequest) Reset() {
	*x = BookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_proto_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRequest) ProtoMessage() {}

func (x *BookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_proto_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRequest.ProtoReflect.Descriptor instead.
func (*BookRequest) Descriptor() ([]byte, []int) {
	return file_internal_pkg_proto_book_proto_rawDescGZIP(), []int{0}
}

func (x *BookRequest) GetIsbn() int32 {
	if x != nil {
		return x.Isbn
	}
	return 0
}

func (x *BookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *BookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*BookRequest `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BooksResponse) Reset() {
	*x = BooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_proto_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksResponse) ProtoMessage() {}

func (x *BooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_proto_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksResponse.ProtoReflect.Descriptor instead.
func (*BooksResponse) Descriptor() ([]byte, []int) {
	return file_internal_pkg_proto_book_proto_rawDescGZIP(), []int{1}
}

func (x *BooksResponse) GetBooks() []*BookRequest {
	if x != nil {
		return x.Books
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_proto_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_proto_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internal_pkg_proto_book_proto_rawDescGZIP(), []int{2}
}

var File_internal_pkg_proto_book_proto protoreflect.FileDescriptor

var file_internal_pkg_proto_book_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x4d, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x6b, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x12, 0x0b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_proto_book_proto_rawDescOnce sync.Once
	file_internal_pkg_proto_book_proto_rawDescData = file_internal_pkg_proto_book_proto_rawDesc
)

func file_internal_pkg_proto_book_proto_rawDescGZIP() []byte {
	file_internal_pkg_proto_book_proto_rawDescOnce.Do(func() {
		file_internal_pkg_proto_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_proto_book_proto_rawDescData)
	})
	return file_internal_pkg_proto_book_proto_rawDescData
}

var file_internal_pkg_proto_book_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_pkg_proto_book_proto_goTypes = []interface{}{
	(*BookRequest)(nil),   // 0: book.BookRequest
	(*BooksResponse)(nil), // 1: book.BooksResponse
	(*Empty)(nil),         // 2: book.Empty
}
var file_internal_pkg_proto_book_proto_depIdxs = []int32{
	0, // 0: book.BooksResponse.books:type_name -> book.BookRequest
	0, // 1: book.BookService.AddBook:input_type -> book.BookRequest
	2, // 2: book.BookService.ListBooks:input_type -> book.Empty
	2, // 3: book.BookService.AddBook:output_type -> book.Empty
	1, // 4: book.BookService.ListBooks:output_type -> book.BooksResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_pkg_proto_book_proto_init() }
func file_internal_pkg_proto_book_proto_init() {
	if File_internal_pkg_proto_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_proto_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookRequest); i {
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
		file_internal_pkg_proto_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksResponse); i {
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
		file_internal_pkg_proto_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_internal_pkg_proto_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pkg_proto_book_proto_goTypes,
		DependencyIndexes: file_internal_pkg_proto_book_proto_depIdxs,
		MessageInfos:      file_internal_pkg_proto_book_proto_msgTypes,
	}.Build()
	File_internal_pkg_proto_book_proto = out.File
	file_internal_pkg_proto_book_proto_rawDesc = nil
	file_internal_pkg_proto_book_proto_goTypes = nil
	file_internal_pkg_proto_book_proto_depIdxs = nil
}
