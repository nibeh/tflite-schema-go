// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OneHotOptionsT struct {
	Axis int32
}

func (t *OneHotOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	OneHotOptionsStart(builder)
	OneHotOptionsAddAxis(builder, t.Axis)
	return OneHotOptionsEnd(builder)
}

func (rcv *OneHotOptions) UnPackTo(t *OneHotOptionsT) {
	t.Axis = rcv.Axis()
}

func (rcv *OneHotOptions) UnPack() *OneHotOptionsT {
	if rcv == nil { return nil }
	t := &OneHotOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type OneHotOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsOneHotOptions(buf []byte, offset flatbuffers.UOffsetT) *OneHotOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OneHotOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OneHotOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OneHotOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OneHotOptions) Axis() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OneHotOptions) MutateAxis(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func OneHotOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func OneHotOptionsAddAxis(builder *flatbuffers.Builder, axis int32) {
	builder.PrependInt32Slot(0, axis, 0)
}
func OneHotOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
