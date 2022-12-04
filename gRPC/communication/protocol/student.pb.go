// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protocol/student.proto

package protocol

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

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_protocol_student_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type SearchByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SearchByID) Reset() {
	*x = SearchByID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchByID) ProtoMessage() {}

func (x *SearchByID) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchByID.ProtoReflect.Descriptor instead.
func (*SearchByID) Descriptor() ([]byte, []int) {
	return file_protocol_student_proto_rawDescGZIP(), []int{1}
}

func (x *SearchByID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SearchByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SearchByName) Reset() {
	*x = SearchByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchByName) ProtoMessage() {}

func (x *SearchByName) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchByName.ProtoReflect.Descriptor instead.
func (*SearchByName) Descriptor() ([]byte, []int) {
	return file_protocol_student_proto_rawDescGZIP(), []int{2}
}

func (x *SearchByName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_protocol_student_proto protoreflect.FileDescriptor

var file_protocol_student_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x6d, 0x0a, 0x0e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x0b, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x08, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0d, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x08, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_student_proto_rawDescOnce sync.Once
	file_protocol_student_proto_rawDescData = file_protocol_student_proto_rawDesc
)

func file_protocol_student_proto_rawDescGZIP() []byte {
	file_protocol_student_proto_rawDescOnce.Do(func() {
		file_protocol_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_student_proto_rawDescData)
	})
	return file_protocol_student_proto_rawDescData
}

var file_protocol_student_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocol_student_proto_goTypes = []interface{}{
	(*Student)(nil),      // 0: Student
	(*SearchByID)(nil),   // 1: SearchByID
	(*SearchByName)(nil), // 2: SearchByName
}
var file_protocol_student_proto_depIdxs = []int32{
	1, // 0: StudentService.GetStudentByID:input_type -> SearchByID
	2, // 1: StudentService.GetStudentsByName:input_type -> SearchByName
	0, // 2: StudentService.GetStudentByID:output_type -> Student
	0, // 3: StudentService.GetStudentsByName:output_type -> Student
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_student_proto_init() }
func file_protocol_student_proto_init() {
	if File_protocol_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_protocol_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchByID); i {
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
		file_protocol_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchByName); i {
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
			RawDescriptor: file_protocol_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_student_proto_goTypes,
		DependencyIndexes: file_protocol_student_proto_depIdxs,
		MessageInfos:      file_protocol_student_proto_msgTypes,
	}.Build()
	File_protocol_student_proto = out.File
	file_protocol_student_proto_rawDesc = nil
	file_protocol_student_proto_goTypes = nil
	file_protocol_student_proto_depIdxs = nil
}
