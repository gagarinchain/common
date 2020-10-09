// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: event.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Event_EventType int32

const (
	Event_RESERVED      Event_EventType = 0
	Event_BLOCK_ADDED   Event_EventType = 1
	Event_EPOCH_STARTED Event_EventType = 2
	Event_VIEW_CHANGED  Event_EventType = 3
	Event_COMMITTED     Event_EventType = 4
	Event_ACCOUNT       Event_EventType = 5
	Event_BLOCK         Event_EventType = 6
)

// Enum value maps for Event_EventType.
var (
	Event_EventType_name = map[int32]string{
		0: "RESERVED",
		1: "BLOCK_ADDED",
		2: "EPOCH_STARTED",
		3: "VIEW_CHANGED",
		4: "COMMITTED",
		5: "ACCOUNT",
		6: "BLOCK",
	}
	Event_EventType_value = map[string]int32{
		"RESERVED":      0,
		"BLOCK_ADDED":   1,
		"EPOCH_STARTED": 2,
		"VIEW_CHANGED":  3,
		"COMMITTED":     4,
		"ACCOUNT":       5,
		"BLOCK":         6,
	}
)

func (x Event_EventType) Enum() *Event_EventType {
	p := new(Event_EventType)
	*p = x
	return p
}

func (x Event_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[0].Descriptor()
}

func (Event_EventType) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[0]
}

func (x Event_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_EventType.Descriptor instead.
func (Event_EventType) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0, 0}
}

type Request_RequestType int32

const (
	Request_ACCOUNT Request_RequestType = 0
	Request_BLOCK   Request_RequestType = 1
)

// Enum value maps for Request_RequestType.
var (
	Request_RequestType_name = map[int32]string{
		0: "ACCOUNT",
		1: "BLOCK",
	}
	Request_RequestType_value = map[string]int32{
		"ACCOUNT": 0,
		"BLOCK":   1,
	}
)

func (x Request_RequestType) Enum() *Request_RequestType {
	p := new(Request_RequestType)
	*p = x
	return p
}

func (x Request_RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Request_RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[1].Descriptor()
}

func (Request_RequestType) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[1]
}

func (x Request_RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_RequestType.Descriptor instead.
func (Request_RequestType) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1, 0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=gagarin.network.event.Event_EventType" json:"type,omitempty"`
	Id      []byte          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Payload *any.Any        `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetType() Event_EventType {
	if x != nil {
		return x.Type
	}
	return Event_RESERVED
}

func (x *Event) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Event) GetPayload() *any.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    Request_RequestType `protobuf:"varint,1,opt,name=type,proto3,enum=gagarin.network.event.Request_RequestType" json:"type,omitempty"`
	Id      []byte              `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Payload *any.Any            `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetType() Request_RequestType {
	if x != nil {
		return x.Type
	}
	return Request_ACCOUNT
}

func (x *Request) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Request) GetPayload() *any.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

type EpochStartedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch int32 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
}

func (x *EpochStartedPayload) Reset() {
	*x = EpochStartedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochStartedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochStartedPayload) ProtoMessage() {}

func (x *EpochStartedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochStartedPayload.ProtoReflect.Descriptor instead.
func (*EpochStartedPayload) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{2}
}

func (x *EpochStartedPayload) GetEpoch() int32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

type ViewChangedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View int32 `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
}

func (x *ViewChangedPayload) Reset() {
	*x = ViewChangedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewChangedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewChangedPayload) ProtoMessage() {}

func (x *ViewChangedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewChangedPayload.ProtoReflect.Descriptor instead.
func (*ViewChangedPayload) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{3}
}

func (x *ViewChangedPayload) GetView() int32 {
	if x != nil {
		return x.View
	}
	return 0
}

type CommittedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *CommittedPayload) Reset() {
	*x = CommittedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommittedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommittedPayload) ProtoMessage() {}

func (x *CommittedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommittedPayload.ProtoReflect.Descriptor instead.
func (*CommittedPayload) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{4}
}

func (x *CommittedPayload) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type AccountRequestPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Block   []byte `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *AccountRequestPayload) Reset() {
	*x = AccountRequestPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountRequestPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRequestPayload) ProtoMessage() {}

func (x *AccountRequestPayload) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRequestPayload.ProtoReflect.Descriptor instead.
func (*AccountRequestPayload) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{5}
}

func (x *AccountRequestPayload) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *AccountRequestPayload) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

type AccountResponsePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountE `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AccountResponsePayload) Reset() {
	*x = AccountResponsePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountResponsePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResponsePayload) ProtoMessage() {}

func (x *AccountResponsePayload) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResponsePayload.ProtoReflect.Descriptor instead.
func (*AccountResponsePayload) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{6}
}

func (x *AccountResponsePayload) GetAccount() *AccountE {
	if x != nil {
		return x.Account
	}
	return nil
}

type AccountUpdatedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Old *AccountE `protobuf:"bytes,1,opt,name=old,proto3" json:"old,omitempty"`
	New *AccountE `protobuf:"bytes,2,opt,name=new,proto3" json:"new,omitempty"`
}

func (x *AccountUpdatedPayload) Reset() {
	*x = AccountUpdatedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountUpdatedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountUpdatedPayload) ProtoMessage() {}

func (x *AccountUpdatedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountUpdatedPayload.ProtoReflect.Descriptor instead.
func (*AccountUpdatedPayload) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{7}
}

func (x *AccountUpdatedPayload) GetOld() *AccountE {
	if x != nil {
		return x.Old
	}
	return nil
}

func (x *AccountUpdatedPayload) GetNew() *AccountE {
	if x != nil {
		return x.New
	}
	return nil
}

type AccountE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Block   []byte   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Nonce   uint64   `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Value   uint64   `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	Proof   [][]byte `protobuf:"bytes,5,rep,name=proof,proto3" json:"proof,omitempty"`
}

func (x *AccountE) Reset() {
	*x = AccountE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountE) ProtoMessage() {}

func (x *AccountE) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountE.ProtoReflect.Descriptor instead.
func (*AccountE) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{8}
}

func (x *AccountE) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *AccountE) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *AccountE) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *AccountE) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *AccountE) GetProof() [][]byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67,
	0x61, 0x67, 0x61, 0x72, 0x69, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfb, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x61, 0x67, 0x61, 0x72, 0x69,
	0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x76, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x50, 0x4f, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54,
	0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x06, 0x22, 0xb0, 0x01,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x61, 0x67, 0x61, 0x72, 0x69,
	0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x25, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x01,
	0x22, 0x2b, 0x0a, 0x13, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x28, 0x0a,
	0x12, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22,
	0x47, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x53, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x39, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x67, 0x61, 0x72, 0x69, 0x6e, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x45, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7d, 0x0a,
	0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x31, 0x0a, 0x03, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x67, 0x61, 0x72, 0x69, 0x6e, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x45, 0x52, 0x03, 0x6f, 0x6c, 0x64, 0x12, 0x31, 0x0a, 0x03, 0x6e, 0x65, 0x77,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x67, 0x61, 0x72, 0x69, 0x6e,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x22, 0x7c, 0x0a, 0x08,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x42, 0x04, 0x5a, 0x02, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_event_proto_goTypes = []interface{}{
	(Event_EventType)(0),           // 0: gagarin.network.event.Event.EventType
	(Request_RequestType)(0),       // 1: gagarin.network.event.Request.RequestType
	(*Event)(nil),                  // 2: gagarin.network.event.Event
	(*Request)(nil),                // 3: gagarin.network.event.Request
	(*EpochStartedPayload)(nil),    // 4: gagarin.network.event.EpochStartedPayload
	(*ViewChangedPayload)(nil),     // 5: gagarin.network.event.ViewChangedPayload
	(*CommittedPayload)(nil),       // 6: gagarin.network.event.CommittedPayload
	(*AccountRequestPayload)(nil),  // 7: gagarin.network.event.AccountRequestPayload
	(*AccountResponsePayload)(nil), // 8: gagarin.network.event.AccountResponsePayload
	(*AccountUpdatedPayload)(nil),  // 9: gagarin.network.event.AccountUpdatedPayload
	(*AccountE)(nil),               // 10: gagarin.network.event.AccountE
	(*any.Any)(nil),                // 11: google.protobuf.Any
}
var file_event_proto_depIdxs = []int32{
	0,  // 0: gagarin.network.event.Event.type:type_name -> gagarin.network.event.Event.EventType
	11, // 1: gagarin.network.event.Event.payload:type_name -> google.protobuf.Any
	1,  // 2: gagarin.network.event.Request.type:type_name -> gagarin.network.event.Request.RequestType
	11, // 3: gagarin.network.event.Request.payload:type_name -> google.protobuf.Any
	10, // 4: gagarin.network.event.AccountResponsePayload.account:type_name -> gagarin.network.event.AccountE
	10, // 5: gagarin.network.event.AccountUpdatedPayload.old:type_name -> gagarin.network.event.AccountE
	10, // 6: gagarin.network.event.AccountUpdatedPayload.new:type_name -> gagarin.network.event.AccountE
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochStartedPayload); i {
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
		file_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewChangedPayload); i {
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
		file_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommittedPayload); i {
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
		file_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountRequestPayload); i {
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
		file_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountResponsePayload); i {
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
		file_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountUpdatedPayload); i {
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
		file_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountE); i {
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
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		EnumInfos:         file_event_proto_enumTypes,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
