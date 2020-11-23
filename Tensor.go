// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TensorT struct {
	Shape []int32
	Type TensorType
	Buffer uint32
	Name string
	Quantization *QuantizationParametersT
	IsVariable bool
	Sparsity *SparsityParametersT
	ShapeSignature []int32
}

func (t *TensorT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	shapeOffset := flatbuffers.UOffsetT(0)
	if t.Shape != nil {
		shapeLength := len(t.Shape)
		TensorStartShapeVector(builder, shapeLength)
		for j := shapeLength - 1; j >= 0; j-- {
			builder.PrependInt32(t.Shape[j])
		}
		shapeOffset = builder.EndVector(shapeLength)
	}
	nameOffset := builder.CreateString(t.Name)
	quantizationOffset := t.Quantization.Pack(builder)
	sparsityOffset := t.Sparsity.Pack(builder)
	shapeSignatureOffset := flatbuffers.UOffsetT(0)
	if t.ShapeSignature != nil {
		shapeSignatureLength := len(t.ShapeSignature)
		TensorStartShapeSignatureVector(builder, shapeSignatureLength)
		for j := shapeSignatureLength - 1; j >= 0; j-- {
			builder.PrependInt32(t.ShapeSignature[j])
		}
		shapeSignatureOffset = builder.EndVector(shapeSignatureLength)
	}
	TensorStart(builder)
	TensorAddShape(builder, shapeOffset)
	TensorAddType(builder, t.Type)
	TensorAddBuffer(builder, t.Buffer)
	TensorAddName(builder, nameOffset)
	TensorAddQuantization(builder, quantizationOffset)
	TensorAddIsVariable(builder, t.IsVariable)
	TensorAddSparsity(builder, sparsityOffset)
	TensorAddShapeSignature(builder, shapeSignatureOffset)
	return TensorEnd(builder)
}

func (rcv *Tensor) UnPackTo(t *TensorT) {
	shapeLength := rcv.ShapeLength()
	t.Shape = make([]int32, shapeLength)
	for j := 0; j < shapeLength; j++ {
		t.Shape[j] = rcv.Shape(j)
	}
	t.Type = rcv.Type()
	t.Buffer = rcv.Buffer()
	t.Name = string(rcv.Name())
	t.Quantization = rcv.Quantization(nil).UnPack()
	t.IsVariable = rcv.IsVariable()
	t.Sparsity = rcv.Sparsity(nil).UnPack()
	shapeSignatureLength := rcv.ShapeSignatureLength()
	t.ShapeSignature = make([]int32, shapeSignatureLength)
	for j := 0; j < shapeSignatureLength; j++ {
		t.ShapeSignature[j] = rcv.ShapeSignature(j)
	}
}

func (rcv *Tensor) UnPack() *TensorT {
	if rcv == nil { return nil }
	t := &TensorT{}
	rcv.UnPackTo(t)
	return t
}

type Tensor struct {
	_tab flatbuffers.Table
}

func GetRootAsTensor(buf []byte, offset flatbuffers.UOffsetT) *Tensor {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Tensor{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Tensor) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Tensor) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Tensor) Shape(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Tensor) ShapeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Tensor) MutateShape(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *Tensor) Type() TensorType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return TensorType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Tensor) MutateType(n TensorType) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func (rcv *Tensor) Buffer() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tensor) MutateBuffer(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *Tensor) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Tensor) Quantization(obj *QuantizationParameters) *QuantizationParameters {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(QuantizationParameters)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Tensor) IsVariable() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Tensor) MutateIsVariable(n bool) bool {
	return rcv._tab.MutateBoolSlot(14, n)
}

func (rcv *Tensor) Sparsity(obj *SparsityParameters) *SparsityParameters {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SparsityParameters)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Tensor) ShapeSignature(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Tensor) ShapeSignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Tensor) MutateShapeSignature(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func TensorStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func TensorAddShape(builder *flatbuffers.Builder, shape flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(shape), 0)
}
func TensorStartShapeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TensorAddType(builder *flatbuffers.Builder, type_ TensorType) {
	builder.PrependInt8Slot(1, int8(type_), 0)
}
func TensorAddBuffer(builder *flatbuffers.Builder, buffer uint32) {
	builder.PrependUint32Slot(2, buffer, 0)
}
func TensorAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(name), 0)
}
func TensorAddQuantization(builder *flatbuffers.Builder, quantization flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(quantization), 0)
}
func TensorAddIsVariable(builder *flatbuffers.Builder, isVariable bool) {
	builder.PrependBoolSlot(5, isVariable, false)
}
func TensorAddSparsity(builder *flatbuffers.Builder, sparsity flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(sparsity), 0)
}
func TensorAddShapeSignature(builder *flatbuffers.Builder, shapeSignature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(shapeSignature), 0)
}
func TensorStartShapeSignatureVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TensorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}