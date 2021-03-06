// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DoneFilter int32

const (
	DoneFilter_NONE     DoneFilter = 0
	DoneFilter_DONE     DoneFilter = 1
	DoneFilter_NOT_DONE DoneFilter = 2
)

var DoneFilter_name = map[int32]string{
	0: "NONE",
	1: "DONE",
	2: "NOT_DONE",
}

var DoneFilter_value = map[string]int32{
	"NONE":     0,
	"DONE":     1,
	"NOT_DONE": 2,
}

func (x DoneFilter) String() string {
	return proto.EnumName(DoneFilter_name, int32(x))
}

func (DoneFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

type AddRequest struct {
	Desc                 string   `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type AddReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReply) Reset()         { *m = AddReply{} }
func (m *AddReply) String() string { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()    {}
func (*AddReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *AddReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReply.Unmarshal(m, b)
}
func (m *AddReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReply.Marshal(b, m, deterministic)
}
func (m *AddReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReply.Merge(m, src)
}
func (m *AddReply) XXX_Size() int {
	return xxx_messageInfo_AddReply.Size(m)
}
func (m *AddReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddReply proto.InternalMessageInfo

type ListRequest struct {
	DoneFilter           DoneFilter `protobuf:"varint,1,opt,name=doneFilter,proto3,enum=pb.DoneFilter" json:"doneFilter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetDoneFilter() DoneFilter {
	if m != nil {
		return m.DoneFilter
	}
	return DoneFilter_NONE
}

type ListReply struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReply) Reset()         { *m = ListReply{} }
func (m *ListReply) String() string { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()    {}
func (*ListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *ListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReply.Unmarshal(m, b)
}
func (m *ListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReply.Marshal(b, m, deterministic)
}
func (m *ListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReply.Merge(m, src)
}
func (m *ListReply) XXX_Size() int {
	return xxx_messageInfo_ListReply.Size(m)
}
func (m *ListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListReply proto.InternalMessageInfo

func (m *ListReply) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type DoneRequest struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoneRequest) Reset()         { *m = DoneRequest{} }
func (m *DoneRequest) String() string { return proto.CompactTextString(m) }
func (*DoneRequest) ProtoMessage()    {}
func (*DoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{4}
}

func (m *DoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoneRequest.Unmarshal(m, b)
}
func (m *DoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoneRequest.Marshal(b, m, deterministic)
}
func (m *DoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoneRequest.Merge(m, src)
}
func (m *DoneRequest) XXX_Size() int {
	return xxx_messageInfo_DoneRequest.Size(m)
}
func (m *DoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoneRequest proto.InternalMessageInfo

func (m *DoneRequest) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type DoneReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoneReply) Reset()         { *m = DoneReply{} }
func (m *DoneReply) String() string { return proto.CompactTextString(m) }
func (*DoneReply) ProtoMessage()    {}
func (*DoneReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{5}
}

func (m *DoneReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoneReply.Unmarshal(m, b)
}
func (m *DoneReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoneReply.Marshal(b, m, deterministic)
}
func (m *DoneReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoneReply.Merge(m, src)
}
func (m *DoneReply) XXX_Size() int {
	return xxx_messageInfo_DoneReply.Size(m)
}
func (m *DoneReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DoneReply.DiscardUnknown(m)
}

var xxx_messageInfo_DoneReply proto.InternalMessageInfo

type Task struct {
	Desc                 string   `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	Done                 bool     `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{6}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Task) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func init() {
	proto.RegisterEnum("pb.DoneFilter", DoneFilter_name, DoneFilter_value)
	proto.RegisterType((*AddRequest)(nil), "pb.AddRequest")
	proto.RegisterType((*AddReply)(nil), "pb.AddReply")
	proto.RegisterType((*ListRequest)(nil), "pb.ListRequest")
	proto.RegisterType((*ListReply)(nil), "pb.ListReply")
	proto.RegisterType((*DoneRequest)(nil), "pb.DoneRequest")
	proto.RegisterType((*DoneReply)(nil), "pb.DoneReply")
	proto.RegisterType((*Task)(nil), "pb.Task")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x6b, 0xeb, 0x30,
	0x18, 0xc4, 0xa3, 0xc4, 0x79, 0x38, 0xe7, 0xbc, 0xb4, 0x88, 0x0e, 0x21, 0x43, 0x31, 0x2a, 0x05,
	0xd3, 0x82, 0x06, 0x77, 0xee, 0x10, 0x48, 0x3b, 0x15, 0x07, 0x84, 0xf7, 0x12, 0x57, 0x1a, 0x4c,
	0x8c, 0xe5, 0x46, 0x2a, 0x34, 0x63, 0xff, 0xf3, 0x22, 0x89, 0xd8, 0x1e, 0xba, 0xdd, 0x77, 0xfc,
	0x3e, 0xe9, 0x4e, 0x02, 0xac, 0x96, 0x9a, 0x77, 0x27, 0x6d, 0x35, 0x9d, 0x76, 0x15, 0x4b, 0x81,
	0xad, 0x94, 0x42, 0x7d, 0x7e, 0x29, 0x63, 0x29, 0x45, 0x24, 0x95, 0xf9, 0x58, 0x93, 0x94, 0x64,
	0x0b, 0xe1, 0x35, 0x03, 0x62, 0x4f, 0x74, 0xcd, 0x99, 0x3d, 0x23, 0x79, 0xab, 0x8d, 0xbd, 0xe0,
	0x1c, 0x90, 0xba, 0x55, 0xaf, 0x75, 0x63, 0xd5, 0xc9, 0x2f, 0xad, 0xf2, 0x15, 0xef, 0x2a, 0xbe,
	0xeb, 0x5d, 0x31, 0x22, 0xd8, 0x23, 0x16, 0x61, 0xbd, 0x6b, 0xce, 0xf4, 0x16, 0x73, 0x7b, 0x30,
	0x47, 0xb3, 0x26, 0xe9, 0x2c, 0x4b, 0xf2, 0xd8, 0xed, 0x95, 0x07, 0x73, 0x14, 0xc1, 0x66, 0x77,
	0x48, 0xdc, 0x31, 0x97, 0xbb, 0x6e, 0x30, 0xaf, 0x5b, 0xa9, 0xbe, 0xfd, 0x35, 0x73, 0x11, 0x06,
	0x96, 0x60, 0x11, 0x20, 0x97, 0x8e, 0x23, 0x72, 0x07, 0xfc, 0xd5, 0xc2, 0x7b, 0xba, 0x55, 0xeb,
	0x69, 0x4a, 0xb2, 0x58, 0x78, 0xfd, 0xc0, 0x81, 0x21, 0x28, 0x8d, 0x11, 0x15, 0xfb, 0xe2, 0xe5,
	0x7a, 0xe2, 0xd4, 0xce, 0x29, 0x42, 0x97, 0x88, 0x8b, 0x7d, 0xf9, 0xee, 0xa7, 0x69, 0xfe, 0x43,
	0x10, 0x95, 0x5a, 0x6a, 0x7a, 0x8f, 0xd9, 0x56, 0x4a, 0xea, 0xab, 0x0e, 0xaf, 0xb7, 0x59, 0xf6,
	0xb3, 0x4b, 0x33, 0xa1, 0x19, 0x22, 0x57, 0x97, 0x5e, 0x39, 0x7f, 0xf4, 0x6e, 0x9b, 0xff, 0x83,
	0xd1, 0x93, 0x2e, 0x49, 0x20, 0x47, 0xad, 0x03, 0x39, 0x34, 0x9c, 0x54, 0xff, 0xfc, 0xd7, 0x3d,
	0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x1d, 0x85, 0xd7, 0xc8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*DoneReply, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := c.cc.Invoke(ctx, "/pb.Todo/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/pb.Todo/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*DoneReply, error) {
	out := new(DoneReply)
	err := c.cc.Invoke(ctx, "/pb.Todo/Done", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
type TodoServer interface {
	Add(context.Context, *AddRequest) (*AddReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Done(context.Context, *DoneRequest) (*DoneReply, error)
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Todo/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Todo/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Done(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Todo/Done",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Done(ctx, req.(*DoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Todo_Add_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Todo_List_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _Todo_Done_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
