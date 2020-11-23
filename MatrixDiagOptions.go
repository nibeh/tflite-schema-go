// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MatrixDiagOptionsT struct {
}

func (t *MatrixDiagOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	MatrixDiagOptionsStart(builder)
	return MatrixDiagOptionsEnd(builder)
}

func (rcv *MatrixDiagOptions) UnPackTo(t *MatrixDiagOptionsT) {
}

func (rcv *MatrixDiagOptions) UnPack() *MatrixDiagOptionsT {
	if rcv == nil { return nil }
	t := &MatrixDiagOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type MatrixDiagOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsMatrixDiagOptions(buf []byte, offset flatbuffers.UOffsetT) *MatrixDiagOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MatrixDiagOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MatrixDiagOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MatrixDiagOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func MatrixDiagOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func MatrixDiagOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
