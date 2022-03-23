// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: bridge/checkpoint/transport/proto/transport.proto

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

type SignedMessage_MsgType int32

const (
	SignedMessage_CHECKPOINT SignedMessage_MsgType = 0
	SignedMessage_ACK        SignedMessage_MsgType = 1
	SignedMessage_NOACK      SignedMessage_MsgType = 2
)

// Enum value maps for SignedMessage_MsgType.
var (
	SignedMessage_MsgType_name = map[int32]string{
		0: "CHECKPOINT",
		1: "ACK",
		2: "NOACK",
	}
	SignedMessage_MsgType_value = map[string]int32{
		"CHECKPOINT": 0,
		"ACK":        1,
		"NOACK":      2,
	}
)

func (x SignedMessage_MsgType) Enum() *SignedMessage_MsgType {
	p := new(SignedMessage_MsgType)
	*p = x
	return p
}

func (x SignedMessage_MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignedMessage_MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_bridge_checkpoint_transport_proto_transport_proto_enumTypes[0].Descriptor()
}

func (SignedMessage_MsgType) Type() protoreflect.EnumType {
	return &file_bridge_checkpoint_transport_proto_transport_proto_enumTypes[0]
}

func (x SignedMessage_MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignedMessage_MsgType.Descriptor instead.
func (SignedMessage_MsgType) EnumDescriptor() ([]byte, []int) {
	return file_bridge_checkpoint_transport_proto_transport_proto_rawDescGZIP(), []int{0, 0}
}

type SignedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  type is the type of message
	Type SignedMessage_MsgType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.SignedMessage_MsgType" json:"type,omitempty"`
	// payload is message body
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// validator signature of the message
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedMessage) Reset() {
	*x = SignedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_checkpoint_transport_proto_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedMessage) ProtoMessage() {}

func (x *SignedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_checkpoint_transport_proto_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedMessage.ProtoReflect.Descriptor instead.
func (*SignedMessage) Descriptor() ([]byte, []int) {
	return file_bridge_checkpoint_transport_proto_transport_proto_rawDescGZIP(), []int{0}
}

func (x *SignedMessage) GetType() SignedMessage_MsgType {
	if x != nil {
		return x.Type
	}
	return SignedMessage_CHECKPOINT
}

func (x *SignedMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SignedMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_bridge_checkpoint_transport_proto_transport_proto protoreflect.FileDescriptor

var file_bridge_checkpoint_transport_proto_transport_proto_rawDesc = []byte{
	0x0a, 0x31, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0xa5, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x2d, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x43,
	0x48, 0x45, 0x43, 0x4b, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x43, 0x4b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x4f, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x42,
	0x24, 0x5a, 0x22, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bridge_checkpoint_transport_proto_transport_proto_rawDescOnce sync.Once
	file_bridge_checkpoint_transport_proto_transport_proto_rawDescData = file_bridge_checkpoint_transport_proto_transport_proto_rawDesc
)

func file_bridge_checkpoint_transport_proto_transport_proto_rawDescGZIP() []byte {
	file_bridge_checkpoint_transport_proto_transport_proto_rawDescOnce.Do(func() {
		file_bridge_checkpoint_transport_proto_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_bridge_checkpoint_transport_proto_transport_proto_rawDescData)
	})
	return file_bridge_checkpoint_transport_proto_transport_proto_rawDescData
}

var file_bridge_checkpoint_transport_proto_transport_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bridge_checkpoint_transport_proto_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bridge_checkpoint_transport_proto_transport_proto_goTypes = []interface{}{
	(SignedMessage_MsgType)(0), // 0: v1.SignedMessage.MsgType
	(*SignedMessage)(nil),      // 1: v1.SignedMessage
}
var file_bridge_checkpoint_transport_proto_transport_proto_depIdxs = []int32{
	0, // 0: v1.SignedMessage.type:type_name -> v1.SignedMessage.MsgType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bridge_checkpoint_transport_proto_transport_proto_init() }
func file_bridge_checkpoint_transport_proto_transport_proto_init() {
	if File_bridge_checkpoint_transport_proto_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bridge_checkpoint_transport_proto_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedMessage); i {
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
			RawDescriptor: file_bridge_checkpoint_transport_proto_transport_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bridge_checkpoint_transport_proto_transport_proto_goTypes,
		DependencyIndexes: file_bridge_checkpoint_transport_proto_transport_proto_depIdxs,
		EnumInfos:         file_bridge_checkpoint_transport_proto_transport_proto_enumTypes,
		MessageInfos:      file_bridge_checkpoint_transport_proto_transport_proto_msgTypes,
	}.Build()
	File_bridge_checkpoint_transport_proto_transport_proto = out.File
	file_bridge_checkpoint_transport_proto_transport_proto_rawDesc = nil
	file_bridge_checkpoint_transport_proto_transport_proto_goTypes = nil
	file_bridge_checkpoint_transport_proto_transport_proto_depIdxs = nil
}