// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RangeOptionsT struct {
}

func (t *RangeOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	RangeOptionsStart(builder)
	return RangeOptionsEnd(builder)
}

func (rcv *RangeOptions) UnPackTo(t *RangeOptionsT) {
}

func (rcv *RangeOptions) UnPack() *RangeOptionsT {
	if rcv == nil { return nil }
	t := &RangeOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type RangeOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsRangeOptions(buf []byte, offset flatbuffers.UOffsetT) *RangeOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RangeOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RangeOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RangeOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func RangeOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func RangeOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
