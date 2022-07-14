// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: routing.proto

package routingpb

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

type RoutingAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *Location `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
}

func (x *RoutingAvailabilityRequest) Reset() {
	*x = RoutingAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingAvailabilityRequest) ProtoMessage() {}

func (x *RoutingAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*RoutingAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{0}
}

func (x *RoutingAvailabilityRequest) GetSource() *Location {
	if x != nil {
		return x.Source
	}
	return nil
}

type RoutingAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarType  string `protobuf:"bytes,1,opt,name=CarType,proto3" json:"CarType,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
	Distance int32  `protobuf:"varint,3,opt,name=Distance,proto3" json:"Distance,omitempty"`
}

func (x *RoutingAvailabilityResponse) Reset() {
	*x = RoutingAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingAvailabilityResponse) ProtoMessage() {}

func (x *RoutingAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*RoutingAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{1}
}

func (x *RoutingAvailabilityResponse) GetCarType() string {
	if x != nil {
		return x.CarType
	}
	return ""
}

func (x *RoutingAvailabilityResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *RoutingAvailabilityResponse) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type ListRoutingAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListRoutingAvailabilityResponse []*RoutingAvailabilityResponse `protobuf:"bytes,1,rep,name=listRoutingAvailabilityResponse,proto3" json:"listRoutingAvailabilityResponse,omitempty"`
}

func (x *ListRoutingAvailabilityResponse) Reset() {
	*x = ListRoutingAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoutingAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutingAvailabilityResponse) ProtoMessage() {}

func (x *ListRoutingAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutingAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*ListRoutingAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{2}
}

func (x *ListRoutingAvailabilityResponse) GetListRoutingAvailabilityResponse() []*RoutingAvailabilityResponse {
	if x != nil {
		return x.ListRoutingAvailabilityResponse
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  int32 `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{3}
}

func (x *Location) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type RoutingRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Date        string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *RoutingRateRequest) Reset() {
	*x = RoutingRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingRateRequest) ProtoMessage() {}

func (x *RoutingRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingRateRequest.ProtoReflect.Descriptor instead.
func (*RoutingRateRequest) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{4}
}

func (x *RoutingRateRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *RoutingRateRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *RoutingRateRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type RoutingRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarType  string  `protobuf:"bytes,1,opt,name=CarType,proto3" json:"CarType,omitempty"`
	Location string  `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
	Distance int32   `protobuf:"varint,3,opt,name=Distance,proto3" json:"Distance,omitempty"`
	Price    float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *RoutingRateResponse) Reset() {
	*x = RoutingRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingRateResponse) ProtoMessage() {}

func (x *RoutingRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingRateResponse.ProtoReflect.Descriptor instead.
func (*RoutingRateResponse) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{5}
}

func (x *RoutingRateResponse) GetCarType() string {
	if x != nil {
		return x.CarType
	}
	return ""
}

func (x *RoutingRateResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *RoutingRateResponse) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *RoutingRateResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_routing_proto protoreflect.FileDescriptor

var file_routing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x22, 0x49, 0x0a, 0x1a, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x6f, 0x0a, 0x1b, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x44, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x1f, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x1f, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0x0a, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x22, 0x62, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x7d, 0x0a, 0x13, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x43, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0xca, 0x01, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x25, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x69, 0x6e, 0x65, 0x73, 0x68, 0x37, 0x38, 0x39, 0x6b, 0x75, 0x6d, 0x61, 0x72, 0x31,
	0x32, 0x2f, 0x43, 0x61, 0x72, 0x41, 0x70, 0x70, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routing_proto_rawDescOnce sync.Once
	file_routing_proto_rawDescData = file_routing_proto_rawDesc
)

func file_routing_proto_rawDescGZIP() []byte {
	file_routing_proto_rawDescOnce.Do(func() {
		file_routing_proto_rawDescData = protoimpl.X.CompressGZIP(file_routing_proto_rawDescData)
	})
	return file_routing_proto_rawDescData
}

var file_routing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_routing_proto_goTypes = []interface{}{
	(*RoutingAvailabilityRequest)(nil),      // 0: routingpb.RoutingAvailabilityRequest
	(*RoutingAvailabilityResponse)(nil),     // 1: routingpb.RoutingAvailabilityResponse
	(*ListRoutingAvailabilityResponse)(nil), // 2: routingpb.ListRoutingAvailabilityResponse
	(*Location)(nil),                        // 3: routingpb.Location
	(*RoutingRateRequest)(nil),              // 4: routingpb.RoutingRateRequest
	(*RoutingRateResponse)(nil),             // 5: routingpb.RoutingRateResponse
}
var file_routing_proto_depIdxs = []int32{
	3, // 0: routingpb.RoutingAvailabilityRequest.Source:type_name -> routingpb.Location
	1, // 1: routingpb.ListRoutingAvailabilityResponse.listRoutingAvailabilityResponse:type_name -> routingpb.RoutingAvailabilityResponse
	0, // 2: routingpb.RoutingService.GetAvailability:input_type -> routingpb.RoutingAvailabilityRequest
	4, // 3: routingpb.RoutingService.GetRate:input_type -> routingpb.RoutingRateRequest
	2, // 4: routingpb.RoutingService.GetAvailability:output_type -> routingpb.ListRoutingAvailabilityResponse
	5, // 5: routingpb.RoutingService.GetRate:output_type -> routingpb.RoutingRateResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_routing_proto_init() }
func file_routing_proto_init() {
	if File_routing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingAvailabilityRequest); i {
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
		file_routing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingAvailabilityResponse); i {
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
		file_routing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoutingAvailabilityResponse); i {
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
		file_routing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_routing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingRateRequest); i {
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
		file_routing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingRateResponse); i {
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
			RawDescriptor: file_routing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routing_proto_goTypes,
		DependencyIndexes: file_routing_proto_depIdxs,
		MessageInfos:      file_routing_proto_msgTypes,
	}.Build()
	File_routing_proto = out.File
	file_routing_proto_rawDesc = nil
	file_routing_proto_goTypes = nil
	file_routing_proto_depIdxs = nil
}
