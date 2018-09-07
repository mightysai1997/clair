// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/build/v1/build_events.proto

/*
Package build is a generated protocol buffer package.

It is generated from these files:
	google/devtools/build/v1/build_events.proto
	google/devtools/build/v1/build_status.proto
	google/devtools/build/v1/publish_build_event.proto

It has these top-level messages:
	BuildEvent
	StreamId
	BuildStatus
	PublishLifecycleEventRequest
	PublishBuildToolEventStreamResponse
	OrderedBuildEvent
*/
package build

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf3 "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of console output stream.
type ConsoleOutputStream int32

const (
	// Unspecified or unknown.
	ConsoleOutputStream_UNKNOWN ConsoleOutputStream = 0
	// Normal output stream.
	ConsoleOutputStream_STDOUT ConsoleOutputStream = 1
	// Error output stream.
	ConsoleOutputStream_STDERR ConsoleOutputStream = 2
)

var ConsoleOutputStream_name = map[int32]string{
	0: "UNKNOWN",
	1: "STDOUT",
	2: "STDERR",
}
var ConsoleOutputStream_value = map[string]int32{
	"UNKNOWN": 0,
	"STDOUT":  1,
	"STDERR":  2,
}

func (x ConsoleOutputStream) String() string {
	return proto.EnumName(ConsoleOutputStream_name, int32(x))
}
func (ConsoleOutputStream) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// How did the event stream finish.
type BuildEvent_BuildComponentStreamFinished_FinishType int32

const (
	// Unknown or unspecified; callers should never set this value.
	BuildEvent_BuildComponentStreamFinished_FINISH_TYPE_UNSPECIFIED BuildEvent_BuildComponentStreamFinished_FinishType = 0
	// Set by the event publisher to indicate a build event stream is
	// finished.
	BuildEvent_BuildComponentStreamFinished_FINISHED BuildEvent_BuildComponentStreamFinished_FinishType = 1
	// Set by the WatchBuild RPC server when the publisher of a build event
	// stream stops publishing events without publishing a
	// BuildComponentStreamFinished event whose type equals FINISHED.
	BuildEvent_BuildComponentStreamFinished_EXPIRED BuildEvent_BuildComponentStreamFinished_FinishType = 2
)

var BuildEvent_BuildComponentStreamFinished_FinishType_name = map[int32]string{
	0: "FINISH_TYPE_UNSPECIFIED",
	1: "FINISHED",
	2: "EXPIRED",
}
var BuildEvent_BuildComponentStreamFinished_FinishType_value = map[string]int32{
	"FINISH_TYPE_UNSPECIFIED": 0,
	"FINISHED":                1,
	"EXPIRED":                 2,
}

func (x BuildEvent_BuildComponentStreamFinished_FinishType) String() string {
	return proto.EnumName(BuildEvent_BuildComponentStreamFinished_FinishType_name, int32(x))
}
func (BuildEvent_BuildComponentStreamFinished_FinishType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 5, 0}
}

// Which build component generates this event stream. Each build component
// may generate one event stream.
type StreamId_BuildComponent int32

const (
	// Unknown or unspecified; callers should never set this value.
	StreamId_UNKNOWN_COMPONENT StreamId_BuildComponent = 0
	// A component that coordinates builds.
	StreamId_CONTROLLER StreamId_BuildComponent = 1
	// A component that runs executables needed to complete a build.
	StreamId_WORKER StreamId_BuildComponent = 2
	// A component that builds something.
	StreamId_TOOL       StreamId_BuildComponent = 3
	StreamId_DEPRECATED StreamId_BuildComponent = 4
)

var StreamId_BuildComponent_name = map[int32]string{
	0: "UNKNOWN_COMPONENT",
	1: "CONTROLLER",
	2: "WORKER",
	3: "TOOL",
	4: "DEPRECATED",
}
var StreamId_BuildComponent_value = map[string]int32{
	"UNKNOWN_COMPONENT": 0,
	"CONTROLLER":        1,
	"WORKER":            2,
	"TOOL":              3,
	"DEPRECATED":        4,
}

func (x StreamId_BuildComponent) String() string {
	return proto.EnumName(StreamId_BuildComponent_name, int32(x))
}
func (StreamId_BuildComponent) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// An event representing some state change that occured in the build. This
// message does not include field for uniquely identifying an event.
type BuildEvent struct {
	// The timestamp of this event.
	EventTime *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime" json:"event_time,omitempty"`
	// //////////////////////////////////////////////////////////////////////////
	// Events that indicate a state change of a build request in the build
	// queue.
	//
	// Types that are valid to be assigned to Event:
	//	*BuildEvent_InvocationAttemptStarted_
	//	*BuildEvent_InvocationAttemptFinished_
	//	*BuildEvent_BuildEnqueued_
	//	*BuildEvent_BuildFinished_
	//	*BuildEvent_ConsoleOutput_
	//	*BuildEvent_ComponentStreamFinished
	//	*BuildEvent_BazelEvent
	//	*BuildEvent_BuildExecutionEvent
	//	*BuildEvent_SourceFetchEvent
	Event isBuildEvent_Event `protobuf_oneof:"event"`
}

func (m *BuildEvent) Reset()                    { *m = BuildEvent{} }
func (m *BuildEvent) String() string            { return proto.CompactTextString(m) }
func (*BuildEvent) ProtoMessage()               {}
func (*BuildEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isBuildEvent_Event interface {
	isBuildEvent_Event()
}

type BuildEvent_InvocationAttemptStarted_ struct {
	InvocationAttemptStarted *BuildEvent_InvocationAttemptStarted `protobuf:"bytes,51,opt,name=invocation_attempt_started,json=invocationAttemptStarted,oneof"`
}
type BuildEvent_InvocationAttemptFinished_ struct {
	InvocationAttemptFinished *BuildEvent_InvocationAttemptFinished `protobuf:"bytes,52,opt,name=invocation_attempt_finished,json=invocationAttemptFinished,oneof"`
}
type BuildEvent_BuildEnqueued_ struct {
	BuildEnqueued *BuildEvent_BuildEnqueued `protobuf:"bytes,53,opt,name=build_enqueued,json=buildEnqueued,oneof"`
}
type BuildEvent_BuildFinished_ struct {
	BuildFinished *BuildEvent_BuildFinished `protobuf:"bytes,55,opt,name=build_finished,json=buildFinished,oneof"`
}
type BuildEvent_ConsoleOutput_ struct {
	ConsoleOutput *BuildEvent_ConsoleOutput `protobuf:"bytes,56,opt,name=console_output,json=consoleOutput,oneof"`
}
type BuildEvent_ComponentStreamFinished struct {
	ComponentStreamFinished *BuildEvent_BuildComponentStreamFinished `protobuf:"bytes,59,opt,name=component_stream_finished,json=componentStreamFinished,oneof"`
}
type BuildEvent_BazelEvent struct {
	BazelEvent *google_protobuf1.Any `protobuf:"bytes,60,opt,name=bazel_event,json=bazelEvent,oneof"`
}
type BuildEvent_BuildExecutionEvent struct {
	BuildExecutionEvent *google_protobuf1.Any `protobuf:"bytes,61,opt,name=build_execution_event,json=buildExecutionEvent,oneof"`
}
type BuildEvent_SourceFetchEvent struct {
	SourceFetchEvent *google_protobuf1.Any `protobuf:"bytes,62,opt,name=source_fetch_event,json=sourceFetchEvent,oneof"`
}

func (*BuildEvent_InvocationAttemptStarted_) isBuildEvent_Event()  {}
func (*BuildEvent_InvocationAttemptFinished_) isBuildEvent_Event() {}
func (*BuildEvent_BuildEnqueued_) isBuildEvent_Event()             {}
func (*BuildEvent_BuildFinished_) isBuildEvent_Event()             {}
func (*BuildEvent_ConsoleOutput_) isBuildEvent_Event()             {}
func (*BuildEvent_ComponentStreamFinished) isBuildEvent_Event()    {}
func (*BuildEvent_BazelEvent) isBuildEvent_Event()                 {}
func (*BuildEvent_BuildExecutionEvent) isBuildEvent_Event()        {}
func (*BuildEvent_SourceFetchEvent) isBuildEvent_Event()           {}

func (m *BuildEvent) GetEvent() isBuildEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *BuildEvent) GetEventTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *BuildEvent) GetInvocationAttemptStarted() *BuildEvent_InvocationAttemptStarted {
	if x, ok := m.GetEvent().(*BuildEvent_InvocationAttemptStarted_); ok {
		return x.InvocationAttemptStarted
	}
	return nil
}

func (m *BuildEvent) GetInvocationAttemptFinished() *BuildEvent_InvocationAttemptFinished {
	if x, ok := m.GetEvent().(*BuildEvent_InvocationAttemptFinished_); ok {
		return x.InvocationAttemptFinished
	}
	return nil
}

func (m *BuildEvent) GetBuildEnqueued() *BuildEvent_BuildEnqueued {
	if x, ok := m.GetEvent().(*BuildEvent_BuildEnqueued_); ok {
		return x.BuildEnqueued
	}
	return nil
}

func (m *BuildEvent) GetBuildFinished() *BuildEvent_BuildFinished {
	if x, ok := m.GetEvent().(*BuildEvent_BuildFinished_); ok {
		return x.BuildFinished
	}
	return nil
}

func (m *BuildEvent) GetConsoleOutput() *BuildEvent_ConsoleOutput {
	if x, ok := m.GetEvent().(*BuildEvent_ConsoleOutput_); ok {
		return x.ConsoleOutput
	}
	return nil
}

func (m *BuildEvent) GetComponentStreamFinished() *BuildEvent_BuildComponentStreamFinished {
	if x, ok := m.GetEvent().(*BuildEvent_ComponentStreamFinished); ok {
		return x.ComponentStreamFinished
	}
	return nil
}

func (m *BuildEvent) GetBazelEvent() *google_protobuf1.Any {
	if x, ok := m.GetEvent().(*BuildEvent_BazelEvent); ok {
		return x.BazelEvent
	}
	return nil
}

func (m *BuildEvent) GetBuildExecutionEvent() *google_protobuf1.Any {
	if x, ok := m.GetEvent().(*BuildEvent_BuildExecutionEvent); ok {
		return x.BuildExecutionEvent
	}
	return nil
}

func (m *BuildEvent) GetSourceFetchEvent() *google_protobuf1.Any {
	if x, ok := m.GetEvent().(*BuildEvent_SourceFetchEvent); ok {
		return x.SourceFetchEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BuildEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BuildEvent_OneofMarshaler, _BuildEvent_OneofUnmarshaler, _BuildEvent_OneofSizer, []interface{}{
		(*BuildEvent_InvocationAttemptStarted_)(nil),
		(*BuildEvent_InvocationAttemptFinished_)(nil),
		(*BuildEvent_BuildEnqueued_)(nil),
		(*BuildEvent_BuildFinished_)(nil),
		(*BuildEvent_ConsoleOutput_)(nil),
		(*BuildEvent_ComponentStreamFinished)(nil),
		(*BuildEvent_BazelEvent)(nil),
		(*BuildEvent_BuildExecutionEvent)(nil),
		(*BuildEvent_SourceFetchEvent)(nil),
	}
}

func _BuildEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BuildEvent)
	// event
	switch x := m.Event.(type) {
	case *BuildEvent_InvocationAttemptStarted_:
		b.EncodeVarint(51<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InvocationAttemptStarted); err != nil {
			return err
		}
	case *BuildEvent_InvocationAttemptFinished_:
		b.EncodeVarint(52<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InvocationAttemptFinished); err != nil {
			return err
		}
	case *BuildEvent_BuildEnqueued_:
		b.EncodeVarint(53<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BuildEnqueued); err != nil {
			return err
		}
	case *BuildEvent_BuildFinished_:
		b.EncodeVarint(55<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BuildFinished); err != nil {
			return err
		}
	case *BuildEvent_ConsoleOutput_:
		b.EncodeVarint(56<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConsoleOutput); err != nil {
			return err
		}
	case *BuildEvent_ComponentStreamFinished:
		b.EncodeVarint(59<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ComponentStreamFinished); err != nil {
			return err
		}
	case *BuildEvent_BazelEvent:
		b.EncodeVarint(60<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BazelEvent); err != nil {
			return err
		}
	case *BuildEvent_BuildExecutionEvent:
		b.EncodeVarint(61<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BuildExecutionEvent); err != nil {
			return err
		}
	case *BuildEvent_SourceFetchEvent:
		b.EncodeVarint(62<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SourceFetchEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BuildEvent.Event has unexpected type %T", x)
	}
	return nil
}

func _BuildEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BuildEvent)
	switch tag {
	case 51: // event.invocation_attempt_started
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildEvent_InvocationAttemptStarted)
		err := b.DecodeMessage(msg)
		m.Event = &BuildEvent_InvocationAttemptStarted_{msg}
		return true, err
	case 52: // event.invocation_attempt_finished
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildEvent_InvocationAttemptFinished)
		err := b.DecodeMessage(msg)
		m.Event = &BuildEvent_InvocationAttemptFinished_{msg}
		return true, err
	case 53: // event.build_enqueued
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildEvent_BuildEnqueued)
		err := b.DecodeMessage(msg)
		m.Event = &BuildEvent_BuildEnqueued_{msg}
		return true, err
	case 55: // event.build_finished
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildEvent_BuildFinished)
		err := b.DecodeMessage(msg)
		m.Event = &BuildEvent_BuildFinished_{msg}
		return true, err
	case 56: // event.console_output
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildEvent_ConsoleOutput)
		err := b.DecodeMessage(msg)
		m.Event = &BuildEvent_ConsoleOutput_{msg}
		return true, err
	case 59: // event.component_stream_finished
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildEvent_BuildComponentStreamFinished)
		err := b.DecodeMessage(msg)
		m.Event = &BuildEvent_ComponentStreamFinished{msg}
		return true, err
	case 60: // event.bazel_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Event = &BuildEvent_BazelEvent{msg}
		return true, err
	case 61: // event.build_execution_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Event = &BuildEvent_BuildExecutionEvent{msg}
		return true, err
	case 62: // event.source_fetch_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Event = &BuildEvent_SourceFetchEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BuildEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BuildEvent)
	// event
	switch x := m.Event.(type) {
	case *BuildEvent_InvocationAttemptStarted_:
		s := proto.Size(x.InvocationAttemptStarted)
		n += proto.SizeVarint(51<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildEvent_InvocationAttemptFinished_:
		s := proto.Size(x.InvocationAttemptFinished)
		n += proto.SizeVarint(52<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildEvent_BuildEnqueued_:
		s := proto.Size(x.BuildEnqueued)
		n += proto.SizeVarint(53<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildEvent_BuildFinished_:
		s := proto.Size(x.BuildFinished)
		n += proto.SizeVarint(55<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildEvent_ConsoleOutput_:
		s := proto.Size(x.ConsoleOutput)
		n += proto.SizeVarint(56<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildEvent_ComponentStreamFinished:
		s := proto.Size(x.ComponentStreamFinished)
		n += proto.SizeVarint(59<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildEvent_BazelEvent:
		s := proto.Size(x.BazelEvent)
		n += proto.SizeVarint(60<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildEvent_BuildExecutionEvent:
		s := proto.Size(x.BuildExecutionEvent)
		n += proto.SizeVarint(61<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildEvent_SourceFetchEvent:
		s := proto.Size(x.SourceFetchEvent)
		n += proto.SizeVarint(62<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Notification that the build system has attempted to run the build tool.
type BuildEvent_InvocationAttemptStarted struct {
	// The number of the invocation attempt, starting at 1 and increasing by 1
	// for each new attempt. Can be used to determine if there is a later
	// invocation attempt replacing the current one a client is processing.
	AttemptNumber int64 `protobuf:"varint,1,opt,name=attempt_number,json=attemptNumber" json:"attempt_number,omitempty"`
}

func (m *BuildEvent_InvocationAttemptStarted) Reset()         { *m = BuildEvent_InvocationAttemptStarted{} }
func (m *BuildEvent_InvocationAttemptStarted) String() string { return proto.CompactTextString(m) }
func (*BuildEvent_InvocationAttemptStarted) ProtoMessage()    {}
func (*BuildEvent_InvocationAttemptStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *BuildEvent_InvocationAttemptStarted) GetAttemptNumber() int64 {
	if m != nil {
		return m.AttemptNumber
	}
	return 0
}

// Notification that an invocation attempt has finished.
type BuildEvent_InvocationAttemptFinished struct {
	// The exit code of the build tool.
	ExitCode *google_protobuf3.Int32Value `protobuf:"bytes,2,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
	// Final status of the invocation.
	InvocationStatus *BuildStatus `protobuf:"bytes,3,opt,name=invocation_status,json=invocationStatus" json:"invocation_status,omitempty"`
}

func (m *BuildEvent_InvocationAttemptFinished) Reset()         { *m = BuildEvent_InvocationAttemptFinished{} }
func (m *BuildEvent_InvocationAttemptFinished) String() string { return proto.CompactTextString(m) }
func (*BuildEvent_InvocationAttemptFinished) ProtoMessage()    {}
func (*BuildEvent_InvocationAttemptFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

func (m *BuildEvent_InvocationAttemptFinished) GetExitCode() *google_protobuf3.Int32Value {
	if m != nil {
		return m.ExitCode
	}
	return nil
}

func (m *BuildEvent_InvocationAttemptFinished) GetInvocationStatus() *BuildStatus {
	if m != nil {
		return m.InvocationStatus
	}
	return nil
}

// Notification that the build request is enqueued. It could happen when
// a new build request is inserted into the build queue, or when a
// build request is put back into the build queue due to a previous build
// failure.
type BuildEvent_BuildEnqueued struct {
}

func (m *BuildEvent_BuildEnqueued) Reset()                    { *m = BuildEvent_BuildEnqueued{} }
func (m *BuildEvent_BuildEnqueued) String() string            { return proto.CompactTextString(m) }
func (*BuildEvent_BuildEnqueued) ProtoMessage()               {}
func (*BuildEvent_BuildEnqueued) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

// Notification that the build request has finished, and no further
// invocations will occur.  Note that this applies to the entire Build.
// Individual invocations trigger InvocationFinished when they finish.
type BuildEvent_BuildFinished struct {
	// Final status of the build.
	Status *BuildStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *BuildEvent_BuildFinished) Reset()                    { *m = BuildEvent_BuildFinished{} }
func (m *BuildEvent_BuildFinished) String() string            { return proto.CompactTextString(m) }
func (*BuildEvent_BuildFinished) ProtoMessage()               {}
func (*BuildEvent_BuildFinished) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

func (m *BuildEvent_BuildFinished) GetStatus() *BuildStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Textual output written to standard output or standard error.
type BuildEvent_ConsoleOutput struct {
	// The output stream type.
	Type ConsoleOutputStream `protobuf:"varint,1,opt,name=type,enum=google.devtools.build.v1.ConsoleOutputStream" json:"type,omitempty"`
	// The output stream content.
	//
	// Types that are valid to be assigned to Output:
	//	*BuildEvent_ConsoleOutput_TextOutput
	//	*BuildEvent_ConsoleOutput_BinaryOutput
	Output isBuildEvent_ConsoleOutput_Output `protobuf_oneof:"output"`
}

func (m *BuildEvent_ConsoleOutput) Reset()                    { *m = BuildEvent_ConsoleOutput{} }
func (m *BuildEvent_ConsoleOutput) String() string            { return proto.CompactTextString(m) }
func (*BuildEvent_ConsoleOutput) ProtoMessage()               {}
func (*BuildEvent_ConsoleOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 4} }

type isBuildEvent_ConsoleOutput_Output interface {
	isBuildEvent_ConsoleOutput_Output()
}

type BuildEvent_ConsoleOutput_TextOutput struct {
	TextOutput string `protobuf:"bytes,2,opt,name=text_output,json=textOutput,oneof"`
}
type BuildEvent_ConsoleOutput_BinaryOutput struct {
	BinaryOutput []byte `protobuf:"bytes,3,opt,name=binary_output,json=binaryOutput,proto3,oneof"`
}

func (*BuildEvent_ConsoleOutput_TextOutput) isBuildEvent_ConsoleOutput_Output()   {}
func (*BuildEvent_ConsoleOutput_BinaryOutput) isBuildEvent_ConsoleOutput_Output() {}

func (m *BuildEvent_ConsoleOutput) GetOutput() isBuildEvent_ConsoleOutput_Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *BuildEvent_ConsoleOutput) GetType() ConsoleOutputStream {
	if m != nil {
		return m.Type
	}
	return ConsoleOutputStream_UNKNOWN
}

func (m *BuildEvent_ConsoleOutput) GetTextOutput() string {
	if x, ok := m.GetOutput().(*BuildEvent_ConsoleOutput_TextOutput); ok {
		return x.TextOutput
	}
	return ""
}

func (m *BuildEvent_ConsoleOutput) GetBinaryOutput() []byte {
	if x, ok := m.GetOutput().(*BuildEvent_ConsoleOutput_BinaryOutput); ok {
		return x.BinaryOutput
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BuildEvent_ConsoleOutput) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BuildEvent_ConsoleOutput_OneofMarshaler, _BuildEvent_ConsoleOutput_OneofUnmarshaler, _BuildEvent_ConsoleOutput_OneofSizer, []interface{}{
		(*BuildEvent_ConsoleOutput_TextOutput)(nil),
		(*BuildEvent_ConsoleOutput_BinaryOutput)(nil),
	}
}

func _BuildEvent_ConsoleOutput_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BuildEvent_ConsoleOutput)
	// output
	switch x := m.Output.(type) {
	case *BuildEvent_ConsoleOutput_TextOutput:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TextOutput)
	case *BuildEvent_ConsoleOutput_BinaryOutput:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BinaryOutput)
	case nil:
	default:
		return fmt.Errorf("BuildEvent_ConsoleOutput.Output has unexpected type %T", x)
	}
	return nil
}

func _BuildEvent_ConsoleOutput_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BuildEvent_ConsoleOutput)
	switch tag {
	case 2: // output.text_output
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Output = &BuildEvent_ConsoleOutput_TextOutput{x}
		return true, err
	case 3: // output.binary_output
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Output = &BuildEvent_ConsoleOutput_BinaryOutput{x}
		return true, err
	default:
		return false, nil
	}
}

func _BuildEvent_ConsoleOutput_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BuildEvent_ConsoleOutput)
	// output
	switch x := m.Output.(type) {
	case *BuildEvent_ConsoleOutput_TextOutput:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TextOutput)))
		n += len(x.TextOutput)
	case *BuildEvent_ConsoleOutput_BinaryOutput:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.BinaryOutput)))
		n += len(x.BinaryOutput)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Notification of the end of a build event stream published by a build
// component other than CONTROLLER (See StreamId.BuildComponents).
type BuildEvent_BuildComponentStreamFinished struct {
	// How the event stream finished.
	Type BuildEvent_BuildComponentStreamFinished_FinishType `protobuf:"varint,1,opt,name=type,enum=google.devtools.build.v1.BuildEvent_BuildComponentStreamFinished_FinishType" json:"type,omitempty"`
}

func (m *BuildEvent_BuildComponentStreamFinished) Reset() {
	*m = BuildEvent_BuildComponentStreamFinished{}
}
func (m *BuildEvent_BuildComponentStreamFinished) String() string { return proto.CompactTextString(m) }
func (*BuildEvent_BuildComponentStreamFinished) ProtoMessage()    {}
func (*BuildEvent_BuildComponentStreamFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 5}
}

func (m *BuildEvent_BuildComponentStreamFinished) GetType() BuildEvent_BuildComponentStreamFinished_FinishType {
	if m != nil {
		return m.Type
	}
	return BuildEvent_BuildComponentStreamFinished_FINISH_TYPE_UNSPECIFIED
}

// Unique identifier for a build event stream.
type StreamId struct {
	// The id of a Build message.
	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	// The unique invocation ID within this build.
	// It should be the same as {invocation} (below) during the migration.
	InvocationId string `protobuf:"bytes,6,opt,name=invocation_id,json=invocationId" json:"invocation_id,omitempty"`
	// The component that emitted this event.
	Component StreamId_BuildComponent `protobuf:"varint,3,opt,name=component,enum=google.devtools.build.v1.StreamId_BuildComponent" json:"component,omitempty"`
	// The unique invocation ID within this build.
	// It should be the same as {invocation_id} below during the migration.
	Invocation string `protobuf:"bytes,4,opt,name=invocation" json:"invocation,omitempty"`
}

func (m *StreamId) Reset()                    { *m = StreamId{} }
func (m *StreamId) String() string            { return proto.CompactTextString(m) }
func (*StreamId) ProtoMessage()               {}
func (*StreamId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StreamId) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *StreamId) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

func (m *StreamId) GetComponent() StreamId_BuildComponent {
	if m != nil {
		return m.Component
	}
	return StreamId_UNKNOWN_COMPONENT
}

func (m *StreamId) GetInvocation() string {
	if m != nil {
		return m.Invocation
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildEvent)(nil), "google.devtools.build.v1.BuildEvent")
	proto.RegisterType((*BuildEvent_InvocationAttemptStarted)(nil), "google.devtools.build.v1.BuildEvent.InvocationAttemptStarted")
	proto.RegisterType((*BuildEvent_InvocationAttemptFinished)(nil), "google.devtools.build.v1.BuildEvent.InvocationAttemptFinished")
	proto.RegisterType((*BuildEvent_BuildEnqueued)(nil), "google.devtools.build.v1.BuildEvent.BuildEnqueued")
	proto.RegisterType((*BuildEvent_BuildFinished)(nil), "google.devtools.build.v1.BuildEvent.BuildFinished")
	proto.RegisterType((*BuildEvent_ConsoleOutput)(nil), "google.devtools.build.v1.BuildEvent.ConsoleOutput")
	proto.RegisterType((*BuildEvent_BuildComponentStreamFinished)(nil), "google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished")
	proto.RegisterType((*StreamId)(nil), "google.devtools.build.v1.StreamId")
	proto.RegisterEnum("google.devtools.build.v1.ConsoleOutputStream", ConsoleOutputStream_name, ConsoleOutputStream_value)
	proto.RegisterEnum("google.devtools.build.v1.BuildEvent_BuildComponentStreamFinished_FinishType", BuildEvent_BuildComponentStreamFinished_FinishType_name, BuildEvent_BuildComponentStreamFinished_FinishType_value)
	proto.RegisterEnum("google.devtools.build.v1.StreamId_BuildComponent", StreamId_BuildComponent_name, StreamId_BuildComponent_value)
}

func init() { proto.RegisterFile("google/devtools/build/v1/build_events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 948 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xed, 0x6e, 0xe3, 0x54,
	0x10, 0x8d, 0xdb, 0xd2, 0x26, 0xd3, 0x24, 0xeb, 0xbd, 0xcb, 0xaa, 0x8e, 0x5b, 0x2d, 0x50, 0x54,
	0x09, 0x81, 0x70, 0xd4, 0x14, 0xb4, 0x0b, 0x4b, 0x57, 0xca, 0x87, 0xab, 0x98, 0x2d, 0x76, 0x74,
	0xe3, 0xb2, 0x7c, 0x08, 0x05, 0xc7, 0xbe, 0xcd, 0x5a, 0x4a, 0x7c, 0x8d, 0x7d, 0x1d, 0x1a, 0x24,
	0x04, 0xcf, 0xc1, 0x03, 0x20, 0xf1, 0x22, 0xbc, 0x12, 0xfc, 0x44, 0xbe, 0xd7, 0x6e, 0x92, 0xb6,
	0xe9, 0x6e, 0xd9, 0x7f, 0xf6, 0xcc, 0x99, 0x73, 0x66, 0xc6, 0x67, 0xa2, 0xc0, 0x47, 0x23, 0x4a,
	0x47, 0x63, 0x52, 0xf7, 0xc8, 0x94, 0x51, 0x3a, 0x8e, 0xeb, 0xc3, 0xc4, 0x1f, 0x7b, 0xf5, 0xe9,
	0xa1, 0x78, 0x18, 0x90, 0x29, 0x09, 0x58, 0xac, 0x85, 0x11, 0x65, 0x14, 0x29, 0x02, 0xac, 0xe5,
	0x60, 0x8d, 0x63, 0xb4, 0xe9, 0xa1, 0xba, 0x97, 0xd1, 0x38, 0xa1, 0x5f, 0x77, 0x82, 0x80, 0x32,
	0x87, 0xf9, 0x34, 0xc8, 0xea, 0xd4, 0x57, 0x89, 0xc4, 0xcc, 0x61, 0x49, 0x0e, 0xae, 0x65, 0x60,
	0xfe, 0x36, 0x4c, 0xce, 0xeb, 0x4e, 0x30, 0xcb, 0x52, 0xef, 0x5c, 0x4d, 0x31, 0x7f, 0x42, 0x62,
	0xe6, 0x4c, 0xc2, 0x0c, 0xf0, 0xe8, 0x2a, 0xe0, 0xe7, 0xc8, 0x09, 0x43, 0x12, 0xe5, 0xdc, 0x3b,
	0x59, 0x3e, 0x0a, 0xdd, 0xfa, 0xa2, 0xe8, 0xfe, 0x3f, 0x65, 0x80, 0x56, 0xda, 0x8b, 0x9e, 0xce,
	0x8b, 0x3e, 0x03, 0xe0, 0x83, 0x0f, 0x52, 0x01, 0x45, 0x7a, 0x57, 0xfa, 0x60, 0xbb, 0xa1, 0x6a,
	0xd9, 0xf4, 0x39, 0xb9, 0x66, 0xe7, 0xea, 0xb8, 0xc4, 0xd1, 0xe9, 0x3b, 0xfa, 0x15, 0x54, 0x3f,
	0x98, 0x52, 0x97, 0x2f, 0x60, 0xe0, 0x30, 0x46, 0x26, 0x21, 0x4b, 0x27, 0x8c, 0x18, 0xf1, 0x94,
	0x23, 0x4e, 0x75, 0xac, 0xad, 0x5a, 0xa4, 0x36, 0x6f, 0x42, 0x33, 0x2e, 0x69, 0x9a, 0x82, 0xa5,
	0x2f, 0x48, 0xba, 0x05, 0xac, 0xf8, 0x2b, 0x72, 0xe8, 0x77, 0x09, 0x76, 0x6f, 0xd0, 0x3f, 0xf7,
	0x03, 0x3f, 0x7e, 0x49, 0x3c, 0xe5, 0x13, 0xde, 0xc0, 0xb3, 0xff, 0xd7, 0xc0, 0x49, 0xc6, 0xd2,
	0x2d, 0xe0, 0x9a, 0xbf, 0x2a, 0x89, 0xbe, 0x87, 0x6a, 0xe6, 0x9d, 0xe0, 0xa7, 0x84, 0x24, 0xc4,
	0x53, 0x3e, 0xe5, 0xa2, 0x8d, 0xd7, 0x12, 0x15, 0x8f, 0x59, 0x65, 0xb7, 0x80, 0x2b, 0xc3, 0xc5,
	0xc0, 0x9c, 0xfc, 0x72, 0xa2, 0xc7, 0x77, 0x25, 0x5f, 0x98, 0x42, 0x90, 0x2f, 0x76, 0xee, 0xd2,
	0x20, 0xa6, 0x63, 0x32, 0xa0, 0x09, 0x0b, 0x13, 0xa6, 0x3c, 0xb9, 0x03, 0x79, 0x5b, 0x94, 0x5a,
	0xbc, 0x32, 0x25, 0x77, 0x17, 0x03, 0xe8, 0x37, 0xa8, 0xb9, 0x74, 0x12, 0xd2, 0x20, 0xf5, 0x55,
	0xcc, 0x22, 0xe2, 0x4c, 0xe6, 0x43, 0x3c, 0xe5, 0x3a, 0xcd, 0xd7, 0x1f, 0xa2, 0x9d, 0x53, 0xf5,
	0x39, 0xd3, 0xc2, 0x4c, 0x3b, 0xee, 0xcd, 0x29, 0xf4, 0x18, 0xb6, 0x87, 0xce, 0x2f, 0x64, 0x2c,
	0x6e, 0x5a, 0xf9, 0x82, 0x4b, 0xbe, 0x7d, 0xcd, 0xd5, 0xcd, 0x60, 0xd6, 0x2d, 0x60, 0xe0, 0x50,
	0x71, 0x0d, 0x5f, 0xc2, 0xc3, 0xec, 0x83, 0x5e, 0x10, 0x37, 0xe1, 0xbe, 0x12, 0x14, 0xc7, 0xb7,
	0x52, 0x3c, 0x10, 0x5f, 0x2e, 0xaf, 0x11, 0x5c, 0x1d, 0x40, 0x31, 0x4d, 0x22, 0x97, 0x0c, 0xce,
	0x09, 0x73, 0x5f, 0x66, 0x44, 0xcf, 0x6e, 0x25, 0x92, 0x45, 0xc5, 0x49, 0x5a, 0xc0, 0x59, 0xd4,
	0x26, 0x28, 0xab, 0xae, 0x03, 0x1d, 0x40, 0x35, 0x77, 0x7d, 0x90, 0x4c, 0x86, 0x24, 0xe2, 0xf7,
	0xbb, 0x8e, 0x2b, 0x59, 0xd4, 0xe4, 0x41, 0xf5, 0x2f, 0x09, 0x6a, 0x2b, 0x0d, 0x8e, 0x9e, 0x40,
	0x89, 0x5c, 0xf8, 0x6c, 0xe0, 0x52, 0x8f, 0x28, 0x6b, 0xbc, 0xbb, 0xdd, 0x6b, 0xdd, 0x19, 0x01,
	0x3b, 0x6a, 0x7c, 0xed, 0x8c, 0x13, 0x82, 0x8b, 0x29, 0xba, 0x4d, 0x3d, 0x82, 0x30, 0xdc, 0x5f,
	0xb8, 0x3f, 0xf1, 0x23, 0xa3, 0xac, 0x73, 0x86, 0x83, 0x57, 0x7c, 0xde, 0x3e, 0x07, 0x63, 0x79,
	0x5e, 0x2f, 0x22, 0xea, 0x3d, 0xa8, 0x2c, 0x9d, 0x85, 0x6a, 0x66, 0x81, 0xcb, 0x7e, 0x8f, 0x61,
	0x33, 0x93, 0x92, 0xee, 0x22, 0x95, 0x15, 0xa9, 0x7f, 0x4a, 0x50, 0x59, 0xb2, 0x2f, 0x6a, 0xc2,
	0x06, 0x9b, 0x85, 0xe2, 0xb7, 0xaf, 0xda, 0xf8, 0x78, 0x35, 0xdd, 0x52, 0x99, 0x70, 0x1c, 0xe6,
	0xa5, 0xe8, 0x3d, 0xd8, 0x66, 0xe4, 0x82, 0xe5, 0xa7, 0x94, 0x6e, 0xb1, 0x94, 0x3a, 0x2b, 0x0d,
	0x66, 0x2a, 0x07, 0x50, 0x19, 0xfa, 0x81, 0x13, 0xcd, 0x72, 0x50, 0xba, 0xa8, 0x72, 0xb7, 0x80,
	0xcb, 0x22, 0x2c, 0x60, 0xad, 0x22, 0x6c, 0x8a, 0xbc, 0xfa, 0xb7, 0x04, 0x7b, 0xb7, 0xf9, 0x1f,
	0xfd, 0xb8, 0xd4, 0xf7, 0xe9, 0x1b, 0x1f, 0x94, 0x26, 0x1e, 0xec, 0x59, 0x48, 0xc4, 0x58, 0xfb,
	0x1d, 0x80, 0x79, 0x0c, 0xed, 0xc2, 0xce, 0x89, 0x61, 0x1a, 0xfd, 0xee, 0xc0, 0xfe, 0xb6, 0xa7,
	0x0f, 0xce, 0xcc, 0x7e, 0x4f, 0x6f, 0x1b, 0x27, 0x86, 0xde, 0x91, 0x0b, 0xa8, 0x0c, 0x45, 0x91,
	0xd4, 0x3b, 0xb2, 0x84, 0xb6, 0x61, 0x4b, 0xff, 0xa6, 0x67, 0x60, 0xbd, 0x23, 0xaf, 0xb5, 0xb6,
	0xe0, 0x2d, 0x6e, 0xfd, 0xfd, 0x3f, 0xd6, 0xa0, 0x28, 0x24, 0x0d, 0x0f, 0xd5, 0xa0, 0x28, 0x2e,
	0xcd, 0xf7, 0xf8, 0x04, 0x25, 0xbc, 0xc5, 0xdf, 0x0d, 0x0f, 0xbd, 0x0f, 0x95, 0x05, 0x5f, 0xf9,
	0x9e, 0xb2, 0xc9, 0xf3, 0xe5, 0x79, 0xd0, 0xf0, 0x90, 0x05, 0xa5, 0xcb, 0xeb, 0xe7, 0xbb, 0xac,
	0x36, 0x0e, 0x57, 0xaf, 0x20, 0x97, 0xbd, 0xb2, 0x00, 0x3c, 0xe7, 0x40, 0x8f, 0x00, 0xe6, 0x02,
	0xca, 0x06, 0x97, 0x5c, 0x88, 0xec, 0xff, 0x00, 0xd5, 0xe5, 0x62, 0xf4, 0x10, 0xee, 0x9f, 0x99,
	0xcf, 0x4d, 0xeb, 0x85, 0x39, 0x68, 0x5b, 0x5f, 0xf5, 0x2c, 0x53, 0x37, 0x6d, 0xb9, 0x80, 0xaa,
	0x00, 0x6d, 0xcb, 0xb4, 0xb1, 0x75, 0x7a, 0xaa, 0x63, 0x59, 0x42, 0x00, 0x9b, 0x2f, 0x2c, 0xfc,
	0x5c, 0xc7, 0xf2, 0x1a, 0x2a, 0xc2, 0x86, 0x6d, 0x59, 0xa7, 0xf2, 0x7a, 0x8a, 0xea, 0xe8, 0x3d,
	0xac, 0xb7, 0x9b, 0xb6, 0xde, 0x91, 0x37, 0x3e, 0xfc, 0x1c, 0x1e, 0xdc, 0xe0, 0xaf, 0x74, 0x93,
	0x99, 0x86, 0x5c, 0x48, 0x99, 0xfa, 0x76, 0xc7, 0x3a, 0xb3, 0x05, 0x6b, 0xdf, 0xee, 0xe8, 0x18,
	0xcb, 0x6b, 0xad, 0x18, 0xf6, 0x5c, 0x3a, 0x59, 0x39, 0x7d, 0xeb, 0xde, 0xdc, 0x01, 0xbd, 0xf4,
	0xa2, 0x7b, 0xd2, 0x77, 0xc7, 0x19, 0x78, 0x44, 0xc7, 0x4e, 0x30, 0xd2, 0x68, 0x34, 0xaa, 0x8f,
	0x48, 0xc0, 0xef, 0xbd, 0x2e, 0x52, 0x4e, 0xe8, 0xc7, 0xd7, 0xff, 0xc6, 0x3c, 0xe5, 0x0f, 0xff,
	0x4a, 0xd2, 0x70, 0x93, 0x83, 0x8f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x95, 0xa8, 0xb7, 0x58,
	0x57, 0x09, 0x00, 0x00,
}