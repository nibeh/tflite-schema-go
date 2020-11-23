// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SplitOptionsT struct {
	NumSplits int32
}

func (t *SplitOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SplitOptionsStart(builder)
	SplitOptionsAddNumSplits(builder, t.NumSplits)
	return SplitOptionsEnd(builder)
}

func (rcv *SplitOptions) UnPackTo(t *SplitOptionsT) {
	t.NumSplits = rcv.NumSplits()
}

func (rcv *SplitOptions) UnPack() *SplitOptionsT {
	if rcv == nil { return nil }
	t := &SplitOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type SplitOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsSplitOptions(buf []byte, offset flatbuffers.UOffsetT) *SplitOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SplitOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SplitOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SplitOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SplitOptions) NumSplits() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SplitOptions) MutateNumSplits(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func SplitOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SplitOptionsAddNumSplits(builder *flatbuffers.Builder, numSplits int32) {
	builder.PrependInt32Slot(0, numSplits, 0)
}
func SplitOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
