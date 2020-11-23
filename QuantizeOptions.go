// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type QuantizeOptionsT struct {
}

func (t *QuantizeOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	QuantizeOptionsStart(builder)
	return QuantizeOptionsEnd(builder)
}

func (rcv *QuantizeOptions) UnPackTo(t *QuantizeOptionsT) {
}

func (rcv *QuantizeOptions) UnPack() *QuantizeOptionsT {
	if rcv == nil { return nil }
	t := &QuantizeOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type QuantizeOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsQuantizeOptions(buf []byte, offset flatbuffers.UOffsetT) *QuantizeOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &QuantizeOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *QuantizeOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *QuantizeOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func QuantizeOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func QuantizeOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}