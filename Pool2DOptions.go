// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Pool2DOptionsT struct {
	Padding Padding
	StrideW int32
	StrideH int32
	FilterWidth int32
	FilterHeight int32
	FusedActivationFunction ActivationFunctionType
}

func (t *Pool2DOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Pool2DOptionsStart(builder)
	Pool2DOptionsAddPadding(builder, t.Padding)
	Pool2DOptionsAddStrideW(builder, t.StrideW)
	Pool2DOptionsAddStrideH(builder, t.StrideH)
	Pool2DOptionsAddFilterWidth(builder, t.FilterWidth)
	Pool2DOptionsAddFilterHeight(builder, t.FilterHeight)
	Pool2DOptionsAddFusedActivationFunction(builder, t.FusedActivationFunction)
	return Pool2DOptionsEnd(builder)
}

func (rcv *Pool2DOptions) UnPackTo(t *Pool2DOptionsT) {
	t.Padding = rcv.Padding()
	t.StrideW = rcv.StrideW()
	t.StrideH = rcv.StrideH()
	t.FilterWidth = rcv.FilterWidth()
	t.FilterHeight = rcv.FilterHeight()
	t.FusedActivationFunction = rcv.FusedActivationFunction()
}

func (rcv *Pool2DOptions) UnPack() *Pool2DOptionsT {
	if rcv == nil { return nil }
	t := &Pool2DOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type Pool2DOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsPool2DOptions(buf []byte, offset flatbuffers.UOffsetT) *Pool2DOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Pool2DOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Pool2DOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Pool2DOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Pool2DOptions) Padding() Padding {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Padding(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Pool2DOptions) MutatePadding(n Padding) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *Pool2DOptions) StrideW() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Pool2DOptions) MutateStrideW(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Pool2DOptions) StrideH() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Pool2DOptions) MutateStrideH(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Pool2DOptions) FilterWidth() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Pool2DOptions) MutateFilterWidth(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Pool2DOptions) FilterHeight() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Pool2DOptions) MutateFilterHeight(n int32) bool {
	return rcv._tab.MutateInt32Slot(12, n)
}

func (rcv *Pool2DOptions) FusedActivationFunction() ActivationFunctionType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return ActivationFunctionType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Pool2DOptions) MutateFusedActivationFunction(n ActivationFunctionType) bool {
	return rcv._tab.MutateInt8Slot(14, int8(n))
}

func Pool2DOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func Pool2DOptionsAddPadding(builder *flatbuffers.Builder, padding Padding) {
	builder.PrependInt8Slot(0, int8(padding), 0)
}
func Pool2DOptionsAddStrideW(builder *flatbuffers.Builder, strideW int32) {
	builder.PrependInt32Slot(1, strideW, 0)
}
func Pool2DOptionsAddStrideH(builder *flatbuffers.Builder, strideH int32) {
	builder.PrependInt32Slot(2, strideH, 0)
}
func Pool2DOptionsAddFilterWidth(builder *flatbuffers.Builder, filterWidth int32) {
	builder.PrependInt32Slot(3, filterWidth, 0)
}
func Pool2DOptionsAddFilterHeight(builder *flatbuffers.Builder, filterHeight int32) {
	builder.PrependInt32Slot(4, filterHeight, 0)
}
func Pool2DOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction ActivationFunctionType) {
	builder.PrependInt8Slot(5, int8(fusedActivationFunction), 0)
}
func Pool2DOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}