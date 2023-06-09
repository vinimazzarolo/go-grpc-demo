// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/greet.proto

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

type NoParam struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoParam) Reset()         { *m = NoParam{} }
func (m *NoParam) String() string { return proto.CompactTextString(m) }
func (*NoParam) ProtoMessage()    {}
func (*NoParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{0}
}

func (m *NoParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoParam.Unmarshal(m, b)
}
func (m *NoParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoParam.Marshal(b, m, deterministic)
}
func (m *NoParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoParam.Merge(m, src)
}
func (m *NoParam) XXX_Size() int {
	return xxx_messageInfo_NoParam.Size(m)
}
func (m *NoParam) XXX_DiscardUnknown() {
	xxx_messageInfo_NoParam.DiscardUnknown(m)
}

var xxx_messageInfo_NoParam proto.InternalMessageInfo

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{1}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{2}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type NamesList struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamesList) Reset()         { *m = NamesList{} }
func (m *NamesList) String() string { return proto.CompactTextString(m) }
func (*NamesList) ProtoMessage()    {}
func (*NamesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{3}
}

func (m *NamesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamesList.Unmarshal(m, b)
}
func (m *NamesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamesList.Marshal(b, m, deterministic)
}
func (m *NamesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamesList.Merge(m, src)
}
func (m *NamesList) XXX_Size() int {
	return xxx_messageInfo_NamesList.Size(m)
}
func (m *NamesList) XXX_DiscardUnknown() {
	xxx_messageInfo_NamesList.DiscardUnknown(m)
}

var xxx_messageInfo_NamesList proto.InternalMessageInfo

func (m *NamesList) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type MessagesList struct {
	Messages             []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagesList) Reset()         { *m = MessagesList{} }
func (m *MessagesList) String() string { return proto.CompactTextString(m) }
func (*MessagesList) ProtoMessage()    {}
func (*MessagesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{4}
}

func (m *MessagesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagesList.Unmarshal(m, b)
}
func (m *MessagesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagesList.Marshal(b, m, deterministic)
}
func (m *MessagesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagesList.Merge(m, src)
}
func (m *MessagesList) XXX_Size() int {
	return xxx_messageInfo_MessagesList.Size(m)
}
func (m *MessagesList) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagesList.DiscardUnknown(m)
}

var xxx_messageInfo_MessagesList proto.InternalMessageInfo

func (m *MessagesList) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*NoParam)(nil), "greet_service.NoParam")
	proto.RegisterType((*HelloRequest)(nil), "greet_service.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "greet_service.HelloResponse")
	proto.RegisterType((*NamesList)(nil), "greet_service.NamesList")
	proto.RegisterType((*MessagesList)(nil), "greet_service.MessagesList")
}

func init() {
	proto.RegisterFile("proto/greet.proto", fileDescriptor_95ca2ad3f55a58e3)
}

var fileDescriptor_95ca2ad3f55a58e3 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0xff, 0xba, 0x1e, 0xba, 0x0b, 0x83, 0x48, 0xe9, 0x44, 0x66, 0xae, 0xaa, 0x17,
	0xdd, 0xd0, 0x17, 0x90, 0x79, 0xa1, 0x17, 0x3a, 0xa4, 0xbd, 0x10, 0xbd, 0x19, 0x71, 0x3b, 0x94,
	0x40, 0xdb, 0xcc, 0x24, 0x0a, 0x3e, 0xb0, 0xef, 0x21, 0x4d, 0x1b, 0xb5, 0x45, 0xc6, 0xae, 0x72,
	0xbe, 0xe4, 0x3b, 0x3f, 0xbe, 0x73, 0x08, 0x1c, 0xae, 0x95, 0x34, 0x72, 0x92, 0x2b, 0x44, 0x93,
	0xd8, 0x9a, 0x0e, 0xad, 0x58, 0x68, 0x54, 0x1f, 0x62, 0x89, 0xcc, 0x07, 0x6f, 0x2e, 0x1f, 0xb9,
	0xe2, 0x25, 0x63, 0x10, 0xdc, 0x61, 0x51, 0xc8, 0x14, 0xdf, 0xde, 0x51, 0x1b, 0x4a, 0x61, 0xaf,
	0xe2, 0x25, 0x86, 0x64, 0x4c, 0x62, 0x3f, 0xb5, 0x35, 0x3b, 0x87, 0x61, 0xeb, 0xd1, 0x6b, 0x59,
	0x69, 0xa4, 0x21, 0x78, 0x25, 0x6a, 0xcd, 0x73, 0xe7, 0x73, 0x92, 0x9d, 0x81, 0x3f, 0xe7, 0x25,
	0xea, 0x7b, 0xa1, 0x0d, 0x3d, 0x82, 0xfd, 0xba, 0x5f, 0x87, 0x64, 0xbc, 0x1b, 0xfb, 0x69, 0x23,
	0xd8, 0x05, 0x04, 0x0f, 0x8d, 0xbb, 0x71, 0x45, 0x30, 0x68, 0xbb, 0x9d, 0xf1, 0x47, 0x5f, 0x7e,
	0xed, 0x40, 0x70, 0x5b, 0x47, 0xcf, 0x9a, 0xe4, 0xf4, 0x1a, 0x06, 0x19, 0xff, 0xb4, 0x69, 0xe8,
	0x71, 0xd2, 0x99, 0x2a, 0x69, 0x47, 0x8a, 0x4e, 0x7a, 0xf7, 0xdd, 0xec, 0x4f, 0x30, 0x72, 0x84,
	0x1a, 0x8a, 0x2a, 0x13, 0x2b, 0xcc, 0x8c, 0x42, 0x5e, 0x8a, 0x2a, 0xa7, 0x61, 0x1f, 0xea, 0xa6,
	0xd9, 0x8c, 0x9d, 0x12, 0xfa, 0xfc, 0x0b, 0xbe, 0x29, 0x04, 0x56, 0xa6, 0x0b, 0x1e, 0xfd, 0xdf,
	0x6e, 0xb7, 0x1e, 0xf5, 0x1f, 0xff, 0x2e, 0x28, 0x26, 0x74, 0x01, 0xa7, 0x0e, 0x3d, 0x13, 0x2b,
	0xa1, 0x70, 0x69, 0x84, 0xac, 0x78, 0xb1, 0x25, 0x7d, 0x63, 0xf2, 0x98, 0x4c, 0xc9, 0xcc, 0x7f,
	0xf1, 0x92, 0x89, 0xfd, 0x2a, 0xaf, 0x07, 0xf6, 0xb8, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x00,
	0xf6, 0x7e, 0x81, 0x46, 0x02, 0x00, 0x00,
}
