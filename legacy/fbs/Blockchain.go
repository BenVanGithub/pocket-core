// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Blockchain struct {
	_tab flatbuffers.Table
}

func GetRootAsBlockchain(buf []byte, offset flatbuffers.UOffsetT) *Blockchain {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Blockchain{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Blockchain) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Blockchain) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Blockchain) Name(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Blockchain) NameLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Blockchain) NameBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Blockchain) Netid() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Blockchain) MutateNetid(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *Blockchain) Version() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Blockchain) MutateVersion(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func BlockchainStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func BlockchainAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func BlockchainStartNameVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func BlockchainAddNetid(builder *flatbuffers.Builder, netid uint32) {
	builder.PrependUint32Slot(1, netid, 0)
}
func BlockchainAddVersion(builder *flatbuffers.Builder, version uint32) {
	builder.PrependUint32Slot(2, version, 0)
}
func BlockchainEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}