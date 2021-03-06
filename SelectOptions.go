// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SelectOptionsT struct {
}

func (t *SelectOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SelectOptionsStart(builder)
	return SelectOptionsEnd(builder)
}

func (rcv *SelectOptions) UnPackTo(t *SelectOptionsT) {
}

func (rcv *SelectOptions) UnPack() *SelectOptionsT {
	if rcv == nil { return nil }
	t := &SelectOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type SelectOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsSelectOptions(buf []byte, offset flatbuffers.UOffsetT) *SelectOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SelectOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SelectOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SelectOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func SelectOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func SelectOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
