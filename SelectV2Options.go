// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SelectV2OptionsT struct {
}

func (t *SelectV2OptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SelectV2OptionsStart(builder)
	return SelectV2OptionsEnd(builder)
}

func (rcv *SelectV2Options) UnPackTo(t *SelectV2OptionsT) {
}

func (rcv *SelectV2Options) UnPack() *SelectV2OptionsT {
	if rcv == nil { return nil }
	t := &SelectV2OptionsT{}
	rcv.UnPackTo(t)
	return t
}

type SelectV2Options struct {
	_tab flatbuffers.Table
}

func GetRootAsSelectV2Options(buf []byte, offset flatbuffers.UOffsetT) *SelectV2Options {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SelectV2Options{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SelectV2Options) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SelectV2Options) Table() flatbuffers.Table {
	return rcv._tab
}

func SelectV2OptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func SelectV2OptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}