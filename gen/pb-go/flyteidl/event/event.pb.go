// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/event/event.proto

package event // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/event"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowExecutionEvent struct {
	// Workflow execution id
	ExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                       `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.WorkflowExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.WorkflowExecution_Phase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the workflow.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*WorkflowExecutionEvent_OutputUri
	//	*WorkflowExecutionEvent_Error
	OutputResult         isWorkflowExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *WorkflowExecutionEvent) Reset()         { *m = WorkflowExecutionEvent{} }
func (m *WorkflowExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEvent) ProtoMessage()    {}
func (*WorkflowExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_032935db9c57c03d, []int{0}
}
func (m *WorkflowExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEvent.Unmarshal(m, b)
}
func (m *WorkflowExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEvent.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEvent.Merge(dst, src)
}
func (m *WorkflowExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEvent.Size(m)
}
func (m *WorkflowExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEvent proto.InternalMessageInfo

func (m *WorkflowExecutionEvent) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetPhase() core.WorkflowExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.WorkflowExecution_UNDEFINED
}

func (m *WorkflowExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

type isWorkflowExecutionEvent_OutputResult interface {
	isWorkflowExecutionEvent_OutputResult()
}

type WorkflowExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,5,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type WorkflowExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,6,opt,name=error,proto3,oneof"`
}

func (*WorkflowExecutionEvent_OutputUri) isWorkflowExecutionEvent_OutputResult() {}

func (*WorkflowExecutionEvent_Error) isWorkflowExecutionEvent_OutputResult() {}

func (m *WorkflowExecutionEvent) GetOutputResult() isWorkflowExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WorkflowExecutionEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WorkflowExecutionEvent_OneofMarshaler, _WorkflowExecutionEvent_OneofUnmarshaler, _WorkflowExecutionEvent_OneofSizer, []interface{}{
		(*WorkflowExecutionEvent_OutputUri)(nil),
		(*WorkflowExecutionEvent_Error)(nil),
	}
}

func _WorkflowExecutionEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WorkflowExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *WorkflowExecutionEvent_OutputUri:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutputUri)
	case *WorkflowExecutionEvent_Error:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("WorkflowExecutionEvent.OutputResult has unexpected type %T", x)
	}
	return nil
}

func _WorkflowExecutionEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WorkflowExecutionEvent)
	switch tag {
	case 5: // output_result.output_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OutputResult = &WorkflowExecutionEvent_OutputUri{x}
		return true, err
	case 6: // output_result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ExecutionError)
		err := b.DecodeMessage(msg)
		m.OutputResult = &WorkflowExecutionEvent_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _WorkflowExecutionEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WorkflowExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *WorkflowExecutionEvent_OutputUri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.OutputUri)))
		n += len(x.OutputUri)
	case *WorkflowExecutionEvent_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NodeExecutionEvent struct {
	// Unique identifier for this node execution
	Id *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                   `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.NodeExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.NodeExecution_Phase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the node.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	InputUri   string               `protobuf:"bytes,5,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*NodeExecutionEvent_OutputUri
	//	*NodeExecutionEvent_Error
	OutputResult isNodeExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	// Additional metadata to do with this event's node target based
	// on the node type
	//
	// Types that are valid to be assigned to TargetMetadata:
	//	*NodeExecutionEvent_WorkflowNodeMetadata
	TargetMetadata isNodeExecutionEvent_TargetMetadata `protobuf_oneof:"target_metadata"`
	// Specifies which task (if any) launched this node.
	ParentTaskMetadata   *ParentTaskExecutionMetadata `protobuf:"bytes,9,opt,name=parent_task_metadata,json=parentTaskMetadata,proto3" json:"parent_task_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *NodeExecutionEvent) Reset()         { *m = NodeExecutionEvent{} }
func (m *NodeExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEvent) ProtoMessage()    {}
func (*NodeExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_032935db9c57c03d, []int{1}
}
func (m *NodeExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEvent.Unmarshal(m, b)
}
func (m *NodeExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEvent.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEvent.Merge(dst, src)
}
func (m *NodeExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEvent.Size(m)
}
func (m *NodeExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEvent proto.InternalMessageInfo

func (m *NodeExecutionEvent) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NodeExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *NodeExecutionEvent) GetPhase() core.NodeExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.NodeExecution_UNDEFINED
}

func (m *NodeExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *NodeExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isNodeExecutionEvent_OutputResult interface {
	isNodeExecutionEvent_OutputResult()
}

type NodeExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,6,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type NodeExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,7,opt,name=error,proto3,oneof"`
}

func (*NodeExecutionEvent_OutputUri) isNodeExecutionEvent_OutputResult() {}

func (*NodeExecutionEvent_Error) isNodeExecutionEvent_OutputResult() {}

func (m *NodeExecutionEvent) GetOutputResult() isNodeExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *NodeExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *NodeExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

type isNodeExecutionEvent_TargetMetadata interface {
	isNodeExecutionEvent_TargetMetadata()
}

type NodeExecutionEvent_WorkflowNodeMetadata struct {
	WorkflowNodeMetadata *WorkflowNodeMetadata `protobuf:"bytes,8,opt,name=workflow_node_metadata,json=workflowNodeMetadata,proto3,oneof"`
}

func (*NodeExecutionEvent_WorkflowNodeMetadata) isNodeExecutionEvent_TargetMetadata() {}

func (m *NodeExecutionEvent) GetTargetMetadata() isNodeExecutionEvent_TargetMetadata {
	if m != nil {
		return m.TargetMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetWorkflowNodeMetadata() *WorkflowNodeMetadata {
	if x, ok := m.GetTargetMetadata().(*NodeExecutionEvent_WorkflowNodeMetadata); ok {
		return x.WorkflowNodeMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetParentTaskMetadata() *ParentTaskExecutionMetadata {
	if m != nil {
		return m.ParentTaskMetadata
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NodeExecutionEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NodeExecutionEvent_OneofMarshaler, _NodeExecutionEvent_OneofUnmarshaler, _NodeExecutionEvent_OneofSizer, []interface{}{
		(*NodeExecutionEvent_OutputUri)(nil),
		(*NodeExecutionEvent_Error)(nil),
		(*NodeExecutionEvent_WorkflowNodeMetadata)(nil),
	}
}

func _NodeExecutionEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NodeExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *NodeExecutionEvent_OutputUri:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutputUri)
	case *NodeExecutionEvent_Error:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NodeExecutionEvent.OutputResult has unexpected type %T", x)
	}
	// target_metadata
	switch x := m.TargetMetadata.(type) {
	case *NodeExecutionEvent_WorkflowNodeMetadata:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WorkflowNodeMetadata); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NodeExecutionEvent.TargetMetadata has unexpected type %T", x)
	}
	return nil
}

func _NodeExecutionEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NodeExecutionEvent)
	switch tag {
	case 6: // output_result.output_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OutputResult = &NodeExecutionEvent_OutputUri{x}
		return true, err
	case 7: // output_result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ExecutionError)
		err := b.DecodeMessage(msg)
		m.OutputResult = &NodeExecutionEvent_Error{msg}
		return true, err
	case 8: // target_metadata.workflow_node_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WorkflowNodeMetadata)
		err := b.DecodeMessage(msg)
		m.TargetMetadata = &NodeExecutionEvent_WorkflowNodeMetadata{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NodeExecutionEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NodeExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *NodeExecutionEvent_OutputUri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.OutputUri)))
		n += len(x.OutputUri)
	case *NodeExecutionEvent_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// target_metadata
	switch x := m.TargetMetadata.(type) {
	case *NodeExecutionEvent_WorkflowNodeMetadata:
		s := proto.Size(x.WorkflowNodeMetadata)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// For Workflow Nodes we need to send information about the workflow that's launched
type WorkflowNodeMetadata struct {
	ExecutionId          *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *WorkflowNodeMetadata) Reset()         { *m = WorkflowNodeMetadata{} }
func (m *WorkflowNodeMetadata) String() string { return proto.CompactTextString(m) }
func (*WorkflowNodeMetadata) ProtoMessage()    {}
func (*WorkflowNodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_032935db9c57c03d, []int{2}
}
func (m *WorkflowNodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowNodeMetadata.Unmarshal(m, b)
}
func (m *WorkflowNodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowNodeMetadata.Marshal(b, m, deterministic)
}
func (dst *WorkflowNodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowNodeMetadata.Merge(dst, src)
}
func (m *WorkflowNodeMetadata) XXX_Size() int {
	return xxx_messageInfo_WorkflowNodeMetadata.Size(m)
}
func (m *WorkflowNodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowNodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowNodeMetadata proto.InternalMessageInfo

func (m *WorkflowNodeMetadata) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

type ParentTaskExecutionMetadata struct {
	Id                   *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ParentTaskExecutionMetadata) Reset()         { *m = ParentTaskExecutionMetadata{} }
func (m *ParentTaskExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*ParentTaskExecutionMetadata) ProtoMessage()    {}
func (*ParentTaskExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_032935db9c57c03d, []int{3}
}
func (m *ParentTaskExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Unmarshal(m, b)
}
func (m *ParentTaskExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Marshal(b, m, deterministic)
}
func (dst *ParentTaskExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentTaskExecutionMetadata.Merge(dst, src)
}
func (m *ParentTaskExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Size(m)
}
func (m *ParentTaskExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentTaskExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ParentTaskExecutionMetadata proto.InternalMessageInfo

func (m *ParentTaskExecutionMetadata) GetId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Plugin specific execution event information. For tasks like Python, Hive, Spark, DynamicJob.
type TaskExecutionEvent struct {
	// ID of the task. In combination with the retryAttempt this will indicate
	// the task execution uniquely for a given parent node execution.
	TaskId *core.Identifier `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// A task execution is always kicked off by a node execution, the event consumer
	// will use the parent_id to relate the task to it's parent node execution
	ParentNodeExecutionId *core.NodeExecutionIdentifier `protobuf:"bytes,2,opt,name=parent_node_execution_id,json=parentNodeExecutionId,proto3" json:"parent_node_execution_id,omitempty"`
	// retry attempt number for this task, ie., 2 for the second attempt
	RetryAttempt uint32 `protobuf:"varint,3,opt,name=retry_attempt,json=retryAttempt,proto3" json:"retry_attempt,omitempty"`
	// Phase associated with the event
	Phase core.TaskExecution_Phase `protobuf:"varint,4,opt,name=phase,proto3,enum=flyteidl.core.TaskExecution_Phase" json:"phase,omitempty"`
	// id of the process that sent this event, mainly for trace debugging
	ProducerId string `protobuf:"bytes,5,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	// log information for the task execution
	Logs []*core.TaskLog `protobuf:"bytes,6,rep,name=logs,proto3" json:"logs,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the task.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// URI of the input file, it encodes all the information
	// including Cloud source provider. ie., s3://...
	InputUri string `protobuf:"bytes,8,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*TaskExecutionEvent_OutputUri
	//	*TaskExecutionEvent_Error
	OutputResult isTaskExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	// Custom data that the task plugin sends back. This is extensible to allow various plugins in the system.
	CustomInfo *_struct.Struct `protobuf:"bytes,11,opt,name=custom_info,json=customInfo,proto3" json:"custom_info,omitempty"`
	// Some phases, like RUNNING, can send multiple events with changed metadata (new logs, additional custom_info, etc)
	// that should be recorded regardless of the lack of phase change.
	// The version field should be incremented when metadata changes across the duration of an individual phase.
	PhaseVersion         uint32   `protobuf:"varint,12,opt,name=phase_version,json=phaseVersion,proto3" json:"phase_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionEvent) Reset()         { *m = TaskExecutionEvent{} }
func (m *TaskExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEvent) ProtoMessage()    {}
func (*TaskExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_032935db9c57c03d, []int{4}
}
func (m *TaskExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEvent.Unmarshal(m, b)
}
func (m *TaskExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEvent.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEvent.Merge(dst, src)
}
func (m *TaskExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEvent.Size(m)
}
func (m *TaskExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEvent proto.InternalMessageInfo

func (m *TaskExecutionEvent) GetTaskId() *core.Identifier {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *TaskExecutionEvent) GetParentNodeExecutionId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.ParentNodeExecutionId
	}
	return nil
}

func (m *TaskExecutionEvent) GetRetryAttempt() uint32 {
	if m != nil {
		return m.RetryAttempt
	}
	return 0
}

func (m *TaskExecutionEvent) GetPhase() core.TaskExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.TaskExecution_UNDEFINED
}

func (m *TaskExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *TaskExecutionEvent) GetLogs() []*core.TaskLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *TaskExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *TaskExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isTaskExecutionEvent_OutputResult interface {
	isTaskExecutionEvent_OutputResult()
}

type TaskExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,9,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type TaskExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,10,opt,name=error,proto3,oneof"`
}

func (*TaskExecutionEvent_OutputUri) isTaskExecutionEvent_OutputResult() {}

func (*TaskExecutionEvent_Error) isTaskExecutionEvent_OutputResult() {}

func (m *TaskExecutionEvent) GetOutputResult() isTaskExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *TaskExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *TaskExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

func (m *TaskExecutionEvent) GetCustomInfo() *_struct.Struct {
	if m != nil {
		return m.CustomInfo
	}
	return nil
}

func (m *TaskExecutionEvent) GetPhaseVersion() uint32 {
	if m != nil {
		return m.PhaseVersion
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TaskExecutionEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TaskExecutionEvent_OneofMarshaler, _TaskExecutionEvent_OneofUnmarshaler, _TaskExecutionEvent_OneofSizer, []interface{}{
		(*TaskExecutionEvent_OutputUri)(nil),
		(*TaskExecutionEvent_Error)(nil),
	}
}

func _TaskExecutionEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TaskExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *TaskExecutionEvent_OutputUri:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutputUri)
	case *TaskExecutionEvent_Error:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TaskExecutionEvent.OutputResult has unexpected type %T", x)
	}
	return nil
}

func _TaskExecutionEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TaskExecutionEvent)
	switch tag {
	case 9: // output_result.output_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OutputResult = &TaskExecutionEvent_OutputUri{x}
		return true, err
	case 10: // output_result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ExecutionError)
		err := b.DecodeMessage(msg)
		m.OutputResult = &TaskExecutionEvent_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TaskExecutionEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TaskExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *TaskExecutionEvent_OutputUri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.OutputUri)))
		n += len(x.OutputUri)
	case *TaskExecutionEvent_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*WorkflowExecutionEvent)(nil), "flyteidl.event.WorkflowExecutionEvent")
	proto.RegisterType((*NodeExecutionEvent)(nil), "flyteidl.event.NodeExecutionEvent")
	proto.RegisterType((*WorkflowNodeMetadata)(nil), "flyteidl.event.WorkflowNodeMetadata")
	proto.RegisterType((*ParentTaskExecutionMetadata)(nil), "flyteidl.event.ParentTaskExecutionMetadata")
	proto.RegisterType((*TaskExecutionEvent)(nil), "flyteidl.event.TaskExecutionEvent")
}

func init() { proto.RegisterFile("flyteidl/event/event.proto", fileDescriptor_event_032935db9c57c03d) }

var fileDescriptor_event_032935db9c57c03d = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xef, 0x6b, 0xd3, 0x40,
	0x18, 0x5e, 0xbb, 0xb6, 0x5b, 0xdf, 0xec, 0x07, 0x1e, 0x73, 0xc6, 0xce, 0xb9, 0x52, 0x45, 0xca,
	0xc4, 0x04, 0x3b, 0x94, 0x81, 0x7e, 0xd9, 0x60, 0xb0, 0x82, 0x93, 0x11, 0x37, 0x05, 0x51, 0x42,
	0x9a, 0x5c, 0xb2, 0x63, 0x69, 0x2e, 0x5c, 0xde, 0x6c, 0xee, 0x5f, 0xf3, 0x6f, 0xf2, 0xbb, 0x5f,
	0x25, 0x97, 0x1f, 0x5d, 0xd2, 0x52, 0x9d, 0xec, 0x4b, 0xa1, 0xcf, 0x3d, 0xf7, 0xe4, 0xb9, 0xe7,
	0xde, 0x87, 0x83, 0x8e, 0xeb, 0xdf, 0x20, 0x65, 0x8e, 0xaf, 0xd3, 0x2b, 0x1a, 0x60, 0xfa, 0xab,
	0x85, 0x82, 0x23, 0x27, 0x6b, 0xf9, 0x9a, 0x26, 0xd1, 0xce, 0x76, 0xc1, 0xb5, 0xb9, 0xa0, 0x3a,
	0xfd, 0x41, 0xed, 0x18, 0x19, 0x0f, 0x52, 0x7a, 0xe7, 0x69, 0x79, 0x99, 0x39, 0x34, 0x40, 0xe6,
	0x32, 0x2a, 0xb2, 0xf5, 0x1d, 0x8f, 0x73, 0xcf, 0xa7, 0xba, 0xfc, 0x37, 0x8a, 0x5d, 0x1d, 0xd9,
	0x98, 0x46, 0x68, 0x8d, 0xc3, 0x8c, 0xf0, 0xa4, 0x4a, 0x88, 0x50, 0xc4, 0x76, 0xe6, 0xa6, 0xf7,
	0xab, 0x0e, 0x9b, 0x5f, 0xb8, 0xb8, 0x74, 0x7d, 0x7e, 0x7d, 0x94, 0x7f, 0xfa, 0x28, 0x31, 0x46,
	0x4e, 0x60, 0xa5, 0x30, 0x63, 0x32, 0x47, 0xad, 0x75, 0x6b, 0x7d, 0x65, 0xb0, 0xab, 0x15, 0xfe,
	0x13, 0x43, 0xda, 0xd4, 0xe6, 0x61, 0xe1, 0xd0, 0x50, 0xe8, 0x04, 0x24, 0x3b, 0xa0, 0x84, 0x82,
	0x3b, 0xb1, 0x4d, 0x45, 0xa2, 0x56, 0xef, 0xd6, 0xfa, 0x6d, 0x03, 0x72, 0x68, 0xe8, 0x90, 0xf7,
	0xd0, 0x0c, 0x2f, 0xac, 0x88, 0xaa, 0x8b, 0xdd, 0x5a, 0x7f, 0x6d, 0xf0, 0xe2, 0x6f, 0x1f, 0xd2,
	0x4e, 0x13, 0xb6, 0x91, 0x6e, 0x22, 0xef, 0x40, 0xe1, 0xb6, 0x1d, 0x0b, 0x41, 0x1d, 0xd3, 0x42,
	0xb5, 0x21, 0xcd, 0x76, 0xb4, 0xf4, 0xf0, 0x5a, 0x7e, 0x78, 0xed, 0x2c, 0x4f, 0xc7, 0x80, 0x9c,
	0x7e, 0x80, 0x64, 0x07, 0x80, 0xc7, 0x18, 0xc6, 0x68, 0xc6, 0x82, 0xa9, 0xcd, 0xc4, 0xda, 0xf1,
	0x82, 0xd1, 0x4e, 0xb1, 0x73, 0xc1, 0xc8, 0x1b, 0x68, 0x52, 0x21, 0xb8, 0x50, 0x5b, 0x52, 0x77,
	0xbb, 0xe2, 0x6d, 0x92, 0x5c, 0x42, 0x3a, 0x5e, 0x30, 0x52, 0xf6, 0xe1, 0x3a, 0xac, 0x66, 0xba,
	0x82, 0x46, 0xb1, 0x8f, 0xbd, 0x9f, 0x0d, 0x20, 0x1f, 0xb9, 0x43, 0x2b, 0x51, 0xbf, 0x85, 0x7a,
	0x11, 0x70, 0xf5, 0xdc, 0x25, 0xfa, 0xad, 0x70, 0xeb, 0xec, 0x1f, 0x32, 0xdd, 0x2f, 0x67, 0xda,
	0x9b, 0xa7, 0x7d, 0x8f, 0x79, 0x6e, 0x41, 0x9b, 0x05, 0xa5, 0x38, 0x8d, 0x65, 0x09, 0x24, 0x59,
	0x96, 0xc3, 0x6e, 0xcd, 0x09, 0x7b, 0xe9, 0x2e, 0x61, 0x93, 0x6f, 0xb0, 0x79, 0x9d, 0xcd, 0x88,
	0x19, 0x70, 0x87, 0x9a, 0x63, 0x8a, 0x96, 0x63, 0xa1, 0xa5, 0x2e, 0x4b, 0x9d, 0xe7, 0x5a, 0xb9,
	0x79, 0xc5, 0x44, 0x25, 0x29, 0x9c, 0x64, 0xdc, 0xe3, 0x9a, 0xb1, 0x71, 0x3d, 0x03, 0x27, 0xdf,
	0x61, 0x23, 0xb4, 0x04, 0x0d, 0xd0, 0x44, 0x2b, 0xba, 0x9c, 0x68, 0xb7, 0xa5, 0xf6, 0xcb, 0xaa,
	0xf6, 0xa9, 0xe4, 0x9e, 0x59, 0xd1, 0x65, 0x61, 0x37, 0x97, 0x32, 0x48, 0x58, 0x2c, 0xe6, 0xd8,
	0xd4, 0xa4, 0x1c, 0x3e, 0x80, 0x75, 0xb4, 0x84, 0x47, 0xb1, 0xf8, 0x54, 0x8f, 0xc2, 0xc6, 0x2c,
	0xcb, 0xf7, 0x5c, 0xd4, 0xde, 0x39, 0x6c, 0xcd, 0x71, 0x3f, 0x77, 0x56, 0x4b, 0x3b, 0xca, 0xb3,
	0xda, 0xfb, 0xdd, 0x00, 0x52, 0x5a, 0x4f, 0x47, 0x7f, 0x00, 0x4b, 0x32, 0xd0, 0x42, 0xf3, 0x71,
	0x45, 0xf3, 0x96, 0x4c, 0x2b, 0x61, 0x0e, 0x1d, 0x62, 0x82, 0x9a, 0xdd, 0x85, 0xbc, 0xe7, 0xd2,
	0xe1, 0xeb, 0x77, 0x2a, 0xd1, 0xc3, 0x54, 0xa7, 0xb2, 0x4c, 0x9e, 0xc1, 0xaa, 0xa0, 0x28, 0x6e,
	0x4c, 0x0b, 0x91, 0x8e, 0x43, 0x94, 0xf5, 0x59, 0x35, 0x56, 0x24, 0x78, 0x90, 0x62, 0x93, 0x6e,
	0x35, 0x66, 0x76, 0xab, 0x74, 0xd6, 0x72, 0xb7, 0x2a, 0xb5, 0x6d, 0x4e, 0xd5, 0x76, 0x17, 0x1a,
	0x3e, 0xf7, 0x22, 0xb5, 0xd5, 0x5d, 0xec, 0x2b, 0x83, 0xcd, 0x19, 0xca, 0x1f, 0xb8, 0x67, 0x48,
	0x4e, 0xb5, 0xa8, 0x4b, 0xff, 0x5f, 0xd4, 0xe5, 0xb9, 0x45, 0x6d, 0xcf, 0x29, 0x2a, 0xdc, 0xa9,
	0xa8, 0xfb, 0xa0, 0xd8, 0x71, 0x84, 0x7c, 0x6c, 0xb2, 0xc0, 0xe5, 0xaa, 0x22, 0x37, 0x3f, 0x9a,
	0x72, 0xfc, 0x49, 0xbe, 0x53, 0x06, 0xa4, 0xdc, 0x61, 0xe0, 0xf2, 0xe4, 0x5e, 0x64, 0x82, 0xe6,
	0x15, 0x15, 0x11, 0xe3, 0x81, 0xba, 0x92, 0xde, 0x8b, 0x04, 0x3f, 0xa7, 0xd8, 0x74, 0x95, 0xf6,
	0xbe, 0xbe, 0xf6, 0x18, 0x5e, 0xc4, 0x23, 0xcd, 0xe6, 0x63, 0xdd, 0xbf, 0x71, 0x51, 0x2f, 0x1e,
	0x55, 0x8f, 0x06, 0x7a, 0x38, 0x7a, 0xe5, 0x71, 0xbd, 0xfc, 0x64, 0x8f, 0x5a, 0xd2, 0xc7, 0xde,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x20, 0xa3, 0x2f, 0xcb, 0x07, 0x00, 0x00,
}