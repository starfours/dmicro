// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: secureBody.proto

package securebody

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

type Encrypt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CipherVersion string `protobuf:"bytes,1,opt,name=cipherVersion,proto3" json:"cipherVersion,omitempty"`
	Ciphertext    string `protobuf:"bytes,2,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
}

func (x *Encrypt) Reset() {
	*x = Encrypt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secureBody_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Encrypt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Encrypt) ProtoMessage() {}

func (x *Encrypt) ProtoReflect() protoreflect.Message {
	mi := &file_secureBody_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Encrypt.ProtoReflect.Descriptor instead.
func (*Encrypt) Descriptor() ([]byte, []int) {
	return file_secureBody_proto_rawDescGZIP(), []int{0}
}

func (x *Encrypt) GetCipherVersion() string {
	if x != nil {
		return x.CipherVersion
	}
	return ""
}

func (x *Encrypt) GetCiphertext() string {
	if x != nil {
		return x.Ciphertext
	}
	return ""
}

var File_secureBody_proto protoreflect.FileDescriptor

var file_secureBody_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x4f,
	0x0a, 0x07, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x62, 0x6f, 0x64, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_secureBody_proto_rawDescOnce sync.Once
	file_secureBody_proto_rawDescData = file_secureBody_proto_rawDesc
)

func file_secureBody_proto_rawDescGZIP() []byte {
	file_secureBody_proto_rawDescOnce.Do(func() {
		file_secureBody_proto_rawDescData = protoimpl.X.CompressGZIP(file_secureBody_proto_rawDescData)
	})
	return file_secureBody_proto_rawDescData
}

var file_secureBody_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_secureBody_proto_goTypes = []interface{}{
	(*Encrypt)(nil), // 0: securebody.Encrypt
}
var file_secureBody_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_secureBody_proto_init() }
func file_secureBody_proto_init() {
	if File_secureBody_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_secureBody_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Encrypt); i {
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
			RawDescriptor: file_secureBody_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_secureBody_proto_goTypes,
		DependencyIndexes: file_secureBody_proto_depIdxs,
		MessageInfos:      file_secureBody_proto_msgTypes,
	}.Build()
	File_secureBody_proto = out.File
	file_secureBody_proto_rawDesc = nil
	file_secureBody_proto_goTypes = nil
	file_secureBody_proto_depIdxs = nil
}
