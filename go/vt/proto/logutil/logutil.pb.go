// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logutil.proto

package logutil // import "vitess.io/vitess/go/vt/proto/logutil"

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

// Level is the level of the log messages.
type Level int32

const (
	// The usual logging levels.
	// Should be logged using logging facility.
	Level_INFO    Level = 0
	Level_WARNING Level = 1
	Level_ERROR   Level = 2
	// For messages that may contains non-logging events.
	// Should be logged to console directly.
	Level_CONSOLE Level = 3
)

var Level_name = map[int32]string{
	0: "INFO",
	1: "WARNING",
	2: "ERROR",
	3: "CONSOLE",
}
var Level_value = map[string]int32{
	"INFO":    0,
	"WARNING": 1,
	"ERROR":   2,
	"CONSOLE": 3,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}
func (Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_logutil_39c26af5691dd7cd, []int{0}
}

// Time represents a time stamp in nanoseconds. In go, use logutil library
// to convert times.
type Time struct {
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds          int32    `protobuf:"varint,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_logutil_39c26af5691dd7cd, []int{0}
}
func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}
func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}
func (dst *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(dst, src)
}
func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}
func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

func (m *Time) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Time) GetNanoseconds() int32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

// Event is a single logging event
type Event struct {
	Time                 *Time    `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Level                Level    `protobuf:"varint,2,opt,name=level,proto3,enum=logutil.Level" json:"level,omitempty"`
	File                 string   `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Line                 int64    `protobuf:"varint,4,opt,name=line,proto3" json:"line,omitempty"`
	Value                string   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_logutil_39c26af5691dd7cd, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetTime() *Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Event) GetLevel() Level {
	if m != nil {
		return m.Level
	}
	return Level_INFO
}

func (m *Event) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Event) GetLine() int64 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *Event) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Time)(nil), "logutil.Time")
	proto.RegisterType((*Event)(nil), "logutil.Event")
	proto.RegisterEnum("logutil.Level", Level_name, Level_value)
}

func init() { proto.RegisterFile("logutil.proto", fileDescriptor_logutil_39c26af5691dd7cd) }

var fileDescriptor_logutil_39c26af5691dd7cd = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xff, 0xdb, 0x64, 0xff, 0xb1, 0x13, 0x5a, 0xc2, 0xe0, 0x21, 0xc7, 0x58, 0x8a, 0x04,
	0x0f, 0x09, 0x54, 0xf0, 0x6e, 0x25, 0x4a, 0xa1, 0x24, 0xb0, 0x0a, 0x82, 0xb7, 0xaa, 0x63, 0x59,
	0xd8, 0x66, 0xc5, 0x6c, 0xf7, 0x63, 0xf8, 0x99, 0x25, 0x93, 0x46, 0xbc, 0xbd, 0xf7, 0x7b, 0xc3,
	0x9b, 0x61, 0x60, 0x66, 0xec, 0xfe, 0xe8, 0xb4, 0x29, 0x3e, 0xbf, 0xac, 0xb3, 0x18, 0x9d, 0xec,
	0x62, 0x0d, 0xe1, 0x93, 0x3e, 0x10, 0xa6, 0x10, 0x75, 0xf4, 0x66, 0xdb, 0xf7, 0x2e, 0x15, 0x99,
	0xc8, 0x03, 0x35, 0x5a, 0xcc, 0x20, 0x6e, 0x77, 0xad, 0x1d, 0xd3, 0x49, 0x26, 0x72, 0xa9, 0xfe,
	0xa2, 0xc5, 0xb7, 0x00, 0x59, 0x79, 0x6a, 0x1d, 0x5e, 0x40, 0xe8, 0xf4, 0x81, 0xb8, 0x22, 0x5e,
	0xcd, 0x8a, 0x71, 0x69, 0xbf, 0x42, 0x71, 0x84, 0x4b, 0x90, 0x86, 0x3c, 0x19, 0x2e, 0x9a, 0xaf,
	0xe6, 0xbf, 0x33, 0xdb, 0x9e, 0xaa, 0x21, 0x44, 0x84, 0xf0, 0x43, 0x1b, 0x4a, 0x83, 0x4c, 0xe4,
	0x53, 0xc5, 0xba, 0x67, 0x46, 0xb7, 0x94, 0x86, 0x7c, 0x1f, 0x6b, 0x3c, 0x07, 0xe9, 0x77, 0xe6,
	0x48, 0xa9, 0xe4, 0xc1, 0xc1, 0x5c, 0xdd, 0x80, 0xe4, 0x36, 0x3c, 0x83, 0x70, 0x53, 0xdf, 0x37,
	0xc9, 0x3f, 0x8c, 0x21, 0x7a, 0xbe, 0x55, 0xf5, 0xa6, 0x7e, 0x48, 0x04, 0x4e, 0x41, 0x56, 0x4a,
	0x35, 0x2a, 0x99, 0xf4, 0xfc, 0xae, 0xa9, 0x1f, 0x9b, 0x6d, 0x95, 0x04, 0xeb, 0xcb, 0x97, 0xa5,
	0xd7, 0x8e, 0xba, 0xae, 0xd0, 0xb6, 0x1c, 0x54, 0xb9, 0xb7, 0xa5, 0x77, 0x25, 0x7f, 0xad, 0x3c,
	0x9d, 0xfa, 0xfa, 0x9f, 0xed, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x02, 0xa1, 0x99,
	0x55, 0x01, 0x00, 0x00,
}
