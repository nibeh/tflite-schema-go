// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SqueezeOptionsT struct {
	SqueezeDims []int32
}

func (t *SqueezeOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	squeezeDimsOffset := flatbuffers.UOffsetT(0)
	if t.SqueezeDims != nil {
		squeezeDimsLength := len(t.SqueezeDims)
		SqueezeOptionsStartSqueezeDimsVector(builder, squeezeDimsLength)
		for j := squeezeDimsLength - 1; j >= 0; j-- {
			builder.PrependInt32(t.SqueezeDims[j])
		}
		squeezeDimsOffset = builder.EndVector(squeezeDimsLength)
	}
	SqueezeOptionsStart(builder)
	SqueezeOptionsAddSqueezeDims(builder, squeezeDimsOffset)
	return SqueezeOptionsEnd(builder)
}

func (rcv *SqueezeOptions) UnPackTo(t *SqueezeOptionsT) {
	squeezeDimsLength := rcv.SqueezeDimsLength()
	t.SqueezeDims = make([]int32, squeezeDimsLength)
	for j := 0; j < squeezeDimsLength; j++ {
		t.SqueezeDims[j] = rcv.SqueezeDims(j)
	}
}

func (rcv *SqueezeOptions) UnPack() *SqueezeOptionsT {
	if rcv == nil { return nil }
	t := &SqueezeOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type SqueezeOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsSqueezeOptions(buf []byte, offset flatbuffers.UOffsetT) *SqueezeOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SqueezeOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SqueezeOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SqueezeOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SqueezeOptions) SqueezeDims(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *SqueezeOptions) SqueezeDimsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SqueezeOptions) MutateSqueezeDims(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func SqueezeOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SqueezeOptionsAddSqueezeDims(builder *flatbuffers.Builder, squeezeDims flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(squeezeDims), 0)
}
func SqueezeOptionsStartSqueezeDimsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SqueezeOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
