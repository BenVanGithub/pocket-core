// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/nodes/nodes.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_pokt_network_pocket_core_types "github.com/pokt-network/pocket-core/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorProto struct {
	Address                 github_com_pokt_network_pocket_core_types.Address     `protobuf:"bytes,1,opt,name=Address,proto3,casttype=github.com/pokt-network/pocket-core/types.Address" json:"address" yaml:"address"`
	PublicKey               string                                                `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"public_key" yaml:"public_key"`
	Jailed                  bool                                                  `protobuf:"varint,3,opt,name=jailed,proto3" json:"jailed,omitempty"`
	Status                  github_com_pokt_network_pocket_core_types.StakeStatus `protobuf:"varint,4,opt,name=status,proto3,casttype=github.com/pokt-network/pocket-core/types.StakeStatus" json:"status,omitempty"`
	Chains                  []string                                              `protobuf:"bytes,5,rep,name=Chains,proto3" json:"chains"`
	ServiceURL              string                                                `protobuf:"bytes,6,opt,name=ServiceURL,proto3" json:"service_url"`
	StakedTokens            github_com_pokt_network_pocket_core_types.Int         `protobuf:"bytes,7,opt,name=StakedTokens,proto3,customtype=github.com/pokt-network/pocket-core/types.Int" json:"tokens"`
	UnstakingCompletionTime time.Time                                             `protobuf:"bytes,8,opt,name=UnstakingCompletionTime,proto3,stdtime" json:"unstaking_time" yaml:"unstaking_time"`
}

func (m *ValidatorProto) Reset()         { *m = ValidatorProto{} }
func (m *ValidatorProto) String() string { return proto.CompactTextString(m) }
func (*ValidatorProto) ProtoMessage()    {}
func (*ValidatorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb49073b61e33a, []int{0}
}
func (m *ValidatorProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorProto.Merge(m, src)
}
func (m *ValidatorProto) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorProto.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorProto proto.InternalMessageInfo

// ValidatorSigningInfo defines the signing info for a validator
type ValidatorSigningInfo struct {
	Address github_com_pokt_network_pocket_core_types.Address `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/pokt-network/pocket-core/types.Address" json:"address"`
	// height at which validator was first a candidate OR was unjailed
	StartHeight int64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height" yaml:"start_height"`
	// index offset into signed block bit array
	Index int64 `protobuf:"varint,3,opt,name=Index,proto3" json:"index_offset" yaml:"index_offset"`
	// timestamp validator cannot be unjailed until
	JailedUntil time.Time `protobuf:"bytes,4,opt,name=jailed_until,json=jailedUntil,proto3,stdtime" json:"jailed_until" yaml:"jailed_until"`
	// missed blocks counter (to avoid scanning the array every time)
	MissedBlocksCounter int64 `protobuf:"varint,5,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter" yaml:"missed_blocks_counter"`
	JailedBlocksCounter int64 `protobuf:"varint,6,opt,name=jailed_blocks_counter,json=jailedBlocksCounter,proto3" json:"jailed_blocks_counter" yaml:"jailed_blocks_counter"`
}

func (m *ValidatorSigningInfo) Reset()      { *m = ValidatorSigningInfo{} }
func (*ValidatorSigningInfo) ProtoMessage() {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb49073b61e33a, []int{1}
}
func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

func (m *ValidatorSigningInfo) GetAddress() github_com_pokt_network_pocket_core_types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ValidatorSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *ValidatorSigningInfo) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ValidatorSigningInfo) GetJailedUntil() time.Time {
	if m != nil {
		return m.JailedUntil
	}
	return time.Time{}
}

func (m *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

func (m *ValidatorSigningInfo) GetJailedBlocksCounter() int64 {
	if m != nil {
		return m.JailedBlocksCounter
	}
	return 0
}

type Power struct {
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Power) Reset()         { *m = Power{} }
func (m *Power) String() string { return proto.CompactTextString(m) }
func (*Power) ProtoMessage()    {}
func (*Power) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb49073b61e33a, []int{2}
}
func (m *Power) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Power) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Power.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Power) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Power.Merge(m, src)
}
func (m *Power) XXX_Size() int {
	return m.Size()
}
func (m *Power) XXX_DiscardUnknown() {
	xxx_messageInfo_Power.DiscardUnknown(m)
}

var xxx_messageInfo_Power proto.InternalMessageInfo

func (m *Power) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ValidatorMissed struct {
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ValidatorMissed) Reset()         { *m = ValidatorMissed{} }
func (m *ValidatorMissed) String() string { return proto.CompactTextString(m) }
func (*ValidatorMissed) ProtoMessage()    {}
func (*ValidatorMissed) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb49073b61e33a, []int{3}
}
func (m *ValidatorMissed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorMissed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorMissed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorMissed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorMissed.Merge(m, src)
}
func (m *ValidatorMissed) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorMissed) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorMissed.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorMissed proto.InternalMessageInfo

func (m *ValidatorMissed) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*ValidatorProto)(nil), "x.nodes.ValidatorProto")
	proto.RegisterType((*ValidatorSigningInfo)(nil), "x.nodes.ValidatorSigningInfo")
	proto.RegisterType((*Power)(nil), "x.nodes.Power")
	proto.RegisterType((*ValidatorMissed)(nil), "x.nodes.ValidatorMissed")
}

func init() { proto.RegisterFile("x/nodes/nodes.proto", fileDescriptor_63cb49073b61e33a) }

var fileDescriptor_63cb49073b61e33a = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xbf, 0x4f, 0x23, 0x47,
	0x14, 0xc7, 0xbd, 0x31, 0xfe, 0xc1, 0xd8, 0x02, 0x65, 0x0d, 0xc9, 0x0a, 0x25, 0x1e, 0x6b, 0x53,
	0xe0, 0x06, 0xaf, 0x12, 0x94, 0x02, 0xa4, 0x28, 0x61, 0x69, 0x42, 0x48, 0x24, 0xb2, 0x86, 0x14,
	0x34, 0xab, 0xf5, 0xee, 0x78, 0x3d, 0xd9, 0x1f, 0x63, 0xed, 0xcc, 0x02, 0xee, 0xae, 0xbc, 0xeb,
	0x28, 0x29, 0xfd, 0xe7, 0x50, 0xd2, 0x9c, 0x74, 0xba, 0x62, 0xee, 0x04, 0xcd, 0xc9, 0xa5, 0x4b,
	0xaa, 0x93, 0x67, 0xd6, 0xd8, 0x3e, 0xf9, 0x4e, 0xe8, 0x1a, 0xc4, 0xfb, 0xcc, 0x7b, 0xdf, 0xef,
	0x5b, 0xbd, 0xaf, 0x0c, 0x6a, 0x57, 0x46, 0x4c, 0x3c, 0x44, 0xe5, 0xdf, 0x56, 0x3f, 0x21, 0x8c,
	0xa8, 0xa5, 0xab, 0x96, 0x28, 0xb7, 0x36, 0x7c, 0xe2, 0x13, 0xc1, 0x8c, 0xc9, 0x7f, 0xf2, 0x79,
	0x0b, 0xfa, 0x84, 0xf8, 0x21, 0x32, 0x44, 0xd5, 0x49, 0xbb, 0x06, 0xc3, 0x11, 0xa2, 0xcc, 0x89,
	0xfa, 0xb2, 0x41, 0x7f, 0x51, 0x00, 0x6b, 0xff, 0x39, 0x21, 0xf6, 0x1c, 0x46, 0x92, 0x13, 0x21,
	0x19, 0x82, 0xd2, 0x81, 0xe7, 0x25, 0x88, 0x52, 0x4d, 0x69, 0x28, 0xcd, 0xaa, 0x69, 0x8d, 0x38,
	0x2c, 0x39, 0x12, 0x8d, 0x39, 0x5c, 0x1b, 0x38, 0x51, 0xb8, 0xaf, 0x67, 0x40, 0x7f, 0xe4, 0xf0,
	0x67, 0x1f, 0xb3, 0x5e, 0xda, 0x69, 0xb9, 0x24, 0x32, 0xfa, 0x24, 0x60, 0x3b, 0x31, 0x62, 0x97,
	0x24, 0x09, 0x8c, 0x3e, 0x71, 0x03, 0xc4, 0x76, 0x5c, 0x92, 0x20, 0x83, 0x0d, 0xfa, 0x88, 0xb6,
	0x32, 0x65, 0x6b, 0x6a, 0xa1, 0x1e, 0x80, 0xd5, 0x93, 0xb4, 0x13, 0x62, 0xf7, 0x18, 0x0d, 0xb4,
	0x6f, 0x1a, 0x4a, 0x73, 0xd5, 0xfc, 0x69, 0xc4, 0x21, 0xe8, 0x0b, 0x68, 0x07, 0x68, 0x30, 0xe6,
	0xf0, 0x5b, 0x69, 0x39, 0x63, 0xba, 0x35, 0x9b, 0x52, 0xbf, 0x03, 0xc5, 0xff, 0x1d, 0x1c, 0x22,
	0x4f, 0xcb, 0x37, 0x94, 0x66, 0xd9, 0xca, 0x2a, 0xf5, 0x5f, 0x50, 0xa4, 0xcc, 0x61, 0x29, 0xd5,
	0x56, 0x1a, 0x4a, 0xb3, 0x60, 0xee, 0x3d, 0x72, 0xf8, 0xeb, 0xf3, 0x57, 0x6d, 0x33, 0x27, 0x40,
	0x6d, 0x21, 0x60, 0x65, 0x42, 0xaa, 0x0e, 0x8a, 0x87, 0x3d, 0x07, 0xc7, 0x54, 0x2b, 0x34, 0xf2,
	0xcd, 0x55, 0x13, 0x8c, 0x38, 0x2c, 0xba, 0x82, 0x58, 0xd9, 0x8b, 0x6a, 0x00, 0xd0, 0x46, 0xc9,
	0x05, 0x76, 0xd1, 0x99, 0xf5, 0xb7, 0x56, 0x14, 0x9f, 0xb4, 0x3e, 0xe2, 0xb0, 0x42, 0x25, 0xb5,
	0xd3, 0x24, 0xb4, 0xe6, 0x5a, 0x54, 0x17, 0x54, 0x85, 0x97, 0x77, 0x4a, 0x02, 0x14, 0x53, 0xad,
	0x24, 0x46, 0x7e, 0xbf, 0xe5, 0x30, 0xf7, 0x96, 0xc3, 0x9d, 0xe7, 0x6f, 0x7c, 0x14, 0xb3, 0xc9,
	0x3e, 0x4c, 0xc8, 0x58, 0x0b, 0xa2, 0xea, 0x2b, 0x05, 0x7c, 0x7f, 0x16, 0x53, 0xe6, 0x04, 0x38,
	0xf6, 0x0f, 0x49, 0xd4, 0x0f, 0x11, 0xc3, 0x24, 0x3e, 0xc5, 0x11, 0xd2, 0xca, 0x0d, 0xa5, 0x59,
	0xf9, 0x65, 0xab, 0x25, 0xc3, 0xd2, 0x9a, 0x86, 0xa5, 0x75, 0x3a, 0x0d, 0x8b, 0xb9, 0x3b, 0x59,
	0x66, 0xc4, 0xe1, 0x5a, 0x3a, 0x95, 0xb0, 0x27, 0x49, 0x1a, 0x73, 0xb8, 0x29, 0x4f, 0xb3, 0xc8,
	0xf5, 0xeb, 0x77, 0x50, 0xb1, 0x3e, 0xe7, 0xb7, 0x5f, 0x7d, 0x39, 0x84, 0xb9, 0x9b, 0x21, 0x54,
	0x3e, 0x0c, 0xa1, 0xa2, 0xbf, 0x5e, 0x01, 0x1b, 0x4f, 0x11, 0x6c, 0x63, 0x3f, 0xc6, 0xb1, 0x7f,
	0x14, 0x77, 0x89, 0x7a, 0x0e, 0xa6, 0xa9, 0xcb, 0x82, 0xf8, 0xc7, 0x5c, 0x10, 0xbf, 0x32, 0x76,
	0xd9, 0xb4, 0xfa, 0x17, 0xa8, 0x52, 0xe6, 0x24, 0xcc, 0xee, 0x21, 0xec, 0xf7, 0x98, 0x48, 0x5e,
	0xde, 0xdc, 0x1e, 0x71, 0xb8, 0xc0, 0xc7, 0x1c, 0xd6, 0xe4, 0x07, 0xce, 0x53, 0xdd, 0xaa, 0x88,
	0xf2, 0x4f, 0x51, 0xa9, 0xbf, 0x81, 0xc2, 0x51, 0xec, 0xa1, 0x2b, 0x11, 0xbf, 0x4c, 0x04, 0x4f,
	0x80, 0x4d, 0xba, 0x5d, 0x8a, 0xe6, 0x44, 0xe6, 0xa9, 0x6e, 0xc9, 0x29, 0x35, 0x06, 0x55, 0x19,
	0x58, 0x3b, 0x8d, 0x19, 0x0e, 0x45, 0x58, 0xbf, 0x7c, 0x0d, 0x23, 0xbb, 0xc6, 0xc2, 0xdc, 0xcc,
	0x65, 0x9e, 0xca, 0x4b, 0x54, 0x24, 0x3a, 0x9b, 0x10, 0x35, 0x02, 0x9b, 0x11, 0xa6, 0x14, 0x79,
	0x76, 0x27, 0x24, 0x6e, 0x40, 0x6d, 0x97, 0xa4, 0x31, 0x43, 0x89, 0x56, 0x10, 0xeb, 0xef, 0x8d,
	0x38, 0x5c, 0xde, 0x30, 0xe6, 0xf0, 0x07, 0xe9, 0xb0, 0xf4, 0x59, 0xb7, 0x6a, 0x92, 0x9b, 0x02,
	0x1f, 0x4a, 0x3a, 0xb1, 0xcb, 0x16, 0xfa, 0xc4, 0xae, 0x38, 0xb3, 0x5b, 0xda, 0x30, 0xb3, 0x5b,
	0xfa, 0xac, 0x5b, 0x35, 0xc9, 0x17, 0xec, 0xf6, 0xcb, 0x37, 0x43, 0x98, 0x13, 0xb9, 0xfa, 0x11,
	0x14, 0x4e, 0xc8, 0x25, 0x4a, 0xd4, 0x0d, 0x50, 0xb8, 0x70, 0xc2, 0x14, 0x89, 0x14, 0xe5, 0x2d,
	0x59, 0xe8, 0xdb, 0x60, 0xfd, 0x29, 0x75, 0xff, 0x88, 0xbd, 0x17, 0x1b, 0xcb, 0x59, 0xa3, 0x79,
	0x7c, 0x7b, 0x5f, 0x57, 0xee, 0xee, 0xeb, 0xca, 0xfb, 0xfb, 0xba, 0x72, 0xfd, 0x50, 0xcf, 0xdd,
	0x3d, 0xd4, 0x73, 0x6f, 0x1e, 0xea, 0xb9, 0xf3, 0x67, 0x05, 0x70, 0xfa, 0xab, 0x2d, 0x82, 0xd8,
	0x29, 0x8a, 0x73, 0xee, 0x7e, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x40, 0xa4, 0x96, 0xb4, 0xcd, 0x05,
	0x00, 0x00,
}

func (this *ValidatorProto) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorProto)
	if !ok {
		that2, ok := that.(ValidatorProto)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if this.PublicKey != that1.PublicKey {
		return false
	}
	if this.Jailed != that1.Jailed {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if len(this.Chains) != len(that1.Chains) {
		return false
	}
	for i := range this.Chains {
		if this.Chains[i] != that1.Chains[i] {
			return false
		}
	}
	if this.ServiceURL != that1.ServiceURL {
		return false
	}
	if !this.StakedTokens.Equal(that1.StakedTokens) {
		return false
	}
	if !this.UnstakingCompletionTime.Equal(that1.UnstakingCompletionTime) {
		return false
	}
	return true
}
func (this *ValidatorSigningInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSigningInfo)
	if !ok {
		that2, ok := that.(ValidatorSigningInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if this.StartHeight != that1.StartHeight {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if !this.JailedUntil.Equal(that1.JailedUntil) {
		return false
	}
	if this.MissedBlocksCounter != that1.MissedBlocksCounter {
		return false
	}
	if this.JailedBlocksCounter != that1.JailedBlocksCounter {
		return false
	}
	return true
}
func (m *ValidatorProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UnstakingCompletionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.UnstakingCompletionTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintNodes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	{
		size := m.StakedTokens.Size()
		i -= size
		if _, err := m.StakedTokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNodes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.ServiceURL) > 0 {
		i -= len(m.ServiceURL)
		copy(dAtA[i:], m.ServiceURL)
		i = encodeVarintNodes(dAtA, i, uint64(len(m.ServiceURL)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Chains[iNdEx])
			copy(dAtA[i:], m.Chains[iNdEx])
			i = encodeVarintNodes(dAtA, i, uint64(len(m.Chains[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Status != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintNodes(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNodes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.JailedBlocksCounter != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.JailedBlocksCounter))
		i--
		dAtA[i] = 0x30
	}
	if m.MissedBlocksCounter != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.MissedBlocksCounter))
		i--
		dAtA[i] = 0x28
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.JailedUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintNodes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if m.Index != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if m.StartHeight != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNodes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Power) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Power) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Power) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorMissed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorMissed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorMissed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value {
		i--
		if m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNodes(dAtA []byte, offset int, v uint64) int {
	offset -= sovNodes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNodes(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovNodes(uint64(l))
	}
	if m.Jailed {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovNodes(uint64(m.Status))
	}
	if len(m.Chains) > 0 {
		for _, s := range m.Chains {
			l = len(s)
			n += 1 + l + sovNodes(uint64(l))
		}
	}
	l = len(m.ServiceURL)
	if l > 0 {
		n += 1 + l + sovNodes(uint64(l))
	}
	l = m.StakedTokens.Size()
	n += 1 + l + sovNodes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UnstakingCompletionTime)
	n += 1 + l + sovNodes(uint64(l))
	return n
}

func (m *ValidatorSigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNodes(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovNodes(uint64(m.StartHeight))
	}
	if m.Index != 0 {
		n += 1 + sovNodes(uint64(m.Index))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil)
	n += 1 + l + sovNodes(uint64(l))
	if m.MissedBlocksCounter != 0 {
		n += 1 + sovNodes(uint64(m.MissedBlocksCounter))
	}
	if m.JailedBlocksCounter != 0 {
		n += 1 + sovNodes(uint64(m.JailedBlocksCounter))
	}
	return n
}

func (m *Power) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovNodes(uint64(m.Value))
	}
	return n
}

func (m *ValidatorMissed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value {
		n += 2
	}
	return n
}

func sovNodes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNodes(x uint64) (n int) {
	return sovNodes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValidatorProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= github_com_pokt_network_pocket_core_types.StakeStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakedTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakingCompletionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UnstakingCompletionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ValidatorSigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValidatorSigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.JailedUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
			}
			m.MissedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedBlocksCounter", wireType)
			}
			m.JailedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JailedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNodes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Power) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Power: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Power: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNodes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ValidatorMissed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValidatorMissed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorMissed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNodes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNodes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNodes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNodes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNodes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNodes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNodes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNodes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNodes = fmt.Errorf("proto: unexpected end of group")
)
