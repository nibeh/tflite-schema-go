// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BufferT struct {
	Data []byte
}

func (t *BufferT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	BufferStart(builder)
	BufferAddData(builder, dataOffset)
	return BufferEnd(builder)
}

func (rcv *Buffer) UnPackTo(t *BufferT) {
	t.Data = rcv.DataBytes()
}

func (rcv *Buffer) UnPack() *BufferT {
	if rcv == nil { return nil }
	t := &BufferT{}
	rcv.UnPackTo(t)
	return t
}

type Buffer struct {
	_tab flatbuffers.Table
}

func GetRootAsBuffer(buf []byte, offset flatbuffers.UOffsetT) *Buffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Buffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Buffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Buffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Buffer) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Buffer) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Buffer) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Buffer) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func BufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BufferAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func BufferStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func BufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
