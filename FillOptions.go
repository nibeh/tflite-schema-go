// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FillOptionsT struct {
}

func (t *FillOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	FillOptionsStart(builder)
	return FillOptionsEnd(builder)
}

func (rcv *FillOptions) UnPackTo(t *FillOptionsT) {
}

func (rcv *FillOptions) UnPack() *FillOptionsT {
	if rcv == nil { return nil }
	t := &FillOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type FillOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsFillOptions(buf []byte, offset flatbuffers.UOffsetT) *FillOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FillOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FillOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FillOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func FillOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func FillOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
