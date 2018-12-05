// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package pb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ListRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsListRequest(buf []byte, offset flatbuffers.UOffsetT) *ListRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ListRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ListRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ListRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ListRequest) DoneFilter() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ListRequest) MutateDoneFilter(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func ListRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ListRequestAddDoneFilter(builder *flatbuffers.Builder, doneFilter int32) {
	builder.PrependInt32Slot(0, doneFilter, 0)
}
func ListRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
