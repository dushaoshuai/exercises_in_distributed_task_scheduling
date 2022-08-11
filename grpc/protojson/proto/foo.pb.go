// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: foo.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TestOneof:
	//	*Bar_Name
	//	*Bar_Empty
	//	*Bar_Null
	TestOneof isBar_TestOneof `protobuf_oneof:"test_oneof"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_foo_proto_rawDescGZIP(), []int{0}
}

func (m *Bar) GetTestOneof() isBar_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (x *Bar) GetName() string {
	if x, ok := x.GetTestOneof().(*Bar_Name); ok {
		return x.Name
	}
	return ""
}

func (x *Bar) GetEmpty() *emptypb.Empty {
	if x, ok := x.GetTestOneof().(*Bar_Empty); ok {
		return x.Empty
	}
	return nil
}

func (x *Bar) GetNull() structpb.NullValue {
	if x, ok := x.GetTestOneof().(*Bar_Null); ok {
		return x.Null
	}
	return structpb.NullValue(0)
}

type isBar_TestOneof interface {
	isBar_TestOneof()
}

type Bar_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type Bar_Empty struct {
	Empty *emptypb.Empty `protobuf:"bytes,2,opt,name=empty,proto3,oneof"`
}

type Bar_Null struct {
	Null structpb.NullValue `protobuf:"varint,3,opt,name=null,proto3,enum=google.protobuf.NullValue,oneof"`
}

func (*Bar_Name) isBar_TestOneof() {}

func (*Bar_Empty) isBar_TestOneof() {}

func (*Bar_Null) isBar_TestOneof() {}

var File_foo_proto protoreflect.FileDescriptor

var file_foo_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x66, 0x6f, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x03,
	0x42, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x75, 0x6c,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x73, 0x68, 0x61, 0x6f, 0x73, 0x68,
	0x75, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x73, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foo_proto_rawDescOnce sync.Once
	file_foo_proto_rawDescData = file_foo_proto_rawDesc
)

func file_foo_proto_rawDescGZIP() []byte {
	file_foo_proto_rawDescOnce.Do(func() {
		file_foo_proto_rawDescData = protoimpl.X.CompressGZIP(file_foo_proto_rawDescData)
	})
	return file_foo_proto_rawDescData
}

var file_foo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_foo_proto_goTypes = []interface{}{
	(*Bar)(nil),             // 0: foo.Bar
	(*emptypb.Empty)(nil),   // 1: google.protobuf.Empty
	(structpb.NullValue)(0), // 2: google.protobuf.NullValue
}
var file_foo_proto_depIdxs = []int32{
	1, // 0: foo.Bar.empty:type_name -> google.protobuf.Empty
	2, // 1: foo.Bar.null:type_name -> google.protobuf.NullValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foo_proto_init() }
func file_foo_proto_init() {
	if File_foo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
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
	file_foo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Bar_Name)(nil),
		(*Bar_Empty)(nil),
		(*Bar_Null)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_foo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foo_proto_goTypes,
		DependencyIndexes: file_foo_proto_depIdxs,
		MessageInfos:      file_foo_proto_msgTypes,
	}.Build()
	File_foo_proto = out.File
	file_foo_proto_rawDesc = nil
	file_foo_proto_goTypes = nil
	file_foo_proto_depIdxs = nil
}
