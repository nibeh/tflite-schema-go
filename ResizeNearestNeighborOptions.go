// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ResizeNearestNeighborOptionsT struct {
	AlignCorners bool
	HalfPixelCenters bool
}

func (t *ResizeNearestNeighborOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ResizeNearestNeighborOptionsStart(builder)
	ResizeNearestNeighborOptionsAddAlignCorners(builder, t.AlignCorners)
	ResizeNearestNeighborOptionsAddHalfPixelCenters(builder, t.HalfPixelCenters)
	return ResizeNearestNeighborOptionsEnd(builder)
}

func (rcv *ResizeNearestNeighborOptions) UnPackTo(t *ResizeNearestNeighborOptionsT) {
	t.AlignCorners = rcv.AlignCorners()
	t.HalfPixelCenters = rcv.HalfPixelCenters()
}

func (rcv *ResizeNearestNeighborOptions) UnPack() *ResizeNearestNeighborOptionsT {
	if rcv == nil { return nil }
	t := &ResizeNearestNeighborOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type ResizeNearestNeighborOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsResizeNearestNeighborOptions(buf []byte, offset flatbuffers.UOffsetT) *ResizeNearestNeighborOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ResizeNearestNeighborOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ResizeNearestNeighborOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ResizeNearestNeighborOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ResizeNearestNeighborOptions) AlignCorners() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ResizeNearestNeighborOptions) MutateAlignCorners(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *ResizeNearestNeighborOptions) HalfPixelCenters() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ResizeNearestNeighborOptions) MutateHalfPixelCenters(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func ResizeNearestNeighborOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ResizeNearestNeighborOptionsAddAlignCorners(builder *flatbuffers.Builder, alignCorners bool) {
	builder.PrependBoolSlot(0, alignCorners, false)
}
func ResizeNearestNeighborOptionsAddHalfPixelCenters(builder *flatbuffers.Builder, halfPixelCenters bool) {
	builder.PrependBoolSlot(1, halfPixelCenters, false)
}
func ResizeNearestNeighborOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}