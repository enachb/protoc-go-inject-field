// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Human
	Robot
	Alien
*/
package pb

import "github.com/golang/protobuf/proto"
import "fmt"
import "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// @inject_field: age int
// @inject_field: spouse *Human
// @inject_field: IgnoreMe
type Human struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Human) Reset()                    { *m = Human{} }
func (m *Human) String() string            { return proto.CompactTextString(m) }
func (*Human) ProtoMessage()               {}
func (*Human) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// @inject_field: model string
type Robot struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Robot) Reset()                    { *m = Robot{} }
func (m *Robot) String() string            { return proto.CompactTextString(m) }
func (*Robot) ProtoMessage()               {}
func (*Robot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Alien struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Alien) Reset()                    { *m = Alien{} }
func (m *Alien) String() string            { return proto.CompactTextString(m) }
func (*Alien) ProtoMessage()               {}
func (*Alien) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Human)(nil), "pb.Human")
	proto.RegisterType((*Robot)(nil), "pb.Robot")
	proto.RegisterType((*Alien)(nil), "pb.Alien")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe6, 0x62, 0xf5, 0x28,
	0xcd, 0x4d, 0xcc, 0x13, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb3, 0x41, 0x92, 0x41, 0xf9, 0x49, 0xf9, 0x25, 0xb8, 0x24, 0x1d, 0x73, 0x32,
	0x53, 0xb1, 0xea, 0x4c, 0x62, 0x03, 0xdb, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x9d,
	0x3a, 0xe6, 0x6f, 0x00, 0x00, 0x00,
}