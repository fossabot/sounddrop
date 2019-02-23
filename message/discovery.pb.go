// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discovery.proto

package message

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

type Announce struct {
	ServiceNumber        uint32   `protobuf:"varint,1,opt,name=service_number,json=serviceNumber,proto3" json:"service_number,omitempty"`
	DeviceName           string   `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Announce) Reset()         { *m = Announce{} }
func (m *Announce) String() string { return proto.CompactTextString(m) }
func (*Announce) ProtoMessage()    {}
func (*Announce) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{0}
}

func (m *Announce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Announce.Unmarshal(m, b)
}
func (m *Announce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Announce.Marshal(b, m, deterministic)
}
func (m *Announce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announce.Merge(m, src)
}
func (m *Announce) XXX_Size() int {
	return xxx_messageInfo_Announce.Size(m)
}
func (m *Announce) XXX_DiscardUnknown() {
	xxx_messageInfo_Announce.DiscardUnknown(m)
}

var xxx_messageInfo_Announce proto.InternalMessageInfo

func (m *Announce) GetServiceNumber() uint32 {
	if m != nil {
		return m.ServiceNumber
	}
	return 0
}

func (m *Announce) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func init() {
	proto.RegisterType((*Announce)(nil), "message.Announce")
}

func init() { proto.RegisterFile("discovery.proto", fileDescriptor_1e7ff60feb39c8d0) }

var fileDescriptor_1e7ff60feb39c8d0 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xc9, 0x2c, 0x4e,
	0xce, 0x2f, 0x4b, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x55, 0x0a, 0xe2, 0xe2, 0x70, 0xcc, 0xcb, 0xcb, 0x2f, 0xcd, 0x4b, 0x4e,
	0x15, 0x52, 0xe5, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x8d, 0xcf, 0x2b, 0xcd, 0x4d,
	0x4a, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0xe2, 0x85, 0x8a, 0xfa, 0x81, 0x05, 0x85,
	0xe4, 0xb9, 0xb8, 0x53, 0x52, 0x21, 0xaa, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38,
	0x83, 0xb8, 0x20, 0x42, 0x7e, 0x89, 0xb9, 0xa9, 0x49, 0x6c, 0x60, 0x3b, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb6, 0xa2, 0xb2, 0x22, 0x76, 0x00, 0x00, 0x00,
}