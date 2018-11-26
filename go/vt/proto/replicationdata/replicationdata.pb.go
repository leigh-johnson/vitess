// Code generated by protoc-gen-go. DO NOT EDIT.
// source: replicationdata.proto

package replicationdata // import "vitess.io/vitess/go/vt/proto/replicationdata"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Status is the replication status for MySQL (returned by 'show slave status'
// and parsed into a Position and fields).
type Status struct {
	Position             string   `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	SlaveIoRunning       bool     `protobuf:"varint,2,opt,name=slave_io_running,json=slaveIoRunning,proto3" json:"slave_io_running,omitempty"`
	SlaveSqlRunning      bool     `protobuf:"varint,3,opt,name=slave_sql_running,json=slaveSqlRunning,proto3" json:"slave_sql_running,omitempty"`
	SecondsBehindMaster  uint32   `protobuf:"varint,4,opt,name=seconds_behind_master,json=secondsBehindMaster,proto3" json:"seconds_behind_master,omitempty"`
	MasterHost           string   `protobuf:"bytes,5,opt,name=master_host,json=masterHost,proto3" json:"master_host,omitempty"`
	MasterPort           int32    `protobuf:"varint,6,opt,name=master_port,json=masterPort,proto3" json:"master_port,omitempty"`
	MasterConnectRetry   int32    `protobuf:"varint,7,opt,name=master_connect_retry,json=masterConnectRetry,proto3" json:"master_connect_retry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_replicationdata_535db925ee5677f7, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Status) GetSlaveIoRunning() bool {
	if m != nil {
		return m.SlaveIoRunning
	}
	return false
}

func (m *Status) GetSlaveSqlRunning() bool {
	if m != nil {
		return m.SlaveSqlRunning
	}
	return false
}

func (m *Status) GetSecondsBehindMaster() uint32 {
	if m != nil {
		return m.SecondsBehindMaster
	}
	return 0
}

func (m *Status) GetMasterHost() string {
	if m != nil {
		return m.MasterHost
	}
	return ""
}

func (m *Status) GetMasterPort() int32 {
	if m != nil {
		return m.MasterPort
	}
	return 0
}

func (m *Status) GetMasterConnectRetry() int32 {
	if m != nil {
		return m.MasterConnectRetry
	}
	return 0
}

func init() {
	proto.RegisterType((*Status)(nil), "replicationdata.Status")
}

func init() {
	proto.RegisterFile("replicationdata.proto", fileDescriptor_replicationdata_535db925ee5677f7)
}

var fileDescriptor_replicationdata_535db925ee5677f7 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0x6a, 0xd7, 0x1a, 0xd1, 0x6a, 0xb4, 0x10, 0xbc, 0xb8, 0x78, 0x5a, 0x44, 0x36,
	0xa2, 0x6f, 0x50, 0x2f, 0x7a, 0x10, 0x24, 0xbd, 0x79, 0x09, 0xe9, 0x6e, 0x68, 0x03, 0x6b, 0x66,
	0x9b, 0x99, 0x2e, 0xf8, 0x3a, 0x3e, 0xa9, 0x34, 0x69, 0x8b, 0xf4, 0x96, 0x7c, 0xdf, 0x77, 0x18,
	0x7e, 0x36, 0x09, 0xb6, 0x6b, 0x5d, 0x6d, 0xc8, 0x81, 0x6f, 0x0c, 0x99, 0xaa, 0x0b, 0x40, 0xc0,
	0xc7, 0x07, 0xf8, 0xfe, 0x77, 0xc0, 0xf2, 0x19, 0x19, 0x5a, 0x23, 0xbf, 0x65, 0xa3, 0x0e, 0xd0,
	0x6d, 0x94, 0xc8, 0x8a, 0xac, 0x3c, 0x55, 0xfb, 0x3f, 0x2f, 0xd9, 0x25, 0xb6, 0xa6, 0xb7, 0xda,
	0x81, 0x0e, 0x6b, 0xef, 0x9d, 0x5f, 0x88, 0x41, 0x91, 0x95, 0x23, 0x75, 0x11, 0xf9, 0x3b, 0xa8,
	0x44, 0xf9, 0x03, 0xbb, 0x4a, 0x25, 0xae, 0xda, 0x7d, 0x7a, 0x14, 0xd3, 0x71, 0x14, 0xb3, 0x55,
	0xbb, 0x6b, 0x9f, 0xd9, 0x04, 0x6d, 0x0d, 0xbe, 0x41, 0x3d, 0xb7, 0x4b, 0xe7, 0x1b, 0xfd, 0x6d,
	0x90, 0x6c, 0x10, 0xc7, 0x45, 0x56, 0x9e, 0xab, 0xeb, 0xad, 0x9c, 0x46, 0xf7, 0x11, 0x15, 0xbf,
	0x63, 0x67, 0x29, 0xd2, 0x4b, 0x40, 0x12, 0xc3, 0x78, 0x28, 0x4b, 0xe8, 0x0d, 0x90, 0xfe, 0x05,
	0x1d, 0x04, 0x12, 0x79, 0x91, 0x95, 0xc3, 0x5d, 0xf0, 0x09, 0x81, 0xf8, 0x13, 0xbb, 0xd9, 0x06,
	0x35, 0x78, 0x6f, 0x6b, 0xd2, 0xc1, 0x52, 0xf8, 0x11, 0x27, 0xb1, 0xe4, 0xc9, 0xbd, 0x26, 0xa5,
	0x36, 0x66, 0x5a, 0x7d, 0x3d, 0xf6, 0x8e, 0x2c, 0x62, 0xe5, 0x40, 0xa6, 0x97, 0x5c, 0x80, 0xec,
	0x49, 0xc6, 0x55, 0xe5, 0xc1, 0xa8, 0xf3, 0x3c, 0xe2, 0x97, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd3, 0x3d, 0xca, 0x01, 0x85, 0x01, 0x00, 0x00,
}
