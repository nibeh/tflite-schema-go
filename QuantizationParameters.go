// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type QuantizationParametersT struct {
	Min []float32
	Max []float32
	Scale []float32
	ZeroPoint []int64
	Details *QuantizationDetailsT
	QuantizedDimension int32
}

func (t *QuantizationParametersT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	minOffset := flatbuffers.UOffsetT(0)
	if t.Min != nil {
		minLength := len(t.Min)
		QuantizationParametersStartMinVector(builder, minLength)
		for j := minLength - 1; j >= 0; j-- {
			builder.PrependFloat32(t.Min[j])
		}
		minOffset = builder.EndVector(minLength)
	}
	maxOffset := flatbuffers.UOffsetT(0)
	if t.Max != nil {
		maxLength := len(t.Max)
		QuantizationParametersStartMaxVector(builder, maxLength)
		for j := maxLength - 1; j >= 0; j-- {
			builder.PrependFloat32(t.Max[j])
		}
		maxOffset = builder.EndVector(maxLength)
	}
	scaleOffset := flatbuffers.UOffsetT(0)
	if t.Scale != nil {
		scaleLength := len(t.Scale)
		QuantizationParametersStartScaleVector(builder, scaleLength)
		for j := scaleLength - 1; j >= 0; j-- {
			builder.PrependFloat32(t.Scale[j])
		}
		scaleOffset = builder.EndVector(scaleLength)
	}
	zeroPointOffset := flatbuffers.UOffsetT(0)
	if t.ZeroPoint != nil {
		zeroPointLength := len(t.ZeroPoint)
		QuantizationParametersStartZeroPointVector(builder, zeroPointLength)
		for j := zeroPointLength - 1; j >= 0; j-- {
			builder.PrependInt64(t.ZeroPoint[j])
		}
		zeroPointOffset = builder.EndVector(zeroPointLength)
	}
	detailsOffset := t.Details.Pack(builder)
	
	QuantizationParametersStart(builder)
	QuantizationParametersAddMin(builder, minOffset)
	QuantizationParametersAddMax(builder, maxOffset)
	QuantizationParametersAddScale(builder, scaleOffset)
	QuantizationParametersAddZeroPoint(builder, zeroPointOffset)
	if t.Details != nil {
		QuantizationParametersAddDetailsType(builder, t.Details.Type)
	}
	QuantizationParametersAddDetails(builder, detailsOffset)
	QuantizationParametersAddQuantizedDimension(builder, t.QuantizedDimension)
	return QuantizationParametersEnd(builder)
}

func (rcv *QuantizationParameters) UnPackTo(t *QuantizationParametersT) {
	minLength := rcv.MinLength()
	t.Min = make([]float32, minLength)
	for j := 0; j < minLength; j++ {
		t.Min[j] = rcv.Min(j)
	}
	maxLength := rcv.MaxLength()
	t.Max = make([]float32, maxLength)
	for j := 0; j < maxLength; j++ {
		t.Max[j] = rcv.Max(j)
	}
	scaleLength := rcv.ScaleLength()
	t.Scale = make([]float32, scaleLength)
	for j := 0; j < scaleLength; j++ {
		t.Scale[j] = rcv.Scale(j)
	}
	zeroPointLength := rcv.ZeroPointLength()
	t.ZeroPoint = make([]int64, zeroPointLength)
	for j := 0; j < zeroPointLength; j++ {
		t.ZeroPoint[j] = rcv.ZeroPoint(j)
	}
	detailsTable := flatbuffers.Table{}
	if rcv.Details(&detailsTable) {
		t.Details = rcv.DetailsType().UnPack(detailsTable)
	}
	t.QuantizedDimension = rcv.QuantizedDimension()
}

func (rcv *QuantizationParameters) UnPack() *QuantizationParametersT {
	if rcv == nil { return nil }
	t := &QuantizationParametersT{}
	rcv.UnPackTo(t)
	return t
}

type QuantizationParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsQuantizationParameters(buf []byte, offset flatbuffers.UOffsetT) *QuantizationParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &QuantizationParameters{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *QuantizationParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *QuantizationParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *QuantizationParameters) Min(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *QuantizationParameters) MinLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *QuantizationParameters) MutateMin(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *QuantizationParameters) Max(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *QuantizationParameters) MaxLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *QuantizationParameters) MutateMax(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *QuantizationParameters) Scale(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *QuantizationParameters) ScaleLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *QuantizationParameters) MutateScale(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *QuantizationParameters) ZeroPoint(j int) int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *QuantizationParameters) ZeroPointLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *QuantizationParameters) MutateZeroPoint(j int, n int64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *QuantizationParameters) DetailsType() QuantizationDetails {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return QuantizationDetails(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *QuantizationParameters) MutateDetailsType(n QuantizationDetails) bool {
	return rcv._tab.MutateByteSlot(12, byte(n))
}

func (rcv *QuantizationParameters) Details(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *QuantizationParameters) QuantizedDimension() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *QuantizationParameters) MutateQuantizedDimension(n int32) bool {
	return rcv._tab.MutateInt32Slot(16, n)
}

func QuantizationParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func QuantizationParametersAddMin(builder *flatbuffers.Builder, min flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(min), 0)
}
func QuantizationParametersStartMinVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func QuantizationParametersAddMax(builder *flatbuffers.Builder, max flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(max), 0)
}
func QuantizationParametersStartMaxVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func QuantizationParametersAddScale(builder *flatbuffers.Builder, scale flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(scale), 0)
}
func QuantizationParametersStartScaleVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func QuantizationParametersAddZeroPoint(builder *flatbuffers.Builder, zeroPoint flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(zeroPoint), 0)
}
func QuantizationParametersStartZeroPointVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func QuantizationParametersAddDetailsType(builder *flatbuffers.Builder, detailsType QuantizationDetails) {
	builder.PrependByteSlot(4, byte(detailsType), 0)
}
func QuantizationParametersAddDetails(builder *flatbuffers.Builder, details flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(details), 0)
}
func QuantizationParametersAddQuantizedDimension(builder *flatbuffers.Builder, quantizedDimension int32) {
	builder.PrependInt32Slot(6, quantizedDimension, 0)
}
func QuantizationParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
