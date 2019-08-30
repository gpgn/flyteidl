// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/execution.proto

package core // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowExecution_Phase int32

const (
	WorkflowExecution_UNDEFINED  WorkflowExecution_Phase = 0
	WorkflowExecution_QUEUED     WorkflowExecution_Phase = 1
	WorkflowExecution_RUNNING    WorkflowExecution_Phase = 2
	WorkflowExecution_SUCCEEDING WorkflowExecution_Phase = 3
	WorkflowExecution_SUCCEEDED  WorkflowExecution_Phase = 4
	WorkflowExecution_FAILING    WorkflowExecution_Phase = 5
	WorkflowExecution_FAILED     WorkflowExecution_Phase = 6
	WorkflowExecution_ABORTED    WorkflowExecution_Phase = 7
	WorkflowExecution_TIMED_OUT  WorkflowExecution_Phase = 8
)

var WorkflowExecution_Phase_name = map[int32]string{
	0: "UNDEFINED",
	1: "QUEUED",
	2: "RUNNING",
	3: "SUCCEEDING",
	4: "SUCCEEDED",
	5: "FAILING",
	6: "FAILED",
	7: "ABORTED",
	8: "TIMED_OUT",
}
var WorkflowExecution_Phase_value = map[string]int32{
	"UNDEFINED":  0,
	"QUEUED":     1,
	"RUNNING":    2,
	"SUCCEEDING": 3,
	"SUCCEEDED":  4,
	"FAILING":    5,
	"FAILED":     6,
	"ABORTED":    7,
	"TIMED_OUT":  8,
}

func (x WorkflowExecution_Phase) String() string {
	return proto.EnumName(WorkflowExecution_Phase_name, int32(x))
}
func (WorkflowExecution_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_9e5183245240f7ca, []int{0, 0}
}

type NodeExecution_Phase int32

const (
	NodeExecution_UNDEFINED NodeExecution_Phase = 0
	NodeExecution_QUEUED    NodeExecution_Phase = 1
	NodeExecution_RUNNING   NodeExecution_Phase = 2
	NodeExecution_SUCCEEDED NodeExecution_Phase = 3
	NodeExecution_FAILING   NodeExecution_Phase = 4
	NodeExecution_FAILED    NodeExecution_Phase = 5
	NodeExecution_ABORTED   NodeExecution_Phase = 6
	NodeExecution_SKIPPED   NodeExecution_Phase = 7
	NodeExecution_TIMED_OUT NodeExecution_Phase = 8
)

var NodeExecution_Phase_name = map[int32]string{
	0: "UNDEFINED",
	1: "QUEUED",
	2: "RUNNING",
	3: "SUCCEEDED",
	4: "FAILING",
	5: "FAILED",
	6: "ABORTED",
	7: "SKIPPED",
	8: "TIMED_OUT",
}
var NodeExecution_Phase_value = map[string]int32{
	"UNDEFINED": 0,
	"QUEUED":    1,
	"RUNNING":   2,
	"SUCCEEDED": 3,
	"FAILING":   4,
	"FAILED":    5,
	"ABORTED":   6,
	"SKIPPED":   7,
	"TIMED_OUT": 8,
}

func (x NodeExecution_Phase) String() string {
	return proto.EnumName(NodeExecution_Phase_name, int32(x))
}
func (NodeExecution_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_9e5183245240f7ca, []int{1, 0}
}

type TaskExecution_Phase int32

const (
	TaskExecution_UNDEFINED TaskExecution_Phase = 0
	TaskExecution_QUEUED    TaskExecution_Phase = 1
	TaskExecution_RUNNING   TaskExecution_Phase = 2
	TaskExecution_SUCCEEDED TaskExecution_Phase = 3
	TaskExecution_ABORTED   TaskExecution_Phase = 4
	TaskExecution_FAILED    TaskExecution_Phase = 5
)

var TaskExecution_Phase_name = map[int32]string{
	0: "UNDEFINED",
	1: "QUEUED",
	2: "RUNNING",
	3: "SUCCEEDED",
	4: "ABORTED",
	5: "FAILED",
}
var TaskExecution_Phase_value = map[string]int32{
	"UNDEFINED": 0,
	"QUEUED":    1,
	"RUNNING":   2,
	"SUCCEEDED": 3,
	"ABORTED":   4,
	"FAILED":    5,
}

func (x TaskExecution_Phase) String() string {
	return proto.EnumName(TaskExecution_Phase_name, int32(x))
}
func (TaskExecution_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_9e5183245240f7ca, []int{2, 0}
}

type TaskLog_MessageFormat int32

const (
	TaskLog_UNKNOWN TaskLog_MessageFormat = 0
	TaskLog_CSV     TaskLog_MessageFormat = 1
	TaskLog_JSON    TaskLog_MessageFormat = 2
)

var TaskLog_MessageFormat_name = map[int32]string{
	0: "UNKNOWN",
	1: "CSV",
	2: "JSON",
}
var TaskLog_MessageFormat_value = map[string]int32{
	"UNKNOWN": 0,
	"CSV":     1,
	"JSON":    2,
}

func (x TaskLog_MessageFormat) String() string {
	return proto.EnumName(TaskLog_MessageFormat_name, int32(x))
}
func (TaskLog_MessageFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_9e5183245240f7ca, []int{4, 0}
}

// Indicates various phases of Workflow Execution
type WorkflowExecution struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecution) Reset()         { *m = WorkflowExecution{} }
func (m *WorkflowExecution) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecution) ProtoMessage()    {}
func (*WorkflowExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_9e5183245240f7ca, []int{0}
}
func (m *WorkflowExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecution.Unmarshal(m, b)
}
func (m *WorkflowExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecution.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecution.Merge(dst, src)
}
func (m *WorkflowExecution) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecution.Size(m)
}
func (m *WorkflowExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecution.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecution proto.InternalMessageInfo

// Indicates various phases of Node Execution
type NodeExecution struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecution) Reset()         { *m = NodeExecution{} }
func (m *NodeExecution) String() string { return proto.CompactTextString(m) }
func (*NodeExecution) ProtoMessage()    {}
func (*NodeExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_9e5183245240f7ca, []int{1}
}
func (m *NodeExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecution.Unmarshal(m, b)
}
func (m *NodeExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecution.Marshal(b, m, deterministic)
}
func (dst *NodeExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecution.Merge(dst, src)
}
func (m *NodeExecution) XXX_Size() int {
	return xxx_messageInfo_NodeExecution.Size(m)
}
func (m *NodeExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecution.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecution proto.InternalMessageInfo

// Phases that task plugins can go through. Not all phases may be applicable to a specific plugin task,
// but this is the cumulative list that customers may want to know about for their task.
type TaskExecution struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecution) Reset()         { *m = TaskExecution{} }
func (m *TaskExecution) String() string { return proto.CompactTextString(m) }
func (*TaskExecution) ProtoMessage()    {}
func (*TaskExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_9e5183245240f7ca, []int{2}
}
func (m *TaskExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecution.Unmarshal(m, b)
}
func (m *TaskExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecution.Marshal(b, m, deterministic)
}
func (dst *TaskExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecution.Merge(dst, src)
}
func (m *TaskExecution) XXX_Size() int {
	return xxx_messageInfo_TaskExecution.Size(m)
}
func (m *TaskExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecution.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecution proto.InternalMessageInfo

// Represents the error message from the execution.
type ExecutionError struct {
	// Error code indicates a grouping of a type of error.
	// More Info: <Link>
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// Detailed description of the error - including stack trace.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Full error contents accessible via a URI
	ErrorUri             string   `protobuf:"bytes,3,opt,name=error_uri,json=errorUri,proto3" json:"error_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionError) Reset()         { *m = ExecutionError{} }
func (m *ExecutionError) String() string { return proto.CompactTextString(m) }
func (*ExecutionError) ProtoMessage()    {}
func (*ExecutionError) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_9e5183245240f7ca, []int{3}
}
func (m *ExecutionError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionError.Unmarshal(m, b)
}
func (m *ExecutionError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionError.Marshal(b, m, deterministic)
}
func (dst *ExecutionError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionError.Merge(dst, src)
}
func (m *ExecutionError) XXX_Size() int {
	return xxx_messageInfo_ExecutionError.Size(m)
}
func (m *ExecutionError) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionError.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionError proto.InternalMessageInfo

func (m *ExecutionError) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ExecutionError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ExecutionError) GetErrorUri() string {
	if m != nil {
		return m.ErrorUri
	}
	return ""
}

// Log information for the task that is specific to a log sink
// When our log story is flushed out, we may have more metadata here like log link expiry
type TaskLog struct {
	Uri                  string                `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MessageFormat        TaskLog_MessageFormat `protobuf:"varint,3,opt,name=message_format,json=messageFormat,proto3,enum=flyteidl.core.TaskLog_MessageFormat" json:"message_format,omitempty"`
	Ttl                  *duration.Duration    `protobuf:"bytes,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TaskLog) Reset()         { *m = TaskLog{} }
func (m *TaskLog) String() string { return proto.CompactTextString(m) }
func (*TaskLog) ProtoMessage()    {}
func (*TaskLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_9e5183245240f7ca, []int{4}
}
func (m *TaskLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLog.Unmarshal(m, b)
}
func (m *TaskLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLog.Marshal(b, m, deterministic)
}
func (dst *TaskLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLog.Merge(dst, src)
}
func (m *TaskLog) XXX_Size() int {
	return xxx_messageInfo_TaskLog.Size(m)
}
func (m *TaskLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskLog.DiscardUnknown(m)
}

var xxx_messageInfo_TaskLog proto.InternalMessageInfo

func (m *TaskLog) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *TaskLog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskLog) GetMessageFormat() TaskLog_MessageFormat {
	if m != nil {
		return m.MessageFormat
	}
	return TaskLog_UNKNOWN
}

func (m *TaskLog) GetTtl() *duration.Duration {
	if m != nil {
		return m.Ttl
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkflowExecution)(nil), "flyteidl.core.WorkflowExecution")
	proto.RegisterType((*NodeExecution)(nil), "flyteidl.core.NodeExecution")
	proto.RegisterType((*TaskExecution)(nil), "flyteidl.core.TaskExecution")
	proto.RegisterType((*ExecutionError)(nil), "flyteidl.core.ExecutionError")
	proto.RegisterType((*TaskLog)(nil), "flyteidl.core.TaskLog")
	proto.RegisterEnum("flyteidl.core.WorkflowExecution_Phase", WorkflowExecution_Phase_name, WorkflowExecution_Phase_value)
	proto.RegisterEnum("flyteidl.core.NodeExecution_Phase", NodeExecution_Phase_name, NodeExecution_Phase_value)
	proto.RegisterEnum("flyteidl.core.TaskExecution_Phase", TaskExecution_Phase_name, TaskExecution_Phase_value)
	proto.RegisterEnum("flyteidl.core.TaskLog_MessageFormat", TaskLog_MessageFormat_name, TaskLog_MessageFormat_value)
}

func init() {
	proto.RegisterFile("flyteidl/core/execution.proto", fileDescriptor_execution_9e5183245240f7ca)
}

var fileDescriptor_execution_9e5183245240f7ca = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x86, 0xb7, 0x14, 0x28, 0x9c, 0x4d, 0xc9, 0x38, 0x57, 0x55, 0xa3, 0xd9, 0x34, 0x5e, 0x6c,
	0x62, 0x6c, 0x0d, 0x3e, 0xc1, 0x2e, 0x1d, 0x4c, 0x65, 0x77, 0xc0, 0x42, 0x25, 0xd1, 0x0b, 0x52,
	0x60, 0x5a, 0x9a, 0x6d, 0x99, 0xcd, 0xd0, 0x46, 0xf7, 0xce, 0xc4, 0x37, 0xf0, 0xce, 0x57, 0xf3,
	0x69, 0xcc, 0x0c, 0x74, 0xa1, 0x89, 0x37, 0xc6, 0xbb, 0x39, 0xf3, 0xff, 0xe7, 0x3f, 0x5f, 0x9b,
	0x33, 0xf0, 0x22, 0xce, 0x1e, 0x0a, 0x96, 0xae, 0x33, 0x77, 0xc5, 0x05, 0x73, 0xd9, 0x37, 0xb6,
	0x2a, 0x8b, 0x94, 0x6f, 0x9d, 0x7b, 0xc1, 0x0b, 0x8e, 0xcd, 0x4a, 0x76, 0xa4, 0xfc, 0xec, 0x65,
	0xc2, 0x79, 0x92, 0x31, 0x57, 0x89, 0xcb, 0x32, 0x76, 0xd7, 0xa5, 0x88, 0x8e, 0x76, 0xfb, 0x97,
	0x06, 0x4f, 0xe6, 0x5c, 0xdc, 0xc5, 0x19, 0xff, 0x4a, 0xaa, 0x28, 0xfb, 0x87, 0x06, 0xad, 0xc9,
	0x26, 0xda, 0x31, 0x6c, 0x42, 0x37, 0xa4, 0x1e, 0x19, 0xfa, 0x94, 0x78, 0xe8, 0x0c, 0x03, 0xb4,
	0x3f, 0x86, 0x24, 0x24, 0x1e, 0xd2, 0xf0, 0x39, 0x18, 0x41, 0x48, 0xa9, 0x4f, 0xdf, 0xa3, 0x06,
	0xee, 0x01, 0x4c, 0xc3, 0xc1, 0x80, 0x10, 0x4f, 0xd6, 0xba, 0xec, 0x3b, 0xd4, 0xc4, 0x43, 0x4d,
	0xe9, 0x1d, 0x5e, 0xf9, 0x37, 0x52, 0x6b, 0xc9, 0x10, 0x59, 0x10, 0x0f, 0xb5, 0xa5, 0x70, 0x75,
	0x3d, 0x0e, 0x66, 0xc4, 0x43, 0x86, 0x6c, 0x9a, 0xf9, 0xb7, 0xc4, 0x5b, 0x8c, 0xc3, 0x19, 0xea,
	0xd8, 0x3f, 0x35, 0x30, 0x29, 0x5f, 0xb3, 0x23, 0xd7, 0xf7, 0x7f, 0xe6, 0xaa, 0x71, 0xe8, 0xa7,
	0x1c, 0xcd, 0x13, 0x8e, 0xd6, 0x29, 0x87, 0x82, 0x9a, 0x8e, 0xfc, 0xc9, 0xe4, 0x6f, 0x50, 0x1b,
	0x30, 0x67, 0xd1, 0xee, 0xee, 0xc8, 0x34, 0xff, 0x7f, 0xa4, 0x6a, 0x72, 0x0d, 0xc9, 0xfe, 0x02,
	0xbd, 0xc7, 0x29, 0x44, 0x08, 0x2e, 0x30, 0x86, 0xe6, 0x8a, 0xaf, 0x99, 0xa5, 0x5d, 0x68, 0x97,
	0xdd, 0x40, 0x9d, 0xb1, 0x05, 0x46, 0xce, 0x76, 0xbb, 0x28, 0x61, 0x56, 0x43, 0x5d, 0x57, 0x25,
	0x7e, 0x0e, 0x5d, 0x26, 0xdb, 0x16, 0xa5, 0x48, 0x2d, 0x5d, 0x69, 0x1d, 0x75, 0x11, 0x8a, 0xd4,
	0xfe, 0xad, 0x81, 0x21, 0xbf, 0xe3, 0x86, 0x27, 0x18, 0x81, 0x2e, 0x2d, 0xfb, 0x54, 0x79, 0x94,
	0x83, 0xb6, 0x51, 0x5e, 0x25, 0xaa, 0x33, 0x1e, 0x41, 0xef, 0x90, 0xbc, 0x88, 0xb9, 0xc8, 0xa3,
	0x42, 0x65, 0xf6, 0xfa, 0xaf, 0x9c, 0xda, 0xc6, 0x39, 0x87, 0x54, 0xe7, 0x76, 0x6f, 0x1e, 0x2a,
	0x6f, 0x60, 0xe6, 0xa7, 0x25, 0x7e, 0x0d, 0x7a, 0x51, 0x64, 0x56, 0xf3, 0x42, 0xbb, 0x3c, 0xef,
	0x3f, 0x75, 0xf6, 0x4b, 0xea, 0x54, 0x4b, 0xea, 0x78, 0x87, 0x25, 0x0d, 0xa4, 0xcb, 0x76, 0xc1,
	0xac, 0x85, 0xc9, 0x5f, 0x16, 0xd2, 0x11, 0x1d, 0xcf, 0x29, 0x3a, 0xc3, 0x06, 0xe8, 0x83, 0xe9,
	0x27, 0xa4, 0xe1, 0x0e, 0x34, 0x3f, 0x4c, 0xc7, 0x14, 0x35, 0xae, 0xfb, 0x9f, 0xdf, 0x26, 0x69,
	0xb1, 0x29, 0x97, 0xce, 0x8a, 0xe7, 0x6e, 0xf6, 0x10, 0x17, 0xee, 0xe3, 0xa3, 0x49, 0xd8, 0xd6,
	0xbd, 0x5f, 0xbe, 0x49, 0xb8, 0x5b, 0x7b, 0x47, 0xcb, 0xb6, 0x1a, 0xfe, 0xee, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0xa6, 0x1d, 0xa9, 0x5f, 0x03, 0x00, 0x00,
}