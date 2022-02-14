// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: pb/session.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type SessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consumer   *ConsumerInfo `protobuf:"bytes,1,opt,name=consumer,proto3" json:"consumer,omitempty"`
	ProposalID int64         `protobuf:"varint,2,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	Config     []byte        `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *SessionRequest) Reset() {
	*x = SessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionRequest) ProtoMessage() {}

func (x *SessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionRequest.ProtoReflect.Descriptor instead.
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return file_pb_session_proto_rawDescGZIP(), []int{0}
}

func (x *SessionRequest) GetConsumer() *ConsumerInfo {
	if x != nil {
		return x.Consumer
	}
	return nil
}

func (x *SessionRequest) GetProposalID() int64 {
	if x != nil {
		return x.ProposalID
	}
	return 0
}

func (x *SessionRequest) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type SessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PaymentInfo string `protobuf:"bytes,2,opt,name=PaymentInfo,proto3" json:"PaymentInfo,omitempty"`
	Config      []byte `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *SessionResponse) Reset() {
	*x = SessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionResponse) ProtoMessage() {}

func (x *SessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionResponse.ProtoReflect.Descriptor instead.
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return file_pb_session_proto_rawDescGZIP(), []int{1}
}

func (x *SessionResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *SessionResponse) GetPaymentInfo() string {
	if x != nil {
		return x.PaymentInfo
	}
	return ""
}

func (x *SessionResponse) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type SessionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerID string `protobuf:"bytes,1,opt,name=consumerID,proto3" json:"consumerID,omitempty"`
	SessionID  string `protobuf:"bytes,2,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
}

func (x *SessionInfo) Reset() {
	*x = SessionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionInfo) ProtoMessage() {}

func (x *SessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionInfo.ProtoReflect.Descriptor instead.
func (*SessionInfo) Descriptor() ([]byte, []int) {
	return file_pb_session_proto_rawDescGZIP(), []int{2}
}

func (x *SessionInfo) GetConsumerID() string {
	if x != nil {
		return x.ConsumerID
	}
	return ""
}

func (x *SessionInfo) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

type ConsumerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HermesID       string        `protobuf:"bytes,2,opt,name=hermesID,proto3" json:"hermesID,omitempty"`
	PaymentVersion string        `protobuf:"bytes,3,opt,name=paymentVersion,proto3" json:"paymentVersion,omitempty"`
	Location       *LocationInfo `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Pricing        *Pricing      `protobuf:"bytes,5,opt,name=pricing,proto3" json:"pricing,omitempty"`
}

func (x *ConsumerInfo) Reset() {
	*x = ConsumerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerInfo) ProtoMessage() {}

func (x *ConsumerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerInfo.ProtoReflect.Descriptor instead.
func (*ConsumerInfo) Descriptor() ([]byte, []int) {
	return file_pb_session_proto_rawDescGZIP(), []int{3}
}

func (x *ConsumerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConsumerInfo) GetHermesID() string {
	if x != nil {
		return x.HermesID
	}
	return ""
}

func (x *ConsumerInfo) GetPaymentVersion() string {
	if x != nil {
		return x.PaymentVersion
	}
	return ""
}

func (x *ConsumerInfo) GetLocation() *LocationInfo {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ConsumerInfo) GetPricing() *Pricing {
	if x != nil {
		return x.Pricing
	}
	return nil
}

type LocationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *LocationInfo) Reset() {
	*x = LocationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationInfo) ProtoMessage() {}

func (x *LocationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationInfo.ProtoReflect.Descriptor instead.
func (*LocationInfo) Descriptor() ([]byte, []int) {
	return file_pb_session_proto_rawDescGZIP(), []int{4}
}

func (x *LocationInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Pricing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PerGib  []byte `protobuf:"bytes,1,opt,name=PerGib,proto3" json:"PerGib,omitempty"`
	PerHour []byte `protobuf:"bytes,2,opt,name=PerHour,proto3" json:"PerHour,omitempty"`
}

func (x *Pricing) Reset() {
	*x = Pricing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pricing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pricing) ProtoMessage() {}

func (x *Pricing) ProtoReflect() protoreflect.Message {
	mi := &file_pb_session_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pricing.ProtoReflect.Descriptor instead.
func (*Pricing) Descriptor() ([]byte, []int) {
	return file_pb_session_proto_rawDescGZIP(), []int{5}
}

func (x *Pricing) GetPerGib() []byte {
	if x != nil {
		return x.PerGib
	}
	return nil
}

func (x *Pricing) GetPerHour() []byte {
	if x != nil {
		return x.PerHour
	}
	return nil
}

type SessionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerID string `protobuf:"bytes,1,opt,name=ConsumerID,proto3" json:"ConsumerID,omitempty"`
	SessionID  string `protobuf:"bytes,2,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Code       uint32 `protobuf:"varint,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Message    string `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *SessionStatus) Reset() {
	*x = SessionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_session_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStatus) ProtoMessage() {}

func (x *SessionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pb_session_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStatus.ProtoReflect.Descriptor instead.
func (*SessionStatus) Descriptor() ([]byte, []int) {
	return file_pb_session_proto_rawDescGZIP(), []int{6}
}

func (x *SessionStatus) GetConsumerID() string {
	if x != nil {
		return x.ConsumerID
	}
	return ""
}

func (x *SessionStatus) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *SessionStatus) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SessionStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pb_session_proto protoreflect.FileDescriptor

var file_pb_session_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x76, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x5b,
	0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4b, 0x0a, 0x0b, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x72,
	0x6d, 0x65, 0x73, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x72,
	0x6d, 0x65, 0x73, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x07, 0x70,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x22, 0x28, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x3b, 0x0a, 0x07,
	0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x47, 0x69,
	0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x50, 0x65, 0x72, 0x47, 0x69, 0x62, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x48, 0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x50, 0x65, 0x72, 0x48, 0x6f, 0x75, 0x72, 0x22, 0x7b, 0x0a, 0x0d, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_session_proto_rawDescOnce sync.Once
	file_pb_session_proto_rawDescData = file_pb_session_proto_rawDesc
)

func file_pb_session_proto_rawDescGZIP() []byte {
	file_pb_session_proto_rawDescOnce.Do(func() {
		file_pb_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_session_proto_rawDescData)
	})
	return file_pb_session_proto_rawDescData
}

var file_pb_session_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_session_proto_goTypes = []interface{}{
	(*SessionRequest)(nil),  // 0: pb.SessionRequest
	(*SessionResponse)(nil), // 1: pb.SessionResponse
	(*SessionInfo)(nil),     // 2: pb.SessionInfo
	(*ConsumerInfo)(nil),    // 3: pb.ConsumerInfo
	(*LocationInfo)(nil),    // 4: pb.LocationInfo
	(*Pricing)(nil),         // 5: pb.Pricing
	(*SessionStatus)(nil),   // 6: pb.SessionStatus
}
var file_pb_session_proto_depIdxs = []int32{
	3, // 0: pb.SessionRequest.consumer:type_name -> pb.ConsumerInfo
	4, // 1: pb.ConsumerInfo.location:type_name -> pb.LocationInfo
	5, // 2: pb.ConsumerInfo.pricing:type_name -> pb.Pricing
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_session_proto_init() }
func file_pb_session_proto_init() {
	if File_pb_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionRequest); i {
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
		file_pb_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionResponse); i {
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
		file_pb_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionInfo); i {
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
		file_pb_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerInfo); i {
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
		file_pb_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationInfo); i {
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
		file_pb_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pricing); i {
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
		file_pb_session_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionStatus); i {
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
			RawDescriptor: file_pb_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_session_proto_goTypes,
		DependencyIndexes: file_pb_session_proto_depIdxs,
		MessageInfos:      file_pb_session_proto_msgTypes,
	}.Build()
	File_pb_session_proto = out.File
	file_pb_session_proto_rawDesc = nil
	file_pb_session_proto_goTypes = nil
	file_pb_session_proto_depIdxs = nil
}