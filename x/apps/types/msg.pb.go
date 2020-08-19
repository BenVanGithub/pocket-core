// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/apps/msg.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_pokt_network_pocket_core_types "github.com/pokt-network/pocket-core/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgAppStake struct {
	PubKey string                                        `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pubkey" yaml:"pubkey"`
	Chains []string                                      `protobuf:"bytes,2,rep,name=chains,proto3" json:"chains" yaml:"chains"`
	Value  github_com_pokt_network_pocket_core_types.Int `protobuf:"bytes,3,opt,name=value,proto3,customtype=github.com/pokt-network/pocket-core/types.Int" json:"value" yaml:"value"`
}

func (m *MsgAppStake) Reset()         { *m = MsgAppStake{} }
func (m *MsgAppStake) String() string { return proto.CompactTextString(m) }
func (*MsgAppStake) ProtoMessage()    {}
func (*MsgAppStake) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd58e5eb64f87460, []int{0}
}
func (m *MsgAppStake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAppStake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAppStake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAppStake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAppStake.Merge(m, src)
}
func (m *MsgAppStake) XXX_Size() int {
	return m.Size()
}
func (m *MsgAppStake) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAppStake.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAppStake proto.InternalMessageInfo

type MsgBeginAppUnstake struct {
	Address github_com_pokt_network_pocket_core_types.Address `protobuf:"bytes,1,opt,name=Address,proto3,casttype=github.com/pokt-network/pocket-core/types.Address" json:"address" yaml:"address"`
}

func (m *MsgBeginAppUnstake) Reset()         { *m = MsgBeginAppUnstake{} }
func (m *MsgBeginAppUnstake) String() string { return proto.CompactTextString(m) }
func (*MsgBeginAppUnstake) ProtoMessage()    {}
func (*MsgBeginAppUnstake) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd58e5eb64f87460, []int{1}
}
func (m *MsgBeginAppUnstake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBeginAppUnstake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBeginAppUnstake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBeginAppUnstake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBeginAppUnstake.Merge(m, src)
}
func (m *MsgBeginAppUnstake) XXX_Size() int {
	return m.Size()
}
func (m *MsgBeginAppUnstake) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBeginAppUnstake.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBeginAppUnstake proto.InternalMessageInfo

type MsgAppUnjail struct {
	AppAddr github_com_pokt_network_pocket_core_types.Address `protobuf:"bytes,1,opt,name=AppAddr,proto3,casttype=github.com/pokt-network/pocket-core/types.Address" json:"address" yaml:"address"`
}

func (m *MsgAppUnjail) Reset()         { *m = MsgAppUnjail{} }
func (m *MsgAppUnjail) String() string { return proto.CompactTextString(m) }
func (*MsgAppUnjail) ProtoMessage()    {}
func (*MsgAppUnjail) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd58e5eb64f87460, []int{2}
}
func (m *MsgAppUnjail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAppUnjail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAppUnjail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAppUnjail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAppUnjail.Merge(m, src)
}
func (m *MsgAppUnjail) XXX_Size() int {
	return m.Size()
}
func (m *MsgAppUnjail) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAppUnjail.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAppUnjail proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAppStake)(nil), "x.apps.MsgAppStake")
	proto.RegisterType((*MsgBeginAppUnstake)(nil), "x.apps.MsgBeginAppUnstake")
	proto.RegisterType((*MsgAppUnjail)(nil), "x.apps.MsgAppUnjail")
}

func init() { proto.RegisterFile("x/apps/msg.proto", fileDescriptor_fd58e5eb64f87460) }

var fileDescriptor_fd58e5eb64f87460 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0x6f, 0x9b, 0x40,
	0x14, 0xc6, 0xb9, 0x5a, 0xc5, 0xea, 0x95, 0x56, 0x15, 0xea, 0x60, 0xb5, 0x12, 0x67, 0x31, 0x75,
	0x31, 0xb4, 0x72, 0x27, 0x6f, 0x66, 0x6b, 0x2b, 0x77, 0xa0, 0xf2, 0xd2, 0x25, 0x02, 0x7c, 0xc2,
	0x04, 0xcc, 0x9d, 0xb8, 0x23, 0x31, 0xff, 0x41, 0x94, 0xc9, 0x63, 0x46, 0xff, 0x39, 0x1e, 0x3d,
	0x46, 0x19, 0x4e, 0x91, 0xbd, 0x44, 0x56, 0x26, 0x8f, 0x99, 0x22, 0xee, 0xb0, 0xa2, 0x64, 0xf2,
	0x92, 0x8d, 0xf7, 0x3d, 0xde, 0xf7, 0xfb, 0x9e, 0xee, 0xc1, 0x4f, 0x73, 0x37, 0xa0, 0x94, 0xb9,
	0x33, 0x16, 0x3b, 0xb4, 0x20, 0x9c, 0x98, 0xfa, 0xdc, 0xa9, 0x95, 0x2f, 0x9f, 0x63, 0x12, 0x13,
	0x29, 0xb9, 0xf5, 0x97, 0xea, 0xda, 0xf7, 0x00, 0xbe, 0x1f, 0xb1, 0x78, 0x48, 0xe9, 0x3f, 0x1e,
	0xa4, 0xd8, 0xfc, 0x09, 0xdb, 0xb4, 0x0c, 0x4f, 0x52, 0x5c, 0x75, 0x40, 0x17, 0x7c, 0x7b, 0xe7,
	0x7d, 0xdd, 0x09, 0xa4, 0xd3, 0x32, 0x4c, 0x71, 0xb5, 0x17, 0xe8, 0x43, 0x15, 0xcc, 0xb2, 0x81,
	0xad, 0x6a, 0xdb, 0xaf, 0x1b, 0x7f, 0x70, 0x65, 0xf6, 0xa1, 0x1e, 0x4d, 0x83, 0x24, 0x67, 0x9d,
	0x37, 0xdd, 0xd6, 0x61, 0x48, 0x29, 0x4f, 0x43, 0xaa, 0xb6, 0xfd, 0xa6, 0x61, 0x4e, 0xe0, 0xdb,
	0xb3, 0x20, 0x2b, 0x71, 0xa7, 0x25, 0x41, 0x7f, 0x57, 0x02, 0x69, 0x37, 0x02, 0xf5, 0xe2, 0x84,
	0x4f, 0xcb, 0xd0, 0x89, 0xc8, 0xcc, 0xa5, 0x24, 0xe5, 0xbd, 0x1c, 0xf3, 0x73, 0x52, 0xa4, 0x2e,
	0x25, 0x51, 0x8a, 0x79, 0x2f, 0x22, 0x05, 0x76, 0x79, 0x45, 0x31, 0x73, 0x7e, 0xe5, 0x7c, 0x27,
	0x90, 0x72, 0xd9, 0x0b, 0x64, 0x28, 0x8e, 0x2c, 0x6d, 0x5f, 0xc9, 0x03, 0xe3, 0x62, 0x89, 0xb4,
	0xab, 0x25, 0x02, 0x77, 0x4b, 0x04, 0xec, 0x05, 0x80, 0xe6, 0x88, 0xc5, 0x1e, 0x8e, 0x93, 0x7c,
	0x48, 0xe9, 0x38, 0x67, 0x72, 0xeb, 0x0c, 0xb6, 0x87, 0x93, 0x49, 0x81, 0x19, 0x93, 0x5b, 0x1b,
	0x9e, 0xbf, 0x13, 0xa8, 0x1d, 0x28, 0x69, 0x2f, 0xd0, 0x47, 0xe5, 0xdc, 0x08, 0xf6, 0x83, 0x40,
	0x3f, 0x8e, 0x4f, 0xd9, 0x38, 0xfb, 0x07, 0xc4, 0x8b, 0x48, 0x97, 0x00, 0x1a, 0xea, 0x05, 0xc6,
	0xf9, 0x69, 0x90, 0x64, 0x32, 0x0c, 0xa5, 0xf5, 0xcf, 0xaf, 0x1a, 0x46, 0x21, 0x9e, 0x87, 0xf1,
	0x7e, 0xaf, 0x36, 0x16, 0x58, 0x6f, 0x2c, 0x70, 0xbb, 0xb1, 0xc0, 0x62, 0x6b, 0x69, 0xeb, 0xad,
	0xa5, 0x5d, 0x6f, 0x2d, 0xed, 0xff, 0xf7, 0x63, 0x18, 0xcd, 0xf1, 0x49, 0x54, 0xa8, 0xcb, 0x0b,
	0xeb, 0x3f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x07, 0x26, 0x90, 0x4a, 0x93, 0x02, 0x00, 0x00,
}

func (this *MsgAppStake) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgAppStake)
	if !ok {
		that2, ok := that.(MsgAppStake)
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
	if this.PubKey != that1.PubKey {
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
	if !this.Value.Equal(that1.Value) {
		return false
	}
	return true
}
func (this *MsgBeginAppUnstake) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgBeginAppUnstake)
	if !ok {
		that2, ok := that.(MsgBeginAppUnstake)
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
	return true
}
func (this *MsgAppUnjail) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgAppUnjail)
	if !ok {
		that2, ok := that.(MsgAppUnjail)
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
	if !bytes.Equal(this.AppAddr, that1.AppAddr) {
		return false
	}
	return true
}
func (m *MsgAppStake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAppStake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAppStake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Value.Size()
		i -= size
		if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Chains[iNdEx])
			copy(dAtA[i:], m.Chains[iNdEx])
			i = encodeVarintMsg(dAtA, i, uint64(len(m.Chains[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBeginAppUnstake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBeginAppUnstake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBeginAppUnstake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAppUnjail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAppUnjail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAppUnjail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AppAddr) > 0 {
		i -= len(m.AppAddr)
		copy(dAtA[i:], m.AppAddr)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.AppAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAppStake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if len(m.Chains) > 0 {
		for _, s := range m.Chains {
			l = len(s)
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	l = m.Value.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgBeginAppUnstake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgAppUnjail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AppAddr)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAppStake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgAppStake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAppStake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgBeginAppUnstake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgBeginAppUnstake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBeginAppUnstake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgAppUnjail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgAppUnjail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAppUnjail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppAddr = append(m.AppAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.AppAddr == nil {
				m.AppAddr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)