// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.1
// source: web-teller.proto

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

type APIREQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxType  string            `protobuf:"bytes,1,opt,name=txType,proto3" json:"txType,omitempty"`
	Headers map[string]string `protobuf:"bytes,2,rep,name=Headers,proto3" json:"Headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Params  map[string]string `protobuf:"bytes,3,rep,name=Params,proto3" json:"Params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *APIREQ) Reset() {
	*x = APIREQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_teller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIREQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIREQ) ProtoMessage() {}

func (x *APIREQ) ProtoReflect() protoreflect.Message {
	mi := &file_web_teller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIREQ.ProtoReflect.Descriptor instead.
func (*APIREQ) Descriptor() ([]byte, []int) {
	return file_web_teller_proto_rawDescGZIP(), []int{0}
}

func (x *APIREQ) GetTxType() string {
	if x != nil {
		return x.TxType
	}
	return ""
}

func (x *APIREQ) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *APIREQ) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type APIRES struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []byte            `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	Headers  map[string]string `protobuf:"bytes,2,rep,name=Headers,proto3" json:"Headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *APIRES) Reset() {
	*x = APIRES{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_teller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIRES) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIRES) ProtoMessage() {}

func (x *APIRES) ProtoReflect() protoreflect.Message {
	mi := &file_web_teller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIRES.ProtoReflect.Descriptor instead.
func (*APIRES) Descriptor() ([]byte, []int) {
	return file_web_teller_proto_rawDescGZIP(), []int{1}
}

func (x *APIRES) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *APIRES) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_web_teller_proto protoreflect.FileDescriptor

var file_web_teller_proto_rawDesc = []byte{
	0x0a, 0x10, 0x77, 0x65, 0x62, 0x2d, 0x74, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02, 0x0a, 0x06,
	0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36,
	0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x98, 0x01, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xca, 0x07,
	0x0a, 0x09, 0x57, 0x65, 0x62, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0e, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e,
	0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f,
	0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12,
	0x33, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49,
	0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50,
	0x49, 0x52, 0x45, 0x53, 0x12, 0x32, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x32, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x36, 0x0a, 0x12,
	0x42, 0x75, 0x6c, 0x6b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49,
	0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50,
	0x49, 0x52, 0x45, 0x53, 0x12, 0x33, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x33, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x77,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e,
	0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x35,
	0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50,
	0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x35, 0x0a, 0x11, 0x43, 0x61, 0x73, 0x68, 0x54, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x38, 0x0a, 0x14,
	0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x4e, 0x6f, 0x6d, 0x6f, 0x72, 0x52, 0x65, 0x6b, 0x65,
	0x6e, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x2f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x65, 0x74, 0x61, 0x6b, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x30, 0x0a, 0x0c, 0x52, 0x65, 0x49, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x4d, 0x50, 0x4e, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x33, 0x0a, 0x0f, 0x49, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x53, 0x69, 0x73, 0x6b, 0x6f, 0x68, 0x61, 0x74, 0x12, 0x0f, 0x2e, 0x77,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e,
	0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x30,
	0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x70, 0x6e, 0x41, 0x6c, 0x6c, 0x12, 0x0f,
	0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a,
	0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53,
	0x12, 0x33, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x69, 0x73, 0x6b, 0x6f, 0x70, 0x61,
	0x74, 0x75, 0x68, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50,
	0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x33, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x12, 0x36, 0x0a, 0x12, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45,
	0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52,
	0x45, 0x53, 0x12, 0x36, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x51, 0x1a, 0x0f, 0x2e, 0x77, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x45, 0x53, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x74, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_teller_proto_rawDescOnce sync.Once
	file_web_teller_proto_rawDescData = file_web_teller_proto_rawDesc
)

func file_web_teller_proto_rawDescGZIP() []byte {
	file_web_teller_proto_rawDescOnce.Do(func() {
		file_web_teller_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_teller_proto_rawDescData)
	})
	return file_web_teller_proto_rawDescData
}

var file_web_teller_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_web_teller_proto_goTypes = []interface{}{
	(*APIREQ)(nil), // 0: wtproto.APIREQ
	(*APIRES)(nil), // 1: wtproto.APIRES
	nil,            // 2: wtproto.APIREQ.HeadersEntry
	nil,            // 3: wtproto.APIREQ.ParamsEntry
	nil,            // 4: wtproto.APIRES.HeadersEntry
}
var file_web_teller_proto_depIdxs = []int32{
	2,  // 0: wtproto.APIREQ.Headers:type_name -> wtproto.APIREQ.HeadersEntry
	3,  // 1: wtproto.APIREQ.Params:type_name -> wtproto.APIREQ.ParamsEntry
	4,  // 2: wtproto.APIRES.Headers:type_name -> wtproto.APIRES.HeadersEntry
	0,  // 3: wtproto.WebTeller.Authentication:input_type -> wtproto.APIREQ
	0,  // 4: wtproto.WebTeller.SessionValidate:input_type -> wtproto.APIREQ
	0,  // 5: wtproto.WebTeller.PaymentInquiry:input_type -> wtproto.APIREQ
	0,  // 6: wtproto.WebTeller.PaymentPosting:input_type -> wtproto.APIREQ
	0,  // 7: wtproto.WebTeller.BulkPaymentPosting:input_type -> wtproto.APIREQ
	0,  // 8: wtproto.WebTeller.TransferInquiry:input_type -> wtproto.APIREQ
	0,  // 9: wtproto.WebTeller.TransferPosting:input_type -> wtproto.APIREQ
	0,  // 10: wtproto.WebTeller.TransactionReport:input_type -> wtproto.APIREQ
	0,  // 11: wtproto.WebTeller.CashTellerInquiry:input_type -> wtproto.APIREQ
	0,  // 12: wtproto.WebTeller.InquiryNomorRekening:input_type -> wtproto.APIREQ
	0,  // 13: wtproto.WebTeller.UpdateCetak:input_type -> wtproto.APIREQ
	0,  // 14: wtproto.WebTeller.ReInquiryMPN:input_type -> wtproto.APIREQ
	0,  // 15: wtproto.WebTeller.InquirySiskohat:input_type -> wtproto.APIREQ
	0,  // 16: wtproto.WebTeller.ReportMpnAll:input_type -> wtproto.APIREQ
	0,  // 17: wtproto.WebTeller.LoginSiskopatuh:input_type -> wtproto.APIREQ
	0,  // 18: wtproto.WebTeller.GetBranchTeller:input_type -> wtproto.APIREQ
	0,  // 19: wtproto.WebTeller.DownloadFileReport:input_type -> wtproto.APIREQ
	0,  // 20: wtproto.WebTeller.GenerateFileReport:input_type -> wtproto.APIREQ
	1,  // 21: wtproto.WebTeller.Authentication:output_type -> wtproto.APIRES
	1,  // 22: wtproto.WebTeller.SessionValidate:output_type -> wtproto.APIRES
	1,  // 23: wtproto.WebTeller.PaymentInquiry:output_type -> wtproto.APIRES
	1,  // 24: wtproto.WebTeller.PaymentPosting:output_type -> wtproto.APIRES
	1,  // 25: wtproto.WebTeller.BulkPaymentPosting:output_type -> wtproto.APIRES
	1,  // 26: wtproto.WebTeller.TransferInquiry:output_type -> wtproto.APIRES
	1,  // 27: wtproto.WebTeller.TransferPosting:output_type -> wtproto.APIRES
	1,  // 28: wtproto.WebTeller.TransactionReport:output_type -> wtproto.APIRES
	1,  // 29: wtproto.WebTeller.CashTellerInquiry:output_type -> wtproto.APIRES
	1,  // 30: wtproto.WebTeller.InquiryNomorRekening:output_type -> wtproto.APIRES
	1,  // 31: wtproto.WebTeller.UpdateCetak:output_type -> wtproto.APIRES
	1,  // 32: wtproto.WebTeller.ReInquiryMPN:output_type -> wtproto.APIRES
	1,  // 33: wtproto.WebTeller.InquirySiskohat:output_type -> wtproto.APIRES
	1,  // 34: wtproto.WebTeller.ReportMpnAll:output_type -> wtproto.APIRES
	1,  // 35: wtproto.WebTeller.LoginSiskopatuh:output_type -> wtproto.APIRES
	1,  // 36: wtproto.WebTeller.GetBranchTeller:output_type -> wtproto.APIRES
	1,  // 37: wtproto.WebTeller.DownloadFileReport:output_type -> wtproto.APIRES
	1,  // 38: wtproto.WebTeller.GenerateFileReport:output_type -> wtproto.APIRES
	21, // [21:39] is the sub-list for method output_type
	3,  // [3:21] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_web_teller_proto_init() }
func file_web_teller_proto_init() {
	if File_web_teller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_teller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIREQ); i {
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
		file_web_teller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIRES); i {
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
			RawDescriptor: file_web_teller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web_teller_proto_goTypes,
		DependencyIndexes: file_web_teller_proto_depIdxs,
		MessageInfos:      file_web_teller_proto_msgTypes,
	}.Build()
	File_web_teller_proto = out.File
	file_web_teller_proto_rawDesc = nil
	file_web_teller_proto_goTypes = nil
	file_web_teller_proto_depIdxs = nil
}
