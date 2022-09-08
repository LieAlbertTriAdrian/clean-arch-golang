// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: domain/proto/todo.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EmptyReq message
type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_domain_proto_todo_proto_rawDescGZIP(), []int{0}
}

// GetUsersResponse message
type FetchTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User message
	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *FetchTodosResponse) Reset() {
	*x = FetchTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTodosResponse) ProtoMessage() {}

func (x *FetchTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTodosResponse.ProtoReflect.Descriptor instead.
func (*FetchTodosResponse) Descriptor() ([]byte, []int) {
	return file_domain_proto_todo_proto_rawDescGZIP(), []int{1}
}

func (x *FetchTodosResponse) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

// Todo message
type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_domain_proto_todo_proto_rawDescGZIP(), []int{2}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Todo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Todo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_domain_proto_todo_proto protoreflect.FileDescriptor

var file_domain_proto_todo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22,
	0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x22, 0x36, 0x0a, 0x12, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f,
	0x64, 0x6f, 0x73, 0x22, 0x66, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x43, 0x0a, 0x07, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x70, 0x63, 0x12, 0x38, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54,
	0x6f, 0x64, 0x6f, 0x73, 0x12, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_domain_proto_todo_proto_rawDescOnce sync.Once
	file_domain_proto_todo_proto_rawDescData = file_domain_proto_todo_proto_rawDesc
)

func file_domain_proto_todo_proto_rawDescGZIP() []byte {
	file_domain_proto_todo_proto_rawDescOnce.Do(func() {
		file_domain_proto_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_proto_todo_proto_rawDescData)
	})
	return file_domain_proto_todo_proto_rawDescData
}

var file_domain_proto_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_domain_proto_todo_proto_goTypes = []interface{}{
	(*EmptyReq)(nil),           // 0: todo.EmptyReq
	(*FetchTodosResponse)(nil), // 1: todo.FetchTodosResponse
	(*Todo)(nil),               // 2: todo.Todo
}
var file_domain_proto_todo_proto_depIdxs = []int32{
	2, // 0: todo.FetchTodosResponse.todos:type_name -> todo.Todo
	0, // 1: todo.TodoRpc.FetchTodos:input_type -> todo.EmptyReq
	1, // 2: todo.TodoRpc.FetchTodos:output_type -> todo.FetchTodosResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_domain_proto_todo_proto_init() }
func file_domain_proto_todo_proto_init() {
	if File_domain_proto_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_proto_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_domain_proto_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTodosResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_domain_proto_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_proto_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_domain_proto_todo_proto_goTypes,
		DependencyIndexes: file_domain_proto_todo_proto_depIdxs,
		MessageInfos:      file_domain_proto_todo_proto_msgTypes,
	}.Build()
	File_domain_proto_todo_proto = out.File
	file_domain_proto_todo_proto_rawDesc = nil
	file_domain_proto_todo_proto_goTypes = nil
	file_domain_proto_todo_proto_depIdxs = nil
}