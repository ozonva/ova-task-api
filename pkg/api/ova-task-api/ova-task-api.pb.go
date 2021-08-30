// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ova-task-api/ova-task-api.proto

package ova_task_api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskState int32

const (
	TaskState_TASK_STATE_UNSPECIFIED TaskState = 0
	TaskState_TASK_STATE_ACTIVE      TaskState = 1
	TaskState_TASK_STATE_RESOLVED    TaskState = 2
)

// Enum value maps for TaskState.
var (
	TaskState_name = map[int32]string{
		0: "TASK_STATE_UNSPECIFIED",
		1: "TASK_STATE_ACTIVE",
		2: "TASK_STATE_RESOLVED",
	}
	TaskState_value = map[string]int32{
		"TASK_STATE_UNSPECIFIED": 0,
		"TASK_STATE_ACTIVE":      1,
		"TASK_STATE_RESOLVED":    2,
	}
)

func (x TaskState) Enum() *TaskState {
	p := new(TaskState)
	*p = x
	return p
}

func (x TaskState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskState) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ova_task_api_ova_task_api_proto_enumTypes[0].Descriptor()
}

func (TaskState) Type() protoreflect.EnumType {
	return &file_api_ova_task_api_ova_task_api_proto_enumTypes[0]
}

func (x TaskState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskState.Descriptor instead.
func (TaskState) EnumDescriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{0}
}

type CreateTaskV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateTaskV1Request) Reset() {
	*x = CreateTaskV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskV1Request) ProtoMessage() {}

func (x *CreateTaskV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskV1Request.ProtoReflect.Descriptor instead.
func (*CreateTaskV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateTaskV1Request) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type TaskV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskId      uint64                 `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	State       TaskState              `protobuf:"varint,5,opt,name=state,proto3,enum=ova.task.api.TaskState" json:"state,omitempty"`
}

func (x *TaskV1) Reset() {
	*x = TaskV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskV1) ProtoMessage() {}

func (x *TaskV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskV1.ProtoReflect.Descriptor instead.
func (*TaskV1) Descriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{1}
}

func (x *TaskV1) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TaskV1) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskV1) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskV1) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

func (x *TaskV1) GetState() TaskState {
	if x != nil {
		return x.State
	}
	return TaskState_TASK_STATE_UNSPECIFIED
}

type CreateTaskV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *TaskV1 `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *CreateTaskV1Response) Reset() {
	*x = CreateTaskV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskV1Response) ProtoMessage() {}

func (x *CreateTaskV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskV1Response.ProtoReflect.Descriptor instead.
func (*CreateTaskV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskV1Response) GetTask() *TaskV1 {
	if x != nil {
		return x.Task
	}
	return nil
}

type DescribeTaskV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *DescribeTaskV1Request) Reset() {
	*x = DescribeTaskV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTaskV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTaskV1Request) ProtoMessage() {}

func (x *DescribeTaskV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTaskV1Request.ProtoReflect.Descriptor instead.
func (*DescribeTaskV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeTaskV1Request) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type DescribeTaskV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *TaskV1 `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *DescribeTaskV1Response) Reset() {
	*x = DescribeTaskV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTaskV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTaskV1Response) ProtoMessage() {}

func (x *DescribeTaskV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTaskV1Response.ProtoReflect.Descriptor instead.
func (*DescribeTaskV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeTaskV1Response) GetTask() *TaskV1 {
	if x != nil {
		return x.Task
	}
	return nil
}

type ListTasksV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskState TaskState `protobuf:"varint,1,opt,name=task_state,json=taskState,proto3,enum=ova.task.api.TaskState" json:"task_state,omitempty"`
	Limit     uint64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    uint64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListTasksV1Request) Reset() {
	*x = ListTasksV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTasksV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksV1Request) ProtoMessage() {}

func (x *ListTasksV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksV1Request.ProtoReflect.Descriptor instead.
func (*ListTasksV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListTasksV1Request) GetTaskState() TaskState {
	if x != nil {
		return x.TaskState
	}
	return TaskState_TASK_STATE_UNSPECIFIED
}

func (x *ListTasksV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTasksV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListTasksV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*TaskV1 `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *ListTasksV1Response) Reset() {
	*x = ListTasksV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTasksV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksV1Response) ProtoMessage() {}

func (x *ListTasksV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksV1Response.ProtoReflect.Descriptor instead.
func (*ListTasksV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListTasksV1Response) GetTasks() []*TaskV1 {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type RemoveTaskV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *RemoveTaskV1Request) Reset() {
	*x = RemoveTaskV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTaskV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTaskV1Request) ProtoMessage() {}

func (x *RemoveTaskV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTaskV1Request.ProtoReflect.Descriptor instead.
func (*RemoveTaskV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveTaskV1Request) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type RemoveTaskV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *RemoveTaskV1Response) Reset() {
	*x = RemoveTaskV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTaskV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTaskV1Response) ProtoMessage() {}

func (x *RemoveTaskV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_task_api_ova_task_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTaskV1Response.ProtoReflect.Descriptor instead.
func (*RemoveTaskV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_task_api_ova_task_api_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveTaskV1Response) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

var File_api_ova_task_api_ova_task_api_proto protoreflect.FileDescriptor

var file_api_ova_task_api_ova_task_api_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x50, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xca, 0x01, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x40,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x22, 0x30, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x22, 0x42, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x76, 0x61,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31,
	0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x7a, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x41, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x2e, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x2a, 0x57, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x15, 0x0a, 0x11, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x10, 0x02, 0x32,
	0xcf, 0x03, 0x0a, 0x0a, 0x4f, 0x76, 0x61, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x70, 0x69, 0x12, 0x72,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x12, 0x21,
	0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x78, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x76, 0x61, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x56, 0x31, 0x12, 0x20, 0x2e, 0x6f, 0x76,
	0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x67, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x56, 0x31, 0x12, 0x21, 0x2e, 0x6f, 0x76, 0x61, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x3a, 0x01,
	0x2a, 0x42, 0x33, 0x5a, 0x31, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2f, 0x6f, 0x76, 0x61, 0x2d,
	0x74, 0x61, 0x73, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x76, 0x61,
	0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x76, 0x61, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ova_task_api_ova_task_api_proto_rawDescOnce sync.Once
	file_api_ova_task_api_ova_task_api_proto_rawDescData = file_api_ova_task_api_ova_task_api_proto_rawDesc
)

func file_api_ova_task_api_ova_task_api_proto_rawDescGZIP() []byte {
	file_api_ova_task_api_ova_task_api_proto_rawDescOnce.Do(func() {
		file_api_ova_task_api_ova_task_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ova_task_api_ova_task_api_proto_rawDescData)
	})
	return file_api_ova_task_api_ova_task_api_proto_rawDescData
}

var file_api_ova_task_api_ova_task_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_ova_task_api_ova_task_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_ova_task_api_ova_task_api_proto_goTypes = []interface{}{
	(TaskState)(0),                 // 0: ova.task.api.TaskState
	(*CreateTaskV1Request)(nil),    // 1: ova.task.api.CreateTaskV1Request
	(*TaskV1)(nil),                 // 2: ova.task.api.TaskV1
	(*CreateTaskV1Response)(nil),   // 3: ova.task.api.CreateTaskV1Response
	(*DescribeTaskV1Request)(nil),  // 4: ova.task.api.DescribeTaskV1Request
	(*DescribeTaskV1Response)(nil), // 5: ova.task.api.DescribeTaskV1Response
	(*ListTasksV1Request)(nil),     // 6: ova.task.api.ListTasksV1Request
	(*ListTasksV1Response)(nil),    // 7: ova.task.api.ListTasksV1Response
	(*RemoveTaskV1Request)(nil),    // 8: ova.task.api.RemoveTaskV1Request
	(*RemoveTaskV1Response)(nil),   // 9: ova.task.api.RemoveTaskV1Response
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 11: google.protobuf.Empty
}
var file_api_ova_task_api_ova_task_api_proto_depIdxs = []int32{
	10, // 0: ova.task.api.TaskV1.date_created:type_name -> google.protobuf.Timestamp
	0,  // 1: ova.task.api.TaskV1.state:type_name -> ova.task.api.TaskState
	2,  // 2: ova.task.api.CreateTaskV1Response.task:type_name -> ova.task.api.TaskV1
	2,  // 3: ova.task.api.DescribeTaskV1Response.task:type_name -> ova.task.api.TaskV1
	0,  // 4: ova.task.api.ListTasksV1Request.task_state:type_name -> ova.task.api.TaskState
	2,  // 5: ova.task.api.ListTasksV1Response.tasks:type_name -> ova.task.api.TaskV1
	1,  // 6: ova.task.api.OvaTaskApi.CreateTaskV1:input_type -> ova.task.api.CreateTaskV1Request
	4,  // 7: ova.task.api.OvaTaskApi.DescribeTaskV1:input_type -> ova.task.api.DescribeTaskV1Request
	6,  // 8: ova.task.api.OvaTaskApi.ListTasksV1:input_type -> ova.task.api.ListTasksV1Request
	8,  // 9: ova.task.api.OvaTaskApi.RemoveTasksV1:input_type -> ova.task.api.RemoveTaskV1Request
	3,  // 10: ova.task.api.OvaTaskApi.CreateTaskV1:output_type -> ova.task.api.CreateTaskV1Response
	5,  // 11: ova.task.api.OvaTaskApi.DescribeTaskV1:output_type -> ova.task.api.DescribeTaskV1Response
	7,  // 12: ova.task.api.OvaTaskApi.ListTasksV1:output_type -> ova.task.api.ListTasksV1Response
	11, // 13: ova.task.api.OvaTaskApi.RemoveTasksV1:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_ova_task_api_ova_task_api_proto_init() }
func file_api_ova_task_api_ova_task_api_proto_init() {
	if File_api_ova_task_api_ova_task_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ova_task_api_ova_task_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskV1Request); i {
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
		file_api_ova_task_api_ova_task_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskV1); i {
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
		file_api_ova_task_api_ova_task_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskV1Response); i {
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
		file_api_ova_task_api_ova_task_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTaskV1Request); i {
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
		file_api_ova_task_api_ova_task_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTaskV1Response); i {
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
		file_api_ova_task_api_ova_task_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTasksV1Request); i {
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
		file_api_ova_task_api_ova_task_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTasksV1Response); i {
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
		file_api_ova_task_api_ova_task_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTaskV1Request); i {
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
		file_api_ova_task_api_ova_task_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTaskV1Response); i {
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
			RawDescriptor: file_api_ova_task_api_ova_task_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ova_task_api_ova_task_api_proto_goTypes,
		DependencyIndexes: file_api_ova_task_api_ova_task_api_proto_depIdxs,
		EnumInfos:         file_api_ova_task_api_ova_task_api_proto_enumTypes,
		MessageInfos:      file_api_ova_task_api_ova_task_api_proto_msgTypes,
	}.Build()
	File_api_ova_task_api_ova_task_api_proto = out.File
	file_api_ova_task_api_ova_task_api_proto_rawDesc = nil
	file_api_ova_task_api_ova_task_api_proto_goTypes = nil
	file_api_ova_task_api_ova_task_api_proto_depIdxs = nil
}
