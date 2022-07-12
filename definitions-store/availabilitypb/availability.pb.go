// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: availability.proto

package availabilitypb

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

type AvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *Location `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
}

func (x *AvailabilityRequest) Reset() {
	*x = AvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_availability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailabilityRequest) ProtoMessage() {}

func (x *AvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_availability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailabilityRequest.ProtoReflect.Descriptor instead.
func (*AvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_availability_proto_rawDescGZIP(), []int{0}
}

func (x *AvailabilityRequest) GetSource() *Location {
	if x != nil {
		return x.Source
	}
	return nil
}

type AvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarType  string `protobuf:"bytes,1,opt,name=CarType,proto3" json:"CarType,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
	Distance int32  `protobuf:"varint,3,opt,name=Distance,proto3" json:"Distance,omitempty"`
}

func (x *AvailabilityResponse) Reset() {
	*x = AvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_availability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailabilityResponse) ProtoMessage() {}

func (x *AvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_availability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailabilityResponse.ProtoReflect.Descriptor instead.
func (*AvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_availability_proto_rawDescGZIP(), []int{1}
}

func (x *AvailabilityResponse) GetCarType() string {
	if x != nil {
		return x.CarType
	}
	return ""
}

func (x *AvailabilityResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *AvailabilityResponse) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
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
		mi := &file_availability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_availability_proto_msgTypes[2]
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
	return file_availability_proto_rawDescGZIP(), []int{2}
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

var File_availability_proto protoreflect.FileDescriptor

var file_availability_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x70, 0x62, 0x22, 0x47, 0x0a, 0x13, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x68, 0x0a,
	0x14, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x44,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x44,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x32, 0x77, 0x0a,
	0x13, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x23, 0x2e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x74, 0x69, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x61, 0x6d,
	0x2f, 0x63, 0x61, 0x72, 0x2d, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2d, 0x70, 0x6f, 0x63, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_availability_proto_rawDescOnce sync.Once
	file_availability_proto_rawDescData = file_availability_proto_rawDesc
)

func file_availability_proto_rawDescGZIP() []byte {
	file_availability_proto_rawDescOnce.Do(func() {
		file_availability_proto_rawDescData = protoimpl.X.CompressGZIP(file_availability_proto_rawDescData)
	})
	return file_availability_proto_rawDescData
}

var file_availability_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_availability_proto_goTypes = []interface{}{
	(*AvailabilityRequest)(nil),  // 0: availabilitypb.AvailabilityRequest
	(*AvailabilityResponse)(nil), // 1: availabilitypb.AvailabilityResponse
	(*Location)(nil),             // 2: availabilitypb.Location
}
var file_availability_proto_depIdxs = []int32{
	2, // 0: availabilitypb.AvailabilityRequest.Source:type_name -> availabilitypb.Location
	0, // 1: availabilitypb.AvailabilityService.GetAvailability:input_type -> availabilitypb.AvailabilityRequest
	1, // 2: availabilitypb.AvailabilityService.GetAvailability:output_type -> availabilitypb.AvailabilityResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_availability_proto_init() }
func file_availability_proto_init() {
	if File_availability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_availability_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailabilityRequest); i {
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
		file_availability_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailabilityResponse); i {
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
		file_availability_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_availability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_availability_proto_goTypes,
		DependencyIndexes: file_availability_proto_depIdxs,
		MessageInfos:      file_availability_proto_msgTypes,
	}.Build()
	File_availability_proto = out.File
	file_availability_proto_rawDesc = nil
	file_availability_proto_goTypes = nil
	file_availability_proto_depIdxs = nil
}
