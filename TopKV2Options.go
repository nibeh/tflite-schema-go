// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TopKV2OptionsT struct {
}

func (t *TopKV2OptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	TopKV2OptionsStart(builder)
	return TopKV2OptionsEnd(builder)
}

func (rcv *TopKV2Options) UnPackTo(t *TopKV2OptionsT) {
}

func (rcv *TopKV2Options) UnPack() *TopKV2OptionsT {
	if rcv == nil { return nil }
	t := &TopKV2OptionsT{}
	rcv.UnPackTo(t)
	return t
}

type TopKV2Options struct {
	_tab flatbuffers.Table
}

func GetRootAsTopKV2Options(buf []byte, offset flatbuffers.UOffsetT) *TopKV2Options {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TopKV2Options{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TopKV2Options) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TopKV2Options) Table() flatbuffers.Table {
	return rcv._tab
}

func TopKV2OptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func TopKV2OptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
