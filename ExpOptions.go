// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ExpOptionsT struct {
}

func (t *ExpOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ExpOptionsStart(builder)
	return ExpOptionsEnd(builder)
}

func (rcv *ExpOptions) UnPackTo(t *ExpOptionsT) {
}

func (rcv *ExpOptions) UnPack() *ExpOptionsT {
	if rcv == nil { return nil }
	t := &ExpOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type ExpOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsExpOptions(buf []byte, offset flatbuffers.UOffsetT) *ExpOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ExpOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ExpOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ExpOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func ExpOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func ExpOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}