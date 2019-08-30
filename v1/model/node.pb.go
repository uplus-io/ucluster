// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package model

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

type ClusterHealth int32

const (
	ClusterHealth_CH_Unknown       ClusterHealth = 0
	ClusterHealth_CH_Normal        ClusterHealth = 1
	ClusterHealth_CH_Unavailable   ClusterHealth = 2
	ClusterHealth_CH_NotInitialize ClusterHealth = 3
	ClusterHealth_CH_NotSafe       ClusterHealth = 4
)

var ClusterHealth_name = map[int32]string{
	0: "CH_Unknown",
	1: "CH_Normal",
	2: "CH_Unavailable",
	3: "CH_NotInitialize",
	4: "CH_NotSafe",
}

var ClusterHealth_value = map[string]int32{
	"CH_Unknown":       0,
	"CH_Normal":        1,
	"CH_Unavailable":   2,
	"CH_NotInitialize": 3,
	"CH_NotSafe":       4,
}

func (x ClusterHealth) String() string {
	return proto.EnumName(ClusterHealth_name, int32(x))
}

func (ClusterHealth) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

type NodeHealth int32

const (
	NodeHealth_Suspect NodeHealth = 0
	NodeHealth_Alive   NodeHealth = 1
	NodeHealth_Dead    NodeHealth = 2
)

var NodeHealth_name = map[int32]string{
	0: "Suspect",
	1: "Alive",
	2: "Dead",
}

var NodeHealth_value = map[string]int32{
	"Suspect": 0,
	"Alive":   1,
	"Dead":    2,
}

func (x NodeHealth) String() string {
	return proto.EnumName(NodeHealth_name, int32(x))
}

func (NodeHealth) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

type NodeStatus int32

const (
	NodeStatus_Unknown     NodeStatus = 0
	NodeStatus_Normal      NodeStatus = 1
	NodeStatus_Unavailable NodeStatus = 2
	NodeStatus_New         NodeStatus = 3
	NodeStatus_Recovering  NodeStatus = 4
)

var NodeStatus_name = map[int32]string{
	0: "Unknown",
	1: "Normal",
	2: "Unavailable",
	3: "New",
	4: "Recovering",
}

var NodeStatus_value = map[string]int32{
	"Unknown":     0,
	"Normal":      1,
	"Unavailable": 2,
	"New":         3,
	"Recovering":  4,
}

func (x NodeStatus) String() string {
	return proto.EnumName(NodeStatus_name, int32(x))
}

func (NodeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{2}
}

type Repository struct {
	DataCenter           int32    `protobuf:"varint,1,opt,name=DataCenter,proto3" json:"DataCenter,omitempty"`
	Area                 int32    `protobuf:"varint,2,opt,name=Area,proto3" json:"Area,omitempty"`
	Rack                 int32    `protobuf:"varint,3,opt,name=Rack,proto3" json:"Rack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetDataCenter() int32 {
	if m != nil {
		return m.DataCenter
	}
	return 0
}

func (m *Repository) GetArea() int32 {
	if m != nil {
		return m.Area
	}
	return 0
}

func (m *Repository) GetRack() int32 {
	if m != nil {
		return m.Rack
	}
	return 0
}

type Partition struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Index                int32    `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Partition) Reset()         { *m = Partition{} }
func (m *Partition) String() string { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()    {}
func (*Partition) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

func (m *Partition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partition.Unmarshal(m, b)
}
func (m *Partition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partition.Marshal(b, m, deterministic)
}
func (m *Partition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partition.Merge(m, src)
}
func (m *Partition) XXX_Size() int {
	return xxx_messageInfo_Partition.Size(m)
}
func (m *Partition) XXX_DiscardUnknown() {
	xxx_messageInfo_Partition.DiscardUnknown(m)
}

var xxx_messageInfo_Partition proto.InternalMessageInfo

func (m *Partition) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Partition) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Partition) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type NodeStorageInfo struct {
	Repository           *Repository  `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	PartitionSize        int32        `protobuf:"varint,2,opt,name=partitionSize,proto3" json:"partitionSize,omitempty"`
	ReplicaSize          int32        `protobuf:"varint,3,opt,name=replicaSize,proto3" json:"replicaSize,omitempty"`
	Health               NodeHealth   `protobuf:"varint,4,opt,name=health,proto3,enum=model.NodeHealth" json:"health,omitempty"`
	Status               NodeStatus   `protobuf:"varint,5,opt,name=status,proto3,enum=model.NodeStatus" json:"status,omitempty"`
	Partitions           []*Partition `protobuf:"bytes,6,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NodeStorageInfo) Reset()         { *m = NodeStorageInfo{} }
func (m *NodeStorageInfo) String() string { return proto.CompactTextString(m) }
func (*NodeStorageInfo) ProtoMessage()    {}
func (*NodeStorageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{2}
}

func (m *NodeStorageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStorageInfo.Unmarshal(m, b)
}
func (m *NodeStorageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStorageInfo.Marshal(b, m, deterministic)
}
func (m *NodeStorageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStorageInfo.Merge(m, src)
}
func (m *NodeStorageInfo) XXX_Size() int {
	return xxx_messageInfo_NodeStorageInfo.Size(m)
}
func (m *NodeStorageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStorageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStorageInfo proto.InternalMessageInfo

func (m *NodeStorageInfo) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *NodeStorageInfo) GetPartitionSize() int32 {
	if m != nil {
		return m.PartitionSize
	}
	return 0
}

func (m *NodeStorageInfo) GetReplicaSize() int32 {
	if m != nil {
		return m.ReplicaSize
	}
	return 0
}

func (m *NodeStorageInfo) GetHealth() NodeHealth {
	if m != nil {
		return m.Health
	}
	return NodeHealth_Suspect
}

func (m *NodeStorageInfo) GetStatus() NodeStatus {
	if m != nil {
		return m.Status
	}
	return NodeStatus_Unknown
}

func (m *NodeStorageInfo) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

type NodeStorageStat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStorageStat) Reset()         { *m = NodeStorageStat{} }
func (m *NodeStorageStat) String() string { return proto.CompactTextString(m) }
func (*NodeStorageStat) ProtoMessage()    {}
func (*NodeStorageStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{3}
}

func (m *NodeStorageStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStorageStat.Unmarshal(m, b)
}
func (m *NodeStorageStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStorageStat.Marshal(b, m, deterministic)
}
func (m *NodeStorageStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStorageStat.Merge(m, src)
}
func (m *NodeStorageStat) XXX_Size() int {
	return xxx_messageInfo_NodeStorageStat.Size(m)
}
func (m *NodeStorageStat) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStorageStat.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStorageStat proto.InternalMessageInfo

type NodeHealthStat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeHealthStat) Reset()         { *m = NodeHealthStat{} }
func (m *NodeHealthStat) String() string { return proto.CompactTextString(m) }
func (*NodeHealthStat) ProtoMessage()    {}
func (*NodeHealthStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{4}
}

func (m *NodeHealthStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeHealthStat.Unmarshal(m, b)
}
func (m *NodeHealthStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeHealthStat.Marshal(b, m, deterministic)
}
func (m *NodeHealthStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeHealthStat.Merge(m, src)
}
func (m *NodeHealthStat) XXX_Size() int {
	return xxx_messageInfo_NodeHealthStat.Size(m)
}
func (m *NodeHealthStat) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeHealthStat.DiscardUnknown(m)
}

var xxx_messageInfo_NodeHealthStat proto.InternalMessageInfo

type Node struct {
	Id                   int32       `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Ip                   string      `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32       `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Repository           *Repository `protobuf:"bytes,5,opt,name=repository,proto3" json:"repository,omitempty"`
	Health               NodeHealth  `protobuf:"varint,7,opt,name=health,proto3,enum=model.NodeHealth" json:"health,omitempty"`
	Status               NodeStatus  `protobuf:"varint,8,opt,name=status,proto3,enum=model.NodeStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{5}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Node) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Node) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *Node) GetHealth() NodeHealth {
	if m != nil {
		return m.Health
	}
	return NodeHealth_Suspect
}

func (m *Node) GetStatus() NodeStatus {
	if m != nil {
		return m.Status
	}
	return NodeStatus_Unknown
}

func init() {
	proto.RegisterEnum("model.ClusterHealth", ClusterHealth_name, ClusterHealth_value)
	proto.RegisterEnum("model.NodeHealth", NodeHealth_name, NodeHealth_value)
	proto.RegisterEnum("model.NodeStatus", NodeStatus_name, NodeStatus_value)
	proto.RegisterType((*Repository)(nil), "model.Repository")
	proto.RegisterType((*Partition)(nil), "model.Partition")
	proto.RegisterType((*NodeStorageInfo)(nil), "model.NodeStorageInfo")
	proto.RegisterType((*NodeStorageStat)(nil), "model.NodeStorageStat")
	proto.RegisterType((*NodeHealthStat)(nil), "model.NodeHealthStat")
	proto.RegisterType((*Node)(nil), "model.Node")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x4e, 0xdb, 0x40,
	0x10, 0x86, 0xb1, 0x63, 0x27, 0x64, 0x2c, 0xcc, 0x32, 0xe2, 0xe0, 0x53, 0x15, 0x59, 0x3d, 0xa4,
	0x39, 0x44, 0x2d, 0x7d, 0x02, 0x14, 0x0e, 0x41, 0x95, 0xd2, 0x6a, 0xdd, 0x9e, 0xd1, 0x12, 0x0f,
	0x64, 0x85, 0xd9, 0xb5, 0xd6, 0x9b, 0xd0, 0xf2, 0x1c, 0x7d, 0xa3, 0xbe, 0x58, 0xe5, 0xb5, 0xb1,
	0x93, 0x72, 0x28, 0xb7, 0xdd, 0x7f, 0x7e, 0xfd, 0xfb, 0xcd, 0x6f, 0x19, 0x40, 0xe9, 0x9c, 0xe6,
	0xa5, 0xd1, 0x56, 0x63, 0xf8, 0xa8, 0x73, 0x2a, 0xd2, 0xef, 0x00, 0x9c, 0x4a, 0x5d, 0x49, 0xab,
	0xcd, 0x2f, 0x7c, 0x07, 0x70, 0x25, 0xac, 0x58, 0x90, 0xb2, 0x64, 0x12, 0x6f, 0xe2, 0x4d, 0x43,
	0xbe, 0xa7, 0x20, 0x42, 0x70, 0x69, 0x48, 0x24, 0xbe, 0x9b, 0xb8, 0x73, 0xad, 0x71, 0xb1, 0x7e,
	0x48, 0x06, 0x8d, 0x56, 0x9f, 0xd3, 0x2f, 0x30, 0xfe, 0x26, 0x8c, 0x95, 0x56, 0x6a, 0x85, 0x09,
	0x8c, 0x76, 0x64, 0x2a, 0xa9, 0x55, 0x9b, 0xf8, 0x72, 0xc5, 0x18, 0x7c, 0x99, 0xb7, 0x61, 0xbe,
	0xcc, 0xf1, 0x1c, 0x42, 0xa9, 0x72, 0xfa, 0xd9, 0x66, 0x35, 0x97, 0xf4, 0xb7, 0x0f, 0xa7, 0x2b,
	0x9d, 0x53, 0x66, 0xb5, 0x11, 0xf7, 0x74, 0xad, 0xee, 0x34, 0x7e, 0x02, 0x30, 0x1d, 0xb6, 0x8b,
	0x8d, 0x2e, 0xce, 0xe6, 0x6e, 0xa5, 0x79, 0xbf, 0x0f, 0xdf, 0x33, 0xe1, 0x7b, 0x38, 0x29, 0x5f,
	0x98, 0x32, 0xf9, 0x4c, 0xed, 0xbb, 0x87, 0x22, 0x4e, 0x20, 0x32, 0x54, 0x16, 0x72, 0x2d, 0x9c,
	0xa7, 0x01, 0xd9, 0x97, 0xf0, 0x03, 0x0c, 0x37, 0x24, 0x0a, 0xbb, 0x49, 0x82, 0x89, 0x37, 0x8d,
	0xbb, 0x67, 0x6b, 0xc4, 0xa5, 0x1b, 0xf0, 0xd6, 0x50, 0x5b, 0x2b, 0x2b, 0xec, 0xb6, 0x4a, 0xc2,
	0x57, 0xd6, 0xcc, 0x0d, 0x78, 0x6b, 0xc0, 0x8f, 0x00, 0x1d, 0x48, 0x95, 0x0c, 0x27, 0x83, 0x69,
	0x74, 0xc1, 0x5a, 0x7b, 0x57, 0x25, 0xdf, 0xf3, 0xa4, 0x67, 0x07, 0xad, 0xd4, 0x71, 0x29, 0x83,
	0xb8, 0xa7, 0x70, 0xca, 0x1f, 0x0f, 0x82, 0x5a, 0x7a, 0x55, 0x75, 0x7d, 0x2f, 0xdd, 0x7a, 0x63,
	0xee, 0xcb, 0xb2, 0xfe, 0x8a, 0xa5, 0x36, 0xd6, 0xed, 0x14, 0x72, 0x77, 0xfe, 0xa7, 0xe4, 0xf0,
	0x2d, 0x25, 0xf7, 0xe5, 0x8c, 0xde, 0x5e, 0xce, 0xf1, 0x7f, 0xca, 0x99, 0x6d, 0xe0, 0x64, 0x51,
	0x6c, 0x2b, 0x4b, 0xa6, 0xc9, 0xc0, 0x18, 0x60, 0xb1, 0xbc, 0xf9, 0xa1, 0x1e, 0x94, 0x7e, 0x52,
	0xec, 0x08, 0x4f, 0x60, 0xbc, 0x58, 0xde, 0xac, 0xb4, 0x79, 0x14, 0x05, 0xf3, 0x10, 0x21, 0x76,
	0x63, 0xb1, 0x13, 0xb2, 0x10, 0xb7, 0x05, 0x31, 0x1f, 0xcf, 0x81, 0x39, 0x8b, 0xbd, 0x56, 0xd2,
	0x4a, 0x51, 0xc8, 0x67, 0x62, 0x83, 0x36, 0x68, 0xa5, 0x6d, 0x26, 0xee, 0x88, 0x05, 0xb3, 0x39,
	0x40, 0x8f, 0x8a, 0x11, 0x8c, 0xb2, 0x6d, 0x55, 0xd2, 0xda, 0xb2, 0x23, 0x1c, 0x43, 0x78, 0x59,
	0xc8, 0x1d, 0x31, 0x0f, 0x8f, 0x21, 0xb8, 0x22, 0x91, 0x33, 0x7f, 0xf6, 0xb5, 0xf1, 0x37, 0xbc,
	0xb5, 0xbf, 0x67, 0x02, 0x18, 0x76, 0x40, 0xa7, 0x10, 0x1d, 0xd2, 0x8c, 0x60, 0xb0, 0xa2, 0xa7,
	0x06, 0x80, 0xd3, 0x5a, 0xef, 0xc8, 0x48, 0x75, 0xcf, 0x82, 0xdb, 0xa1, 0xfb, 0x3b, 0x3f, 0xff,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x61, 0x59, 0xd5, 0xab, 0x03, 0x00, 0x00,
}