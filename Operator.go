// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OperatorT struct {
	OpcodeIndex uint32
	Inputs []int32
	Outputs []int32
	BuiltinOptions *BuiltinOptionsT
	CustomOptions []byte
	CustomOptionsFormat CustomOptionsFormat
	MutatingVariableInputs []bool
	Intermediates []int32
}

func (t *OperatorT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	inputsOffset := flatbuffers.UOffsetT(0)
	if t.Inputs != nil {
		inputsLength := len(t.Inputs)
		OperatorStartInputsVector(builder, inputsLength)
		for j := inputsLength - 1; j >= 0; j-- {
			builder.PrependInt32(t.Inputs[j])
		}
		inputsOffset = builder.EndVector(inputsLength)
	}
	outputsOffset := flatbuffers.UOffsetT(0)
	if t.Outputs != nil {
		outputsLength := len(t.Outputs)
		OperatorStartOutputsVector(builder, outputsLength)
		for j := outputsLength - 1; j >= 0; j-- {
			builder.PrependInt32(t.Outputs[j])
		}
		outputsOffset = builder.EndVector(outputsLength)
	}
	builtinOptionsOffset := t.BuiltinOptions.Pack(builder)
	
	customOptionsOffset := flatbuffers.UOffsetT(0)
	if t.CustomOptions != nil {
		customOptionsOffset = builder.CreateByteString(t.CustomOptions)
	}
	mutatingVariableInputsOffset := flatbuffers.UOffsetT(0)
	if t.MutatingVariableInputs != nil {
		mutatingVariableInputsLength := len(t.MutatingVariableInputs)
		OperatorStartMutatingVariableInputsVector(builder, mutatingVariableInputsLength)
		for j := mutatingVariableInputsLength - 1; j >= 0; j-- {
			builder.PrependBool(t.MutatingVariableInputs[j])
		}
		mutatingVariableInputsOffset = builder.EndVector(mutatingVariableInputsLength)
	}
	intermediatesOffset := flatbuffers.UOffsetT(0)
	if t.Intermediates != nil {
		intermediatesLength := len(t.Intermediates)
		OperatorStartIntermediatesVector(builder, intermediatesLength)
		for j := intermediatesLength - 1; j >= 0; j-- {
			builder.PrependInt32(t.Intermediates[j])
		}
		intermediatesOffset = builder.EndVector(intermediatesLength)
	}
	OperatorStart(builder)
	OperatorAddOpcodeIndex(builder, t.OpcodeIndex)
	OperatorAddInputs(builder, inputsOffset)
	OperatorAddOutputs(builder, outputsOffset)
	if t.BuiltinOptions != nil {
		OperatorAddBuiltinOptionsType(builder, t.BuiltinOptions.Type)
	}
	OperatorAddBuiltinOptions(builder, builtinOptionsOffset)
	OperatorAddCustomOptions(builder, customOptionsOffset)
	OperatorAddCustomOptionsFormat(builder, t.CustomOptionsFormat)
	OperatorAddMutatingVariableInputs(builder, mutatingVariableInputsOffset)
	OperatorAddIntermediates(builder, intermediatesOffset)
	return OperatorEnd(builder)
}

func (rcv *Operator) UnPackTo(t *OperatorT) {
	t.OpcodeIndex = rcv.OpcodeIndex()
	inputsLength := rcv.InputsLength()
	t.Inputs = make([]int32, inputsLength)
	for j := 0; j < inputsLength; j++ {
		t.Inputs[j] = rcv.Inputs(j)
	}
	outputsLength := rcv.OutputsLength()
	t.Outputs = make([]int32, outputsLength)
	for j := 0; j < outputsLength; j++ {
		t.Outputs[j] = rcv.Outputs(j)
	}
	builtinOptionsTable := flatbuffers.Table{}
	if rcv.BuiltinOptions(&builtinOptionsTable) {
		t.BuiltinOptions = rcv.BuiltinOptionsType().UnPack(builtinOptionsTable)
	}
	t.CustomOptions = rcv.CustomOptionsBytes()
	t.CustomOptionsFormat = rcv.CustomOptionsFormat()
	mutatingVariableInputsLength := rcv.MutatingVariableInputsLength()
	t.MutatingVariableInputs = make([]bool, mutatingVariableInputsLength)
	for j := 0; j < mutatingVariableInputsLength; j++ {
		t.MutatingVariableInputs[j] = rcv.MutatingVariableInputs(j)
	}
	intermediatesLength := rcv.IntermediatesLength()
	t.Intermediates = make([]int32, intermediatesLength)
	for j := 0; j < intermediatesLength; j++ {
		t.Intermediates[j] = rcv.Intermediates(j)
	}
}

func (rcv *Operator) UnPack() *OperatorT {
	if rcv == nil { return nil }
	t := &OperatorT{}
	rcv.UnPackTo(t)
	return t
}

type Operator struct {
	_tab flatbuffers.Table
}

func GetRootAsOperator(buf []byte, offset flatbuffers.UOffsetT) *Operator {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Operator{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Operator) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Operator) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Operator) OpcodeIndex() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Operator) MutateOpcodeIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Operator) Inputs(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Operator) InputsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Operator) MutateInputs(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *Operator) Outputs(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Operator) OutputsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Operator) MutateOutputs(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *Operator) BuiltinOptionsType() BuiltinOptions {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return BuiltinOptions(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Operator) MutateBuiltinOptionsType(n BuiltinOptions) bool {
	return rcv._tab.MutateByteSlot(10, byte(n))
}

func (rcv *Operator) BuiltinOptions(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *Operator) CustomOptions(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Operator) CustomOptionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Operator) CustomOptionsBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Operator) MutateCustomOptions(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Operator) CustomOptionsFormat() CustomOptionsFormat {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return CustomOptionsFormat(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Operator) MutateCustomOptionsFormat(n CustomOptionsFormat) bool {
	return rcv._tab.MutateInt8Slot(16, int8(n))
}

func (rcv *Operator) MutatingVariableInputs(j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetBool(a + flatbuffers.UOffsetT(j*1))
	}
	return false
}

func (rcv *Operator) MutatingVariableInputsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Operator) MutateMutatingVariableInputs(j int, n bool) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateBool(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Operator) Intermediates(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Operator) IntermediatesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Operator) MutateIntermediates(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func OperatorStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func OperatorAddOpcodeIndex(builder *flatbuffers.Builder, opcodeIndex uint32) {
	builder.PrependUint32Slot(0, opcodeIndex, 0)
}
func OperatorAddInputs(builder *flatbuffers.Builder, inputs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(inputs), 0)
}
func OperatorStartInputsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func OperatorAddOutputs(builder *flatbuffers.Builder, outputs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(outputs), 0)
}
func OperatorStartOutputsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func OperatorAddBuiltinOptionsType(builder *flatbuffers.Builder, builtinOptionsType BuiltinOptions) {
	builder.PrependByteSlot(3, byte(builtinOptionsType), 0)
}
func OperatorAddBuiltinOptions(builder *flatbuffers.Builder, builtinOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(builtinOptions), 0)
}
func OperatorAddCustomOptions(builder *flatbuffers.Builder, customOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(customOptions), 0)
}
func OperatorStartCustomOptionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func OperatorAddCustomOptionsFormat(builder *flatbuffers.Builder, customOptionsFormat CustomOptionsFormat) {
	builder.PrependInt8Slot(6, int8(customOptionsFormat), 0)
}
func OperatorAddMutatingVariableInputs(builder *flatbuffers.Builder, mutatingVariableInputs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(mutatingVariableInputs), 0)
}
func OperatorStartMutatingVariableInputsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func OperatorAddIntermediates(builder *flatbuffers.Builder, intermediates flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(intermediates), 0)
}
func OperatorStartIntermediatesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func OperatorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}