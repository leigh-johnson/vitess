// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tableacl.proto

package tableacl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// TableGroupSpec defines ACLs for a group of tables.
type TableGroupSpec struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// either tables or a table name prefixes (if it ends in a %)
	TableNamesOrPrefixes []string `protobuf:"bytes,2,rep,name=table_names_or_prefixes,json=tableNamesOrPrefixes,proto3" json:"table_names_or_prefixes,omitempty"`
	Readers              []string `protobuf:"bytes,3,rep,name=readers,proto3" json:"readers,omitempty"`
	Writers              []string `protobuf:"bytes,4,rep,name=writers,proto3" json:"writers,omitempty"`
	Admins               []string `protobuf:"bytes,5,rep,name=admins,proto3" json:"admins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableGroupSpec) Reset()         { *m = TableGroupSpec{} }
func (m *TableGroupSpec) String() string { return proto.CompactTextString(m) }
func (*TableGroupSpec) ProtoMessage()    {}
func (*TableGroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d0bedb248a1632e, []int{0}
}

func (m *TableGroupSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableGroupSpec.Unmarshal(m, b)
}
func (m *TableGroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableGroupSpec.Marshal(b, m, deterministic)
}
func (m *TableGroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableGroupSpec.Merge(m, src)
}
func (m *TableGroupSpec) XXX_Size() int {
	return xxx_messageInfo_TableGroupSpec.Size(m)
}
func (m *TableGroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TableGroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TableGroupSpec proto.InternalMessageInfo

func (m *TableGroupSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableGroupSpec) GetTableNamesOrPrefixes() []string {
	if m != nil {
		return m.TableNamesOrPrefixes
	}
	return nil
}

func (m *TableGroupSpec) GetReaders() []string {
	if m != nil {
		return m.Readers
	}
	return nil
}

func (m *TableGroupSpec) GetWriters() []string {
	if m != nil {
		return m.Writers
	}
	return nil
}

func (m *TableGroupSpec) GetAdmins() []string {
	if m != nil {
		return m.Admins
	}
	return nil
}

type Config struct {
	TableGroups          []*TableGroupSpec `protobuf:"bytes,1,rep,name=table_groups,json=tableGroups,proto3" json:"table_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d0bedb248a1632e, []int{1}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetTableGroups() []*TableGroupSpec {
	if m != nil {
		return m.TableGroups
	}
	return nil
}

func init() {
	proto.RegisterType((*TableGroupSpec)(nil), "tableacl.TableGroupSpec")
	proto.RegisterType((*Config)(nil), "tableacl.Config")
}

func init() { proto.RegisterFile("tableacl.proto", fileDescriptor_7d0bedb248a1632e) }

var fileDescriptor_7d0bedb248a1632e = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0x9d, 0xd5, 0xbd, 0xc9, 0x0e, 0x41, 0x34, 0xc7, 0x32, 0x10, 0x7b, 0x6a, 0x40,
	0xf1, 0xe4, 0x4d, 0x11, 0x6f, 0x2a, 0xd5, 0x93, 0x97, 0x92, 0x6d, 0x6f, 0x25, 0xb0, 0x35, 0xe1,
	0xbd, 0x38, 0xfd, 0x8f, 0xfc, 0x37, 0x25, 0x69, 0x3b, 0xf0, 0xf6, 0xfd, 0xf8, 0x25, 0xe1, 0xfb,
	0x02, 0xf3, 0x60, 0x96, 0x5b, 0x34, 0xab, 0x6d, 0xe5, 0xc9, 0x05, 0x27, 0x4f, 0x47, 0x5e, 0xfc,
	0x0a, 0x98, 0x7f, 0x44, 0x78, 0x26, 0xf7, 0xe5, 0xdf, 0x3d, 0xae, 0xa4, 0x84, 0x49, 0x67, 0x76,
	0xa8, 0x44, 0x21, 0xca, 0x69, 0x9d, 0xb2, 0xbc, 0x83, 0xcb, 0x74, 0xa5, 0x89, 0xc4, 0x8d, 0xa3,
	0xc6, 0x13, 0x6e, 0xec, 0x0f, 0xb2, 0x3a, 0x2a, 0xb2, 0x72, 0x5a, 0x9f, 0x27, 0xfd, 0x12, 0xed,
	0x2b, 0xbd, 0x0d, 0x4e, 0x2a, 0x38, 0x21, 0x34, 0x6b, 0x24, 0x56, 0x59, 0x3a, 0x36, 0x62, 0x34,
	0xdf, 0x64, 0x43, 0x34, 0x93, 0xde, 0x0c, 0x28, 0x2f, 0x20, 0x37, 0xeb, 0x9d, 0xed, 0x58, 0x1d,
	0x27, 0x31, 0xd0, 0xe2, 0x09, 0xf2, 0x47, 0xd7, 0x6d, 0x6c, 0x2b, 0xef, 0xe1, 0xac, 0x2f, 0xd3,
	0xc6, 0xce, 0xac, 0x44, 0x91, 0x95, 0xb3, 0x1b, 0x55, 0x1d, 0x46, 0xfe, 0x1f, 0x54, 0xcf, 0xc2,
	0x81, 0xf9, 0xe1, 0xfa, 0xf3, 0x6a, 0x6f, 0x03, 0x32, 0x57, 0xd6, 0xe9, 0x3e, 0xe9, 0xd6, 0xe9,
	0x7d, 0xd0, 0xe9, 0x6b, 0xf4, 0xf8, 0xc8, 0x32, 0x4f, 0x7c, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff,
	0x09, 0x82, 0xf5, 0x82, 0x3c, 0x01, 0x00, 0x00,
}
