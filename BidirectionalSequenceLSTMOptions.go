// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BidirectionalSequenceLSTMOptionsT struct {
	FusedActivationFunction ActivationFunctionType
	CellClip float32
	ProjClip float32
	MergeOutputs bool
	TimeMajor bool
	AsymmetricQuantizeInputs bool
}

func (t *BidirectionalSequenceLSTMOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	BidirectionalSequenceLSTMOptionsStart(builder)
	BidirectionalSequenceLSTMOptionsAddFusedActivationFunction(builder, t.FusedActivationFunction)
	BidirectionalSequenceLSTMOptionsAddCellClip(builder, t.CellClip)
	BidirectionalSequenceLSTMOptionsAddProjClip(builder, t.ProjClip)
	BidirectionalSequenceLSTMOptionsAddMergeOutputs(builder, t.MergeOutputs)
	BidirectionalSequenceLSTMOptionsAddTimeMajor(builder, t.TimeMajor)
	BidirectionalSequenceLSTMOptionsAddAsymmetricQuantizeInputs(builder, t.AsymmetricQuantizeInputs)
	return BidirectionalSequenceLSTMOptionsEnd(builder)
}

func (rcv *BidirectionalSequenceLSTMOptions) UnPackTo(t *BidirectionalSequenceLSTMOptionsT) {
	t.FusedActivationFunction = rcv.FusedActivationFunction()
	t.CellClip = rcv.CellClip()
	t.ProjClip = rcv.ProjClip()
	t.MergeOutputs = rcv.MergeOutputs()
	t.TimeMajor = rcv.TimeMajor()
	t.AsymmetricQuantizeInputs = rcv.AsymmetricQuantizeInputs()
}

func (rcv *BidirectionalSequenceLSTMOptions) UnPack() *BidirectionalSequenceLSTMOptionsT {
	if rcv == nil { return nil }
	t := &BidirectionalSequenceLSTMOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type BidirectionalSequenceLSTMOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsBidirectionalSequenceLSTMOptions(buf []byte, offset flatbuffers.UOffsetT) *BidirectionalSequenceLSTMOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BidirectionalSequenceLSTMOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *BidirectionalSequenceLSTMOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BidirectionalSequenceLSTMOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BidirectionalSequenceLSTMOptions) FusedActivationFunction() ActivationFunctionType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return ActivationFunctionType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *BidirectionalSequenceLSTMOptions) MutateFusedActivationFunction(n ActivationFunctionType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *BidirectionalSequenceLSTMOptions) CellClip() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *BidirectionalSequenceLSTMOptions) MutateCellClip(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *BidirectionalSequenceLSTMOptions) ProjClip() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *BidirectionalSequenceLSTMOptions) MutateProjClip(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *BidirectionalSequenceLSTMOptions) MergeOutputs() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *BidirectionalSequenceLSTMOptions) MutateMergeOutputs(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func (rcv *BidirectionalSequenceLSTMOptions) TimeMajor() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *BidirectionalSequenceLSTMOptions) MutateTimeMajor(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func (rcv *BidirectionalSequenceLSTMOptions) AsymmetricQuantizeInputs() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *BidirectionalSequenceLSTMOptions) MutateAsymmetricQuantizeInputs(n bool) bool {
	return rcv._tab.MutateBoolSlot(14, n)
}

func BidirectionalSequenceLSTMOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func BidirectionalSequenceLSTMOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction ActivationFunctionType) {
	builder.PrependInt8Slot(0, int8(fusedActivationFunction), 0)
}
func BidirectionalSequenceLSTMOptionsAddCellClip(builder *flatbuffers.Builder, cellClip float32) {
	builder.PrependFloat32Slot(1, cellClip, 0.0)
}
func BidirectionalSequenceLSTMOptionsAddProjClip(builder *flatbuffers.Builder, projClip float32) {
	builder.PrependFloat32Slot(2, projClip, 0.0)
}
func BidirectionalSequenceLSTMOptionsAddMergeOutputs(builder *flatbuffers.Builder, mergeOutputs bool) {
	builder.PrependBoolSlot(3, mergeOutputs, false)
}
func BidirectionalSequenceLSTMOptionsAddTimeMajor(builder *flatbuffers.Builder, timeMajor bool) {
	builder.PrependBoolSlot(4, timeMajor, true)
}
func BidirectionalSequenceLSTMOptionsAddAsymmetricQuantizeInputs(builder *flatbuffers.Builder, asymmetricQuantizeInputs bool) {
	builder.PrependBoolSlot(5, asymmetricQuantizeInputs, false)
}
func BidirectionalSequenceLSTMOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
