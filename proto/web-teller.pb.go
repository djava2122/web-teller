// Code generated by protoc-gen-go. DO NOT EDIT.
// source: web-teller.proto

package proto

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
	proto.RegisterType((*APIREQ)(nil), "proto.APIREQ")
	proto.RegisterMapType((map[string]string)(nil), "proto.APIREQ.HeadersEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.APIREQ.ParamsEntry")
	proto.RegisterType((*APIRES)(nil), "proto.APIRES")
	proto.RegisterMapType((map[string]string)(nil), "proto.APIRES.HeadersEntry")
}

func init() { proto.RegisterFile("web-teller.proto", fileDescriptor_0ef744ad4a9dc6f9) }

var fileDescriptor_0ef744ad4a9dc6f9 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x95, 0x44, 0x0d, 0xf4, 0x6f, 0x29, 0x95, 0x85, 0x50, 0xc8, 0x54, 0x75, 0xea, 0x42,
	0x10, 0x85, 0x01, 0xba, 0x75, 0xa8, 0x44, 0xb7, 0x90, 0x54, 0x30, 0x3b, 0xf4, 0x17, 0x58, 0xa4,
	0x76, 0xb0, 0x1d, 0x20, 0x1b, 0x4f, 0xc1, 0x53, 0xf1, 0x50, 0x88, 0xd8, 0xa0, 0x46, 0x42, 0x14,
	0xc4, 0x64, 0x9f, 0xff, 0xfb, 0x4e, 0xf6, 0x19, 0xfa, 0x4f, 0x98, 0x1d, 0x6a, 0xcc, 0x73, 0x94,
	0x51, 0x21, 0x85, 0x16, 0xa4, 0x55, 0x2f, 0xc3, 0x17, 0x17, 0xfc, 0x69, 0x3c, 0x4f, 0x66, 0x97,
	0x64, 0x1f, 0x7c, 0xfd, 0xbc, 0xa8, 0x0a, 0x0c, 0x9c, 0x81, 0x33, 0x6a, 0x27, 0x56, 0x91, 0x53,
	0xd8, 0xba, 0x40, 0xba, 0x44, 0xa9, 0x02, 0x77, 0xe0, 0x8d, 0x3a, 0xe3, 0xd0, 0x44, 0x44, 0x86,
	0x8b, 0xec, 0x70, 0xc6, 0xb5, 0xac, 0x92, 0x4f, 0x2b, 0x39, 0x06, 0x3f, 0xa6, 0x92, 0xae, 0x54,
	0xe0, 0xd5, 0xd0, 0x41, 0x13, 0x32, 0x33, 0xc3, 0x58, 0x63, 0x38, 0x81, 0xee, 0x7a, 0x16, 0xe9,
	0x83, 0x77, 0x8f, 0x95, 0xbd, 0xcd, 0xc7, 0x96, 0xec, 0x41, 0xeb, 0x91, 0xe6, 0x25, 0x06, 0x6e,
	0x7d, 0x66, 0xc4, 0xc4, 0x3d, 0x73, 0xc2, 0x73, 0xe8, 0xac, 0x45, 0xfe, 0x05, 0x1d, 0xbe, 0x3a,
	0xb6, 0x82, 0x94, 0x84, 0xb0, 0x9d, 0xa0, 0x2a, 0x04, 0x57, 0xa6, 0x84, 0x6e, 0xf2, 0xa5, 0x37,
	0xd4, 0x90, 0x7e, 0x5f, 0xc3, 0x7f, 0xde, 0x34, 0x7e, 0x73, 0xa0, 0x7d, 0x8d, 0xd9, 0xa2, 0xfe,
	0x36, 0x12, 0x41, 0x6f, 0x5a, 0xea, 0x3b, 0xe4, 0x9a, 0xdd, 0x50, 0xcd, 0x04, 0x27, 0x3b, 0x8d,
	0x4a, 0xc3, 0x86, 0x4c, 0xc9, 0x11, 0xec, 0xa6, 0xa8, 0x14, 0x13, 0xfc, 0x8a, 0xe6, 0x6c, 0x49,
	0x35, 0x6e, 0x00, 0x22, 0xe8, 0xc5, 0xb4, 0x5a, 0x21, 0xd7, 0x73, 0xfe, 0x50, 0x32, 0x59, 0xfd,
	0xda, 0x1f, 0x0b, 0xa5, 0x19, 0xbf, 0xfd, 0xd9, 0x9f, 0xf9, 0xb5, 0x3a, 0x79, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0xf8, 0xe8, 0xff, 0x03, 0x8c, 0x02, 0x00, 0x00,
}