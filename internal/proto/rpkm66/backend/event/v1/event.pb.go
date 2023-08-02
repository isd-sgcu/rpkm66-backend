// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpkm66/backend/event/v1/event.proto

package v1

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NameTH        string `protobuf:"bytes,2,opt,name=nameTH,proto3" json:"nameTH,omitempty"`
	DescriptionTH string `protobuf:"bytes,3,opt,name=descriptionTH,proto3" json:"descriptionTH,omitempty"`
	NameEN        string `protobuf:"bytes,4,opt,name=nameEN,proto3" json:"nameEN,omitempty"`
	DescriptionEN string `protobuf:"bytes,5,opt,name=descriptionEN,proto3" json:"descriptionEN,omitempty"`
	Code          string `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`
	ImageURL      string `protobuf:"bytes,7,opt,name=imageURL,proto3" json:"imageURL,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[0]
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
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetNameTH() string {
	if x != nil {
		return x.NameTH
	}
	return ""
}

func (x *Event) GetDescriptionTH() string {
	if x != nil {
		return x.DescriptionTH
	}
	return ""
}

func (x *Event) GetNameEN() string {
	if x != nil {
		return x.NameEN
	}
	return ""
}

func (x *Event) GetDescriptionEN() string {
	if x != nil {
		return x.DescriptionEN
	}
	return ""
}

func (x *Event) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Event) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

type FindAllEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindAllEventRequest) Reset() {
	*x = FindAllEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEventRequest) ProtoMessage() {}

func (x *FindAllEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEventRequest.ProtoReflect.Descriptor instead.
func (*FindAllEventRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindAllEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event []*Event `protobuf:"bytes,1,rep,name=event,proto3" json:"event,omitempty"`
}

func (x *FindAllEventResponse) Reset() {
	*x = FindAllEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEventResponse) ProtoMessage() {}

func (x *FindAllEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEventResponse.ProtoReflect.Descriptor instead.
func (*FindAllEventResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllEventResponse) GetEvent() []*Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type FindEventByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindEventByIDRequest) Reset() {
	*x = FindEventByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEventByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEventByIDRequest) ProtoMessage() {}

func (x *FindEventByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEventByIDRequest.ProtoReflect.Descriptor instead.
func (*FindEventByIDRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{3}
}

func (x *FindEventByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindEventByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *FindEventByIDResponse) Reset() {
	*x = FindEventByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEventByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEventByIDResponse) ProtoMessage() {}

func (x *FindEventByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEventByIDResponse.ProtoReflect.Descriptor instead.
func (*FindEventByIDResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{4}
}

func (x *FindEventByIDResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{5}
}

func (x *CreateEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{6}
}

func (x *CreateEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type UpdateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *UpdateEventRequest) Reset() {
	*x = UpdateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventRequest) ProtoMessage() {}

func (x *UpdateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventRequest.ProtoReflect.Descriptor instead.
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type UpdateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *UpdateEventResponse) Reset() {
	*x = UpdateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventResponse) ProtoMessage() {}

func (x *UpdateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventResponse.ProtoReflect.Descriptor instead.
func (*UpdateEventResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type DeleteEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEventRequest) Reset() {
	*x = DeleteEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRequest) ProtoMessage() {}

func (x *DeleteEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteEventResponse) Reset() {
	*x = DeleteEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventResponse) ProtoMessage() {}

func (x *DeleteEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventResponse.ProtoReflect.Descriptor instead.
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteEventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type FindAllEventWithTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType string `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
}

func (x *FindAllEventWithTypeRequest) Reset() {
	*x = FindAllEventWithTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEventWithTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEventWithTypeRequest) ProtoMessage() {}

func (x *FindAllEventWithTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEventWithTypeRequest.ProtoReflect.Descriptor instead.
func (*FindAllEventWithTypeRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{11}
}

func (x *FindAllEventWithTypeRequest) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

type FindAllEventWithTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event []*Event `protobuf:"bytes,1,rep,name=event,proto3" json:"event,omitempty"`
}

func (x *FindAllEventWithTypeResponse) Reset() {
	*x = FindAllEventWithTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEventWithTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEventWithTypeResponse) ProtoMessage() {}

func (x *FindAllEventWithTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_backend_event_v1_event_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEventWithTypeResponse.ProtoReflect.Descriptor instead.
func (*FindAllEventWithTypeResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_backend_event_v1_event_proto_rawDescGZIP(), []int{12}
}

func (x *FindAllEventWithTypeResponse) GetEvent() []*Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_rpkm66_backend_event_v1_event_proto protoreflect.FileDescriptor

var file_rpkm66_backend_event_v1_event_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xc3,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x6d, 0x65,
	0x54, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x48,
	0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x48, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x48, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x4e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x4e, 0x12, 0x24,
	0x0a, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x4e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x52, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x52, 0x4c, 0x22, 0x25, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x14, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4d, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d,
	0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x4a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70,
	0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3b, 0x0a, 0x1b, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x54, 0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0xac, 0x05, 0x0a,
	0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a,
	0x0c, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x2e,
	0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x72, 0x70,
	0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0d,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2d, 0x2e,
	0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x72,
	0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36,
	0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x2b, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72,
	0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x2e, 0x72,
	0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x72,
	0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpkm66_backend_event_v1_event_proto_rawDescOnce sync.Once
	file_rpkm66_backend_event_v1_event_proto_rawDescData = file_rpkm66_backend_event_v1_event_proto_rawDesc
)

func file_rpkm66_backend_event_v1_event_proto_rawDescGZIP() []byte {
	file_rpkm66_backend_event_v1_event_proto_rawDescOnce.Do(func() {
		file_rpkm66_backend_event_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpkm66_backend_event_v1_event_proto_rawDescData)
	})
	return file_rpkm66_backend_event_v1_event_proto_rawDescData
}

var file_rpkm66_backend_event_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_rpkm66_backend_event_v1_event_proto_goTypes = []interface{}{
	(*Event)(nil),                        // 0: rpkm66.backend.event.v1.Event
	(*FindAllEventRequest)(nil),          // 1: rpkm66.backend.event.v1.FindAllEventRequest
	(*FindAllEventResponse)(nil),         // 2: rpkm66.backend.event.v1.FindAllEventResponse
	(*FindEventByIDRequest)(nil),         // 3: rpkm66.backend.event.v1.FindEventByIDRequest
	(*FindEventByIDResponse)(nil),        // 4: rpkm66.backend.event.v1.FindEventByIDResponse
	(*CreateEventRequest)(nil),           // 5: rpkm66.backend.event.v1.CreateEventRequest
	(*CreateEventResponse)(nil),          // 6: rpkm66.backend.event.v1.CreateEventResponse
	(*UpdateEventRequest)(nil),           // 7: rpkm66.backend.event.v1.UpdateEventRequest
	(*UpdateEventResponse)(nil),          // 8: rpkm66.backend.event.v1.UpdateEventResponse
	(*DeleteEventRequest)(nil),           // 9: rpkm66.backend.event.v1.DeleteEventRequest
	(*DeleteEventResponse)(nil),          // 10: rpkm66.backend.event.v1.DeleteEventResponse
	(*FindAllEventWithTypeRequest)(nil),  // 11: rpkm66.backend.event.v1.FindAllEventWithTypeRequest
	(*FindAllEventWithTypeResponse)(nil), // 12: rpkm66.backend.event.v1.FindAllEventWithTypeResponse
}
var file_rpkm66_backend_event_v1_event_proto_depIdxs = []int32{
	0,  // 0: rpkm66.backend.event.v1.FindAllEventResponse.event:type_name -> rpkm66.backend.event.v1.Event
	0,  // 1: rpkm66.backend.event.v1.FindEventByIDResponse.event:type_name -> rpkm66.backend.event.v1.Event
	0,  // 2: rpkm66.backend.event.v1.CreateEventRequest.event:type_name -> rpkm66.backend.event.v1.Event
	0,  // 3: rpkm66.backend.event.v1.CreateEventResponse.event:type_name -> rpkm66.backend.event.v1.Event
	0,  // 4: rpkm66.backend.event.v1.UpdateEventRequest.event:type_name -> rpkm66.backend.event.v1.Event
	0,  // 5: rpkm66.backend.event.v1.UpdateEventResponse.event:type_name -> rpkm66.backend.event.v1.Event
	0,  // 6: rpkm66.backend.event.v1.FindAllEventWithTypeResponse.event:type_name -> rpkm66.backend.event.v1.Event
	1,  // 7: rpkm66.backend.event.v1.EventService.FindAllEvent:input_type -> rpkm66.backend.event.v1.FindAllEventRequest
	3,  // 8: rpkm66.backend.event.v1.EventService.FindEventByID:input_type -> rpkm66.backend.event.v1.FindEventByIDRequest
	5,  // 9: rpkm66.backend.event.v1.EventService.Create:input_type -> rpkm66.backend.event.v1.CreateEventRequest
	7,  // 10: rpkm66.backend.event.v1.EventService.Update:input_type -> rpkm66.backend.event.v1.UpdateEventRequest
	9,  // 11: rpkm66.backend.event.v1.EventService.Delete:input_type -> rpkm66.backend.event.v1.DeleteEventRequest
	11, // 12: rpkm66.backend.event.v1.EventService.FindAllEventWithType:input_type -> rpkm66.backend.event.v1.FindAllEventWithTypeRequest
	2,  // 13: rpkm66.backend.event.v1.EventService.FindAllEvent:output_type -> rpkm66.backend.event.v1.FindAllEventResponse
	4,  // 14: rpkm66.backend.event.v1.EventService.FindEventByID:output_type -> rpkm66.backend.event.v1.FindEventByIDResponse
	6,  // 15: rpkm66.backend.event.v1.EventService.Create:output_type -> rpkm66.backend.event.v1.CreateEventResponse
	8,  // 16: rpkm66.backend.event.v1.EventService.Update:output_type -> rpkm66.backend.event.v1.UpdateEventResponse
	10, // 17: rpkm66.backend.event.v1.EventService.Delete:output_type -> rpkm66.backend.event.v1.DeleteEventResponse
	12, // 18: rpkm66.backend.event.v1.EventService.FindAllEventWithType:output_type -> rpkm66.backend.event.v1.FindAllEventWithTypeResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_rpkm66_backend_event_v1_event_proto_init() }
func file_rpkm66_backend_event_v1_event_proto_init() {
	if File_rpkm66_backend_event_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpkm66_backend_event_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEventRequest); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEventResponse); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEventByIDRequest); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEventByIDResponse); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventRequest); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventResponse); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventRequest); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventResponse); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEventWithTypeRequest); i {
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
		file_rpkm66_backend_event_v1_event_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEventWithTypeResponse); i {
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
			RawDescriptor: file_rpkm66_backend_event_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpkm66_backend_event_v1_event_proto_goTypes,
		DependencyIndexes: file_rpkm66_backend_event_v1_event_proto_depIdxs,
		MessageInfos:      file_rpkm66_backend_event_v1_event_proto_msgTypes,
	}.Build()
	File_rpkm66_backend_event_v1_event_proto = out.File
	file_rpkm66_backend_event_v1_event_proto_rawDesc = nil
	file_rpkm66_backend_event_v1_event_proto_goTypes = nil
	file_rpkm66_backend_event_v1_event_proto_depIdxs = nil
}
