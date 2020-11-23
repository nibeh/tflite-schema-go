// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DepthToSpaceOptionsT struct {
	BlockSize int32
}

func (t *DepthToSpaceOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	DepthToSpaceOptionsStart(builder)
	DepthToSpaceOptionsAddBlockSize(builder, t.BlockSize)
	return DepthToSpaceOptionsEnd(builder)
}

func (rcv *DepthToSpaceOptions) UnPackTo(t *DepthToSpaceOptionsT) {
	t.BlockSize = rcv.BlockSize()
}

func (rcv *DepthToSpaceOptions) UnPack() *DepthToSpaceOptionsT {
	if rcv == nil { return nil }
	t := &DepthToSpaceOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type DepthToSpaceOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsDepthToSpaceOptions(buf []byte, offset flatbuffers.UOffsetT) *DepthToSpaceOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DepthToSpaceOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DepthToSpaceOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DepthToSpaceOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DepthToSpaceOptions) BlockSize() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DepthToSpaceOptions) MutateBlockSize(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func DepthToSpaceOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DepthToSpaceOptionsAddBlockSize(builder *flatbuffers.Builder, blockSize int32) {
	builder.PrependInt32Slot(0, blockSize, 0)
}
func DepthToSpaceOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
