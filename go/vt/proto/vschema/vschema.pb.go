// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vschema.proto

package vschema

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	query "vitess.io/vitess/go/vt/proto/query"
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

// RoutingRules specify the high level routing rules for the VSchema.
type RoutingRules struct {
	// rules should ideally be a map. However protos dont't allow
	// repeated fields as elements of a map. So, we use a list
	// instead.
	Rules                []*RoutingRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RoutingRules) Reset()         { *m = RoutingRules{} }
func (m *RoutingRules) String() string { return proto.CompactTextString(m) }
func (*RoutingRules) ProtoMessage()    {}
func (*RoutingRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f6849254fea3e77, []int{0}
}

func (m *RoutingRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutingRules.Unmarshal(m, b)
}
func (m *RoutingRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutingRules.Marshal(b, m, deterministic)
}
func (m *RoutingRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutingRules.Merge(m, src)
}
func (m *RoutingRules) XXX_Size() int {
	return xxx_messageInfo_RoutingRules.Size(m)
}
func (m *RoutingRules) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutingRules.DiscardUnknown(m)
}

var xxx_messageInfo_RoutingRules proto.InternalMessageInfo

func (m *RoutingRules) GetRules() []*RoutingRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// RoutingRule specifies a routing rule.
type RoutingRule struct {
	FromTable            string   `protobuf:"bytes,1,opt,name=from_table,json=fromTable,proto3" json:"from_table,omitempty"`
	ToTables             []string `protobuf:"bytes,2,rep,name=to_tables,json=toTables,proto3" json:"to_tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutingRule) Reset()         { *m = RoutingRule{} }
func (m *RoutingRule) String() string { return proto.CompactTextString(m) }
func (*RoutingRule) ProtoMessage()    {}
func (*RoutingRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f6849254fea3e77, []int{1}
}

func (m *RoutingRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutingRule.Unmarshal(m, b)
}
func (m *RoutingRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutingRule.Marshal(b, m, deterministic)
}
func (m *RoutingRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutingRule.Merge(m, src)
}
func (m *RoutingRule) XXX_Size() int {
	return xxx_messageInfo_RoutingRule.Size(m)
}
func (m *RoutingRule) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutingRule.DiscardUnknown(m)
}

var xxx_messageInfo_RoutingRule proto.InternalMessageInfo

func (m *RoutingRule) GetFromTable() string {
	if m != nil {
		return m.FromTable
	}
	return ""
}

func (m *RoutingRule) GetToTables() []string {
	if m != nil {
		return m.ToTables
	}
	return nil
}

// Keyspace is the vschema for a keyspace.
type Keyspace struct {
	// If sharded is false, vindexes and tables are ignored.
	Sharded  bool               `protobuf:"varint,1,opt,name=sharded,proto3" json:"sharded,omitempty"`
	Vindexes map[string]*Vindex `protobuf:"bytes,2,rep,name=vindexes,proto3" json:"vindexes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tables   map[string]*Table  `protobuf:"bytes,3,rep,name=tables,proto3" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If require_explicit_routing is true, vindexes and tables are not added to global routing
	RequireExplicitRouting bool     `protobuf:"varint,4,opt,name=require_explicit_routing,json=requireExplicitRouting,proto3" json:"require_explicit_routing,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Keyspace) Reset()         { *m = Keyspace{} }
func (m *Keyspace) String() string { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()    {}
func (*Keyspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f6849254fea3e77, []int{2}
}

func (m *Keyspace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyspace.Unmarshal(m, b)
}
func (m *Keyspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyspace.Marshal(b, m, deterministic)
}
func (m *Keyspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyspace.Merge(m, src)
}
func (m *Keyspace) XXX_Size() int {
	return xxx_messageInfo_Keyspace.Size(m)
}
func (m *Keyspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyspace.DiscardUnknown(m)
}

var xxx_messageInfo_Keyspace proto.InternalMessageInfo

func (m *Keyspace) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *Keyspace) GetVindexes() map[string]*Vindex {
	if m != nil {
		return m.Vindexes
	}
	return nil
}

func (m *Keyspace) GetTables() map[string]*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

func (m *Keyspace) GetRequireExplicitRouting() bool {
	if m != nil {
		return m.RequireExplicitRouting
	}
	return false
}

// Vindex is the vindex info for a Keyspace.
type Vindex struct {
	// The type must match one of the predefined
	// (or plugged in) vindex names.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// params is a map of attribute value pairs
	// that must be defined as required by the
	// vindex constructors. The values can only
	// be strings.
	Params map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A lookup vindex can have an owner table defined.
	// If so, rows in the lookup table are created or
	// deleted in sync with corresponding rows in the
	// owner table.
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vindex) Reset()         { *m = Vindex{} }
func (m *Vindex) String() string { return proto.CompactTextString(m) }
func (*Vindex) ProtoMessage()    {}
func (*Vindex) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f6849254fea3e77, []int{3}
}

func (m *Vindex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vindex.Unmarshal(m, b)
}
func (m *Vindex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vindex.Marshal(b, m, deterministic)
}
func (m *Vindex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vindex.Merge(m, src)
}
func (m *Vindex) XXX_Size() int {
	return xxx_messageInfo_Vindex.Size(m)
}
func (m *Vindex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vindex.DiscardUnknown(m)
}

var xxx_messageInfo_Vindex proto.InternalMessageInfo

func (m *Vindex) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Vindex) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Vindex) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// Table is the table info for a Keyspace.
type Table struct {
	// If the table is a sequence, type must be
	// "sequence". Otherwise, it should be empty.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// column_vindexes associates columns to vindexes.
	ColumnVindexes []*ColumnVindex `protobuf:"bytes,2,rep,name=column_vindexes,json=columnVindexes,proto3" json:"column_vindexes,omitempty"`
	// auto_increment is specified if a column needs
	// to be associated with a sequence.
	AutoIncrement *AutoIncrement `protobuf:"bytes,3,opt,name=auto_increment,json=autoIncrement,proto3" json:"auto_increment,omitempty"`
	// columns lists the columns for the table.
	Columns []*Column `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns,omitempty"`
	// pinned pins an unsharded table to a specific
	// shard, as dictated by the keyspace id.
	// The keyspace id is represented in hex form
	// like in keyranges.
	Pinned string `protobuf:"bytes,5,opt,name=pinned,proto3" json:"pinned,omitempty"`
	// column_list_authoritative is set to true if columns is
	// an authoritative list for the table. This allows
	// us to expand 'select *' expressions.
	ColumnListAuthoritative bool     `protobuf:"varint,6,opt,name=column_list_authoritative,json=columnListAuthoritative,proto3" json:"column_list_authoritative,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f6849254fea3e77, []int{4}
}

func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Table) GetColumnVindexes() []*ColumnVindex {
	if m != nil {
		return m.ColumnVindexes
	}
	return nil
}

func (m *Table) GetAutoIncrement() *AutoIncrement {
	if m != nil {
		return m.AutoIncrement
	}
	return nil
}

func (m *Table) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Table) GetPinned() string {
	if m != nil {
		return m.Pinned
	}
	return ""
}

func (m *Table) GetColumnListAuthoritative() bool {
	if m != nil {
		return m.ColumnListAuthoritative
	}
	return false
}

// ColumnVindex is used to associate a column to a vindex.
type ColumnVindex struct {
	// Legacy implementation, moving forward all vindexes should define a list of columns.
	Column string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	// The name must match a vindex defined in Keyspace.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// List of columns that define this Vindex
	Columns              []string `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColumnVindex) Reset()         { *m = ColumnVindex{} }
func (m *ColumnVindex) String() string { return proto.CompactTextString(m) }
func (*ColumnVindex) ProtoMessage()    {}
func (*ColumnVindex) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f6849254fea3e77, []int{5}
}

func (m *ColumnVindex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnVindex.Unmarshal(m, b)
}
func (m *ColumnVindex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnVindex.Marshal(b, m, deterministic)
}
func (m *ColumnVindex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnVindex.Merge(m, src)
}
func (m *ColumnVindex) XXX_Size() int {
	return xxx_messageInfo_ColumnVindex.Size(m)
}
func (m *ColumnVindex) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnVindex.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnVindex proto.InternalMessageInfo

func (m *ColumnVindex) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *ColumnVindex) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColumnVindex) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

// Autoincrement is used to designate a column as auto-inc.
type AutoIncrement struct {
	Column string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	// The sequence must match a table of type SEQUENCE.
	Sequence             string   `protobuf:"bytes,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoIncrement) Reset()         { *m = AutoIncrement{} }
func (m *AutoIncrement) String() string { return proto.CompactTextString(m) }
func (*AutoIncrement) ProtoMessage()    {}
func (*AutoIncrement) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f6849254fea3e77, []int{6}
}

func (m *AutoIncrement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoIncrement.Unmarshal(m, b)
}
func (m *AutoIncrement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoIncrement.Marshal(b, m, deterministic)
}
func (m *AutoIncrement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoIncrement.Merge(m, src)
}
func (m *AutoIncrement) XXX_Size() int {
	return xxx_messageInfo_AutoIncrement.Size(m)
}
func (m *AutoIncrement) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoIncrement.DiscardUnknown(m)
}

var xxx_messageInfo_AutoIncrement proto.InternalMessageInfo

func (m *AutoIncrement) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *AutoIncrement) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

// Column describes a column.
type Column struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 query.Type `protobuf:"varint,2,opt,name=type,proto3,enum=query.Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f6849254fea3e77, []int{7}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Column) GetType() query.Type {
	if m != nil {
		return m.Type
	}
	return query.Type_NULL_TYPE
}

// SrvVSchema is the roll-up of all the Keyspace schema for a cell.
type SrvVSchema struct {
	// keyspaces is a map of keyspace name -> Keyspace object.
	Keyspaces            map[string]*Keyspace `protobuf:"bytes,1,rep,name=keyspaces,proto3" json:"keyspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RoutingRules         *RoutingRules        `protobuf:"bytes,2,opt,name=routing_rules,json=routingRules,proto3" json:"routing_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SrvVSchema) Reset()         { *m = SrvVSchema{} }
func (m *SrvVSchema) String() string { return proto.CompactTextString(m) }
func (*SrvVSchema) ProtoMessage()    {}
func (*SrvVSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f6849254fea3e77, []int{8}
}

func (m *SrvVSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrvVSchema.Unmarshal(m, b)
}
func (m *SrvVSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrvVSchema.Marshal(b, m, deterministic)
}
func (m *SrvVSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrvVSchema.Merge(m, src)
}
func (m *SrvVSchema) XXX_Size() int {
	return xxx_messageInfo_SrvVSchema.Size(m)
}
func (m *SrvVSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_SrvVSchema.DiscardUnknown(m)
}

var xxx_messageInfo_SrvVSchema proto.InternalMessageInfo

func (m *SrvVSchema) GetKeyspaces() map[string]*Keyspace {
	if m != nil {
		return m.Keyspaces
	}
	return nil
}

func (m *SrvVSchema) GetRoutingRules() *RoutingRules {
	if m != nil {
		return m.RoutingRules
	}
	return nil
}

func init() {
	proto.RegisterType((*RoutingRules)(nil), "vschema.RoutingRules")
	proto.RegisterType((*RoutingRule)(nil), "vschema.RoutingRule")
	proto.RegisterType((*Keyspace)(nil), "vschema.Keyspace")
	proto.RegisterMapType((map[string]*Table)(nil), "vschema.Keyspace.TablesEntry")
	proto.RegisterMapType((map[string]*Vindex)(nil), "vschema.Keyspace.VindexesEntry")
	proto.RegisterType((*Vindex)(nil), "vschema.Vindex")
	proto.RegisterMapType((map[string]string)(nil), "vschema.Vindex.ParamsEntry")
	proto.RegisterType((*Table)(nil), "vschema.Table")
	proto.RegisterType((*ColumnVindex)(nil), "vschema.ColumnVindex")
	proto.RegisterType((*AutoIncrement)(nil), "vschema.AutoIncrement")
	proto.RegisterType((*Column)(nil), "vschema.Column")
	proto.RegisterType((*SrvVSchema)(nil), "vschema.SrvVSchema")
	proto.RegisterMapType((map[string]*Keyspace)(nil), "vschema.SrvVSchema.KeyspacesEntry")
}

func init() { proto.RegisterFile("vschema.proto", fileDescriptor_3f6849254fea3e77) }

var fileDescriptor_3f6849254fea3e77 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0x96, 0x13, 0x62, 0x92, 0x31, 0x09, 0xbf, 0xdf, 0x0a, 0xa8, 0x1b, 0x84, 0x88, 0x2c, 0xda,
	0xa6, 0x3d, 0x24, 0x52, 0x50, 0x25, 0x9a, 0x8a, 0xaa, 0x14, 0x71, 0x40, 0x45, 0x6a, 0x65, 0x10,
	0x87, 0x5e, 0x2c, 0xe3, 0x6c, 0x61, 0x45, 0xe2, 0x35, 0xbb, 0x6b, 0x97, 0x3c, 0x4a, 0xaf, 0x7d,
	0xad, 0x3e, 0x42, 0x5f, 0xa2, 0xf2, 0xfe, 0x31, 0x1b, 0x48, 0x6f, 0x3b, 0x3b, 0xf3, 0x7d, 0xf3,
	0xed, 0xec, 0xcc, 0x40, 0xbb, 0xe0, 0xc9, 0x0d, 0x9e, 0xc5, 0x83, 0x8c, 0x51, 0x41, 0xd1, 0xaa,
	0x36, 0xbb, 0xde, 0x5d, 0x8e, 0xd9, 0x5c, 0xdd, 0x06, 0x63, 0x58, 0x0b, 0x69, 0x2e, 0x48, 0x7a,
	0x1d, 0xe6, 0x53, 0xcc, 0xd1, 0x1b, 0x68, 0xb0, 0xf2, 0xe0, 0x3b, 0xbd, 0x7a, 0xdf, 0x1b, 0x6d,
	0x0c, 0x0c, 0x89, 0x15, 0x15, 0xaa, 0x90, 0xe0, 0x14, 0x3c, 0xeb, 0x16, 0xed, 0x00, 0x7c, 0x67,
	0x74, 0x16, 0x89, 0xf8, 0x6a, 0x8a, 0x7d, 0xa7, 0xe7, 0xf4, 0x5b, 0x61, 0xab, 0xbc, 0xb9, 0x28,
	0x2f, 0xd0, 0x36, 0xb4, 0x04, 0x55, 0x4e, 0xee, 0xd7, 0x7a, 0xf5, 0x7e, 0x2b, 0x6c, 0x0a, 0x2a,
	0x7d, 0x3c, 0xf8, 0x53, 0x83, 0xe6, 0x67, 0x3c, 0xe7, 0x59, 0x9c, 0x60, 0xe4, 0xc3, 0x2a, 0xbf,
	0x89, 0xd9, 0x04, 0x4f, 0x24, 0x4b, 0x33, 0x34, 0x26, 0x7a, 0x0f, 0xcd, 0x82, 0xa4, 0x13, 0x7c,
	0xaf, 0x29, 0xbc, 0xd1, 0x6e, 0x25, 0xd0, 0xc0, 0x07, 0x97, 0x3a, 0xe2, 0x24, 0x15, 0x6c, 0x1e,
	0x56, 0x00, 0xf4, 0x16, 0x5c, 0x9d, 0xbd, 0x2e, 0xa1, 0x3b, 0x4f, 0xa1, 0x4a, 0x8d, 0x02, 0xea,
	0x60, 0x74, 0x00, 0x3e, 0xc3, 0x77, 0x39, 0x61, 0x38, 0xc2, 0xf7, 0xd9, 0x94, 0x24, 0x44, 0x44,
	0x4c, 0x3d, 0xdb, 0x5f, 0x91, 0xf2, 0xb6, 0xb4, 0xff, 0x44, 0xbb, 0x75, 0x51, 0xba, 0x67, 0xd0,
	0x5e, 0xd0, 0x82, 0xfe, 0x83, 0xfa, 0x2d, 0x9e, 0xeb, 0xd2, 0x94, 0x47, 0xf4, 0x02, 0x1a, 0x45,
	0x3c, 0xcd, 0xb1, 0x5f, 0xeb, 0x39, 0x7d, 0x6f, 0xb4, 0x5e, 0x49, 0x52, 0xc0, 0x50, 0x79, 0xc7,
	0xb5, 0x03, 0xa7, 0x7b, 0x0a, 0x9e, 0x25, 0x6f, 0x09, 0xd7, 0xde, 0x22, 0x57, 0xa7, 0xe2, 0x92,
	0x30, 0x8b, 0x2a, 0xf8, 0xe5, 0x80, 0xab, 0x12, 0x20, 0x04, 0x2b, 0x62, 0x9e, 0x99, 0xef, 0x92,
	0x67, 0xb4, 0x0f, 0x6e, 0x16, 0xb3, 0x78, 0x66, 0x6a, 0xbc, 0xfd, 0x48, 0xd5, 0xe0, 0xab, 0xf4,
	0xea, 0x32, 0xa9, 0x50, 0xb4, 0x01, 0x0d, 0xfa, 0x23, 0xc5, 0xcc, 0xaf, 0x4b, 0x26, 0x65, 0x74,
	0xdf, 0x81, 0x67, 0x05, 0x2f, 0x11, 0xbd, 0x61, 0x8b, 0x6e, 0xd9, 0x22, 0x7f, 0xd6, 0xa0, 0xa1,
	0x3a, 0x67, 0x99, 0xc6, 0x0f, 0xb0, 0x9e, 0xd0, 0x69, 0x3e, 0x4b, 0xa3, 0x47, 0x0d, 0xb1, 0x59,
	0x89, 0x3d, 0x96, 0x7e, 0x5d, 0xc8, 0x4e, 0x62, 0x59, 0x98, 0xa3, 0x43, 0xe8, 0xc4, 0xb9, 0xa0,
	0x11, 0x49, 0x13, 0x86, 0x67, 0x38, 0x15, 0x52, 0xb7, 0x37, 0xda, 0xaa, 0xe0, 0x47, 0xb9, 0xa0,
	0xa7, 0xc6, 0x1b, 0xb6, 0x63, 0xdb, 0x44, 0xaf, 0x61, 0x55, 0x11, 0x72, 0x7f, 0x45, 0xa6, 0x5d,
	0x7f, 0x94, 0x36, 0x34, 0x7e, 0xb4, 0x05, 0x6e, 0x46, 0xd2, 0x14, 0x4f, 0xfc, 0x86, 0xd4, 0xaf,
	0x2d, 0x34, 0x86, 0xe7, 0xfa, 0x05, 0x53, 0xc2, 0x45, 0x14, 0xe7, 0xe2, 0x86, 0x32, 0x22, 0x62,
	0x41, 0x0a, 0xec, 0xbb, 0xb2, 0xb1, 0x9e, 0xa9, 0x80, 0x33, 0xc2, 0xc5, 0x91, 0xed, 0x0e, 0x2e,
	0x60, 0xcd, 0x7e, 0x5d, 0x99, 0x43, 0x85, 0xea, 0x1a, 0x69, 0xab, 0xac, 0x5c, 0x1a, 0xcf, 0x4c,
	0x71, 0xe5, 0xb9, 0x9c, 0x2e, 0x23, 0xbd, 0x2e, 0xa7, 0xd0, 0x98, 0xc1, 0x31, 0xb4, 0x17, 0x1e,
	0xfd, 0x4f, 0xda, 0x2e, 0x34, 0x39, 0xbe, 0xcb, 0x71, 0x9a, 0x18, 0xea, 0xca, 0x0e, 0x0e, 0xc1,
	0x3d, 0x5e, 0x4c, 0xee, 0x58, 0xc9, 0x77, 0xf5, 0x57, 0x96, 0xa8, 0xce, 0xc8, 0x1b, 0xa8, 0x55,
	0x74, 0x31, 0xcf, 0xb0, 0xfa, 0xd7, 0xe0, 0xb7, 0x03, 0x70, 0xce, 0x8a, 0xcb, 0x73, 0x59, 0x4c,
	0xf4, 0x11, 0x5a, 0xb7, 0x7a, 0x38, 0xcd, 0x4a, 0x0a, 0xaa, 0x4a, 0x3f, 0xc4, 0x55, 0x13, 0xac,
	0x9b, 0xf2, 0x01, 0x84, 0xc6, 0xd0, 0xd6, 0xd3, 0x1a, 0xa9, 0xc5, 0xa6, 0xa6, 0x63, 0x73, 0xd9,
	0x62, 0xe3, 0xe1, 0x1a, 0xb3, 0xac, 0xee, 0x17, 0xe8, 0x2c, 0x12, 0x2f, 0x69, 0xe0, 0x57, 0x8b,
	0x53, 0xf7, 0xff, 0x93, 0xa5, 0x62, 0xf5, 0xf4, 0xa7, 0x97, 0xdf, 0xf6, 0x0a, 0x22, 0x30, 0xe7,
	0x03, 0x42, 0x87, 0xea, 0x34, 0xbc, 0xa6, 0xc3, 0x42, 0x0c, 0xe5, 0x36, 0x1e, 0x6a, 0xec, 0x95,
	0x2b, 0xcd, 0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xa7, 0x99, 0x19, 0xc3, 0x05, 0x00,
	0x00,
}
