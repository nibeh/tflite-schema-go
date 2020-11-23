// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CallOptionsT struct {
	Subgraph uint32
}

func (t *CallOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	CallOptionsStart(builder)
	CallOptionsAddSubgraph(builder, t.Subgraph)
	return CallOptionsEnd(builder)
}

func (rcv *CallOptions) UnPackTo(t *CallOptionsT) {
	t.Subgraph = rcv.Subgraph()
}

func (rcv *CallOptions) UnPack() *CallOptionsT {
	if rcv == nil { return nil }
	t := &CallOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type CallOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsCallOptions(buf []byte, offset flatbuffers.UOffsetT) *CallOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CallOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *CallOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CallOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CallOptions) Subgraph() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CallOptions) MutateSubgraph(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func CallOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func CallOptionsAddSubgraph(builder *flatbuffers.Builder, subgraph uint32) {
	builder.PrependUint32Slot(0, subgraph, 0)
}
func CallOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
