// Code generated by protoc-gen-go. DO NOT EDIT.
// source: web-teller.proto

package wtproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type APIREQ struct {
	TxType               string            `protobuf:"bytes,1,opt,name=txType,proto3" json:"txType,omitempty"`
	Headers              map[string]string `protobuf:"bytes,2,rep,name=Headers,proto3" json:"Headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Params               map[string]string `protobuf:"bytes,3,rep,name=Params,proto3" json:"Params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *APIREQ) Reset()         { *m = APIREQ{} }
func (m *APIREQ) String() string { return proto.CompactTextString(m) }
func (*APIREQ) ProtoMessage()    {}
func (*APIREQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ef744ad4a9dc6f9, []int{0}
}

func (m *APIREQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIREQ.Unmarshal(m, b)
}
func (m *APIREQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIREQ.Marshal(b, m, deterministic)
}
func (m *APIREQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIREQ.Merge(m, src)
}
func (m *APIREQ) XXX_Size() int {
	return xxx_messageInfo_APIREQ.Size(m)
}
func (m *APIREQ) XXX_DiscardUnknown() {
	xxx_messageInfo_APIREQ.DiscardUnknown(m)
}

var xxx_messageInfo_APIREQ proto.InternalMessageInfo

func (m *APIREQ) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

func (m *APIREQ) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *APIREQ) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type APIRES struct {
	Response             []byte            `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	Headers              map[string]string `protobuf:"bytes,2,rep,name=Headers,proto3" json:"Headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *APIRES) Reset()         { *m = APIRES{} }
func (m *APIRES) String() string { return proto.CompactTextString(m) }
func (*APIRES) ProtoMessage()    {}
func (*APIRES) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ef744ad4a9dc6f9, []int{1}
}

func (m *APIRES) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIRES.Unmarshal(m, b)
}
func (m *APIRES) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIRES.Marshal(b, m, deterministic)
}
func (m *APIRES) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIRES.Merge(m, src)
}
func (m *APIRES) XXX_Size() int {
	return xxx_messageInfo_APIRES.Size(m)
}
func (m *APIRES) XXX_DiscardUnknown() {
	xxx_messageInfo_APIRES.DiscardUnknown(m)
}

var xxx_messageInfo_APIRES proto.InternalMessageInfo

func (m *APIRES) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *APIRES) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*APIREQ)(nil), "wtproto.APIREQ")
	proto.RegisterMapType((map[string]string)(nil), "wtproto.APIREQ.HeadersEntry")
	proto.RegisterMapType((map[string]string)(nil), "wtproto.APIREQ.ParamsEntry")
	proto.RegisterType((*APIRES)(nil), "wtproto.APIRES")
	proto.RegisterMapType((map[string]string)(nil), "wtproto.APIRES.HeadersEntry")
}

func init() { proto.RegisterFile("web-teller.proto", fileDescriptor_0ef744ad4a9dc6f9) }

var fileDescriptor_0ef744ad4a9dc6f9 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xd3, 0x36, 0x16, 0x19, 0x08, 0xe0, 0xc6, 0x98, 0xa6, 0x7a, 0x20, 0x9c, 0xb8, 0xd8,
	0x43, 0x89, 0x46, 0xb9, 0x11, 0x43, 0x22, 0xb7, 0xda, 0x12, 0x3d, 0x2f, 0x30, 0x4a, 0x63, 0xd9,
	0xad, 0xbb, 0x8b, 0xd8, 0xbb, 0x0f, 0xe1, 0x83, 0xf8, 0x80, 0x86, 0xb6, 0x1a, 0xd2, 0x98, 0xb8,
	0xc4, 0xdb, 0xfc, 0x33, 0xf3, 0xcd, 0xcc, 0xfe, 0x0b, 0x9d, 0x0d, 0xce, 0xce, 0x15, 0x26, 0x09,
	0x0a, 0x2f, 0x15, 0x5c, 0x71, 0x52, 0xdb, 0xa8, 0x3c, 0xe8, 0xbd, 0x9b, 0x60, 0x8f, 0x82, 0x49,
	0x38, 0xbe, 0x23, 0x27, 0x60, 0xab, 0xb7, 0x69, 0x96, 0xa2, 0x63, 0x74, 0x8d, 0x7e, 0x3d, 0x2c,
	0x15, 0xb9, 0x84, 0xda, 0x2d, 0xd2, 0x05, 0x0a, 0xe9, 0x98, 0x5d, 0xab, 0xdf, 0xf0, 0xcf, 0xbc,
	0x92, 0xf6, 0x0a, 0xd2, 0x2b, 0xcb, 0x63, 0xa6, 0x44, 0x16, 0x7e, 0x37, 0x93, 0x01, 0xd8, 0x01,
	0x15, 0x74, 0x25, 0x1d, 0x2b, 0xc7, 0x4e, 0xab, 0x58, 0x51, 0x2d, 0xa8, 0xb2, 0xd5, 0x1d, 0x42,
	0x73, 0x77, 0x1a, 0xe9, 0x80, 0xf5, 0x8c, 0x59, 0x79, 0xd1, 0x36, 0x24, 0xc7, 0x70, 0xf0, 0x4a,
	0x93, 0x35, 0x3a, 0x66, 0x9e, 0x2b, 0xc4, 0xd0, 0xbc, 0x32, 0xdc, 0x6b, 0x68, 0xec, 0x8c, 0xdc,
	0x07, 0xed, 0x7d, 0x18, 0xa5, 0x0d, 0x11, 0x71, 0xe1, 0x30, 0x44, 0x99, 0x72, 0x26, 0x0b, 0x23,
	0x9a, 0xe1, 0x8f, 0xfe, 0xd3, 0x8a, 0xe8, 0x77, 0x2b, 0xfe, 0xf3, 0x2a, 0xff, 0xd3, 0x82, 0xfa,
	0x03, 0xce, 0xa6, 0xf9, 0xf7, 0x11, 0x1f, 0x5a, 0xa3, 0xb5, 0x5a, 0x22, 0x53, 0xf1, 0x9c, 0xaa,
	0x98, 0x33, 0xd2, 0xae, 0xd8, 0xea, 0x56, 0x12, 0x11, 0x19, 0x40, 0x3b, 0x42, 0x29, 0x63, 0xce,
	0xee, 0x69, 0x12, 0x2f, 0xa8, 0x42, 0x0d, 0xc8, 0x87, 0x56, 0x40, 0xb3, 0x15, 0x32, 0x35, 0x61,
	0x2f, 0xeb, 0x58, 0x64, 0x7b, 0x31, 0x01, 0x97, 0x2a, 0x66, 0x4f, 0x7a, 0xc7, 0x4d, 0x05, 0x65,
	0xf2, 0x11, 0x85, 0xfe, 0xa2, 0x1d, 0x48, 0x7f, 0xd3, 0x05, 0x1c, 0xe5, 0x10, 0x9d, 0x6f, 0x7d,
	0x0b, 0x31, 0xe5, 0x42, 0xe9, 0x61, 0x37, 0x54, 0x2e, 0x0b, 0xff, 0xb5, 0x4f, 0x9c, 0xd9, 0xb9,
	0x1a, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xbf, 0x3e, 0xb5, 0x7c, 0x03, 0x00, 0x00,
}
