// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/address.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Addresses struct {
	Arr addresses `protobuf:"bytes,1,rep,name=arr,proto3,castrepeated=addresses" json:"arr,omitempty"`
}

func (m *Addresses) Reset()         { *m = Addresses{} }
func (m *Addresses) String() string { return proto.CompactTextString(m) }
func (*Addresses) ProtoMessage()    {}
func (*Addresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_468dd63a469488c9, []int{0}
}
func (m *Addresses) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Addresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Addresses.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Addresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addresses.Merge(m, src)
}
func (m *Addresses) XXX_Size() int {
	return m.Size()
}
func (m *Addresses) XXX_DiscardUnknown() {
	xxx_messageInfo_Addresses.DiscardUnknown(m)
}

var xxx_messageInfo_Addresses proto.InternalMessageInfo

func (m *Addresses) GetArr() addresses {
	if m != nil {
		return m.Arr
	}
	return nil
}

type AddressProto struct {
	Address Address `protobuf:"bytes,1,opt,name=address,proto3,casttype=Address" json:"address,omitempty"`
}

func (m *AddressProto) Reset()         { *m = AddressProto{} }
func (m *AddressProto) String() string { return proto.CompactTextString(m) }
func (*AddressProto) ProtoMessage()    {}
func (*AddressProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_468dd63a469488c9, []int{1}
}
func (m *AddressProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddressProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddressProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddressProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressProto.Merge(m, src)
}
func (m *AddressProto) XXX_Size() int {
	return m.Size()
}
func (m *AddressProto) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressProto.DiscardUnknown(m)
}

var xxx_messageInfo_AddressProto proto.InternalMessageInfo

func (m *AddressProto) GetAddress() Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type BoolProto struct {
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BoolProto) Reset()         { *m = BoolProto{} }
func (m *BoolProto) String() string { return proto.CompactTextString(m) }
func (*BoolProto) ProtoMessage()    {}
func (*BoolProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_468dd63a469488c9, []int{2}
}
func (m *BoolProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BoolProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BoolProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BoolProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolProto.Merge(m, src)
}
func (m *BoolProto) XXX_Size() int {
	return m.Size()
}
func (m *BoolProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolProto.DiscardUnknown(m)
}

var xxx_messageInfo_BoolProto proto.InternalMessageInfo

func (m *BoolProto) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*Addresses)(nil), "types.Addresses")
	proto.RegisterType((*AddressProto)(nil), "types.AddressProto")
	proto.RegisterType((*BoolProto)(nil), "types.BoolProto")
}

func init() { proto.RegisterFile("types/address.proto", fileDescriptor_468dd63a469488c9) }

var fileDescriptor_468dd63a469488c9 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x0b, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x45, 0xf4, 0x41, 0x2c, 0x88, 0xa4,
	0x92, 0x0e, 0x17, 0xa7, 0x23, 0x44, 0x75, 0x6a, 0xb1, 0x90, 0x3c, 0x17, 0x73, 0x62, 0x51, 0x91,
	0x04, 0xa3, 0x02, 0xb3, 0x06, 0x8f, 0x13, 0xef, 0xaa, 0xfb, 0xf2, 0x9c, 0x89, 0x30, 0xb9, 0x20,
	0x90, 0x8c, 0x92, 0x29, 0x17, 0x0f, 0x54, 0x75, 0x00, 0xd8, 0x68, 0x55, 0x2e, 0x76, 0xa8, 0x0a,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x1e, 0x27, 0xee, 0x5f, 0xf7, 0xe4, 0xd9, 0xa1, 0x4a, 0x82, 0x60,
	0x72, 0x4a, 0x8a, 0x5c, 0x9c, 0x4e, 0xf9, 0xf9, 0x39, 0x10, 0x3d, 0x22, 0x5c, 0xac, 0x65, 0x89,
	0x39, 0xa5, 0xa9, 0x60, 0x1d, 0x1c, 0x41, 0x10, 0x8e, 0x93, 0xf3, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c,
	0x37, 0x1e, 0xcb, 0x31, 0x44, 0x69, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0x17, 0xe4, 0x67, 0x97, 0xe8, 0xe6, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0xeb, 0x17, 0xe4,
	0x27, 0x67, 0xa7, 0x96, 0xe8, 0x26, 0xe7, 0x17, 0xa5, 0xea, 0x83, 0xbd, 0x98, 0xc4, 0x06, 0xf6,
	0x93, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x67, 0x4d, 0x3d, 0xfc, 0x07, 0x01, 0x00, 0x00,
}

func (m *Addresses) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Addresses) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Addresses) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Arr) > 0 {
		for iNdEx := len(m.Arr) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Arr[iNdEx])
			copy(dAtA[i:], m.Arr[iNdEx])
			i = encodeVarintAddress(dAtA, i, uint64(len(m.Arr[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AddressProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddressProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BoolProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BoolProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BoolProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintAddress(dAtA []byte, offset int, v uint64) int {
	offset -= sovAddress(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Addresses) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Arr) > 0 {
		for _, b := range m.Arr {
			l = len(b)
			n += 1 + l + sovAddress(uint64(l))
		}
	}
	return n
}

func (m *AddressProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	return n
}

func (m *BoolProto) Size() (n int) {
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

func sovAddress(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAddress(x uint64) (n int) {
	return sovAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Addresses) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
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
			return fmt.Errorf("proto: Addresses: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Addresses: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Arr = append(m.Arr, make([]byte, postIndex-iNdEx))
			copy(m.Arr[len(m.Arr)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAddress
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAddress
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
func (m *AddressProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
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
			return fmt.Errorf("proto: AddressProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
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
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAddress
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAddress
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
func (m *BoolProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
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
			return fmt.Errorf("proto: BoolProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BoolProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAddress
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAddress
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
func skipAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAddress
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
					return 0, ErrIntOverflowAddress
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
					return 0, ErrIntOverflowAddress
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
				return 0, ErrInvalidLengthAddress
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAddress
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAddress
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAddress        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAddress          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAddress = fmt.Errorf("proto: unexpected end of group")
)
