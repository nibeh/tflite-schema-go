// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GreaterEqualOptionsT struct {
}

func (t *GreaterEqualOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	GreaterEqualOptionsStart(builder)
	return GreaterEqualOptionsEnd(builder)
}

func (rcv *GreaterEqualOptions) UnPackTo(t *GreaterEqualOptionsT) {
}

func (rcv *GreaterEqualOptions) UnPack() *GreaterEqualOptionsT {
	if rcv == nil { return nil }
	t := &GreaterEqualOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type GreaterEqualOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsGreaterEqualOptions(buf []byte, offset flatbuffers.UOffsetT) *GreaterEqualOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GreaterEqualOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *GreaterEqualOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GreaterEqualOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func GreaterEqualOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func GreaterEqualOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
