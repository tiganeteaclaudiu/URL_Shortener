// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: service.proto

package urlshortener

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

// Represents an empty return
type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

// Represents an accessible URL.
// Used as either input or output of rpc calls which receive/return URLs
type Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full URL (Ex. https://facebook.com)
	Url string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
}

func (x *Url) Reset() {
	*x = Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *Url) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Represents a shortened URL key (a sequence of n length of random letters)
// Is used as suffix for shortened URLs
type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents the key of a shortened URL (EX. localhost:40/<lkjhbalsjdhbl>,  key between <>)
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// input used for SetShortenedUrl rpc method
type SetShortenedUrlInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL to be shortened
	Url string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	// Expiry time in minutes for entry
	// If a 0 value is given, the entry will be persistent.
	ExpiryMinutes int64 `protobuf:"varint,2,opt,name=ExpiryMinutes,proto3" json:"ExpiryMinutes,omitempty"`
}

func (x *SetShortenedUrlInput) Reset() {
	*x = SetShortenedUrlInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetShortenedUrlInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetShortenedUrlInput) ProtoMessage() {}

func (x *SetShortenedUrlInput) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetShortenedUrlInput.ProtoReflect.Descriptor instead.
func (*SetShortenedUrlInput) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *SetShortenedUrlInput) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SetShortenedUrlInput) GetExpiryMinutes() int64 {
	if x != nil {
		return x.ExpiryMinutes
	}
	return 0
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x17,
	0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x22, 0x17, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79,
	0x22, 0x4e, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64,
	0x55, 0x72, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73,
	0x32, 0xab, 0x01, 0x0a, 0x13, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x0a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x72, 0x6c, 0x12, 0x3a, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x72, 0x6c, 0x12,
	0x2d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x1a,
	0x5a, 0x18, 0x6a, 0x6f, 0x62, 0x74, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x72,
	0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_proto_goTypes = []interface{}{
	(*Void)(nil),                 // 0: proto.Void
	(*Url)(nil),                  // 1: proto.Url
	(*Key)(nil),                  // 2: proto.Key
	(*SetShortenedUrlInput)(nil), // 3: proto.SetShortenedUrlInput
}
var file_service_proto_depIdxs = []int32{
	2, // 0: proto.UrlShortenerService.GetShortenedUrl:input_type -> proto.Key
	3, // 1: proto.UrlShortenerService.SetShortenedUrl:input_type -> proto.SetShortenedUrlInput
	2, // 2: proto.UrlShortenerService.DeleteShortenedUrl:input_type -> proto.Key
	1, // 3: proto.UrlShortenerService.GetShortenedUrl:output_type -> proto.Url
	1, // 4: proto.UrlShortenerService.SetShortenedUrl:output_type -> proto.Url
	0, // 5: proto.UrlShortenerService.DeleteShortenedUrl:output_type -> proto.Void
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetShortenedUrlInput); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}